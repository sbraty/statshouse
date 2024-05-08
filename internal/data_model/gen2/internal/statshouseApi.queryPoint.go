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

type StatshouseApiQueryPoint struct {
	FieldsMask uint32
	Version    int32
	TopN       int32
	MetricName string
	TimeFrom   int64
	TimeTo     int64
	Function   StatshouseApiFunction
	GroupBy    []string
	Filter     []StatshouseApiFilter
	TimeShift  []int64
	What       []StatshouseApiFunction // Conditional: item.FieldsMask.1
}

func (StatshouseApiQueryPoint) TLName() string { return "statshouseApi.queryPoint" }
func (StatshouseApiQueryPoint) TLTag() uint32  { return 0xc9951bbb }

func (item *StatshouseApiQueryPoint) SetWhat(v []StatshouseApiFunction) {
	item.What = v
	item.FieldsMask |= 1 << 1
}
func (item *StatshouseApiQueryPoint) ClearWhat() {
	item.What = item.What[:0]
	item.FieldsMask &^= 1 << 1
}
func (item StatshouseApiQueryPoint) IsSetWhat() bool { return item.FieldsMask&(1<<1) != 0 }

func (item *StatshouseApiQueryPoint) Reset() {
	item.FieldsMask = 0
	item.Version = 0
	item.TopN = 0
	item.MetricName = ""
	item.TimeFrom = 0
	item.TimeTo = 0
	item.Function.Reset()
	item.GroupBy = item.GroupBy[:0]
	item.Filter = item.Filter[:0]
	item.TimeShift = item.TimeShift[:0]
	item.What = item.What[:0]
}

func (item *StatshouseApiQueryPoint) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.Version); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.TopN); err != nil {
		return w, err
	}
	if w, err = basictl.StringRead(w, &item.MetricName); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.TimeFrom); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.TimeTo); err != nil {
		return w, err
	}
	if w, err = item.Function.ReadBoxed(w); err != nil {
		return w, err
	}
	if w, err = BuiltinVectorStringRead(w, &item.GroupBy); err != nil {
		return w, err
	}
	if w, err = BuiltinVectorStatshouseApiFilterRead(w, &item.Filter); err != nil {
		return w, err
	}
	if w, err = BuiltinVectorLongRead(w, &item.TimeShift); err != nil {
		return w, err
	}
	if item.FieldsMask&(1<<1) != 0 {
		if w, err = BuiltinVectorStatshouseApiFunctionRead(w, &item.What); err != nil {
			return w, err
		}
	} else {
		item.What = item.What[:0]
	}
	return w, nil
}

func (item *StatshouseApiQueryPoint) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	w = basictl.IntWrite(w, item.Version)
	w = basictl.IntWrite(w, item.TopN)
	w = basictl.StringWrite(w, item.MetricName)
	w = basictl.LongWrite(w, item.TimeFrom)
	w = basictl.LongWrite(w, item.TimeTo)
	if w, err = item.Function.WriteBoxed(w); err != nil {
		return w, err
	}
	if w, err = BuiltinVectorStringWrite(w, item.GroupBy); err != nil {
		return w, err
	}
	if w, err = BuiltinVectorStatshouseApiFilterWrite(w, item.Filter); err != nil {
		return w, err
	}
	if w, err = BuiltinVectorLongWrite(w, item.TimeShift); err != nil {
		return w, err
	}
	if item.FieldsMask&(1<<1) != 0 {
		if w, err = BuiltinVectorStatshouseApiFunctionWrite(w, item.What); err != nil {
			return w, err
		}
	}
	return w, nil
}

func (item *StatshouseApiQueryPoint) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xc9951bbb); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseApiQueryPoint) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xc9951bbb)
	return item.Write(w)
}

