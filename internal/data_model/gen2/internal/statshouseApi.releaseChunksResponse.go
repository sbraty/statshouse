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

type StatshouseApiReleaseChunksResponse struct {
	FieldsMask         uint32
	ReleasedChunkCount int32
}

func (StatshouseApiReleaseChunksResponse) TLName() string {
	return "statshouseApi.releaseChunksResponse"
}
func (StatshouseApiReleaseChunksResponse) TLTag() uint32 { return 0xd12dc2bd }

func (item *StatshouseApiReleaseChunksResponse) Reset() {
	item.FieldsMask = 0
	item.ReleasedChunkCount = 0
}

func (item *StatshouseApiReleaseChunksResponse) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.ReleasedChunkCount)
}

func (item *StatshouseApiReleaseChunksResponse) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	return basictl.IntWrite(w, item.ReleasedChunkCount), nil
}

func (item *StatshouseApiReleaseChunksResponse) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xd12dc2bd); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseApiReleaseChunksResponse) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xd12dc2bd)
	return item.Write(w)
}

func (item StatshouseApiReleaseChunksResponse) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *StatshouseApiReleaseChunksResponse) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFieldsMaskPresented bool
	var propReleasedChunkCountPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "fields_mask":
				if propFieldsMaskPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouseApi.releaseChunksResponse", "fields_mask")
				}
				if err := Json2ReadUint32(in, &item.FieldsMask); err != nil {
					return err
				}
				propFieldsMaskPresented = true
			case "releasedChunkCount":
				if propReleasedChunkCountPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouseApi.releaseChunksResponse", "releasedChunkCount")
				}
				if err := Json2ReadInt32(in, &item.ReleasedChunkCount); err != nil {
					return err
				}
				propReleasedChunkCountPresented = true
			default:
				return ErrorInvalidJSONExcessElement("statshouseApi.releaseChunksResponse", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propFieldsMaskPresented {
		item.FieldsMask = 0
	}
	if !propReleasedChunkCountPresented {
		item.ReleasedChunkCount = 0
	}
	return nil
}

func (item *StatshouseApiReleaseChunksResponse) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *StatshouseApiReleaseChunksResponse) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexFieldsMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"fields_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldsMask)
	if (item.FieldsMask != 0) == false {
		w = w[:backupIndexFieldsMask]
	}
	backupIndexReleasedChunkCount := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"releasedChunkCount":`...)
	w = basictl.JSONWriteInt32(w, item.ReleasedChunkCount)
	if (item.ReleasedChunkCount != 0) == false {
		w = w[:backupIndexReleasedChunkCount]
	}
	return append(w, '}'), nil
}

func (item *StatshouseApiReleaseChunksResponse) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseApiReleaseChunksResponse) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("statshouseApi.releaseChunksResponse", err.Error())
	}
	return nil
}
