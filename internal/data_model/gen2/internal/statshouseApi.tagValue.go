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

type StatshouseApiTagValue struct {
	FieldsMask uint32
	In         bool
	Value      string
	Flag       StatshouseApiFlag
}

func (StatshouseApiTagValue) TLName() string { return "statshouseApi.tagValue" }
func (StatshouseApiTagValue) TLTag() uint32  { return 0x43eeb763 }

func (item *StatshouseApiTagValue) Reset() {
	item.FieldsMask = 0
	item.In = false
	item.Value = ""
	item.Flag.Reset()
}

func (item *StatshouseApiTagValue) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = BoolReadBoxed(w, &item.In); err != nil {
		return w, err
	}
	if w, err = basictl.StringRead(w, &item.Value); err != nil {
		return w, err
	}
	return item.Flag.ReadBoxed(w)
}

func (item *StatshouseApiTagValue) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	if w, err = BoolWriteBoxed(w, item.In); err != nil {
		return w, err
	}
	if w, err = basictl.StringWrite(w, item.Value); err != nil {
		return w, err
	}
	return item.Flag.WriteBoxed(w)
}

func (item *StatshouseApiTagValue) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x43eeb763); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseApiTagValue) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x43eeb763)
	return item.Write(w)
}

func (item StatshouseApiTagValue) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func StatshouseApiTagValue__ReadJSON(item *StatshouseApiTagValue, j interface{}) error {
	return item.readJSON(j)
}
func (item *StatshouseApiTagValue) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouseApi.tagValue", "expected json object")
	}
	_jFieldsMask := _jm["fields_mask"]
	delete(_jm, "fields_mask")
	if err := JsonReadUint32(_jFieldsMask, &item.FieldsMask); err != nil {
		return err
	}
	_jIn := _jm["in"]
	delete(_jm, "in")
	_jValue := _jm["value"]
	delete(_jm, "value")
	if err := JsonReadString(_jValue, &item.Value); err != nil {
		return err
	}
	_jFlag := _jm["flag"]
	delete(_jm, "flag")
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouseApi.tagValue", k)
	}
	if err := JsonReadBool(_jIn, &item.In); err != nil {
		return err
	}
	if err := StatshouseApiFlag__ReadJSON(&item.Flag, _jFlag); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseApiTagValue) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldsMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"fields_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldsMask)
	}
	if item.In {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"in":`...)
		w = basictl.JSONWriteBool(w, item.In)
	}
	if len(item.Value) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"value":`...)
		w = basictl.JSONWriteString(w, item.Value)
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"flag":`...)
	if w, err = item.Flag.WriteJSON(w); err != nil {
		return w, err
	}
	return append(w, '}'), nil
}

func (item *StatshouseApiTagValue) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseApiTagValue) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("statshouseApi.tagValue", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("statshouseApi.tagValue", err.Error())
	}
	return nil
}

func VectorStatshouseApiTagValue0Read(w []byte, vec *[]StatshouseApiTagValue) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]StatshouseApiTagValue, l)
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

func VectorStatshouseApiTagValue0Write(w []byte, vec []StatshouseApiTagValue) (_ []byte, err error) {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		if w, err = elem.Write(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorStatshouseApiTagValue0ReadJSON(j interface{}, vec *[]StatshouseApiTagValue) error {
	l, _arr, err := JsonReadArray("[]StatshouseApiTagValue", j)
	if err != nil {
		return err
	}
	if cap(*vec) < l {
		*vec = make([]StatshouseApiTagValue, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if err := StatshouseApiTagValue__ReadJSON(&(*vec)[i], _arr[i]); err != nil {
			return err
		}
	}
	return nil
}

func VectorStatshouseApiTagValue0WriteJSON(w []byte, vec []StatshouseApiTagValue) (_ []byte, err error) {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		if w, err = elem.WriteJSON(w); err != nil {
			return w, err
		}
	}
	return append(w, ']'), nil
}
