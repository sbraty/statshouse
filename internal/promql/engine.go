// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package promql

import (
	"context"
	"errors"
	"fmt"
	"hash"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gogo/protobuf/sortkeys"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/vkcom/statshouse-go"
	"github.com/vkcom/statshouse/internal/data_model"
	"github.com/vkcom/statshouse/internal/format"
	"github.com/vkcom/statshouse/internal/promql/parser"
	"github.com/vkcom/statshouse/internal/vkgo/srvfunc"
	"golang.org/x/sync/errgroup"
	"pgregory.net/rand"
)

const (
	LabelWhat    = "__what__"
	labelBy      = "__by__"
	labelBind    = "__bind__"
	LabelShard   = "__shard__"
	LabelOffset  = "__offset__"
	LabelMinHost = "__minhost__"
	LabelMaxHost = "__maxhost__"

	maxSeriesRows = 10_000_000
)

type Query struct {
	Start int64 // inclusive
	End   int64 // exclusive
	Step  int64
	Expr  string

	Options Options // StatsHouse specific
}

// NB! If you add an option make sure that default Options{} corresponds to Prometheus behavior.
type Options struct {
	Version          string
	Namespace        string
	AvoidCache       bool
	TimeNow          int64
	ScreenWidth      int64
	Collapse         bool // aka "point" query
	Extend           bool
	TagWhat          bool
	TagOffset        bool
	TagTotal         bool
	ExplicitGrouping bool
	MinHost          bool
	MaxHost          bool
	QuerySequential  bool
	Offsets          []int64
	Limit            int
	Rand             *rand.Rand
	Vars             map[string]Variable

	ExprQueriesSingleMetricCallback MetricMetaValueCallback
}

type (
	MetricMetaValueCallback func(*format.MetricMetaValue)
	SeriesQueryCallback     func(version string, key string, pq any, lod any, avoidCache bool)
)

type Variable struct {
	Value  []string
	Group  bool
	Negate bool
}

type Engine struct {
	h  Handler
	tz data_model.TimeZone
}

type evaluator struct {
	Engine

	ctx context.Context
	opt Options
	ast parser.Expr
	ars map[parser.Expr]parser.Expr // ast reductions
	t   data_model.Timescale
	r   int64 // matrix selector range
	hh  hash.Hash64

	// metric -> tag index -> offset -> tag value id -> tag value
	tags map[*format.MetricMetaValue][]map[int64]map[int32]string
	// metric -> offset -> String TOP
	stop map[*format.MetricMetaValue]map[int64][]string

	// memory management
	allocMap         map[*[]float64]bool
	freeList         []*[]float64
	reuseList        []*[]float64
	cancellationList []func()

	// diagnostics
	trace              *[]string
	debug              bool
	timeStart          time.Time
	timeQueryParseEnd  time.Time
	dataAccessDuration time.Duration
}

type seriesQueryX struct { // SeriesQuery eXtended
	SeriesQuery
	prefixSum bool
	histogram histogramQuery
}

type histogramQuery struct {
	restore bool
	filter  bool
	eq      bool // compare "==" if true, "!=" otherwise
	le      float32
}

type contextKey int

type traceContext struct {
	s *[]string // sink
	v bool      // verbose
}

const (
	offsetContextKey contextKey = iota
	traceContextKey
)

func NewEngine(h Handler, tz data_model.TimeZone) Engine {
	return Engine{h, tz}
}

func (ng Engine) Exec(ctx context.Context, qry Query) (parser.Value, func(), error) {
	// parse query
	ev, err := ng.newEvaluator(ctx, qry)
	if err != nil {
		return nil, nil, Error{what: err}
	}
	if ev.t.Empty() {
		return &TimeSeries{Time: []int64{}}, func() {}, nil
	}
	if e, ok := ev.ast.(*parser.StringLiteral); ok {
		return String{T: qry.Start, V: e.Val}, func() {}, nil
	}
	// evaluate query
	if ev.trace != nil {
		*ev.trace = append(*ev.trace, ev.ast.String())
		if ev.debug {
			ev.tracef("requested from %d to %d, timescale from %d to %d", qry.Start, qry.End, ev.t.Time[ev.t.StartX], ev.t.Time[len(ev.t.Time)-1])
		}
	}
	var ok bool
	defer func() {
		if !ok {
			ev.cancel()
		}
	}()
	res, err := ev.exec()
	if err != nil {
		return nil, nil, Error{what: err}
	}
	// resolve int32 tag values into strings
	for _, dat := range res.Series.Data {
		for _, tg := range dat.Tags.ID2Tag {
			tg.stringify(&ev)
		}
	}
	if ev.trace != nil {
		ev.tracef("buffers alloc #%d, reuse #%d, %s", len(ev.allocMap)+len(ev.freeList), len(ev.reuseList), res.String())
	}
	ev.reportStat(qry, time.Now())
	ok = true // prevents deffered "cancel"
	return &res, ev.cancel, nil
}

