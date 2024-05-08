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

type MetadataGetInvertMapping struct {
	FieldMask uint32
	Id        int32
}

func (MetadataGetInvertMapping) TLName() string { return "metadata.getInvertMapping" }
func (MetadataGetInvertMapping) TLTag() uint32  { return 0x9faf5280 }

func (item *MetadataGetInvertMapping) Reset() {
	item.FieldMask = 0
	item.Id = 0
}

func (item *MetadataGetInvertMapping) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldMask); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.Id)
}

func (item *MetadataGetInvertMapping) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldMask)
	return basictl.IntWrite(w, item.Id), nil
}

func (item *MetadataGetInvertMapping) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9faf5280); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *MetadataGetInvertMapping) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x9faf5280)
	return item.Write(w)
}

func (item *MetadataGetInvertMapping) ReadResult(w []byte, ret *MetadataGetInvertMappingResponse) (_ []byte, err error) {
	return ret.ReadBoxed(w, item.FieldMask)
}

func (item *MetadataGetInvertMapping) WriteResult(w []byte, ret MetadataGetInvertMappingResponse) (_ []byte, err error) {
	return ret.WriteBoxed(w, item.FieldMask)
}

func (item *MetadataGetInvertMapping) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *MetadataGetInvertMappingResponse) error {
	if err := ret.ReadJSON(legacyTypeNames, in, item.FieldMask); err != nil {
		return err
	}
	return nil
}

func (item *MetadataGetInvertMapping) WriteResultJSON(w []byte, ret MetadataGetInvertMappingResponse) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *MetadataGetInvertMapping) writeResultJSON(newTypeNames bool, short bool, w []byte, ret MetadataGetInvertMappingResponse) (_ []byte, err error) {
	if w, err = ret.WriteJSONOpt(newTypeNames, short, w, item.FieldMask); err != nil {
		return w, err
	}
	return w, nil
}

func (item *MetadataGetInvertMapping) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret MetadataGetInvertMappingResponse
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *MetadataGetInvertMapping) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret MetadataGetInvertMappingResponse
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *MetadataGetInvertMapping) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret MetadataGetInvertMappingResponse
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item MetadataGetInvertMapping) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *MetadataGetInvertMapping) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFieldMaskPresented bool
	var propIdPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "field_mask":
				if propFieldMaskPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("metadata.getInvertMapping", "field_mask")
				}
				if err := Json2ReadUint32(in, &item.FieldMask); err != nil {
					return err
				}
				propFieldMaskPresented = true
			case "id":
				if propIdPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("metadata.getInvertMapping", "id")
				}
				if err := Json2ReadInt32(in, &item.Id); err != nil {
					return err
				}
				propIdPresented = true
			default:
				return ErrorInvalidJSONExcessElement("metadata.getInvertMapping", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propFieldMaskPresented {
		item.FieldMask = 0
	}
	if !propIdPresented {
		item.Id = 0
	}
	return nil
}

func (item *MetadataGetInvertMapping) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *MetadataGetInvertMapping) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexFieldMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"field_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldMask)
	if (item.FieldMask != 0) == false {
		w = w[:backupIndexFieldMask]
	}
	backupIndexId := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"id":`...)
	w = basictl.JSONWriteInt32(w, item.Id)
	if (item.Id != 0) == false {
		w = w[:backupIndexId]
	}
	return append(w, '}'), nil
}

func (item *MetadataGetInvertMapping) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *MetadataGetInvertMapping) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("metadata.getInvertMapping", err.Error())
	}
	return nil
}
