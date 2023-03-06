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

type StatshousePromTarget struct {
	FieldsMask     uint32
	JobName        string
	Url            string
	Labels         map[string]string
	ScrapeInterval int64
	// HonorTimestamps (TrueType) // Conditional: item.FieldsMask.0
	// HonorLabels (TrueType) // Conditional: item.FieldsMask.1
	ScrapeTimeout         int64
	BodySizeLimit         int64
	LabelLimit            int64
	LabelNameLengthLimit  int64
	LabelValueLengthLimit int64
	HttpClientConfig      string
}

func (StatshousePromTarget) TLName() string { return "statshouse.promTarget" }
func (StatshousePromTarget) TLTag() uint32  { return 0xac5296df }

func (item *StatshousePromTarget) SetHonorTimestamps(v bool) {
	if v {
		item.FieldsMask |= 1 << 0
	} else {
		item.FieldsMask &^= 1 << 0
	}
}
func (item StatshousePromTarget) IsSetHonorTimestamps() bool { return item.FieldsMask&(1<<0) != 0 }

func (item *StatshousePromTarget) SetHonorLabels(v bool) {
	if v {
		item.FieldsMask |= 1 << 1
	} else {
		item.FieldsMask &^= 1 << 1
	}
}
func (item StatshousePromTarget) IsSetHonorLabels() bool { return item.FieldsMask&(1<<1) != 0 }

func (item *StatshousePromTarget) Reset() {
	item.FieldsMask = 0
	item.JobName = ""
	item.Url = ""
	VectorDictionaryFieldString0Reset(item.Labels)
	item.ScrapeInterval = 0
	item.ScrapeTimeout = 0
	item.BodySizeLimit = 0
	item.LabelLimit = 0
	item.LabelNameLengthLimit = 0
	item.LabelValueLengthLimit = 0
	item.HttpClientConfig = ""
}

func (item *StatshousePromTarget) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = basictl.StringRead(w, &item.JobName); err != nil {
		return w, err
	}
	if w, err = basictl.StringRead(w, &item.Url); err != nil {
		return w, err
	}
	if w, err = VectorDictionaryFieldString0Read(w, &item.Labels); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.ScrapeInterval); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.ScrapeTimeout); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.BodySizeLimit); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.LabelLimit); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.LabelNameLengthLimit); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.LabelValueLengthLimit); err != nil {
		return w, err
	}
	return basictl.StringRead(w, &item.HttpClientConfig)
}

func (item *StatshousePromTarget) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	if w, err = basictl.StringWrite(w, item.JobName); err != nil {
		return w, err
	}
	if w, err = basictl.StringWrite(w, item.Url); err != nil {
		return w, err
	}
	if w, err = VectorDictionaryFieldString0Write(w, item.Labels); err != nil {
		return w, err
	}
	w = basictl.LongWrite(w, item.ScrapeInterval)
	w = basictl.LongWrite(w, item.ScrapeTimeout)
	w = basictl.LongWrite(w, item.BodySizeLimit)
	w = basictl.LongWrite(w, item.LabelLimit)
	w = basictl.LongWrite(w, item.LabelNameLengthLimit)
	w = basictl.LongWrite(w, item.LabelValueLengthLimit)
	return basictl.StringWrite(w, item.HttpClientConfig)
}

func (item *StatshousePromTarget) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xac5296df); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshousePromTarget) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xac5296df)
	return item.Write(w)
}

