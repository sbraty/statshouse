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

type DictionaryEngineMetafilesStatBoxed map[string]EngineMetafilesStat

func (DictionaryEngineMetafilesStatBoxed) TLName() string { return "dictionary" }
func (DictionaryEngineMetafilesStatBoxed) TLTag() uint32  { return 0x1f4c618f }

func (item *DictionaryEngineMetafilesStatBoxed) Reset() {
	ptr := (*map[string]EngineMetafilesStat)(item)
	VectorDictionaryFieldEngineMetafilesStatBoxed0Reset(*ptr)
}

func (item *DictionaryEngineMetafilesStatBoxed) Read(w []byte) (_ []byte, err error) {
	ptr := (*map[string]EngineMetafilesStat)(item)
	return VectorDictionaryFieldEngineMetafilesStatBoxed0Read(w, ptr)
}

func (item *DictionaryEngineMetafilesStatBoxed) Write(w []byte) (_ []byte, err error) {
	ptr := (*map[string]EngineMetafilesStat)(item)
	return VectorDictionaryFieldEngineMetafilesStatBoxed0Write(w, *ptr)
}

func (item *DictionaryEngineMetafilesStatBoxed) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1f4c618f); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *DictionaryEngineMetafilesStatBoxed) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x1f4c618f)
	return item.Write(w)
}

func (item DictionaryEngineMetafilesStatBoxed) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func DictionaryEngineMetafilesStatBoxed__ReadJSON(item *DictionaryEngineMetafilesStatBoxed, j interface{}) error {
	return item.readJSON(j)
}
func (item *DictionaryEngineMetafilesStatBoxed) readJSON(j interface{}) error {
	ptr := (*map[string]EngineMetafilesStat)(item)
	if err := VectorDictionaryFieldEngineMetafilesStatBoxed0ReadJSON(j, ptr); err != nil {
		return err
	}
	return nil
}

func (item *DictionaryEngineMetafilesStatBoxed) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}

func (item *DictionaryEngineMetafilesStatBoxed) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	ptr := (*map[string]EngineMetafilesStat)(item)
	if w, err = VectorDictionaryFieldEngineMetafilesStatBoxed0WriteJSONOpt(short, w, *ptr); err != nil {
		return w, err
	}
	return w, nil
}
func (item *DictionaryEngineMetafilesStatBoxed) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *DictionaryEngineMetafilesStatBoxed) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("dictionary", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("dictionary", err.Error())
	}
	return nil
}

type DictionaryString map[string]string

func (DictionaryString) TLName() string { return "dictionary" }
func (DictionaryString) TLTag() uint32  { return 0x1f4c618f }

func (item *DictionaryString) Reset() {
	ptr := (*map[string]string)(item)
	VectorDictionaryFieldString0Reset(*ptr)
}

func (item *DictionaryString) Read(w []byte) (_ []byte, err error) {
	ptr := (*map[string]string)(item)
	return VectorDictionaryFieldString0Read(w, ptr)
}

func (item *DictionaryString) Write(w []byte) (_ []byte, err error) {
	ptr := (*map[string]string)(item)
	return VectorDictionaryFieldString0Write(w, *ptr)
}

func (item *DictionaryString) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1f4c618f); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *DictionaryString) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x1f4c618f)
	return item.Write(w)
}

func (item DictionaryString) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func DictionaryString__ReadJSON(item *DictionaryString, j interface{}) error { return item.readJSON(j) }
func (item *DictionaryString) readJSON(j interface{}) error {
	ptr := (*map[string]string)(item)
	if err := VectorDictionaryFieldString0ReadJSON(j, ptr); err != nil {
		return err
	}
	return nil
}

func (item *DictionaryString) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}

func (item *DictionaryString) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	ptr := (*map[string]string)(item)
	if w, err = VectorDictionaryFieldString0WriteJSONOpt(short, w, *ptr); err != nil {
		return w, err
	}
	return w, nil
}
func (item *DictionaryString) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *DictionaryString) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("dictionary", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("dictionary", err.Error())
	}
	return nil
}

type DictionaryStringBytes []DictionaryFieldStringBytes

func (DictionaryStringBytes) TLName() string { return "dictionary" }
func (DictionaryStringBytes) TLTag() uint32  { return 0x1f4c618f }

func (item *DictionaryStringBytes) Reset() {
	ptr := (*[]DictionaryFieldStringBytes)(item)
	*ptr = (*ptr)[:0]
}

func (item *DictionaryStringBytes) Read(w []byte) (_ []byte, err error) {
	ptr := (*[]DictionaryFieldStringBytes)(item)
	return VectorDictionaryFieldString0BytesRead(w, ptr)
}

func (item *DictionaryStringBytes) Write(w []byte) (_ []byte, err error) {
	ptr := (*[]DictionaryFieldStringBytes)(item)
	return VectorDictionaryFieldString0BytesWrite(w, *ptr)
}

func (item *DictionaryStringBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1f4c618f); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *DictionaryStringBytes) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x1f4c618f)
	return item.Write(w)
}

func (item DictionaryStringBytes) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func DictionaryStringBytes__ReadJSON(item *DictionaryStringBytes, j interface{}) error {
	return item.readJSON(j)
}
func (item *DictionaryStringBytes) readJSON(j interface{}) error {
	ptr := (*[]DictionaryFieldStringBytes)(item)
	if err := VectorDictionaryFieldString0BytesReadJSON(j, ptr); err != nil {
		return err
	}
	return nil
}

func (item *DictionaryStringBytes) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}

func (item *DictionaryStringBytes) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	ptr := (*[]DictionaryFieldStringBytes)(item)
	if w, err = VectorDictionaryFieldString0BytesWriteJSONOpt(short, w, *ptr); err != nil {
		return w, err
	}
	return w, nil
}
func (item *DictionaryStringBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *DictionaryStringBytes) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("dictionary", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("dictionary", err.Error())
	}
	return nil
}
