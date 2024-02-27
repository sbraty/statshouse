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

type StatshouseTestConnection2 struct {
	FieldsMask         uint32
	Header             StatshouseCommonProxyHeader
	Payload            string
	ResponseSize       int32
	ResponseTimeoutSec int32
}

func (StatshouseTestConnection2) TLName() string { return "statshouse.testConnection2" }
func (StatshouseTestConnection2) TLTag() uint32  { return 0x4285ff58 }

func (item *StatshouseTestConnection2) Reset() {
	item.FieldsMask = 0
	item.Header.Reset()
	item.Payload = ""
	item.ResponseSize = 0
	item.ResponseTimeoutSec = 0
}

func (item *StatshouseTestConnection2) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = item.Header.Read(w, item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = basictl.StringRead(w, &item.Payload); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.ResponseSize); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.ResponseTimeoutSec)
}

func (item *StatshouseTestConnection2) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	if w, err = item.Header.Write(w, item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = basictl.StringWrite(w, item.Payload); err != nil {
		return w, err
	}
	w = basictl.IntWrite(w, item.ResponseSize)
	return basictl.IntWrite(w, item.ResponseTimeoutSec), nil
}

func (item *StatshouseTestConnection2) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x4285ff58); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseTestConnection2) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x4285ff58)
	return item.Write(w)
}

func (item *StatshouseTestConnection2) ReadResult(w []byte, ret *string) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xb5286e24); err != nil {
		return w, err
	}
	return basictl.StringRead(w, ret)
}

func (item *StatshouseTestConnection2) WriteResult(w []byte, ret string) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0xb5286e24)
	return basictl.StringWrite(w, ret)
}

func (item *StatshouseTestConnection2) ReadResultJSON(j interface{}, ret *string) error {
	if err := JsonReadString(j, ret); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseTestConnection2) WriteResultJSON(w []byte, ret string) (_ []byte, err error) {
	return item.writeResultJSON(false, w, ret)
}

func (item *StatshouseTestConnection2) writeResultJSON(short bool, w []byte, ret string) (_ []byte, err error) {
	w = basictl.JSONWriteString(w, ret)
	return w, nil
}

func (item *StatshouseTestConnection2) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret string
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *StatshouseTestConnection2) ReadResultWriteResultJSONShort(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret string
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(true, w, ret)
	return r, w, err
}

func (item *StatshouseTestConnection2) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := JsonBytesToInterface(r)
	if err != nil {
		return r, w, ErrorInvalidJSON("statshouse.testConnection2", err.Error())
	}
	var ret string
	if err = item.ReadResultJSON(j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item StatshouseTestConnection2) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func StatshouseTestConnection2__ReadJSON(item *StatshouseTestConnection2, j interface{}) error {
	return item.readJSON(j)
}
func (item *StatshouseTestConnection2) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouse.testConnection2", "expected json object")
	}
	_jFieldsMask := _jm["fields_mask"]
	delete(_jm, "fields_mask")
	if err := JsonReadUint32(_jFieldsMask, &item.FieldsMask); err != nil {
		return err
	}
	_jHeader := _jm["header"]
	delete(_jm, "header")
	_jPayload := _jm["payload"]
	delete(_jm, "payload")
	if err := JsonReadString(_jPayload, &item.Payload); err != nil {
		return err
	}
	_jResponseSize := _jm["response_size"]
	delete(_jm, "response_size")
	if err := JsonReadInt32(_jResponseSize, &item.ResponseSize); err != nil {
		return err
	}
	_jResponseTimeoutSec := _jm["response_timeout_sec"]
	delete(_jm, "response_timeout_sec")
	if err := JsonReadInt32(_jResponseTimeoutSec, &item.ResponseTimeoutSec); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouse.testConnection2", k)
	}
	if err := StatshouseCommonProxyHeader__ReadJSON(&item.Header, _jHeader, item.FieldsMask); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseTestConnection2) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *StatshouseTestConnection2) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldsMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"fields_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldsMask)
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"header":`...)
	if w, err = item.Header.WriteJSONOpt(short, w, item.FieldsMask); err != nil {
		return w, err
	}
	if len(item.Payload) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"payload":`...)
		w = basictl.JSONWriteString(w, item.Payload)
	}
	if item.ResponseSize != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"response_size":`...)
		w = basictl.JSONWriteInt32(w, item.ResponseSize)
	}
	if item.ResponseTimeoutSec != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"response_timeout_sec":`...)
		w = basictl.JSONWriteInt32(w, item.ResponseTimeoutSec)
	}
	return append(w, '}'), nil
}

func (item *StatshouseTestConnection2) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseTestConnection2) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("statshouse.testConnection2", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("statshouse.testConnection2", err.Error())
	}
	return nil
}

type StatshouseTestConnection2Bytes struct {
	FieldsMask         uint32
	Header             StatshouseCommonProxyHeaderBytes
	Payload            []byte
	ResponseSize       int32
	ResponseTimeoutSec int32
}