func (ng Engine) newEvaluator(ctx context.Context, qry Query) (evaluator, error) {
	timeStart := time.Now()
	if qry.Options.TimeNow == 0 {
		// fix the time "now"
		qry.Options.TimeNow = timeStart.Unix()
	}
	if qry.Options.TimeNow < qry.Start {
		// the future is unknown
		return evaluator{}, nil
	}
	ev := evaluator{
		Engine:    ng,
		ctx:       ctx,
		opt:       qry.Options,
		timeStart: timeStart,
	}
	// init diagnostics
	if v, ok := ctx.Value(traceContextKey).(*traceContext); ok {
		ev.trace = v.s
		ev.debug = v.v
	}
	// parse PromQL expression
	var err error
	ev.ast, err = parser.ParseExpr(qry.Expr)
	if err != nil {
		return evaluator{}, err
	}
	if v, ok := evalLiteral(ev.ast); ok {
		ev.ast = v
	}
	ev.opt.Offsets = normalizeOffsets(append(ev.opt.Offsets, 0))
	// match metrics
	var (
		maxRange     int64
		metricOffset = make(map[*format.MetricMetaValue]int64)
	)
	parser.Inspect(ev.ast, func(node parser.Node, path []parser.Node) error {
		switch e := node.(type) {
		case *parser.VectorSelector:
			if len(e.OriginalOffsetEx) != 0 {
				e.Offsets = normalizeOffsets(e.OriginalOffsetEx)
			} else {
				e.Offsets = []int64{e.OriginalOffset}
			}
			if err = ev.bindVariables(e); err == nil {
				err = ev.matchMetrics(e, path, metricOffset, ev.opt.Offsets[0])
			}
		case *parser.MatrixSelector:
			if maxRange < e.Range {
				maxRange = e.Range
			}
		case *parser.SubqueryExpr:
			if maxRange < e.Range {
				maxRange = e.Range
			}
		}
		return err
	})
	if err != nil {
		return evaluator{}, err
	}
	// init timescale
	qry.Start -= maxRange // widen time range to accommodate range selectors
	ev.t, err = ng.tz.GetTimescale(data_model.GetTimescaleArgs{
		Version:      qry.Options.Version,
		Start:        qry.Start,
		End:          qry.End,
		Step:         qry.Step,
		TimeNow:      qry.Options.TimeNow,
		ScreenWidth:  qry.Options.ScreenWidth,
		Collapse:     qry.Options.Collapse,
		Extend:       qry.Options.Extend,
		MetricOffset: metricOffset,
	})

	if err != nil || ev.t.Empty() {
		return evaluator{}, err
	}
	// evaluate reduction rules
	ev.ars = make(map[parser.Expr]parser.Expr)
	stepMin := ev.t.LODs[len(ev.t.LODs)-1].Step
	parser.Inspect(ev.ast, func(node parser.Node, nodes []parser.Node) error {
		switch s := node.(type) {
		case *parser.VectorSelector:
			var grouped bool
			if s.GroupBy != nil {
				grouped = true
			}
			if !grouped && len(s.What) <= 1 {
				if ar, ok := evalReductionRules(s, nodes, stepMin); ok {
					s.What = ar.what
					s.GroupBy = ar.groupBy
					s.GroupWithout = ar.groupWithout
					s.Range = ar.step
					s.OmitNameTag = true
					ev.ars[ar.expr] = s
					grouped = ar.grouped
				}
			}
			if !grouped && !ev.opt.ExplicitGrouping {
				s.GroupByAll = true
			}
		}
		return nil
	})
	// callback
	if ev.opt.ExprQueriesSingleMetricCallback != nil && len(metricOffset) == 1 {
		for metric := range metricOffset {
			ev.opt.ExprQueriesSingleMetricCallback(metric)
			break
		}
	}
	// final touch
	ev.tags = make(map[*format.MetricMetaValue][]map[int64]map[int32]string)
	ev.stop = make(map[*format.MetricMetaValue]map[int64][]string)
	ev.timeQueryParseEnd = time.Now()
	return ev, nil
}

func (ev *evaluator) bindVariables(sel *parser.VectorSelector) error {
	var s []*labels.Matcher
	for _, matcher := range sel.LabelMatchers {
		if matcher.Name == labelBind {
			if matcher.Type != labels.MatchEqual {
				return fmt.Errorf("%s supports only strict equality", matcher.Name)
			}
			s = append(s, matcher)
		}
	}
	for _, matcher := range s {
		for _, bind := range strings.Split(matcher.Value, ",") {
			s := strings.Split(bind, ":")
			if len(s) != 2 || len(s[0]) == 0 || len(s[1]) == 0 {
				return fmt.Errorf("%s invalid value format: expected \"tag:var\", got %q", matcher.Name, bind)
			}
			var (
				vn = s[1]   // variable name
				vv Variable // variable value
				ok bool
			)
			if vv, ok = ev.opt.Vars[vn]; !ok {
				if ev.trace != nil && ev.debug {
					ev.tracef("variable %q not specified", vn)
				}
				continue
			}
			var mt labels.MatchType
			if vv.Negate {
				mt = labels.MatchNotEqual
			} else {
				mt = labels.MatchEqual
			}
			var (
				tn  = s[0] // tag name
				m   *labels.Matcher
				err error
			)
			for _, v := range vv.Value {
				if m, err = labels.NewMatcher(mt, tn, v); err != nil {
					return err
				}
				sel.LabelMatchers = append(sel.LabelMatchers, m)
			}
			if vv.Group {
				sel.GroupBy = append(sel.GroupBy, tn)
			}
		}
	}
	return nil
}

func (ev *evaluator) matchMetrics(sel *parser.VectorSelector, path []parser.Node, metricOffset map[*format.MetricMetaValue]int64, offset int64) error {
	sel.MinHost = ev.opt.MinHost
	sel.MaxHost = ev.opt.MaxHost
	for _, matcher := range sel.LabelMatchers {
		switch matcher.Name {
		case labels.MetricName:
			metrics, names, err := ev.h.MatchMetrics(ev.ctx, matcher, ev.opt.Namespace)
			if err != nil {
				return err
			}
			if len(metrics) == 0 {
				if ev.trace != nil && ev.debug {
					ev.tracef("no metric matches %v", matcher)
				}
				return nil // metric does not exist, not an error
			}
			if ev.trace != nil && ev.debug {
				ev.tracef("found %d metrics for %v", len(metrics), matcher)
			}
			for i, m := range metrics {
				var selOffset int64
				for _, v := range sel.Offsets {
					if selOffset < v {
						selOffset = v
					}
				}
				var (
					curOffset, ok = metricOffset[m]
					newOffset     = selOffset + offset
				)
				if !ok || curOffset < newOffset {
					metricOffset[m] = newOffset
				}
				sel.MatchingMetrics = append(sel.MatchingMetrics, m)
				sel.MatchingNames = append(sel.MatchingNames, names[i])
			}
		case LabelWhat:
			if matcher.Type != labels.MatchEqual {
				return fmt.Errorf("%s supports only strict equality", LabelWhat)
			}
			for _, what := range strings.Split(matcher.Value, ",") {
				switch what {
				case "":
					// ignore empty "what" value
				case MinHost:
					sel.MinHost = true
				case MaxHost:
					sel.MaxHost = true
				default:
					sel.Whats = append(sel.Whats, what)
				}
			}
		case labelBy:
			if matcher.Type != labels.MatchEqual {
				return fmt.Errorf("%s supports only strict equality", labelBy)
			}
			if len(matcher.Value) != 0 {
				sel.GroupBy = append(sel.GroupBy, strings.Split(matcher.Value, ",")...)
			} else if sel.GroupBy == nil {
				sel.GroupBy = make([]string, 0)
			}
		case LabelMinHost:
			sel.MinHost = true
			sel.MinHostMatchers = append(sel.MinHostMatchers, matcher)
		case LabelMaxHost:
			sel.MaxHost = true
			sel.MaxHostMatchers = append(sel.MaxHostMatchers, matcher)
		}
	}
	for i := len(path); len(sel.MetricKindHint) == 0 && i != 0; i-- {
		switch e := path[i-1].(type) {
		case *parser.Call:
			switch e.Func.Name {
			case "delta", "deriv", "holt_winters", "idelta", "predict_linear":
				sel.MetricKindHint = format.MetricKindValue
			case "increase", "irate", "rate", "resets":
				sel.MetricKindHint = format.MetricKindCounter
			}
		}
	}
	for i := len(path); !(sel.MinHost && sel.MaxHost) && i != 0; i-- {
		if e, ok := path[i-1].(*parser.Call); ok {
			switch e.Func.Name {
			case "label_minhost":
				sel.MinHost = true
			case "label_maxhost":
				sel.MaxHost = true
			}
		}
	}
	for i := len(path); i != 0; i-- {
		if e, ok := path[i-1].(*parser.Call); ok && e.Func.Name == "histogram_quantile" {
			sel.GroupBy = append(sel.GroupBy, format.LETagName)
			break
		}
	}
	return nil
}

