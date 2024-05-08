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

type StatshouseAddMetricsBatch struct {
	FieldsMask uint32
	Metrics    []StatshouseMetric
}

func (StatshouseAddMetricsBatch) TLName() string { return "statshouse.addMetricsBatch" }
func (StatshouseAddMetricsBatch) TLTag() uint32  { return 0x56580239 }

func (item *StatshouseAddMetricsBatch) Reset() {
	item.FieldsMask = 0
	item.Metrics = item.Metrics[:0]
}

func (item *StatshouseAddMetricsBatch) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	return BuiltinVectorStatshouseMetricRead(w, &item.Metrics)
}

func (item *StatshouseAddMetricsBatch) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	return BuiltinVectorStatshouseMetricWrite(w, item.Metrics)
}

func (item *StatshouseAddMetricsBatch) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x56580239); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseAddMetricsBatch) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x56580239)
	return item.Write(w)
}

func (item *StatshouseAddMetricsBatch) ReadResult(w []byte, ret *True) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *StatshouseAddMetricsBatch) WriteResult(w []byte, ret True) (_ []byte, err error) {
	return ret.WriteBoxed(w)
}

func (item *StatshouseAddMetricsBatch) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *True) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseAddMetricsBatch) WriteResultJSON(w []byte, ret True) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *StatshouseAddMetricsBatch) writeResultJSON(newTypeNames bool, short bool, w []byte, ret True) (_ []byte, err error) {
	if w, err = ret.WriteJSONOpt(newTypeNames, short, w); err != nil {
		return w, err
	}
	return w, nil
}

func (item *StatshouseAddMetricsBatch) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *StatshouseAddMetricsBatch) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *StatshouseAddMetricsBatch) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret True
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item StatshouseAddMetricsBatch) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *StatshouseAddMetricsBatch) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFieldsMaskPresented bool
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
			case "fields_mask":
				if propFieldsMaskPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.addMetricsBatch", "fields_mask")
				}
				if err := Json2ReadUint32(in, &item.FieldsMask); err != nil {
					return err
				}
				propFieldsMaskPresented = true
			case "metrics":
				if propMetricsPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.addMetricsBatch", "metrics")
				}
				if err := BuiltinVectorStatshouseMetricReadJSON(legacyTypeNames, in, &item.Metrics); err != nil {
					return err
				}
				propMetricsPresented = true
			default:
				return ErrorInvalidJSONExcessElement("statshouse.addMetricsBatch", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propFieldsMaskPresented {
		item.FieldsMask = 0
	}
	if !propMetricsPresented {
		item.Metrics = item.Metrics[:0]
	}
	return nil
}

func (item *StatshouseAddMetricsBatch) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *StatshouseAddMetricsBatch) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexFieldsMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"fields_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldsMask)
	if (item.FieldsMask != 0) == false {
		w = w[:backupIndexFieldsMask]
	}
	backupIndexMetrics := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"metrics":`...)
	if w, err = BuiltinVectorStatshouseMetricWriteJSONOpt(newTypeNames, short, w, item.Metrics); err != nil {
		return w, err
	}
	if (len(item.Metrics) != 0) == false {
		w = w[:backupIndexMetrics]
	}
	return append(w, '}'), nil
}

func (item *StatshouseAddMetricsBatch) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseAddMetricsBatch) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("statshouse.addMetricsBatch", err.Error())
	}
	return nil
}

type StatshouseAddMetricsBatchBytes struct {
	FieldsMask uint32
	Metrics    []StatshouseMetricBytes
}

func (StatshouseAddMetricsBatchBytes) TLName() string { return "statshouse.addMetricsBatch" }
func (StatshouseAddMetricsBatchBytes) TLTag() uint32  { return 0x56580239 }

func (item *StatshouseAddMetricsBatchBytes) Reset() {
	item.FieldsMask = 0
	item.Metrics = item.Metrics[:0]
}

func (item *StatshouseAddMetricsBatchBytes) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	return BuiltinVectorStatshouseMetricBytesRead(w, &item.Metrics)
}

func (item *StatshouseAddMetricsBatchBytes) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	return BuiltinVectorStatshouseMetricBytesWrite(w, item.Metrics)
}

func (item *StatshouseAddMetricsBatchBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x56580239); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseAddMetricsBatchBytes) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x56580239)
	return item.Write(w)
}

func (item *StatshouseAddMetricsBatchBytes) ReadResult(w []byte, ret *True) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *StatshouseAddMetricsBatchBytes) WriteResult(w []byte, ret True) (_ []byte, err error) {
	return ret.WriteBoxed(w)
}

func (item *StatshouseAddMetricsBatchBytes) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *True) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseAddMetricsBatchBytes) WriteResultJSON(w []byte, ret True) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *StatshouseAddMetricsBatchBytes) writeResultJSON(newTypeNames bool, short bool, w []byte, ret True) (_ []byte, err error) {
	if w, err = ret.WriteJSONOpt(newTypeNames, short, w); err != nil {
		return w, err
	}
	return w, nil
}

func (item *StatshouseAddMetricsBatchBytes) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *StatshouseAddMetricsBatchBytes) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *StatshouseAddMetricsBatchBytes) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret True
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item StatshouseAddMetricsBatchBytes) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *StatshouseAddMetricsBatchBytes) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFieldsMaskPresented bool
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
			case "fields_mask":
				if propFieldsMaskPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.addMetricsBatch", "fields_mask")
				}
				if err := Json2ReadUint32(in, &item.FieldsMask); err != nil {
					return err
				}
				propFieldsMaskPresented = true
			case "metrics":
				if propMetricsPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.addMetricsBatch", "metrics")
				}
				if err := BuiltinVectorStatshouseMetricBytesReadJSON(legacyTypeNames, in, &item.Metrics); err != nil {
					return err
				}
				propMetricsPresented = true
			default:
				return ErrorInvalidJSONExcessElement("statshouse.addMetricsBatch", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propFieldsMaskPresented {
		item.FieldsMask = 0
	}
	if !propMetricsPresented {
		item.Metrics = item.Metrics[:0]
	}
	return nil
}

func (item *StatshouseAddMetricsBatchBytes) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *StatshouseAddMetricsBatchBytes) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexFieldsMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"fields_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldsMask)
	if (item.FieldsMask != 0) == false {
		w = w[:backupIndexFieldsMask]
	}
	backupIndexMetrics := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"metrics":`...)
	if w, err = BuiltinVectorStatshouseMetricBytesWriteJSONOpt(newTypeNames, short, w, item.Metrics); err != nil {
		return w, err
	}
	if (len(item.Metrics) != 0) == false {
		w = w[:backupIndexMetrics]
	}
	return append(w, '}'), nil
}

func (item *StatshouseAddMetricsBatchBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseAddMetricsBatchBytes) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("statshouse.addMetricsBatch", err.Error())
	}
	return nil
}
