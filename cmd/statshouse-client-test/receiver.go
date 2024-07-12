package main

import (
	"bytes"
	"log"
	"slices"
	"sort"
	"sync"

	"github.com/vkcom/statshouse/internal/data_model"
	"github.com/vkcom/statshouse/internal/data_model/gen2/tlstatshouse"
	"github.com/vkcom/statshouse/internal/receiver"
)

type series map[tags]map[uint32]*value
type tag [2]string // name, value
type tags [17]tag  // metric name + 16 tags

type value struct {
	counter float64
	values  []float64
	uniques []int64
}

type handler struct {
	sink func(*tlstatshouse.MetricBytes)
}

func (h handler) HandleMetrics(args data_model.HandlerArgs) (data_model.MappedMetricHeader, bool) {
	h.sink(args.MetricBytes)
	return data_model.MappedMetricHeader{}, false
}

func (handler) HandleParseError(pkt []byte, err error) {
	log.Fatalln(pkt, err)
}

func listenUDP(args argv, ch chan series) (func(), error) {
	addr := ":13337"
	ln, err := receiver.ListenUDP("udp4", addr, 16*1024*1024, false, nil, nil)
	if err != nil {
		return nil, err
	}
	log.Printf("listen UDP %s\n", addr)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer ln.Close()
		res := make(series, args.m)
		var n int
		if err := ln.Serve(handler{func(b *tlstatshouse.MetricBytes) {
			if bytes.Equal(b.Name, endOfIterationMarkBytes) {
				n += int(b.Counter) // iteration #n
				if n == args.n {
					ch <- res
					res = make(series, len(res))
					n = 0
				}
			} else {
				res.addMetricBytes(b)
			}
		}}); err != nil {
			log.Fatalln(err)
		}
	}()
	return func() {
		ln.Close()
		wg.Wait()
	}, nil
}

func (s series) addMetricBytes(b *tlstatshouse.MetricBytes) {
	tags := tags{{"", string(b.Name)}}
	for i := 0; i < len(b.Tags); i++ {
		if len(b.Tags[i].Value) != 0 {
			tags[i+1] = tag{string(b.Tags[i].Key), string(b.Tags[i].Value)}
		}
	}
	sort.Slice(tags[:], func(i, j int) bool {
		return slices.Compare(tags[i][:], tags[j][:]) < 0
	})
	vals := s[tags]
	if vals == nil {
		vals = map[uint32]*value{}
		s[tags] = vals
	}
	val := vals[b.Ts]
	if val == nil {
		val = &value{}
		vals[b.Ts] = val
	}
	val.addMetricBytes(b)
}

func (v *value) addMetricBytes(b *tlstatshouse.MetricBytes) {
	v.counter += b.Counter
	v.values = append(v.values, b.Value...)
	v.uniques = append(v.uniques, b.Unique...)
}

func (v *value) addMetric(b metric) {
	v.counter += b.Count
	v.values = append(v.values, b.Values...)
	v.uniques = append(v.uniques, b.Uniques...)
}
