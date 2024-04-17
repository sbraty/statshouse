// Copyright 2024 V Kontakte LLC
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

type DictionaryLong map[string]int64

func (DictionaryLong) TLName() string { return "dictionary" }
func (DictionaryLong) TLTag() uint32  { return 0x1f4c618f }

func (item *DictionaryLong) Reset() {
	ptr := (*map[string]int64)(item)
	BuiltinVectorDictionaryFieldLongReset(*ptr)
}

func (item *DictionaryLong) Read(w []byte) (_ []byte, err error) {
	ptr := (*map[string]int64)(item)
	return BuiltinVectorDictionaryFieldLongRead(w, ptr)
}

func (item *DictionaryLong) Write(w []byte) (_ []byte, err error) {
	ptr := (*map[string]int64)(item)
	return BuiltinVectorDictionaryFieldLongWrite(w, *ptr)
}

func (item *DictionaryLong) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1f4c618f); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *DictionaryLong) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x1f4c618f)
	return item.Write(w)
}

func (item DictionaryLong) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *DictionaryLong) ReadJSONLegacy(legacyTypeNames bool, j interface{}) error {
	ptr := (*map[string]int64)(item)
	if err := BuiltinVectorDictionaryFieldLongReadJSONLegacy(legacyTypeNames, j, ptr); err != nil {
		return err
	}
	return nil
}

func (item *DictionaryLong) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}

func (item *DictionaryLong) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	ptr := (*map[string]int64)(item)
	if w, err = BuiltinVectorDictionaryFieldLongWriteJSONOpt(newTypeNames, short, w, *ptr); err != nil {
		return w, err
	}
	return w, nil
}
func (item *DictionaryLong) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *DictionaryLong) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("dictionary", err.Error())
	}
	if err = item.ReadJSONLegacy(true, j); err != nil {
		return ErrorInvalidJSON("dictionary", err.Error())
	}
	return nil
}

type DictionaryString map[string]string

func (DictionaryString) TLName() string { return "dictionary" }
func (DictionaryString) TLTag() uint32  { return 0x1f4c618f }

func (item *DictionaryString) Reset() {
	ptr := (*map[string]string)(item)
	BuiltinVectorDictionaryFieldStringReset(*ptr)
}

func (item *DictionaryString) Read(w []byte) (_ []byte, err error) {
	ptr := (*map[string]string)(item)
	return BuiltinVectorDictionaryFieldStringRead(w, ptr)
}

func (item *DictionaryString) Write(w []byte) (_ []byte, err error) {
	ptr := (*map[string]string)(item)
	return BuiltinVectorDictionaryFieldStringWrite(w, *ptr)
}

func (item *DictionaryString) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1f4c618f); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *DictionaryString) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x1f4c618f)
	return item.Write(w)
}

func (item DictionaryString) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *DictionaryString) ReadJSONLegacy(legacyTypeNames bool, j interface{}) error {
	ptr := (*map[string]string)(item)
	if err := BuiltinVectorDictionaryFieldStringReadJSONLegacy(legacyTypeNames, j, ptr); err != nil {
		return err
	}
	return nil
}

func (item *DictionaryString) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}

func (item *DictionaryString) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	ptr := (*map[string]string)(item)
	if w, err = BuiltinVectorDictionaryFieldStringWriteJSONOpt(newTypeNames, short, w, *ptr); err != nil {
		return w, err
	}
	return w, nil
}
func (item *DictionaryString) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *DictionaryString) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("dictionary", err.Error())
	}
	if err = item.ReadJSONLegacy(true, j); err != nil {
		return ErrorInvalidJSON("dictionary", err.Error())
	}
	return nil
}
