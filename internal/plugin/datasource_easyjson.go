// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package plugin

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson4e226ef3DecodeGithubComVkcomStatshouseInternalPlugin(in *jlexer.Lexer, out *pluginProperties) {
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
		case "apiURL":
			out.ApiURL = string(in.String())
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
func easyjson4e226ef3EncodeGithubComVkcomStatshouseInternalPlugin(out *jwriter.Writer, in pluginProperties) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"apiURL\":"
		out.RawString(prefix[1:])
		out.String(string(in.ApiURL))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v pluginProperties) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4e226ef3EncodeGithubComVkcomStatshouseInternalPlugin(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *pluginProperties) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4e226ef3DecodeGithubComVkcomStatshouseInternalPlugin(l, v)
}
