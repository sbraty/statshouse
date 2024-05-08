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

type EngineDumpNextQueries struct {
	NumQueries int32
}

func (EngineDumpNextQueries) TLName() string { return "engine.dumpNextQueries" }
func (EngineDumpNextQueries) TLTag() uint32  { return 0xe65872ad }

func (item *EngineDumpNextQueries) Reset() {
	item.NumQueries = 0
}

func (item *EngineDumpNextQueries) Read(w []byte) (_ []byte, err error) {
	return basictl.IntRead(w, &item.NumQueries)
}

func (item *EngineDumpNextQueries) Write(w []byte) (_ []byte, err error) {
	return basictl.IntWrite(w, item.NumQueries), nil
}

func (item *EngineDumpNextQueries) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xe65872ad); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *EngineDumpNextQueries) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xe65872ad)
	return item.Write(w)
}

func (item *EngineDumpNextQueries) ReadResult(w []byte, ret *True) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *EngineDumpNextQueries) WriteResult(w []byte, ret True) (_ []byte, err error) {
	return ret.WriteBoxed(w)
}

func (item *EngineDumpNextQueries) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *True) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *EngineDumpNextQueries) WriteResultJSON(w []byte, ret True) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *EngineDumpNextQueries) writeResultJSON(newTypeNames bool, short bool, w []byte, ret True) (_ []byte, err error) {
	if w, err = ret.WriteJSONOpt(newTypeNames, short, w); err != nil {
		return w, err
	}
	return w, nil
}

func (item *EngineDumpNextQueries) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *EngineDumpNextQueries) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *EngineDumpNextQueries) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret True
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item EngineDumpNextQueries) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *EngineDumpNextQueries) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propNumQueriesPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "num_queries":
				if propNumQueriesPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("engine.dumpNextQueries", "num_queries")
				}
				if err := Json2ReadInt32(in, &item.NumQueries); err != nil {
					return err
				}
				propNumQueriesPresented = true
			default:
				return ErrorInvalidJSONExcessElement("engine.dumpNextQueries", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propNumQueriesPresented {
		item.NumQueries = 0
	}
	return nil
}

func (item *EngineDumpNextQueries) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *EngineDumpNextQueries) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexNumQueries := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"num_queries":`...)
	w = basictl.JSONWriteInt32(w, item.NumQueries)
	if (item.NumQueries != 0) == false {
		w = w[:backupIndexNumQueries]
	}
	return append(w, '}'), nil
}

func (item *EngineDumpNextQueries) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EngineDumpNextQueries) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("engine.dumpNextQueries", err.Error())
	}
	return nil
}
