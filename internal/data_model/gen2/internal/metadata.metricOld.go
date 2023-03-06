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

type MetadataMetricOld struct {
	Id         int64
	Name       string
	EventType  int32  // Conditional: nat_field_mask.2
	Unused     uint32 // Conditional: nat_field_mask.3
	Version    int64
	UpdateTime uint32
	Data       string
}

func (MetadataMetricOld) TLName() string { return "metadata.metricOld" }
func (MetadataMetricOld) TLTag() uint32  { return 0x9286abfa }

func (item *MetadataMetricOld) SetEventType(v int32, nat_field_mask *uint32) {
	item.EventType = v
	if nat_field_mask != nil {
		*nat_field_mask |= 1 << 2
	}
}
func (item *MetadataMetricOld) ClearEventType(nat_field_mask *uint32) {
	item.EventType = 0
	if nat_field_mask != nil {
		*nat_field_mask &^= 1 << 2
	}
}
func (item MetadataMetricOld) IsSetEventType(nat_field_mask uint32) bool {
	return nat_field_mask&(1<<2) != 0
}

func (item *MetadataMetricOld) SetUnused(v uint32, nat_field_mask *uint32) {
	item.Unused = v
	if nat_field_mask != nil {
		*nat_field_mask |= 1 << 3
	}
}
func (item *MetadataMetricOld) ClearUnused(nat_field_mask *uint32) {
	item.Unused = 0
	if nat_field_mask != nil {
		*nat_field_mask &^= 1 << 3
	}
}
func (item MetadataMetricOld) IsSetUnused(nat_field_mask uint32) bool {
	return nat_field_mask&(1<<3) != 0
}

func (item *MetadataMetricOld) Reset() {
	item.Id = 0
	item.Name = ""
	item.EventType = 0
	item.Unused = 0
	item.Version = 0
	item.UpdateTime = 0
	item.Data = ""
}

func (item *MetadataMetricOld) Read(w []byte, nat_field_mask uint32) (_ []byte, err error) {
	if w, err = basictl.LongRead(w, &item.Id); err != nil {
		return w, err
	}
	if w, err = basictl.StringRead(w, &item.Name); err != nil {
		return w, err
	}
	if nat_field_mask&(1<<2) != 0 {
		if w, err = basictl.IntRead(w, &item.EventType); err != nil {
			return w, err
		}
	} else {
		item.EventType = 0
	}
	if nat_field_mask&(1<<3) != 0 {
		if w, err = basictl.NatRead(w, &item.Unused); err != nil {
			return w, err
		}
	} else {
		item.Unused = 0
	}
	if w, err = basictl.LongRead(w, &item.Version); err != nil {
		return w, err
	}
	if w, err = basictl.NatRead(w, &item.UpdateTime); err != nil {
		return w, err
	}
	return basictl.StringRead(w, &item.Data)
}

func (item *MetadataMetricOld) Write(w []byte, nat_field_mask uint32) (_ []byte, err error) {
	w = basictl.LongWrite(w, item.Id)
	if w, err = basictl.StringWrite(w, item.Name); err != nil {
		return w, err
	}
	if nat_field_mask&(1<<2) != 0 {
		w = basictl.IntWrite(w, item.EventType)
	}
	if nat_field_mask&(1<<3) != 0 {
		w = basictl.NatWrite(w, item.Unused)
	}
	w = basictl.LongWrite(w, item.Version)
	w = basictl.NatWrite(w, item.UpdateTime)
	return basictl.StringWrite(w, item.Data)
}

func (item *MetadataMetricOld) ReadBoxed(w []byte, nat_field_mask uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9286abfa); err != nil {
		return w, err
	}
	return item.Read(w, nat_field_mask)
}

func (item *MetadataMetricOld) WriteBoxed(w []byte, nat_field_mask uint32) ([]byte, error) {
	w = basictl.NatWrite(w, 0x9286abfa)
	return item.Write(w, nat_field_mask)
}

