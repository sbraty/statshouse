// Copyright 2023 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tl

import (
	"github.com/vkcom/statshouse/internal/data_model/gen2/internal"
)

type (
	BoolStat                                      = internal.BoolStat
	DictionaryEngineMetafilesStatBoxed            = internal.DictionaryEngineMetafilesStatBoxed
	DictionaryFieldEngineMetafilesStatBoxed       = internal.DictionaryFieldEngineMetafilesStatBoxed
	DictionaryFieldString                         = internal.DictionaryFieldString
	DictionaryFieldStringBytes                    = internal.DictionaryFieldStringBytes
	DictionaryString                              = internal.DictionaryString
	DictionaryStringBytes                         = internal.DictionaryStringBytes
	Int                                           = internal.Int
	Stat                                          = internal.Stat
	String                                        = internal.String
	StringBytes                                   = internal.StringBytes
	True                                          = internal.True
	VectorDictionaryFieldEngineMetafilesStatBoxed = internal.VectorDictionaryFieldEngineMetafilesStatBoxed
	VectorDictionaryFieldString                   = internal.VectorDictionaryFieldString
	VectorDictionaryFieldStringBytes              = internal.VectorDictionaryFieldStringBytes
	VectorDouble                                  = internal.VectorDouble
	VectorEngineBinlogPrefix                      = internal.VectorEngineBinlogPrefix
	VectorEngineMetafilesOneMemoryStat            = internal.VectorEngineMetafilesOneMemoryStat
	VectorInt                                     = internal.VectorInt
	VectorKvEngineKvBoxed                         = internal.VectorKvEngineKvBoxed
	VectorLong                                    = internal.VectorLong
	VectorMetadataEvent                           = internal.VectorMetadataEvent
	VectorMetadataEventBytes                      = internal.VectorMetadataEventBytes
	VectorMetadataHistoryShortResponseEvent       = internal.VectorMetadataHistoryShortResponseEvent
	VectorMetadataMetricOld                       = internal.VectorMetadataMetricOld
	VectorStatshouseApiFilter                     = internal.VectorStatshouseApiFilter
	VectorStatshouseApiFunction                   = internal.VectorStatshouseApiFunction
	VectorStatshouseApiPointMeta                  = internal.VectorStatshouseApiPointMeta
	VectorStatshouseApiSeriesMeta                 = internal.VectorStatshouseApiSeriesMeta
	VectorStatshouseApiTagValue                   = internal.VectorStatshouseApiTagValue
	VectorStatshouseCentroid                      = internal.VectorStatshouseCentroid
	VectorStatshouseIngestionStatus2              = internal.VectorStatshouseIngestionStatus2
	VectorStatshouseMapping                       = internal.VectorStatshouseMapping
	VectorStatshouseMappingBytes                  = internal.VectorStatshouseMappingBytes
	VectorStatshouseMetric                        = internal.VectorStatshouseMetric
	VectorStatshouseMetricBytes                   = internal.VectorStatshouseMetricBytes
	VectorStatshouseMultiItem                     = internal.VectorStatshouseMultiItem
	VectorStatshouseMultiItemBytes                = internal.VectorStatshouseMultiItemBytes
	VectorStatshousePromTarget                    = internal.VectorStatshousePromTarget
	VectorStatshousePromTargetBytes               = internal.VectorStatshousePromTargetBytes
	VectorStatshouseSampleFactor                  = internal.VectorStatshouseSampleFactor
	VectorStatshouseTopElement                    = internal.VectorStatshouseTopElement
	VectorStatshouseTopElementBytes               = internal.VectorStatshouseTopElementBytes
	VectorString                                  = internal.VectorString
	VectorStringBytes                             = internal.VectorStringBytes
	VectorVectorDouble                            = internal.VectorVectorDouble
)

func BoolReadBoxed(w []byte, v *bool) ([]byte, error) {
	return internal.BoolReadBoxed(w, v)
}
func BoolWriteBoxed(w []byte, v bool) []byte {
	return internal.BoolWriteBoxed(w, v)
}