func (ev *evaluator) time() []int64 {
	return ev.t.Time
}

func (ev *evaluator) exec() (TimeSeries, error) {
	srs, err := ev.eval(ev.ast)
	if err != nil {
		return TimeSeries{}, err
	}
	ev.stableRemoveEmptySeries(srs)
	res := TimeSeries{Time: ev.t.Time[ev.t.StartX:]}
	limit := ev.opt.Limit
	if limit == 0 {
		limit = math.MaxInt
	} else if limit < 0 {
		limit = -limit
	}
	for _, sr := range srs {
		if len(res.Series.Data) == 0 {
			// get resulting meta from first time shift
			res.Series.Meta = sr.Meta
		}
		i := 0
		end := len(sr.Data)
		if limit < len(sr.Data) {
			if ev.opt.Limit < 0 {
				i = len(sr.Data) - limit
			} else {
				end = limit
			}
		}
		for ; i < end; i++ {
			// trim time outside [StartX:]
			vs := (*sr.Data[i].Values)[ev.t.StartX:]
			sr.Data[i].Values = &vs
			for j := range sr.Data[i].MinMaxHost {
				if len(sr.Data[i].MinMaxHost[j]) != 0 {
					sr.Data[i].MinMaxHost[j] = sr.Data[i].MinMaxHost[j][ev.t.StartX:]
				}
			}
			res.Series.appendSome(sr, i)
		}
	}
	return res, nil
}

func (ev *evaluator) eval(expr parser.Expr) (res []Series, err error) {
	if ev.ctx.Err() != nil {
		return nil, ev.ctx.Err()
	}
	if e, ok := ev.ars[expr]; ok {
		if ev.trace != nil && ev.debug {
			ev.tracef("replace %s with %s", string(expr.Type()), string(e.Type()))
		}
		return ev.eval(e)
	}
	switch e := expr.(type) {
	case *parser.AggregateExpr:
		fn := aggregates[e.Op]
		if fn == nil {
			return nil, fmt.Errorf("not implemented aggregate %q", e.Op)
		}
		if res, err = fn(ev, e); err != nil {
			return nil, err
		}
		if e.Op != parser.TOPK && e.Op != parser.BOTTOMK {
			removeMetricName(res)
		}
	case *parser.BinaryExpr:
		switch l := e.LHS.(type) {
		case *parser.NumberLiteral:
			res, err = ev.eval(e.RHS)
			if err != nil {
				return nil, err
			}
			fn := getBinaryFunc(scalarSliceFuncM, e.Op, e.ReturnBool)
			if fn == nil {
				return nil, fmt.Errorf("binary operator %q is not defined on (%q, %q) pair", e.Op, e.LHS.Type(), e.RHS.Type())
			}
			for i := range res {
				for j := range res[i].Data {
					fn(l.Val, *res[i].Data[j].Values)
					res[i].Data[j].What = SelectorWhat{}
				}
				res[i].Meta = evalSeriesMeta(e, SeriesMeta{}, res[i].Meta)
			}
		default:
			switch r := e.RHS.(type) {
			case *parser.NumberLiteral:
				if res, err = ev.eval(e.LHS); err != nil {
					return nil, err
				}
				fn := getBinaryFunc(sliceScalarFuncM, e.Op, e.ReturnBool)
				if fn == nil {
					return nil, fmt.Errorf("binary operator %q is not defined on (%q, %q) pair", e.Op, e.LHS.Type(), e.RHS.Type())
				}
				for i := range res {
					if e.Op == parser.LDEFAULT && res[i].empty() {
						res[i] = ev.newVector(r.Val)
					} else {
						for j := range res[i].Data {
							fn(*res[i].Data[j].Values, r.Val)
							res[i].Data[j].What = SelectorWhat{}
						}
						res[i].Meta = evalSeriesMeta(e, res[i].Meta, SeriesMeta{})
					}
				}
			default:
				if res, err = ev.evalBinary(e); err != nil {
					return nil, err
				}
				for i := range res {
					res[i].Meta.Total = len(res[i].Data)
				}
			}
		}
		if e.ReturnBool || len(e.VectorMatching.MatchingLabels) != 0 || shouldDropMetricName(e.Op) {
			removeMetricName(res)
		}
	case *parser.Call:
		fn, ok := calls[e.Func.Name]
		if !ok {
			return nil, fmt.Errorf("not implemented function %q", e.Func.Name)
		}
		if res, err = fn(ev, e.Args); err != nil {
			return nil, err
		}
		ev.r = 0
		switch e.Func.Name {
		case "label_join", "label_replace", "last_over_time":
			break // keep metric name
		default:
			removeMetricName(res)
		}
	case *parser.MatrixSelector:
		if res, err = ev.eval(e.VectorSelector); err != nil {
			return nil, err
		}
		ev.r = e.Range
	case *parser.NumberLiteral:
		res = make([]Series, len(ev.opt.Offsets))
		for i := range res {
			res[i] = ev.newVector(e.Val)
		}
	case *parser.ParenExpr:
		res, err = ev.eval(e.Expr)
	case *parser.SubqueryExpr:
		if res, err = ev.eval(e.Expr); err != nil {
			return nil, err
		}
		ev.r = e.Range
	case *parser.UnaryExpr:
		if res, err = ev.eval(e.Expr); err != nil {
			return nil, err
		}
		if e.Op == parser.SUB {
			for i := range res {
				for j := range res[i].Data {
					s := *res[i].Data[j].Values
					for k := range s {
						s[k] = -s[k]
					}
				}
				res[i].removeMetricName()
			}
		}
	case *parser.VectorSelector:
		res, err = ev.querySeries(e)
	default:
		return nil, fmt.Errorf("not implemented %T", expr)
	}
	return res, err
}