func (item StatshouseApiQueryPoint) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *StatshouseApiQueryPoint) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFieldsMaskPresented bool
	var propVersionPresented bool
	var propTopNPresented bool
	var propMetricNamePresented bool
	var propTimeFromPresented bool
	var propTimeToPresented bool
	var propFunctionPresented bool
	var propGroupByPresented bool
	var propFilterPresented bool
	var propTimeShiftPresented bool
	var propWhatPresented bool

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
					return ErrorInvalidJSONWithDuplicatingKeys("statshouseApi.queryPoint", "fields_mask")
				}
				if err := Json2ReadUint32(in, &item.FieldsMask); err != nil {
					return err
				}
				propFieldsMaskPresented = true
			case "version":
				if propVersionPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouseApi.queryPoint", "version")
				}
				if err := Json2ReadInt32(in, &item.Version); err != nil {
					return err
				}
				propVersionPresented = true
			case "top_n":
				if propTopNPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouseApi.queryPoint", "top_n")
				}
				if err := Json2ReadInt32(in, &item.TopN); err != nil {
					return err
				}
				propTopNPresented = true
			case "metric_name":
				if propMetricNamePresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouseApi.queryPoint", "metric_name")
				}
				if err := Json2ReadString(in, &item.MetricName); err != nil {
					return err
				}
				propMetricNamePresented = true
			case "time_from":
				if propTimeFromPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouseApi.queryPoint", "time_from")
				}
				if err := Json2ReadInt64(in, &item.TimeFrom); err != nil {
					return err
				}
				propTimeFromPresented = true
			case "time_to":
				if propTimeToPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouseApi.queryPoint", "time_to")
				}
				if err := Json2ReadInt64(in, &item.TimeTo); err != nil {
					return err
				}
				propTimeToPresented = true
			case "function":
				if propFunctionPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouseApi.queryPoint", "function")
				}
				if err := item.Function.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propFunctionPresented = true
			case "group_by":
				if propGroupByPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouseApi.queryPoint", "group_by")
				}
				if err := BuiltinVectorStringReadJSON(legacyTypeNames, in, &item.GroupBy); err != nil {
					return err
				}
				propGroupByPresented = true
			case "filter":
				if propFilterPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouseApi.queryPoint", "filter")
				}
				if err := BuiltinVectorStatshouseApiFilterReadJSON(legacyTypeNames, in, &item.Filter); err != nil {
					return err
				}
				propFilterPresented = true
			case "time_shift":
				if propTimeShiftPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouseApi.queryPoint", "time_shift")
				}
				if err := BuiltinVectorLongReadJSON(legacyTypeNames, in, &item.TimeShift); err != nil {
					return err
				}
				propTimeShiftPresented = true
			case "what":
				if propWhatPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("statshouseApi.queryPoint", "what")
				}
				if err := BuiltinVectorStatshouseApiFunctionReadJSON(legacyTypeNames, in, &item.What); err != nil {
					return err
				}
				propWhatPresented = true
			default:
				return ErrorInvalidJSONExcessElement("statshouseApi.queryPoint", key)
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
	if !propVersionPresented {
		item.Version = 0
	}
	if !propTopNPresented {
		item.TopN = 0
	}
	if !propMetricNamePresented {
		item.MetricName = ""
	}
	if !propTimeFromPresented {
		item.TimeFrom = 0
	}
	if !propTimeToPresented {
		item.TimeTo = 0
	}
	if !propFunctionPresented {
		item.Function.Reset()
	}
	if !propGroupByPresented {
		item.GroupBy = item.GroupBy[:0]
	}
	if !propFilterPresented {
		item.Filter = item.Filter[:0]
	}
	if !propTimeShiftPresented {
		item.TimeShift = item.TimeShift[:0]
	}
	if !propWhatPresented {
		item.What = item.What[:0]
	}
	if propWhatPresented {
		item.FieldsMask |= 1 << 1
	}
	return nil
}

func (item *StatshouseApiQueryPoint) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *StatshouseApiQueryPoint) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexFieldsMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"fields_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldsMask)
	if (item.FieldsMask != 0) == false {
		w = w[:backupIndexFieldsMask]
	}
	backupIndexVersion := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"version":`...)
	w = basictl.JSONWriteInt32(w, item.Version)
	if (item.Version != 0) == false {
		w = w[:backupIndexVersion]
	}
	backupIndexTopN := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"top_n":`...)
	w = basictl.JSONWriteInt32(w, item.TopN)
	if (item.TopN != 0) == false {
		w = w[:backupIndexTopN]
	}
	backupIndexMetricName := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"metric_name":`...)
	w = basictl.JSONWriteString(w, item.MetricName)
	if (len(item.MetricName) != 0) == false {
		w = w[:backupIndexMetricName]
	}
	backupIndexTimeFrom := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"time_from":`...)
	w = basictl.JSONWriteInt64(w, item.TimeFrom)
	if (item.TimeFrom != 0) == false {
		w = w[:backupIndexTimeFrom]
	}
	backupIndexTimeTo := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"time_to":`...)
	w = basictl.JSONWriteInt64(w, item.TimeTo)
	if (item.TimeTo != 0) == false {
		w = w[:backupIndexTimeTo]
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"function":`...)
	if w, err = item.Function.WriteJSONOpt(newTypeNames, short, w); err != nil {
		return w, err
	}
	backupIndexGroupBy := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"group_by":`...)
	if w, err = BuiltinVectorStringWriteJSONOpt(newTypeNames, short, w, item.GroupBy); err != nil {
		return w, err
	}
	if (len(item.GroupBy) != 0) == false {
		w = w[:backupIndexGroupBy]
	}
	backupIndexFilter := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"filter":`...)
	if w, err = BuiltinVectorStatshouseApiFilterWriteJSONOpt(newTypeNames, short, w, item.Filter); err != nil {
		return w, err
	}
	if (len(item.Filter) != 0) == false {
		w = w[:backupIndexFilter]
	}
	backupIndexTimeShift := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"time_shift":`...)
	if w, err = BuiltinVectorLongWriteJSONOpt(newTypeNames, short, w, item.TimeShift); err != nil {
		return w, err
	}
	if (len(item.TimeShift) != 0) == false {
		w = w[:backupIndexTimeShift]
	}
	if item.FieldsMask&(1<<1) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"what":`...)
		if w, err = BuiltinVectorStatshouseApiFunctionWriteJSONOpt(newTypeNames, short, w, item.What); err != nil {
			return w, err
		}
	}
	return append(w, '}'), nil
}

func (item *StatshouseApiQueryPoint) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseApiQueryPoint) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("statshouseApi.queryPoint", err.Error())
	}
	return nil
}
