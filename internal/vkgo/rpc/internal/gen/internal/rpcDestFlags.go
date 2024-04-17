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

type RpcDestFlags struct {
	Extra RpcInvokeReqExtra
}

func (RpcDestFlags) TLName() string { return "rpcDestFlags" }
func (RpcDestFlags) TLTag() uint32  { return 0xe352035e }

func (item *RpcDestFlags) Reset() {
	item.Extra.Reset()
}

func (item *RpcDestFlags) Read(w []byte) (_ []byte, err error) {
	return item.Extra.Read(w)
}

func (item *RpcDestFlags) Write(w []byte) (_ []byte, err error) {
	return item.Extra.Write(w)
}

func (item *RpcDestFlags) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xe352035e); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *RpcDestFlags) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xe352035e)
	return item.Write(w)
}

func (item RpcDestFlags) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *RpcDestFlags) ReadJSONLegacy(legacyTypeNames bool, j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("rpcDestFlags", "expected json object")
	}
	_jExtra := _jm["extra"]
	delete(_jm, "extra")
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("rpcDestFlags", k)
	}
	if err := item.Extra.ReadJSONLegacy(legacyTypeNames, _jExtra); err != nil {
		return err
	}
	return nil
}

func (item *RpcDestFlags) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *RpcDestFlags) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"extra":`...)
	if w, err = item.Extra.WriteJSONOpt(newTypeNames, short, w); err != nil {
		return w, err
	}
	return append(w, '}'), nil
}

func (item *RpcDestFlags) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *RpcDestFlags) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("rpcDestFlags", err.Error())
	}
	if err = item.ReadJSONLegacy(true, j); err != nil {
		return ErrorInvalidJSON("rpcDestFlags", err.Error())
	}
	return nil
}
