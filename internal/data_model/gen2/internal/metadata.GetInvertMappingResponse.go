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

func (item MetadataGetInvertMappingResponse) AsUnion() MetadataGetInvertMappingResponseUnion {
	var ret MetadataGetInvertMappingResponseUnion
	ret.SetGetInvertMappingResponse(item)
	return ret
}

// AsUnion will be here
type MetadataGetInvertMappingResponse struct {
	Key string
}

func (MetadataGetInvertMappingResponse) TLName() string { return "metadata.getInvertMappingResponse" }
func (MetadataGetInvertMappingResponse) TLTag() uint32  { return 0x9286abac }

func (item *MetadataGetInvertMappingResponse) Reset() {
	item.Key = ""
}

func (item *MetadataGetInvertMappingResponse) Read(w []byte, nat_field_mask uint32) (_ []byte, err error) {
	return basictl.StringRead(w, &item.Key)
}

func (item *MetadataGetInvertMappingResponse) Write(w []byte, nat_field_mask uint32) (_ []byte, err error) {
	return basictl.StringWrite(w, item.Key)
}

func (item *MetadataGetInvertMappingResponse) ReadBoxed(w []byte, nat_field_mask uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9286abac); err != nil {
		return w, err
	}
	return item.Read(w, nat_field_mask)
}

func (item *MetadataGetInvertMappingResponse) WriteBoxed(w []byte, nat_field_mask uint32) ([]byte, error) {
	w = basictl.NatWrite(w, 0x9286abac)
	return item.Write(w, nat_field_mask)
}

func MetadataGetInvertMappingResponse__ReadJSON(item *MetadataGetInvertMappingResponse, j interface{}, nat_field_mask uint32) error {
	return item.readJSON(j, nat_field_mask)
}
func (item *MetadataGetInvertMappingResponse) readJSON(j interface{}, nat_field_mask uint32) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("metadata.getInvertMappingResponse", "expected json object")
	}
	_jKey := _jm["key"]
	delete(_jm, "key")
	if err := JsonReadString(_jKey, &item.Key); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("metadata.getInvertMappingResponse", k)
	}
	return nil
}

