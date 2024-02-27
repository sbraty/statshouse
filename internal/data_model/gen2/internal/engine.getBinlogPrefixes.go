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

type EngineGetBinlogPrefixes struct {
}

func (EngineGetBinlogPrefixes) TLName() string { return "engine.getBinlogPrefixes" }
func (EngineGetBinlogPrefixes) TLTag() uint32  { return 0xef14db93 }

func (item *EngineGetBinlogPrefixes) Reset()                         {}
func (item *EngineGetBinlogPrefixes) Read(w []byte) ([]byte, error)  { return w, nil }
func (item *EngineGetBinlogPrefixes) Write(w []byte) ([]byte, error) { return w, nil }
func (item *EngineGetBinlogPrefixes) ReadBoxed(w []byte) ([]byte, error) {
	return basictl.NatReadExactTag(w, 0xef14db93)
}
func (item *EngineGetBinlogPrefixes) WriteBoxed(w []byte) ([]byte, error) {
	return basictl.NatWrite(w, 0xef14db93), nil
}

func (item *EngineGetBinlogPrefixes) ReadResult(w []byte, ret *[]EngineBinlogPrefix) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1cb5c415); err != nil {
		return w, err
	}
	return BuiltinVectorEngineBinlogPrefixRead(w, ret)
}

func (item *EngineGetBinlogPrefixes) WriteResult(w []byte, ret []EngineBinlogPrefix) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0x1cb5c415)
	return BuiltinVectorEngineBinlogPrefixWrite(w, ret)
}

func (item *EngineGetBinlogPrefixes) ReadResultJSON(j interface{}, ret *[]EngineBinlogPrefix) error {
	if err := BuiltinVectorEngineBinlogPrefixReadJSON(j, ret); err != nil {
		return err
	}
	return nil
}

func (item *EngineGetBinlogPrefixes) WriteResultJSON(w []byte, ret []EngineBinlogPrefix) (_ []byte, err error) {
	return item.writeResultJSON(false, w, ret)
}

func (item *EngineGetBinlogPrefixes) writeResultJSON(short bool, w []byte, ret []EngineBinlogPrefix) (_ []byte, err error) {
	if w, err = BuiltinVectorEngineBinlogPrefixWriteJSONOpt(short, w, ret); err != nil {
		return w, err
	}
	return w, nil
}

func (item *EngineGetBinlogPrefixes) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret []EngineBinlogPrefix
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *EngineGetBinlogPrefixes) ReadResultWriteResultJSONShort(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret []EngineBinlogPrefix
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(true, w, ret)
	return r, w, err
}

func (item *EngineGetBinlogPrefixes) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := JsonBytesToInterface(r)
	if err != nil {
		return r, w, ErrorInvalidJSON("engine.getBinlogPrefixes", err.Error())
	}
	var ret []EngineBinlogPrefix
	if err = item.ReadResultJSON(j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item EngineGetBinlogPrefixes) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func EngineGetBinlogPrefixes__ReadJSON(item *EngineGetBinlogPrefixes, j interface{}) error {
	return item.readJSON(j)
}
func (item *EngineGetBinlogPrefixes) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("engine.getBinlogPrefixes", "expected json object")
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("engine.getBinlogPrefixes", k)
	}
	return nil
}

func (item *EngineGetBinlogPrefixes) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *EngineGetBinlogPrefixes) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	return append(w, '}'), nil
}

func (item *EngineGetBinlogPrefixes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EngineGetBinlogPrefixes) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("engine.getBinlogPrefixes", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("engine.getBinlogPrefixes", err.Error())
	}
	return nil
}
