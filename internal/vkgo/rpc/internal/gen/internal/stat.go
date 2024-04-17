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

type Stat map[string]string

func (Stat) TLName() string { return "stat" }
func (Stat) TLTag() uint32  { return 0x9d56e6b2 }

func (item *Stat) Reset() {
	ptr := (*map[string]string)(item)
	BuiltinVectorDictionaryFieldStringReset(*ptr)
}

func (item *Stat) Read(w []byte) (_ []byte, err error) {
	ptr := (*map[string]string)(item)
	return BuiltinVectorDictionaryFieldStringRead(w, ptr)
}

func (item *Stat) Write(w []byte) (_ []byte, err error) {
	ptr := (*map[string]string)(item)
	return BuiltinVectorDictionaryFieldStringWrite(w, *ptr)
}

func (item *Stat) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9d56e6b2); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *Stat) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x9d56e6b2)
	return item.Write(w)
}

func (item Stat) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *Stat) ReadJSONLegacy(legacyTypeNames bool, j interface{}) error {
	ptr := (*map[string]string)(item)
	if err := BuiltinVectorDictionaryFieldStringReadJSONLegacy(legacyTypeNames, j, ptr); err != nil {
		return err
	}
	return nil
}

func (item *Stat) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}

func (item *Stat) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	ptr := (*map[string]string)(item)
	if w, err = BuiltinVectorDictionaryFieldStringWriteJSONOpt(newTypeNames, short, w, *ptr); err != nil {
		return w, err
	}
	return w, nil
}
func (item *Stat) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *Stat) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("stat", err.Error())
	}
	if err = item.ReadJSONLegacy(true, j); err != nil {
		return ErrorInvalidJSON("stat", err.Error())
	}
	return nil
}
