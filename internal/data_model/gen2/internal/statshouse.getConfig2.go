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

type StatshouseGetConfig2 struct {
	FieldsMask uint32
	Header     StatshouseCommonProxyHeader
	Cluster    string
	// Ts (TrueType) // Conditional: item.FieldsMask.0
}

func (StatshouseGetConfig2) TLName() string { return "statshouse.getConfig2" }
func (StatshouseGetConfig2) TLTag() uint32  { return 0x4285ff57 }

func (item *StatshouseGetConfig2) SetTs(v bool) {
	if v {
		item.FieldsMask |= 1 << 0
	} else {
		item.FieldsMask &^= 1 << 0
	}
}
func (item StatshouseGetConfig2) IsSetTs() bool { return item.FieldsMask&(1<<0) != 0 }

func (item *StatshouseGetConfig2) Reset() {
	item.FieldsMask = 0
	item.Header.Reset()
	item.Cluster = ""
}

func (item *StatshouseGetConfig2) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = item.Header.Read(w, item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = basictl.StringRead(w, &item.Cluster); err != nil {
		return w, err
	}
	return w, nil
}

// This method is general version of Write, use it instead!
func (item *StatshouseGetConfig2) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *StatshouseGetConfig2) Write(w []byte) []byte {
	w = basictl.NatWrite(w, item.FieldsMask)
	w = item.Header.Write(w, item.FieldsMask)
	w = basictl.StringWrite(w, item.Cluster)
	return w
}

func (item *StatshouseGetConfig2) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x4285ff57); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *StatshouseGetConfig2) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *StatshouseGetConfig2) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x4285ff57)
	return item.Write(w)
}

func (item *StatshouseGetConfig2) ReadResult(w []byte, ret *StatshouseGetConfigResult) (_ []byte, err error) {
	return ret.ReadBoxed(w, item.FieldsMask)
}

func (item *StatshouseGetConfig2) WriteResult(w []byte, ret StatshouseGetConfigResult) (_ []byte, err error) {
	w = ret.WriteBoxed(w, item.FieldsMask)
	return w, nil
}

func (item *StatshouseGetConfig2) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *StatshouseGetConfigResult) error {
	if err := ret.ReadJSON(legacyTypeNames, in, item.FieldsMask); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseGetConfig2) WriteResultJSON(w []byte, ret StatshouseGetConfigResult) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *StatshouseGetConfig2) writeResultJSON(newTypeNames bool, short bool, w []byte, ret StatshouseGetConfigResult) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w, item.FieldsMask)
	return w, nil
}

func (item *StatshouseGetConfig2) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret StatshouseGetConfigResult
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *StatshouseGetConfig2) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret StatshouseGetConfigResult
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *StatshouseGetConfig2) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret StatshouseGetConfigResult
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item StatshouseGetConfig2) String() string {
	return string(item.WriteJSON(nil))
}

