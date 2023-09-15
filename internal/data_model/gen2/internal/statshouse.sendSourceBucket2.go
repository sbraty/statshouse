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

type StatshouseSendSourceBucket2 struct {
	FieldsMask uint32
	Header     StatshouseCommonProxyHeader
	Time       uint32
	// Historic (TrueType) // Conditional: item.FieldsMask.0
	// Spare (TrueType) // Conditional: item.FieldsMask.1
	BuildCommit            string
	BuildCommitDate        int32
	BuildCommitTs          int32
	QueueSizeDisk          int32
	QueueSizeMemory        int32
	QueueSizeDiskSum       int32 // Conditional: item.FieldsMask.2
	QueueSizeMemorySum     int32 // Conditional: item.FieldsMask.2
	QueueSizeDiskUnsent    int32 // Conditional: item.FieldsMask.3
	QueueSizeDiskSumUnsent int32 // Conditional: item.FieldsMask.3
	OriginalSize           int32
	CompressedData         string
}

func (StatshouseSendSourceBucket2) TLName() string { return "statshouse.sendSourceBucket2" }
func (StatshouseSendSourceBucket2) TLTag() uint32  { return 0x44575940 }

func (item *StatshouseSendSourceBucket2) SetHistoric(v bool) {
	if v {
		item.FieldsMask |= 1 << 0
	} else {
		item.FieldsMask &^= 1 << 0
	}
}
func (item StatshouseSendSourceBucket2) IsSetHistoric() bool { return item.FieldsMask&(1<<0) != 0 }

func (item *StatshouseSendSourceBucket2) SetSpare(v bool) {
	if v {
		item.FieldsMask |= 1 << 1
	} else {
		item.FieldsMask &^= 1 << 1
	}
}
func (item StatshouseSendSourceBucket2) IsSetSpare() bool { return item.FieldsMask&(1<<1) != 0 }

func (item *StatshouseSendSourceBucket2) SetQueueSizeDiskSum(v int32) {
	item.QueueSizeDiskSum = v
	item.FieldsMask |= 1 << 2
}
func (item *StatshouseSendSourceBucket2) ClearQueueSizeDiskSum() {
	item.QueueSizeDiskSum = 0
	item.FieldsMask &^= 1 << 2
}
func (item StatshouseSendSourceBucket2) IsSetQueueSizeDiskSum() bool {
	return item.FieldsMask&(1<<2) != 0
}

func (item *StatshouseSendSourceBucket2) SetQueueSizeMemorySum(v int32) {
	item.QueueSizeMemorySum = v
	item.FieldsMask |= 1 << 2
}
func (item *StatshouseSendSourceBucket2) ClearQueueSizeMemorySum() {
	item.QueueSizeMemorySum = 0
	item.FieldsMask &^= 1 << 2
}
func (item StatshouseSendSourceBucket2) IsSetQueueSizeMemorySum() bool {
	return item.FieldsMask&(1<<2) != 0
}

func (item *StatshouseSendSourceBucket2) SetQueueSizeDiskUnsent(v int32) {
	item.QueueSizeDiskUnsent = v
	item.FieldsMask |= 1 << 3
}
func (item *StatshouseSendSourceBucket2) ClearQueueSizeDiskUnsent() {
	item.QueueSizeDiskUnsent = 0
	item.FieldsMask &^= 1 << 3
}
func (item StatshouseSendSourceBucket2) IsSetQueueSizeDiskUnsent() bool {
	return item.FieldsMask&(1<<3) != 0
}

func (item *StatshouseSendSourceBucket2) SetQueueSizeDiskSumUnsent(v int32) {
	item.QueueSizeDiskSumUnsent = v
	item.FieldsMask |= 1 << 3
}
func (item *StatshouseSendSourceBucket2) ClearQueueSizeDiskSumUnsent() {
	item.QueueSizeDiskSumUnsent = 0
	item.FieldsMask &^= 1 << 3
}
func (item StatshouseSendSourceBucket2) IsSetQueueSizeDiskSumUnsent() bool {
	return item.FieldsMask&(1<<3) != 0
}

func (item *StatshouseSendSourceBucket2) Reset() {
	item.FieldsMask = 0
	item.Header.Reset()
	item.Time = 0
	item.BuildCommit = ""
	item.BuildCommitDate = 0
	item.BuildCommitTs = 0
	item.QueueSizeDisk = 0
	item.QueueSizeMemory = 0
	item.QueueSizeDiskSum = 0
	item.QueueSizeMemorySum = 0
	item.QueueSizeDiskUnsent = 0
	item.QueueSizeDiskSumUnsent = 0
	item.OriginalSize = 0
	item.CompressedData = ""
}

