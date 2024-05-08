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

type MetadataResetFlood struct {
	Metric string
}

func (MetadataResetFlood) TLName() string { return "metadata.resetFlood" }
func (MetadataResetFlood) TLTag() uint32  { return 0x9faf5282 }

func (item *MetadataResetFlood) Reset() {
	item.Metric = ""
}

func (item *MetadataResetFlood) Read(w []byte) (_ []byte, err error) {
	return basictl.StringRead(w, &item.Metric)
}

func (item *MetadataResetFlood) Write(w []byte) (_ []byte, err error) {
	return basictl.StringWrite(w, item.Metric), nil
}

func (item *MetadataResetFlood) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9faf5282); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *MetadataResetFlood) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x9faf5282)
	return item.Write(w)
}

func (item *MetadataResetFlood) ReadResult(w []byte, ret *MetadataResetFloodResponse) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *MetadataResetFlood) WriteResult(w []byte, ret MetadataResetFloodResponse) (_ []byte, err error) {
	return ret.WriteBoxed(w)
}

func (item *MetadataResetFlood) ReadResultJSON(legacyTypeNames bool, in *basictl.JsonLexer, ret *MetadataResetFloodResponse) error {
	if err := ret.ReadJSON(legacyTypeNames, in); err != nil {
		return err
	}
	return nil
}

func (item *MetadataResetFlood) WriteResultJSON(w []byte, ret MetadataResetFloodResponse) (_ []byte, err error) {
	return item.writeResultJSON(true, false, w, ret)
}

func (item *MetadataResetFlood) writeResultJSON(newTypeNames bool, short bool, w []byte, ret MetadataResetFloodResponse) (_ []byte, err error) {
	if w, err = ret.WriteJSONOpt(newTypeNames, short, w); err != nil {
		return w, err
	}
	return w, nil
}

func (item *MetadataResetFlood) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret MetadataResetFloodResponse
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *MetadataResetFlood) ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret MetadataResetFloodResponse
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(newTypeNames, short, w, ret)
	return r, w, err
}

func (item *MetadataResetFlood) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	var ret MetadataResetFloodResponse
	err := item.ReadResultJSON(true, &basictl.JsonLexer{Data: r}, &ret)
	if err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item MetadataResetFlood) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *MetadataResetFlood) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propMetricPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "metric":
				if propMetricPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("metadata.resetFlood", "metric")
				}
				if err := Json2ReadString(in, &item.Metric); err != nil {
					return err
				}
				propMetricPresented = true
			default:
				return ErrorInvalidJSONExcessElement("metadata.resetFlood", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propMetricPresented {
		item.Metric = ""
	}
	return nil
}

func (item *MetadataResetFlood) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *MetadataResetFlood) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexMetric := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"metric":`...)
	w = basictl.JSONWriteString(w, item.Metric)
	if (len(item.Metric) != 0) == false {
		w = w[:backupIndexMetric]
	}
	return append(w, '}'), nil
}

func (item *MetadataResetFlood) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *MetadataResetFlood) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("metadata.resetFlood", err.Error())
	}
	return nil
}
