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

type Tuple8 [8]uint32

func (Tuple8) TLName() string { return "tuple" }
func (Tuple8) TLTag() uint32  { return 0x9770768a }

func (item *Tuple8) Reset() {
	ptr := (*[8]uint32)(item)
	BuiltinTuple8Reset(ptr)
}

func (item *Tuple8) Read(w []byte) (_ []byte, err error) {
	ptr := (*[8]uint32)(item)
	return BuiltinTuple8Read(w, ptr)
}

func (item *Tuple8) Write(w []byte) (_ []byte, err error) {
	ptr := (*[8]uint32)(item)
	return BuiltinTuple8Write(w, ptr)
}

func (item *Tuple8) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9770768a); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *Tuple8) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x9770768a)
	return item.Write(w)
}

func (item Tuple8) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *Tuple8) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*[8]uint32)(item)
	if err := BuiltinTuple8ReadJSON(legacyTypeNames, in, ptr); err != nil {
		return err
	}
	return nil
}

func (item *Tuple8) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}

func (item *Tuple8) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	ptr := (*[8]uint32)(item)
	if w, err = BuiltinTuple8WriteJSONOpt(newTypeNames, short, w, ptr); err != nil {
		return w, err
	}
	return w, nil
}
func (item *Tuple8) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *Tuple8) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("tuple", err.Error())
	}
	return nil
}
