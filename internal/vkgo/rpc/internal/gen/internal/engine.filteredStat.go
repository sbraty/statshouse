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

type EngineFilteredStat struct {
	StatNames []string
}

func (EngineFilteredStat) TLName() string { return "engine.filteredStat" }
func (EngineFilteredStat) TLTag() uint32  { return 0x594870d6 }

func (item *EngineFilteredStat) Reset() {
	item.StatNames = item.StatNames[:0]
}

func (item *EngineFilteredStat) Read(w []byte) (_ []byte, err error) {
	return BuiltinVectorStringRead(w, &item.StatNames)
}

func (item *EngineFilteredStat) Write(w []byte) (_ []byte, err error) {
	return BuiltinVectorStringWrite(w, item.StatNames)
}

func (item *EngineFilteredStat) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x594870d6); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *EngineFilteredStat) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x594870d6)
	return item.Write(w)
}

func (item *EngineFilteredStat) ReadResult(w []byte, ret *Stat) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *EngineFilteredStat) WriteResult(w []byte, ret Stat) (_ []byte, err error) {
	return ret.WriteBoxed(w)
}

func (item *EngineFilteredStat) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *Stat) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *EngineFilteredStat) WriteResultJSON(w []byte, ret Stat) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *EngineFilteredStat) writeResultJSON(newTypeNames bool, short bool, w []byte, ret Stat) (_ []byte, err error) {
	if w, err = ret.WriteJSONOpt(newTypeNames, short, w); err != nil {
		return w, err
	}
	return w, nil
}

func (item *EngineFilteredStat) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret Stat
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *EngineFilteredStat) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret Stat
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *EngineFilteredStat) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret Stat
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item EngineFilteredStat) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *EngineFilteredStat) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propStatNamesPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "stat_names":
				if propStatNamesPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("engine.filteredStat", "stat_names")
				}
				if err := BuiltinVectorStringReadJSON(legacyTypeNames, in, &item.StatNames); err != nil {
					return err
				}
				propStatNamesPresented = true
			default:
				return ErrorInvalidJSONExcessElement("engine.filteredStat", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propStatNamesPresented {
		item.StatNames = item.StatNames[:0]
	}
	return nil
}

func (item *EngineFilteredStat) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *EngineFilteredStat) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexStatNames := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"stat_names":`...)
	if w, err = BuiltinVectorStringWriteJSONOpt(newTypeNames, short, w, item.StatNames); err != nil {
		return w, err
	}
	if (len(item.StatNames) != 0) == false {
		w = w[:backupIndexStatNames]
	}
	return append(w, '}'), nil
}

func (item *EngineFilteredStat) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EngineFilteredStat) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("engine.filteredStat", err.Error())
	}
	return nil
}
