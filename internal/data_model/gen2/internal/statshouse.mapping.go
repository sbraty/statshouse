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

func BuiltinVectorStatshouseMappingRead(w []byte, vec *[]StatshouseMapping) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]StatshouseMapping, l)
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

func BuiltinVectorStatshouseMappingWrite(w []byte, vec []StatshouseMapping) []byte {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		w = elem.Write(w)
	}
	return w
}

func BuiltinVectorStatshouseMappingReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]StatshouseMapping) error {
	*vec = (*vec)[:cap(*vec)]
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return ErrorInvalidJSON("[]StatshouseMapping", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if len(*vec) <= index {
				var newValue StatshouseMapping
				*vec = append(*vec, newValue)
				*vec = (*vec)[:cap(*vec)]
			}
			if err := (*vec)[index].ReadJSON(legacyTypeNames, in); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return ErrorInvalidJSON("[]StatshouseMapping", "expected json array's end")
		}
	}
	*vec = (*vec)[:index]
	return nil
}

func BuiltinVectorStatshouseMappingWriteJSON(w []byte, vec []StatshouseMapping) []byte {
	return BuiltinVectorStatshouseMappingWriteJSONOpt(true, false, w, vec)
}
func BuiltinVectorStatshouseMappingWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec []StatshouseMapping) []byte {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = elem.WriteJSONOpt(newTypeNames, short, w)
	}
	return append(w, ']')
}

func BuiltinVectorStatshouseMappingBytesRead(w []byte, vec *[]StatshouseMappingBytes) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]StatshouseMappingBytes, l)
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

func BuiltinVectorStatshouseMappingBytesWrite(w []byte, vec []StatshouseMappingBytes) []byte {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		w = elem.Write(w)
	}
	return w
}

func BuiltinVectorStatshouseMappingBytesReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]StatshouseMappingBytes) error {
	*vec = (*vec)[:cap(*vec)]
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return ErrorInvalidJSON("[]StatshouseMappingBytes", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if len(*vec) <= index {
				var newValue StatshouseMappingBytes
				*vec = append(*vec, newValue)
				*vec = (*vec)[:cap(*vec)]
			}
			if err := (*vec)[index].ReadJSON(legacyTypeNames, in); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return ErrorInvalidJSON("[]StatshouseMappingBytes", "expected json array's end")
		}
	}
	*vec = (*vec)[:index]
	return nil
}

func BuiltinVectorStatshouseMappingBytesWriteJSON(w []byte, vec []StatshouseMappingBytes) []byte {
	return BuiltinVectorStatshouseMappingBytesWriteJSONOpt(true, false, w, vec)
}
func BuiltinVectorStatshouseMappingBytesWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec []StatshouseMappingBytes) []byte {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = elem.WriteJSONOpt(newTypeNames, short, w)
	}
	return append(w, ']')
}

type StatshouseMapping struct {
	Str   string
	Value int32
}

func (StatshouseMapping) TLName() string { return "statshouse.mapping" }
func (StatshouseMapping) TLTag() uint32  { return 0xbf401d4b }

func (item *StatshouseMapping) Reset() {
	item.Str = ""
	item.Value = 0
}

func (item *StatshouseMapping) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringRead(w, &item.Str); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.Value)
}

// This method is general version of Write, use it instead!
func (item *StatshouseMapping) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *StatshouseMapping) Write(w []byte) []byte {
	w = basictl.StringWrite(w, item.Str)
	w = basictl.IntWrite(w, item.Value)
	return w
}

func (item *StatshouseMapping) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xbf401d4b); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *StatshouseMapping) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *StatshouseMapping) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xbf401d4b)
	return item.Write(w)
}

func (item StatshouseMapping) String() string {
	return string(item.WriteJSON(nil))
}

func (item *StatshouseMapping) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propStrPresented bool
	var propValuePresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "str":
				if propStrPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.mapping", "str")
				}
				if err := Json2ReadString(in, &item.Str); err != nil {
					return err
				}
				propStrPresented = true
			case "value":
				if propValuePresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.mapping", "value")
				}
				if err := Json2ReadInt32(in, &item.Value); err != nil {
					return err
				}
				propValuePresented = true
			default:
				return ErrorInvalidJSONExcessElement("statshouse.mapping", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propStrPresented {
		item.Str = ""
	}
	if !propValuePresented {
		item.Value = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *StatshouseMapping) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *StatshouseMapping) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *StatshouseMapping) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexStr := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"str":`...)
	w = basictl.JSONWriteString(w, item.Str)
	if (len(item.Str) != 0) == false {
		w = w[:backupIndexStr]
	}
	backupIndexValue := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"value":`...)
	w = basictl.JSONWriteInt32(w, item.Value)
	if (item.Value != 0) == false {
		w = w[:backupIndexValue]
	}
	return append(w, '}')
}

func (item *StatshouseMapping) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *StatshouseMapping) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("statshouse.mapping", err.Error())
	}
	return nil
}

type StatshouseMappingBytes struct {
	Str   []byte
	Value int32
}

func (StatshouseMappingBytes) TLName() string { return "statshouse.mapping" }
func (StatshouseMappingBytes) TLTag() uint32  { return 0xbf401d4b }

func (item *StatshouseMappingBytes) Reset() {
	item.Str = item.Str[:0]
	item.Value = 0
}

func (item *StatshouseMappingBytes) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringReadBytes(w, &item.Str); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.Value)
}

// This method is general version of Write, use it instead!
func (item *StatshouseMappingBytes) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *StatshouseMappingBytes) Write(w []byte) []byte {
	w = basictl.StringWriteBytes(w, item.Str)
	w = basictl.IntWrite(w, item.Value)
	return w
}

func (item *StatshouseMappingBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xbf401d4b); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *StatshouseMappingBytes) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *StatshouseMappingBytes) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xbf401d4b)
	return item.Write(w)
}

func (item StatshouseMappingBytes) String() string {
	return string(item.WriteJSON(nil))
}

func (item *StatshouseMappingBytes) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propStrPresented bool
	var propValuePresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "str":
				if propStrPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.mapping", "str")
				}
				if err := Json2ReadStringBytes(in, &item.Str); err != nil {
					return err
				}
				propStrPresented = true
			case "value":
				if propValuePresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.mapping", "value")
				}
				if err := Json2ReadInt32(in, &item.Value); err != nil {
					return err
				}
				propValuePresented = true
			default:
				return ErrorInvalidJSONExcessElement("statshouse.mapping", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propStrPresented {
		item.Str = item.Str[:0]
	}
	if !propValuePresented {
		item.Value = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *StatshouseMappingBytes) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *StatshouseMappingBytes) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *StatshouseMappingBytes) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexStr := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"str":`...)
	w = basictl.JSONWriteStringBytes(w, item.Str)
	if (len(item.Str) != 0) == false {
		w = w[:backupIndexStr]
	}
	backupIndexValue := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"value":`...)
	w = basictl.JSONWriteInt32(w, item.Value)
	if (item.Value != 0) == false {
		w = w[:backupIndexValue]
	}
	return append(w, '}')
}

func (item *StatshouseMappingBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *StatshouseMappingBytes) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("statshouse.mapping", err.Error())
	}
	return nil
}
