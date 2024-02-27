// Copyright 2023 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package internal

import (
	"github.com/vkcom/statshouse/internal/vkgo/basictl"
)

var _ = basictl.NatWrite

type StatshouseApiQuery struct {
	FieldsMask uint32
	Version    int32
	TopN       int32
	MetricName string
	TimeFrom   int64
	TimeTo     int64
	Interval   string
	Function   StatshouseApiFunction
	GroupBy    []string
	Filter     []StatshouseApiFilter
	TimeShift  []int64
	Promql     string                  // Conditional: item.FieldsMask.0
	What       []StatshouseApiFunction // Conditional: item.FieldsMask.1
	// ExcessPointsFlag (TrueType) // Conditional: item.FieldsMask.2
	WidthAgg string // Conditional: item.FieldsMask.3
	// NameFlag (TrueType) // Conditional: item.FieldsMask.4
	// ColorFlag (TrueType) // Conditional: item.FieldsMask.5
	// TotalFlag (TrueType) // Conditional: item.FieldsMask.6
	// MaxHostFlag (TrueType) // Conditional: item.FieldsMask.7
}

func (StatshouseApiQuery) TLName() string { return "statshouseApi.query" }
func (StatshouseApiQuery) TLTag() uint32  { return 0xc9951bb9 }

func (item *StatshouseApiQuery) SetPromql(v string) {
	item.Promql = v
	item.FieldsMask |= 1 << 0
}
func (item *StatshouseApiQuery) ClearPromql() {
	item.Promql = ""
	item.FieldsMask &^= 1 << 0
}
func (item StatshouseApiQuery) IsSetPromql() bool { return item.FieldsMask&(1<<0) != 0 }

func (item *StatshouseApiQuery) SetWhat(v []StatshouseApiFunction) {
	item.What = v
	item.FieldsMask |= 1 << 1
}
func (item *StatshouseApiQuery) ClearWhat() {
	item.What = item.What[:0]
	item.FieldsMask &^= 1 << 1
}
func (item StatshouseApiQuery) IsSetWhat() bool { return item.FieldsMask&(1<<1) != 0 }

func (item *StatshouseApiQuery) SetExcessPointsFlag(v bool) {
	if v {
		item.FieldsMask |= 1 << 2
	} else {
		item.FieldsMask &^= 1 << 2
	}
}
func (item StatshouseApiQuery) IsSetExcessPointsFlag() bool { return item.FieldsMask&(1<<2) != 0 }

func (item *StatshouseApiQuery) SetWidthAgg(v string) {
	item.WidthAgg = v
	item.FieldsMask |= 1 << 3
}
func (item *StatshouseApiQuery) ClearWidthAgg() {
	item.WidthAgg = ""
	item.FieldsMask &^= 1 << 3
}
func (item StatshouseApiQuery) IsSetWidthAgg() bool { return item.FieldsMask&(1<<3) != 0 }

func (item *StatshouseApiQuery) SetNameFlag(v bool) {
	if v {
		item.FieldsMask |= 1 << 4
	} else {
		item.FieldsMask &^= 1 << 4
	}
}
func (item StatshouseApiQuery) IsSetNameFlag() bool { return item.FieldsMask&(1<<4) != 0 }

func (item *StatshouseApiQuery) SetColorFlag(v bool) {
	if v {
		item.FieldsMask |= 1 << 5
	} else {
		item.FieldsMask &^= 1 << 5
	}
}
func (item StatshouseApiQuery) IsSetColorFlag() bool { return item.FieldsMask&(1<<5) != 0 }

func (item *StatshouseApiQuery) SetTotalFlag(v bool) {
	if v {
		item.FieldsMask |= 1 << 6
	} else {
		item.FieldsMask &^= 1 << 6
	}
}
func (item StatshouseApiQuery) IsSetTotalFlag() bool { return item.FieldsMask&(1<<6) != 0 }

