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

func BuiltinVectorStatshouseTopElementRead(w []byte, vec *[]StatshouseTopElement) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]StatshouseTopElement, l)
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

func BuiltinVectorStatshouseTopElementWrite(w []byte, vec []StatshouseTopElement) (_ []byte, err error) {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		if w, err = elem.Write(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinVectorStatshouseTopElementReadJSON(j interface{}, vec *[]StatshouseTopElement) error {
	l, _arr, err := JsonReadArray("[]StatshouseTopElement", j)
	if err != nil {
		return err
	}
	if cap(*vec) < l {
		*vec = make([]StatshouseTopElement, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if err := StatshouseTopElement__ReadJSON(&(*vec)[i], _arr[i]); err != nil {
			return err
		}
	}
	return nil
}

func BuiltinVectorStatshouseTopElementWriteJSON(w []byte, vec []StatshouseTopElement) (_ []byte, err error) {
	return BuiltinVectorStatshouseTopElementWriteJSONOpt(false, w, vec)
}
func BuiltinVectorStatshouseTopElementWriteJSONOpt(short bool, w []byte, vec []StatshouseTopElement) (_ []byte, err error) {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		if w, err = elem.WriteJSONOpt(short, w); err != nil {
			return w, err
		}
	}
	return append(w, ']'), nil
}

func BuiltinVectorStatshouseTopElementBytesRead(w []byte, vec *[]StatshouseTopElementBytes) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]StatshouseTopElementBytes, l)
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

func BuiltinVectorStatshouseTopElementBytesWrite(w []byte, vec []StatshouseTopElementBytes) (_ []byte, err error) {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		if w, err = elem.Write(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinVectorStatshouseTopElementBytesReadJSON(j interface{}, vec *[]StatshouseTopElementBytes) error {
	l, _arr, err := JsonReadArray("[]StatshouseTopElementBytes", j)
	if err != nil {
		return err
	}
	if cap(*vec) < l {
		*vec = make([]StatshouseTopElementBytes, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if err := StatshouseTopElementBytes__ReadJSON(&(*vec)[i], _arr[i]); err != nil {
			return err
		}
	}
	return nil
}

func BuiltinVectorStatshouseTopElementBytesWriteJSON(w []byte, vec []StatshouseTopElementBytes) (_ []byte, err error) {
	return BuiltinVectorStatshouseTopElementBytesWriteJSONOpt(false, w, vec)
}
func BuiltinVectorStatshouseTopElementBytesWriteJSONOpt(short bool, w []byte, vec []StatshouseTopElementBytes) (_ []byte, err error) {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		if w, err = elem.WriteJSONOpt(short, w); err != nil {
			return w, err
		}
	}
	return append(w, ']'), nil
}

type StatshouseTopElement struct {
	Key        string
	FieldsMask uint32
	Value      StatshouseMultiValue
}

func (StatshouseTopElement) TLName() string { return "statshouse.top_element" }
func (StatshouseTopElement) TLTag() uint32  { return 0x9ffdea42 }

func (item *StatshouseTopElement) Reset() {
	item.Key = ""
	item.FieldsMask = 0
	item.Value.Reset()
}

func (item *StatshouseTopElement) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringRead(w, &item.Key); err != nil {
		return w, err
	}
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	return item.Value.Read(w, item.FieldsMask)
}

func (item *StatshouseTopElement) Write(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringWrite(w, item.Key); err != nil {
		return w, err
	}
	w = basictl.NatWrite(w, item.FieldsMask)
	return item.Value.Write(w, item.FieldsMask)
}

func (item *StatshouseTopElement) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9ffdea42); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseTopElement) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x9ffdea42)
	return item.Write(w)
}

func (item StatshouseTopElement) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func StatshouseTopElement__ReadJSON(item *StatshouseTopElement, j interface{}) error {
	return item.readJSON(j)
}
func (item *StatshouseTopElement) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouse.top_element", "expected json object")
	}
	_jKey := _jm["key"]
	delete(_jm, "key")
	if err := JsonReadString(_jKey, &item.Key); err != nil {
		return err
	}
	_jFieldsMask := _jm["fields_mask"]
	delete(_jm, "fields_mask")
	if err := JsonReadUint32(_jFieldsMask, &item.FieldsMask); err != nil {
		return err
	}
	_jValue := _jm["value"]
	delete(_jm, "value")
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouse.top_element", k)
	}
	if err := StatshouseMultiValue__ReadJSON(&item.Value, _jValue, item.FieldsMask); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseTopElement) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *StatshouseTopElement) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if len(item.Key) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"key":`...)
		w = basictl.JSONWriteString(w, item.Key)
	}
	if item.FieldsMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"fields_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldsMask)
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"value":`...)
	if w, err = item.Value.WriteJSONOpt(short, w, item.FieldsMask); err != nil {
		return w, err
	}
	return append(w, '}'), nil
}

func (item *StatshouseTopElement) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseTopElement) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("statshouse.top_element", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("statshouse.top_element", err.Error())
	}
	return nil
}

type StatshouseTopElementBytes struct {
	Key        []byte
	FieldsMask uint32
	Value      StatshouseMultiValueBytes
}

func (StatshouseTopElementBytes) TLName() string { return "statshouse.top_element" }
func (StatshouseTopElementBytes) TLTag() uint32  { return 0x9ffdea42 }

func (item *StatshouseTopElementBytes) Reset() {
	item.Key = item.Key[:0]
	item.FieldsMask = 0
	item.Value.Reset()
}

func (item *StatshouseTopElementBytes) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringReadBytes(w, &item.Key); err != nil {
		return w, err
	}
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	return item.Value.Read(w, item.FieldsMask)
}

func (item *StatshouseTopElementBytes) Write(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringWriteBytes(w, item.Key); err != nil {
		return w, err
	}
	w = basictl.NatWrite(w, item.FieldsMask)
	return item.Value.Write(w, item.FieldsMask)
}

func (item *StatshouseTopElementBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9ffdea42); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseTopElementBytes) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x9ffdea42)
	return item.Write(w)
}

func (item StatshouseTopElementBytes) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func StatshouseTopElementBytes__ReadJSON(item *StatshouseTopElementBytes, j interface{}) error {
	return item.readJSON(j)
}
func (item *StatshouseTopElementBytes) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouse.top_element", "expected json object")
	}
	_jKey := _jm["key"]
	delete(_jm, "key")
	if err := JsonReadStringBytes(_jKey, &item.Key); err != nil {
		return err
	}
	_jFieldsMask := _jm["fields_mask"]
	delete(_jm, "fields_mask")
	if err := JsonReadUint32(_jFieldsMask, &item.FieldsMask); err != nil {
		return err
	}
	_jValue := _jm["value"]
	delete(_jm, "value")
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouse.top_element", k)
	}
	if err := StatshouseMultiValueBytes__ReadJSON(&item.Value, _jValue, item.FieldsMask); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseTopElementBytes) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *StatshouseTopElementBytes) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if len(item.Key) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"key":`...)
		w = basictl.JSONWriteStringBytes(w, item.Key)
	}
	if item.FieldsMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"fields_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldsMask)
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"value":`...)
	if w, err = item.Value.WriteJSONOpt(short, w, item.FieldsMask); err != nil {
		return w, err
	}
	return append(w, '}'), nil
}

func (item *StatshouseTopElementBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseTopElementBytes) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("statshouse.top_element", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("statshouse.top_element", err.Error())
	}
	return nil
}
