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

type EngineIsProduction struct {
}

func (EngineIsProduction) TLName() string { return "engine.isProduction" }
func (EngineIsProduction) TLTag() uint32  { return 0xccdea0ac }

func (item *EngineIsProduction) Reset() {}

func (item *EngineIsProduction) Read(w []byte) (_ []byte, err error) { return w, nil }

// This method is general version of Write, use it instead!
func (item *EngineIsProduction) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *EngineIsProduction) Write(w []byte) []byte {
	return w
}

func (item *EngineIsProduction) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xccdea0ac); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *EngineIsProduction) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *EngineIsProduction) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xccdea0ac)
	return item.Write(w)
}

func (item *EngineIsProduction) ReadResult(w []byte, ret *bool) (_ []byte, err error) {
	return BoolReadBoxed(w, ret)
}

func (item *EngineIsProduction) WriteResult(w []byte, ret bool) (_ []byte, err error) {
	w = BoolWriteBoxed(w, ret)
	return w, nil
}

func (item *EngineIsProduction) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *bool) error {
	if err := Json2ReadBool(in, ret); err != nil {
		return err
	}
	return nil
}

func (item *EngineIsProduction) WriteResultJSON(w []byte, ret bool) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *EngineIsProduction) writeResultJSON(newTypeNames bool, short bool, w []byte, ret bool) (_ []byte, err error) {
	w = basictl.JSONWriteBool(w, ret)
	return w, nil
}

func (item *EngineIsProduction) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret bool
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *EngineIsProduction) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret bool
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *EngineIsProduction) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret bool
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item EngineIsProduction) String() string {
	return string(item.WriteJSON(nil))
}

func (item *EngineIsProduction) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			return ErrorInvalidJSON("engine.isProduction", "this object can't have properties")
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *EngineIsProduction) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *EngineIsProduction) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *EngineIsProduction) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	return append(w, '}')
}

func (item *EngineIsProduction) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *EngineIsProduction) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("engine.isProduction", err.Error())
	}
	return nil
}
