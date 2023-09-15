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

type StatshouseApiGetQueryPointResponse struct {
	FieldsMask uint32
	Data       []float64
	Meta       []StatshouseApiPointMeta
}

func (StatshouseApiGetQueryPointResponse) TLName() string { return "statshouseApi.queryPointResponse" }
func (StatshouseApiGetQueryPointResponse) TLTag() uint32  { return 0x4487e41a }

func (item *StatshouseApiGetQueryPointResponse) Reset() {
	item.FieldsMask = 0
	item.Data = item.Data[:0]
	item.Meta = item.Meta[:0]
}

func (item *StatshouseApiGetQueryPointResponse) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = VectorDouble0Read(w, &item.Data); err != nil {
		return w, err
	}
	return VectorStatshouseApiPointMeta0Read(w, &item.Meta)
}

func (item *StatshouseApiGetQueryPointResponse) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	if w, err = VectorDouble0Write(w, item.Data); err != nil {
		return w, err
	}
	return VectorStatshouseApiPointMeta0Write(w, item.Meta)
}

func (item *StatshouseApiGetQueryPointResponse) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x4487e41a); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseApiGetQueryPointResponse) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x4487e41a)
	return item.Write(w)
}

func (item StatshouseApiGetQueryPointResponse) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func StatshouseApiGetQueryPointResponse__ReadJSON(item *StatshouseApiGetQueryPointResponse, j interface{}) error {
	return item.readJSON(j)
}
func (item *StatshouseApiGetQueryPointResponse) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouseApi.queryPointResponse", "expected json object")
	}
	_jFieldsMask := _jm["fields_mask"]
	delete(_jm, "fields_mask")
	if err := JsonReadUint32(_jFieldsMask, &item.FieldsMask); err != nil {
		return err
	}
	_jData := _jm["data"]
	delete(_jm, "data")
	_jMeta := _jm["meta"]
	delete(_jm, "meta")
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouseApi.queryPointResponse", k)
	}
	if err := VectorDouble0ReadJSON(_jData, &item.Data); err != nil {
		return err
	}
	if err := VectorStatshouseApiPointMeta0ReadJSON(_jMeta, &item.Meta); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseApiGetQueryPointResponse) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *StatshouseApiGetQueryPointResponse) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldsMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"fields_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldsMask)
	}
	if len(item.Data) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"data":`...)
		if w, err = VectorDouble0WriteJSONOpt(short, w, item.Data); err != nil {
			return w, err
		}
	}
	if len(item.Meta) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"meta":`...)
		if w, err = VectorStatshouseApiPointMeta0WriteJSONOpt(short, w, item.Meta); err != nil {
			return w, err
		}
	}
	return append(w, '}'), nil
}

func (item *StatshouseApiGetQueryPointResponse) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseApiGetQueryPointResponse) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("statshouseApi.queryPointResponse", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("statshouseApi.queryPointResponse", err.Error())
	}
	return nil
}
