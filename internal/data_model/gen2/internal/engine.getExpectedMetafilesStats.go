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

type EngineGetExpectedMetafilesStats struct {
}

func (EngineGetExpectedMetafilesStats) TLName() string { return "engine.getExpectedMetafilesStats" }
func (EngineGetExpectedMetafilesStats) TLTag() uint32  { return 0x342f391 }

func (item *EngineGetExpectedMetafilesStats) Reset()                         {}
func (item *EngineGetExpectedMetafilesStats) Read(w []byte) ([]byte, error)  { return w, nil }
func (item *EngineGetExpectedMetafilesStats) Write(w []byte) ([]byte, error) { return w, nil }
func (item *EngineGetExpectedMetafilesStats) ReadBoxed(w []byte) ([]byte, error) {
	return basictl.NatReadExactTag(w, 0x342f391)
}
func (item *EngineGetExpectedMetafilesStats) WriteBoxed(w []byte) ([]byte, error) {
	return basictl.NatWrite(w, 0x342f391), nil
}

func (item *EngineGetExpectedMetafilesStats) ReadResult(w []byte, ret *map[string]EngineMetafilesStat) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1f4c618f); err != nil {
		return w, err
	}
	return BuiltinVectorDictionaryFieldEngineMetafilesStatBoxedRead(w, ret)
}

func (item *EngineGetExpectedMetafilesStats) WriteResult(w []byte, ret map[string]EngineMetafilesStat) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0x1f4c618f)
	return BuiltinVectorDictionaryFieldEngineMetafilesStatBoxedWrite(w, ret)
}

func (item *EngineGetExpectedMetafilesStats) ReadResultJSON(j interface{}, ret *map[string]EngineMetafilesStat) error {
	if err := BuiltinVectorDictionaryFieldEngineMetafilesStatBoxedReadJSON(j, ret); err != nil {
		return err
	}
	return nil
}

func (item *EngineGetExpectedMetafilesStats) WriteResultJSON(w []byte, ret map[string]EngineMetafilesStat) (_ []byte, err error) {
	return item.writeResultJSON(false, w, ret)
}

func (item *EngineGetExpectedMetafilesStats) writeResultJSON(short bool, w []byte, ret map[string]EngineMetafilesStat) (_ []byte, err error) {
	if w, err = BuiltinVectorDictionaryFieldEngineMetafilesStatBoxedWriteJSONOpt(short, w, ret); err != nil {
		return w, err
	}
	return w, nil
}

func (item *EngineGetExpectedMetafilesStats) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret map[string]EngineMetafilesStat
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *EngineGetExpectedMetafilesStats) ReadResultWriteResultJSONShort(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret map[string]EngineMetafilesStat
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(true, w, ret)
	return r, w, err
}

func (item *EngineGetExpectedMetafilesStats) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := JsonBytesToInterface(r)
	if err != nil {
		return r, w, ErrorInvalidJSON("engine.getExpectedMetafilesStats", err.Error())
	}
	var ret map[string]EngineMetafilesStat
	if err = item.ReadResultJSON(j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item EngineGetExpectedMetafilesStats) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func EngineGetExpectedMetafilesStats__ReadJSON(item *EngineGetExpectedMetafilesStats, j interface{}) error {
	return item.readJSON(j)
}
func (item *EngineGetExpectedMetafilesStats) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("engine.getExpectedMetafilesStats", "expected json object")
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("engine.getExpectedMetafilesStats", k)
	}
	return nil
}

func (item *EngineGetExpectedMetafilesStats) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *EngineGetExpectedMetafilesStats) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	return append(w, '}'), nil
}

func (item *EngineGetExpectedMetafilesStats) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EngineGetExpectedMetafilesStats) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("engine.getExpectedMetafilesStats", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("engine.getExpectedMetafilesStats", err.Error())
	}
	return nil
}
