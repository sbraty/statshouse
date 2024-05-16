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

type EngineReadWriteMode struct {
	ReadEnabled  bool // Conditional: nat_fields_mask.0
	WriteEnabled bool // Conditional: nat_fields_mask.1
}

func (EngineReadWriteMode) TLName() string { return "engine.mode" }
func (EngineReadWriteMode) TLTag() uint32  { return 0xb9b7b6c9 }

func (item *EngineReadWriteMode) SetReadEnabled(v bool, nat_fields_mask *uint32) {
	item.ReadEnabled = v
	if nat_fields_mask != nil {
		*nat_fields_mask |= 1 << 0
	}
}
func (item *EngineReadWriteMode) ClearReadEnabled(nat_fields_mask *uint32) {
	item.ReadEnabled = false
	if nat_fields_mask != nil {
		*nat_fields_mask &^= 1 << 0
	}
}
func (item EngineReadWriteMode) IsSetReadEnabled(nat_fields_mask uint32) bool {
	return nat_fields_mask&(1<<0) != 0
}

func (item *EngineReadWriteMode) SetWriteEnabled(v bool, nat_fields_mask *uint32) {
	item.WriteEnabled = v
	if nat_fields_mask != nil {
		*nat_fields_mask |= 1 << 1
	}
}
func (item *EngineReadWriteMode) ClearWriteEnabled(nat_fields_mask *uint32) {
	item.WriteEnabled = false
	if nat_fields_mask != nil {
		*nat_fields_mask &^= 1 << 1
	}
}
func (item EngineReadWriteMode) IsSetWriteEnabled(nat_fields_mask uint32) bool {
	return nat_fields_mask&(1<<1) != 0
}

func (item *EngineReadWriteMode) Reset() {
	item.ReadEnabled = false
	item.WriteEnabled = false
}

func (item *EngineReadWriteMode) Read(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	if nat_fields_mask&(1<<0) != 0 {
		if w, err = BoolReadBoxed(w, &item.ReadEnabled); err != nil {
			return w, err
		}
	} else {
		item.ReadEnabled = false
	}
	if nat_fields_mask&(1<<1) != 0 {
		if w, err = BoolReadBoxed(w, &item.WriteEnabled); err != nil {
			return w, err
		}
	} else {
		item.WriteEnabled = false
	}
	return w, nil
}

// This method is general version of Write, use it instead!
func (item *EngineReadWriteMode) WriteGeneral(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	return item.Write(w, nat_fields_mask), nil
}

func (item *EngineReadWriteMode) Write(w []byte, nat_fields_mask uint32) []byte {
	if nat_fields_mask&(1<<0) != 0 {
		w = BoolWriteBoxed(w, item.ReadEnabled)
	}
	if nat_fields_mask&(1<<1) != 0 {
		w = BoolWriteBoxed(w, item.WriteEnabled)
	}
	return w
}

func (item *EngineReadWriteMode) ReadBoxed(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xb9b7b6c9); err != nil {
		return w, err
	}
	return item.Read(w, nat_fields_mask)
}

// This method is general version of WriteBoxed, use it instead!
func (item *EngineReadWriteMode) WriteBoxedGeneral(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	return item.WriteBoxed(w, nat_fields_mask), nil
}

func (item *EngineReadWriteMode) WriteBoxed(w []byte, nat_fields_mask uint32) []byte {
	w = basictl.NatWrite(w, 0xb9b7b6c9)
	return item.Write(w, nat_fields_mask)
}

func (item *EngineReadWriteMode) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_fields_mask uint32) error {
	var propReadEnabledPresented bool
	var propWriteEnabledPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "read_enabled":
				if propReadEnabledPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("engine.mode", "read_enabled")
				}
				if nat_fields_mask&(1<<0) == 0 {
					return ErrorInvalidJSON("engine.mode", "field 'read_enabled' is defined, while corresponding implicit fieldmask bit is 0")
				}
				if err := Json2ReadBool(in, &item.ReadEnabled); err != nil {
					return err
				}
				propReadEnabledPresented = true
			case "write_enabled":
				if propWriteEnabledPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("engine.mode", "write_enabled")
				}
				if nat_fields_mask&(1<<1) == 0 {
					return ErrorInvalidJSON("engine.mode", "field 'write_enabled' is defined, while corresponding implicit fieldmask bit is 0")
				}
				if err := Json2ReadBool(in, &item.WriteEnabled); err != nil {
					return err
				}
				propWriteEnabledPresented = true
			default:
				return ErrorInvalidJSONExcessElement("engine.mode", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propReadEnabledPresented {
		item.ReadEnabled = false
	}
	if !propWriteEnabledPresented {
		item.WriteEnabled = false
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *EngineReadWriteMode) WriteJSONGeneral(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_fields_mask), nil
}

func (item *EngineReadWriteMode) WriteJSON(w []byte, nat_fields_mask uint32) []byte {
	return item.WriteJSONOpt(true, false, w, nat_fields_mask)
}
func (item *EngineReadWriteMode) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_fields_mask uint32) []byte {
	w = append(w, '{')
	if nat_fields_mask&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"read_enabled":`...)
		w = basictl.JSONWriteBool(w, item.ReadEnabled)
	}
	if nat_fields_mask&(1<<1) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"write_enabled":`...)
		w = basictl.JSONWriteBool(w, item.WriteEnabled)
	}
	return append(w, '}')
}
