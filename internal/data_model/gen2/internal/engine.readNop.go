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

type EngineReadNop struct {
}

func (EngineReadNop) TLName() string { return "engine.readNop" }
func (EngineReadNop) TLTag() uint32  { return 0x9d2b841f }

func (item *EngineReadNop) Reset()                         {}
func (item *EngineReadNop) Read(w []byte) ([]byte, error)  { return w, nil }
func (item *EngineReadNop) Write(w []byte) ([]byte, error) { return w, nil }
func (item *EngineReadNop) ReadBoxed(w []byte) ([]byte, error) {
	return basictl.NatReadExactTag(w, 0x9d2b841f)
}
func (item *EngineReadNop) WriteBoxed(w []byte) ([]byte, error) {
	return basictl.NatWrite(w, 0x9d2b841f), nil
}

func (item *EngineReadNop) ReadResult(w []byte, ret *True) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *EngineReadNop) WriteResult(w []byte, ret True) (_ []byte, err error) {
	return ret.WriteBoxed(w)
}

func (item *EngineReadNop) ReadResultJSON(j interface{}, ret *True) error {
	if err := True__ReadJSON(ret, j); err != nil {
		return err
	}
	return nil
}

func (item *EngineReadNop) WriteResultJSON(w []byte, ret True) (_ []byte, err error) {
	return item.writeResultJSON(false, w, ret)
}

func (item *EngineReadNop) writeResultJSON(short bool, w []byte, ret True) (_ []byte, err error) {
	if w, err = ret.WriteJSONOpt(short, w); err != nil {
		return w, err
	}
	return w, nil
}

func (item *EngineReadNop) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *EngineReadNop) ReadResultWriteResultJSONShort(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(true, w, ret)
	return r, w, err
}

func (item *EngineReadNop) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := JsonBytesToInterface(r)
	if err != nil {
		return r, w, ErrorInvalidJSON("engine.readNop", err.Error())
	}
	var ret True
	if err = item.ReadResultJSON(j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item EngineReadNop) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func EngineReadNop__ReadJSON(item *EngineReadNop, j interface{}) error { return item.readJSON(j) }
func (item *EngineReadNop) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("engine.readNop", "expected json object")
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("engine.readNop", k)
	}
	return nil
}

func (item *EngineReadNop) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *EngineReadNop) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	return append(w, '}'), nil
}

func (item *EngineReadNop) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EngineReadNop) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("engine.readNop", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("engine.readNop", err.Error())
	}
	return nil
}
