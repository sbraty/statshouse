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

type EngineSetNoPersistentConfigValue struct {
	Name  string
	Value int32
}

func (EngineSetNoPersistentConfigValue) TLName() string { return "engine.setNoPersistentConfigValue" }
func (EngineSetNoPersistentConfigValue) TLTag() uint32  { return 0x92aaa5b9 }

func (item *EngineSetNoPersistentConfigValue) Reset() {
	item.Name = ""
	item.Value = 0
}

func (item *EngineSetNoPersistentConfigValue) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringRead(w, &item.Name); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.Value)
}

func (item *EngineSetNoPersistentConfigValue) Write(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringWrite(w, item.Name); err != nil {
		return w, err
	}
	return basictl.IntWrite(w, item.Value), nil
}

func (item *EngineSetNoPersistentConfigValue) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x92aaa5b9); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *EngineSetNoPersistentConfigValue) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x92aaa5b9)
	return item.Write(w)
}

func (item *EngineSetNoPersistentConfigValue) ReadResult(w []byte, ret *True) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *EngineSetNoPersistentConfigValue) WriteResult(w []byte, ret True) (_ []byte, err error) {
	return ret.WriteBoxed(w)
}

func (item *EngineSetNoPersistentConfigValue) ReadResultJSON(j interface{}, ret *True) error {
	if err := True__ReadJSON(ret, j); err != nil {
		return err
	}
	return nil
}

func (item *EngineSetNoPersistentConfigValue) WriteResultJSON(w []byte, ret True) (_ []byte, err error) {
	if w, err = ret.WriteJSON(w); err != nil {
		return w, err
	}
	return w, nil
}

func (item *EngineSetNoPersistentConfigValue) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *EngineSetNoPersistentConfigValue) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := JsonBytesToInterface(r)
	if err != nil {
		return r, w, ErrorInvalidJSON("engine.setNoPersistentConfigValue", err.Error())
	}
	var ret True
	if err = item.ReadResultJSON(j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item EngineSetNoPersistentConfigValue) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func EngineSetNoPersistentConfigValue__ReadJSON(item *EngineSetNoPersistentConfigValue, j interface{}) error {
	return item.readJSON(j)
}
func (item *EngineSetNoPersistentConfigValue) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("engine.setNoPersistentConfigValue", "expected json object")
	}
	_jName := _jm["name"]
	delete(_jm, "name")
	if err := JsonReadString(_jName, &item.Name); err != nil {
		return err
	}
	_jValue := _jm["value"]
	delete(_jm, "value")
	if err := JsonReadInt32(_jValue, &item.Value); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("engine.setNoPersistentConfigValue", k)
	}
	return nil
}

func (item *EngineSetNoPersistentConfigValue) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if len(item.Name) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"name":`...)
		w = basictl.JSONWriteString(w, item.Name)
	}
	if item.Value != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"value":`...)
		w = basictl.JSONWriteInt32(w, item.Value)
	}
	return append(w, '}'), nil
}

func (item *EngineSetNoPersistentConfigValue) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EngineSetNoPersistentConfigValue) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("engine.setNoPersistentConfigValue", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("engine.setNoPersistentConfigValue", err.Error())
	}
	return nil
}
