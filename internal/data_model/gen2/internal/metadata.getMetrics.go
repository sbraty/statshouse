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

type MetadataGetMetrics struct {
	FieldMask uint32
	From      int64
	Limit     int64
}

func (MetadataGetMetrics) TLName() string { return "metadata.getMetrics" }
func (MetadataGetMetrics) TLTag() uint32  { return 0x93ba92f5 }

func (item *MetadataGetMetrics) Reset() {
	item.FieldMask = 0
	item.From = 0
	item.Limit = 0
}

func (item *MetadataGetMetrics) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldMask); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.From); err != nil {
		return w, err
	}
	return basictl.LongRead(w, &item.Limit)
}

// This method is general version of Write, use it instead!
func (item *MetadataGetMetrics) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *MetadataGetMetrics) Write(w []byte) []byte {
	w = basictl.NatWrite(w, item.FieldMask)
	w = basictl.LongWrite(w, item.From)
	w = basictl.LongWrite(w, item.Limit)
	return w
}

func (item *MetadataGetMetrics) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x93ba92f5); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *MetadataGetMetrics) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *MetadataGetMetrics) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x93ba92f5)
	return item.Write(w)
}

func (item *MetadataGetMetrics) ReadResult(w []byte, ret *MetadataGetMetricsResponse) (_ []byte, err error) {
	return ret.ReadBoxed(w, item.FieldMask)
}

func (item *MetadataGetMetrics) WriteResult(w []byte, ret MetadataGetMetricsResponse) (_ []byte, err error) {
	w = ret.WriteBoxed(w, item.FieldMask)
	return w, nil
}

func (item *MetadataGetMetrics) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *MetadataGetMetricsResponse) error {
	if err := ret.ReadJSON(legacyTypeNames, in, item.FieldMask); err != nil {
		return err
	}
	return nil
}

func (item *MetadataGetMetrics) WriteResultJSON(w []byte, ret MetadataGetMetricsResponse) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *MetadataGetMetrics) writeResultJSON(newTypeNames bool, short bool, w []byte, ret MetadataGetMetricsResponse) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w, item.FieldMask)
	return w, nil
}

func (item *MetadataGetMetrics) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret MetadataGetMetricsResponse
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *MetadataGetMetrics) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret MetadataGetMetricsResponse
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *MetadataGetMetrics) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret MetadataGetMetricsResponse
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item MetadataGetMetrics) String() string {
	return string(item.WriteJSON(nil))
}

func (item *MetadataGetMetrics) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFieldMaskPresented bool
	var propFromPresented bool
	var propLimitPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "field_mask":
				if propFieldMaskPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("metadata.getMetrics", "field_mask")
				}
				if err := Json2ReadUint32(in, &item.FieldMask); err != nil {
					return err
				}
				propFieldMaskPresented = true
			case "from":
				if propFromPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("metadata.getMetrics", "from")
				}
				if err := Json2ReadInt64(in, &item.From); err != nil {
					return err
				}
				propFromPresented = true
			case "limit":
				if propLimitPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("metadata.getMetrics", "limit")
				}
				if err := Json2ReadInt64(in, &item.Limit); err != nil {
					return err
				}
				propLimitPresented = true
			default:
				return ErrorInvalidJSONExcessElement("metadata.getMetrics", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propFieldMaskPresented {
		item.FieldMask = 0
	}
	if !propFromPresented {
		item.From = 0
	}
	if !propLimitPresented {
		item.Limit = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *MetadataGetMetrics) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *MetadataGetMetrics) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *MetadataGetMetrics) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexFieldMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"field_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldMask)
	if (item.FieldMask != 0) == false {
		w = w[:backupIndexFieldMask]
	}
	backupIndexFrom := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"from":`...)
	w = basictl.JSONWriteInt64(w, item.From)
	if (item.From != 0) == false {
		w = w[:backupIndexFrom]
	}
	backupIndexLimit := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"limit":`...)
	w = basictl.JSONWriteInt64(w, item.Limit)
	if (item.Limit != 0) == false {
		w = w[:backupIndexLimit]
	}
	return append(w, '}')
}

func (item *MetadataGetMetrics) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *MetadataGetMetrics) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("metadata.getMetrics", err.Error())
	}
	return nil
}
