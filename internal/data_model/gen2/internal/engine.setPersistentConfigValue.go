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

type EngineSetPersistentConfigValue struct {
	Name  string
	Value int32
}

func (EngineSetPersistentConfigValue) TLName() string { return "engine.setPersistentConfigValue" }
func (EngineSetPersistentConfigValue) TLTag() uint32  { return 0x4cc8953f }

func (item *EngineSetPersistentConfigValue) Reset() {
	item.Name = ""
	item.Value = 0
}

func (item *EngineSetPersistentConfigValue) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringRead(w, &item.Name); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.Value)
}

func (item *EngineSetPersistentConfigValue) Write(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringWrite(w, item.Name); err != nil {
		return w, err
	}
	return basictl.IntWrite(w, item.Value), nil
}

func (item *EngineSetPersistentConfigValue) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x4cc8953f); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *EngineSetPersistentConfigValue) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x4cc8953f)
	return item.Write(w)
}

func (item *EngineSetPersistentConfigValue) ReadResult(w []byte, ret *True) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *EngineSetPersistentConfigValue) WriteResult(w []byte, ret True) (_ []byte, err error) {
	return ret.WriteBoxed(w)
}

func (item *EngineSetPersistentConfigValue) ReadResultJSON(j interface{}, ret *True) error {
	if err := True__ReadJSON(ret, j); err != nil {
		return err
	}
	return nil
}

func (item *EngineSetPersistentConfigValue) WriteResultJSON(w []byte, ret True) (_ []byte, err error) {
	return item.writeResultJSON(false, w, ret)
}

func (item *EngineSetPersistentConfigValue) writeResultJSON(short bool, w []byte, ret True) (_ []byte, err error) {
	if w, err = ret.WriteJSONOpt(short, w); err != nil {
		return w, err
	}
	return w, nil
}

func (item *EngineSetPersistentConfigValue) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *EngineSetPersistentConfigValue) ReadResultWriteResultJSONShort(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(true, w, ret)
	return r, w, err
}

func (item *EngineSetPersistentConfigValue) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := JsonBytesToInterface(r)
	if err != nil {
		return r, w, ErrorInvalidJSON("engine.setPersistentConfigValue", err.Error())
	}
	var ret True
	if err = item.ReadResultJSON(j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item EngineSetPersistentConfigValue) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func EngineSetPersistentConfigValue__ReadJSON(item *EngineSetPersistentConfigValue, j interface{}) error {
	return item.readJSON(j)
}
func (item *EngineSetPersistentConfigValue) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("engine.setPersistentConfigValue", "expected json object")
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
		return ErrorInvalidJSONExcessElement("engine.setPersistentConfigValue", k)
	}
	return nil
}

func (item *EngineSetPersistentConfigValue) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *EngineSetPersistentConfigValue) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
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

func (item *EngineSetPersistentConfigValue) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EngineSetPersistentConfigValue) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("engine.setPersistentConfigValue", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("engine.setPersistentConfigValue", err.Error())
	}
	return nil
}
