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

type StatshouseGetConfigResult struct {
	Addresses         []string
	MaxAddressesCount int32
	PreviousAddresses int32
}

func (StatshouseGetConfigResult) TLName() string { return "statshouse.getConfigResult" }
func (StatshouseGetConfigResult) TLTag() uint32  { return 0xc803d07 }

func (item *StatshouseGetConfigResult) Reset() {
	item.Addresses = item.Addresses[:0]
	item.MaxAddressesCount = 0
	item.PreviousAddresses = 0
}

func (item *StatshouseGetConfigResult) Read(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	if w, err = BuiltinVectorStringRead(w, &item.Addresses); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.MaxAddressesCount); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.PreviousAddresses)
}

func (item *StatshouseGetConfigResult) Write(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	if w, err = BuiltinVectorStringWrite(w, item.Addresses); err != nil {
		return w, err
	}
	w = basictl.IntWrite(w, item.MaxAddressesCount)
	return basictl.IntWrite(w, item.PreviousAddresses), nil
}

func (item *StatshouseGetConfigResult) ReadBoxed(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xc803d07); err != nil {
		return w, err
	}
	return item.Read(w, nat_fields_mask)
}

func (item *StatshouseGetConfigResult) WriteBoxed(w []byte, nat_fields_mask uint32) ([]byte, error) {
	w = basictl.NatWrite(w, 0xc803d07)
	return item.Write(w, nat_fields_mask)
}

func (item *StatshouseGetConfigResult) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_fields_mask uint32) error {
	var propAddressesPresented bool
	var propMaxAddressesCountPresented bool
	var propPreviousAddressesPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "addresses":
				if propAddressesPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.getConfigResult", "addresses")
				}
				if err := BuiltinVectorStringReadJSON(legacyTypeNames, in, &item.Addresses); err != nil {
					return err
				}
				propAddressesPresented = true
			case "max_addresses_count":
				if propMaxAddressesCountPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.getConfigResult", "max_addresses_count")
				}
				if err := Json2ReadInt32(in, &item.MaxAddressesCount); err != nil {
					return err
				}
				propMaxAddressesCountPresented = true
			case "previous_addresses":
				if propPreviousAddressesPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.getConfigResult", "previous_addresses")
				}
				if err := Json2ReadInt32(in, &item.PreviousAddresses); err != nil {
					return err
				}
				propPreviousAddressesPresented = true
			default:
				return ErrorInvalidJSONExcessElement("statshouse.getConfigResult", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propAddressesPresented {
		item.Addresses = item.Addresses[:0]
	}
	if !propMaxAddressesCountPresented {
		item.MaxAddressesCount = 0
	}
	if !propPreviousAddressesPresented {
		item.PreviousAddresses = 0
	}
	return nil
}

