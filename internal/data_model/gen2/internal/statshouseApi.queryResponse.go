// Copyright 2022 V Kontakte LLC
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
}

func (StatshouseApiGetQueryResponse) TLName() string { return "statshouseApi.queryResponse" }
func (StatshouseApiGetQueryResponse) TLTag() uint32  { return 0x4487e49a }

func (item *StatshouseApiGetQueryResponse) Reset() {
	item.FieldsMask = 0
	item.Series.Reset()
	item.SeriesMeta = item.SeriesMeta[:0]
	item.ChunkIds = item.ChunkIds[:0]
	item.TotalTimePoints = 0
	item.ResponseId = 0
}

func (item *StatshouseApiGetQueryResponse) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = item.Series.Read(w); err != nil {
		return w, err
	}
	if w, err = VectorStatshouseApiSeriesMeta0Read(w, &item.SeriesMeta); err != nil {
		return w, err
	}
	if w, err = VectorInt0Read(w, &item.ChunkIds); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.TotalTimePoints); err != nil {
		return w, err
	}
	return basictl.LongRead(w, &item.ResponseId)
}

func (item *StatshouseApiGetQueryResponse) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	if w, err = item.Series.Write(w); err != nil {
		return w, err
	}
	if w, err = VectorStatshouseApiSeriesMeta0Write(w, item.SeriesMeta); err != nil {
		return w, err
	}
	if w, err = VectorInt0Write(w, item.ChunkIds); err != nil {
		return w, err
	}
	w = basictl.IntWrite(w, item.TotalTimePoints)
	return basictl.LongWrite(w, item.ResponseId), nil
}

func (item *StatshouseApiGetQueryResponse) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x4487e49a); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseApiGetQueryResponse) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x4487e49a)
	return item.Write(w)
}

func (item StatshouseApiGetQueryResponse) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func StatshouseApiGetQueryResponse__ReadJSON(item *StatshouseApiGetQueryResponse, j interface{}) error {
	return item.readJSON(j)
}
func (item *StatshouseApiGetQueryResponse) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouseApi.queryResponse", "expected json object")
	}
	_jFieldsMask := _jm["fields_mask"]
	delete(_jm, "fields_mask")
	if err := JsonReadUint32(_jFieldsMask, &item.FieldsMask); err != nil {
		return err
	}
	_jSeries := _jm["series"]
	delete(_jm, "series")
	_jSeriesMeta := _jm["series_meta"]
	delete(_jm, "series_meta")
	_jChunkIds := _jm["chunk_ids"]
	delete(_jm, "chunk_ids")
	_jTotalTimePoints := _jm["total_time_points"]
	delete(_jm, "total_time_points")
	if err := JsonReadInt32(_jTotalTimePoints, &item.TotalTimePoints); err != nil {
		return err
	}
	_jResponseId := _jm["response_id"]
	delete(_jm, "response_id")
	if err := JsonReadInt64(_jResponseId, &item.ResponseId); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouseApi.queryResponse", k)
	}
	if err := StatshouseApiSeries__ReadJSON(&item.Series, _jSeries); err != nil {
		return err
	}
	if err := VectorStatshouseApiSeriesMeta0ReadJSON(_jSeriesMeta, &item.SeriesMeta); err != nil {
		return err
	}
	if err := VectorInt0ReadJSON(_jChunkIds, &item.ChunkIds); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseApiGetQueryResponse) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldsMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"fields_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldsMask)
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"series":`...)
	if w, err = item.Series.WriteJSON(w); err != nil {
		return w, err
	}
	if len(item.SeriesMeta) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"series_meta":`...)
		if w, err = VectorStatshouseApiSeriesMeta0WriteJSON(w, item.SeriesMeta); err != nil {
			return w, err
		}
	}
	if len(item.ChunkIds) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"chunk_ids":`...)
		if w, err = VectorInt0WriteJSON(w, item.ChunkIds); err != nil {
			return w, err
		}
	}
	if item.TotalTimePoints != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"total_time_points":`...)
		w = basictl.JSONWriteInt32(w, item.TotalTimePoints)
	}
	if item.ResponseId != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"response_id":`...)
		w = basictl.JSONWriteInt64(w, item.ResponseId)
	}
	return append(w, '}'), nil
}

func (item *StatshouseApiGetQueryResponse) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseApiGetQueryResponse) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("statshouseApi.queryResponse", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("statshouseApi.queryResponse", err.Error())
	}
	return nil
}
