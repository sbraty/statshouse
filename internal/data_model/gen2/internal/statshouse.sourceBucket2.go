// Copyright 2022 V Kontakte LLC
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

type StatshouseSourceBucket2 struct {
	Metrics            []StatshouseMultiItem
	SampleFactors      []StatshouseSampleFactor
	IngestionStatusOk  []StatshouseSampleFactor
	MissedSeconds      uint32
	LegacyAgentEnv     int32
	IngestionStatusOk2 []StatshouseIngestionStatus2
}

func (StatshouseSourceBucket2) TLName() string { return "statshouse.sourceBucket2" }
func (StatshouseSourceBucket2) TLTag() uint32  { return 0x3af6e822 }

func (item *StatshouseSourceBucket2) Reset() {
	item.Metrics = item.Metrics[:0]
	item.SampleFactors = item.SampleFactors[:0]
	item.IngestionStatusOk = item.IngestionStatusOk[:0]
	item.MissedSeconds = 0
	item.LegacyAgentEnv = 0
	item.IngestionStatusOk2 = item.IngestionStatusOk2[:0]
}

func (item *StatshouseSourceBucket2) Read(w []byte) (_ []byte, err error) {
	if w, err = VectorStatshouseMultiItem0Read(w, &item.Metrics); err != nil {
		return w, err
	}
	if w, err = VectorStatshouseSampleFactor0Read(w, &item.SampleFactors); err != nil {
		return w, err
	}
	if w, err = VectorStatshouseSampleFactor0Read(w, &item.IngestionStatusOk); err != nil {
		return w, err
	}
	if w, err = basictl.NatRead(w, &item.MissedSeconds); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.LegacyAgentEnv); err != nil {
		return w, err
	}
	return VectorStatshouseIngestionStatus20Read(w, &item.IngestionStatusOk2)
}

func (item *StatshouseSourceBucket2) Write(w []byte) (_ []byte, err error) {
	if w, err = VectorStatshouseMultiItem0Write(w, item.Metrics); err != nil {
		return w, err
	}
	if w, err = VectorStatshouseSampleFactor0Write(w, item.SampleFactors); err != nil {
		return w, err
	}
	if w, err = VectorStatshouseSampleFactor0Write(w, item.IngestionStatusOk); err != nil {
		return w, err
	}
	w = basictl.NatWrite(w, item.MissedSeconds)
	w = basictl.IntWrite(w, item.LegacyAgentEnv)
	return VectorStatshouseIngestionStatus20Write(w, item.IngestionStatusOk2)
}

func (item *StatshouseSourceBucket2) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x3af6e822); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseSourceBucket2) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x3af6e822)
	return item.Write(w)
}

