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

type EngineDumpForceQueries struct {
	BuffersPressureThreshold float64
}

func (EngineDumpForceQueries) TLName() string { return "engine.dumpForceQueries" }
func (EngineDumpForceQueries) TLTag() uint32  { return 0xf1f90880 }

func (item *EngineDumpForceQueries) Reset() {
	item.BuffersPressureThreshold = 0
}

func (item *EngineDumpForceQueries) Read(w []byte) (_ []byte, err error) {
	return basictl.DoubleRead(w, &item.BuffersPressureThreshold)
}

func (item *EngineDumpForceQueries) Write(w []byte) (_ []byte, err error) {
	return basictl.DoubleWrite(w, item.BuffersPressureThreshold), nil
}

func (item *EngineDumpForceQueries) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xf1f90880); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *EngineDumpForceQueries) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xf1f90880)
	return item.Write(w)
}

func (item *EngineDumpForceQueries) ReadResult(w []byte, ret *True) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *EngineDumpForceQueries) WriteResult(w []byte, ret True) (_ []byte, err error) {
	return ret.WriteBoxed(w)
}

func (item *EngineDumpForceQueries) ReadResultJSON(j interface{}, ret *True) error {
	if err := True__ReadJSON(ret, j); err != nil {
		return err
	}
	return nil
}

func (item *EngineDumpForceQueries) WriteResultJSON(w []byte, ret True) (_ []byte, err error) {
	return item.writeResultJSON(false, w, ret)
}

func (item *EngineDumpForceQueries) writeResultJSON(short bool, w []byte, ret True) (_ []byte, err error) {
	if w, err = ret.WriteJSONOpt(short, w); err != nil {
		return w, err
	}
	return w, nil
}

func (item *EngineDumpForceQueries) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *EngineDumpForceQueries) ReadResultWriteResultJSONShort(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(true, w, ret)
	return r, w, err
}

func (item *EngineDumpForceQueries) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := JsonBytesToInterface(r)
	if err != nil {
		return r, w, ErrorInvalidJSON("engine.dumpForceQueries", err.Error())
	}
	var ret True
	if err = item.ReadResultJSON(j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item EngineDumpForceQueries) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func EngineDumpForceQueries__ReadJSON(item *EngineDumpForceQueries, j interface{}) error {
	return item.readJSON(j)
}
func (item *EngineDumpForceQueries) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("engine.dumpForceQueries", "expected json object")
	}
	_jBuffersPressureThreshold := _jm["buffers_pressure_threshold"]
	delete(_jm, "buffers_pressure_threshold")
	if err := JsonReadFloat64(_jBuffersPressureThreshold, &item.BuffersPressureThreshold); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("engine.dumpForceQueries", k)
	}
	return nil
}

func (item *EngineDumpForceQueries) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *EngineDumpForceQueries) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.BuffersPressureThreshold != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"buffers_pressure_threshold":`...)
		w = basictl.JSONWriteFloat64(w, item.BuffersPressureThreshold)
	}
	return append(w, '}'), nil
}

func (item *EngineDumpForceQueries) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EngineDumpForceQueries) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("engine.dumpForceQueries", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("engine.dumpForceQueries", err.Error())
	}
	return nil
}
