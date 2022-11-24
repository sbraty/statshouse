// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package internal

import (
	"github.com/vkcom/statshouse/internal/vkgo/basictl"
	"sort"
)

var _ = basictl.NatWrite

type DictionaryFieldEngineMetafilesStatBoxed struct {
	Key   string
	Value EngineMetafilesStat
}

func (DictionaryFieldEngineMetafilesStatBoxed) TLName() string { return "dictionaryField" }
func (DictionaryFieldEngineMetafilesStatBoxed) TLTag() uint32  { return 0x239c1b62 }

func (item *DictionaryFieldEngineMetafilesStatBoxed) Reset() {
	item.Key = ""
	item.Value.Reset()
}

func (item *DictionaryFieldEngineMetafilesStatBoxed) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringRead(w, &item.Key); err != nil {
		return w, err
	}
	return item.Value.ReadBoxed(w)
}

func (item *DictionaryFieldEngineMetafilesStatBoxed) Write(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringWrite(w, item.Key); err != nil {
		return w, err
	}
	return item.Value.WriteBoxed(w)
}

func (item *DictionaryFieldEngineMetafilesStatBoxed) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x239c1b62); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *DictionaryFieldEngineMetafilesStatBoxed) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x239c1b62)
	return item.Write(w)
}

func (item DictionaryFieldEngineMetafilesStatBoxed) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func DictionaryFieldEngineMetafilesStatBoxed__ReadJSON(item *DictionaryFieldEngineMetafilesStatBoxed, j interface{}) error {
	return item.readJSON(j)
}
func (item *DictionaryFieldEngineMetafilesStatBoxed) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("dictionaryField", "expected json object")
	}
	_jKey := _jm["key"]
	delete(_jm, "key")
	if err := JsonReadString(_jKey, &item.Key); err != nil {
		return err
	}
	_jValue := _jm["value"]
	delete(_jm, "value")
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("dictionaryField", k)
	}
	if err := EngineMetafilesStat__ReadJSON(&item.Value, _jValue); err != nil {
		return err
	}
	return nil
}

func (item *DictionaryFieldEngineMetafilesStatBoxed) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if len(item.Key) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"key":`...)
		w = basictl.JSONWriteString(w, item.Key)
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"value":`...)
	if w, err = item.Value.WriteJSON(w); err != nil {
		return w, err
	}
	return append(w, '}'), nil
}

func (item *DictionaryFieldEngineMetafilesStatBoxed) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *DictionaryFieldEngineMetafilesStatBoxed) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("dictionaryField", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("dictionaryField", err.Error())
	}
	return nil
}

type DictionaryFieldString struct {
	Key   string
	Value string
}

func (DictionaryFieldString) TLName() string { return "dictionaryField" }
func (DictionaryFieldString) TLTag() uint32  { return 0x239c1b62 }

func (item *DictionaryFieldString) Reset() {
	item.Key = ""
	item.Value = ""
}

func (item *DictionaryFieldString) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringRead(w, &item.Key); err != nil {
		return w, err
	}
	return basictl.StringRead(w, &item.Value)
}

func (item *DictionaryFieldString) Write(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringWrite(w, item.Key); err != nil {
		return w, err
	}
	return basictl.StringWrite(w, item.Value)
}

func (item *DictionaryFieldString) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x239c1b62); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *DictionaryFieldString) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x239c1b62)
	return item.Write(w)
}

func (item DictionaryFieldString) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func DictionaryFieldString__ReadJSON(item *DictionaryFieldString, j interface{}) error {
	return item.readJSON(j)
}
func (item *DictionaryFieldString) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("dictionaryField", "expected json object")
	}
	_jKey := _jm["key"]
	delete(_jm, "key")
	if err := JsonReadString(_jKey, &item.Key); err != nil {
		return err
	}
	_jValue := _jm["value"]
	delete(_jm, "value")
	if err := JsonReadString(_jValue, &item.Value); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("dictionaryField", k)
	}
	return nil
}

func (item *DictionaryFieldString) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if len(item.Key) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"key":`...)
		w = basictl.JSONWriteString(w, item.Key)
	}
	if len(item.Value) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"value":`...)
		w = basictl.JSONWriteString(w, item.Value)
	}
	return append(w, '}'), nil
}

