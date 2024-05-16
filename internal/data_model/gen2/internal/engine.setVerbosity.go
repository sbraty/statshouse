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

type EngineSetVerbosity struct {
	Verbosity int32
}

func (EngineSetVerbosity) TLName() string { return "engine.setVerbosity" }
func (EngineSetVerbosity) TLTag() uint32  { return 0x9d980926 }

func (item *EngineSetVerbosity) Reset() {
	item.Verbosity = 0
}

func (item *EngineSetVerbosity) Read(w []byte) (_ []byte, err error) {
	return basictl.IntRead(w, &item.Verbosity)
}

// This method is general version of Write, use it instead!
func (item *EngineSetVerbosity) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *EngineSetVerbosity) Write(w []byte) []byte {
	w = basictl.IntWrite(w, item.Verbosity)
	return w
}

func (item *EngineSetVerbosity) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9d980926); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *EngineSetVerbosity) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *EngineSetVerbosity) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x9d980926)
	return item.Write(w)
}

func (item *EngineSetVerbosity) ReadResult(w []byte, ret *True) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *EngineSetVerbosity) WriteResult(w []byte, ret True) (_ []byte, err error) {
	w = ret.WriteBoxed(w)
	return w, nil
}

func (item *EngineSetVerbosity) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *True) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *EngineSetVerbosity) WriteResultJSON(w []byte, ret True) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *EngineSetVerbosity) writeResultJSON(newTypeNames bool, short bool, w []byte, ret True) (_ []byte, err error) {
	w = ret.WriteJSONOpt(newTypeNames, short, w)
	return w, nil
}

func (item *EngineSetVerbosity) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *EngineSetVerbosity) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *EngineSetVerbosity) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret True
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item EngineSetVerbosity) String() string {
	return string(item.WriteJSON(nil))
}

func (item *EngineSetVerbosity) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propVerbosityPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "verbosity":
				if propVerbosityPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("engine.setVerbosity", "verbosity")
				}
				if err := Json2ReadInt32(in, &item.Verbosity); err != nil {
					return err
				}
				propVerbosityPresented = true
			default:
				return ErrorInvalidJSONExcessElement("engine.setVerbosity", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propVerbosityPresented {
		item.Verbosity = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *EngineSetVerbosity) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *EngineSetVerbosity) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *EngineSetVerbosity) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexVerbosity := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"verbosity":`...)
	w = basictl.JSONWriteInt32(w, item.Verbosity)
	if (item.Verbosity != 0) == false {
		w = w[:backupIndexVerbosity]
	}
	return append(w, '}')
}

func (item *EngineSetVerbosity) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *EngineSetVerbosity) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("engine.setVerbosity", err.Error())
	}
	return nil
}
