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

type EngineSwitchToMasterMode struct {
}

func (EngineSwitchToMasterMode) TLName() string { return "engine.switchToMasterMode" }
func (EngineSwitchToMasterMode) TLTag() uint32  { return 0x8cdcb5f9 }

func (item *EngineSwitchToMasterMode) Reset() {}

func (item *EngineSwitchToMasterMode) Read(w []byte) (_ []byte, err error) { return w, nil }

func (item *EngineSwitchToMasterMode) Write(w []byte) (_ []byte, err error) { return w, nil }

func (item *EngineSwitchToMasterMode) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x8cdcb5f9); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *EngineSwitchToMasterMode) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x8cdcb5f9)
	return item.Write(w)
}

func (item *EngineSwitchToMasterMode) ReadResult(w []byte, ret *EngineSwitchMasterReplicaModeResult) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *EngineSwitchToMasterMode) WriteResult(w []byte, ret EngineSwitchMasterReplicaModeResult) (_ []byte, err error) {
	return ret.WriteBoxed(w)
}

func (item *EngineSwitchToMasterMode) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *EngineSwitchMasterReplicaModeResult) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *EngineSwitchToMasterMode) WriteResultJSON(w []byte, ret EngineSwitchMasterReplicaModeResult) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *EngineSwitchToMasterMode) writeResultJSON(newTypeNames bool, short bool, w []byte, ret EngineSwitchMasterReplicaModeResult) (_ []byte, err error) {
	if w, err = ret.WriteJSONOpt(newTypeNames, short, w); err != nil {
		return w, err
	}
	return w, nil
}

func (item *EngineSwitchToMasterMode) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret EngineSwitchMasterReplicaModeResult
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *EngineSwitchToMasterMode) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret EngineSwitchMasterReplicaModeResult
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *EngineSwitchToMasterMode) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret EngineSwitchMasterReplicaModeResult
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item EngineSwitchToMasterMode) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *EngineSwitchToMasterMode) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			return ErrorInvalidJSON("engine.switchToMasterMode", "this object can't have properties")
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	return nil
}

func (item *EngineSwitchToMasterMode) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *EngineSwitchToMasterMode) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	return append(w, '}'), nil
}

func (item *EngineSwitchToMasterMode) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EngineSwitchToMasterMode) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("engine.switchToMasterMode", err.Error())
	}
	return nil
}
