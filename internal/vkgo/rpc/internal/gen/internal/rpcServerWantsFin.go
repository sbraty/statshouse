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

type RpcServerWantsFin struct {
}

func (RpcServerWantsFin) TLName() string { return "rpcServerWantsFin" }
func (RpcServerWantsFin) TLTag() uint32  { return 0xa8ddbc46 }

func (item *RpcServerWantsFin) Reset() {}

func (item *RpcServerWantsFin) Read(w []byte) (_ []byte, err error) { return w, nil }

func (item *RpcServerWantsFin) Write(w []byte) (_ []byte, err error) { return w, nil }

func (item *RpcServerWantsFin) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xa8ddbc46); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *RpcServerWantsFin) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xa8ddbc46)
	return item.Write(w)
}

func (item RpcServerWantsFin) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *RpcServerWantsFin) ReadJSONLegacy(legacyTypeNames bool, j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("rpcServerWantsFin", "expected json object")
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("rpcServerWantsFin", k)
	}
	return nil
}

func (item *RpcServerWantsFin) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *RpcServerWantsFin) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	return append(w, '}'), nil
}

func (item *RpcServerWantsFin) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *RpcServerWantsFin) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("rpcServerWantsFin", err.Error())
	}
	if err = item.ReadJSONLegacy(true, j); err != nil {
		return ErrorInvalidJSON("rpcServerWantsFin", err.Error())
	}
	return nil
}