func (item *StatshouseSendSourceBucket2) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = item.Header.Read(w, item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = basictl.NatRead(w, &item.Time); err != nil {
		return w, err
	}
	if w, err = basictl.StringRead(w, &item.BuildCommit); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.BuildCommitDate); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.BuildCommitTs); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.QueueSizeDisk); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.QueueSizeMemory); err != nil {
		return w, err
	}
	if item.FieldsMask&(1<<2) != 0 {
		if w, err = basictl.IntRead(w, &item.QueueSizeDiskSum); err != nil {
			return w, err
		}
	} else {
		item.QueueSizeDiskSum = 0
	}
	if item.FieldsMask&(1<<2) != 0 {
		if w, err = basictl.IntRead(w, &item.QueueSizeMemorySum); err != nil {
			return w, err
		}
	} else {
		item.QueueSizeMemorySum = 0
	}
	if item.FieldsMask&(1<<3) != 0 {
		if w, err = basictl.IntRead(w, &item.QueueSizeDiskUnsent); err != nil {
			return w, err
		}
	} else {
		item.QueueSizeDiskUnsent = 0
	}
	if item.FieldsMask&(1<<3) != 0 {
		if w, err = basictl.IntRead(w, &item.QueueSizeDiskSumUnsent); err != nil {
			return w, err
		}
	} else {
		item.QueueSizeDiskSumUnsent = 0
	}
	if w, err = basictl.IntRead(w, &item.OriginalSize); err != nil {
		return w, err
	}
	return basictl.StringRead(w, &item.CompressedData)
}

func (item *StatshouseSendSourceBucket2) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	if w, err = item.Header.Write(w, item.FieldsMask); err != nil {
		return w, err
	}
	w = basictl.NatWrite(w, item.Time)
	if w, err = basictl.StringWrite(w, item.BuildCommit); err != nil {
		return w, err
	}
	w = basictl.IntWrite(w, item.BuildCommitDate)
	w = basictl.IntWrite(w, item.BuildCommitTs)
	w = basictl.IntWrite(w, item.QueueSizeDisk)
	w = basictl.IntWrite(w, item.QueueSizeMemory)
	if item.FieldsMask&(1<<2) != 0 {
		w = basictl.IntWrite(w, item.QueueSizeDiskSum)
	}
	if item.FieldsMask&(1<<2) != 0 {
		w = basictl.IntWrite(w, item.QueueSizeMemorySum)
	}
	if item.FieldsMask&(1<<3) != 0 {
		w = basictl.IntWrite(w, item.QueueSizeDiskUnsent)
	}
	if item.FieldsMask&(1<<3) != 0 {
		w = basictl.IntWrite(w, item.QueueSizeDiskSumUnsent)
	}
	w = basictl.IntWrite(w, item.OriginalSize)
	return basictl.StringWrite(w, item.CompressedData)
}

func (item *StatshouseSendSourceBucket2) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x44575940); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseSendSourceBucket2) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x44575940)
	return item.Write(w)
}

func (item *StatshouseSendSourceBucket2) ReadResult(w []byte, ret *string) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xb5286e24); err != nil {
		return w, err
	}
	return basictl.StringRead(w, ret)
}

func (item *StatshouseSendSourceBucket2) WriteResult(w []byte, ret string) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0xb5286e24)
	return basictl.StringWrite(w, ret)
}

func (item *StatshouseSendSourceBucket2) ReadResultJSON(j interface{}, ret *string) error {
	if err := JsonReadString(j, ret); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseSendSourceBucket2) WriteResultJSON(w []byte, ret string) (_ []byte, err error) {
	return item.writeResultJSON(false, w, ret)
}

func (item *StatshouseSendSourceBucket2) writeResultJSON(short bool, w []byte, ret string) (_ []byte, err error) {
	w = basictl.JSONWriteString(w, ret)
	return w, nil
}

func (item *StatshouseSendSourceBucket2) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret string
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *StatshouseSendSourceBucket2) ReadResultWriteResultJSONShort(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret string
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(true, w, ret)
	return r, w, err
}

