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

func BuiltinVectorStatshouseTopElementRead(w []byte, vec *[]StatshouseTopElement) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]StatshouseTopElement, l)
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

func BuiltinVectorStatshouseTopElementWrite(w []byte, vec []StatshouseTopElement) (_ []byte, err error) {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		if w, err = elem.Write(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinVectorStatshouseTopElementReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]StatshouseTopElement) error {
	*vec = (*vec)[:cap(*vec)]
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return ErrorInvalidJSON("[]StatshouseTopElement", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if len(*vec) <= index {
				var newValue StatshouseTopElement
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
			return ErrorInvalidJSON("[]StatshouseTopElement", "expected json array's end")
		}
	}
	*vec = (*vec)[:index]
	return nil
}

func BuiltinVectorStatshouseTopElementWriteJSON(w []byte, vec []StatshouseTopElement) (_ []byte, err error) {
	return BuiltinVectorStatshouseTopElementWriteJSONOpt(true, false, w, vec)
}
func BuiltinVectorStatshouseTopElementWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec []StatshouseTopElement) (_ []byte, err error) {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		if w, err = elem.WriteJSONOpt(newTypeNames, short, w); err != nil {
			return w, err
		}
	}
	return append(w, ']'), nil
}

func BuiltinVectorStatshouseTopElementBytesRead(w []byte, vec *[]StatshouseTopElementBytes) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]StatshouseTopElementBytes, l)
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

func BuiltinVectorStatshouseTopElementBytesWrite(w []byte, vec []StatshouseTopElementBytes) (_ []byte, err error) {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		if w, err = elem.Write(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func BuiltinVectorStatshouseTopElementBytesReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, vec *[]StatshouseTopElementBytes) error {
	*vec = (*vec)[:cap(*vec)]
	index := 0
	if in != nil {
		in.Delim('[')
		if !in.Ok() {
			return ErrorInvalidJSON("[]StatshouseTopElementBytes", "expected json array")
		}
		for ; !in.IsDelim(']'); index++ {
			if len(*vec) <= index {
				var newValue StatshouseTopElementBytes
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
			return ErrorInvalidJSON("[]StatshouseTopElementBytes", "expected json array's end")
		}
	}
	*vec = (*vec)[:index]
	return nil
}

func BuiltinVectorStatshouseTopElementBytesWriteJSON(w []byte, vec []StatshouseTopElementBytes) (_ []byte, err error) {
	return BuiltinVectorStatshouseTopElementBytesWriteJSONOpt(true, false, w, vec)
}
func BuiltinVectorStatshouseTopElementBytesWriteJSONOpt(newTypeNames bool, short bool, w []byte, vec []StatshouseTopElementBytes) (_ []byte, err error) {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		if w, err = elem.WriteJSONOpt(newTypeNames, short, w); err != nil {
			return w, err
		}
	}
	return append(w, ']'), nil
}

type StatshouseTopElement struct {
	Key        string
	FieldsMask uint32
	Value      StatshouseMultiValue
}

func (StatshouseTopElement) TLName() string { return "statshouse.top_element" }
func (StatshouseTopElement) TLTag() uint32  { return 0x9ffdea42 }

func (item *StatshouseTopElement) Reset() {
	item.Key = ""
	item.FieldsMask = 0
	item.Value.Reset()
}

func (item *StatshouseTopElement) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringRead(w, &item.Key); err != nil {
		return w, err
	}
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	return item.Value.Read(w, item.FieldsMask)
}

func (item *StatshouseTopElement) Write(w []byte) (_ []byte, err error) {
	w = basictl.StringWrite(w, item.Key)
	w = basictl.NatWrite(w, item.FieldsMask)
	return item.Value.Write(w, item.FieldsMask)
}

func (item *StatshouseTopElement) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9ffdea42); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseTopElement) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x9ffdea42)
	return item.Write(w)
}

func (item StatshouseTopElement) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *StatshouseTopElement) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propKeyPresented bool
	var propFieldsMaskPresented bool
	var rawValue []byte

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "key":
				if propKeyPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.top_element", "key")
				}
				if err := Json2ReadString(in, &item.Key); err != nil {
					return err
				}
				propKeyPresented = true
			case "fields_mask":
				if propFieldsMaskPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.top_element", "fields_mask")
				}
				if err := Json2ReadUint32(in, &item.FieldsMask); err != nil {
					return err
				}
				propFieldsMaskPresented = true
			case "value":
				if rawValue != nil {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.top_element", "value")
				}
				rawValue = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			default:
				return ErrorInvalidJSONExcessElement("statshouse.top_element", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propKeyPresented {
		item.Key = ""
	}
	if !propFieldsMaskPresented {
		item.FieldsMask = 0
	}
	var inValuePointer *basictl.JsonLexer
	inValue := basictl.JsonLexer{Data: rawValue}
	if rawValue != nil {
		inValuePointer = &inValue
	}
	if err := item.Value.ReadJSON(legacyTypeNames, inValuePointer, item.FieldsMask); err != nil {
		return err
	}

	return nil
}

