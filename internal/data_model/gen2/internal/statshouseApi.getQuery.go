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

type StatshouseApiGetQuery struct {
	FieldsMask  uint32
	AccessToken string
	Query       StatshouseApiQuery
}

func (StatshouseApiGetQuery) TLName() string { return "statshouseApi.getQuery" }
func (StatshouseApiGetQuery) TLTag() uint32  { return 0xc7349bb }

func (item *StatshouseApiGetQuery) Reset() {
	item.FieldsMask = 0
	item.AccessToken = ""
	item.Query.Reset()
}

func (item *StatshouseApiGetQuery) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = basictl.StringRead(w, &item.AccessToken); err != nil {
		return w, err
	}
	return item.Query.Read(w)
}

func (item *StatshouseApiGetQuery) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	w = basictl.StringWrite(w, item.AccessToken)
	return item.Query.Write(w)
}

func (item *StatshouseApiGetQuery) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xc7349bb); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseApiGetQuery) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xc7349bb)
	return item.Write(w)
}

func (item *StatshouseApiGetQuery) ReadResult(w []byte, ret *StatshouseApiGetQueryResponse) (_ []byte, err error) {
	return ret.ReadBoxed(w, item.FieldsMask)
}

func (item *StatshouseApiGetQuery) WriteResult(w []byte, ret StatshouseApiGetQueryResponse) (_ []byte, err error) {
	return ret.WriteBoxed(w, item.FieldsMask)
}

func (item *StatshouseApiGetQuery) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *StatshouseApiGetQueryResponse) error {
	if err := ret.ReadJSON(legacyTypeNames, in, item.FieldsMask); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseApiGetQuery) WriteResultJSON(w []byte, ret StatshouseApiGetQueryResponse) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *StatshouseApiGetQuery) writeResultJSON(newTypeNames bool, short bool, w []byte, ret StatshouseApiGetQueryResponse) (_ []byte, err error) {
	if w, err = ret.WriteJSONOpt(newTypeNames, short, w, item.FieldsMask); err != nil {
		return w, err
	}
	return w, nil
}

func (item *StatshouseApiGetQuery) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret StatshouseApiGetQueryResponse
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *StatshouseApiGetQuery) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret StatshouseApiGetQueryResponse
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *StatshouseApiGetQuery) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret StatshouseApiGetQueryResponse
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item StatshouseApiGetQuery) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *StatshouseApiGetQuery) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFieldsMaskPresented bool
	var propAccessTokenPresented bool
	var propQueryPresented bool

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
					return ErrorInvalidJSONWithDuplicatingKeys("statshouseApi.getQuery", "fields_mask")
				}
				if err := Json2ReadUint32(in, &item.FieldsMask); err != nil {
					return err
				}
				propFieldsMaskPresented = true
			case "access_token":
				if propAccessTokenPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouseApi.getQuery", "access_token")
				}
				if err := Json2ReadString(in, &item.AccessToken); err != nil {
					return err
				}
				propAccessTokenPresented = true
			case "query":
				if propQueryPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouseApi.getQuery", "query")
				}
				if err := item.Query.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propQueryPresented = true
			default:
				return ErrorInvalidJSONExcessElement("statshouseApi.getQuery", key)
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
	if !propAccessTokenPresented {
		item.AccessToken = ""
	}
	if !propQueryPresented {
		item.Query.Reset()
	}
	return nil
}

func (item *StatshouseApiGetQuery) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *StatshouseApiGetQuery) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexFieldsMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"fields_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldsMask)
	if (item.FieldsMask != 0) == false {
		w = w[:backupIndexFieldsMask]
	}
	backupIndexAccessToken := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"access_token":`...)
	w = basictl.JSONWriteString(w, item.AccessToken)
	if (len(item.AccessToken) != 0) == false {
		w = w[:backupIndexAccessToken]
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"query":`...)
	if w, err = item.Query.WriteJSONOpt(newTypeNames, short, w); err != nil {
		return w, err
	}
	return append(w, '}'), nil
}

func (item *StatshouseApiGetQuery) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseApiGetQuery) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("statshouseApi.getQuery", err.Error())
	}
	return nil
}
