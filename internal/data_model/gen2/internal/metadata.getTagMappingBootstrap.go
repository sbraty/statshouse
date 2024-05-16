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

type MetadataGetTagMappingBootstrap struct {
	FieldsMask uint32
}

func (MetadataGetTagMappingBootstrap) TLName() string { return "metadata.getTagMappingBootstrap" }
func (MetadataGetTagMappingBootstrap) TLTag() uint32  { return 0x5fc81a9b }

func (item *MetadataGetTagMappingBootstrap) Reset() {
	item.FieldsMask = 0
}

func (item *MetadataGetTagMappingBootstrap) Read(w []byte) (_ []byte, err error) {
	return basictl.NatRead(w, &item.FieldsMask)
}

// This method is general version of Write, use it instead!
func (item *MetadataGetTagMappingBootstrap) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *MetadataGetTagMappingBootstrap) Write(w []byte) []byte {
	w = basictl.NatWrite(w, item.FieldsMask)
	return w
}

func (item *MetadataGetTagMappingBootstrap) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x5fc81a9b); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *MetadataGetTagMappingBootstrap) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *MetadataGetTagMappingBootstrap) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x5fc81a9b)
	return item.Write(w)
}

func (item *MetadataGetTagMappingBootstrap) ReadResult(w []byte, ret *StatshouseGetTagMappingBootstrapResult) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *MetadataGetTagMappingBootstrap) WriteResult(w []byte, ret StatshouseGetTagMappingBootstrapResult) (_ []byte, err error) {
	w = ret.WriteBoxed(w)
	return w, nil
}

func (item *MetadataGetTagMappingBootstrap) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *StatshouseGetTagMappingBootstrapResult) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *MetadataGetTagMappingBootstrap) WriteResultJSON(w []byte, ret StatshouseGetTagMappingBootstrapResult) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *MetadataGetTagMappingBootstrap) writeResultJSON(newTypeNames bool, short bool, w []byte, ret StatshouseGetTagMappingBootstrapResult) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w)
	return w, nil
}

func (item *MetadataGetTagMappingBootstrap) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret StatshouseGetTagMappingBootstrapResult
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *MetadataGetTagMappingBootstrap) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret StatshouseGetTagMappingBootstrapResult
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *MetadataGetTagMappingBootstrap) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret StatshouseGetTagMappingBootstrapResult
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item MetadataGetTagMappingBootstrap) String() string {
	return string(item.WriteJSON(nil))
}

func (item *MetadataGetTagMappingBootstrap) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFieldsMaskPresented bool

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
					return ErrorInvalidJSONWithDuplicatingKeys("metadata.getTagMappingBootstrap", "fields_mask")
				}
				if err := Json2ReadUint32(in, &item.FieldsMask); err != nil {
					return err
				}
				propFieldsMaskPresented = true
			default:
				return ErrorInvalidJSONExcessElement("metadata.getTagMappingBootstrap", key)
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
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *MetadataGetTagMappingBootstrap) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *MetadataGetTagMappingBootstrap) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *MetadataGetTagMappingBootstrap) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexFieldsMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"fields_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldsMask)
	if (item.FieldsMask != 0) == false {
		w = w[:backupIndexFieldsMask]
	}
	return append(w, '}')
}

func (item *MetadataGetTagMappingBootstrap) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *MetadataGetTagMappingBootstrap) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("metadata.getTagMappingBootstrap", err.Error())
	}
	return nil
}
