// Copyright 2022 V Kontakte LLC
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

type StatshouseGetMetrics2 struct {
	FieldsMask uint32
	Header     StatshouseCommonProxyHeader
	Version    string
}

func (StatshouseGetMetrics2) TLName() string { return "statshouse.getMetrics2" }
func (StatshouseGetMetrics2) TLTag() uint32  { return 0x4285ff54 }

func (item *StatshouseGetMetrics2) Reset() {
	item.FieldsMask = 0
	item.Header.Reset()
	item.Version = ""
}

func (item *StatshouseGetMetrics2) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = item.Header.Read(w, item.FieldsMask); err != nil {
		return w, err
	}
	return basictl.StringRead(w, &item.Version)
}

func (item *StatshouseGetMetrics2) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	if w, err = item.Header.Write(w, item.FieldsMask); err != nil {
		return w, err
	}
	return basictl.StringWrite(w, item.Version)
}

func (item *StatshouseGetMetrics2) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x4285ff54); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseGetMetrics2) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x4285ff54)
	return item.Write(w)
}

func (item *StatshouseGetMetrics2) ReadResult(w []byte, ret *StatshouseGetMetricsResult) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *StatshouseGetMetrics2) WriteResult(w []byte, ret StatshouseGetMetricsResult) (_ []byte, err error) {
	return ret.WriteBoxed(w)
}

func (item *StatshouseGetMetrics2) ReadResultJSON(j interface{}, ret *StatshouseGetMetricsResult) error {
	if err := StatshouseGetMetricsResult__ReadJSON(ret, j); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseGetMetrics2) WriteResultJSON(w []byte, ret StatshouseGetMetricsResult) (_ []byte, err error) {
	if w, err = ret.WriteJSON(w); err != nil {
		return w, err
	}
	return w, nil
}

func (item *StatshouseGetMetrics2) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret StatshouseGetMetricsResult
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *StatshouseGetMetrics2) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := JsonBytesToInterface(r)
	if err != nil {
		return r, w, ErrorInvalidJSON("statshouse.getMetrics2", err.Error())
	}
	var ret StatshouseGetMetricsResult
	if err = item.ReadResultJSON(j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item StatshouseGetMetrics2) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func StatshouseGetMetrics2__ReadJSON(item *StatshouseGetMetrics2, j interface{}) error {
	return item.readJSON(j)
}
func (item *StatshouseGetMetrics2) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouse.getMetrics2", "expected json object")
	}
	_jFieldsMask := _jm["fields_mask"]
	delete(_jm, "fields_mask")
	if err := JsonReadUint32(_jFieldsMask, &item.FieldsMask); err != nil {
		return err
	}
	_jHeader := _jm["header"]
	delete(_jm, "header")
	_jVersion := _jm["version"]
	delete(_jm, "version")
	if err := JsonReadString(_jVersion, &item.Version); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouse.getMetrics2", k)
	}
	if err := StatshouseCommonProxyHeader__ReadJSON(&item.Header, _jHeader, item.FieldsMask); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseGetMetrics2) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldsMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"fields_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldsMask)
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"header":`...)
	if w, err = item.Header.WriteJSON(w, item.FieldsMask); err != nil {
		return w, err
	}
	if len(item.Version) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"version":`...)
		w = basictl.JSONWriteString(w, item.Version)
	}
	return append(w, '}'), nil
}

func (item *StatshouseGetMetrics2) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseGetMetrics2) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("statshouse.getMetrics2", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("statshouse.getMetrics2", err.Error())
	}
	return nil
}

type StatshouseGetMetrics2Bytes struct {
	FieldsMask uint32
	Header     StatshouseCommonProxyHeaderBytes
	Version    []byte
}

func (StatshouseGetMetrics2Bytes) TLName() string { return "statshouse.getMetrics2" }
func (StatshouseGetMetrics2Bytes) TLTag() uint32  { return 0x4285ff54 }

func (item *StatshouseGetMetrics2Bytes) Reset() {
	item.FieldsMask = 0
	item.Header.Reset()
	item.Version = item.Version[:0]
}

func (item *StatshouseGetMetrics2Bytes) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = item.Header.Read(w, item.FieldsMask); err != nil {
		return w, err
	}
	return basictl.StringReadBytes(w, &item.Version)
}

func (item *StatshouseGetMetrics2Bytes) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	if w, err = item.Header.Write(w, item.FieldsMask); err != nil {
		return w, err
	}
	return basictl.StringWriteBytes(w, item.Version)
}

func (item *StatshouseGetMetrics2Bytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x4285ff54); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseGetMetrics2Bytes) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x4285ff54)
	return item.Write(w)
}

func (item *StatshouseGetMetrics2Bytes) ReadResult(w []byte, ret *StatshouseGetMetricsResultBytes) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *StatshouseGetMetrics2Bytes) WriteResult(w []byte, ret StatshouseGetMetricsResultBytes) (_ []byte, err error) {
	return ret.WriteBoxed(w)
}

func (item *StatshouseGetMetrics2Bytes) ReadResultJSON(j interface{}, ret *StatshouseGetMetricsResultBytes) error {
	if err := StatshouseGetMetricsResultBytes__ReadJSON(ret, j); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseGetMetrics2Bytes) WriteResultJSON(w []byte, ret StatshouseGetMetricsResultBytes) (_ []byte, err error) {
	if w, err = ret.WriteJSON(w); err != nil {
		return w, err
	}
	return w, nil
}

func (item *StatshouseGetMetrics2Bytes) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret StatshouseGetMetricsResultBytes
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *StatshouseGetMetrics2Bytes) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := JsonBytesToInterface(r)
	if err != nil {
		return r, w, ErrorInvalidJSON("statshouse.getMetrics2", err.Error())
	}
	var ret StatshouseGetMetricsResultBytes
	if err = item.ReadResultJSON(j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item StatshouseGetMetrics2Bytes) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func StatshouseGetMetrics2Bytes__ReadJSON(item *StatshouseGetMetrics2Bytes, j interface{}) error {
	return item.readJSON(j)
}
func (item *StatshouseGetMetrics2Bytes) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouse.getMetrics2", "expected json object")
	}
	_jFieldsMask := _jm["fields_mask"]
	delete(_jm, "fields_mask")
	if err := JsonReadUint32(_jFieldsMask, &item.FieldsMask); err != nil {
		return err
	}
	_jHeader := _jm["header"]
	delete(_jm, "header")
	_jVersion := _jm["version"]
	delete(_jm, "version")
	if err := JsonReadStringBytes(_jVersion, &item.Version); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouse.getMetrics2", k)
	}
	if err := StatshouseCommonProxyHeaderBytes__ReadJSON(&item.Header, _jHeader, item.FieldsMask); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseGetMetrics2Bytes) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldsMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"fields_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldsMask)
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"header":`...)
	if w, err = item.Header.WriteJSON(w, item.FieldsMask); err != nil {
		return w, err
	}
	if len(item.Version) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"version":`...)
		w = basictl.JSONWriteStringBytes(w, item.Version)
	}
	return append(w, '}'), nil
}

func (item *StatshouseGetMetrics2Bytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseGetMetrics2Bytes) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("statshouse.getMetrics2", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("statshouse.getMetrics2", err.Error())
	}
	return nil
}