func (item *StatshouseSendSourceBucket2) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := JsonBytesToInterface(r)
	if err != nil {
		return r, w, ErrorInvalidJSON("statshouse.sendSourceBucket2", err.Error())
	}
	var ret string
	if err = item.ReadResultJSON(j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item StatshouseSendSourceBucket2) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func StatshouseSendSourceBucket2__ReadJSON(item *StatshouseSendSourceBucket2, j interface{}) error {
	return item.readJSON(j)
}
func (item *StatshouseSendSourceBucket2) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouse.sendSourceBucket2", "expected json object")
	}
	_jFieldsMask := _jm["fields_mask"]
	delete(_jm, "fields_mask")
	if err := JsonReadUint32(_jFieldsMask, &item.FieldsMask); err != nil {
		return err
	}
	_jHeader := _jm["header"]
	delete(_jm, "header")
	_jTime := _jm["time"]
	delete(_jm, "time")
	if err := JsonReadUint32(_jTime, &item.Time); err != nil {
		return err
	}
	_jHistoric := _jm["historic"]
	delete(_jm, "historic")
	_jSpare := _jm["spare"]
	delete(_jm, "spare")
	_jBuildCommit := _jm["build_commit"]
	delete(_jm, "build_commit")
	if err := JsonReadString(_jBuildCommit, &item.BuildCommit); err != nil {
		return err
	}
	_jBuildCommitDate := _jm["build_commit_date"]
	delete(_jm, "build_commit_date")
	if err := JsonReadInt32(_jBuildCommitDate, &item.BuildCommitDate); err != nil {
		return err
	}
	_jBuildCommitTs := _jm["build_commit_ts"]
	delete(_jm, "build_commit_ts")
	if err := JsonReadInt32(_jBuildCommitTs, &item.BuildCommitTs); err != nil {
		return err
	}
	_jQueueSizeDisk := _jm["queue_size_disk"]
	delete(_jm, "queue_size_disk")
	if err := JsonReadInt32(_jQueueSizeDisk, &item.QueueSizeDisk); err != nil {
		return err
	}
	_jQueueSizeMemory := _jm["queue_size_memory"]
	delete(_jm, "queue_size_memory")
	if err := JsonReadInt32(_jQueueSizeMemory, &item.QueueSizeMemory); err != nil {
		return err
	}
	_jQueueSizeDiskSum := _jm["queue_size_disk_sum"]
	delete(_jm, "queue_size_disk_sum")
	_jQueueSizeMemorySum := _jm["queue_size_memory_sum"]
	delete(_jm, "queue_size_memory_sum")
	_jQueueSizeDiskUnsent := _jm["queue_size_disk_unsent"]
	delete(_jm, "queue_size_disk_unsent")
	_jQueueSizeDiskSumUnsent := _jm["queue_size_disk_sum_unsent"]
	delete(_jm, "queue_size_disk_sum_unsent")
	_jOriginalSize := _jm["original_size"]
	delete(_jm, "original_size")
	if err := JsonReadInt32(_jOriginalSize, &item.OriginalSize); err != nil {
		return err
	}
	_jCompressedData := _jm["compressed_data"]
	delete(_jm, "compressed_data")
	if err := JsonReadString(_jCompressedData, &item.CompressedData); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouse.sendSourceBucket2", k)
	}
	if _jHistoric != nil {
		_bit := false
		if err := JsonReadBool(_jHistoric, &_bit); err != nil {
			return err
		}
		if _bit {
			item.FieldsMask |= 1 << 0
		} else {
			item.FieldsMask &^= 1 << 0
		}
	}
	if _jSpare != nil {
		_bit := false
		if err := JsonReadBool(_jSpare, &_bit); err != nil {
			return err
		}
		if _bit {
			item.FieldsMask |= 1 << 1
		} else {
			item.FieldsMask &^= 1 << 1
		}
	}
	if _jQueueSizeDiskSum != nil {
		item.FieldsMask |= 1 << 2
	}
	if _jQueueSizeMemorySum != nil {
		item.FieldsMask |= 1 << 2
	}
	if _jQueueSizeDiskUnsent != nil {
		item.FieldsMask |= 1 << 3
	}
	if _jQueueSizeDiskSumUnsent != nil {
		item.FieldsMask |= 1 << 3
	}
	if err := StatshouseCommonProxyHeader__ReadJSON(&item.Header, _jHeader, item.FieldsMask); err != nil {
		return err
	}
	if _jQueueSizeDiskSum != nil {
		if err := JsonReadInt32(_jQueueSizeDiskSum, &item.QueueSizeDiskSum); err != nil {
			return err
		}
	} else {
		item.QueueSizeDiskSum = 0
	}
	if _jQueueSizeMemorySum != nil {
		if err := JsonReadInt32(_jQueueSizeMemorySum, &item.QueueSizeMemorySum); err != nil {
			return err
		}
	} else {
		item.QueueSizeMemorySum = 0
	}
	if _jQueueSizeDiskUnsent != nil {
		if err := JsonReadInt32(_jQueueSizeDiskUnsent, &item.QueueSizeDiskUnsent); err != nil {
			return err
		}
	} else {
		item.QueueSizeDiskUnsent = 0
	}
	if _jQueueSizeDiskSumUnsent != nil {
		if err := JsonReadInt32(_jQueueSizeDiskSumUnsent, &item.QueueSizeDiskSumUnsent); err != nil {
			return err
		}
	} else {
		item.QueueSizeDiskSumUnsent = 0
	}
	return nil
}

