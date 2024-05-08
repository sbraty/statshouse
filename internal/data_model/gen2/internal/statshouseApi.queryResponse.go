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

type StatshouseApiGetQueryResponse struct {
	FieldsMask      uint32
	Series          StatshouseApiSeries
	SeriesMeta      []StatshouseApiSeriesMeta
	ChunkIds        []int32
	TotalTimePoints int32
	ResponseId      int64
	// ExcessPointLeft (TrueType) // Conditional: item.FieldsMask.0
	// ExcessPointRight (TrueType) // Conditional: item.FieldsMask.1
}

func (StatshouseApiGetQueryResponse) TLName() string { return "statshouseApi.queryResponse" }
func (StatshouseApiGetQueryResponse) TLTag() uint32  { return 0x4487e49a }

func (item *StatshouseApiGetQueryResponse) SetExcessPointLeft(v bool) {
	if v {
		item.FieldsMask |= 1 << 0
	} else {
		item.FieldsMask &^= 1 << 0
	}
}
func (item StatshouseApiGetQueryResponse) IsSetExcessPointLeft() bool {
	return item.FieldsMask&(1<<0) != 0
}

func (item *StatshouseApiGetQueryResponse) SetExcessPointRight(v bool) {
	if v {
		item.FieldsMask |= 1 << 1
	} else {
		item.FieldsMask &^= 1 << 1
	}
}
func (item StatshouseApiGetQueryResponse) IsSetExcessPointRight() bool {
	return item.FieldsMask&(1<<1) != 0
}

func (item *StatshouseApiGetQueryResponse) Reset() {
	item.FieldsMask = 0
	item.Series.Reset()
	item.SeriesMeta = item.SeriesMeta[:0]
	item.ChunkIds = item.ChunkIds[:0]
	item.TotalTimePoints = 0
	item.ResponseId = 0
}

func (item *StatshouseApiGetQueryResponse) Read(w []byte, nat_query_fields_mask uint32) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = item.Series.Read(w); err != nil {
		return w, err
	}
	if w, err = BuiltinVectorStatshouseApiSeriesMetaRead(w, &item.SeriesMeta, nat_query_fields_mask); err != nil {
		return w, err
	}
	if w, err = BuiltinVectorIntRead(w, &item.ChunkIds); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.TotalTimePoints); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.ResponseId); err != nil {
		return w, err
	}
	return w, nil
}

func (item *StatshouseApiGetQueryResponse) Write(w []byte, nat_query_fields_mask uint32) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	if w, err = item.Series.Write(w); err != nil {
		return w, err
	}
	if w, err = BuiltinVectorStatshouseApiSeriesMetaWrite(w, item.SeriesMeta, nat_query_fields_mask); err != nil {
		return w, err
	}
	if w, err = BuiltinVectorIntWrite(w, item.ChunkIds); err != nil {
		return w, err
	}
	w = basictl.IntWrite(w, item.TotalTimePoints)
	w = basictl.LongWrite(w, item.ResponseId)
	return w, nil
}

func (item *StatshouseApiGetQueryResponse) ReadBoxed(w []byte, nat_query_fields_mask uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x4487e49a); err != nil {
		return w, err
	}
	return item.Read(w, nat_query_fields_mask)
}

func (item *StatshouseApiGetQueryResponse) WriteBoxed(w []byte, nat_query_fields_mask uint32) ([]byte, error) {
	w = basictl.NatWrite(w, 0x4487e49a)
	return item.Write(w, nat_query_fields_mask)
}

