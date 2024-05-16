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

func BuiltinVectorStringRead(w []byte, vec *[]string) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]string, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if w, err = basictl.StringRead(w, &(*vec)[i]); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinVectorStringWrite(w []byte, vec []string) []byte {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		w = basictl.StringWrite(w, elem)
	}
	return w
}

func BuiltinVectorStringReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]string) error {
	*vec = (*vec)[:cap(*vec)]
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return ErrorInvalidJSON("[]string", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if len(*vec) <= index {
				var newValue string
				*vec = append(*vec, newValue)
				*vec = (*vec)[:cap(*vec)]
			}
			if err := Json2ReadString(in, &(*vec)[index]); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return ErrorInvalidJSON("[]string", "expected json array's end")
		}
	}
	*vec = (*vec)[:index]
	return nil
}

func BuiltinVectorStringWriteJSON(w []byte, vec []string) []byte {
	return BuiltinVectorStringWriteJSONOpt(true, false, w, vec)
}
func BuiltinVectorStringWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec []string) []byte {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = basictl.JSONWriteString(w, elem)
	}
	return append(w, ']')
}

func BuiltinVectorStringBytesRead(w []byte, vec *[][]byte) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([][]byte, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if w, err = basictl.StringReadBytes(w, &(*vec)[i]); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinVectorStringBytesWrite(w []byte, vec [][]byte) []byte {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		w = basictl.StringWriteBytes(w, elem)
	}
	return w
}

func BuiltinVectorStringBytesReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[][]byte) error {
	*vec = (*vec)[:cap(*vec)]
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return ErrorInvalidJSON("[][]byte", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if len(*vec) <= index {
				var newValue []byte
				*vec = append(*vec, newValue)
				*vec = (*vec)[:cap(*vec)]
			}
			if err := Json2ReadStringBytes(in, &(*vec)[index]); err != nil {
				return err
			}
			in.WantComma()
		}
		in.Delim(']')
		if !in.Ok() {
			return ErrorInvalidJSON("[][]byte", "expected json array's end")
		}
	}
	*vec = (*vec)[:index]
	return nil
}

func BuiltinVectorStringBytesWriteJSON(w []byte, vec [][]byte) []byte {
	return BuiltinVectorStringBytesWriteJSONOpt(true, false, w, vec)
}
func BuiltinVectorStringBytesWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec [][]byte) []byte {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = basictl.JSONWriteStringBytes(w, elem)
	}
	return append(w, ']')
}

type String string

func (String) TLName() string { return "string" }
func (String) TLTag() uint32  { return 0xb5286e24 }

func (item *String) Reset() {
	ptr := (*string)(item)
	*ptr = ""
}

func (item *String) Read(w []byte) (_ []byte, err error) {
	ptr := (*string)(item)
	return basictl.StringRead(w, ptr)
}

// This method is general version of Write, use it instead!
func (item *String) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *String) Write(w []byte) []byte {
	ptr := (*string)(item)
	return basictl.StringWrite(w, *ptr)
}

func (item *String) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xb5286e24); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *String) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *String) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xb5286e24)
	return item.Write(w)
}

func (item String) String() string {
	return string(item.WriteJSON(nil))
}

func (item *String) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*string)(item)
	if err := Json2ReadString(in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *String) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}

func (item *String) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}

func (item *String) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	ptr := (*string)(item)
	w = basictl.JSONWriteString(w, *ptr)
	return w
}
func (item *String) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *String) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("string", err.Error())
	}
	return nil
}

type StringBytes []byte

func (StringBytes) TLName() string { return "string" }
func (StringBytes) TLTag() uint32  { return 0xb5286e24 }

func (item *StringBytes) Reset() {
	ptr := (*[]byte)(item)
	*ptr = (*ptr)[:0]
}

func (item *StringBytes) Read(w []byte) (_ []byte, err error) {
	ptr := (*[]byte)(item)
	return basictl.StringReadBytes(w, ptr)
}

// This method is general version of Write, use it instead!
func (item *StringBytes) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *StringBytes) Write(w []byte) []byte {
	ptr := (*[]byte)(item)
	return basictl.StringWriteBytes(w, *ptr)
}

func (item *StringBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xb5286e24); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *StringBytes) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *StringBytes) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xb5286e24)
	return item.Write(w)
}

func (item StringBytes) String() string {
	return string(item.WriteJSON(nil))
}

func (item *StringBytes) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	ptr := (*[]byte)(item)
	if err := Json2ReadStringBytes(in, ptr); err != nil {
		return err
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *StringBytes) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}

func (item *StringBytes) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}

func (item *StringBytes) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	ptr := (*[]byte)(item)
	w = basictl.JSONWriteStringBytes(w, *ptr)
	return w
}
func (item *StringBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *StringBytes) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("string", err.Error())
	}
	return nil
}