func (item *StatshouseSendSourceBucket2) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *StatshouseSendSourceBucket2) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldsMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"fields_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldsMask)
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"header":`...)
	if w, err = item.Header.WriteJSONOpt(short, w, item.FieldsMask); err != nil {
		return w, err
	}
	if item.Time != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"time":`...)
		w = basictl.JSONWriteUint32(w, item.Time)
	}
	if item.FieldsMask&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"historic":true`...)
	}
	if item.FieldsMask&(1<<1) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"spare":true`...)
	}
	if len(item.BuildCommit) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"build_commit":`...)
		w = basictl.JSONWriteString(w, item.BuildCommit)
	}
	if item.BuildCommitDate != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"build_commit_date":`...)
		w = basictl.JSONWriteInt32(w, item.BuildCommitDate)
	}
	if item.BuildCommitTs != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"build_commit_ts":`...)
		w = basictl.JSONWriteInt32(w, item.BuildCommitTs)
	}
	if item.QueueSizeDisk != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"queue_size_disk":`...)
		w = basictl.JSONWriteInt32(w, item.QueueSizeDisk)
	}
	if item.QueueSizeMemory != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"queue_size_memory":`...)
		w = basictl.JSONWriteInt32(w, item.QueueSizeMemory)
	}
	if item.FieldsMask&(1<<2) != 0 {
		if item.QueueSizeDiskSum != 0 {
			w = basictl.JSONAddCommaIfNeeded(w)
			w = append(w, `"queue_size_disk_sum":`...)
			w = basictl.JSONWriteInt32(w, item.QueueSizeDiskSum)
		}
	}
	if item.FieldsMask&(1<<2) != 0 {
		if item.QueueSizeMemorySum != 0 {
			w = basictl.JSONAddCommaIfNeeded(w)
			w = append(w, `"queue_size_memory_sum":`...)
			w = basictl.JSONWriteInt32(w, item.QueueSizeMemorySum)
		}
	}
	if item.FieldsMask&(1<<3) != 0 {
		if item.QueueSizeDiskUnsent != 0 {
			w = basictl.JSONAddCommaIfNeeded(w)
			w = append(w, `"queue_size_disk_unsent":`...)
			w = basictl.JSONWriteInt32(w, item.QueueSizeDiskUnsent)
		}
	}
	if item.FieldsMask&(1<<3) != 0 {
		if item.QueueSizeDiskSumUnsent != 0 {
			w = basictl.JSONAddCommaIfNeeded(w)
			w = append(w, `"queue_size_disk_sum_unsent":`...)
			w = basictl.JSONWriteInt32(w, item.QueueSizeDiskSumUnsent)
		}
	}
	if item.OriginalSize != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"original_size":`...)
		w = basictl.JSONWriteInt32(w, item.OriginalSize)
	}
	if len(item.CompressedData) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"compressed_data":`...)
		w = basictl.JSONWriteString(w, item.CompressedData)
	}
	return append(w, '}'), nil
}

func (item *StatshouseSendSourceBucket2) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseSendSourceBucket2) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("statshouse.sendSourceBucket2", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("statshouse.sendSourceBucket2", err.Error())
	}
	return nil
}

type StatshouseSendSourceBucket2Bytes struct {
	FieldsMask uint32
	Header     StatshouseCommonProxyHeaderBytes
	Time       uint32
	// Historic (TrueType) // Conditional: item.FieldsMask.0
	// Spare (TrueType) // Conditional: item.FieldsMask.1
	BuildCommit            []byte
	BuildCommitDate        int32
	BuildCommitTs          int32
	QueueSizeDisk          int32
	QueueSizeMemory        int32
	QueueSizeDiskSum       int32 // Conditional: item.FieldsMask.2
	QueueSizeMemorySum     int32 // Conditional: item.FieldsMask.2
	QueueSizeDiskUnsent    int32 // Conditional: item.FieldsMask.3
	QueueSizeDiskSumUnsent int32 // Conditional: item.FieldsMask.3
	OriginalSize           int32
	CompressedData         []byte
}

func (StatshouseSendSourceBucket2Bytes) TLName() string { return "statshouse.sendSourceBucket2" }
func (StatshouseSendSourceBucket2Bytes) TLTag() uint32  { return 0x44575940 }

func (item *StatshouseSendSourceBucket2Bytes) SetHistoric(v bool) {
	if v {
		item.FieldsMask |= 1 << 0
	} else {
		item.FieldsMask &^= 1 << 0
	}
}
func (item StatshouseSendSourceBucket2Bytes) IsSetHistoric() bool { return item.FieldsMask&(1<<0) != 0 }

func (item *StatshouseSendSourceBucket2Bytes) SetSpare(v bool) {
	if v {
		item.FieldsMask |= 1 << 1
	} else {
		item.FieldsMask &^= 1 << 1
	}
}
func (item StatshouseSendSourceBucket2Bytes) IsSetSpare() bool { return item.FieldsMask&(1<<1) != 0 }

func (item *StatshouseSendSourceBucket2Bytes) SetQueueSizeDiskSum(v int32) {
	item.QueueSizeDiskSum = v
	item.FieldsMask |= 1 << 2
}
func (item *StatshouseSendSourceBucket2Bytes) ClearQueueSizeDiskSum() {
	item.QueueSizeDiskSum = 0
	item.FieldsMask &^= 1 << 2
}
func (item StatshouseSendSourceBucket2Bytes) IsSetQueueSizeDiskSum() bool {
	return item.FieldsMask&(1<<2) != 0
}

func (item *StatshouseSendSourceBucket2Bytes) SetQueueSizeMemorySum(v int32) {
	item.QueueSizeMemorySum = v
	item.FieldsMask |= 1 << 2
}
func (item *StatshouseSendSourceBucket2Bytes) ClearQueueSizeMemorySum() {
	item.QueueSizeMemorySum = 0
	item.FieldsMask &^= 1 << 2
}
func (item StatshouseSendSourceBucket2Bytes) IsSetQueueSizeMemorySum() bool {
	return item.FieldsMask&(1<<2) != 0
}

func (item *StatshouseSendSourceBucket2Bytes) SetQueueSizeDiskUnsent(v int32) {
	item.QueueSizeDiskUnsent = v
	item.FieldsMask |= 1 << 3
}
func (item *StatshouseSendSourceBucket2Bytes) ClearQueueSizeDiskUnsent() {
	item.QueueSizeDiskUnsent = 0
	item.FieldsMask &^= 1 << 3
}
func (item StatshouseSendSourceBucket2Bytes) IsSetQueueSizeDiskUnsent() bool {
	return item.FieldsMask&(1<<3) != 0
}

func (item *StatshouseSendSourceBucket2Bytes) SetQueueSizeDiskSumUnsent(v int32) {
	item.QueueSizeDiskSumUnsent = v
	item.FieldsMask |= 1 << 3
}
func (item *StatshouseSendSourceBucket2Bytes) ClearQueueSizeDiskSumUnsent() {
	item.QueueSizeDiskSumUnsent = 0
	item.FieldsMask &^= 1 << 3
}
func (item StatshouseSendSourceBucket2Bytes) IsSetQueueSizeDiskSumUnsent() bool {
	return item.FieldsMask&(1<<3) != 0
}

func (item *StatshouseSendSourceBucket2Bytes) Reset() {
	item.FieldsMask = 0
	item.Header.Reset()
	item.Time = 0
	item.BuildCommit = item.BuildCommit[:0]
	item.BuildCommitDate = 0
	item.BuildCommitTs = 0
	item.QueueSizeDisk = 0
	item.QueueSizeMemory = 0
	item.QueueSizeDiskSum = 0
	item.QueueSizeMemorySum = 0
	item.QueueSizeDiskUnsent = 0
	item.QueueSizeDiskSumUnsent = 0
	item.OriginalSize = 0
	item.CompressedData = item.CompressedData[:0]
}

func (item *StatshouseSendSourceBucket2Bytes) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = item.Header.Read(w, item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = basictl.NatRead(w, &item.Time); err != nil {
		return w, err
	}
	if w, err = basictl.StringReadBytes(w, &item.BuildCommit); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.BuildCommitDate); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.BuildCommitTs); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.QueueSizeDisk); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.QueueSizeMemory); err != nil {
		return w, err
	}
	if item.FieldsMask&(1<<2) != 0 {
		if w, err = basictl.IntRead(w, &item.QueueSizeDiskSum); err != nil {
			return w, err
		}
	} else {
		item.QueueSizeDiskSum = 0
	}
	if item.FieldsMask&(1<<2) != 0 {
		if w, err = basictl.IntRead(w, &item.QueueSizeMemorySum); err != nil {
			return w, err
		}
	} else {
		item.QueueSizeMemorySum = 0
	}
	if item.FieldsMask&(1<<3) != 0 {
		if w, err = basictl.IntRead(w, &item.QueueSizeDiskUnsent); err != nil {
			return w, err
		}
	} else {
		item.QueueSizeDiskUnsent = 0
	}
	if item.FieldsMask&(1<<3) != 0 {
		if w, err = basictl.IntRead(w, &item.QueueSizeDiskSumUnsent); err != nil {
			return w, err
		}
	} else {
		item.QueueSizeDiskSumUnsent = 0
	}
	if w, err = basictl.IntRead(w, &item.OriginalSize); err != nil {
		return w, err
	}
	return basictl.StringReadBytes(w, &item.CompressedData)
}

func (item *StatshouseSendSourceBucket2Bytes) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	if w, err = item.Header.Write(w, item.FieldsMask); err != nil {
		return w, err
	}
	w = basictl.NatWrite(w, item.Time)
	if w, err = basictl.StringWriteBytes(w, item.BuildCommit); err != nil {
		return w, err
	}
	w = basictl.IntWrite(w, item.BuildCommitDate)
	w = basictl.IntWrite(w, item.BuildCommitTs)
	w = basictl.IntWrite(w, item.QueueSizeDisk)
	w = basictl.IntWrite(w, item.QueueSizeMemory)
	if item.FieldsMask&(1<<2) != 0 {
		w = basictl.IntWrite(w, item.QueueSizeDiskSum)
	}
	if item.FieldsMask&(1<<2) != 0 {
		w = basictl.IntWrite(w, item.QueueSizeMemorySum)
	}
	if item.FieldsMask&(1<<3) != 0 {
		w = basictl.IntWrite(w, item.QueueSizeDiskUnsent)
	}
	if item.FieldsMask&(1<<3) != 0 {
		w = basictl.IntWrite(w, item.QueueSizeDiskSumUnsent)
	}
	w = basictl.IntWrite(w, item.OriginalSize)
	return basictl.StringWriteBytes(w, item.CompressedData)
}

func (item *StatshouseSendSourceBucket2Bytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x44575940); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseSendSourceBucket2Bytes) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x44575940)
	return item.Write(w)
}

func (item *StatshouseSendSourceBucket2Bytes) ReadResult(w []byte, ret *[]byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xb5286e24); err != nil {
		return w, err
	}
	return basictl.StringReadBytes(w, ret)
}

func (item *StatshouseSendSourceBucket2Bytes) WriteResult(w []byte, ret []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, 0xb5286e24)
	return basictl.StringWriteBytes(w, ret)
}

func (item *StatshouseSendSourceBucket2Bytes) ReadResultJSON(j interface{}, ret *[]byte) error {
	if err := JsonReadStringBytes(j, ret); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseSendSourceBucket2Bytes) WriteResultJSON(w []byte, ret []byte) (_ []byte, err error) {
	return item.writeResultJSON(false, w, ret)
}

func (item *StatshouseSendSourceBucket2Bytes) writeResultJSON(short bool, w []byte, ret []byte) (_ []byte, err error) {
	w = basictl.JSONWriteStringBytes(w, ret)
	return w, nil
}

func (item *StatshouseSendSourceBucket2Bytes) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret []byte
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *StatshouseSendSourceBucket2Bytes) ReadResultWriteResultJSONShort(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret []byte
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(true, w, ret)
	return r, w, err
}

func (item *StatshouseSendSourceBucket2Bytes) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := JsonBytesToInterface(r)
	if err != nil {
		return r, w, ErrorInvalidJSON("statshouse.sendSourceBucket2", err.Error())
	}
	var ret []byte
	if err = item.ReadResultJSON(j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item StatshouseSendSourceBucket2Bytes) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func StatshouseSendSourceBucket2Bytes__ReadJSON(item *StatshouseSendSourceBucket2Bytes, j interface{}) error {
	return item.readJSON(j)
}
func (item *StatshouseSendSourceBucket2Bytes) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouse.sendSourceBucket2", "expected json object")
	}
	_jFieldsMask := _jm["fields_mask"]
	delete(_jm, "fields_mask")
	if err := JsonReadUint32(_jFieldsMask, &item.FieldsMask); err != nil {
		return err
	}
	_jHeader := _jm["header"]
	delete(_jm, "header")
	_jTime := _jm["time"]
	delete(_jm, "time")
	if err := JsonReadUint32(_jTime, &item.Time); err != nil {
		return err
	}
	_jHistoric := _jm["historic"]
	delete(_jm, "historic")
	_jSpare := _jm["spare"]
	delete(_jm, "spare")
	_jBuildCommit := _jm["build_commit"]
	delete(_jm, "build_commit")
	if err := JsonReadStringBytes(_jBuildCommit, &item.BuildCommit); err != nil {
		return err
	}
	_jBuildCommitDate := _jm["build_commit_date"]
	delete(_jm, "build_commit_date")
	if err := JsonReadInt32(_jBuildCommitDate, &item.BuildCommitDate); err != nil {
		return err
	}
	_jBuildCommitTs := _jm["build_commit_ts"]
	delete(_jm, "build_commit_ts")
	if err := JsonReadInt32(_jBuildCommitTs, &item.BuildCommitTs); err != nil {
		return err
	}
	_jQueueSizeDisk := _jm["queue_size_disk"]
	delete(_jm, "queue_size_disk")
	if err := JsonReadInt32(_jQueueSizeDisk, &item.QueueSizeDisk); err != nil {
		return err
	}
	_jQueueSizeMemory := _jm["queue_size_memory"]
	delete(_jm, "queue_size_memory")
	if err := JsonReadInt32(_jQueueSizeMemory, &item.QueueSizeMemory); err != nil {
		return err
	}
	_jQueueSizeDiskSum := _jm["queue_size_disk_sum"]
	delete(_jm, "queue_size_disk_sum")
	_jQueueSizeMemorySum := _jm["queue_size_memory_sum"]
	delete(_jm, "queue_size_memory_sum")
	_jQueueSizeDiskUnsent := _jm["queue_size_disk_unsent"]
	delete(_jm, "queue_size_disk_unsent")
	_jQueueSizeDiskSumUnsent := _jm["queue_size_disk_sum_unsent"]
	delete(_jm, "queue_size_disk_sum_unsent")
	_jOriginalSize := _jm["original_size"]
	delete(_jm, "original_size")
	if err := JsonReadInt32(_jOriginalSize, &item.OriginalSize); err != nil {
		return err
	}
	_jCompressedData := _jm["compressed_data"]
	delete(_jm, "compressed_data")
	if err := JsonReadStringBytes(_jCompressedData, &item.CompressedData); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouse.sendSourceBucket2", k)
	}
	if _jHistoric != nil {
		_bit := false
		if err := JsonReadBool(_jHistoric, &_bit); err != nil {
			return err
		}
		if _bit {
			item.FieldsMask |= 1 << 0
		} else {
			item.FieldsMask &^= 1 << 0
		}
	}
	if _jSpare != nil {
		_bit := false
		if err := JsonReadBool(_jSpare, &_bit); err != nil {
			return err
		}
		if _bit {
			item.FieldsMask |= 1 << 1
		} else {
			item.FieldsMask &^= 1 << 1
		}
	}
	if _jQueueSizeDiskSum != nil {
		item.FieldsMask |= 1 << 2
	}
	if _jQueueSizeMemorySum != nil {
		item.FieldsMask |= 1 << 2
	}
	if _jQueueSizeDiskUnsent != nil {
		item.FieldsMask |= 1 << 3
	}
	if _jQueueSizeDiskSumUnsent != nil {
		item.FieldsMask |= 1 << 3
	}
	if err := StatshouseCommonProxyHeaderBytes__ReadJSON(&item.Header, _jHeader, item.FieldsMask); err != nil {
		return err
	}
	if _jQueueSizeDiskSum != nil {
		if err := JsonReadInt32(_jQueueSizeDiskSum, &item.QueueSizeDiskSum); err != nil {
			return err
		}
	} else {
		item.QueueSizeDiskSum = 0
	}
	if _jQueueSizeMemorySum != nil {
		if err := JsonReadInt32(_jQueueSizeMemorySum, &item.QueueSizeMemorySum); err != nil {
			return err
		}
	} else {
		item.QueueSizeMemorySum = 0
	}
	if _jQueueSizeDiskUnsent != nil {
		if err := JsonReadInt32(_jQueueSizeDiskUnsent, &item.QueueSizeDiskUnsent); err != nil {
			return err
		}
	} else {
		item.QueueSizeDiskUnsent = 0
	}
	if _jQueueSizeDiskSumUnsent != nil {
		if err := JsonReadInt32(_jQueueSizeDiskSumUnsent, &item.QueueSizeDiskSumUnsent); err != nil {
			return err
		}
	} else {
		item.QueueSizeDiskSumUnsent = 0
	}
	return nil
}

func (item *StatshouseSendSourceBucket2Bytes) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *StatshouseSendSourceBucket2Bytes) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldsMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"fields_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldsMask)
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"header":`...)
	if w, err = item.Header.WriteJSONOpt(short, w, item.FieldsMask); err != nil {
		return w, err
	}
	if item.Time != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"time":`...)
		w = basictl.JSONWriteUint32(w, item.Time)
	}
	if item.FieldsMask&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"historic":true`...)
	}
	if item.FieldsMask&(1<<1) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"spare":true`...)
	}
	if len(item.BuildCommit) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"build_commit":`...)
		w = basictl.JSONWriteStringBytes(w, item.BuildCommit)
	}
	if item.BuildCommitDate != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"build_commit_date":`...)
		w = basictl.JSONWriteInt32(w, item.BuildCommitDate)
	}
	if item.BuildCommitTs != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"build_commit_ts":`...)
		w = basictl.JSONWriteInt32(w, item.BuildCommitTs)
	}
	if item.QueueSizeDisk != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"queue_size_disk":`...)
		w = basictl.JSONWriteInt32(w, item.QueueSizeDisk)
	}
	if item.QueueSizeMemory != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"queue_size_memory":`...)
		w = basictl.JSONWriteInt32(w, item.QueueSizeMemory)
	}
	if item.FieldsMask&(1<<2) != 0 {
		if item.QueueSizeDiskSum != 0 {
			w = basictl.JSONAddCommaIfNeeded(w)
			w = append(w, `"queue_size_disk_sum":`...)
			w = basictl.JSONWriteInt32(w, item.QueueSizeDiskSum)
		}
	}
	if item.FieldsMask&(1<<2) != 0 {
		if item.QueueSizeMemorySum != 0 {
			w = basictl.JSONAddCommaIfNeeded(w)
			w = append(w, `"queue_size_memory_sum":`...)
			w = basictl.JSONWriteInt32(w, item.QueueSizeMemorySum)
		}
	}
	if item.FieldsMask&(1<<3) != 0 {
		if item.QueueSizeDiskUnsent != 0 {
			w = basictl.JSONAddCommaIfNeeded(w)
			w = append(w, `"queue_size_disk_unsent":`...)
			w = basictl.JSONWriteInt32(w, item.QueueSizeDiskUnsent)
		}
	}
	if item.FieldsMask&(1<<3) != 0 {
		if item.QueueSizeDiskSumUnsent != 0 {
			w = basictl.JSONAddCommaIfNeeded(w)
			w = append(w, `"queue_size_disk_sum_unsent":`...)
			w = basictl.JSONWriteInt32(w, item.QueueSizeDiskSumUnsent)
		}
	}
	if item.OriginalSize != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"original_size":`...)
		w = basictl.JSONWriteInt32(w, item.OriginalSize)
	}
	if len(item.CompressedData) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"compressed_data":`...)
		w = basictl.JSONWriteStringBytes(w, item.CompressedData)
	}
	return append(w, '}'), nil
}

func (item *StatshouseSendSourceBucket2Bytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseSendSourceBucket2Bytes) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("statshouse.sendSourceBucket2", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("statshouse.sendSourceBucket2", err.Error())
	}
	return nil
}