func (StatshouseTestConnection2Bytes) TLName() string { return "statshouse.testConnection2" }
func (StatshouseTestConnection2Bytes) TLTag() uint32  { return 0x4285ff58 }

func (item *StatshouseTestConnection2Bytes) Reset() {
	item.FieldsMask = 0
	item.Header.Reset()
	item.Payload = item.Payload[:0]
	item.ResponseSize = 0
	item.ResponseTimeoutSec = 0
}

func (item *StatshouseTestConnection2Bytes) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = item.Header.Read(w, item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = basictl.StringReadBytes(w, &item.Payload); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.ResponseSize); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.ResponseTimeoutSec)
}

func (item *StatshouseTestConnection2Bytes) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	if w, err = item.Header.Write(w, item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = basictl.StringWriteBytes(w, item.Payload); err != nil {
		return w, err
	}
	w = basictl.IntWrite(w, item.ResponseSize)
	return basictl.IntWrite(w, item.ResponseTimeoutSec), nil
}

func (item *StatshouseTestConnection2Bytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x4285ff58); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseTestConnection2Bytes) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x4285ff58)
	return item.Write(w)
}

func (item *StatshouseTestConnection2Bytes) ReadResult(w []byte, ret *[]byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xb5286e24); err != nil {
		return w, err
	}
	return basictl.StringReadBytes(w, ret)
}

func (item *StatshouseTestConnection2Bytes) WriteResult(w []byte, ret []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0xb5286e24)
	return basictl.StringWriteBytes(w, ret)
}

func (item *StatshouseTestConnection2Bytes) ReadResultJSON(j interface{}, ret *[]byte) error {
	if err := JsonReadStringBytes(j, ret); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseTestConnection2Bytes) WriteResultJSON(w []byte, ret []byte) (_ []byte, err error) {
	return item.writeResultJSON(false, w, ret)
}

func (item *StatshouseTestConnection2Bytes) writeResultJSON(short bool, w []byte, ret []byte) (_ []byte, err error) {
	w = basictl.JSONWriteStringBytes(w, ret)
	return w, nil
}

func (item *StatshouseTestConnection2Bytes) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret []byte
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *StatshouseTestConnection2Bytes) ReadResultWriteResultJSONShort(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret []byte
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(true, w, ret)
	return r, w, err
}

func (item *StatshouseTestConnection2Bytes) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := JsonBytesToInterface(r)
	if err != nil {
		return r, w, ErrorInvalidJSON("statshouse.testConnection2", err.Error())
	}
	var ret []byte
	if err = item.ReadResultJSON(j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item StatshouseTestConnection2Bytes) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func StatshouseTestConnection2Bytes__ReadJSON(item *StatshouseTestConnection2Bytes, j interface{}) error {
	return item.readJSON(j)
}
func (item *StatshouseTestConnection2Bytes) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouse.testConnection2", "expected json object")
	}
	_jFieldsMask := _jm["fields_mask"]
	delete(_jm, "fields_mask")
	if err := JsonReadUint32(_jFieldsMask, &item.FieldsMask); err != nil {
		return err
	}
	_jHeader := _jm["header"]
	delete(_jm, "header")
	_jPayload := _jm["payload"]
	delete(_jm, "payload")
	if err := JsonReadStringBytes(_jPayload, &item.Payload); err != nil {
		return err
	}
	_jResponseSize := _jm["response_size"]
	delete(_jm, "response_size")
	if err := JsonReadInt32(_jResponseSize, &item.ResponseSize); err != nil {
		return err
	}
	_jResponseTimeoutSec := _jm["response_timeout_sec"]
	delete(_jm, "response_timeout_sec")
	if err := JsonReadInt32(_jResponseTimeoutSec, &item.ResponseTimeoutSec); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouse.testConnection2", k)
	}
	if err := StatshouseCommonProxyHeaderBytes__ReadJSON(&item.Header, _jHeader, item.FieldsMask); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseTestConnection2Bytes) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *StatshouseTestConnection2Bytes) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldsMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"fields_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldsMask)
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"header":`...)
	if w, err = item.Header.WriteJSONOpt(short, w, item.FieldsMask); err != nil {
		return w, err
	}
	if len(item.Payload) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"payload":`...)
		w = basictl.JSONWriteStringBytes(w, item.Payload)
	}
	if item.ResponseSize != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"response_size":`...)
		w = basictl.JSONWriteInt32(w, item.ResponseSize)
	}
	if item.ResponseTimeoutSec != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"response_timeout_sec":`...)
		w = basictl.JSONWriteInt32(w, item.ResponseTimeoutSec)
	}
	return append(w, '}'), nil
}

func (item *StatshouseTestConnection2Bytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseTestConnection2Bytes) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("statshouse.testConnection2", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("statshouse.testConnection2", err.Error())
	}
	return nil
}
