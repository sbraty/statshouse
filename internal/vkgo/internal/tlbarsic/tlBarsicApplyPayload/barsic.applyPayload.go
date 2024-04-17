// Copyright 2024 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlBarsicApplyPayload

import (
	"github.com/vkcom/statshouse/internal/vkgo/basictl"
	"github.com/vkcom/statshouse/internal/vkgo/internal"
	"github.com/vkcom/statshouse/internal/vkgo/internal/tl/tlTrue"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type BarsicApplyPayload struct {
	FieldsMask uint32
	// CommitASAP (TrueType) // Conditional: item.FieldsMask.0
	// Committed (TrueType) // Conditional: item.FieldsMask.1
	Offset  int64
	Payload string
}

func (BarsicApplyPayload) TLName() string { return "barsic.applyPayload" }
func (BarsicApplyPayload) TLTag() uint32  { return 0x8c981e13 }

func (item *BarsicApplyPayload) SetCommitASAP(v bool) {
	if v {
		item.FieldsMask |= 1 << 0
	} else {
		item.FieldsMask &^= 1 << 0
	}
}
func (item BarsicApplyPayload) IsSetCommitASAP() bool { return item.FieldsMask&(1<<0) != 0 }

func (item *BarsicApplyPayload) SetCommitted(v bool) {
	if v {
		item.FieldsMask |= 1 << 1
	} else {
		item.FieldsMask &^= 1 << 1
	}
}
func (item BarsicApplyPayload) IsSetCommitted() bool { return item.FieldsMask&(1<<1) != 0 }

func (item *BarsicApplyPayload) Reset() {
	item.FieldsMask = 0
	item.Offset = 0
	item.Payload = ""
}

func (item *BarsicApplyPayload) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.Offset); err != nil {
		return w, err
	}
	return basictl.StringRead(w, &item.Payload)
}

func (item *BarsicApplyPayload) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	w = basictl.LongWrite(w, item.Offset)
	return basictl.StringWrite(w, item.Payload), nil
}

func (item *BarsicApplyPayload) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x8c981e13); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *BarsicApplyPayload) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x8c981e13)
	return item.Write(w)
}

func (item *BarsicApplyPayload) ReadResult(w []byte, ret *tlTrue.True) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *BarsicApplyPayload) WriteResult(w []byte, ret tlTrue.True) (_ []byte, err error) {
	return ret.WriteBoxed(w)
}

func (item *BarsicApplyPayload) ReadResultJSON(legacyTypeNames bool, j interface{}, ret *tlTrue.True) error {
	if err := ret.ReadJSONLegacy(legacyTypeNames, j); err != nil {
		return err
	}
	return nil
}

func (item *BarsicApplyPayload) WriteResultJSON(w []byte, ret tlTrue.True) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *BarsicApplyPayload) writeResultJSON(newTypeNames bool, short bool, w []byte, ret tlTrue.True) (_ []byte, err error) {
	if w, err = ret.WriteJSONOpt(newTypeNames, short, w); err != nil {
		return w, err
	}
	return w, nil
}

