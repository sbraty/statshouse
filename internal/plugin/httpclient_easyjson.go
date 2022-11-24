// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

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

func easyjsonB4ad7d41DecodeGitlabMvkComGoVkgoProjectsStatshouseInternalPlugin(in *jlexer.Lexer, out *GetQueryResponse) {
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
		case "series":
			easyjsonB4ad7d41Decode(in, &out.Series)
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
func easyjsonB4ad7d41EncodeGitlabMvkComGoVkgoProjectsStatshouseInternalPlugin(out *jwriter.Writer, in GetQueryResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"series\":"
		out.RawString(prefix[1:])
		easyjsonB4ad7d41Encode(out, in.Series)
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetQueryResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB4ad7d41EncodeGitlabMvkComGoVkgoProjectsStatshouseInternalPlugin(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetQueryResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB4ad7d41DecodeGitlabMvkComGoVkgoProjectsStatshouseInternalPlugin(l, v)
}
func easyjsonB4ad7d41Decode(in *jlexer.Lexer, out *struct {
	Time       []int64                 `json:"time"`
	SeriesMeta []api.QuerySeriesMetaV2 `json:"series_meta"`
	SeriesData [][]*float64            `json:"series_data"`
}) {
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
		case "time":
			if in.IsNull() {
				in.Skip()
				out.Time = nil
			} else {
				in.Delim('[')
				if out.Time == nil {
					if !in.IsDelim(']') {
						out.Time = make([]int64, 0, 8)
					} else {
						out.Time = []int64{}
					}
				} else {
					out.Time = (out.Time)[:0]
				}
				for !in.IsDelim(']') {
					var v1 int64
					v1 = int64(in.Int64())
					out.Time = append(out.Time, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "series_meta":
			if in.IsNull() {
				in.Skip()
				out.SeriesMeta = nil
			} else {
				in.Delim('[')
				if out.SeriesMeta == nil {
					if !in.IsDelim(']') {
						out.SeriesMeta = make([]api.QuerySeriesMetaV2, 0, 1)
					} else {
						out.SeriesMeta = []api.QuerySeriesMetaV2{}
					}
				} else {
					out.SeriesMeta = (out.SeriesMeta)[:0]
				}
				for !in.IsDelim(']') {
					var v2 api.QuerySeriesMetaV2
					easyjsonB4ad7d41DecodeGitlabMvkComGoVkgoProjectsStatshouseInternalApi(in, &v2)
					out.SeriesMeta = append(out.SeriesMeta, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "series_data":
			if in.IsNull() {
				in.Skip()
				out.SeriesData = nil
			} else {
				in.Delim('[')
				if out.SeriesData == nil {
					if !in.IsDelim(']') {
						out.SeriesData = make([][]*float64, 0, 2)
					} else {
						out.SeriesData = [][]*float64{}
					}
				} else {
					out.SeriesData = (out.SeriesData)[:0]
				}
				for !in.IsDelim(']') {
					var v3 []*float64
					if in.IsNull() {
						in.Skip()
						v3 = nil
					} else {
						in.Delim('[')
						if v3 == nil {
							if !in.IsDelim(']') {
								v3 = make([]*float64, 0, 8)
							} else {
								v3 = []*float64{}
							}
						} else {
							v3 = (v3)[:0]
						}
						for !in.IsDelim(']') {
							var v4 *float64
							if in.IsNull() {
								in.Skip()
								v4 = nil
							} else {
								if v4 == nil {
									v4 = new(float64)
								}
								*v4 = float64(in.Float64())
							}
							v3 = append(v3, v4)
							in.WantComma()
						}
						in.Delim(']')
					}
					out.SeriesData = append(out.SeriesData, v3)
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
func easyjsonB4ad7d41Encode(out *jwriter.Writer, in struct {
	Time       []int64                 `json:"time"`
	SeriesMeta []api.QuerySeriesMetaV2 `json:"series_meta"`
	SeriesData [][]*float64            `json:"series_data"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"time\":"
		out.RawString(prefix[1:])
		if in.Time == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Time {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.Int64(int64(v6))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"series_meta\":"
		out.RawString(prefix)
		if in.SeriesMeta == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v7, v8 := range in.SeriesMeta {
				if v7 > 0 {
					out.RawByte(',')
				}
				easyjsonB4ad7d41EncodeGitlabMvkComGoVkgoProjectsStatshouseInternalApi(out, v8)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"series_data\":"
		out.RawString(prefix)
		if in.SeriesData == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v9, v10 := range in.SeriesData {
				if v9 > 0 {
					out.RawByte(',')
				}
				if v10 == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
					out.RawString("null")
				} else {
					out.RawByte('[')
					for v11, v12 := range v10 {
						if v11 > 0 {
							out.RawByte(',')
						}
						if v12 == nil {
							out.RawString("null")
						} else {
							out.Float64(float64(*v12))
						}
					}
					out.RawByte(']')
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonB4ad7d41DecodeGitlabMvkComGoVkgoProjectsStatshouseInternalApi(in *jlexer.Lexer, out *api.QuerySeriesMetaV2) {
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
		case "time_shift":
			out.TimeShift = int64(in.Int64())
		case "tags":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Tags = make(map[string]api.SeriesMetaTag)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v13 api.SeriesMetaTag
					easyjsonB4ad7d41DecodeGitlabMvkComGoVkgoProjectsStatshouseInternalApi1(in, &v13)
					(out.Tags)[key] = v13
					in.WantComma()
				}
				in.Delim('}')
			}
		case "max_hosts":
			if in.IsNull() {
				in.Skip()
				out.MaxHosts = nil
			} else {
				in.Delim('[')
				if out.MaxHosts == nil {
					if !in.IsDelim(']') {
						out.MaxHosts = make([]string, 0, 4)
					} else {
						out.MaxHosts = []string{}
					}
				} else {
					out.MaxHosts = (out.MaxHosts)[:0]
				}
				for !in.IsDelim(']') {
					var v14 string
					v14 = string(in.String())
					out.MaxHosts = append(out.MaxHosts, v14)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "what":
			(out.What).UnmarshalEasyJSON(in)
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
func easyjsonB4ad7d41EncodeGitlabMvkComGoVkgoProjectsStatshouseInternalApi(out *jwriter.Writer, in api.QuerySeriesMetaV2) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"time_shift\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.TimeShift))
	}
	{
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v15First := true
			for v15Name, v15Value := range in.Tags {
				if v15First {
					v15First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v15Name))
				out.RawByte(':')
				easyjsonB4ad7d41EncodeGitlabMvkComGoVkgoProjectsStatshouseInternalApi1(out, v15Value)
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"max_hosts\":"
		out.RawString(prefix)
		if in.MaxHosts == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v16, v17 := range in.MaxHosts {
				if v16 > 0 {
					out.RawByte(',')
				}
				out.String(string(v17))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"what\":"
		out.RawString(prefix)
		(in.What).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}
func easyjsonB4ad7d41DecodeGitlabMvkComGoVkgoProjectsStatshouseInternalApi1(in *jlexer.Lexer, out *api.SeriesMetaTag) {
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
		case "comment":
			out.Comment = string(in.String())
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
func easyjsonB4ad7d41EncodeGitlabMvkComGoVkgoProjectsStatshouseInternalApi1(out *jwriter.Writer, in api.SeriesMetaTag) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"value\":"
		out.RawString(prefix[1:])
		out.String(string(in.Value))
	}
	if in.Comment != "" {
		const prefix string = ",\"comment\":"
		out.RawString(prefix)
		out.String(string(in.Comment))
	}
	out.RawByte('}')
}