func (item StatshousePromTarget) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func StatshousePromTarget__ReadJSON(item *StatshousePromTarget, j interface{}) error {
	return item.readJSON(j)
}
func (item *StatshousePromTarget) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouse.promTarget", "expected json object")
	}
	_jFieldsMask := _jm["fields_mask"]
	delete(_jm, "fields_mask")
	if err := JsonReadUint32(_jFieldsMask, &item.FieldsMask); err != nil {
		return err
	}
	_jJobName := _jm["job_name"]
	delete(_jm, "job_name")
	if err := JsonReadString(_jJobName, &item.JobName); err != nil {
		return err
	}
	_jUrl := _jm["url"]
	delete(_jm, "url")
	if err := JsonReadString(_jUrl, &item.Url); err != nil {
		return err
	}
	_jLabels := _jm["labels"]
	delete(_jm, "labels")
	_jScrapeInterval := _jm["scrape_interval"]
	delete(_jm, "scrape_interval")
	if err := JsonReadInt64(_jScrapeInterval, &item.ScrapeInterval); err != nil {
		return err
	}
	_jHonorTimestamps := _jm["honor_timestamps"]
	delete(_jm, "honor_timestamps")
	_jHonorLabels := _jm["honor_labels"]
	delete(_jm, "honor_labels")
	_jScrapeTimeout := _jm["scrape_timeout"]
	delete(_jm, "scrape_timeout")
	if err := JsonReadInt64(_jScrapeTimeout, &item.ScrapeTimeout); err != nil {
		return err
	}
	_jBodySizeLimit := _jm["body_size_limit"]
	delete(_jm, "body_size_limit")
	if err := JsonReadInt64(_jBodySizeLimit, &item.BodySizeLimit); err != nil {
		return err
	}
	_jLabelLimit := _jm["label_limit"]
	delete(_jm, "label_limit")
	if err := JsonReadInt64(_jLabelLimit, &item.LabelLimit); err != nil {
		return err
	}
	_jLabelNameLengthLimit := _jm["label_name_length_limit"]
	delete(_jm, "label_name_length_limit")
	if err := JsonReadInt64(_jLabelNameLengthLimit, &item.LabelNameLengthLimit); err != nil {
		return err
	}
	_jLabelValueLengthLimit := _jm["label_value_length_limit"]
	delete(_jm, "label_value_length_limit")
	if err := JsonReadInt64(_jLabelValueLengthLimit, &item.LabelValueLengthLimit); err != nil {
		return err
	}
	_jHttpClientConfig := _jm["http_client_config"]
	delete(_jm, "http_client_config")
	if err := JsonReadString(_jHttpClientConfig, &item.HttpClientConfig); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouse.promTarget", k)
	}
	if _jHonorTimestamps != nil {
		_bit := false
		if err := JsonReadBool(_jHonorTimestamps, &_bit); err != nil {
			return err
		}
		if _bit {
			item.FieldsMask |= 1 << 0
		} else {
			item.FieldsMask &^= 1 << 0
		}
	}
	if _jHonorLabels != nil {
		_bit := false
		if err := JsonReadBool(_jHonorLabels, &_bit); err != nil {
			return err
		}
		if _bit {
			item.FieldsMask |= 1 << 1
		} else {
			item.FieldsMask &^= 1 << 1
		}
	}
	if err := VectorDictionaryFieldString0ReadJSON(_jLabels, &item.Labels); err != nil {
		return err
	}
	return nil
}

func (item *StatshousePromTarget) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldsMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"fields_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldsMask)
	}
	if len(item.JobName) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"job_name":`...)
		w = basictl.JSONWriteString(w, item.JobName)
	}
	if len(item.Url) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"url":`...)
		w = basictl.JSONWriteString(w, item.Url)
	}
	if len(item.Labels) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"labels":`...)
		if w, err = VectorDictionaryFieldString0WriteJSON(w, item.Labels); err != nil {
			return w, err
		}
	}
	if item.ScrapeInterval != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"scrape_interval":`...)
		w = basictl.JSONWriteInt64(w, item.ScrapeInterval)
	}
	if item.FieldsMask&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"honor_timestamps":true`...)
	}
	if item.FieldsMask&(1<<1) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"honor_labels":true`...)
	}
	if item.ScrapeTimeout != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"scrape_timeout":`...)
		w = basictl.JSONWriteInt64(w, item.ScrapeTimeout)
	}
	if item.BodySizeLimit != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"body_size_limit":`...)
		w = basictl.JSONWriteInt64(w, item.BodySizeLimit)
	}
	if item.LabelLimit != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"label_limit":`...)
		w = basictl.JSONWriteInt64(w, item.LabelLimit)
	}
	if item.LabelNameLengthLimit != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"label_name_length_limit":`...)
		w = basictl.JSONWriteInt64(w, item.LabelNameLengthLimit)
	}
	if item.LabelValueLengthLimit != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"label_value_length_limit":`...)
		w = basictl.JSONWriteInt64(w, item.LabelValueLengthLimit)
	}
	if len(item.HttpClientConfig) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"http_client_config":`...)
		w = basictl.JSONWriteString(w, item.HttpClientConfig)
	}
	return append(w, '}'), nil
}

