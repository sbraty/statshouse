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

type True struct {
}

func (True) TLName() string { return "true" }
func (True) TLTag() uint32  { return 0x3fedd339 }

func (item *True) Reset()                              {}
func (item *True) Read(w []byte) ([]byte, error)       { return w, nil }
func (item *True) Write(w []byte) ([]byte, error)      { return w, nil }
func (item *True) ReadBoxed(w []byte) ([]byte, error)  { return basictl.NatReadExactTag(w, 0x3fedd339) }
func (item *True) WriteBoxed(w []byte) ([]byte, error) { return basictl.NatWrite(w, 0x3fedd339), nil }

func (item True) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func True__ReadJSON(item *True, j interface{}) error { return item.readJSON(j) }
func (item *True) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("true", "expected json object")
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("true", k)
	}
	return nil
}

func (item *True) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *True) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	return append(w, '}'), nil
}

func (item *True) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *True) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("true", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("true", err.Error())
	}
	return nil
}
