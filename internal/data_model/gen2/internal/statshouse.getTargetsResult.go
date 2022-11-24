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

type StatshouseGetTargetsResult struct {
	Targets []StatshousePromTarget
	Hash    string
}

func (StatshouseGetTargetsResult) TLName() string { return "statshouse.getTargetsResult" }
func (StatshouseGetTargetsResult) TLTag() uint32  { return 0x51ac86df }

func (item *StatshouseGetTargetsResult) Reset() {
	item.Targets = item.Targets[:0]
	item.Hash = ""
}

func (item *StatshouseGetTargetsResult) Read(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	if w, err = VectorStatshousePromTarget0Read(w, &item.Targets); err != nil {
		return w, err
	}
	return basictl.StringRead(w, &item.Hash)
}

func (item *StatshouseGetTargetsResult) Write(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	if w, err = VectorStatshousePromTarget0Write(w, item.Targets); err != nil {
		return w, err
	}
	return basictl.StringWrite(w, item.Hash)
}

func (item *StatshouseGetTargetsResult) ReadBoxed(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x51ac86df); err != nil {
		return w, err
	}
	return item.Read(w, nat_fields_mask)
}

func (item *StatshouseGetTargetsResult) WriteBoxed(w []byte, nat_fields_mask uint32) ([]byte, error) {
	w = basictl.NatWrite(w, 0x51ac86df)
	return item.Write(w, nat_fields_mask)
}

func StatshouseGetTargetsResult__ReadJSON(item *StatshouseGetTargetsResult, j interface{}, nat_fields_mask uint32) error {
	return item.readJSON(j, nat_fields_mask)
}
func (item *StatshouseGetTargetsResult) readJSON(j interface{}, nat_fields_mask uint32) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouse.getTargetsResult", "expected json object")
	}
	_jTargets := _jm["targets"]
	delete(_jm, "targets")
	_jHash := _jm["hash"]
	delete(_jm, "hash")
	if err := JsonReadString(_jHash, &item.Hash); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouse.getTargetsResult", k)
	}
	if err := VectorStatshousePromTarget0ReadJSON(_jTargets, &item.Targets); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseGetTargetsResult) WriteJSON(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	w = append(w, '{')
	if len(item.Targets) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"targets":`...)
		if w, err = VectorStatshousePromTarget0WriteJSON(w, item.Targets); err != nil {
			return w, err
		}
	}
	if len(item.Hash) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"hash":`...)
		w = basictl.JSONWriteString(w, item.Hash)
	}
	return append(w, '}'), nil
}

type StatshouseGetTargetsResultBytes struct {
	Targets []StatshousePromTargetBytes
	Hash    []byte
}

func (StatshouseGetTargetsResultBytes) TLName() string { return "statshouse.getTargetsResult" }
func (StatshouseGetTargetsResultBytes) TLTag() uint32  { return 0x51ac86df }

func (item *StatshouseGetTargetsResultBytes) Reset() {
	item.Targets = item.Targets[:0]
	item.Hash = item.Hash[:0]
}

func (item *StatshouseGetTargetsResultBytes) Read(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	if w, err = VectorStatshousePromTarget0BytesRead(w, &item.Targets); err != nil {
		return w, err
	}
	return basictl.StringReadBytes(w, &item.Hash)
}

func (item *StatshouseGetTargetsResultBytes) Write(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	if w, err = VectorStatshousePromTarget0BytesWrite(w, item.Targets); err != nil {
		return w, err
	}
	return basictl.StringWriteBytes(w, item.Hash)
}

func (item *StatshouseGetTargetsResultBytes) ReadBoxed(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x51ac86df); err != nil {
		return w, err
	}
	return item.Read(w, nat_fields_mask)
}

func (item *StatshouseGetTargetsResultBytes) WriteBoxed(w []byte, nat_fields_mask uint32) ([]byte, error) {
	w = basictl.NatWrite(w, 0x51ac86df)
	return item.Write(w, nat_fields_mask)
}

func StatshouseGetTargetsResultBytes__ReadJSON(item *StatshouseGetTargetsResultBytes, j interface{}, nat_fields_mask uint32) error {
	return item.readJSON(j, nat_fields_mask)
}
func (item *StatshouseGetTargetsResultBytes) readJSON(j interface{}, nat_fields_mask uint32) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouse.getTargetsResult", "expected json object")
	}
	_jTargets := _jm["targets"]
	delete(_jm, "targets")
	_jHash := _jm["hash"]
	delete(_jm, "hash")
	if err := JsonReadStringBytes(_jHash, &item.Hash); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouse.getTargetsResult", k)
	}
	if err := VectorStatshousePromTarget0BytesReadJSON(_jTargets, &item.Targets); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseGetTargetsResultBytes) WriteJSON(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	w = append(w, '{')
	if len(item.Targets) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"targets":`...)
		if w, err = VectorStatshousePromTarget0BytesWriteJSON(w, item.Targets); err != nil {
			return w, err
		}
	}
	if len(item.Hash) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"hash":`...)
		w = basictl.JSONWriteStringBytes(w, item.Hash)
	}
	return append(w, '}'), nil
}