func (item StatshouseSourceBucket2) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func StatshouseSourceBucket2__ReadJSON(item *StatshouseSourceBucket2, j interface{}) error {
	return item.readJSON(j)
}
func (item *StatshouseSourceBucket2) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouse.sourceBucket2", "expected json object")
	}
	_jMetrics := _jm["metrics"]
	delete(_jm, "metrics")
	_jSampleFactors := _jm["sample_factors"]
	delete(_jm, "sample_factors")
	_jIngestionStatusOk := _jm["ingestion_status_ok"]
	delete(_jm, "ingestion_status_ok")
	_jMissedSeconds := _jm["missed_seconds"]
	delete(_jm, "missed_seconds")
	if err := JsonReadUint32(_jMissedSeconds, &item.MissedSeconds); err != nil {
		return err
	}
	_jLegacyAgentEnv := _jm["legacy_agent_env"]
	delete(_jm, "legacy_agent_env")
	if err := JsonReadInt32(_jLegacyAgentEnv, &item.LegacyAgentEnv); err != nil {
		return err
	}
	_jIngestionStatusOk2 := _jm["ingestion_status_ok2"]
	delete(_jm, "ingestion_status_ok2")
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouse.sourceBucket2", k)
	}
	if err := VectorStatshouseMultiItem0ReadJSON(_jMetrics, &item.Metrics); err != nil {
		return err
	}
	if err := VectorStatshouseSampleFactor0ReadJSON(_jSampleFactors, &item.SampleFactors); err != nil {
		return err
	}
	if err := VectorStatshouseSampleFactor0ReadJSON(_jIngestionStatusOk, &item.IngestionStatusOk); err != nil {
		return err
	}
	if err := VectorStatshouseIngestionStatus20ReadJSON(_jIngestionStatusOk2, &item.IngestionStatusOk2); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseSourceBucket2) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if len(item.Metrics) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"metrics":`...)
		if w, err = VectorStatshouseMultiItem0WriteJSON(w, item.Metrics); err != nil {
			return w, err
		}
	}
	if len(item.SampleFactors) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"sample_factors":`...)
		if w, err = VectorStatshouseSampleFactor0WriteJSON(w, item.SampleFactors); err != nil {
			return w, err
		}
	}
	if len(item.IngestionStatusOk) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"ingestion_status_ok":`...)
		if w, err = VectorStatshouseSampleFactor0WriteJSON(w, item.IngestionStatusOk); err != nil {
			return w, err
		}
	}
	if item.MissedSeconds != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"missed_seconds":`...)
		w = basictl.JSONWriteUint32(w, item.MissedSeconds)
	}
	if item.LegacyAgentEnv != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"legacy_agent_env":`...)
		w = basictl.JSONWriteInt32(w, item.LegacyAgentEnv)
	}
	if len(item.IngestionStatusOk2) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"ingestion_status_ok2":`...)
		if w, err = VectorStatshouseIngestionStatus20WriteJSON(w, item.IngestionStatusOk2); err != nil {
			return w, err
		}
	}
	return append(w, '}'), nil
}

func (item *StatshouseSourceBucket2) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseSourceBucket2) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("statshouse.sourceBucket2", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("statshouse.sourceBucket2", err.Error())
	}
	return nil
}

type StatshouseSourceBucket2Bytes struct {
	Metrics            []StatshouseMultiItemBytes
	SampleFactors      []StatshouseSampleFactor
	IngestionStatusOk  []StatshouseSampleFactor
	MissedSeconds      uint32
	LegacyAgentEnv     int32
	IngestionStatusOk2 []StatshouseIngestionStatus2
}

func (StatshouseSourceBucket2Bytes) TLName() string { return "statshouse.sourceBucket2" }
func (StatshouseSourceBucket2Bytes) TLTag() uint32  { return 0x3af6e822 }

func (item *StatshouseSourceBucket2Bytes) Reset() {
	item.Metrics = item.Metrics[:0]
	item.SampleFactors = item.SampleFactors[:0]
	item.IngestionStatusOk = item.IngestionStatusOk[:0]
	item.MissedSeconds = 0
	item.LegacyAgentEnv = 0
	item.IngestionStatusOk2 = item.IngestionStatusOk2[:0]
}

func (item *StatshouseSourceBucket2Bytes) Read(w []byte) (_ []byte, err error) {
	if w, err = VectorStatshouseMultiItem0BytesRead(w, &item.Metrics); err != nil {
		return w, err
	}
	if w, err = VectorStatshouseSampleFactor0Read(w, &item.SampleFactors); err != nil {
		return w, err
	}
	if w, err = VectorStatshouseSampleFactor0Read(w, &item.IngestionStatusOk); err != nil {
		return w, err
	}
	if w, err = basictl.NatRead(w, &item.MissedSeconds); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.LegacyAgentEnv); err != nil {
		return w, err
	}
	return VectorStatshouseIngestionStatus20Read(w, &item.IngestionStatusOk2)
}