func (item *MetadataGetInvertMappingResponse) WriteJSON(w []byte, nat_field_mask uint32) (_ []byte, err error) {
	w = append(w, '{')
	if len(item.Key) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"key":`...)
		w = basictl.JSONWriteString(w, item.Key)
	}
	return append(w, '}'), nil
}

func (item MetadataGetInvertMappingResponseKeyNotExists) AsUnion() MetadataGetInvertMappingResponseUnion {
	var ret MetadataGetInvertMappingResponseUnion
	ret.SetKeyNotExists()
	return ret
}

// AsUnion will be here
type MetadataGetInvertMappingResponseKeyNotExists struct {
}

func (MetadataGetInvertMappingResponseKeyNotExists) TLName() string {
	return "metadata.getInvertMappingResponseKeyNotExists"
}
func (MetadataGetInvertMappingResponseKeyNotExists) TLTag() uint32 { return 0x9286abab }

func (item *MetadataGetInvertMappingResponseKeyNotExists) Reset() {}
func (item *MetadataGetInvertMappingResponseKeyNotExists) Read(w []byte, nat_field_mask uint32) ([]byte, error) {
	return w, nil
}
func (item *MetadataGetInvertMappingResponseKeyNotExists) Write(w []byte, nat_field_mask uint32) ([]byte, error) {
	return w, nil
}
func (item *MetadataGetInvertMappingResponseKeyNotExists) ReadBoxed(w []byte, nat_field_mask uint32) ([]byte, error) {
	return basictl.NatReadExactTag(w, 0x9286abab)
}
func (item *MetadataGetInvertMappingResponseKeyNotExists) WriteBoxed(w []byte, nat_field_mask uint32) ([]byte, error) {
	return basictl.NatWrite(w, 0x9286abab), nil
}

func MetadataGetInvertMappingResponseKeyNotExists__ReadJSON(item *MetadataGetInvertMappingResponseKeyNotExists, j interface{}, nat_field_mask uint32) error {
	return item.readJSON(j, nat_field_mask)
}
func (item *MetadataGetInvertMappingResponseKeyNotExists) readJSON(j interface{}, nat_field_mask uint32) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("metadata.getInvertMappingResponseKeyNotExists", "expected json object")
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("metadata.getInvertMappingResponseKeyNotExists", k)
	}
	return nil
}

func (item *MetadataGetInvertMappingResponseKeyNotExists) WriteJSON(w []byte, nat_field_mask uint32) (_ []byte, err error) {
	w = append(w, '{')
	return append(w, '}'), nil
}

var _MetadataGetInvertMappingResponseUnion = [2]UnionElement{
	{TLTag: 0x9286abac, TLName: "metadata.getInvertMappingResponse", TLString: "metadata.getInvertMappingResponse#9286abac"},
	{TLTag: 0x9286abab, TLName: "metadata.getInvertMappingResponseKeyNotExists", TLString: "metadata.getInvertMappingResponseKeyNotExists#9286abab"},
}

type MetadataGetInvertMappingResponseUnion struct {
	valueGetInvertMappingResponse MetadataGetInvertMappingResponse
	index                         int
}

func (item MetadataGetInvertMappingResponseUnion) TLName() string {
	return _MetadataGetInvertMappingResponseUnion[item.index].TLName
}
func (item MetadataGetInvertMappingResponseUnion) TLTag() uint32 {
	return _MetadataGetInvertMappingResponseUnion[item.index].TLTag
}

func (item *MetadataGetInvertMappingResponseUnion) Reset() { item.ResetToGetInvertMappingResponse() }

func (item *MetadataGetInvertMappingResponseUnion) IsGetInvertMappingResponse() bool {
	return item.index == 0
}

func (item *MetadataGetInvertMappingResponseUnion) AsGetInvertMappingResponse() (*MetadataGetInvertMappingResponse, bool) {
	if item.index != 0 {
		return nil, false
	}
	return &item.valueGetInvertMappingResponse, true
}
func (item *MetadataGetInvertMappingResponseUnion) ResetToGetInvertMappingResponse() *MetadataGetInvertMappingResponse {
	item.index = 0
	item.valueGetInvertMappingResponse.Reset()
	return &item.valueGetInvertMappingResponse
}
func (item *MetadataGetInvertMappingResponseUnion) SetGetInvertMappingResponse(value MetadataGetInvertMappingResponse) {
	item.index = 0
	item.valueGetInvertMappingResponse = value
}

func (item *MetadataGetInvertMappingResponseUnion) IsKeyNotExists() bool { return item.index == 1 }

func (item *MetadataGetInvertMappingResponseUnion) AsKeyNotExists() (MetadataGetInvertMappingResponseKeyNotExists, bool) {
	var value MetadataGetInvertMappingResponseKeyNotExists
	return value, item.index == 1
}
func (item *MetadataGetInvertMappingResponseUnion) ResetToKeyNotExists() { item.index = 1 }
func (item *MetadataGetInvertMappingResponseUnion) SetKeyNotExists()     { item.index = 1 }

func (item *MetadataGetInvertMappingResponseUnion) ReadBoxed(w []byte, nat_field_mask uint32) (_ []byte, err error) {
	var tag uint32
	if w, err = basictl.NatRead(w, &tag); err != nil {
		return w, err
	}
	switch tag {
	case 0x9286abac:
		item.index = 0
		return item.valueGetInvertMappingResponse.Read(w, nat_field_mask)
	case 0x9286abab:
		item.index = 1
		return w, nil
	default:
		return w, ErrorInvalidUnionTag("metadata.GetInvertMappingResponse", tag)
	}
}

func (item *MetadataGetInvertMappingResponseUnion) WriteBoxed(w []byte, nat_field_mask uint32) (_ []byte, err error) {
	w = basictl.NatWrite(w, _MetadataGetInvertMappingResponseUnion[item.index].TLTag)
	switch item.index {
	case 0:
		return item.valueGetInvertMappingResponse.Write(w, nat_field_mask)
	case 1:
		return w, nil
	default: // Impossible due to panic above
		return w, nil
	}
}

func MetadataGetInvertMappingResponseUnion__ReadJSON(item *MetadataGetInvertMappingResponseUnion, j interface{}, nat_field_mask uint32) error {
	return item.readJSON(j, nat_field_mask)
}
func (item *MetadataGetInvertMappingResponseUnion) readJSON(j interface{}, nat_field_mask uint32) error {
	_jm, _tag, err := JsonReadUnionType("metadata.GetInvertMappingResponse", j)
	if err != nil {
		return err
	}
	jvalue := _jm["value"]
	switch _tag {
	case "metadata.getInvertMappingResponse#9286abac", "metadata.getInvertMappingResponse", "#9286abac":
		item.index = 0
		if err := MetadataGetInvertMappingResponse__ReadJSON(&item.valueGetInvertMappingResponse, jvalue, nat_field_mask); err != nil {
			return err
		}
		delete(_jm, "value")
	case "metadata.getInvertMappingResponseKeyNotExists#9286abab", "metadata.getInvertMappingResponseKeyNotExists", "#9286abab":
		item.index = 1
	default:
		return ErrorInvalidUnionTagJSON("metadata.GetInvertMappingResponse", _tag)
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("metadata.GetInvertMappingResponse", k)
	}
	return nil
}

func (item *MetadataGetInvertMappingResponseUnion) WriteJSON(w []byte, nat_field_mask uint32) (_ []byte, err error) {
	switch item.index {
	case 0:
		w = append(w, `{"type":"metadata.getInvertMappingResponse#9286abac","value":`...)
		if w, err = item.valueGetInvertMappingResponse.WriteJSON(w, nat_field_mask); err != nil {
			return w, err
		}
		return append(w, '}'), nil
	case 1:
		return append(w, `{"type":"metadata.getInvertMappingResponseKeyNotExists#9286abab"}`...), nil
	default: // Impossible due to panic above
		return w, nil
	}
}