func (item *BarsicApplyPayload) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlTrue.True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *BarsicApplyPayload) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlTrue.True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *BarsicApplyPayload) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := internal.JsonBytesToInterface(r)
	if err != nil {
		return r, w, internal.ErrorInvalidJSON("barsic.applyPayload", err.Error())
	}
	var ret tlTrue.True
	if err = item.ReadResultJSON(true, j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item BarsicApplyPayload) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *BarsicApplyPayload) ReadJSONLegacy(legacyTypeNames bool, j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return internal.ErrorInvalidJSON("barsic.applyPayload", "expected json object")
	}
	_jFieldsMask := _jm["fields_mask"]
	delete(_jm, "fields_mask")
	if err := internal.JsonReadUint32(_jFieldsMask, &item.FieldsMask); err != nil {
		return err
	}
	_jCommitASAP := _jm["commitASAP"]
	delete(_jm, "commitASAP")
	_jCommitted := _jm["committed"]
	delete(_jm, "committed")
	_jOffset := _jm["offset"]
	delete(_jm, "offset")
	if err := internal.JsonReadInt64(_jOffset, &item.Offset); err != nil {
		return err
	}
	_jPayload := _jm["payload"]
	delete(_jm, "payload")
	if err := internal.JsonReadString(_jPayload, &item.Payload); err != nil {
		return err
	}
	for k := range _jm {
		return internal.ErrorInvalidJSONExcessElement("barsic.applyPayload", k)
	}
	if _jCommitASAP != nil {
		_bit := false
		if err := internal.JsonReadBool(_jCommitASAP, &_bit); err != nil {
			return err
		}
		if _bit {
			item.FieldsMask |= 1 << 0
		} else {
			item.FieldsMask &^= 1 << 0
		}
	}
	if _jCommitted != nil {
		_bit := false
		if err := internal.JsonReadBool(_jCommitted, &_bit); err != nil {
			return err
		}
		if _bit {
			item.FieldsMask |= 1 << 1
		} else {
			item.FieldsMask &^= 1 << 1
		}
	}
	return nil
}

func (item *BarsicApplyPayload) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *BarsicApplyPayload) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldsMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"fields_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldsMask)
	}
	if item.FieldsMask&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"commitASAP":true`...)
	}
	if item.FieldsMask&(1<<1) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"committed":true`...)
	}
	if item.Offset != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"offset":`...)
		w = basictl.JSONWriteInt64(w, item.Offset)
	}
	if len(item.Payload) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"payload":`...)
		w = basictl.JSONWriteString(w, item.Payload)
	}
	return append(w, '}'), nil
}

func (item *BarsicApplyPayload) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *BarsicApplyPayload) UnmarshalJSON(b []byte) error {
	j, err := internal.JsonBytesToInterface(b)
	if err != nil {
		return internal.ErrorInvalidJSON("barsic.applyPayload", err.Error())
	}
	if err = item.ReadJSONLegacy(true, j); err != nil {
		return internal.ErrorInvalidJSON("barsic.applyPayload", err.Error())
	}
	return nil
}

type BarsicApplyPayloadBytes struct {
	FieldsMask uint32
	// CommitASAP (TrueType) // Conditional: item.FieldsMask.0
	// Committed (TrueType) // Conditional: item.FieldsMask.1
	Offset  int64
	Payload []byte
}

func (BarsicApplyPayloadBytes) TLName() string { return "barsic.applyPayload" }
func (BarsicApplyPayloadBytes) TLTag() uint32  { return 0x8c981e13 }

func (item *BarsicApplyPayloadBytes) SetCommitASAP(v bool) {
	if v {
		item.FieldsMask |= 1 << 0
	} else {
		item.FieldsMask &^= 1 << 0
	}
}
func (item BarsicApplyPayloadBytes) IsSetCommitASAP() bool { return item.FieldsMask&(1<<0) != 0 }

func (item *BarsicApplyPayloadBytes) SetCommitted(v bool) {
	if v {
		item.FieldsMask |= 1 << 1
	} else {
		item.FieldsMask &^= 1 << 1
	}
}
func (item BarsicApplyPayloadBytes) IsSetCommitted() bool { return item.FieldsMask&(1<<1) != 0 }

func (item *BarsicApplyPayloadBytes) Reset() {
	item.FieldsMask = 0
	item.Offset = 0
	item.Payload = item.Payload[:0]
}

func (item *BarsicApplyPayloadBytes) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.Offset); err != nil {
		return w, err
	}
	return basictl.StringReadBytes(w, &item.Payload)
}

func (item *BarsicApplyPayloadBytes) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	w = basictl.LongWrite(w, item.Offset)
	return basictl.StringWriteBytes(w, item.Payload), nil
}

func (item *BarsicApplyPayloadBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x8c981e13); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *BarsicApplyPayloadBytes) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x8c981e13)
	return item.Write(w)
}