func (item *StatshouseGetConfig2) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFieldsMaskPresented bool
	var rawHeader []byte
	var propClusterPresented bool
	var trueTypeTsPresented bool
	var trueTypeTsValue bool

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
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.getConfig2", "fields_mask")
				}
				if err := Json2ReadUint32(in, &item.FieldsMask); err != nil {
					return err
				}
				propFieldsMaskPresented = true
			case "header":
				if rawHeader != nil {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.getConfig2", "header")
				}
				rawHeader = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			case "cluster":
				if propClusterPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.getConfig2", "cluster")
				}
				if err := Json2ReadString(in, &item.Cluster); err != nil {
					return err
				}
				propClusterPresented = true
			case "ts":
				if trueTypeTsPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.getConfig2", "ts")
				}
				if err := Json2ReadBool(in, &trueTypeTsValue); err != nil {
					return err
				}
				trueTypeTsPresented = true
			default:
				return ErrorInvalidJSONExcessElement("statshouse.getConfig2", key)
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
	if !propClusterPresented {
		item.Cluster = ""
	}
	if trueTypeTsPresented {
		if trueTypeTsValue {
			item.FieldsMask |= 1 << 0
		}
	}
	var inHeaderPointer *basictl.JsonLexer
	inHeader := basictl.JsonLexer{Data: rawHeader}
	if rawHeader != nil {
		inHeaderPointer = &inHeader
	}
	if err := item.Header.ReadJSON(legacyTypeNames, inHeaderPointer, item.FieldsMask); err != nil {
		return err
	}

	// tries to set bit to zero if it is 1
	if trueTypeTsPresented && !trueTypeTsValue && (item.FieldsMask&(1<<0) != 0) {
		return ErrorInvalidJSON("statshouse.getConfig2", "fieldmask bit fields_mask.0 is indefinite because of the contradictions in values")
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *StatshouseGetConfig2) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *StatshouseGetConfig2) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *StatshouseGetConfig2) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexFieldsMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"fields_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldsMask)
	if (item.FieldsMask != 0) == false {
		w = w[:backupIndexFieldsMask]
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"header":`...)
	w = item.Header.WriteJSONOpt(newTypeNames, short, w, item.FieldsMask)
	backupIndexCluster := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"cluster":`...)
	w = basictl.JSONWriteString(w, item.Cluster)
	if (len(item.Cluster) != 0) == false {
		w = w[:backupIndexCluster]
	}
	if item.FieldsMask&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"ts":true`...)
	}
	return append(w, '}')
}

func (item *StatshouseGetConfig2) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *StatshouseGetConfig2) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("statshouse.getConfig2", err.Error())
	}
	return nil
}

type StatshouseGetConfig2Bytes struct {
	FieldsMask uint32
	Header     StatshouseCommonProxyHeaderBytes
	Cluster    []byte
	// Ts (TrueType) // Conditional: item.FieldsMask.0
}

func (StatshouseGetConfig2Bytes) TLName() string { return "statshouse.getConfig2" }
func (StatshouseGetConfig2Bytes) TLTag() uint32  { return 0x4285ff57 }

func (item *StatshouseGetConfig2Bytes) SetTs(v bool) {
	if v {
		item.FieldsMask |= 1 << 0
	} else {
		item.FieldsMask &^= 1 << 0
	}
}
func (item StatshouseGetConfig2Bytes) IsSetTs() bool { return item.FieldsMask&(1<<0) != 0 }

func (item *StatshouseGetConfig2Bytes) Reset() {
	item.FieldsMask = 0
	item.Header.Reset()
	item.Cluster = item.Cluster[:0]
}

func (item *StatshouseGetConfig2Bytes) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = item.Header.Read(w, item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = basictl.StringReadBytes(w, &item.Cluster); err != nil {
		return w, err
	}
	return w, nil
}

// This method is general version of Write, use it instead!
func (item *StatshouseGetConfig2Bytes) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *StatshouseGetConfig2Bytes) Write(w []byte) []byte {
	w = basictl.NatWrite(w, item.FieldsMask)
	w = item.Header.Write(w, item.FieldsMask)
	w = basictl.StringWriteBytes(w, item.Cluster)
	return w
}

func (item *StatshouseGetConfig2Bytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x4285ff57); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *StatshouseGetConfig2Bytes) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *StatshouseGetConfig2Bytes) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x4285ff57)
	return item.Write(w)
}

func (item *StatshouseGetConfig2Bytes) ReadResult(w []byte, ret *StatshouseGetConfigResultBytes) (_ []byte, err error) {
	return ret.ReadBoxed(w, item.FieldsMask)
}

func (item *StatshouseGetConfig2Bytes) WriteResult(w []byte, ret StatshouseGetConfigResultBytes) (_ []byte, err error) {
	w = ret.WriteBoxed(w, item.FieldsMask)
	return w, nil
}

func (item *StatshouseGetConfig2Bytes) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *StatshouseGetConfigResultBytes) error {
	if err := ret.ReadJSON(legacyTypeNames, in, item.FieldsMask); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseGetConfig2Bytes) WriteResultJSON(w []byte, ret StatshouseGetConfigResultBytes) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *StatshouseGetConfig2Bytes) writeResultJSON(newTypeNames bool, short bool, w []byte, ret StatshouseGetConfigResultBytes) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w, item.FieldsMask)
	return w, nil
}

func (item *StatshouseGetConfig2Bytes) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret StatshouseGetConfigResultBytes
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *StatshouseGetConfig2Bytes) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret StatshouseGetConfigResultBytes
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *StatshouseGetConfig2Bytes) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret StatshouseGetConfigResultBytes
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item StatshouseGetConfig2Bytes) String() string {
	return string(item.WriteJSON(nil))
}

func (item *StatshouseGetConfig2Bytes) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFieldsMaskPresented bool
	var rawHeader []byte
	var propClusterPresented bool
	var trueTypeTsPresented bool
	var trueTypeTsValue bool

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
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.getConfig2", "fields_mask")
				}
				if err := Json2ReadUint32(in, &item.FieldsMask); err != nil {
					return err
				}
				propFieldsMaskPresented = true
			case "header":
				if rawHeader != nil {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.getConfig2", "header")
				}
				rawHeader = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			case "cluster":
				if propClusterPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.getConfig2", "cluster")
				}
				if err := Json2ReadStringBytes(in, &item.Cluster); err != nil {
					return err
				}
				propClusterPresented = true
			case "ts":
				if trueTypeTsPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.getConfig2", "ts")
				}
				if err := Json2ReadBool(in, &trueTypeTsValue); err != nil {
					return err
				}
				trueTypeTsPresented = true
			default:
				return ErrorInvalidJSONExcessElement("statshouse.getConfig2", key)
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
	if !propClusterPresented {
		item.Cluster = item.Cluster[:0]
	}
	if trueTypeTsPresented {
		if trueTypeTsValue {
			item.FieldsMask |= 1 << 0
		}
	}
	var inHeaderPointer *basictl.JsonLexer
	inHeader := basictl.JsonLexer{Data: rawHeader}
	if rawHeader != nil {
		inHeaderPointer = &inHeader
	}
	if err := item.Header.ReadJSON(legacyTypeNames, inHeaderPointer, item.FieldsMask); err != nil {
		return err
	}

	// tries to set bit to zero if it is 1
	if trueTypeTsPresented && !trueTypeTsValue && (item.FieldsMask&(1<<0) != 0) {
		return ErrorInvalidJSON("statshouse.getConfig2", "fieldmask bit fields_mask.0 is indefinite because of the contradictions in values")
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *StatshouseGetConfig2Bytes) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *StatshouseGetConfig2Bytes) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *StatshouseGetConfig2Bytes) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexFieldsMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"fields_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldsMask)
	if (item.FieldsMask != 0) == false {
		w = w[:backupIndexFieldsMask]
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"header":`...)
	w = item.Header.WriteJSONOpt(newTypeNames, short, w, item.FieldsMask)
	backupIndexCluster := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"cluster":`...)
	w = basictl.JSONWriteStringBytes(w, item.Cluster)
	if (len(item.Cluster) != 0) == false {
		w = w[:backupIndexCluster]
	}
	if item.FieldsMask&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"ts":true`...)
	}
	return append(w, '}')
}

func (item *StatshouseGetConfig2Bytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *StatshouseGetConfig2Bytes) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("statshouse.getConfig2", err.Error())
	}
	return nil
}
