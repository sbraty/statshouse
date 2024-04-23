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

type RpcPong struct {
	PingId int64
}

func (RpcPong) TLName() string { return "rpcPong" }
func (RpcPong) TLTag() uint32  { return 0x8430eaa7 }

func (item *RpcPong) Reset() {
	item.PingId = 0
}

func (item *RpcPong) Read(w []byte) (_ []byte, err error) {
	return basictl.LongRead(w, &item.PingId)
}

func (item *RpcPong) Write(w []byte) (_ []byte, err error) {
	return basictl.LongWrite(w, item.PingId), nil
}

func (item *RpcPong) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x8430eaa7); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *RpcPong) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x8430eaa7)
	return item.Write(w)
}

func (item RpcPong) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *RpcPong) ReadJSONLegacy(legacyTypeNames bool, j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("rpcPong", "expected json object")
	}
	_jPingId := _jm["ping_id"]
	delete(_jm, "ping_id")
	if err := JsonReadInt64(_jPingId, &item.PingId); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("rpcPong", k)
	}
	return nil
}

func (item *RpcPong) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propPingIdPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "ping_id":
				if propPingIdPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("rpcPong", "ping_id")
				}
				if err := Json2ReadInt64(in, &item.PingId); err != nil {
					return err
				}
				propPingIdPresented = true
			default:
				return ErrorInvalidJSONExcessElement("rpcPong", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propPingIdPresented {
		item.PingId = 0
	}
	return nil
}

func (item *RpcPong) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *RpcPong) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.PingId != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"ping_id":`...)
		w = basictl.JSONWriteInt64(w, item.PingId)
	}
	return append(w, '}'), nil
}

func (item *RpcPong) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *RpcPong) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("rpcPong", err.Error())
	}
	if err = item.ReadJSONLegacy(true, j); err != nil {
		return ErrorInvalidJSON("rpcPong", err.Error())
	}
	return nil
}
