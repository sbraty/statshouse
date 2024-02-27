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

type EngineCount struct {
}

func (EngineCount) TLName() string { return "engine.count" }
func (EngineCount) TLTag() uint32  { return 0x19d0f020 }

func (item *EngineCount) Reset()                         {}
func (item *EngineCount) Read(w []byte) ([]byte, error)  { return w, nil }
func (item *EngineCount) Write(w []byte) ([]byte, error) { return w, nil }
func (item *EngineCount) ReadBoxed(w []byte) ([]byte, error) {
	return basictl.NatReadExactTag(w, 0x19d0f020)
}
func (item *EngineCount) WriteBoxed(w []byte) ([]byte, error) {
	return basictl.NatWrite(w, 0x19d0f020), nil
}

func (item *EngineCount) ReadResult(w []byte, ret *BoolStat) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *EngineCount) WriteResult(w []byte, ret BoolStat) (_ []byte, err error) {
	return ret.WriteBoxed(w)
}

func (item *EngineCount) ReadResultJSON(j interface{}, ret *BoolStat) error {
	if err := BoolStat__ReadJSON(ret, j); err != nil {
		return err
	}
	return nil
}

func (item *EngineCount) WriteResultJSON(w []byte, ret BoolStat) (_ []byte, err error) {
	return item.writeResultJSON(false, w, ret)
}

func (item *EngineCount) writeResultJSON(short bool, w []byte, ret BoolStat) (_ []byte, err error) {
	if w, err = ret.WriteJSONOpt(short, w); err != nil {
		return w, err
	}
	return w, nil
}

func (item *EngineCount) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret BoolStat
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *EngineCount) ReadResultWriteResultJSONShort(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret BoolStat
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(true, w, ret)
	return r, w, err
}

func (item *EngineCount) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := JsonBytesToInterface(r)
	if err != nil {
		return r, w, ErrorInvalidJSON("engine.count", err.Error())
	}
	var ret BoolStat
	if err = item.ReadResultJSON(j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item EngineCount) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func EngineCount__ReadJSON(item *EngineCount, j interface{}) error { return item.readJSON(j) }
func (item *EngineCount) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("engine.count", "expected json object")
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("engine.count", k)
	}
	return nil
}

func (item *EngineCount) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *EngineCount) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	return append(w, '}'), nil
}

func (item *EngineCount) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EngineCount) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("engine.count", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("engine.count", err.Error())
	}
	return nil
}
