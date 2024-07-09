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

type SqliteMetainfo struct {
	FieldMask   uint32
	Offset      int64
	ControlMeta string // Conditional: item.FieldMask.0
}

func (SqliteMetainfo) TLName() string { return "sqlite.metainfo" }
func (SqliteMetainfo) TLTag() uint32  { return 0x9286affa }

func (item *SqliteMetainfo) SetControlMeta(v string) {
	item.ControlMeta = v
	item.FieldMask |= 1 << 0
}
func (item *SqliteMetainfo) ClearControlMeta() {
	item.ControlMeta = ""
	item.FieldMask &^= 1 << 0
}
func (item SqliteMetainfo) IsSetControlMeta() bool { return item.FieldMask&(1<<0) != 0 }

func (item *SqliteMetainfo) Reset() {
	item.FieldMask = 0
	item.Offset = 0
	item.ControlMeta = ""
}

func (item *SqliteMetainfo) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldMask); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.Offset); err != nil {
		return w, err
	}
	if item.FieldMask&(1<<0) != 0 {
		if w, err = basictl.StringRead(w, &item.ControlMeta); err != nil {
			return w, err
		}
	} else {
		item.ControlMeta = ""
	}
	return w, nil
}

// This method is general version of Write, use it instead!
func (item *SqliteMetainfo) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *SqliteMetainfo) Write(w []byte) []byte {
	w = basictl.NatWrite(w, item.FieldMask)
	w = basictl.LongWrite(w, item.Offset)
	if item.FieldMask&(1<<0) != 0 {
		w = basictl.StringWrite(w, item.ControlMeta)
	}
	return w
}

func (item *SqliteMetainfo) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9286affa); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *SqliteMetainfo) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *SqliteMetainfo) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x9286affa)
	return item.Write(w)
}

func (item SqliteMetainfo) String() string {
	return string(item.WriteJSON(nil))
}

func (item *SqliteMetainfo) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFieldMaskPresented bool
	var propOffsetPresented bool
	var propControlMetaPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "field_mask":
				if propFieldMaskPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("sqlite.metainfo", "field_mask")
				}
				if err := Json2ReadUint32(in, &item.FieldMask); err != nil {
					return err
				}
				propFieldMaskPresented = true
			case "offset":
				if propOffsetPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("sqlite.metainfo", "offset")
				}
				if err := Json2ReadInt64(in, &item.Offset); err != nil {
					return err
				}
				propOffsetPresented = true
			case "control_meta":
				if propControlMetaPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("sqlite.metainfo", "control_meta")
				}
				if err := Json2ReadString(in, &item.ControlMeta); err != nil {
					return err
				}
				propControlMetaPresented = true
			default:
				return ErrorInvalidJSONExcessElement("sqlite.metainfo", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propFieldMaskPresented {
		item.FieldMask = 0
	}
	if !propOffsetPresented {
		item.Offset = 0
	}
	if !propControlMetaPresented {
		item.ControlMeta = ""
	}
	if propControlMetaPresented {
		item.FieldMask |= 1 << 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *SqliteMetainfo) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *SqliteMetainfo) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *SqliteMetainfo) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexFieldMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"field_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldMask)
	if (item.FieldMask != 0) == false {
		w = w[:backupIndexFieldMask]
	}
	backupIndexOffset := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"offset":`...)
	w = basictl.JSONWriteInt64(w, item.Offset)
	if (item.Offset != 0) == false {
		w = w[:backupIndexOffset]
	}
	if item.FieldMask&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"control_meta":`...)
		w = basictl.JSONWriteString(w, item.ControlMeta)
	}
	return append(w, '}')
}

func (item *SqliteMetainfo) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *SqliteMetainfo) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("sqlite.metainfo", err.Error())
	}
	return nil
}

type SqliteMetainfoBytes struct {
	FieldMask   uint32
	Offset      int64
	ControlMeta []byte // Conditional: item.FieldMask.0
}