func (item *StatshouseApiQuery) SetMaxHostFlag(v bool) {
	if v {
		item.FieldsMask |= 1 << 7
	} else {
		item.FieldsMask &^= 1 << 7
	}
}
func (item StatshouseApiQuery) IsSetMaxHostFlag() bool { return item.FieldsMask&(1<<7) != 0 }

func (item *StatshouseApiQuery) Reset() {
	item.FieldsMask = 0
	item.Version = 0
	item.TopN = 0
	item.MetricName = ""
	item.TimeFrom = 0
	item.TimeTo = 0
	item.Interval = ""
	item.Function.Reset()
	item.GroupBy = item.GroupBy[:0]
	item.Filter = item.Filter[:0]
	item.TimeShift = item.TimeShift[:0]
	item.Promql = ""
	item.What = item.What[:0]
	item.WidthAgg = ""
}

func (item *StatshouseApiQuery) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.Version); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.TopN); err != nil {
		return w, err
	}
	if w, err = basictl.StringRead(w, &item.MetricName); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.TimeFrom); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.TimeTo); err != nil {
		return w, err
	}
	if w, err = basictl.StringRead(w, &item.Interval); err != nil {
		return w, err
	}
	if w, err = item.Function.ReadBoxed(w); err != nil {
		return w, err
	}
	if w, err = BuiltinVectorStringRead(w, &item.GroupBy); err != nil {
		return w, err
	}
	if w, err = BuiltinVectorStatshouseApiFilterRead(w, &item.Filter); err != nil {
		return w, err
	}
	if w, err = BuiltinVectorLongRead(w, &item.TimeShift); err != nil {
		return w, err
	}
	if item.FieldsMask&(1<<0) != 0 {
		if w, err = basictl.StringRead(w, &item.Promql); err != nil {
			return w, err
		}
	} else {
		item.Promql = ""
	}
	if item.FieldsMask&(1<<1) != 0 {
		if w, err = BuiltinVectorStatshouseApiFunctionBoxedRead(w, &item.What); err != nil {
			return w, err
		}
	} else {
		item.What = item.What[:0]
	}
	if item.FieldsMask&(1<<3) != 0 {
		if w, err = basictl.StringRead(w, &item.WidthAgg); err != nil {
			return w, err
		}
	} else {
		item.WidthAgg = ""
	}
	return w, nil
}

func (item *StatshouseApiQuery) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	w = basictl.IntWrite(w, item.Version)
	w = basictl.IntWrite(w, item.TopN)
	if w, err = basictl.StringWrite(w, item.MetricName); err != nil {
		return w, err
	}
	w = basictl.LongWrite(w, item.TimeFrom)
	w = basictl.LongWrite(w, item.TimeTo)
	if w, err = basictl.StringWrite(w, item.Interval); err != nil {
		return w, err
	}
	if w, err = item.Function.WriteBoxed(w); err != nil {
		return w, err
	}
	if w, err = BuiltinVectorStringWrite(w, item.GroupBy); err != nil {
		return w, err
	}
	if w, err = BuiltinVectorStatshouseApiFilterWrite(w, item.Filter); err != nil {
		return w, err
	}
	if w, err = BuiltinVectorLongWrite(w, item.TimeShift); err != nil {
		return w, err
	}
	if item.FieldsMask&(1<<0) != 0 {
		if w, err = basictl.StringWrite(w, item.Promql); err != nil {
			return w, err
		}
	}
	if item.FieldsMask&(1<<1) != 0 {
		if w, err = BuiltinVectorStatshouseApiFunctionBoxedWrite(w, item.What); err != nil {
			return w, err
		}
	}
	if item.FieldsMask&(1<<3) != 0 {
		if w, err = basictl.StringWrite(w, item.WidthAgg); err != nil {
			return w, err
		}
	}
	return w, nil
}

func (item *StatshouseApiQuery) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xc9951bb9); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseApiQuery) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xc9951bb9)
	return item.Write(w)
}

