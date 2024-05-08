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

type MetadataEditEntityEvent struct {
	FieldsMask uint32
	Metric     MetadataEvent
	OldVersion int64
}

func (MetadataEditEntityEvent) TLName() string { return "metadata.editEntityEvent" }
func (MetadataEditEntityEvent) TLTag() uint32  { return 0x1234b677 }

func (item *MetadataEditEntityEvent) Reset() {
	item.FieldsMask = 0
	item.Metric.Reset()
	item.OldVersion = 0
}

func (item *MetadataEditEntityEvent) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = item.Metric.Read(w); err != nil {
		return w, err
	}
	return basictl.LongRead(w, &item.OldVersion)
}

func (item *MetadataEditEntityEvent) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	if w, err = item.Metric.Write(w); err != nil {
		return w, err
	}
	return basictl.LongWrite(w, item.OldVersion), nil
}

func (item *MetadataEditEntityEvent) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x1234b677); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *MetadataEditEntityEvent) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x1234b677)
	return item.Write(w)
}

func (item MetadataEditEntityEvent) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *MetadataEditEntityEvent) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFieldsMaskPresented bool
	var propMetricPresented bool
	var propOldVersionPresented bool

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
					return ErrorInvalidJSONWithDuplicatingKeys("metadata.editEntityEvent", "fields_mask")
				}
				if err := Json2ReadUint32(in, &item.FieldsMask); err != nil {
					return err
				}
				propFieldsMaskPresented = true
			case "metric":
				if propMetricPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("metadata.editEntityEvent", "metric")
				}
				if err := item.Metric.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propMetricPresented = true
			case "old_version":
				if propOldVersionPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("metadata.editEntityEvent", "old_version")
				}
				if err := Json2ReadInt64(in, &item.OldVersion); err != nil {
					return err
				}
				propOldVersionPresented = true
			default:
				return ErrorInvalidJSONExcessElement("metadata.editEntityEvent", key)
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
	if !propMetricPresented {
		item.Metric.Reset()
	}
	if !propOldVersionPresented {
		item.OldVersion = 0
	}
	return nil
}

func (item *MetadataEditEntityEvent) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *MetadataEditEntityEvent) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexFieldsMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"fields_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldsMask)
	if (item.FieldsMask != 0) == false {
		w = w[:backupIndexFieldsMask]
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"metric":`...)
	if w, err = item.Metric.WriteJSONOpt(newTypeNames, short, w); err != nil {
		return w, err
	}
	backupIndexOldVersion := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"old_version":`...)
	w = basictl.JSONWriteInt64(w, item.OldVersion)
	if (item.OldVersion != 0) == false {
		w = w[:backupIndexOldVersion]
	}
	return append(w, '}'), nil
}

func (item *MetadataEditEntityEvent) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *MetadataEditEntityEvent) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("metadata.editEntityEvent", err.Error())
	}
	return nil
}
