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

type EngineEnableMetafilesAnalyzer struct {
	Enable bool
}

func (EngineEnableMetafilesAnalyzer) TLName() string { return "engine.enableMetafilesAnalyzer" }
func (EngineEnableMetafilesAnalyzer) TLTag() uint32  { return 0x88836cdc }

func (item *EngineEnableMetafilesAnalyzer) Reset() {
	item.Enable = false
}

func (item *EngineEnableMetafilesAnalyzer) Read(w []byte) (_ []byte, err error) {
	return BoolReadBoxed(w, &item.Enable)
}

func (item *EngineEnableMetafilesAnalyzer) Write(w []byte) (_ []byte, err error) {
	return BoolWriteBoxed(w, item.Enable)
}

func (item *EngineEnableMetafilesAnalyzer) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x88836cdc); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *EngineEnableMetafilesAnalyzer) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x88836cdc)
	return item.Write(w)
}

func (item *EngineEnableMetafilesAnalyzer) ReadResult(w []byte, ret *bool) (_ []byte, err error) {
	return BoolReadBoxed(w, ret)
}

func (item *EngineEnableMetafilesAnalyzer) WriteResult(w []byte, ret bool) (_ []byte, err error) {
	return BoolWriteBoxed(w, ret)
}

func (item *EngineEnableMetafilesAnalyzer) ReadResultJSON(j interface{}, ret *bool) error {
	if err := JsonReadBool(j, ret); err != nil {
		return err
	}
	return nil
}

func (item *EngineEnableMetafilesAnalyzer) WriteResultJSON(w []byte, ret bool) (_ []byte, err error) {
	return item.writeResultJSON(false, w, ret)
}

func (item *EngineEnableMetafilesAnalyzer) writeResultJSON(short bool, w []byte, ret bool) (_ []byte, err error) {
	w = basictl.JSONWriteBool(w, ret)
	return w, nil
}

func (item *EngineEnableMetafilesAnalyzer) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret bool
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *EngineEnableMetafilesAnalyzer) ReadResultWriteResultJSONShort(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret bool
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(true, w, ret)
	return r, w, err
}

func (item *EngineEnableMetafilesAnalyzer) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := JsonBytesToInterface(r)
	if err != nil {
		return r, w, ErrorInvalidJSON("engine.enableMetafilesAnalyzer", err.Error())
	}
	var ret bool
	if err = item.ReadResultJSON(j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item EngineEnableMetafilesAnalyzer) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func EngineEnableMetafilesAnalyzer__ReadJSON(item *EngineEnableMetafilesAnalyzer, j interface{}) error {
	return item.readJSON(j)
}
func (item *EngineEnableMetafilesAnalyzer) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("engine.enableMetafilesAnalyzer", "expected json object")
	}
	_jEnable := _jm["enable"]
	delete(_jm, "enable")
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("engine.enableMetafilesAnalyzer", k)
	}
	if err := JsonReadBool(_jEnable, &item.Enable); err != nil {
		return err
	}
	return nil
}

func (item *EngineEnableMetafilesAnalyzer) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *EngineEnableMetafilesAnalyzer) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.Enable {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"enable":`...)
		w = basictl.JSONWriteBool(w, item.Enable)
	}
	return append(w, '}'), nil
}

func (item *EngineEnableMetafilesAnalyzer) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EngineEnableMetafilesAnalyzer) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("engine.enableMetafilesAnalyzer", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("engine.enableMetafilesAnalyzer", err.Error())
	}
	return nil
}
