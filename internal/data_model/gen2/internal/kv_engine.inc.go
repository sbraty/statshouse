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

type KvEngineInc struct {
	Key  int64
	Incr int64
}

func (KvEngineInc) TLName() string { return "kv_engine.inc" }
func (KvEngineInc) TLTag() uint32  { return 0x3c7239bb }

func (item *KvEngineInc) Reset() {
	item.Key = 0
	item.Incr = 0
}

func (item *KvEngineInc) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.LongRead(w, &item.Key); err != nil {
		return w, err
	}
	return basictl.LongRead(w, &item.Incr)
}

func (item *KvEngineInc) Write(w []byte) (_ []byte, err error) {
	w = basictl.LongWrite(w, item.Key)
	return basictl.LongWrite(w, item.Incr), nil
}

func (item *KvEngineInc) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x3c7239bb); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *KvEngineInc) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x3c7239bb)
	return item.Write(w)
}

func (item *KvEngineInc) ReadResult(w []byte, ret *KvEngineChangeResponse) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *KvEngineInc) WriteResult(w []byte, ret KvEngineChangeResponse) (_ []byte, err error) {
	return ret.WriteBoxed(w)
}

func (item *KvEngineInc) ReadResultJSON(j interface{}, ret *KvEngineChangeResponse) error {
	if err := KvEngineChangeResponse__ReadJSON(ret, j); err != nil {
		return err
	}
	return nil
}

func (item *KvEngineInc) WriteResultJSON(w []byte, ret KvEngineChangeResponse) (_ []byte, err error) {
	return item.writeResultJSON(false, w, ret)
}

func (item *KvEngineInc) writeResultJSON(short bool, w []byte, ret KvEngineChangeResponse) (_ []byte, err error) {
	if w, err = ret.WriteJSONOpt(short, w); err != nil {
		return w, err
	}
	return w, nil
}

func (item *KvEngineInc) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret KvEngineChangeResponse
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *KvEngineInc) ReadResultWriteResultJSONShort(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret KvEngineChangeResponse
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(true, w, ret)
	return r, w, err
}

func (item *KvEngineInc) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := JsonBytesToInterface(r)
	if err != nil {
		return r, w, ErrorInvalidJSON("kv_engine.inc", err.Error())
	}
	var ret KvEngineChangeResponse
	if err = item.ReadResultJSON(j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item KvEngineInc) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func KvEngineInc__ReadJSON(item *KvEngineInc, j interface{}) error { return item.readJSON(j) }
func (item *KvEngineInc) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("kv_engine.inc", "expected json object")
	}
	_jKey := _jm["key"]
	delete(_jm, "key")
	if err := JsonReadInt64(_jKey, &item.Key); err != nil {
		return err
	}
	_jIncr := _jm["incr"]
	delete(_jm, "incr")
	if err := JsonReadInt64(_jIncr, &item.Incr); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("kv_engine.inc", k)
	}
	return nil
}

func (item *KvEngineInc) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *KvEngineInc) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.Key != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"key":`...)
		w = basictl.JSONWriteInt64(w, item.Key)
	}
	if item.Incr != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"incr":`...)
		w = basictl.JSONWriteInt64(w, item.Incr)
	}
	return append(w, '}'), nil
}

func (item *KvEngineInc) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *KvEngineInc) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("kv_engine.inc", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("kv_engine.inc", err.Error())
	}
	return nil
}
