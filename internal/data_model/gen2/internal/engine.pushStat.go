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

type EnginePushStat struct {
	FieldsMask uint32
	Stat       Stat // Conditional: item.FieldsMask.0
}

func (EnginePushStat) TLName() string { return "engine.pushStat" }
func (EnginePushStat) TLTag() uint32  { return 0xf4b19fa2 }

func (item *EnginePushStat) SetStat(v Stat) {
	item.Stat = v
	item.FieldsMask |= 1 << 0
}
func (item *EnginePushStat) ClearStat() {
	item.Stat.Reset()
	item.FieldsMask &^= 1 << 0
}
func (item EnginePushStat) IsSetStat() bool { return item.FieldsMask&(1<<0) != 0 }

func (item *EnginePushStat) Reset() {
	item.FieldsMask = 0
	item.Stat.Reset()
}

func (item *EnginePushStat) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if item.FieldsMask&(1<<0) != 0 {
		if w, err = item.Stat.Read(w); err != nil {
			return w, err
		}
	} else {
		item.Stat.Reset()
	}
	return w, nil
}

func (item *EnginePushStat) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	if item.FieldsMask&(1<<0) != 0 {
		if w, err = item.Stat.Write(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func (item *EnginePushStat) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xf4b19fa2); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *EnginePushStat) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xf4b19fa2)
	return item.Write(w)
}

func (item *EnginePushStat) ReadResult(w []byte, ret *bool) (_ []byte, err error) {
	return BoolReadBoxed(w, ret)
}

func (item *EnginePushStat) WriteResult(w []byte, ret bool) (_ []byte, err error) {
	return BoolWriteBoxed(w, ret)
}

func (item *EnginePushStat) ReadResultJSON(j interface{}, ret *bool) error {
	if err := JsonReadBool(j, ret); err != nil {
		return err
	}
	return nil
}

func (item *EnginePushStat) WriteResultJSON(w []byte, ret bool) (_ []byte, err error) {
	return item.writeResultJSON(false, w, ret)
}

func (item *EnginePushStat) writeResultJSON(short bool, w []byte, ret bool) (_ []byte, err error) {
	w = basictl.JSONWriteBool(w, ret)
	return w, nil
}

func (item *EnginePushStat) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret bool
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *EnginePushStat) ReadResultWriteResultJSONShort(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret bool
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(true, w, ret)
	return r, w, err
}

func (item *EnginePushStat) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := JsonBytesToInterface(r)
	if err != nil {
		return r, w, ErrorInvalidJSON("engine.pushStat", err.Error())
	}
	var ret bool
	if err = item.ReadResultJSON(j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item EnginePushStat) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func EnginePushStat__ReadJSON(item *EnginePushStat, j interface{}) error { return item.readJSON(j) }
func (item *EnginePushStat) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("engine.pushStat", "expected json object")
	}
	_jFieldsMask := _jm["fields_mask"]
	delete(_jm, "fields_mask")
	if err := JsonReadUint32(_jFieldsMask, &item.FieldsMask); err != nil {
		return err
	}
	_jStat := _jm["stat"]
	delete(_jm, "stat")
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("engine.pushStat", k)
	}
	if _jStat != nil {
		item.FieldsMask |= 1 << 0
	}
	if _jStat != nil {
		if err := Stat__ReadJSON(&item.Stat, _jStat); err != nil {
			return err
		}
	} else {
		item.Stat.Reset()
	}
	return nil
}

func (item *EnginePushStat) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *EnginePushStat) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldsMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"fields_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldsMask)
	}
	if item.FieldsMask&(1<<0) != 0 {
		if len(item.Stat) != 0 {
			w = basictl.JSONAddCommaIfNeeded(w)
			w = append(w, `"stat":`...)
			if w, err = item.Stat.WriteJSONOpt(short, w); err != nil {
				return w, err
			}
		}
	}
	return append(w, '}'), nil
}

func (item *EnginePushStat) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EnginePushStat) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("engine.pushStat", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("engine.pushStat", err.Error())
	}
	return nil
}
