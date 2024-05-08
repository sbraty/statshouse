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

type StatshouseGetMetricsResult struct {
	Version string
	Metrics []string
}

func (StatshouseGetMetricsResult) TLName() string { return "statshouse.getMetricsResult" }
func (StatshouseGetMetricsResult) TLTag() uint32  { return 0xc803d05 }

func (item *StatshouseGetMetricsResult) Reset() {
	item.Version = ""
	item.Metrics = item.Metrics[:0]
}

func (item *StatshouseGetMetricsResult) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringRead(w, &item.Version); err != nil {
		return w, err
	}
	return BuiltinVectorStringRead(w, &item.Metrics)
}

func (item *StatshouseGetMetricsResult) Write(w []byte) (_ []byte, err error) {
	w = basictl.StringWrite(w, item.Version)
	return BuiltinVectorStringWrite(w, item.Metrics)
}

func (item *StatshouseGetMetricsResult) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xc803d05); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseGetMetricsResult) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xc803d05)
	return item.Write(w)
}

func (item StatshouseGetMetricsResult) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *StatshouseGetMetricsResult) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propVersionPresented bool
	var propMetricsPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "version":
				if propVersionPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.getMetricsResult", "version")
				}
				if err := Json2ReadString(in, &item.Version); err != nil {
					return err
				}
				propVersionPresented = true
			case "metrics":
				if propMetricsPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.getMetricsResult", "metrics")
				}
				if err := BuiltinVectorStringReadJSON(legacyTypeNames, in, &item.Metrics); err != nil {
					return err
				}
				propMetricsPresented = true
			default:
				return ErrorInvalidJSONExcessElement("statshouse.getMetricsResult", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propVersionPresented {
		item.Version = ""
	}
	if !propMetricsPresented {
		item.Metrics = item.Metrics[:0]
	}
	return nil
}

func (item *StatshouseGetMetricsResult) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *StatshouseGetMetricsResult) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexVersion := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"version":`...)
	w = basictl.JSONWriteString(w, item.Version)
	if (len(item.Version) != 0) == false {
		w = w[:backupIndexVersion]
	}
	backupIndexMetrics := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"metrics":`...)
	if w, err = BuiltinVectorStringWriteJSONOpt(newTypeNames, short, w, item.Metrics); err != nil {
		return w, err
	}
	if (len(item.Metrics) != 0) == false {
		w = w[:backupIndexMetrics]
	}
	return append(w, '}'), nil
}

func (item *StatshouseGetMetricsResult) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseGetMetricsResult) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("statshouse.getMetricsResult", err.Error())
	}
	return nil
}

type StatshouseGetMetricsResultBytes struct {
	Version []byte
	Metrics [][]byte
}

func (StatshouseGetMetricsResultBytes) TLName() string { return "statshouse.getMetricsResult" }
func (StatshouseGetMetricsResultBytes) TLTag() uint32  { return 0xc803d05 }

func (item *StatshouseGetMetricsResultBytes) Reset() {
	item.Version = item.Version[:0]
	item.Metrics = item.Metrics[:0]
}

func (item *StatshouseGetMetricsResultBytes) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringReadBytes(w, &item.Version); err != nil {
		return w, err
	}
	return BuiltinVectorStringBytesRead(w, &item.Metrics)
}

func (item *StatshouseGetMetricsResultBytes) Write(w []byte) (_ []byte, err error) {
	w = basictl.StringWriteBytes(w, item.Version)
	return BuiltinVectorStringBytesWrite(w, item.Metrics)
}

func (item *StatshouseGetMetricsResultBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xc803d05); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseGetMetricsResultBytes) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xc803d05)
	return item.Write(w)
}

func (item StatshouseGetMetricsResultBytes) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *StatshouseGetMetricsResultBytes) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propVersionPresented bool
	var propMetricsPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "version":
				if propVersionPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.getMetricsResult", "version")
				}
				if err := Json2ReadStringBytes(in, &item.Version); err != nil {
					return err
				}
				propVersionPresented = true
			case "metrics":
				if propMetricsPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.getMetricsResult", "metrics")
				}
				if err := BuiltinVectorStringBytesReadJSON(legacyTypeNames, in, &item.Metrics); err != nil {
					return err
				}
				propMetricsPresented = true
			default:
				return ErrorInvalidJSONExcessElement("statshouse.getMetricsResult", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propVersionPresented {
		item.Version = item.Version[:0]
	}
	if !propMetricsPresented {
		item.Metrics = item.Metrics[:0]
	}
	return nil
}

func (item *StatshouseGetMetricsResultBytes) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *StatshouseGetMetricsResultBytes) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexVersion := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"version":`...)
	w = basictl.JSONWriteStringBytes(w, item.Version)
	if (len(item.Version) != 0) == false {
		w = w[:backupIndexVersion]
	}
	backupIndexMetrics := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"metrics":`...)
	if w, err = BuiltinVectorStringBytesWriteJSONOpt(newTypeNames, short, w, item.Metrics); err != nil {
		return w, err
	}
	if (len(item.Metrics) != 0) == false {
		w = w[:backupIndexMetrics]
	}
	return append(w, '}'), nil
}

func (item *StatshouseGetMetricsResultBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseGetMetricsResultBytes) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("statshouse.getMetricsResult", err.Error())
	}
	return nil
}