func (item *DictionaryFieldString) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *DictionaryFieldString) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("dictionaryField", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("dictionaryField", err.Error())
	}
	return nil
}

type DictionaryFieldStringBytes struct {
	Key   []byte
	Value []byte
}

func (DictionaryFieldStringBytes) TLName() string { return "dictionaryField" }
func (DictionaryFieldStringBytes) TLTag() uint32  { return 0x239c1b62 }

func (item *DictionaryFieldStringBytes) Reset() {
	item.Key = item.Key[:0]
	item.Value = item.Value[:0]
}

func (item *DictionaryFieldStringBytes) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringReadBytes(w, &item.Key); err != nil {
		return w, err
	}
	return basictl.StringReadBytes(w, &item.Value)
}

func (item *DictionaryFieldStringBytes) Write(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringWriteBytes(w, item.Key); err != nil {
		return w, err
	}
	return basictl.StringWriteBytes(w, item.Value)
}

func (item *DictionaryFieldStringBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x239c1b62); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *DictionaryFieldStringBytes) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x239c1b62)
	return item.Write(w)
}

func (item DictionaryFieldStringBytes) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func DictionaryFieldStringBytes__ReadJSON(item *DictionaryFieldStringBytes, j interface{}) error {
	return item.readJSON(j)
}
func (item *DictionaryFieldStringBytes) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("dictionaryField", "expected json object")
	}
	_jKey := _jm["key"]
	delete(_jm, "key")
	if err := JsonReadStringBytes(_jKey, &item.Key); err != nil {
		return err
	}
	_jValue := _jm["value"]
	delete(_jm, "value")
	if err := JsonReadStringBytes(_jValue, &item.Value); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("dictionaryField", k)
	}
	return nil
}

func (item *DictionaryFieldStringBytes) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if len(item.Key) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"key":`...)
		w = basictl.JSONWriteStringBytes(w, item.Key)
	}
	if len(item.Value) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"value":`...)
		w = basictl.JSONWriteStringBytes(w, item.Value)
	}
	return append(w, '}'), nil
}

func (item *DictionaryFieldStringBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *DictionaryFieldStringBytes) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("dictionaryField", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("dictionaryField", err.Error())
	}
	return nil
}

func VectorDictionaryFieldEngineMetafilesStatBoxed0Reset(m map[string]EngineMetafilesStat) {
	for k := range m {
		delete(m, k)
	}
}

func VectorDictionaryFieldEngineMetafilesStatBoxed0Read(w []byte, m *map[string]EngineMetafilesStat) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil { // No sanity check required for map
		return w, err
	}
	var data map[string]EngineMetafilesStat
	if *m == nil {
		if l == 0 {
			return w, nil
		}
		data = make(map[string]EngineMetafilesStat, l)
		*m = data
	} else {
		data = *m
		for k := range data {
			delete(data, k)
		}
	}
	for i := 0; i < int(l); i++ {
		var elem DictionaryFieldEngineMetafilesStatBoxed
		if w, err = elem.Read(w); err != nil {
			return w, err
		}
		data[elem.Key] = elem.Value
	}
	return w, nil
}

