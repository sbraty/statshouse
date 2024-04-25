// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package plugin

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	api "github.com/vkcom/statshouse/internal/api"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonA2f992c4DecodeGithubComVkcomStatshouseInternalPlugin(in *jlexer.Lexer, out *metricTagValuesResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "tag_values":
			if in.IsNull() {
				in.Skip()
				out.TagValues = nil
			} else {
				in.Delim('[')
				if out.TagValues == nil {
					if !in.IsDelim(']') {
						out.TagValues = make([]api.MetricTagValueInfo, 0, 2)
					} else {
						out.TagValues = []api.MetricTagValueInfo{}
					}
				} else {
					out.TagValues = (out.TagValues)[:0]
				}
				for !in.IsDelim(']') {
					var v1 api.MetricTagValueInfo
					easyjsonA2f992c4DecodeGithubComVkcomStatshouseInternalApi(in, &v1)
					out.TagValues = append(out.TagValues, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA2f992c4EncodeGithubComVkcomStatshouseInternalPlugin(out *jwriter.Writer, in metricTagValuesResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"tag_values\":"
		out.RawString(prefix[1:])
		if in.TagValues == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.TagValues {
				if v2 > 0 {
					out.RawByte(',')
				}
				easyjsonA2f992c4EncodeGithubComVkcomStatshouseInternalApi(out, v3)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v metricTagValuesResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA2f992c4EncodeGithubComVkcomStatshouseInternalPlugin(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *metricTagValuesResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA2f992c4DecodeGithubComVkcomStatshouseInternalPlugin(l, v)
}
func easyjsonA2f992c4DecodeGithubComVkcomStatshouseInternalApi(in *jlexer.Lexer, out *api.MetricTagValueInfo) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "value":
			out.Value = string(in.String())
		case "count":
			out.Count = float64(in.Float64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA2f992c4EncodeGithubComVkcomStatshouseInternalApi(out *jwriter.Writer, in api.MetricTagValueInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"value\":"
		out.RawString(prefix[1:])
		out.String(string(in.Value))
	}
	{
		const prefix string = ",\"count\":"
		out.RawString(prefix)
		out.Float64(float64(in.Count))
	}
	out.RawByte('}')
}
func easyjsonA2f992c4DecodeGithubComVkcomStatshouseInternalPlugin1(in *jlexer.Lexer, out *metricResourceResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "functions":
			if in.IsNull() {
				in.Skip()
				out.Functions = nil
			} else {
				in.Delim('[')
				if out.Functions == nil {
					if !in.IsDelim(']') {
						out.Functions = make([]string, 0, 4)
					} else {
						out.Functions = []string{}
					}
				} else {
					out.Functions = (out.Functions)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Functions = append(out.Functions, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]Tag, 0, 1)
					} else {
						out.Tags = []Tag{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v5 Tag
					easyjsonA2f992c4DecodeGithubComVkcomStatshouseInternalPlugin2(in, &v5)
					out.Tags = append(out.Tags, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA2f992c4EncodeGithubComVkcomStatshouseInternalPlugin1(out *jwriter.Writer, in metricResourceResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"functions\":"
		out.RawString(prefix[1:])
		if in.Functions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v6, v7 := range in.Functions {
				if v6 > 0 {
					out.RawByte(',')
				}
				out.String(string(v7))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Tags {
				if v8 > 0 {
					out.RawByte(',')
				}
				easyjsonA2f992c4EncodeGithubComVkcomStatshouseInternalPlugin2(out, v9)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v metricResourceResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA2f992c4EncodeGithubComVkcomStatshouseInternalPlugin1(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *metricResourceResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA2f992c4DecodeGithubComVkcomStatshouseInternalPlugin1(l, v)
}
func easyjsonA2f992c4DecodeGithubComVkcomStatshouseInternalPlugin2(in *jlexer.Lexer, out *Tag) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "is_raw":
			out.IsRaw = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA2f992c4EncodeGithubComVkcomStatshouseInternalPlugin2(out *jwriter.Writer, in Tag) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	if in.Description != "" {
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"is_raw\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsRaw))
	}
	out.RawByte('}')
}
func easyjsonA2f992c4DecodeGithubComVkcomStatshouseInternalPlugin3(in *jlexer.Lexer, out *metricNamesResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "metric_names":
			if in.IsNull() {
				in.Skip()
				out.MetricNames = nil
			} else {
				in.Delim('[')
				if out.MetricNames == nil {
					if !in.IsDelim(']') {
						out.MetricNames = make([]string, 0, 4)
					} else {
						out.MetricNames = []string{}
					}
				} else {
					out.MetricNames = (out.MetricNames)[:0]
				}
				for !in.IsDelim(']') {
					var v10 string
					v10 = string(in.String())
					out.MetricNames = append(out.MetricNames, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA2f992c4EncodeGithubComVkcomStatshouseInternalPlugin3(out *jwriter.Writer, in metricNamesResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"metric_names\":"
		out.RawString(prefix[1:])
		if in.MetricNames == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.MetricNames {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.String(string(v12))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v metricNamesResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA2f992c4EncodeGithubComVkcomStatshouseInternalPlugin3(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *metricNamesResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA2f992c4DecodeGithubComVkcomStatshouseInternalPlugin3(l, v)
}
