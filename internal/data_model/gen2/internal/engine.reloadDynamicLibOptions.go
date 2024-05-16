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

// This method is general version of Write, use it instead!
func (item *EngineReloadDynamicLibOptions) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *EngineReloadDynamicLibOptions) Write(w []byte) []byte {
	w = basictl.NatWrite(w, item.FieldsMask)
	w = basictl.StringWrite(w, item.LibId)
	w = basictl.StringWrite(w, item.LibFileName)
	if item.FieldsMask&(1<<0) != 0 {
		w = basictl.DoubleWrite(w, item.SlicesPart)
	}
	return w
}

func (item *EngineReloadDynamicLibOptions) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xf3d0fb1); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *EngineReloadDynamicLibOptions) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *EngineReloadDynamicLibOptions) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xf3d0fb1)
	return item.Write(w)
}

func (item EngineReloadDynamicLibOptions) String() string {
	return string(item.WriteJSON(nil))
}

func (item *EngineReloadDynamicLibOptions) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFieldsMaskPresented bool
	var propLibIdPresented bool
	var propLibFileNamePresented bool
	var propSlicesPartPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "fields_mask":
				if propFieldsMaskPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("engine.reloadDynamicLibOptions", "fields_mask")
				}
				if err := Json2ReadUint32(in, &item.FieldsMask); err != nil {
					return err
				}
				propFieldsMaskPresented = true
			case "lib_id":
				if propLibIdPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("engine.reloadDynamicLibOptions", "lib_id")
				}
				if err := Json2ReadString(in, &item.LibId); err != nil {
					return err
				}
				propLibIdPresented = true
			case "lib_file_name":
				if propLibFileNamePresented {
					return ErrorInvalidJSONWithDuplicatingKeys("engine.reloadDynamicLibOptions", "lib_file_name")
				}
				if err := Json2ReadString(in, &item.LibFileName); err != nil {
					return err
				}
				propLibFileNamePresented = true
			case "slices_part":
				if propSlicesPartPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("engine.reloadDynamicLibOptions", "slices_part")
				}
				if err := Json2ReadFloat64(in, &item.SlicesPart); err != nil {
					return err
				}
				propSlicesPartPresented = true
			default:
				return ErrorInvalidJSONExcessElement("engine.reloadDynamicLibOptions", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propFieldsMaskPresented {
		item.FieldsMask = 0
	}
	if !propLibIdPresented {
		item.LibId = ""
	}
	if !propLibFileNamePresented {
		item.LibFileName = ""
	}
	if !propSlicesPartPresented {
		item.SlicesPart = 0
	}
	if propSlicesPartPresented {
		item.FieldsMask |= 1 << 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *EngineReloadDynamicLibOptions) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *EngineReloadDynamicLibOptions) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *EngineReloadDynamicLibOptions) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexFieldsMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"fields_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldsMask)
	if (item.FieldsMask != 0) == false {
		w = w[:backupIndexFieldsMask]
	}
	backupIndexLibId := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"lib_id":`...)
	w = basictl.JSONWriteString(w, item.LibId)
	if (len(item.LibId) != 0) == false {
		w = w[:backupIndexLibId]
	}
	backupIndexLibFileName := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"lib_file_name":`...)
	w = basictl.JSONWriteString(w, item.LibFileName)
	if (len(item.LibFileName) != 0) == false {
		w = w[:backupIndexLibFileName]
	}
	if item.FieldsMask&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"slices_part":`...)
		w = basictl.JSONWriteFloat64(w, item.SlicesPart)
	}
	return append(w, '}')
}

func (item *EngineReloadDynamicLibOptions) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *EngineReloadDynamicLibOptions) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("engine.reloadDynamicLibOptions", err.Error())
	}
	return nil
}
