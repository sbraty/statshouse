// Copyright 2023 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package factory

import (
	"github.com/vkcom/statshouse/internal/data_model/gen2/internal"
	"github.com/vkcom/statshouse/internal/data_model/gen2/meta"
)

func CreateFunction(tag uint32) meta.Function {
	return meta.CreateFunction(tag)
}

func CreateObject(tag uint32) meta.Object {
	return meta.CreateObject(tag)
}

// name can be in any of 3 forms "ch_proxy.insert#7cf362ba", "ch_proxy.insert" or "#7cf362ba"
func CreateFunctionFromName(name string) meta.Function {
	return meta.CreateFunctionFromName(name)
}

// name can be in any of 3 forms "ch_proxy.insert#7cf362ba", "ch_proxy.insert" or "#7cf362ba"
func CreateObjectFromName(name string) meta.Object {
	return meta.CreateObjectFromName(name)
}

func init() {
	meta.SetGlobalFactoryCreateForObject(0x92cbcbfa, func() meta.Object { var ret internal.BoolStat; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x402409cb, func() meta.Object { var ret internal.EngineAlreadyInMasterMode; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xebd80142, func() meta.Object { var ret internal.EngineAlreadyInReplicaMode; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0x60e50d3d, func() meta.Object { var ret internal.EngineAsyncSleep; return &ret }, func() meta.Function { var ret internal.EngineAsyncSleep; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0x4c09c894, func() meta.Object { var ret internal.EngineBinlogPrefix; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0x19d0f020, func() meta.Object { var ret internal.EngineCount; return &ret }, func() meta.Function { var ret internal.EngineCount; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0xf1f90880, func() meta.Object { var ret internal.EngineDumpForceQueries; return &ret }, func() meta.Function { var ret internal.EngineDumpForceQueries; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0xc060a29f, func() meta.Object { var ret internal.EngineDumpLastQueries; return &ret }, func() meta.Function { var ret internal.EngineDumpLastQueries; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0xe65872ad, func() meta.Object { var ret internal.EngineDumpNextQueries; return &ret }, func() meta.Function { var ret internal.EngineDumpNextQueries; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x88836cdc, func() meta.Object { var ret internal.EngineEnableMetafilesAnalyzer; return &ret }, func() meta.Function { var ret internal.EngineEnableMetafilesAnalyzer; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0x17418662, func() meta.Object { var ret internal.EngineFailedToSwitchMode; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0x594870d6, func() meta.Object { var ret internal.EngineFilteredStat; return &ret }, func() meta.Function { var ret internal.EngineFilteredStat; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0xef14db93, func() meta.Object { var ret internal.EngineGetBinlogPrefixes; return &ret }, func() meta.Function { var ret internal.EngineGetBinlogPrefixes; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x0342f391, func() meta.Object { var ret internal.EngineGetExpectedMetafilesStats; return &ret }, func() meta.Function { var ret internal.EngineGetExpectedMetafilesStats; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x61b3f593, func() meta.Object { var ret internal.EngineGetReadWriteMode; return &ret }, func() meta.Function { var ret internal.EngineGetReadWriteMode; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0xf492042e, func() meta.Object { var ret internal.EngineGetReindexStatus; return &ret }, func() meta.Function { var ret internal.EngineGetReindexStatus; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0x58300321, func() meta.Object { var ret internal.EngineHttpQuery; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x284852fc, func() meta.Object { var ret internal.EngineHttpQueryResponse; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0xf4c73c0b, func() meta.Object { var ret internal.EngineInvokeHttpQuery; return &ret }, func() meta.Function { var ret internal.EngineInvokeHttpQuery; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0xccdea0ac, func() meta.Object { var ret internal.EngineIsProduction; return &ret }, func() meta.Function { var ret internal.EngineIsProduction; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0xc292e4a6, func() meta.Object { var ret internal.EngineMetafilesOneMemoryStat; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xb673669b, func() meta.Object { var ret internal.EngineMetafilesStat; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0x166bb7c6, func() meta.Object { var ret internal.EngineNop; return &ret }, func() meta.Function { var ret internal.EngineNop; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x559d6e36, func() meta.Object { var ret internal.EnginePid; return &ret }, func() meta.Function { var ret internal.EnginePid; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0xf4b19fa2, func() meta.Object { var ret internal.EnginePushStat; return &ret }, func() meta.Function { var ret internal.EnginePushStat; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0xac4d6fe9, func() meta.Object { var ret internal.EngineQueryResult0; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xee2879b0, func() meta.Object { var ret internal.EngineQueryResultAio; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x2b4dd0ba, func() meta.Object { var ret internal.EngineQueryResultError; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0x9d2b841f, func() meta.Object { var ret internal.EngineReadNop; return &ret }, func() meta.Function { var ret internal.EngineReadNop; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x0001e9d6, func() meta.Object { var ret internal.EngineRecordNextQueries; return &ret }, func() meta.Function { var ret internal.EngineRecordNextQueries; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x2f86f276, func() meta.Object { var ret internal.EngineRegisterDynamicLib; return &ret }, func() meta.Function { var ret internal.EngineRegisterDynamicLib; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0x0f67569a, func() meta.Object { var ret internal.EngineReindexStatusDone; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xafdbd505, func() meta.Object { var ret internal.EngineReindexStatusDoneOld; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x10533721, func() meta.Object { var ret internal.EngineReindexStatusFailed; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x7f6a89b9, func() meta.Object { var ret internal.EngineReindexStatusNever; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xfa198b59, func() meta.Object { var ret internal.EngineReindexStatusRunning; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xac530b46, func() meta.Object { var ret internal.EngineReindexStatusRunningOld; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x756e878b, func() meta.Object { var ret internal.EngineReindexStatusSignaled; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0x602d62c1, func() meta.Object { var ret internal.EngineReloadDynamicLib; return &ret }, func() meta.Function { var ret internal.EngineReloadDynamicLib; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0x0f3d0fb1, func() meta.Object { var ret internal.EngineReloadDynamicLibOptions; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0x5fcd8e77, func() meta.Object { var ret internal.EngineReplaceConfigServer; return &ret }, func() meta.Function { var ret internal.EngineReplaceConfigServer; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x1a7708a3, func() meta.Object { var ret internal.EngineSendSignal; return &ret }, func() meta.Function { var ret internal.EngineSendSignal; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x665d2ab7, func() meta.Object { var ret internal.EngineSetFsyncInterval; return &ret }, func() meta.Function { var ret internal.EngineSetFsyncInterval; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x7bdcf404, func() meta.Object { var ret internal.EngineSetMetafileMemory; return &ret }, func() meta.Function { var ret internal.EngineSetMetafileMemory; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x5806a520, func() meta.Object { var ret internal.EngineSetNoPersistentConfigArray; return &ret }, func() meta.Function { var ret internal.EngineSetNoPersistentConfigArray; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x92aaa5b9, func() meta.Object { var ret internal.EngineSetNoPersistentConfigValue; return &ret }, func() meta.Function { var ret internal.EngineSetNoPersistentConfigValue; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0xfc99af0b, func() meta.Object { var ret internal.EngineSetPersistentConfigArray; return &ret }, func() meta.Function { var ret internal.EngineSetPersistentConfigArray; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x4cc8953f, func() meta.Object { var ret internal.EngineSetPersistentConfigValue; return &ret }, func() meta.Function { var ret internal.EngineSetPersistentConfigValue; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x9d980926, func() meta.Object { var ret internal.EngineSetVerbosity; return &ret }, func() meta.Function { var ret internal.EngineSetVerbosity; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x5388c0ae, func() meta.Object { var ret internal.EngineSetVerbosityType; return &ret }, func() meta.Function { var ret internal.EngineSetVerbosityType; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x3d3bcd48, func() meta.Object { var ret internal.EngineSleep; return &ret }, func() meta.Function { var ret internal.EngineSleep; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0xefb3c36b, func() meta.Object { var ret internal.EngineStat; return &ret }, func() meta.Function { var ret internal.EngineStat; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x8cdcb5f9, func() meta.Object { var ret internal.EngineSwitchToMasterMode; return &ret }, func() meta.Function { var ret internal.EngineSwitchToMasterMode; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x1973fa8f, func() meta.Object { var ret internal.EngineSwitchToMasterModeForcefully; return &ret }, func() meta.Function { var ret internal.EngineSwitchToMasterModeForcefully; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x23c3a87e, func() meta.Object { var ret internal.EngineSwitchToReplicaMode; return &ret }, func() meta.Function { var ret internal.EngineSwitchToReplicaMode; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0x95b13964, func() meta.Object { var ret internal.EngineSwitchedToMasterMode; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xec61b4be, func() meta.Object { var ret internal.EngineSwitchedToMasterModeForcefully; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xad642a0b, func() meta.Object { var ret internal.EngineSwitchedToReplicaMode; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0x84d5fcb9, func() meta.Object { var ret internal.EngineUnregisterDynamicLib; return &ret }, func() meta.Function { var ret internal.EngineUnregisterDynamicLib; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x1a2e06fa, func() meta.Object { var ret internal.EngineVersion; return &ret }, func() meta.Function { var ret internal.EngineVersion; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x58160af4, func() meta.Object { var ret internal.EngineWriteNop; return &ret }, func() meta.Function { var ret internal.EngineWriteNop; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0xa8509bda, func() meta.Object { var ret internal.Int; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x1a345674, func() meta.Object { var ret internal.MetadataCreateEntityEvent; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x12345678, func() meta.Object { var ret internal.MetadataCreateMappingEvent; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x12345674, func() meta.Object { var ret internal.MetadataCreateMetricEvent; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x1234b677, func() meta.Object { var ret internal.MetadataEditEntityEvent; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0x86df475f, func() meta.Object { var ret internal.MetadataEditEntitynew; return &ret }, func() meta.Function { var ret internal.MetadataEditEntitynew; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0x12345677, func() meta.Object { var ret internal.MetadataEditMetricEvent; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x9286affa, func() meta.Object { var ret internal.MetadataEvent; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0x72b132f8, func() meta.Object { var ret internal.MetadataGetEntity; return &ret }, func() meta.Function { var ret internal.MetadataGetEntity; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x22ff6a79, func() meta.Object { var ret internal.MetadataGetHistoryShortInfo; return &ret }, func() meta.Function { var ret internal.MetadataGetHistoryShortInfo; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x9faf5280, func() meta.Object { var ret internal.MetadataGetInvertMapping; return &ret }, func() meta.Function { var ret internal.MetadataGetInvertMapping; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x93ba92f8, func() meta.Object { var ret internal.MetadataGetJournalnew; return &ret }, func() meta.Function { var ret internal.MetadataGetJournalnew; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x9dfa7a83, func() meta.Object { var ret internal.MetadataGetMapping; return &ret }, func() meta.Function { var ret internal.MetadataGetMapping; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x93ba92f5, func() meta.Object { var ret internal.MetadataGetMetrics; return &ret }, func() meta.Function { var ret internal.MetadataGetMetrics; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x5fc81a9b, func() meta.Object { var ret internal.MetadataGetTagMappingBootstrap; return &ret }, func() meta.Function { var ret internal.MetadataGetTagMappingBootstrap; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0x5854dfaf, func() meta.Object { var ret internal.MetadataPutBootstrapEvent; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0x9faf5281, func() meta.Object { var ret internal.MetadataPutMapping; return &ret }, func() meta.Function { var ret internal.MetadataPutMapping; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0x12345676, func() meta.Object { var ret internal.MetadataPutMappingEvent; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x9286abfe, func() meta.Object { var ret internal.MetadataPutMappingResponse; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0x5fc8ab9b, func() meta.Object { var ret internal.MetadataPutTagMappingBootstrap; return &ret }, func() meta.Function { var ret internal.MetadataPutTagMappingBootstrap; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x9faf5282, func() meta.Object { var ret internal.MetadataResetFlood; return &ret }, func() meta.Function { var ret internal.MetadataResetFlood; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x88d0fd5e, func() meta.Object { var ret internal.MetadataResetFlood2; return &ret }, func() meta.Function { var ret internal.MetadataResetFlood2; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0x9286abee, func() meta.Object { var ret internal.MetadataResetFloodResponse; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x9286abef, func() meta.Object { var ret internal.MetadataResetFloodResponse2; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x46409ccf, func() meta.Object { var ret internal.NetPid; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x9d56e6b2, func() meta.Object { var ret internal.Stat; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0x56580239, func() meta.Object { var ret internal.StatshouseAddMetricsBatch; return &ret }, func() meta.Function { var ret internal.StatshouseAddMetricsBatch; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0x511276a6, func() meta.Object { var ret internal.StatshouseApiFilter; return &ret })
	meta.SetGlobalFactoryCreateForEnumElement(0x2a6e4c14)
	meta.SetGlobalFactoryCreateForEnumElement(0x670ab89c)
	meta.SetGlobalFactoryCreateForEnumElement(0x4ca979c0)
	meta.SetGlobalFactoryCreateForEnumElement(0x6323c2f6)
	meta.SetGlobalFactoryCreateForEnumElement(0x89689775)
	meta.SetGlobalFactoryCreateForEnumElement(0x60e68b5c)
	meta.SetGlobalFactoryCreateForEnumElement(0xf4d9ad09)
	meta.SetGlobalFactoryCreateForEnumElement(0x871201c4)
	meta.SetGlobalFactoryCreateForEnumElement(0x42fc39b6)
	meta.SetGlobalFactoryCreateForEnumElement(0x60d2b603)
	meta.SetGlobalFactoryCreateForEnumElement(0xe617771c)
	meta.SetGlobalFactoryCreateForEnumElement(0xbfb5f7fc)
	meta.SetGlobalFactoryCreateForEnumElement(0x43eeb810)
	meta.SetGlobalFactoryCreateForEnumElement(0x4817df2b)
	meta.SetGlobalFactoryCreateForEnumElement(0xa3a43781)
	meta.SetGlobalFactoryCreateForEnumElement(0x96683390)
	meta.SetGlobalFactoryCreateForEnumElement(0x5745a0a3)
	meta.SetGlobalFactoryCreateForEnumElement(0x4bd4f327)
	meta.SetGlobalFactoryCreateForEnumElement(0xf90de384)
	meta.SetGlobalFactoryCreateForEnumElement(0x885e665b)
	meta.SetGlobalFactoryCreateForEnumElement(0xb4790064)
	meta.SetGlobalFactoryCreateForEnumElement(0xb4cb2644)
	meta.SetGlobalFactoryCreateForEnumElement(0x381b1cee)
	meta.SetGlobalFactoryCreateForEnumElement(0xbbb36a23)
	meta.SetGlobalFactoryCreateForEnumElement(0x56071d39)
	meta.SetGlobalFactoryCreateForEnumElement(0xcf9ad7bf)
	meta.SetGlobalFactoryCreateForEnumElement(0xbcdeae3a)
	meta.SetGlobalFactoryCreateForEnumElement(0x77c5de5c)
	meta.SetGlobalFactoryCreateForEnumElement(0x0e674272)
	meta.SetGlobalFactoryCreateForEnumElement(0xd4c8c793)
	meta.SetGlobalFactoryCreateForEnumElement(0x9a92b76f)
	meta.SetGlobalFactoryCreateForEnumElement(0x71992e9a)
	meta.SetGlobalFactoryCreateForEnumElement(0xa3434c26)
	meta.SetGlobalFactoryCreateForEnumElement(0x2043e480)
	meta.SetGlobalFactoryCreateForEnumElement(0x80ce3cf1)
	meta.SetGlobalFactoryCreateForEnumElement(0x361963d5)
	meta.SetGlobalFactoryCreateForEnumElement(0xf20fb854)
	meta.SetGlobalFactoryCreateForEnumElement(0x9ceb6f68)
	meta.SetGlobalFactoryCreateForFunction(0x52721884, func() meta.Object { var ret internal.StatshouseApiGetChunk; return &ret }, func() meta.Function { var ret internal.StatshouseApiGetChunk; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0x63928b42, func() meta.Object { var ret internal.StatshouseApiGetChunkResponse; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0x0c7349bb, func() meta.Object { var ret internal.StatshouseApiGetQuery; return &ret }, func() meta.Function { var ret internal.StatshouseApiGetQuery; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x0c7348bb, func() meta.Object { var ret internal.StatshouseApiGetQueryPoint; return &ret }, func() meta.Function { var ret internal.StatshouseApiGetQueryPoint; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0x4487e41a, func() meta.Object { var ret internal.StatshouseApiGetQueryPointResponse; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x5c2bf296, func() meta.Object { var ret internal.StatshouseApiPointMeta; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xc9951bb9, func() meta.Object { var ret internal.StatshouseApiQuery; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xc9951bbb, func() meta.Object { var ret internal.StatshouseApiQueryPoint; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0x62adc773, func() meta.Object { var ret internal.StatshouseApiReleaseChunks; return &ret }, func() meta.Function { var ret internal.StatshouseApiReleaseChunks; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0xd12dc2bd, func() meta.Object { var ret internal.StatshouseApiReleaseChunksResponse; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x07a3e919, func() meta.Object { var ret internal.StatshouseApiSeries; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x43eeb763, func() meta.Object { var ret internal.StatshouseApiTagValue; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0x28bea524, func() meta.Object { var ret internal.StatshouseAutoCreate; return &ret }, func() meta.Function { var ret internal.StatshouseAutoCreate; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0x73fd01e0, func() meta.Object { var ret internal.StatshouseCentroid; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0x4285ff57, func() meta.Object { var ret internal.StatshouseGetConfig2; return &ret }, func() meta.Function { var ret internal.StatshouseGetConfig2; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x42855554, func() meta.Object { var ret internal.StatshouseGetMetrics3; return &ret }, func() meta.Function { var ret internal.StatshouseGetMetrics3; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0x0c803d05, func() meta.Object { var ret internal.StatshouseGetMetricsResult; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0x4285ff56, func() meta.Object { var ret internal.StatshouseGetTagMapping2; return &ret }, func() meta.Function { var ret internal.StatshouseGetTagMapping2; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x75a7f68e, func() meta.Object { var ret internal.StatshouseGetTagMappingBootstrap; return &ret }, func() meta.Function { var ret internal.StatshouseGetTagMappingBootstrap; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0x486a40de, func() meta.Object { var ret internal.StatshouseGetTagMappingBootstrapResult; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x1a7d91fd, func() meta.Object { var ret internal.StatshouseGetTagMappingResult; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0x41df72a3, func() meta.Object { var ret internal.StatshouseGetTargets2; return &ret }, func() meta.Function { var ret internal.StatshouseGetTargets2; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0x2e17a6d3, func() meta.Object { var ret internal.StatshouseIngestionStatus2; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xbf401d4b, func() meta.Object { var ret internal.StatshouseMapping; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x3325d884, func() meta.Object { var ret internal.StatshouseMetric; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x0c803e07, func() meta.Object { var ret internal.StatshouseMultiItem; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x486affde, func() meta.Object { var ret internal.StatshousePutTagMappingBootstrapResult; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x4f7b7822, func() meta.Object { var ret internal.StatshouseSampleFactor; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0x4285ff53, func() meta.Object { var ret internal.StatshouseSendKeepAlive2; return &ret }, func() meta.Function { var ret internal.StatshouseSendKeepAlive2; return &ret }, nil)
	meta.SetGlobalFactoryCreateForFunction(0x44575940, func() meta.Object { var ret internal.StatshouseSendSourceBucket2; return &ret }, func() meta.Function { var ret internal.StatshouseSendSourceBucket2; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0x3af6e822, func() meta.Object { var ret internal.StatshouseSourceBucket2; return &ret })
	meta.SetGlobalFactoryCreateForFunction(0x4285ff58, func() meta.Object { var ret internal.StatshouseTestConnection2; return &ret }, func() meta.Function { var ret internal.StatshouseTestConnection2; return &ret }, nil)
	meta.SetGlobalFactoryCreateForObject(0x9ffdea42, func() meta.Object { var ret internal.StatshouseTopElement; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xb5286e24, func() meta.Object { var ret internal.String; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x3fedd339, func() meta.Object { var ret internal.True; return &ret })
}
