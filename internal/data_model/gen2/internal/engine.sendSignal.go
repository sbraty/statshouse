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

type EngineSendSignal struct {
	Signal int32
}

func (EngineSendSignal) TLName() string { return "engine.sendSignal" }
func (EngineSendSignal) TLTag() uint32  { return 0x1a7708a3 }

func (item *EngineSendSignal) Reset() {
	item.Signal = 0
}

func (item *EngineSendSignal) Read(w []byte) (_ []byte, err error) {
	return basictl.IntRead(w, &item.Signal)
}

func (item *EngineSendSignal) Write(w []byte) (_ []byte, err error) {
	return basictl.IntWrite(w, item.Signal), nil
}

func (item *EngineSendSignal) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1a7708a3); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *EngineSendSignal) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x1a7708a3)
	return item.Write(w)
}

func (item *EngineSendSignal) ReadResult(w []byte, ret *True) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *EngineSendSignal) WriteResult(w []byte, ret True) (_ []byte, err error) {
	return ret.WriteBoxed(w)
}

func (item *EngineSendSignal) ReadResultJSON(j interface{}, ret *True) error {
	if err := True__ReadJSON(ret, j); err != nil {
		return err
	}
	return nil
}

func (item *EngineSendSignal) WriteResultJSON(w []byte, ret True) (_ []byte, err error) {
	if w, err = ret.WriteJSON(w); err != nil {
		return w, err
	}
	return w, nil
}

func (item *EngineSendSignal) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *EngineSendSignal) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := JsonBytesToInterface(r)
	if err != nil {
		return r, w, ErrorInvalidJSON("engine.sendSignal", err.Error())
	}
	var ret True
	if err = item.ReadResultJSON(j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item EngineSendSignal) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func EngineSendSignal__ReadJSON(item *EngineSendSignal, j interface{}) error { return item.readJSON(j) }
func (item *EngineSendSignal) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("engine.sendSignal", "expected json object")
	}
	_jSignal := _jm["signal"]
	delete(_jm, "signal")
	if err := JsonReadInt32(_jSignal, &item.Signal); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("engine.sendSignal", k)
	}
	return nil
}

func (item *EngineSendSignal) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.Signal != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"signal":`...)
		w = basictl.JSONWriteInt32(w, item.Signal)
	}
	return append(w, '}'), nil
}

func (item *EngineSendSignal) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EngineSendSignal) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("engine.sendSignal", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("engine.sendSignal", err.Error())
	}
	return nil
}