func VectorDictionaryFieldEngineMetafilesStatBoxed0Write(w []byte, m map[string]EngineMetafilesStat) (_ []byte, err error) {
	w = basictl.NatWrite(w, uint32(len(m)))
	for key, val := range m {
		elem := DictionaryFieldEngineMetafilesStatBoxed{Key: key, Value: val}
		if w, err = elem.Write(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorDictionaryFieldEngineMetafilesStatBoxed0ReadJSON(j interface{}, m *map[string]EngineMetafilesStat) error {
	var _map map[string]interface{}
	var _mapok bool
	if j != nil {
		_map, _mapok = j.(map[string]interface{})
		if !_mapok {
			return ErrorInvalidJSON("map[string]EngineMetafilesStat", "expected json object") // TODO - better name
		}
	}
	l := len(_map)
	var data map[string]EngineMetafilesStat
	if *m == nil {
		if l == 0 {
			return nil
		}
		data = make(map[string]EngineMetafilesStat, l)
		*m = data
	} else {
		data = *m
		for k := range data {
			delete(data, k)
		}
	}
	for key, _jvalue := range _map {
		var value EngineMetafilesStat
		if err := EngineMetafilesStat__ReadJSON(&value, _jvalue); err != nil {
			return err
		}
		data[key] = value
	}
	return nil
}

func VectorDictionaryFieldEngineMetafilesStatBoxed0WriteJSON(w []byte, m map[string]EngineMetafilesStat) (_ []byte, err error) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	w = append(w, '{')
	for _, key := range keys {
		value := m[key]
		w = basictl.JSONAddCommaIfNeeded(w)
		w = basictl.JSONWriteString(w, key) // StringKey
		w = append(w, ':')
		if w, err = value.WriteJSON(w); err != nil {
			return w, err
		}
	}
	return append(w, '}'), nil
}

func VectorDictionaryFieldString0Reset(m map[string]string) {
	for k := range m {
		delete(m, k)
	}
}

func VectorDictionaryFieldString0Read(w []byte, m *map[string]string) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil { // No sanity check required for map
		return w, err
	}
	var data map[string]string
	if *m == nil {
		if l == 0 {
			return w, nil
		}
		data = make(map[string]string, l)
		*m = data
	} else {
		data = *m
		for k := range data {
			delete(data, k)
		}
	}
	for i := 0; i < int(l); i++ {
		var elem DictionaryFieldString
		if w, err = elem.Read(w); err != nil {
			return w, err
		}
		data[elem.Key] = elem.Value
	}
	return w, nil
}

func VectorDictionaryFieldString0Write(w []byte, m map[string]string) (_ []byte, err error) {
	w = basictl.NatWrite(w, uint32(len(m)))
	for key, val := range m {
		elem := DictionaryFieldString{Key: key, Value: val}
		if w, err = elem.Write(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorDictionaryFieldString0ReadJSON(j interface{}, m *map[string]string) error {
	var _map map[string]interface{}
	var _mapok bool
	if j != nil {
		_map, _mapok = j.(map[string]interface{})
		if !_mapok {
			return ErrorInvalidJSON("map[string]string", "expected json object") // TODO - better name
		}
	}
	l := len(_map)
	var data map[string]string
	if *m == nil {
		if l == 0 {
			return nil
		}
		data = make(map[string]string, l)
		*m = data
	} else {
		data = *m
		for k := range data {
			delete(data, k)
		}
	}
	for key, _jvalue := range _map {
		var value string
		if err := JsonReadString(_jvalue, &value); err != nil {
			return err
		}
		data[key] = value
	}
	return nil
}

func VectorDictionaryFieldString0WriteJSON(w []byte, m map[string]string) (_ []byte, err error) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	w = append(w, '{')
	for _, key := range keys {
		value := m[key]
		w = basictl.JSONAddCommaIfNeeded(w)
		w = basictl.JSONWriteString(w, key) // StringKey
		w = append(w, ':')
		w = basictl.JSONWriteString(w, value)
	}
	return append(w, '}'), nil
}

func VectorDictionaryFieldString0BytesRead(w []byte, vec *[]DictionaryFieldStringBytes) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]DictionaryFieldStringBytes, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if w, err = (*vec)[i].Read(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorDictionaryFieldString0BytesWrite(w []byte, vec []DictionaryFieldStringBytes) (_ []byte, err error) {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		if w, err = elem.Write(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorDictionaryFieldString0BytesReadJSON(j interface{}, vec *[]DictionaryFieldStringBytes) error {
	var _map map[string]interface{}
	var _mapok bool
	if j != nil {
		_map, _mapok = j.(map[string]interface{})
		if !_mapok {
			return ErrorInvalidJSON("[]DictionaryFieldStringBytes", "expected json object") // TODO - better name
		}
	}
	l := len(_map)
	if cap(*vec) < l {
		*vec = make([]DictionaryFieldStringBytes, l)
	} else {
		*vec = (*vec)[:l]
	}
	i := 0
	arr := *vec
	for key, _jvalue := range _map {
		arr[i].Key = append(arr[i].Key[:0], key...)
		if err := JsonReadStringBytes(_jvalue, &arr[i].Value); err != nil {
			return err
		}
		i++
	}
	return nil
}

func VectorDictionaryFieldString0BytesWriteJSON(w []byte, vec []DictionaryFieldStringBytes) (_ []byte, err error) {
	w = append(w, '{')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = basictl.JSONWriteStringBytes(w, elem.Key)
		w = append(w, ':')
		w = basictl.JSONWriteStringBytes(w, elem.Value)
	}
	return append(w, '}'), nil
}