func (item *StatshouseApiGetQueryResponse) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_query_fields_mask uint32) error {
	var propFieldsMaskPresented bool
	var propSeriesPresented bool
	var rawSeriesMeta []byte
	var propChunkIdsPresented bool
	var propTotalTimePointsPresented bool
	var propResponseIdPresented bool
	var trueTypeExcessPointLeftPresented bool
	var trueTypeExcessPointLeftValue bool
	var trueTypeExcessPointRightPresented bool
	var trueTypeExcessPointRightValue bool

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
					return ErrorInvalidJSONWithDuplicatingKeys("statshouseApi.queryResponse", "fields_mask")
				}
				if err := Json2ReadUint32(in, &item.FieldsMask); err != nil {
					return err
				}
				propFieldsMaskPresented = true
			case "series":
				if propSeriesPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouseApi.queryResponse", "series")
				}
				if err := item.Series.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propSeriesPresented = true
			case "series_meta":
				if rawSeriesMeta != nil {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouseApi.queryResponse", "series_meta")
				}
				rawSeriesMeta = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			case "chunk_ids":
				if propChunkIdsPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouseApi.queryResponse", "chunk_ids")
				}
				if err := BuiltinVectorIntReadJSON(legacyTypeNames, in, &item.ChunkIds); err != nil {
					return err
				}
				propChunkIdsPresented = true
			case "total_time_points":
				if propTotalTimePointsPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouseApi.queryResponse", "total_time_points")
				}
				if err := Json2ReadInt32(in, &item.TotalTimePoints); err != nil {
					return err
				}
				propTotalTimePointsPresented = true
			case "response_id":
				if propResponseIdPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouseApi.queryResponse", "response_id")
				}
				if err := Json2ReadInt64(in, &item.ResponseId); err != nil {
					return err
				}
				propResponseIdPresented = true
			case "excess_point_left":
				if trueTypeExcessPointLeftPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouseApi.queryResponse", "excess_point_left")
				}
				if err := Json2ReadBool(in, &trueTypeExcessPointLeftValue); err != nil {
					return err
				}
				trueTypeExcessPointLeftPresented = true
			case "excess_point_right":
				if trueTypeExcessPointRightPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouseApi.queryResponse", "excess_point_right")
				}
				if err := Json2ReadBool(in, &trueTypeExcessPointRightValue); err != nil {
					return err
				}
				trueTypeExcessPointRightPresented = true
			default:
				return ErrorInvalidJSONExcessElement("statshouseApi.queryResponse", key)
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
	if !propSeriesPresented {
		item.Series.Reset()
	}
	if !propChunkIdsPresented {
		item.ChunkIds = item.ChunkIds[:0]
	}
	if !propTotalTimePointsPresented {
		item.TotalTimePoints = 0
	}
	if !propResponseIdPresented {
		item.ResponseId = 0
	}
	if trueTypeExcessPointLeftPresented {
		if trueTypeExcessPointLeftValue {
			item.FieldsMask |= 1 << 0
		}
	}
	if trueTypeExcessPointRightPresented {
		if trueTypeExcessPointRightValue {
			item.FieldsMask |= 1 << 1
		}
	}
	var inSeriesMetaPointer *basictl.JsonLexer
	inSeriesMeta := basictl.JsonLexer{Data: rawSeriesMeta}
	if rawSeriesMeta != nil {
		inSeriesMetaPointer = &inSeriesMeta
	}
	if err := BuiltinVectorStatshouseApiSeriesMetaReadJSON(legacyTypeNames, inSeriesMetaPointer, &item.SeriesMeta, nat_query_fields_mask); err != nil {
		return err
	}

	// tries to set bit to zero if it is 1
	if trueTypeExcessPointLeftPresented && !trueTypeExcessPointLeftValue && (item.FieldsMask&(1<<0) != 0) {
		return ErrorInvalidJSON("statshouseApi.queryResponse", "fieldmask bit fields_mask.0 is indefinite because of the contradictions in values")
	}
	// tries to set bit to zero if it is 1
	if trueTypeExcessPointRightPresented && !trueTypeExcessPointRightValue && (item.FieldsMask&(1<<1) != 0) {
		return ErrorInvalidJSON("statshouseApi.queryResponse", "fieldmask bit fields_mask.0 is indefinite because of the contradictions in values")
	}
	return nil
}

func (item *StatshouseApiGetQueryResponse) WriteJSON(w []byte, nat_query_fields_mask uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_query_fields_mask)
}
func (item *StatshouseApiGetQueryResponse) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_query_fields_mask uint32) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexFieldsMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"fields_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldsMask)
	if (item.FieldsMask != 0) == false {
		w = w[:backupIndexFieldsMask]
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"series":`...)
	if w, err = item.Series.WriteJSONOpt(newTypeNames, short, w); err != nil {
		return w, err
	}
	backupIndexSeriesMeta := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"series_meta":`...)
	if w, err = BuiltinVectorStatshouseApiSeriesMetaWriteJSONOpt(newTypeNames, short, w, item.SeriesMeta, nat_query_fields_mask); err != nil {
		return w, err
	}
	if (len(item.SeriesMeta) != 0) == false {
		w = w[:backupIndexSeriesMeta]
	}
	backupIndexChunkIds := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"chunk_ids":`...)
	if w, err = BuiltinVectorIntWriteJSONOpt(newTypeNames, short, w, item.ChunkIds); err != nil {
		return w, err
	}
	if (len(item.ChunkIds) != 0) == false {
		w = w[:backupIndexChunkIds]
	}
	backupIndexTotalTimePoints := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"total_time_points":`...)
	w = basictl.JSONWriteInt32(w, item.TotalTimePoints)
	if (item.TotalTimePoints != 0) == false {
		w = w[:backupIndexTotalTimePoints]
	}
	backupIndexResponseId := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"response_id":`...)
	w = basictl.JSONWriteInt64(w, item.ResponseId)
	if (item.ResponseId != 0) == false {
		w = w[:backupIndexResponseId]
	}
	if item.FieldsMask&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"excess_point_left":true`...)
	}
	if item.FieldsMask&(1<<1) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"excess_point_right":true`...)
	}
	return append(w, '}'), nil
}
