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

type EngineReloadDynamicLibOptions struct {
	FieldsMask  uint32
	LibId       string
	LibFileName string
	SlicesPart  float64 // Conditional: item.FieldsMask.0
}

func (EngineReloadDynamicLibOptions) TLName() string { return "engine.reloadDynamicLibOptions" }
func (EngineReloadDynamicLibOptions) TLTag() uint32  { return 0xf3d0fb1 }

func (item *EngineReloadDynamicLibOptions) SetSlicesPart(v float64) {
	item.SlicesPart = v
	item.FieldsMask |= 1 << 0
}
func (item *EngineReloadDynamicLibOptions) ClearSlicesPart() {
	item.SlicesPart = 0
	item.FieldsMask &^= 1 << 0
}
func (item EngineReloadDynamicLibOptions) IsSetSlicesPart() bool { return item.FieldsMask&(1<<0) != 0 }

func (item *EngineReloadDynamicLibOptions) Reset() {
	item.FieldsMask = 0
	item.LibId = ""
	item.LibFileName = ""
	item.SlicesPart = 0
}

func (item *EngineReloadDynamicLibOptions) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = basictl.StringRead(w, &item.LibId); err != nil {
		return w, err
	}
	if w, err = basictl.StringRead(w, &item.LibFileName); err != nil {
		return w, err
	}
	if item.FieldsMask&(1<<0) != 0 {
		if w, err = basictl.DoubleRead(w, &item.SlicesPart); err != nil {
			return w, err
		}
	} else {
		item.SlicesPart = 0
	}
	return w, nil
}

func (item *EngineReloadDynamicLibOptions) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	if w, err = basictl.StringWrite(w, item.LibId); err != nil {
		return w, err
	}
	if w, err = basictl.StringWrite(w, item.LibFileName); err != nil {
		return w, err
	}
	if item.FieldsMask&(1<<0) != 0 {
		w = basictl.DoubleWrite(w, item.SlicesPart)
	}
	return w, nil
}

func (item *EngineReloadDynamicLibOptions) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xf3d0fb1); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *EngineReloadDynamicLibOptions) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xf3d0fb1)
	return item.Write(w)
}

func (item EngineReloadDynamicLibOptions) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func EngineReloadDynamicLibOptions__ReadJSON(item *EngineReloadDynamicLibOptions, j interface{}) error {
	return item.readJSON(j)
}
func (item *EngineReloadDynamicLibOptions) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("engine.reloadDynamicLibOptions", "expected json object")
	}
	_jFieldsMask := _jm["fields_mask"]
	delete(_jm, "fields_mask")
	if err := JsonReadUint32(_jFieldsMask, &item.FieldsMask); err != nil {
		return err
	}
	_jLibId := _jm["lib_id"]
	delete(_jm, "lib_id")
	if err := JsonReadString(_jLibId, &item.LibId); err != nil {
		return err
	}
	_jLibFileName := _jm["lib_file_name"]
	delete(_jm, "lib_file_name")
	if err := JsonReadString(_jLibFileName, &item.LibFileName); err != nil {
		return err
	}
	_jSlicesPart := _jm["slices_part"]
	delete(_jm, "slices_part")
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("engine.reloadDynamicLibOptions", k)
	}
	if _jSlicesPart != nil {
		item.FieldsMask |= 1 << 0
	}
	if _jSlicesPart != nil {
		if err := JsonReadFloat64(_jSlicesPart, &item.SlicesPart); err != nil {
			return err
		}
	} else {
		item.SlicesPart = 0
	}
	return nil
}

func (item *EngineReloadDynamicLibOptions) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *EngineReloadDynamicLibOptions) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldsMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"fields_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldsMask)
	}
	if len(item.LibId) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"lib_id":`...)
		w = basictl.JSONWriteString(w, item.LibId)
	}
	if len(item.LibFileName) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"lib_file_name":`...)
		w = basictl.JSONWriteString(w, item.LibFileName)
	}
	if item.FieldsMask&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"slices_part":`...)
		w = basictl.JSONWriteFloat64(w, item.SlicesPart)
	}
	return append(w, '}'), nil
}

func (item *EngineReloadDynamicLibOptions) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EngineReloadDynamicLibOptions) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("engine.reloadDynamicLibOptions", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("engine.reloadDynamicLibOptions", err.Error())
	}
	return nil
}
