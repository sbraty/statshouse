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

type EngineAsyncSleep struct {
	TimeMs int32
}

func (EngineAsyncSleep) TLName() string { return "engine.asyncSleep" }
func (EngineAsyncSleep) TLTag() uint32  { return 0x60e50d3d }

func (item *EngineAsyncSleep) Reset() {
	item.TimeMs = 0
}

func (item *EngineAsyncSleep) Read(w []byte) (_ []byte, err error) {
	return basictl.IntRead(w, &item.TimeMs)
}

func (item *EngineAsyncSleep) Write(w []byte) (_ []byte, err error) {
	return basictl.IntWrite(w, item.TimeMs), nil
}

func (item *EngineAsyncSleep) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x60e50d3d); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *EngineAsyncSleep) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x60e50d3d)
	return item.Write(w)
}

func (item *EngineAsyncSleep) ReadResult(w []byte, ret *bool) (_ []byte, err error) {
	return BoolReadBoxed(w, ret)
}

func (item *EngineAsyncSleep) WriteResult(w []byte, ret bool) (_ []byte, err error) {
	return BoolWriteBoxed(w, ret)
}

func (item *EngineAsyncSleep) ReadResultJSON(j interface{}, ret *bool) error {
	if err := JsonReadBool(j, ret); err != nil {
		return err
	}
	return nil
}

func (item *EngineAsyncSleep) WriteResultJSON(w []byte, ret bool) (_ []byte, err error) {
	w = basictl.JSONWriteBool(w, ret)
	return w, nil
}

func (item *EngineAsyncSleep) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret bool
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *EngineAsyncSleep) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := JsonBytesToInterface(r)
	if err != nil {
		return r, w, ErrorInvalidJSON("engine.asyncSleep", err.Error())
	}
	var ret bool
	if err = item.ReadResultJSON(j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item EngineAsyncSleep) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func EngineAsyncSleep__ReadJSON(item *EngineAsyncSleep, j interface{}) error { return item.readJSON(j) }
func (item *EngineAsyncSleep) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("engine.asyncSleep", "expected json object")
	}
	_jTimeMs := _jm["time_ms"]
	delete(_jm, "time_ms")
	if err := JsonReadInt32(_jTimeMs, &item.TimeMs); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("engine.asyncSleep", k)
	}
	return nil
}

func (item *EngineAsyncSleep) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.TimeMs != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"time_ms":`...)
		w = basictl.JSONWriteInt32(w, item.TimeMs)
	}
	return append(w, '}'), nil
}

func (item *EngineAsyncSleep) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EngineAsyncSleep) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("engine.asyncSleep", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("engine.asyncSleep", err.Error())
	}
	return nil
}
