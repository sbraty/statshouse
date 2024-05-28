package stats

import (
	"strconv"

	"github.com/vkcom/statshouse-go"
	"github.com/vkcom/statshouse/internal/data_model"
	"github.com/vkcom/statshouse/internal/data_model/gen2/tl"
	"github.com/vkcom/statshouse/internal/data_model/gen2/tlstatshouse"
	"github.com/vkcom/statshouse/internal/env"
	"github.com/vkcom/statshouse/internal/receiver"
)

type MetricWriter interface {
	WriteSystemMetricValue(nowUnix int64, name string, value float64, tagsList ...int32)
	WriteSystemMetricCount(nowUnix int64, name string, count float64, tagsList ...int32)
	WriteSystemMetricCountValue(nowUnix int64, name string, count, value float64, tagsList ...int32)
	WriteSystemMetricValueWithoutHost(nowUnix int64, name string, value float64, tagsList ...int32)

	WriteSystemMetricCountExtendedTag(nowUnix int64, name string, count float64, tagsList ...Tag)
	WriteSystemMetricCountValueExtendedTag(nowUnix int64, name string, count, value float64, tagsList ...Tag)
}

type MetricWriterRemoteImpl struct {
	HostName  string
	envLoader *env.Loader
}

// MetricWriterSHImpl isn't thread safe
type MetricWriterSHImpl struct {
	HostName  []byte
	handler   receiver.Handler
	metric    *tlstatshouse.MetricBytes
	envLoader *env.Loader
}

type Tag struct {
	Raw int32
	Str string
}

func buildTags[A Tag | int32](useHost bool, tags ...A) statshouse.Tags {
	res := statshouse.Tags{}
	// Tag1 is reserved for host
	for index, tagV := range tags {
		tagI := any(tagV)
		tag := ""
		switch t := tagI.(type) {
		case int32:
			tag = strconv.FormatInt(int64(t), 10)
		case Tag:
			if t.Str == "" {
				tag = strconv.FormatInt(int64(t.Raw), 10)
			} else {
				tag = t.Str
			}
		}
		i := index
		if useHost {
			i++
		}
		if i+1 < len(res) {
			res[i+1] = tag
		}
	}
	return res
}

const (
	reservedKeys = 2
	usedCommon   = 6 // envLoader, hostname, dc, group, region, owner
)

func fillTags[A Tag | int32](metric *tlstatshouse.MetricBytes, reservedKeysPrefix int, startFillFrom int, tags ...A) {
	// Tag1 is reserved for host
	for i, tag := range tags {
		i = i + reservedKeysPrefix
		t := &metric.Tags[startFillFrom]
		t.Key = strconv.AppendInt(t.Key[:0], int64(i), 10)
		tagI := any(tag)
		switch tagI := tagI.(type) {
		case int32:
			t.Value = strconv.AppendInt(t.Value[:0], int64(tagI), 10)
		case Tag:
			if tagI.Str == "" {
				t.Value = strconv.AppendInt(t.Value[:0], int64(tagI.Raw), 10)
			} else {
				t.Value = append(t.Value[:0], tagI.Str...)
			}
		}
		startFillFrom++
	}
}

func (p *MetricWriterRemoteImpl) fillCommonTags(tags *statshouse.Tags) {
	env := p.envLoader.Load()
	tags[0] = env.EnvT
	tags[1] = p.HostName
	tags[11] = env.DC
	tags[12] = env.Group
	tags[13] = env.Region
	tags[14] = env.Owner
}

func (p *MetricWriterRemoteImpl) WriteSystemMetricValue(nowUnix int64, name string, value float64, tagsList ...int32) {
	tags := buildTags(true, tagsList...)
	p.fillCommonTags(&tags)
	statshouse.Metric(name, tags).Value(value)
}

func (p *MetricWriterRemoteImpl) WriteSystemMetricCountValue(nowUnix int64, name string, count, value float64, tagsList ...int32) {
	tags := buildTags(true, tagsList...)
	p.fillCommonTags(&tags)
	statshouse.Metric(name, tags).Count(count)
	statshouse.Metric(name, tags).Value(value)
}

func (p *MetricWriterRemoteImpl) WriteSystemMetricValueWithoutHost(nowUnix int64, name string, value float64, tagsList ...int32) {
	tags := buildTags(false, tagsList...)
	statshouse.Metric(name, tags).Value(value)
}

func (p *MetricWriterRemoteImpl) WriteSystemMetricCount(nowUnix int64, name string, count float64, tagsList ...int32) {
	tags := buildTags(true, tagsList...)
	p.fillCommonTags(&tags)
	statshouse.Metric(name, tags).Count(count)
}

