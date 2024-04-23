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

type ReqError struct {
	ErrorCode int32
	Error     string
}

func (ReqError) TLName() string { return "reqError" }
func (ReqError) TLTag() uint32  { return 0xb527877d }

func (item *ReqError) Reset() {
	item.ErrorCode = 0
	item.Error = ""
}

func (item *ReqError) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.IntRead(w, &item.ErrorCode); err != nil {
		return w, err
	}
	return basictl.StringRead(w, &item.Error)
}

func (item *ReqError) Write(w []byte) (_ []byte, err error) {
	w = basictl.IntWrite(w, item.ErrorCode)
	return basictl.StringWrite(w, item.Error), nil
}

func (item *ReqError) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xb527877d); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *ReqError) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xb527877d)
	return item.Write(w)
}

func (item ReqError) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *ReqError) ReadJSONLegacy(legacyTypeNames bool, j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("reqError", "expected json object")
	}
	_jErrorCode := _jm["error_code"]
	delete(_jm, "error_code")
	if err := JsonReadInt32(_jErrorCode, &item.ErrorCode); err != nil {
		return err
	}
	_jError := _jm["error"]
	delete(_jm, "error")
	if err := JsonReadString(_jError, &item.Error); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("reqError", k)
	}
	return nil
}

func (item *ReqError) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propErrorCodePresented bool
	var propErrorPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "error_code":
				if propErrorCodePresented {
					return ErrorInvalidJSONWithDuplicatingKeys("reqError", "error_code")
				}
				if err := Json2ReadInt32(in, &item.ErrorCode); err != nil {
					return err
				}
				propErrorCodePresented = true
			case "error":
				if propErrorPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("reqError", "error")
				}
				if err := Json2ReadString(in, &item.Error); err != nil {
					return err
				}
				propErrorPresented = true
			default:
				return ErrorInvalidJSONExcessElement("reqError", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propErrorCodePresented {
		item.ErrorCode = 0
	}
	if !propErrorPresented {
		item.Error = ""
	}
	return nil
}

func (item *ReqError) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *ReqError) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.ErrorCode != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"error_code":`...)
		w = basictl.JSONWriteInt32(w, item.ErrorCode)
	}
	if len(item.Error) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"error":`...)
		w = basictl.JSONWriteString(w, item.Error)
	}
	return append(w, '}'), nil
}

func (item *ReqError) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *ReqError) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("reqError", err.Error())
	}
	if err = item.ReadJSONLegacy(true, j); err != nil {
		return ErrorInvalidJSON("reqError", err.Error())
	}
	return nil
}