func (item *StatshousePromTarget) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshousePromTarget) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("statshouse.promTarget", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("statshouse.promTarget", err.Error())
	}
	return nil
}

type StatshousePromTargetBytes struct {
	FieldsMask     uint32
	JobName        []byte
	Url            []byte
	Labels         []DictionaryFieldStringBytes
	ScrapeInterval int64
	// HonorTimestamps (TrueType) // Conditional: item.FieldsMask.0
	// HonorLabels (TrueType) // Conditional: item.FieldsMask.1
	ScrapeTimeout         int64
	BodySizeLimit         int64
	LabelLimit            int64
	LabelNameLengthLimit  int64
	LabelValueLengthLimit int64
	HttpClientConfig      []byte
}

func (StatshousePromTargetBytes) TLName() string { return "statshouse.promTarget" }
func (StatshousePromTargetBytes) TLTag() uint32  { return 0xac5296df }

func (item *StatshousePromTargetBytes) SetHonorTimestamps(v bool) {
	if v {
		item.FieldsMask |= 1 << 0
	} else {
		item.FieldsMask &^= 1 << 0
	}
}
func (item StatshousePromTargetBytes) IsSetHonorTimestamps() bool { return item.FieldsMask&(1<<0) != 0 }

func (item *StatshousePromTargetBytes) SetHonorLabels(v bool) {
	if v {
		item.FieldsMask |= 1 << 1
	} else {
		item.FieldsMask &^= 1 << 1
	}
}
func (item StatshousePromTargetBytes) IsSetHonorLabels() bool { return item.FieldsMask&(1<<1) != 0 }

func (item *StatshousePromTargetBytes) Reset() {
	item.FieldsMask = 0
	item.JobName = item.JobName[:0]
	item.Url = item.Url[:0]
	item.Labels = item.Labels[:0]
	item.ScrapeInterval = 0
	item.ScrapeTimeout = 0
	item.BodySizeLimit = 0
	item.LabelLimit = 0
	item.LabelNameLengthLimit = 0
	item.LabelValueLengthLimit = 0
	item.HttpClientConfig = item.HttpClientConfig[:0]
}

func (item *StatshousePromTargetBytes) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = basictl.StringReadBytes(w, &item.JobName); err != nil {
		return w, err
	}
	if w, err = basictl.StringReadBytes(w, &item.Url); err != nil {
		return w, err
	}
	if w, err = VectorDictionaryFieldString0BytesRead(w, &item.Labels); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.ScrapeInterval); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.ScrapeTimeout); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.BodySizeLimit); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.LabelLimit); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.LabelNameLengthLimit); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.LabelValueLengthLimit); err != nil {
		return w, err
	}
	return basictl.StringReadBytes(w, &item.HttpClientConfig)
}

func (item *StatshousePromTargetBytes) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	if w, err = basictl.StringWriteBytes(w, item.JobName); err != nil {
		return w, err
	}
	if w, err = basictl.StringWriteBytes(w, item.Url); err != nil {
		return w, err
	}
	if w, err = VectorDictionaryFieldString0BytesWrite(w, item.Labels); err != nil {
		return w, err
	}
	w = basictl.LongWrite(w, item.ScrapeInterval)
	w = basictl.LongWrite(w, item.ScrapeTimeout)
	w = basictl.LongWrite(w, item.BodySizeLimit)
	w = basictl.LongWrite(w, item.LabelLimit)
	w = basictl.LongWrite(w, item.LabelNameLengthLimit)
	w = basictl.LongWrite(w, item.LabelValueLengthLimit)
	return basictl.StringWriteBytes(w, item.HttpClientConfig)
}

