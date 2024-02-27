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

type MetadataResetFlood2 struct {
	FieldMask uint32
	Metric    string
	Value     int32 // Conditional: item.FieldMask.1
}

func (MetadataResetFlood2) TLName() string { return "metadata.resetFlood2" }
func (MetadataResetFlood2) TLTag() uint32  { return 0x88d0fd5e }

func (item *MetadataResetFlood2) SetValue(v int32) {
	item.Value = v
	item.FieldMask |= 1 << 1
}
func (item *MetadataResetFlood2) ClearValue() {
	item.Value = 0
	item.FieldMask &^= 1 << 1
}
func (item MetadataResetFlood2) IsSetValue() bool { return item.FieldMask&(1<<1) != 0 }

func (item *MetadataResetFlood2) Reset() {
	item.FieldMask = 0
	item.Metric = ""
	item.Value = 0
}

func (item *MetadataResetFlood2) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldMask); err != nil {
		return w, err
	}
	if w, err = basictl.StringRead(w, &item.Metric); err != nil {
		return w, err
	}
	if item.FieldMask&(1<<1) != 0 {
		if w, err = basictl.IntRead(w, &item.Value); err != nil {
			return w, err
		}
	} else {
		item.Value = 0
	}
	return w, nil
}

func (item *MetadataResetFlood2) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldMask)
	if w, err = basictl.StringWrite(w, item.Metric); err != nil {
		return w, err
	}
	if item.FieldMask&(1<<1) != 0 {
		w = basictl.IntWrite(w, item.Value)
	}
	return w, nil
}

func (item *MetadataResetFlood2) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x88d0fd5e); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *MetadataResetFlood2) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x88d0fd5e)
	return item.Write(w)
}

func (item *MetadataResetFlood2) ReadResult(w []byte, ret *MetadataResetFloodResponse2) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *MetadataResetFlood2) WriteResult(w []byte, ret MetadataResetFloodResponse2) (_ []byte, err error) {
	return ret.WriteBoxed(w)
}

func (item *MetadataResetFlood2) ReadResultJSON(j interface{}, ret *MetadataResetFloodResponse2) error {
	if err := MetadataResetFloodResponse2__ReadJSON(ret, j); err != nil {
		return err
	}
	return nil
}

func (item *MetadataResetFlood2) WriteResultJSON(w []byte, ret MetadataResetFloodResponse2) (_ []byte, err error) {
	return item.writeResultJSON(false, w, ret)
}

func (item *MetadataResetFlood2) writeResultJSON(short bool, w []byte, ret MetadataResetFloodResponse2) (_ []byte, err error) {
	if w, err = ret.WriteJSONOpt(short, w); err != nil {
		return w, err
	}
	return w, nil
}

func (item *MetadataResetFlood2) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret MetadataResetFloodResponse2
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *MetadataResetFlood2) ReadResultWriteResultJSONShort(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret MetadataResetFloodResponse2
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(true, w, ret)
	return r, w, err
}

func (item *MetadataResetFlood2) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := JsonBytesToInterface(r)
	if err != nil {
		return r, w, ErrorInvalidJSON("metadata.resetFlood2", err.Error())
	}
	var ret MetadataResetFloodResponse2
	if err = item.ReadResultJSON(j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item MetadataResetFlood2) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func MetadataResetFlood2__ReadJSON(item *MetadataResetFlood2, j interface{}) error {
	return item.readJSON(j)
}
func (item *MetadataResetFlood2) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("metadata.resetFlood2", "expected json object")
	}
	_jFieldMask := _jm["field_mask"]
	delete(_jm, "field_mask")
	if err := JsonReadUint32(_jFieldMask, &item.FieldMask); err != nil {
		return err
	}
	_jMetric := _jm["metric"]
	delete(_jm, "metric")
	if err := JsonReadString(_jMetric, &item.Metric); err != nil {
		return err
	}
	_jValue := _jm["value"]
	delete(_jm, "value")
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("metadata.resetFlood2", k)
	}
	if _jValue != nil {
		item.FieldMask |= 1 << 1
	}
	if _jValue != nil {
		if err := JsonReadInt32(_jValue, &item.Value); err != nil {
			return err
		}
	} else {
		item.Value = 0
	}
	return nil
}

func (item *MetadataResetFlood2) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *MetadataResetFlood2) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"field_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldMask)
	}
	if len(item.Metric) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"metric":`...)
		w = basictl.JSONWriteString(w, item.Metric)
	}
	if item.FieldMask&(1<<1) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"value":`...)
		w = basictl.JSONWriteInt32(w, item.Value)
	}
	return append(w, '}'), nil
}

func (item *MetadataResetFlood2) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *MetadataResetFlood2) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("metadata.resetFlood2", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("metadata.resetFlood2", err.Error())
	}
	return nil
}
