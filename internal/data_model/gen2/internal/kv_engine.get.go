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

type KvEngineGet struct {
	Key int64
}

func (KvEngineGet) TLName() string { return "kv_engine.get" }
func (KvEngineGet) TLTag() uint32  { return 0x1c7349bb }

func (item *KvEngineGet) Reset() {
	item.Key = 0
}

func (item *KvEngineGet) Read(w []byte) (_ []byte, err error) {
	return basictl.LongRead(w, &item.Key)
}

// This method is general version of Write, use it instead!
func (item *KvEngineGet) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *KvEngineGet) Write(w []byte) []byte {
	w = basictl.LongWrite(w, item.Key)
	return w
}

func (item *KvEngineGet) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1c7349bb); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *KvEngineGet) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *KvEngineGet) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x1c7349bb)
	return item.Write(w)
}

func (item *KvEngineGet) ReadResult(w []byte, ret *KvEngineGetResponse) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *KvEngineGet) WriteResult(w []byte, ret KvEngineGetResponse) (_ []byte, err error) {
	w = ret.WriteBoxed(w)
	return w, nil
}

func (item *KvEngineGet) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *KvEngineGetResponse) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *KvEngineGet) WriteResultJSON(w []byte, ret KvEngineGetResponse) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *KvEngineGet) writeResultJSON(newTypeNames bool, short bool, w []byte, ret KvEngineGetResponse) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w)
	return w, nil
}

func (item *KvEngineGet) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret KvEngineGetResponse
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *KvEngineGet) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret KvEngineGetResponse
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *KvEngineGet) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret KvEngineGetResponse
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item KvEngineGet) String() string {
	return string(item.WriteJSON(nil))
}

func (item *KvEngineGet) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propKeyPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "key":
				if propKeyPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("kv_engine.get", "key")
				}
				if err := Json2ReadInt64(in, &item.Key); err != nil {
					return err
				}
				propKeyPresented = true
			default:
				return ErrorInvalidJSONExcessElement("kv_engine.get", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propKeyPresented {
		item.Key = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *KvEngineGet) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *KvEngineGet) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *KvEngineGet) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexKey := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"key":`...)
	w = basictl.JSONWriteInt64(w, item.Key)
	if (item.Key != 0) == false {
		w = w[:backupIndexKey]
	}
	return append(w, '}')
}

func (item *KvEngineGet) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *KvEngineGet) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("kv_engine.get", err.Error())
	}
	return nil
}
