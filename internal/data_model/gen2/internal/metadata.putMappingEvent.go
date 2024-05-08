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

type MetadataPutMappingEvent struct {
	FieldsMask uint32
	Keys       []string
	Value      []int32
}

func (MetadataPutMappingEvent) TLName() string { return "metadata.putMappingEvent" }
func (MetadataPutMappingEvent) TLTag() uint32  { return 0x12345676 }

func (item *MetadataPutMappingEvent) Reset() {
	item.FieldsMask = 0
	item.Keys = item.Keys[:0]
	item.Value = item.Value[:0]
}

func (item *MetadataPutMappingEvent) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = BuiltinVectorStringRead(w, &item.Keys); err != nil {
		return w, err
	}
	return BuiltinVectorIntRead(w, &item.Value)
}

func (item *MetadataPutMappingEvent) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	if w, err = BuiltinVectorStringWrite(w, item.Keys); err != nil {
		return w, err
	}
	return BuiltinVectorIntWrite(w, item.Value)
}

func (item *MetadataPutMappingEvent) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x12345676); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *MetadataPutMappingEvent) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x12345676)
	return item.Write(w)
}

func (item MetadataPutMappingEvent) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *MetadataPutMappingEvent) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFieldsMaskPresented bool
	var propKeysPresented bool
	var propValuePresented bool

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
					return ErrorInvalidJSONWithDuplicatingKeys("metadata.putMappingEvent", "fields_mask")
				}
				if err := Json2ReadUint32(in, &item.FieldsMask); err != nil {
					return err
				}
				propFieldsMaskPresented = true
			case "keys":
				if propKeysPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("metadata.putMappingEvent", "keys")
				}
				if err := BuiltinVectorStringReadJSON(legacyTypeNames, in, &item.Keys); err != nil {
					return err
				}
				propKeysPresented = true
			case "value":
				if propValuePresented {
					return ErrorInvalidJSONWithDuplicatingKeys("metadata.putMappingEvent", "value")
				}
				if err := BuiltinVectorIntReadJSON(legacyTypeNames, in, &item.Value); err != nil {
					return err
				}
				propValuePresented = true
			default:
				return ErrorInvalidJSONExcessElement("metadata.putMappingEvent", key)
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
	if !propKeysPresented {
		item.Keys = item.Keys[:0]
	}
	if !propValuePresented {
		item.Value = item.Value[:0]
	}
	return nil
}

func (item *MetadataPutMappingEvent) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *MetadataPutMappingEvent) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexFieldsMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"fields_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldsMask)
	if (item.FieldsMask != 0) == false {
		w = w[:backupIndexFieldsMask]
	}
	backupIndexKeys := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"keys":`...)
	if w, err = BuiltinVectorStringWriteJSONOpt(newTypeNames, short, w, item.Keys); err != nil {
		return w, err
	}
	if (len(item.Keys) != 0) == false {
		w = w[:backupIndexKeys]
	}
	backupIndexValue := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"value":`...)
	if w, err = BuiltinVectorIntWriteJSONOpt(newTypeNames, short, w, item.Value); err != nil {
		return w, err
	}
	if (len(item.Value) != 0) == false {
		w = w[:backupIndexValue]
	}
	return append(w, '}'), nil
}

func (item *MetadataPutMappingEvent) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *MetadataPutMappingEvent) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("metadata.putMappingEvent", err.Error())
	}
	return nil
}
