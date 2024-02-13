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

type MetadataHistoryShortResponseEvent struct {
	Version  int64
	Metadata string
}

func (MetadataHistoryShortResponseEvent) TLName() string {
	return "metadata.history_short_response_event"
}
func (MetadataHistoryShortResponseEvent) TLTag() uint32 { return 0x1186baaf }

func (item *MetadataHistoryShortResponseEvent) Reset() {
	item.Version = 0
	item.Metadata = ""
}

func (item *MetadataHistoryShortResponseEvent) Read(w []byte, nat_field_mask uint32) (_ []byte, err error) {
	if w, err = basictl.LongRead(w, &item.Version); err != nil {
		return w, err
	}
	return basictl.StringRead(w, &item.Metadata)
}

func (item *MetadataHistoryShortResponseEvent) Write(w []byte, nat_field_mask uint32) (_ []byte, err error) {
	w = basictl.LongWrite(w, item.Version)
	return basictl.StringWrite(w, item.Metadata)
}

func (item *MetadataHistoryShortResponseEvent) ReadBoxed(w []byte, nat_field_mask uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1186baaf); err != nil {
		return w, err
	}
	return item.Read(w, nat_field_mask)
}

func (item *MetadataHistoryShortResponseEvent) WriteBoxed(w []byte, nat_field_mask uint32) ([]byte, error) {
	w = basictl.NatWrite(w, 0x1186baaf)
	return item.Write(w, nat_field_mask)
}

func MetadataHistoryShortResponseEvent__ReadJSON(item *MetadataHistoryShortResponseEvent, j interface{}, nat_field_mask uint32) error {
	return item.readJSON(j, nat_field_mask)
}
func (item *MetadataHistoryShortResponseEvent) readJSON(j interface{}, nat_field_mask uint32) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("metadata.history_short_response_event", "expected json object")
	}
	_jVersion := _jm["version"]
	delete(_jm, "version")
	if err := JsonReadInt64(_jVersion, &item.Version); err != nil {
		return err
	}
	_jMetadata := _jm["metadata"]
	delete(_jm, "metadata")
	if err := JsonReadString(_jMetadata, &item.Metadata); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("metadata.history_short_response_event", k)
	}
	return nil
}

func (item *MetadataHistoryShortResponseEvent) WriteJSON(w []byte, nat_field_mask uint32) (_ []byte, err error) {
	w = append(w, '{')
	if item.Version != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"version":`...)
		w = basictl.JSONWriteInt64(w, item.Version)
	}
	if len(item.Metadata) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"metadata":`...)
		w = basictl.JSONWriteString(w, item.Metadata)
	}
	return append(w, '}'), nil
}

func VectorMetadataHistoryShortResponseEvent0Read(w []byte, vec *[]MetadataHistoryShortResponseEvent, nat_t uint32) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]MetadataHistoryShortResponseEvent, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if w, err = (*vec)[i].Read(w, nat_t); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorMetadataHistoryShortResponseEvent0Write(w []byte, vec []MetadataHistoryShortResponseEvent, nat_t uint32) (_ []byte, err error) {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		if w, err = elem.Write(w, nat_t); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorMetadataHistoryShortResponseEvent0ReadJSON(j interface{}, vec *[]MetadataHistoryShortResponseEvent, nat_t uint32) error {
	l, _arr, err := JsonReadArray("[]MetadataHistoryShortResponseEvent", j)
	if err != nil {
		return err
	}
	if cap(*vec) < l {
		*vec = make([]MetadataHistoryShortResponseEvent, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if err := MetadataHistoryShortResponseEvent__ReadJSON(&(*vec)[i], _arr[i], nat_t); err != nil {
			return err
		}
	}
	return nil
}

func VectorMetadataHistoryShortResponseEvent0WriteJSON(w []byte, vec []MetadataHistoryShortResponseEvent, nat_t uint32) (_ []byte, err error) {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		if w, err = elem.WriteJSON(w, nat_t); err != nil {
			return w, err
		}
	}
	return append(w, ']'), nil
}