func (item *StatshouseGetConfigResult) WriteJSON(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_fields_mask)
}
func (item *StatshouseGetConfigResult) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexAddresses := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"addresses":`...)
	if w, err = BuiltinVectorStringWriteJSONOpt(newTypeNames, short, w, item.Addresses); err != nil {
		return w, err
	}
	if (len(item.Addresses) != 0) == false {
		w = w[:backupIndexAddresses]
	}
	backupIndexMaxAddressesCount := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"max_addresses_count":`...)
	w = basictl.JSONWriteInt32(w, item.MaxAddressesCount)
	if (item.MaxAddressesCount != 0) == false {
		w = w[:backupIndexMaxAddressesCount]
	}
	backupIndexPreviousAddresses := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"previous_addresses":`...)
	w = basictl.JSONWriteInt32(w, item.PreviousAddresses)
	if (item.PreviousAddresses != 0) == false {
		w = w[:backupIndexPreviousAddresses]
	}
	return append(w, '}'), nil
}

type StatshouseGetConfigResultBytes struct {
	Addresses         [][]byte
	MaxAddressesCount int32
	PreviousAddresses int32
}

func (StatshouseGetConfigResultBytes) TLName() string { return "statshouse.getConfigResult" }
func (StatshouseGetConfigResultBytes) TLTag() uint32  { return 0xc803d07 }

func (item *StatshouseGetConfigResultBytes) Reset() {
	item.Addresses = item.Addresses[:0]
	item.MaxAddressesCount = 0
	item.PreviousAddresses = 0
}

func (item *StatshouseGetConfigResultBytes) Read(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	if w, err = BuiltinVectorStringBytesRead(w, &item.Addresses); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.MaxAddressesCount); err != nil {
		return w, err
	}
	return basictl.IntRead(w, &item.PreviousAddresses)
}

func (item *StatshouseGetConfigResultBytes) Write(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	if w, err = BuiltinVectorStringBytesWrite(w, item.Addresses); err != nil {
		return w, err
	}
	w = basictl.IntWrite(w, item.MaxAddressesCount)
	return basictl.IntWrite(w, item.PreviousAddresses), nil
}

func (item *StatshouseGetConfigResultBytes) ReadBoxed(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xc803d07); err != nil {
		return w, err
	}
	return item.Read(w, nat_fields_mask)
}

func (item *StatshouseGetConfigResultBytes) WriteBoxed(w []byte, nat_fields_mask uint32) ([]byte, error) {
	w = basictl.NatWrite(w, 0xc803d07)
	return item.Write(w, nat_fields_mask)
}

func (item *StatshouseGetConfigResultBytes) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer, nat_fields_mask uint32) error {
	var propAddressesPresented bool
	var propMaxAddressesCountPresented bool
	var propPreviousAddressesPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "addresses":
				if propAddressesPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.getConfigResult", "addresses")
				}
				if err := BuiltinVectorStringBytesReadJSON(legacyTypeNames, in, &item.Addresses); err != nil {
					return err
				}
				propAddressesPresented = true
			case "max_addresses_count":
				if propMaxAddressesCountPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.getConfigResult", "max_addresses_count")
				}
				if err := Json2ReadInt32(in, &item.MaxAddressesCount); err != nil {
					return err
				}
				propMaxAddressesCountPresented = true
			case "previous_addresses":
				if propPreviousAddressesPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouse.getConfigResult", "previous_addresses")
				}
				if err := Json2ReadInt32(in, &item.PreviousAddresses); err != nil {
					return err
				}
				propPreviousAddressesPresented = true
			default:
				return ErrorInvalidJSONExcessElement("statshouse.getConfigResult", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propAddressesPresented {
		item.Addresses = item.Addresses[:0]
	}
	if !propMaxAddressesCountPresented {
		item.MaxAddressesCount = 0
	}
	if !propPreviousAddressesPresented {
		item.PreviousAddresses = 0
	}
	return nil
}

func (item *StatshouseGetConfigResultBytes) WriteJSON(w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w, nat_fields_mask)
}
func (item *StatshouseGetConfigResultBytes) WriteJSONOpt(newTypeNames bool, short bool, w []byte, nat_fields_mask uint32) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexAddresses := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"addresses":`...)
	if w, err = BuiltinVectorStringBytesWriteJSONOpt(newTypeNames, short, w, item.Addresses); err != nil {
		return w, err
	}
	if (len(item.Addresses) != 0) == false {
		w = w[:backupIndexAddresses]
	}
	backupIndexMaxAddressesCount := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"max_addresses_count":`...)
	w = basictl.JSONWriteInt32(w, item.MaxAddressesCount)
	if (item.MaxAddressesCount != 0) == false {
		w = w[:backupIndexMaxAddressesCount]
	}
	backupIndexPreviousAddresses := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"previous_addresses":`...)
	w = basictl.JSONWriteInt32(w, item.PreviousAddresses)
	if (item.PreviousAddresses != 0) == false {
		w = w[:backupIndexPreviousAddresses]
	}
	return append(w, '}'), nil
}
