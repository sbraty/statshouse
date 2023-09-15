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

type StatshouseMapping struct {
	Str   string
	Value int32
}

func (StatshouseMapping) TLName() string { return "statshouse.mapping" }
func (StatshouseMapping) TLTag() uint32  { return 0xbf401d4b }

func (item *StatshouseMapping) Reset() {
	item.Str = ""
	item.Value = 0
}

func (item *StatshouseMapping) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringRead(w, &item.Str); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.Value)
}

func (item *StatshouseMapping) Write(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringWrite(w, item.Str); err != nil {
		return w, err
	}
	return basictl.IntWrite(w, item.Value), nil
}

func (item *StatshouseMapping) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xbf401d4b); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseMapping) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xbf401d4b)
	return item.Write(w)
}

func (item StatshouseMapping) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func StatshouseMapping__ReadJSON(item *StatshouseMapping, j interface{}) error {
	return item.readJSON(j)
}
func (item *StatshouseMapping) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouse.mapping", "expected json object")
	}
	_jStr := _jm["str"]
	delete(_jm, "str")
	if err := JsonReadString(_jStr, &item.Str); err != nil {
		return err
	}
	_jValue := _jm["value"]
	delete(_jm, "value")
	if err := JsonReadInt32(_jValue, &item.Value); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouse.mapping", k)
	}
	return nil
}

func (item *StatshouseMapping) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *StatshouseMapping) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if len(item.Str) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"str":`...)
		w = basictl.JSONWriteString(w, item.Str)
	}
	if item.Value != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"value":`...)
		w = basictl.JSONWriteInt32(w, item.Value)
	}
	return append(w, '}'), nil
}

func (item *StatshouseMapping) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseMapping) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("statshouse.mapping", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("statshouse.mapping", err.Error())
	}
	return nil
}

type StatshouseMappingBytes struct {
	Str   []byte
	Value int32
}

func (StatshouseMappingBytes) TLName() string { return "statshouse.mapping" }
func (StatshouseMappingBytes) TLTag() uint32  { return 0xbf401d4b }

func (item *StatshouseMappingBytes) Reset() {
	item.Str = item.Str[:0]
	item.Value = 0
}

func (item *StatshouseMappingBytes) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringReadBytes(w, &item.Str); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.Value)
}

func (item *StatshouseMappingBytes) Write(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringWriteBytes(w, item.Str); err != nil {
		return w, err
	}
	return basictl.IntWrite(w, item.Value), nil
}

func (item *StatshouseMappingBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xbf401d4b); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseMappingBytes) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xbf401d4b)
	return item.Write(w)
}

func (item StatshouseMappingBytes) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func StatshouseMappingBytes__ReadJSON(item *StatshouseMappingBytes, j interface{}) error {
	return item.readJSON(j)
}
func (item *StatshouseMappingBytes) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouse.mapping", "expected json object")
	}
	_jStr := _jm["str"]
	delete(_jm, "str")
	if err := JsonReadStringBytes(_jStr, &item.Str); err != nil {
		return err
	}
	_jValue := _jm["value"]
	delete(_jm, "value")
	if err := JsonReadInt32(_jValue, &item.Value); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouse.mapping", k)
	}
	return nil
}

func (item *StatshouseMappingBytes) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *StatshouseMappingBytes) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if len(item.Str) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"str":`...)
		w = basictl.JSONWriteStringBytes(w, item.Str)
	}
	if item.Value != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"value":`...)
		w = basictl.JSONWriteInt32(w, item.Value)
	}
	return append(w, '}'), nil
}

func (item *StatshouseMappingBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseMappingBytes) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("statshouse.mapping", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("statshouse.mapping", err.Error())
	}
	return nil
}

func VectorStatshouseMapping0Read(w []byte, vec *[]StatshouseMapping) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]StatshouseMapping, l)
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

func VectorStatshouseMapping0Write(w []byte, vec []StatshouseMapping) (_ []byte, err error) {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		if w, err = elem.Write(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorStatshouseMapping0ReadJSON(j interface{}, vec *[]StatshouseMapping) error {
	l, _arr, err := JsonReadArray("[]StatshouseMapping", j)
	if err != nil {
		return err
	}
	if cap(*vec) < l {
		*vec = make([]StatshouseMapping, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if err := StatshouseMapping__ReadJSON(&(*vec)[i], _arr[i]); err != nil {
			return err
		}
	}
	return nil
}

func VectorStatshouseMapping0WriteJSON(w []byte, vec []StatshouseMapping) (_ []byte, err error) {
	return VectorStatshouseMapping0WriteJSONOpt(false, w, vec)
}
func VectorStatshouseMapping0WriteJSONOpt(short bool, w []byte, vec []StatshouseMapping) (_ []byte, err error) {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		if w, err = elem.WriteJSONOpt(short, w); err != nil {
			return w, err
		}
	}
	return append(w, ']'), nil
}

func VectorStatshouseMapping0BytesRead(w []byte, vec *[]StatshouseMappingBytes) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]StatshouseMappingBytes, l)
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

func VectorStatshouseMapping0BytesWrite(w []byte, vec []StatshouseMappingBytes) (_ []byte, err error) {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		if w, err = elem.Write(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorStatshouseMapping0BytesReadJSON(j interface{}, vec *[]StatshouseMappingBytes) error {
	l, _arr, err := JsonReadArray("[]StatshouseMappingBytes", j)
	if err != nil {
		return err
	}
	if cap(*vec) < l {
		*vec = make([]StatshouseMappingBytes, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if err := StatshouseMappingBytes__ReadJSON(&(*vec)[i], _arr[i]); err != nil {
			return err
		}
	}
	return nil
}

func VectorStatshouseMapping0BytesWriteJSON(w []byte, vec []StatshouseMappingBytes) (_ []byte, err error) {
	return VectorStatshouseMapping0BytesWriteJSONOpt(false, w, vec)
}
func VectorStatshouseMapping0BytesWriteJSONOpt(short bool, w []byte, vec []StatshouseMappingBytes) (_ []byte, err error) {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		if w, err = elem.WriteJSONOpt(short, w); err != nil {
			return w, err
		}
	}
	return append(w, ']'), nil
}