func (ev *evaluator) evalBinary(expr *parser.BinaryExpr) ([]Series, error) {
	if expr.Op.IsSetOperator() {
		expr.VectorMatching.Card = parser.CardManyToMany
	}
	lhs, err := ev.eval(expr.LHS)
	if err != nil {
		return nil, err
	}
	rhs, err := ev.eval(expr.RHS)
	if err != nil {
		return nil, err
	}
	if len(lhs) != len(rhs) {
		panic("LHS length isn't equal to RHS length")
	}
	fn := getBinaryFunc(sliceBinaryFuncM, expr.Op, expr.ReturnBool)
	res := make([]Series, len(lhs))
	for x, lhs := range lhs {
		rhs := rhs[x]
		switch expr.VectorMatching.Card {
		case parser.CardOneToOne:
			if rhs.scalar() {
				for i := range lhs.Data {
					fn(*lhs.Data[i].Values, *lhs.Data[i].Values, *rhs.Data[0].Values)
					lhs.Data[i].What = SelectorWhat{}
				}
				res[x].appendAll(lhs)
				rhs.free(ev)
			} else if lhs.scalar() {
				op := expr.Op
				var swapArgs bool
				switch expr.Op {
				case parser.EQLC:
					swapArgs = true
				case parser.GTE:
					swapArgs = true
					op = parser.LSS
				case parser.GTR:
					swapArgs = true
					op = parser.LTE
				case parser.LSS:
					swapArgs = true
					op = parser.GTE
				case parser.LTE:
					swapArgs = true
					op = parser.GTR
				case parser.NEQ:
					swapArgs = true
				}
				fn = getBinaryFunc(sliceBinaryFuncM, op, expr.ReturnBool)
				if swapArgs {
					for i := range rhs.Data {
						fn(*rhs.Data[i].Values, *rhs.Data[i].Values, *lhs.Data[0].Values)
						rhs.Data[i].What = SelectorWhat{}
					}
				} else {
					for i := range rhs.Data {
						fn(*rhs.Data[i].Values, *lhs.Data[0].Values, *rhs.Data[i].Values)
						rhs.Data[i].What = SelectorWhat{}
					}
				}
				res[x].appendAll(rhs)
				lhs.Data[0].free(ev)
			} else {
				var lhsM map[uint64]hashMeta
				lhsM, err = lhs.hash(ev, hashOptions{
					on:         expr.VectorMatching.On,
					tags:       expr.VectorMatching.MatchingLabels,
					stags:      rhs.Meta.STags,
					listUnused: true,
				})
				if err != nil {
					return nil, err
				}
				var rhsM map[uint64]hashMeta
				rhsM, err = rhs.hash(ev, hashOptions{
					on:    expr.VectorMatching.On,
					tags:  expr.VectorMatching.MatchingLabels,
					stags: lhs.Meta.STags,
				})
				if err != nil {
					return nil, err
				}
				res[x] = ev.newSeries(len(lhsM), evalSeriesMeta(expr, lhs.Meta, rhs.Meta))
				for lhsH, lhsMt := range lhsM {
					if rhsMt, ok := rhsM[lhsH]; ok {
						fn(*lhs.Data[lhsMt.x].Values, *lhs.Data[lhsMt.x].Values, *rhs.Data[rhsMt.x].Values)
						if lhs.Data[lhsMt.x].What != rhs.Data[rhsMt.x].What {
							lhs.Data[lhsMt.x].What = SelectorWhat{}
						}
						for _, v := range lhsMt.unused {
							lhs.Data[lhsMt.x].Tags.remove(v)
						}
						res[x].appendSome(lhs, lhsMt.x)
					} else {
						lhs.Data[lhsMt.x].free(ev)
					}
				}
				rhs.free(ev)
			}
		case parser.CardManyToOne:
			lhs, rhs = rhs, lhs
			fallthrough
		case parser.CardOneToMany:
			var lhsM map[uint64]hashMeta
			lhsM, err = lhs.hash(ev, hashOptions{
				on:    expr.VectorMatching.On,
				tags:  expr.VectorMatching.MatchingLabels,
				stags: rhs.Meta.STags,
			})
			if err != nil {
				return nil, err
			}
			var rhsM map[uint64][]int
			rhsM, _, err = rhs.group(ev, hashOptions{
				on:    expr.VectorMatching.On,
				tags:  expr.VectorMatching.MatchingLabels,
				stags: lhs.Meta.STags,
			})
			if err != nil {
				return nil, err
			}
			res[x] = ev.newSeries(len(rhsM), evalSeriesMeta(expr, lhs.Meta, rhs.Meta))
			for rhsH, rhsXs := range rhsM {
				if lhsMt, ok := lhsM[rhsH]; ok {
					for _, rhsX := range rhsXs {
						fn(*rhs.Data[rhsX].Values, *rhs.Data[rhsX].Values, *lhs.Data[lhsMt.x].Values)
						if rhs.Data[rhsX].What != lhs.Data[lhsMt.x].What {
							rhs.Data[rhsX].What = SelectorWhat{}
						}
						for _, v := range expr.VectorMatching.Include {
							if tag, ok := lhs.Data[lhsMt.x].Tags.get(v); ok {
								rhs.AddTagAt(rhsX, tag)
							}
						}
					}
					res[x].appendSome(rhs, rhsXs...)
				} else {
					ev.freeSome(rhs.Data, rhsXs...)
				}
			}
			lhs.free(ev)
		case parser.CardManyToMany:
			var lhsM map[uint64][]int
			lhsM, _, err = lhs.group(ev, hashOptions{
				on:    expr.VectorMatching.On,
				tags:  expr.VectorMatching.MatchingLabels,
				stags: rhs.Meta.STags,
			})
			if err != nil {
				return nil, err
			}
			switch expr.Op {
			case parser.LAND:
				var rhsM map[uint64][]int
				rhsM, _, err = rhs.group(ev, hashOptions{
					on:    expr.VectorMatching.On,
					tags:  expr.VectorMatching.MatchingLabels,
					stags: lhs.Meta.STags,
				})
				if err != nil {
					return nil, err
				}
				res[x] = ev.newSeries(len(lhsM), lhs.Meta)
				for lhsH, lhsX := range lhsM {
					if rhsX, ok := rhsM[lhsH]; ok {
						for _, rx := range rhsX[1:] {
							sliceOr(*rhs.Data[rhsX[0]].Values, *rhs.Data[rhsX[0]].Values, *rhs.Data[rx].Values)
						}
						for _, lx := range lhsX {
							sliceAnd(*lhs.Data[lx].Values, *lhs.Data[lx].Values, *rhs.Data[rhsX[0]].Values)
						}
						res[x].appendSome(lhs, lhsX...)
					} else {
						ev.freeSome(lhs.Data, lhsX...)
					}
				}
				rhs.free(ev)
			case parser.LDEFAULT:
				if lhs.empty() {
					res[x] = rhs
				} else {
					var rhsM map[uint64]hashMeta
					rhsM, err = rhs.hash(ev, hashOptions{
						on:    expr.VectorMatching.On,
						tags:  expr.VectorMatching.MatchingLabels,
						stags: lhs.Meta.STags,
					})
					if err != nil {
						return nil, err
					}
					res[x] = lhs
					res[x].Meta = evalSeriesMeta(expr, lhs.Meta, rhs.Meta)
					for lhsH, lhsXs := range lhsM {
						if rhsMt, ok := rhsM[lhsH]; ok {
							for _, lhsX := range lhsXs {
								sliceOr(*res[x].Data[lhsX].Values, *lhs.Data[lhsX].Values, *rhs.Data[rhsMt.x].Values)
							}
						}
					}
					rhs.free(ev)
				}
			case parser.LOR:
				var rhsM map[uint64][]int
				rhsM, _, err = rhs.group(ev, hashOptions{
					on:    expr.VectorMatching.On,
					tags:  expr.VectorMatching.MatchingLabels,
					stags: lhs.Meta.STags,
				})
				if err != nil {
					return nil, err
				}
				for rhsH, rhsXs := range rhsM {
					if lhsXs, ok := lhsM[rhsH]; ok {
						for i := range ev.time() {
							for _, lhsX := range lhsXs {
								if !math.IsNaN((*lhs.Data[lhsX].Values)[i]) {
									for _, rhsX := range rhsXs {
										(*rhs.Data[rhsX].Values)[i] = NilValue
									}
									break
								}
							}
						}
					}
				}
				res[x] = lhs
				res[x].appendAll(rhs)
				res[x].Meta = evalSeriesMeta(expr, lhs.Meta, rhs.Meta)
			case parser.LUNLESS:
				var rhsM map[uint64][]int
				rhsM, _, err = rhs.group(ev, hashOptions{
					on:    expr.VectorMatching.On,
					tags:  expr.VectorMatching.MatchingLabels,
					stags: lhs.Meta.STags,
				})
				if err != nil {
					return nil, err
				}
				res[x] = ev.newSeries(len(lhsM), lhs.Meta)
				for lhsH, lhsXs := range lhsM {
					if rhsXs, ok := rhsM[lhsH]; ok {
						for _, rhsX := range rhsXs[1:] {
							sliceOr(*rhs.Data[rhsXs[0]].Values, *rhs.Data[rhsXs[0]].Values, *rhs.Data[rhsX].Values)
						}
						for _, lhsX := range lhsXs {
							sliceUnless(*lhs.Data[lhsX].Values, *lhs.Data[lhsX].Values, *rhs.Data[rhsXs[0]].Values)
						}
					}
					res[x].appendSome(lhs, lhsXs...)
				}
				rhs.free(ev)
			default:
				err = fmt.Errorf("not implemented binary operator %q", expr.Op)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ev *evaluator) querySeries(sel *parser.VectorSelector) (srs []Series, err error) {
	res := make([]Series, len(ev.opt.Offsets))
	run := func(j int, mu *sync.Mutex) error {
		var locked bool
		if mu != nil {
			mu.Lock()
			locked = true
			defer func() {
				if locked {
					mu.Unlock()
				}
			}()
		}
		offset := ev.opt.Offsets[j]
		for i, metric := range sel.MatchingMetrics {
			if ev.trace != nil && ev.debug {
				ev.tracef("#%d request %s: %s", i, metric.Name, sel.What)
			}
			for _, selOffset := range sel.Offsets {
				qry, err := ev.buildSeriesQuery(ev.ctx, sel, metric, sel.Whats, selOffset+offset)
				if err != nil {
					return err
				}
				if qry.empty() {
					if ev.trace != nil && ev.debug {
						ev.tracef("#%d query is empty", i)
					}
					continue
				}
				if mu != nil {
					mu.Unlock()
					locked = false
				}
				sr, cancel, err := ev.h.QuerySeries(ev.ctx, &qry.SeriesQuery)
				if err != nil {
					return err
				}
				if mu != nil {
					mu.Lock()
					locked = true
				}
				ev.cancellationList = append(ev.cancellationList, cancel)
				if ev.trace != nil && ev.debug {
					ev.tracef("#%d series count %d", i, len(sr.Data))
				}
				for k, s := range [2][]*labels.Matcher{sel.MinHostMatchers, sel.MaxHostMatchers} {
					if len(s) != 0 {
						sr.filterMinMaxHost(ev, k, s)
					}
				}
				if qry.prefixSum {
					sr = ev.funcPrefixSum(sr)
				}
				for k := range sr.Data {
					if !sel.OmitNameTag {
						sr.AddTagAt(k, &SeriesTag{
							ID:     labels.MetricName,
							SValue: sel.MatchingNames[i]})
					}
					if len(sel.OriginalOffsetEx) != 0 {
						sr.AddTagAt(k, &SeriesTag{
							ID:    LabelOffset,
							Value: int32(selOffset)})
					}
					sr.Data[k].Offset = offset
				}
				if qry.histogram.restore {
					sr, err = ev.restoreHistogram(sr, qry)
					if err != nil {
						return err
					}
				}
				if len(res[j].Data) == 0 {
					res[j].Meta = sr.Meta
				}
				res[j].appendAll(sr)
			}
		}
		return nil // success
	}
	timeStart := time.Now()
	if ev.opt.QuerySequential || len(res) == 1 {
		for i := 0; i < len(res); i++ {
			err = run(i, nil)
			if err != nil {
				return nil, err
			}
		}
	} else {
		var mu sync.Mutex
		for i := 0; i < len(res); {
			// Limit the number of parallel queries to the number of time shifts
			// available in UI (currently 7) plus always present zero offset.
			var g errgroup.Group
			for n := 8; n != 0 && i < len(res); i, n = i+1, n-1 {
				ii := i
				g.Go(func() error {
					return run(ii, &mu)
				})
			}
			err = g.Wait()
			if err != nil {
				return nil, err
			}
		}
	}
	ev.dataAccessDuration += time.Since(timeStart)
	return res, nil
}

func (ev *evaluator) restoreHistogram(sr Series, qry seriesQueryX) (Series, error) {
	hs, err := sr.histograms(ev)
	if err != nil {
		return Series{}, err
	}
	for _, h := range hs {
		for i := 0; i < len(ev.time()); i++ {
			var acc float64
			for j := 0; j < len(h.buckets); j++ {
				v := (*sr.Data[h.buckets[j].x].Values)[i]
				if math.IsNaN(v) {
					continue
				}
				acc += v
				(*sr.Data[h.buckets[j].x].Values)[i] = acc
			}
		}
	}
	if !qry.histogram.filter {
		return sr, nil
	}
	res := ev.newSeries(len(sr.Data), sr.Meta)
	for _, h := range hs {
		for _, b := range h.buckets {
			if qry.histogram.le == b.le == qry.histogram.eq {
				res.appendSome(sr, b.x)
				break
			}
		}
	}
	return res, nil
}

func (ev *evaluator) buildSeriesQuery(ctx context.Context, sel *parser.VectorSelector, metric *format.MetricMetaValue, selWhats []string, offset int64) (seriesQueryX, error) {
	// whats
	var (
		sels      []SelectorWhat
		prefixSum bool
	)
	for _, selWhat := range selWhats {
		if selWhat == "" {
			continue
		}
		if what, queryFunc, ok := parseSelectorWhat(selWhat); ok {
			sels = append(sels, SelectorWhat{what, queryFunc})
		} else {
			return seriesQueryX{}, fmt.Errorf("unrecognized %s value %q", LabelWhat, selWhat)
		}
	}
	if len(sels) == 0 {
		var what data_model.DigestWhat
		if metric.Kind == format.MetricKindCounter || sel.MetricKindHint == format.MetricKindCounter {
			what = data_model.DigestCountRaw
			prefixSum = true
		} else {
			what = data_model.DigestAvg
		}
		sels = append(sels, SelectorWhat{Digest: what})
	}
	// grouping
	var (
		groupBy    []string
		metricH    = len(sels) == 1 && sels[0].Digest == data_model.DigestCount && metric.Name2Tag[format.LETagName].Raw
		histogramQ histogramQuery
		addGroupBy = func(t format.MetricMetaTag) {
			groupBy = append(groupBy, format.TagID(t.Index))
			if t.Name == format.LETagName {
				histogramQ.restore = true
			}
		}
	)
	if sel.GroupByAll {
		if !sel.GroupWithout {
			for _, t := range metric.Tags {
				if len(t.Name) != 0 {
					addGroupBy(t)
				}
			}
		}
	} else if sel.GroupWithout {
		skip := make(map[int]bool)
		for _, name := range sel.GroupBy {
			t, ok, _ := metric.APICompatGetTag(name)
			if ok {
				skip[t.Index] = true
			}
		}
		for _, t := range metric.Tags {
			if !skip[t.Index] {
				addGroupBy(t)
			}
		}
	} else if len(sel.GroupBy) != 0 {
		groupBy = make([]string, 0, len(sel.GroupBy))
		for _, k := range sel.GroupBy {
			if k == LabelShard {
				groupBy = append(groupBy, format.ShardTagID)
			} else if t, ok, _ := metric.APICompatGetTag(k); ok {
				if t.Index == format.StringTopTagIndex {
					groupBy = append(groupBy, format.StringTopTagID)
				} else {
					addGroupBy(t)
				}
			}
		}
	}
	// filtering
	var (
		filterIn   [format.MaxTags]map[int32]string // tag index -> tag value ID -> tag value
		filterOut  [format.MaxTags]map[int32]string // as above
		sFilterIn  []string
		sFilterOut []string
		emptyCount [format.MaxTags + 1]int // number of "MatchEqual" or "MatchRegexp" filters which are guaranteed to yield empty response
	)
	for _, matcher := range sel.LabelMatchers {
		if strings.HasPrefix(matcher.Name, "__") {
			continue
		}
		tag, ok, _ := metric.APICompatGetTag(matcher.Name)
		if !ok {
			return seriesQueryX{}, fmt.Errorf("not found tag %q", matcher.Name)
		}
		if tag.Index == format.StringTopTagIndex {
			switch matcher.Type {
			case labels.MatchEqual:
				sFilterIn = append(sFilterIn, matcher.Value)
			case labels.MatchNotEqual:
				sFilterOut = append(sFilterOut, matcher.Value)
			case labels.MatchRegexp:
				stop, err := ev.getStringTop(ctx, metric, offset)
				if err != nil {
					return seriesQueryX{}, err
				}
				var n int
				for _, v := range stop {
					if matcher.Matches(v) {
						sFilterIn = append(sFilterIn, v)
						n++
					}
				}
				if n == 0 {
					// there no data satisfying the filter
					emptyCount[format.MaxTags]++
					continue
				}
			case labels.MatchNotRegexp:
				stop, err := ev.getStringTop(ctx, metric, offset)
				if err != nil {
					return seriesQueryX{}, err
				}
				for _, v := range stop {
					if !matcher.Matches(v) {
						sFilterOut = append(sFilterOut, v)
					}
				}
			}
		} else {
			i := tag.Index
			switch matcher.Type {
			case labels.MatchEqual:
				id, err := ev.getTagValueID(metric, i, matcher.Value)
				if err != nil {
					if errors.Is(err, ErrNotFound) {
						// string is not mapped, result is guaranteed to be empty
						emptyCount[i]++
						continue
					} else {
						return seriesQueryX{}, fmt.Errorf("failed to map string %q: %v", matcher.Value, err)
					}
				}
				if metricH && !histogramQ.restore && matcher.Name == format.LETagName {
					histogramQ.filter = true
					histogramQ.eq = true
					histogramQ.le = statshouse.LexDecode(id)
				} else if filterIn[i] != nil {
					filterIn[i][id] = matcher.Value
				} else {
					filterIn[i] = map[int32]string{id: matcher.Value}
				}
			case labels.MatchNotEqual:
				id, err := ev.getTagValueID(metric, i, matcher.Value)
				if err != nil {
					if errors.Is(err, ErrNotFound) {
						continue // ignore values with no mapping
					}
					return seriesQueryX{}, err
				}
				if metricH && !histogramQ.restore && matcher.Name == format.LETagName {
					histogramQ.filter = true
					histogramQ.eq = false
					histogramQ.le = statshouse.LexDecode(id)
				} else if filterOut[i] != nil {
					filterOut[i][id] = matcher.Value
				} else {
					filterOut[i] = map[int32]string{id: matcher.Value}
				}
			case labels.MatchRegexp:
				m, err := ev.getTagValues(ctx, metric, i, offset)
				if err != nil {
					return seriesQueryX{}, err
				}
				in := make(map[int32]string)
				for id, str := range m {
					if matcher.Matches(str) {
						in[id] = str
					}
				}
				if len(in) == 0 {
					// there no data satisfying the filter
					emptyCount[i]++
					continue
				}
				filterIn[i] = in
			case labels.MatchNotRegexp:
				m, err := ev.getTagValues(ctx, metric, i, offset)
				if err != nil {
					return seriesQueryX{}, err
				}
				out := make(map[int32]string)
				for id, str := range m {
					if !matcher.Matches(str) {
						out[id] = str
					}
				}
				filterOut[i] = out
			}
		}
	}
	for i, n := range emptyCount {
		if n == 0 {
			continue
		}
		var m int
		if i == format.MaxTags {
			m = len(sFilterIn)
		} else {
			m = len(filterIn[i])
		}
		if m == 0 {
			// All "MatchEqual" and "MatchRegexp" filters give an empty result and
			// there are no other such filters, overall result is guaranteed to be empty
			return seriesQueryX{}, nil
		}
	}
	if histogramQ.filter && !histogramQ.restore {
		groupBy = append(groupBy, format.TagID(format.LETagIndex))
		histogramQ.restore = true
	}
	// remove dublicates in "groupBy"
	sort.Strings(groupBy)
	for i := 1; i < len(groupBy); i++ {
		j := i
		for j < len(groupBy) && groupBy[i-1] == groupBy[j] {
			j++
		}
		if i != j {
			groupBy = append(groupBy[:i], groupBy[j:]...)
		}
	}
	return seriesQueryX{
			SeriesQuery{
				Metric:     metric,
				Whats:      sels,
				Timescale:  ev.t,
				Offset:     offset,
				Range:      sel.Range,
				GroupBy:    groupBy,
				FilterIn:   filterIn,
				FilterOut:  filterOut,
				SFilterIn:  sFilterIn,
				SFilterOut: sFilterOut,
				MinMaxHost: [2]bool{sel.MinHost, sel.MaxHost},
				Options:    ev.opt,
			},
			prefixSum,
			histogramQ},
		nil
}

func (ev *evaluator) newSeries(capacity int, meta SeriesMeta) Series {
	if capacity == 0 {
		return Series{Meta: meta}
	}
	return Series{
		Data: make([]SeriesData, 0, capacity),
		Meta: meta,
	}
}

func (ev *evaluator) newVector(v float64) Series {
	res := Series{
		Data: []SeriesData{{Values: ev.alloc()}},
	}
	for i := range *res.Data[0].Values {
		(*res.Data[0].Values)[i] = v
	}
	return res
}

func (ev *evaluator) getTagValues(ctx context.Context, metric *format.MetricMetaValue, tagX int, offset int64) (map[int32]string, error) {
	m, ok := ev.tags[metric]
	if !ok {
		// tag index -> offset -> tag value ID -> tag value
		m = make([]map[int64]map[int32]string, format.MaxTags)
		ev.tags[metric] = m
	}
	m2 := m[tagX]
	if m2 == nil {
		// offset -> tag value ID -> tag value
		m2 = make(map[int64]map[int32]string)
		m[tagX] = m2
	}
	var res map[int32]string
	if res, ok = m2[offset]; ok {
		return res, nil
	}
	ids, err := ev.h.QueryTagValueIDs(ctx, TagValuesQuery{
		Metric:    metric,
		TagIndex:  tagX,
		Timescale: ev.t,
		Offset:    offset,
		Options:   ev.opt,
	})
	if err != nil {
		return nil, err
	}
	// tag value ID -> tag value
	res = make(map[int32]string, len(ids))
	for _, id := range ids {
		res[id] = ev.h.GetTagValue(TagValueQuery{
			Version:    ev.opt.Version,
			Metric:     metric,
			TagIndex:   tagX,
			TagValueID: id,
		})
	}
	m2[offset] = res
	return res, nil
}

func (ev *evaluator) getTagValueID(metric *format.MetricMetaValue, tagX int, tagV string) (int32, error) {
	if format.HasRawValuePrefix(tagV) {
		return format.ParseCodeTagValue(tagV)
	}
	if tagX < 0 || len(metric.Tags) <= tagX {
		return 0, ErrNotFound
	}
	t := metric.Tags[tagX]
	if t.Name == labels.BucketLabel && t.Raw {
		if v, err := strconv.ParseFloat(tagV, 32); err == nil {
			return statshouse.LexEncode(float32(v)), nil
		}
	}
	return ev.h.GetTagValueID(TagValueIDQuery{
		Version:  ev.opt.Version,
		Metric:   metric,
		TagIndex: tagX,
		TagValue: tagV,
	})
}

func (ev *evaluator) getStringTop(ctx context.Context, metric *format.MetricMetaValue, offset int64) ([]string, error) {
	m, ok := ev.stop[metric]
	if !ok {
		// offset -> tag values
		m = make(map[int64][]string)
		ev.stop[metric] = m
	}
	var res []string
	if res, ok = m[offset]; ok {
		return res, nil
	}
	var err error
	res, err = ev.h.QueryStringTop(ctx, TagValuesQuery{
		Metric:    metric,
		Timescale: ev.t,
		Offset:    offset,
		Options:   ev.opt,
	})
	if err == nil {
		m[offset] = res
	}
	return res, err
}

func (ev *evaluator) weight(ds []SeriesData) []float64 {
	var (
		w      = make([]float64, len(ds))
		nodecN int // number of non-decreasing series
	)
	for i, d := range ds {
		var (
			j     int
			acc   float64
			prev  = -math.MaxFloat64
			nodec = true // non-decreasing
		)
		for _, lod := range ev.t.LODs {
			for m := 0; m < lod.Len; m++ {
				k := j + m
				// do not count invisible
				if k < ev.t.ViewStartX {
					continue
				} else if k >= ev.t.ViewEndX {
					break
				}
				v := (*d.Values)[k]
				if !math.IsNaN(v) {
					acc += v * v * float64(lod.Step)
					if v < prev {
						nodec = false
					}
					prev = v
				}
			}
			j += lod.Len
		}
		w[i] = acc
		if nodec {
			nodecN++
		}
	}
	if nodecN == len(w) {
		// all series are non-decreasing, weight is a last value
		for i, s := range ds {
			last := -math.MaxFloat64
			for i := ev.t.ViewEndX; i > 0; i-- {
				v := (*s.Values)[i-1]
				if !math.IsNaN(v) {
					last = v
					break
				}
			}
			w[i] = last
		}
	}
	return w
}

func (ev *evaluator) alloc() *[]float64 {
	var s *[]float64
	if len(ev.reuseList) != 0 {
		ev.reuseList, s = removeLast(ev.reuseList)
	} else if len(ev.freeList) != 0 {
		ev.freeList, s = removeLast(ev.freeList)
	} else {
		if ev.allocMap == nil {
			ev.allocMap = make(map[*[]float64]bool)
		}
		s = ev.h.Alloc(len(ev.time()))
		ev.allocMap[s] = true
	}
	return s
}

func (ev *evaluator) free(s *[]float64) {
	if ev.allocMap != nil && ev.allocMap[s] {
		delete(ev.allocMap, s)
		ev.freeList = append(ev.freeList, s)
	} else {
		ev.reuseList = append(ev.reuseList, s)
	}
}

func (ev *evaluator) freeAll(ds []SeriesData) {
	for x := range ds {
		ds[x].free(ev)
	}
}

func (ev *evaluator) freeSome(ds []SeriesData, xs ...int) {
	for _, x := range xs {
		ds[x].free(ev)
	}
}

func (ev *evaluator) cancel() {
	if ev.allocMap != nil {
		for s := range ev.allocMap {
			ev.h.Free(s)
		}
		ev.allocMap = nil
	}
	for _, s := range ev.freeList {
		ev.h.Free(s)
	}
	ev.freeList = nil
	for _, cancel := range ev.cancellationList {
		cancel()
	}
	ev.reuseList = nil
	ev.cancellationList = nil
}

func (ev *evaluator) newWindow(v []float64, s bool) window {
	return newWindow(ev.time(), v, ev.r, ev.t.LODs[len(ev.t.LODs)-1].Step, s)
}

func (ev *evaluator) reportStat(qry Query, timeEnd time.Time) {
	tags := statshouse.Tags{
		1: srvfunc.HostnameForStatshouse(),
	}
	r := qry.End - qry.Start
	x := 2
	switch {
	// add one because UI always requests one second more
	case r <= 1+1:
		tags[x] = "1" // "1_second"
	case r <= 300+1:
		tags[x] = "2" // "5_minutes"
	case r <= 900+1:
		tags[x] = "3" // "15_minutes"
	case r <= 3600+1:
		tags[x] = "4" // "1_hour"
	case r <= 7200+1:
		tags[x] = "5" // "2_hours"
	case r <= 21600+1:
		tags[x] = "6" // "6_hours"
	case r <= 43200+1:
		tags[x] = "7" // "12_hours"
	case r <= 86400+1:
		tags[x] = "8" // "1_day"
	case r <= 172800+1:
		tags[x] = "9" // "2_days"
	case r <= 259200+1:
		tags[x] = "10" // "3_days"
	case r <= 604800+1:
		tags[x] = "11" // "1_week"
	case r <= 1209600+1:
		tags[x] = "12" // "2_weeks"
	case r <= 2592000+1:
		tags[x] = "13" // "1_month"
	case r <= 7776000+1:
		tags[x] = "14" // "3_months"
	case r <= 15552000+1:
		tags[x] = "15" // "6_months"
	case r <= 31536000+1:
		tags[x] = "16" // "1_year"
	case r <= 63072000+1:
		tags[x] = "17" // "2_years"
	default:
		tags[x] = "18"
	}
	n := len(ev.t.Time) - ev.t.StartX // number of points
	x = 3
	switch {
	case n <= 1:
		tags[x] = "1"
	case n <= 1024:
		tags[x] = "2" // "1K"
	case n <= 2048:
		tags[x] = "3" // "2K"
	case n <= 3072:
		tags[x] = "4" // "3K"
	case n <= 4096:
		tags[x] = "5" // "4K"
	case n <= 5120:
		tags[x] = "6" // "5K"
	case n <= 6144:
		tags[x] = "7" // "6K"
	case n <= 7168:
		tags[x] = "8" // "7K"
	case n <= 8192:
		tags[x] = "9" // "8K"
	default:
		tags[x] = "10"
	}
	x = 4
	tags[x] = "1" // "query_parsing"
	value := ev.timeQueryParseEnd.Sub(ev.timeStart).Seconds()
	statshouse.Metric(format.BuiltinMetricNamePromQLEngineTime, tags).Value(value)
	tags[x] = "2" // "data_access"
	value = ev.dataAccessDuration.Seconds()
	statshouse.Metric(format.BuiltinMetricNamePromQLEngineTime, tags).Value(value)
	tags[x] = "3" // "data_processing"
	value = (timeEnd.Sub(ev.timeQueryParseEnd) - ev.dataAccessDuration).Seconds()
	statshouse.Metric(format.BuiltinMetricNamePromQLEngineTime, tags).Value(value)
}

func (qry *seriesQueryX) empty() bool {
	return qry.Metric == nil
}

func TraceContext(ctx context.Context, s *[]string) context.Context {
	return context.WithValue(ctx, traceContextKey, &traceContext{s, false})
}

func DebugContext(ctx context.Context, s *[]string) context.Context {
	return context.WithValue(ctx, traceContextKey, &traceContext{s, true})
}

func (ev *evaluator) tracef(format string, a ...any) {
	*ev.trace = append(*ev.trace, fmt.Sprintf(format, a...))
}

func evalLiteral(expr parser.Expr) (*parser.NumberLiteral, bool) {
	switch e := expr.(type) {
	case *parser.BinaryExpr:
		l, okL := evalLiteral(e.LHS)
		r, okR := evalLiteral(e.RHS)
		if okL && okR {
			return &parser.NumberLiteral{Val: getBinaryFunc(scalarBinaryFuncM, e.Op, e.ReturnBool)(l.Val, r.Val)}, true
		} else if okL {
			e.LHS = l
		} else if okR {
			e.RHS = r
		}
	case *parser.NumberLiteral:
		return e, true
	case *parser.ParenExpr:
		return evalLiteral(e.Expr)
	case *parser.UnaryExpr:
		if l, ok := evalLiteral(e.Expr); ok {
			if e.Op == parser.SUB {
				l.Val = -l.Val
			}
			return l, true
		}
	default:
		for _, node := range parser.Children(expr) {
			if e2, ok := node.(parser.Expr); ok {
				evalLiteral(e2)
			}
		}
	}
	return nil, false
}

func removeLast[V any](s []V) ([]V, V) {
	return s[:len(s)-1], s[len(s)-1]
}

func normalizeOffsets(s []int64) []int64 {
	res := make([]int64, 0, len(s))
	res = append(res, s...)
	// sort descending
	sort.Sort(sort.Reverse(sortkeys.Int64Slice(res)))
	// remove duplicates
	var i int
	for j := 1; j < len(res); j++ {
		if res[i] != res[j] {
			i++
			res[i] = res[j]
		}
	}
	return res[:i+1]
}

func parseSelectorWhat(str string) (data_model.DigestWhat, string, bool) {
	var digestWhat, queryFunc string
	if i := strings.Index(str, ":"); i != -1 {
		digestWhat = str[:i]
		queryFunc = str[i+1:]
	} else {
		digestWhat = str
	}
	var res data_model.DigestWhat
	switch digestWhat {
	case Count:
		res = data_model.DigestCount
	case CountSec:
		res = data_model.DigestCountSec
	case CountRaw:
		res = data_model.DigestCountRaw
	case Min:
		res = data_model.DigestMin
	case Max:
		res = data_model.DigestMax
	case Sum:
		res = data_model.DigestSum
	case SumSec:
		res = data_model.DigestSumSec
	case SumRaw:
		res = data_model.DigestSumRaw
	case Avg:
		res = data_model.DigestAvg
	case StdDev:
		res = data_model.DigestStdDev
	case StdVar:
		res = data_model.DigestStdVar
	case P0_1:
		res = data_model.DigestP0_1
	case P1:
		res = data_model.DigestP1
	case P5:
		res = data_model.DigestP5
	case P10:
		res = data_model.DigestP10
	case P25:
		res = data_model.DigestP25
	case P50:
		res = data_model.DigestP50
	case P75:
		res = data_model.DigestP75
	case P90:
		res = data_model.DigestP90
	case P95:
		res = data_model.DigestP95
	case P99:
		res = data_model.DigestP99
	case P999:
		res = data_model.DigestP999
	case Cardinality:
		res = data_model.DigestCardinality
	case CardinalitySec:
		res = data_model.DigestCardinalitySec
	case CardinalityRaw:
		res = data_model.DigestCardinalityRaw
	case Unique:
		res = data_model.DigestUnique
	case UniqueSec:
		res = data_model.DigestUniqueSec
	default:
		return data_model.DigestUnspecified, "", false
	}
	return res, queryFunc, true
}