func (item *StatshouseTopElement) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *StatshouseTopElement) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexKey := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"key":`...)
	w = basictl.JSONWriteString(w, item.Key)
	if (len(item.Key) != 0) == false {
		w = w[:backupIndexKey]
	}
	backupIndexFieldsMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"fields_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldsMask)
	if (item.FieldsMask != 0) == false {
		w = w[:backupIndexFieldsMask]
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"value":`...)
	if w, err = item.Value.WriteJSONOpt(newTypeNames, short, w, item.FieldsMask); err != nil {
		return w, err
	}
	return append(w, '}'), nil
}

func (item *StatshouseTopElement) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseTopElement) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("statshouse.top_element", err.Error())
	}
	return nil
}

type StatshouseTopElementBytes struct {
	Key        []byte
	FieldsMask uint32
	Value      StatshouseMultiValueBytes
}

func (StatshouseTopElementBytes) TLName() string { return "statshouse.top_element" }
func (StatshouseTopElementBytes) TLTag() uint32  { return 0x9ffdea42 }

func (item *StatshouseTopElementBytes) Reset() {
	item.Key = item.Key[:0]
	item.FieldsMask = 0
	item.Value.Reset()
}

func (item *StatshouseTopElementBytes) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.StringReadBytes(w, &item.Key); err != nil {
		return w, err
	}
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	return item.Value.Read(w, item.FieldsMask)
}

func (item *StatshouseTopElementBytes) Write(w []byte) (_ []byte, err error) {
	w = basictl.StringWriteBytes(w, item.Key)
	w = basictl.NatWrite(w, item.FieldsMask)
	return item.Value.Write(w, item.FieldsMask)
}

func (item *StatshouseTopElementBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9ffdea42); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseTopElementBytes) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x9ffdea42)
	return item.Write(w)
}

func (item StatshouseTopElementBytes) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *StatshouseTopElementBytes) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propKeyPresented bool
	var propFieldsMaskPresented bool
	var rawValue []byte

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "key":
				if propKeyPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.top_element", "key")
				}
				if err := Json2ReadStringBytes(in, &item.Key); err != nil {
					return err
				}
				propKeyPresented = true
			case "fields_mask":
				if propFieldsMaskPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.top_element", "fields_mask")
				}
				if err := Json2ReadUint32(in, &item.FieldsMask); err != nil {
					return err
				}
				propFieldsMaskPresented = true
			case "value":
				if rawValue != nil {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.top_element", "value")
				}
				rawValue = in.Raw()
				if !in.Ok() {
					return in.Error()
				}
			default:
				return ErrorInvalidJSONExcessElement("statshouse.top_element", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propKeyPresented {
		item.Key = item.Key[:0]
	}
	if !propFieldsMaskPresented {
		item.FieldsMask = 0
	}
	var inValuePointer *basictl.JsonLexer
	inValue := basictl.JsonLexer{Data: rawValue}
	if rawValue != nil {
		inValuePointer = &inValue
	}
	if err := item.Value.ReadJSON(legacyTypeNames, inValuePointer, item.FieldsMask); err != nil {
		return err
	}

	return nil
}

func (item *StatshouseTopElementBytes) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *StatshouseTopElementBytes) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexKey := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"key":`...)
	w = basictl.JSONWriteStringBytes(w, item.Key)
	if (len(item.Key) != 0) == false {
		w = w[:backupIndexKey]
	}
	backupIndexFieldsMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"fields_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldsMask)
	if (item.FieldsMask != 0) == false {
		w = w[:backupIndexFieldsMask]
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"value":`...)
	if w, err = item.Value.WriteJSONOpt(newTypeNames, short, w, item.FieldsMask); err != nil {
		return w, err
	}
	return append(w, '}'), nil
}

func (item *StatshouseTopElementBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseTopElementBytes) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("statshouse.top_element", err.Error())
	}
	return nil
}