func (item *StatshousePromTargetBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xac5296df); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshousePromTargetBytes) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xac5296df)
	return item.Write(w)
}

func (item StatshousePromTargetBytes) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func StatshousePromTargetBytes__ReadJSON(item *StatshousePromTargetBytes, j interface{}) error {
	return item.readJSON(j)
}
func (item *StatshousePromTargetBytes) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouse.promTarget", "expected json object")
	}
	_jFieldsMask := _jm["fields_mask"]
	delete(_jm, "fields_mask")
	if err := JsonReadUint32(_jFieldsMask, &item.FieldsMask); err != nil {
		return err
	}
	_jJobName := _jm["job_name"]
	delete(_jm, "job_name")
	if err := JsonReadStringBytes(_jJobName, &item.JobName); err != nil {
		return err
	}
	_jUrl := _jm["url"]
	delete(_jm, "url")
	if err := JsonReadStringBytes(_jUrl, &item.Url); err != nil {
		return err
	}
	_jLabels := _jm["labels"]
	delete(_jm, "labels")
	_jScrapeInterval := _jm["scrape_interval"]
	delete(_jm, "scrape_interval")
	if err := JsonReadInt64(_jScrapeInterval, &item.ScrapeInterval); err != nil {
		return err
	}
	_jHonorTimestamps := _jm["honor_timestamps"]
	delete(_jm, "honor_timestamps")
	_jHonorLabels := _jm["honor_labels"]
	delete(_jm, "honor_labels")
	_jScrapeTimeout := _jm["scrape_timeout"]
	delete(_jm, "scrape_timeout")
	if err := JsonReadInt64(_jScrapeTimeout, &item.ScrapeTimeout); err != nil {
		return err
	}
	_jBodySizeLimit := _jm["body_size_limit"]
	delete(_jm, "body_size_limit")
	if err := JsonReadInt64(_jBodySizeLimit, &item.BodySizeLimit); err != nil {
		return err
	}
	_jLabelLimit := _jm["label_limit"]
	delete(_jm, "label_limit")
	if err := JsonReadInt64(_jLabelLimit, &item.LabelLimit); err != nil {
		return err
	}
	_jLabelNameLengthLimit := _jm["label_name_length_limit"]
	delete(_jm, "label_name_length_limit")
	if err := JsonReadInt64(_jLabelNameLengthLimit, &item.LabelNameLengthLimit); err != nil {
		return err
	}
	_jLabelValueLengthLimit := _jm["label_value_length_limit"]
	delete(_jm, "label_value_length_limit")
	if err := JsonReadInt64(_jLabelValueLengthLimit, &item.LabelValueLengthLimit); err != nil {
		return err
	}
	_jHttpClientConfig := _jm["http_client_config"]
	delete(_jm, "http_client_config")
	if err := JsonReadStringBytes(_jHttpClientConfig, &item.HttpClientConfig); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouse.promTarget", k)
	}
	if _jHonorTimestamps != nil {
		_bit := false
		if err := JsonReadBool(_jHonorTimestamps, &_bit); err != nil {
			return err
		}
		if _bit {
			item.FieldsMask |= 1 << 0
		} else {
			item.FieldsMask &^= 1 << 0
		}
	}
	if _jHonorLabels != nil {
		_bit := false
		if err := JsonReadBool(_jHonorLabels, &_bit); err != nil {
			return err
		}
		if _bit {
			item.FieldsMask |= 1 << 1
		} else {
			item.FieldsMask &^= 1 << 1
		}
	}
	if err := VectorDictionaryFieldString0BytesReadJSON(_jLabels, &item.Labels); err != nil {
		return err
	}
	return nil
}