func MetadataMetricOld__ReadJSON(item *MetadataMetricOld, j interface{}, nat_field_mask uint32) error {
	return item.readJSON(j, nat_field_mask)
}
func (item *MetadataMetricOld) readJSON(j interface{}, nat_field_mask uint32) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("metadata.metricOld", "expected json object")
	}
	_jId := _jm["id"]
	delete(_jm, "id")
	if err := JsonReadInt64(_jId, &item.Id); err != nil {
		return err
	}
	_jName := _jm["name"]
	delete(_jm, "name")
	if err := JsonReadString(_jName, &item.Name); err != nil {
		return err
	}
	_jEventType := _jm["event_type"]
	delete(_jm, "event_type")
	_jUnused := _jm["unused"]
	delete(_jm, "unused")
	_jVersion := _jm["version"]
	delete(_jm, "version")
	if err := JsonReadInt64(_jVersion, &item.Version); err != nil {
		return err
	}
	_jUpdateTime := _jm["update_time"]
	delete(_jm, "update_time")
	if err := JsonReadUint32(_jUpdateTime, &item.UpdateTime); err != nil {
		return err
	}
	_jData := _jm["data"]
	delete(_jm, "data")
	if err := JsonReadString(_jData, &item.Data); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("metadata.metricOld", k)
	}
	if nat_field_mask&(1<<2) == 0 && _jEventType != nil {
		return ErrorInvalidJSON("metadata.metricOld", "field 'event_type' is defined, while corresponding implicit fieldmask bit is 0")
	}
	if nat_field_mask&(1<<3) == 0 && _jUnused != nil {
		return ErrorInvalidJSON("metadata.metricOld", "field 'unused' is defined, while corresponding implicit fieldmask bit is 0")
	}
	if nat_field_mask&(1<<2) != 0 {
		if err := JsonReadInt32(_jEventType, &item.EventType); err != nil {
			return err
		}
	} else {
		item.EventType = 0
	}
	if nat_field_mask&(1<<3) != 0 {
		if err := JsonReadUint32(_jUnused, &item.Unused); err != nil {
			return err
		}
	} else {
		item.Unused = 0
	}
	return nil
}

func (item *MetadataMetricOld) WriteJSON(w []byte, nat_field_mask uint32) (_ []byte, err error) {
	w = append(w, '{')
	if item.Id != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"id":`...)
		w = basictl.JSONWriteInt64(w, item.Id)
	}
	if len(item.Name) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"name":`...)
		w = basictl.JSONWriteString(w, item.Name)
	}
	if nat_field_mask&(1<<2) != 0 {
		if item.EventType != 0 {
			w = basictl.JSONAddCommaIfNeeded(w)
			w = append(w, `"event_type":`...)
			w = basictl.JSONWriteInt32(w, item.EventType)
		}
	}
	if nat_field_mask&(1<<3) != 0 {
		if item.Unused != 0 {
			w = basictl.JSONAddCommaIfNeeded(w)
			w = append(w, `"unused":`...)
			w = basictl.JSONWriteUint32(w, item.Unused)
		}
	}
	if item.Version != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"version":`...)
		w = basictl.JSONWriteInt64(w, item.Version)
	}
	if item.UpdateTime != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"update_time":`...)
		w = basictl.JSONWriteUint32(w, item.UpdateTime)
	}
	if len(item.Data) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"data":`...)
		w = basictl.JSONWriteString(w, item.Data)
	}
	return append(w, '}'), nil
}

func VectorMetadataMetricOld0Read(w []byte, vec *[]MetadataMetricOld, nat_t uint32) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]MetadataMetricOld, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if w, err = (*vec)[i].Read(w, nat_t); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorMetadataMetricOld0Write(w []byte, vec []MetadataMetricOld, nat_t uint32) (_ []byte, err error) {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		if w, err = elem.Write(w, nat_t); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorMetadataMetricOld0ReadJSON(j interface{}, vec *[]MetadataMetricOld, nat_t uint32) error {
	l, _arr, err := JsonReadArray("[]MetadataMetricOld", j)
	if err != nil {
		return err
	}
	if cap(*vec) < l {
		*vec = make([]MetadataMetricOld, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if err := MetadataMetricOld__ReadJSON(&(*vec)[i], _arr[i], nat_t); err != nil {
			return err
		}
	}
	return nil
}

func VectorMetadataMetricOld0WriteJSON(w []byte, vec []MetadataMetricOld, nat_t uint32) (_ []byte, err error) {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		if w, err = elem.WriteJSON(w, nat_t); err != nil {
			return w, err
		}
	}
	return append(w, ']'), nil
}