func (item *StatshouseSourceBucket2Bytes) Write(w []byte) (_ []byte, err error) {
	if w, err = VectorStatshouseMultiItem0BytesWrite(w, item.Metrics); err != nil {
		return w, err
	}
	if w, err = VectorStatshouseSampleFactor0Write(w, item.SampleFactors); err != nil {
		return w, err
	}
	if w, err = VectorStatshouseSampleFactor0Write(w, item.IngestionStatusOk); err != nil {
		return w, err
	}
	w = basictl.NatWrite(w, item.MissedSeconds)
	w = basictl.IntWrite(w, item.LegacyAgentEnv)
	return VectorStatshouseIngestionStatus20Write(w, item.IngestionStatusOk2)
}

func (item *StatshouseSourceBucket2Bytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x3af6e822); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseSourceBucket2Bytes) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x3af6e822)
	return item.Write(w)
}

func (item StatshouseSourceBucket2Bytes) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func StatshouseSourceBucket2Bytes__ReadJSON(item *StatshouseSourceBucket2Bytes, j interface{}) error {
	return item.readJSON(j)
}
func (item *StatshouseSourceBucket2Bytes) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouse.sourceBucket2", "expected json object")
	}
	_jMetrics := _jm["metrics"]
	delete(_jm, "metrics")
	_jSampleFactors := _jm["sample_factors"]
	delete(_jm, "sample_factors")
	_jIngestionStatusOk := _jm["ingestion_status_ok"]
	delete(_jm, "ingestion_status_ok")
	_jMissedSeconds := _jm["missed_seconds"]
	delete(_jm, "missed_seconds")
	if err := JsonReadUint32(_jMissedSeconds, &item.MissedSeconds); err != nil {
		return err
	}
	_jLegacyAgentEnv := _jm["legacy_agent_env"]
	delete(_jm, "legacy_agent_env")
	if err := JsonReadInt32(_jLegacyAgentEnv, &item.LegacyAgentEnv); err != nil {
		return err
	}
	_jIngestionStatusOk2 := _jm["ingestion_status_ok2"]
	delete(_jm, "ingestion_status_ok2")
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouse.sourceBucket2", k)
	}
	if err := VectorStatshouseMultiItem0BytesReadJSON(_jMetrics, &item.Metrics); err != nil {
		return err
	}
	if err := VectorStatshouseSampleFactor0ReadJSON(_jSampleFactors, &item.SampleFactors); err != nil {
		return err
	}
	if err := VectorStatshouseSampleFactor0ReadJSON(_jIngestionStatusOk, &item.IngestionStatusOk); err != nil {
		return err
	}
	if err := VectorStatshouseIngestionStatus20ReadJSON(_jIngestionStatusOk2, &item.IngestionStatusOk2); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseSourceBucket2Bytes) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if len(item.Metrics) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"metrics":`...)
		if w, err = VectorStatshouseMultiItem0BytesWriteJSON(w, item.Metrics); err != nil {
			return w, err
		}
	}
	if len(item.SampleFactors) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"sample_factors":`...)
		if w, err = VectorStatshouseSampleFactor0WriteJSON(w, item.SampleFactors); err != nil {
			return w, err
		}
	}
	if len(item.IngestionStatusOk) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"ingestion_status_ok":`...)
		if w, err = VectorStatshouseSampleFactor0WriteJSON(w, item.IngestionStatusOk); err != nil {
			return w, err
		}
	}
	if item.MissedSeconds != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"missed_seconds":`...)
		w = basictl.JSONWriteUint32(w, item.MissedSeconds)
	}
	if item.LegacyAgentEnv != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"legacy_agent_env":`...)
		w = basictl.JSONWriteInt32(w, item.LegacyAgentEnv)
	}
	if len(item.IngestionStatusOk2) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"ingestion_status_ok2":`...)
		if w, err = VectorStatshouseIngestionStatus20WriteJSON(w, item.IngestionStatusOk2); err != nil {
			return w, err
		}
	}
	return append(w, '}'), nil
}

func (item *StatshouseSourceBucket2Bytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseSourceBucket2Bytes) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("statshouse.sourceBucket2", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("statshouse.sourceBucket2", err.Error())
	}
	return nil
}