func (SqliteMetainfoBytes) TLName() string { return "sqlite.metainfo" }
func (SqliteMetainfoBytes) TLTag() uint32  { return 0x9286affa }

func (item *SqliteMetainfoBytes) SetControlMeta(v []byte) {
	item.ControlMeta = v
	item.FieldMask |= 1 << 0
}
func (item *SqliteMetainfoBytes) ClearControlMeta() {
	item.ControlMeta = item.ControlMeta[:0]
	item.FieldMask &^= 1 << 0
}
func (item SqliteMetainfoBytes) IsSetControlMeta() bool { return item.FieldMask&(1<<0) != 0 }

func (item *SqliteMetainfoBytes) Reset() {
	item.FieldMask = 0
	item.Offset = 0
	item.ControlMeta = item.ControlMeta[:0]
}

func (item *SqliteMetainfoBytes) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldMask); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.Offset); err != nil {
		return w, err
	}
	if item.FieldMask&(1<<0) != 0 {
		if w, err = basictl.StringReadBytes(w, &item.ControlMeta); err != nil {
			return w, err
		}
	} else {
		item.ControlMeta = item.ControlMeta[:0]
	}
	return w, nil
}

// This method is general version of Write, use it instead!
func (item *SqliteMetainfoBytes) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *SqliteMetainfoBytes) Write(w []byte) []byte {
	w = basictl.NatWrite(w, item.FieldMask)
	w = basictl.LongWrite(w, item.Offset)
	if item.FieldMask&(1<<0) != 0 {
		w = basictl.StringWriteBytes(w, item.ControlMeta)
	}
	return w
}

func (item *SqliteMetainfoBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9286affa); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *SqliteMetainfoBytes) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *SqliteMetainfoBytes) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x9286affa)
	return item.Write(w)
}

func (item SqliteMetainfoBytes) String() string {
	return string(item.WriteJSON(nil))
}

func (item *SqliteMetainfoBytes) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFieldMaskPresented bool
	var propOffsetPresented bool
	var propControlMetaPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "field_mask":
				if propFieldMaskPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("sqlite.metainfo", "field_mask")
				}
				if err := Json2ReadUint32(in, &item.FieldMask); err != nil {
					return err
				}
				propFieldMaskPresented = true
			case "offset":
				if propOffsetPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("sqlite.metainfo", "offset")
				}
				if err := Json2ReadInt64(in, &item.Offset); err != nil {
					return err
				}
				propOffsetPresented = true
			case "control_meta":
				if propControlMetaPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("sqlite.metainfo", "control_meta")
				}
				if err := Json2ReadStringBytes(in, &item.ControlMeta); err != nil {
					return err
				}
				propControlMetaPresented = true
			default:
				return ErrorInvalidJSONExcessElement("sqlite.metainfo", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propFieldMaskPresented {
		item.FieldMask = 0
	}
	if !propOffsetPresented {
		item.Offset = 0
	}
	if !propControlMetaPresented {
		item.ControlMeta = item.ControlMeta[:0]
	}
	if propControlMetaPresented {
		item.FieldMask |= 1 << 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *SqliteMetainfoBytes) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *SqliteMetainfoBytes) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *SqliteMetainfoBytes) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexFieldMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"field_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldMask)
	if (item.FieldMask != 0) == false {
		w = w[:backupIndexFieldMask]
	}
	backupIndexOffset := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"offset":`...)
	w = basictl.JSONWriteInt64(w, item.Offset)
	if (item.Offset != 0) == false {
		w = w[:backupIndexOffset]
	}
	if item.FieldMask&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"control_meta":`...)
		w = basictl.JSONWriteStringBytes(w, item.ControlMeta)
	}
	return append(w, '}')
}

func (item *SqliteMetainfoBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *SqliteMetainfoBytes) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("sqlite.metainfo", err.Error())
	}
	return nil
}
