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

func BuiltinVectorStatshouseCentroidRead(w []byte, vec *[]StatshouseCentroid) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]StatshouseCentroid, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if w, err = (*vec)[i].Read(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinVectorStatshouseCentroidWrite(w []byte, vec []StatshouseCentroid) (_ []byte, err error) {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		if w, err = elem.Write(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinVectorStatshouseCentroidReadJSON(j interface{}, vec *[]StatshouseCentroid) error {
	l, _arr, err := JsonReadArray("[]StatshouseCentroid", j)
	if err != nil {
		return err
	}
	if cap(*vec) < l {
		*vec = make([]StatshouseCentroid, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if err := StatshouseCentroid__ReadJSON(&(*vec)[i], _arr[i]); err != nil {
			return err
		}
	}
	return nil
}

func BuiltinVectorStatshouseCentroidWriteJSON(w []byte, vec []StatshouseCentroid) (_ []byte, err error) {
	return BuiltinVectorStatshouseCentroidWriteJSONOpt(false, w, vec)
}
func BuiltinVectorStatshouseCentroidWriteJSONOpt(short bool, w []byte, vec []StatshouseCentroid) (_ []byte, err error) {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		if w, err = elem.WriteJSONOpt(short, w); err != nil {
			return w, err
		}
	}
	return append(w, ']'), nil
}

type StatshouseCentroid struct {
	Value  float32
	Weight float32
}

func (StatshouseCentroid) TLName() string { return "statshouse.centroid" }
func (StatshouseCentroid) TLTag() uint32  { return 0x73fd01e0 }

func (item *StatshouseCentroid) Reset() {
	item.Value = 0
	item.Weight = 0
}

func (item *StatshouseCentroid) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.FloatRead(w, &item.Value); err != nil {
		return w, err
	}
	return basictl.FloatRead(w, &item.Weight)
}

func (item *StatshouseCentroid) Write(w []byte) (_ []byte, err error) {
	w = basictl.FloatWrite(w, item.Value)
	return basictl.FloatWrite(w, item.Weight), nil
}

func (item *StatshouseCentroid) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x73fd01e0); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseCentroid) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x73fd01e0)
	return item.Write(w)
}

func (item StatshouseCentroid) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func StatshouseCentroid__ReadJSON(item *StatshouseCentroid, j interface{}) error {
	return item.readJSON(j)
}
func (item *StatshouseCentroid) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouse.centroid", "expected json object")
	}
	_jValue := _jm["value"]
	delete(_jm, "value")
	if err := JsonReadFloat32(_jValue, &item.Value); err != nil {
		return err
	}
	_jWeight := _jm["weight"]
	delete(_jm, "weight")
	if err := JsonReadFloat32(_jWeight, &item.Weight); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouse.centroid", k)
	}
	return nil
}

func (item *StatshouseCentroid) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *StatshouseCentroid) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.Value != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"value":`...)
		w = basictl.JSONWriteFloat32(w, item.Value)
	}
	if item.Weight != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"weight":`...)
		w = basictl.JSONWriteFloat32(w, item.Weight)
	}
	return append(w, '}'), nil
}

func (item *StatshouseCentroid) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseCentroid) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("statshouse.centroid", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("statshouse.centroid", err.Error())
	}
	return nil
}