func (item StatshouseApiQuery) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func StatshouseApiQuery__ReadJSON(item *StatshouseApiQuery, j interface{}) error {
	return item.readJSON(j)
}
func (item *StatshouseApiQuery) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouseApi.query", "expected json object")
	}
	_jFieldsMask := _jm["fields_mask"]
	delete(_jm, "fields_mask")
	if err := JsonReadUint32(_jFieldsMask, &item.FieldsMask); err != nil {
		return err
	}
	_jVersion := _jm["version"]
	delete(_jm, "version")
	if err := JsonReadInt32(_jVersion, &item.Version); err != nil {
		return err
	}
	_jTopN := _jm["top_n"]
	delete(_jm, "top_n")
	if err := JsonReadInt32(_jTopN, &item.TopN); err != nil {
		return err
	}
	_jMetricName := _jm["metric_name"]
	delete(_jm, "metric_name")
	if err := JsonReadString(_jMetricName, &item.MetricName); err != nil {
		return err
	}
	_jTimeFrom := _jm["time_from"]
	delete(_jm, "time_from")
	if err := JsonReadInt64(_jTimeFrom, &item.TimeFrom); err != nil {
		return err
	}
	_jTimeTo := _jm["time_to"]
	delete(_jm, "time_to")
	if err := JsonReadInt64(_jTimeTo, &item.TimeTo); err != nil {
		return err
	}
	_jInterval := _jm["interval"]
	delete(_jm, "interval")
	if err := JsonReadString(_jInterval, &item.Interval); err != nil {
		return err
	}
	_jFunction := _jm["function"]
	delete(_jm, "function")
	_jGroupBy := _jm["group_by"]
	delete(_jm, "group_by")
	_jFilter := _jm["filter"]
	delete(_jm, "filter")
	_jTimeShift := _jm["time_shift"]
	delete(_jm, "time_shift")
	_jPromql := _jm["promql"]
	delete(_jm, "promql")
	_jWhat := _jm["what"]
	delete(_jm, "what")
	_jExcessPointsFlag := _jm["excess_points_flag"]
	delete(_jm, "excess_points_flag")
	_jWidthAgg := _jm["widthAgg"]
	delete(_jm, "widthAgg")
	_jNameFlag := _jm["name_flag"]
	delete(_jm, "name_flag")
	_jColorFlag := _jm["color_flag"]
	delete(_jm, "color_flag")
	_jTotalFlag := _jm["total_flag"]
	delete(_jm, "total_flag")
	_jMaxHostFlag := _jm["max_host_flag"]
	delete(_jm, "max_host_flag")
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouseApi.query", k)
	}
	if _jPromql != nil {
		item.FieldsMask |= 1 << 0
	}
	if _jWhat != nil {
		item.FieldsMask |= 1 << 1
	}
	if _jExcessPointsFlag != nil {
		_bit := false
		if err := JsonReadBool(_jExcessPointsFlag, &_bit); err != nil {
			return err
		}
		if _bit {
			item.FieldsMask |= 1 << 2
		} else {
			item.FieldsMask &^= 1 << 2
		}
	}
	if _jWidthAgg != nil {
		item.FieldsMask |= 1 << 3
	}
	if _jNameFlag != nil {
		_bit := false
		if err := JsonReadBool(_jNameFlag, &_bit); err != nil {
			return err
		}
		if _bit {
			item.FieldsMask |= 1 << 4
		} else {
			item.FieldsMask &^= 1 << 4
		}
	}
	if _jColorFlag != nil {
		_bit := false
		if err := JsonReadBool(_jColorFlag, &_bit); err != nil {
			return err
		}
		if _bit {
			item.FieldsMask |= 1 << 5
		} else {
			item.FieldsMask &^= 1 << 5
		}
	}
	if _jTotalFlag != nil {
		_bit := false
		if err := JsonReadBool(_jTotalFlag, &_bit); err != nil {
			return err
		}
		if _bit {
			item.FieldsMask |= 1 << 6
		} else {
			item.FieldsMask &^= 1 << 6
		}
	}
	if _jMaxHostFlag != nil {
		_bit := false
		if err := JsonReadBool(_jMaxHostFlag, &_bit); err != nil {
			return err
		}
		if _bit {
			item.FieldsMask |= 1 << 7
		} else {
			item.FieldsMask &^= 1 << 7
		}
	}
	if err := StatshouseApiFunction__ReadJSON(&item.Function, _jFunction); err != nil {
		return err
	}
	if err := BuiltinVectorStringReadJSON(_jGroupBy, &item.GroupBy); err != nil {
		return err
	}
	if err := BuiltinVectorStatshouseApiFilterReadJSON(_jFilter, &item.Filter); err != nil {
		return err
	}
	if err := BuiltinVectorLongReadJSON(_jTimeShift, &item.TimeShift); err != nil {
		return err
	}
	if _jPromql != nil {
		if err := JsonReadString(_jPromql, &item.Promql); err != nil {
			return err
		}
	} else {
		item.Promql = ""
	}
	if _jWhat != nil {
		if err := BuiltinVectorStatshouseApiFunctionBoxedReadJSON(_jWhat, &item.What); err != nil {
			return err
		}
	} else {
		item.What = item.What[:0]
	}
	if _jWidthAgg != nil {
		if err := JsonReadString(_jWidthAgg, &item.WidthAgg); err != nil {
			return err
		}
	} else {
		item.WidthAgg = ""
	}
	return nil
}