func (p *MetricWriterRemoteImpl) WriteSystemMetricCountValueExtendedTag(nowUnix int64, name string, count, value float64, tagsList ...Tag) {
	tags := buildTags(true, tagsList...)
	p.fillCommonTags(&tags)
	statshouse.Metric(name, tags).Count(count)
	statshouse.Metric(name, tags).Value(value)
}

func (p *MetricWriterRemoteImpl) WriteSystemMetricCountExtendedTag(nowUnix int64, name string, count float64, tagsList ...Tag) {
	tags := buildTags(true, tagsList...)
	p.fillCommonTags(&tags)
	statshouse.Metric(name, tags).Count(count)
}

func fillTag[T string | []byte](m *tlstatshouse.MetricBytes, i int, key string, value T) {
	m.Tags[i].Key = append(m.Tags[i].Key[:0], key...)
	m.Tags[i].Value = append(m.Tags[i].Value[:0], value...)
}

func fillCommonMetric[A Tag | int32](p *MetricWriterSHImpl, m *tlstatshouse.MetricBytes, useHostTags bool, name string, nowUnix int64, tagsList ...A) {
	m.Reset()
	m.Name = append(m.Name, name...)
	m.Ts = uint32(nowUnix)
	used := usedCommon
	reserved := reservedKeys
	if !useHostTags {
		used = 0
		reserved = 1
	}
	tagsLength := len(tagsList) + used
	if cap(m.Tags) < tagsLength {
		m.Tags = make([]tl.DictionaryFieldStringBytes, tagsLength)
	} else {
		m.Tags = m.Tags[:tagsLength]
	}
	env := p.envLoader.Load()
	if useHostTags {
		fillTag(m, 0, "0", env.EnvT)
		fillTag(m, 1, "1", p.HostName)
		fillTag(m, 2, "11", env.DC)
		fillTag(m, 3, "12", env.Group)
		fillTag(m, 4, "13", env.Region)
		fillTag(m, 5, "14", env.Owner)
	}
	fillTags(m, reserved, used, tagsList...)
}

func (p *MetricWriterSHImpl) WriteSystemMetricValue(nowUnix int64, name string, value float64, tagsList ...int32) {
	m := p.metric
	fillCommonMetric(p, m, true, name, nowUnix, tagsList...)
	m.Value = append(m.Value, value)
	_, _ = p.handler.HandleMetrics(data_model.HandlerArgs{MetricBytes: m})
}

func (p *MetricWriterSHImpl) WriteSystemMetricCountValue(nowUnix int64, name string, count, value float64, tagsList ...int32) {
	m := p.metric
	fillCommonMetric(p, m, true, name, nowUnix, tagsList...)
	m.Counter = count
	m.Value = append(m.Value, value)
	_, _ = p.handler.HandleMetrics(data_model.HandlerArgs{MetricBytes: m})
}

func (p *MetricWriterSHImpl) WriteSystemMetricValueWithoutHost(nowUnix int64, name string, value float64, tagsList ...int32) {
	m := p.metric
	fillCommonMetric(p, m, false, name, nowUnix, tagsList...)
	m.Value = append(m.Value, value)
	_, _ = p.handler.HandleMetrics(data_model.HandlerArgs{MetricBytes: m})
}

func (p *MetricWriterSHImpl) WriteSystemMetricCount(nowUnix int64, name string, count float64, tagsList ...int32) {
	m := p.metric
	fillCommonMetric(p, m, true, name, nowUnix, tagsList...)
	m.Counter = count
	_, _ = p.handler.HandleMetrics(data_model.HandlerArgs{MetricBytes: m})
}

func (p *MetricWriterSHImpl) WriteSystemMetricCountValueExtendedTag(nowUnix int64, name string, count, value float64, tagsList ...Tag) {
	m := p.metric
	fillCommonMetric(p, m, true, name, nowUnix, tagsList...)
	m.Counter = count
	m.Value = append(m.Value, value)
	_, _ = p.handler.HandleMetrics(data_model.HandlerArgs{MetricBytes: m})
}

func (p *MetricWriterSHImpl) WriteSystemMetricCountExtendedTag(nowUnix int64, name string, count float64, tagsList ...Tag) {
	m := p.metric
	fillCommonMetric(p, m, true, name, nowUnix, tagsList...)
	m.Counter = count
	_, _ = p.handler.HandleMetrics(data_model.HandlerArgs{MetricBytes: m})
}