func (item *StatshousePromTargetBytes) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldsMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"fields_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldsMask)
	}
	if len(item.JobName) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"job_name":`...)
		w = basictl.JSONWriteStringBytes(w, item.JobName)
	}
	if len(item.Url) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"url":`...)
		w = basictl.JSONWriteStringBytes(w, item.Url)
	}
	if len(item.Labels) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"labels":`...)
		if w, err = VectorDictionaryFieldString0BytesWriteJSON(w, item.Labels); err != nil {
			return w, err
		}
	}
	if item.ScrapeInterval != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"scrape_interval":`...)
		w = basictl.JSONWriteInt64(w, item.ScrapeInterval)
	}
	if item.FieldsMask&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"honor_timestamps":true`...)
	}
	if item.FieldsMask&(1<<1) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"honor_labels":true`...)
	}
	if item.ScrapeTimeout != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"scrape_timeout":`...)
		w = basictl.JSONWriteInt64(w, item.ScrapeTimeout)
	}
	if item.BodySizeLimit != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"body_size_limit":`...)
		w = basictl.JSONWriteInt64(w, item.BodySizeLimit)
	}
	if item.LabelLimit != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"label_limit":`...)
		w = basictl.JSONWriteInt64(w, item.LabelLimit)
	}
	if item.LabelNameLengthLimit != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"label_name_length_limit":`...)
		w = basictl.JSONWriteInt64(w, item.LabelNameLengthLimit)
	}
	if item.LabelValueLengthLimit != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"label_value_length_limit":`...)
		w = basictl.JSONWriteInt64(w, item.LabelValueLengthLimit)
	}
	if len(item.HttpClientConfig) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"http_client_config":`...)
		w = basictl.JSONWriteStringBytes(w, item.HttpClientConfig)
	}
	return append(w, '}'), nil
}

func (item *StatshousePromTargetBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshousePromTargetBytes) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("statshouse.promTarget", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("statshouse.promTarget", err.Error())
	}
	return nil
}

func VectorStatshousePromTarget0Read(w []byte, vec *[]StatshousePromTarget) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]StatshousePromTarget, l)
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

func VectorStatshousePromTarget0Write(w []byte, vec []StatshousePromTarget) (_ []byte, err error) {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		if w, err = elem.Write(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorStatshousePromTarget0ReadJSON(j interface{}, vec *[]StatshousePromTarget) error {
	l, _arr, err := JsonReadArray("[]StatshousePromTarget", j)
	if err != nil {
		return err
	}
	if cap(*vec) < l {
		*vec = make([]StatshousePromTarget, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if err := StatshousePromTarget__ReadJSON(&(*vec)[i], _arr[i]); err != nil {
			return err
		}
	}
	return nil
}

func VectorStatshousePromTarget0WriteJSON(w []byte, vec []StatshousePromTarget) (_ []byte, err error) {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		if w, err = elem.WriteJSON(w); err != nil {
			return w, err
		}
	}
	return append(w, ']'), nil
}

func VectorStatshousePromTarget0BytesRead(w []byte, vec *[]StatshousePromTargetBytes) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]StatshousePromTargetBytes, l)
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

func VectorStatshousePromTarget0BytesWrite(w []byte, vec []StatshousePromTargetBytes) (_ []byte, err error) {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		if w, err = elem.Write(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorStatshousePromTarget0BytesReadJSON(j interface{}, vec *[]StatshousePromTargetBytes) error {
	l, _arr, err := JsonReadArray("[]StatshousePromTargetBytes", j)
	if err != nil {
		return err
	}
	if cap(*vec) < l {
		*vec = make([]StatshousePromTargetBytes, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if err := StatshousePromTargetBytes__ReadJSON(&(*vec)[i], _arr[i]); err != nil {
			return err
		}
	}
	return nil
}

func VectorStatshousePromTarget0BytesWriteJSON(w []byte, vec []StatshousePromTargetBytes) (_ []byte, err error) {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		if w, err = elem.WriteJSON(w); err != nil {
			return w, err
		}
	}
	return append(w, ']'), nil
}