func (item *StatshouseApiQuery) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *StatshouseApiQuery) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldsMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"fields_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldsMask)
	}
	if item.Version != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"version":`...)
		w = basictl.JSONWriteInt32(w, item.Version)
	}
	if item.TopN != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"top_n":`...)
		w = basictl.JSONWriteInt32(w, item.TopN)
	}
	if len(item.MetricName) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"metric_name":`...)
		w = basictl.JSONWriteString(w, item.MetricName)
	}
	if item.TimeFrom != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"time_from":`...)
		w = basictl.JSONWriteInt64(w, item.TimeFrom)
	}
	if item.TimeTo != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"time_to":`...)
		w = basictl.JSONWriteInt64(w, item.TimeTo)
	}
	if len(item.Interval) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"interval":`...)
		w = basictl.JSONWriteString(w, item.Interval)
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"function":`...)
	if w, err = item.Function.WriteJSONOpt(short, w); err != nil {
		return w, err
	}
	if len(item.GroupBy) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"group_by":`...)
		if w, err = BuiltinVectorStringWriteJSONOpt(short, w, item.GroupBy); err != nil {
			return w, err
		}
	}
	if len(item.Filter) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"filter":`...)
		if w, err = BuiltinVectorStatshouseApiFilterWriteJSONOpt(short, w, item.Filter); err != nil {
			return w, err
		}
	}
	if len(item.TimeShift) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"time_shift":`...)
		if w, err = BuiltinVectorLongWriteJSONOpt(short, w, item.TimeShift); err != nil {
			return w, err
		}
	}
	if item.FieldsMask&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"promql":`...)
		w = basictl.JSONWriteString(w, item.Promql)
	}
	if item.FieldsMask&(1<<1) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"what":`...)
		if w, err = BuiltinVectorStatshouseApiFunctionBoxedWriteJSONOpt(short, w, item.What); err != nil {
			return w, err
		}
	}
	if item.FieldsMask&(1<<2) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"excess_points_flag":true`...)
	}
	if item.FieldsMask&(1<<3) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"widthAgg":`...)
		w = basictl.JSONWriteString(w, item.WidthAgg)
	}
	if item.FieldsMask&(1<<4) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"name_flag":true`...)
	}
	if item.FieldsMask&(1<<5) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"color_flag":true`...)
	}
	if item.FieldsMask&(1<<6) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"total_flag":true`...)
	}
	if item.FieldsMask&(1<<7) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"max_host_flag":true`...)
	}
	return append(w, '}'), nil
}

func (item *StatshouseApiQuery) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseApiQuery) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("statshouseApi.query", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("statshouseApi.query", err.Error())
	}
	return nil
}