func (item *BarsicApplyPayloadBytes) ReadResult(w []byte, ret *tlTrue.True) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *BarsicApplyPayloadBytes) WriteResult(w []byte, ret tlTrue.True) (_ []byte, err error) {
	return ret.WriteBoxed(w)
}

func (item *BarsicApplyPayloadBytes) ReadResultJSON(legacyTypeNames bool, j interface{}, ret *tlTrue.True) error {
	if err := ret.ReadJSONLegacy(legacyTypeNames, j); err != nil {
		return err
	}
	return nil
}

func (item *BarsicApplyPayloadBytes) WriteResultJSON(w []byte, ret tlTrue.True) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *BarsicApplyPayloadBytes) writeResultJSON(newTypeNames bool, short bool, w []byte, ret tlTrue.True) (_ []byte, err error) {
	if w, err = ret.WriteJSONOpt(newTypeNames, short, w); err != nil {
		return w, err
	}
	return w, nil
}

func (item *BarsicApplyPayloadBytes) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlTrue.True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *BarsicApplyPayloadBytes) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret tlTrue.True
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *BarsicApplyPayloadBytes) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := internal.JsonBytesToInterface(r)
	if err != nil {
		return r, w, internal.ErrorInvalidJSON("barsic.applyPayload", err.Error())
	}
	var ret tlTrue.True
	if err = item.ReadResultJSON(true, j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item BarsicApplyPayloadBytes) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *BarsicApplyPayloadBytes) ReadJSONLegacy(legacyTypeNames bool, j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return internal.ErrorInvalidJSON("barsic.applyPayload", "expected json object")
	}
	_jFieldsMask := _jm["fields_mask"]
	delete(_jm, "fields_mask")
	if err := internal.JsonReadUint32(_jFieldsMask, &item.FieldsMask); err != nil {
		return err
	}
	_jCommitASAP := _jm["commitASAP"]
	delete(_jm, "commitASAP")
	_jCommitted := _jm["committed"]
	delete(_jm, "committed")
	_jOffset := _jm["offset"]
	delete(_jm, "offset")
	if err := internal.JsonReadInt64(_jOffset, &item.Offset); err != nil {
		return err
	}
	_jPayload := _jm["payload"]
	delete(_jm, "payload")
	if err := internal.JsonReadStringBytes(_jPayload, &item.Payload); err != nil {
		return err
	}
	for k := range _jm {
		return internal.ErrorInvalidJSONExcessElement("barsic.applyPayload", k)
	}
	if _jCommitASAP != nil {
		_bit := false
		if err := internal.JsonReadBool(_jCommitASAP, &_bit); err != nil {
			return err
		}
		if _bit {
			item.FieldsMask |= 1 << 0
		} else {
			item.FieldsMask &^= 1 << 0
		}
	}
	if _jCommitted != nil {
		_bit := false
		if err := internal.JsonReadBool(_jCommitted, &_bit); err != nil {
			return err
		}
		if _bit {
			item.FieldsMask |= 1 << 1
		} else {
			item.FieldsMask &^= 1 << 1
		}
	}
	return nil
}

func (item *BarsicApplyPayloadBytes) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *BarsicApplyPayloadBytes) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldsMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"fields_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldsMask)
	}
	if item.FieldsMask&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"commitASAP":true`...)
	}
	if item.FieldsMask&(1<<1) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"committed":true`...)
	}
	if item.Offset != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"offset":`...)
		w = basictl.JSONWriteInt64(w, item.Offset)
	}
	if len(item.Payload) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"payload":`...)
		w = basictl.JSONWriteStringBytes(w, item.Payload)
	}
	return append(w, '}'), nil
}

func (item *BarsicApplyPayloadBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *BarsicApplyPayloadBytes) UnmarshalJSON(b []byte) error {
	j, err := internal.JsonBytesToInterface(b)
	if err != nil {
		return internal.ErrorInvalidJSON("barsic.applyPayload", err.Error())
	}
	if err = item.ReadJSONLegacy(true, j); err != nil {
		return internal.ErrorInvalidJSON("barsic.applyPayload", err.Error())
	}
	return nil
}
