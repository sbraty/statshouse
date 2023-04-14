// Copyright 2023 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package constants

const (
	BoolFalse                                    = 0xbc799737 // boolFalse
	BoolStat                                     = 0x92cbcbfa // boolStat
	BoolTrue                                     = 0x997275b5 // boolTrue
	Dictionary                                   = 0x1f4c618f // dictionary
	DictionaryField                              = 0x239c1b62 // dictionaryField
	EngineAlreadyInMasterMode                    = 0x402409cb // engine.alreadyInMasterMode
	EngineAlreadyInReplicaMode                   = 0xebd80142 // engine.alreadyInReplicaMode
	EngineAsyncSleep                             = 0x60e50d3d // engine.asyncSleep
	EngineBinlogPrefix                           = 0x4c09c894 // engine.binlogPrefix
	EngineCount                                  = 0x19d0f020 // engine.count
	EngineDumpForceQueries                       = 0xf1f90880 // engine.dumpForceQueries
	EngineDumpLastQueries                        = 0xc060a29f // engine.dumpLastQueries
	EngineDumpNextQueries                        = 0xe65872ad // engine.dumpNextQueries
	EngineEnableMetafilesAnalyzer                = 0x88836cdc // engine.enableMetafilesAnalyzer
	EngineFailedToSwitchMode                     = 0x17418662 // engine.failedToSwitchMode
	EngineFilteredStat                           = 0x594870d6 // engine.filteredStat
	EngineGetBinlogPrefixes                      = 0xef14db93 // engine.getBinlogPrefixes
	EngineGetExpectedMetafilesStats              = 0x0342f391 // engine.getExpectedMetafilesStats
	EngineGetReadWriteMode                       = 0x61b3f593 // engine.getReadWriteMode
	EngineGetReindexStatus                       = 0xf492042e // engine.getReindexStatus
	EngineHttpQuery                              = 0x58300321 // engine.httpQuery
	EngineHttpQueryResponse                      = 0x284852fc // engine.httpQueryResponse
	EngineInvokeHttpQuery                        = 0xf4c73c0b // engine.invokeHttpQuery
	EngineIsProduction                           = 0xccdea0ac // engine.isProduction
	EngineMetafilesOneMemoryStat                 = 0xc292e4a6 // engine.metafilesOneMemoryStat
	EngineMetafilesStatData                      = 0xb673669b // engine.metafilesStatData
	EngineMode                                   = 0xb9b7b6c9 // engine.mode
	EngineNop                                    = 0x166bb7c6 // engine.nop
	EnginePid                                    = 0x559d6e36 // engine.pid
	EnginePushStat                               = 0xf4b19fa2 // engine.pushStat
	EngineQueryResult                            = 0xac4d6fe9 // engine.queryResult
	EngineQueryResultAio                         = 0xee2879b0 // engine.queryResultAio
	EngineQueryResultError                       = 0x2b4dd0ba // engine.queryResultError
	EngineReadNop                                = 0x9d2b841f // engine.readNop
	EngineRecordNextQueries                      = 0x0001e9d6 // engine.recordNextQueries
	EngineRegisterDynamicLib                     = 0x2f86f276 // engine.registerDynamicLib
	EngineReindexStatusDone                      = 0x0f67569a // engine.reindexStatusDone
	EngineReindexStatusDoneOld                   = 0xafdbd505 // engine.reindexStatusDoneOld
	EngineReindexStatusFailed                    = 0x10533721 // engine.reindexStatusFailed
	EngineReindexStatusNever                     = 0x7f6a89b9 // engine.reindexStatusNever
	EngineReindexStatusRunning                   = 0xfa198b59 // engine.reindexStatusRunning
	EngineReindexStatusRunningOld                = 0xac530b46 // engine.reindexStatusRunningOld
	EngineReindexStatusSignaled                  = 0x756e878b // engine.reindexStatusSignaled
	EngineReloadDynamicLib                       = 0x602d62c1 // engine.reloadDynamicLib
	EngineReloadDynamicLibOptions                = 0x0f3d0fb1 // engine.reloadDynamicLibOptions
	EngineReplaceConfigServer                    = 0x5fcd8e77 // engine.replaceConfigServer
	EngineSendResponseTo                         = 0x31e9061d // engine.sendResponseTo
	EngineSendSignal                             = 0x1a7708a3 // engine.sendSignal
	EngineSetFsyncInterval                       = 0x665d2ab7 // engine.setFsyncInterval
	EngineSetMetafileMemory                      = 0x7bdcf404 // engine.setMetafileMemory
	EngineSetNoPersistentConfigArray             = 0x5806a520 // engine.setNoPersistentConfigArray
	EngineSetNoPersistentConfigValue             = 0x92aaa5b9 // engine.setNoPersistentConfigValue
	EngineSetPersistentConfigArray               = 0xfc99af0b // engine.setPersistentConfigArray
	EngineSetPersistentConfigValue               = 0x4cc8953f // engine.setPersistentConfigValue
	EngineSetVerbosity                           = 0x9d980926 // engine.setVerbosity
	EngineSetVerbosityType                       = 0x5388c0ae // engine.setVerbosityType
	EngineSleep                                  = 0x3d3bcd48 // engine.sleep
	EngineStat                                   = 0xefb3c36b // engine.stat
	EngineSwitchToMasterMode                     = 0x8cdcb5f9 // engine.switchToMasterMode
	EngineSwitchToMasterModeForcefully           = 0x1973fa8f // engine.switchToMasterModeForcefully
	EngineSwitchToReplicaMode                    = 0x23c3a87e // engine.switchToReplicaMode
	EngineSwitchedToMasterMode                   = 0x95b13964 // engine.switchedToMasterMode
	EngineSwitchedToMasterModeForcefully         = 0xec61b4be // engine.switchedToMasterModeForcefully
	EngineSwitchedToReplicaMode                  = 0xad642a0b // engine.switchedToReplicaMode
	EngineUnregisterDynamicLib                   = 0x84d5fcb9 // engine.unregisterDynamicLib
	EngineVersion                                = 0x1a2e06fa // engine.version
	EngineWriteNop                               = 0x58160af4 // engine.writeNop
	Map                                          = 0x79c473a4 // map
	MetadataCreateEntityEvent                    = 0x1a345674 // metadata.createEntityEvent
	MetadataCreateMappingEvent                   = 0x12345678 // metadata.createMappingEvent
	MetadataCreateMetricEvent                    = 0x12345674 // metadata.createMetricEvent
	MetadataEditEntityEvent                      = 0x1234b677 // metadata.editEntityEvent
	MetadataEditEntitynew                        = 0x86df475f // metadata.editEntitynew
	MetadataEditMetricEvent                      = 0x12345677 // metadata.editMetricEvent
	MetadataEvent                                = 0x9286affa // metadata.event
	MetadataGetInvertMapping                     = 0x9faf5280 // metadata.getInvertMapping
	MetadataGetInvertMappingResponse             = 0x9286abac // metadata.getInvertMappingResponse
	MetadataGetInvertMappingResponseKeyNotExists = 0x9286abab // metadata.getInvertMappingResponseKeyNotExists
	MetadataGetJournalResponsenew                = 0x9286aaaa // metadata.getJournalResponsenew
	MetadataGetJournalnew                        = 0x93ba92f8 // metadata.getJournalnew
	MetadataGetMapping                           = 0x9dfa7a83 // metadata.getMapping
	MetadataGetMappingResponse                   = 0x9286abfc // metadata.getMappingResponse
	MetadataGetMappingResponseCreated            = 0x9286abbb // metadata.getMappingResponseCreated
	MetadataGetMappingResponseFloodLimitError    = 0x9286abfd // metadata.getMappingResponseFloodLimitError
	MetadataGetMappingResponseKeyNotExists       = 0x9286abff // metadata.getMappingResponseKeyNotExists
	MetadataGetMetrics                           = 0x93ba92f5 // metadata.getMetrics
	MetadataGetMetricsResponse                   = 0x9286abfb // metadata.getMetricsResponse
	MetadataGetTagMappingBootstrap               = 0x5fc81a9b // metadata.getTagMappingBootstrap
	MetadataMetricOld                            = 0x9286abfa // metadata.metricOld
	MetadataPutBootstrapEvent                    = 0x5854dfaf // metadata.putBootstrapEvent
	MetadataPutMapping                           = 0x9faf5281 // metadata.putMapping
	MetadataPutMappingEvent                      = 0x12345676 // metadata.putMappingEvent
	MetadataPutMappingResponse                   = 0x9286abfe // metadata.putMappingResponse
	MetadataPutTagMappingBootstrap               = 0x5fc8ab9b // metadata.putTagMappingBootstrap
	MetadataResetFlood                           = 0x9faf5282 // metadata.resetFlood
	MetadataResetFlood2                          = 0x88d0fd5e // metadata.resetFlood2
	MetadataResetFloodResponse                   = 0x9286abee // metadata.resetFloodResponse
	MetadataResetFloodResponse2                  = 0x9286abef // metadata.resetFloodResponse2
	NetPid                                       = 0x46409ccf // net.pid
	Pair                                         = 0x0f3c47ab // pair
	ResultFalse                                  = 0x27930a7b // resultFalse
	ResultTrue                                   = 0x3f9c8ef8 // resultTrue
	Stat                                         = 0x9d56e6b2 // stat
	StatshouseAddMetricsBatch                    = 0x56580239 // statshouse.addMetricsBatch
	StatshouseAutoCreate                         = 0x28bea524 // statshouse.autoCreate
	StatshouseCentroid                           = 0x73fd01e0 // statshouse.centroid
	StatshouseCommonProxyHeader                  = 0x6c803d07 // statshouse.commonProxyHeader
	StatshouseGetConfig2                         = 0x4285ff57 // statshouse.getConfig2
	StatshouseGetConfigResult                    = 0x0c803d07 // statshouse.getConfigResult
	StatshouseGetMetrics3                        = 0x42855554 // statshouse.getMetrics3
	StatshouseGetMetricsResult                   = 0x0c803d05 // statshouse.getMetricsResult
	StatshouseGetTagMapping2                     = 0x4285ff56 // statshouse.getTagMapping2
	StatshouseGetTagMappingBootstrap             = 0x75a7f68e // statshouse.getTagMappingBootstrap
	StatshouseGetTagMappingBootstrapResult       = 0x486a40de // statshouse.getTagMappingBootstrapResult
	StatshouseGetTagMappingResult                = 0x1a7d91fd // statshouse.getTagMappingResult
	StatshouseGetTargets2                        = 0x41df72a3 // statshouse.getTargets2
	StatshouseGetTargetsResult                   = 0x51ac86df // statshouse.getTargetsResult
	StatshouseIngestionStatus2                   = 0x2e17a6d3 // statshouse.ingestion_status2
	StatshouseMapping                            = 0xbf401d4b // statshouse.mapping
	StatshouseMetric                             = 0x3325d884 // statshouse.metric
	StatshouseMultiItem                          = 0x0c803e07 // statshouse.multi_item
	StatshouseMultiValue                         = 0x0c803e06 // statshouse.multi_value
	StatshousePromTarget                         = 0xac5296df // statshouse.promTarget
	StatshousePutTagMappingBootstrapResult       = 0x486affde // statshouse.putTagMappingBootstrapResult
	StatshouseSampleFactor                       = 0x4f7b7822 // statshouse.sample_factor
	StatshouseSendKeepAlive2                     = 0x4285ff53 // statshouse.sendKeepAlive2
	StatshouseSendSourceBucket2                  = 0x44575940 // statshouse.sendSourceBucket2
	StatshouseSourceBucket2                      = 0x3af6e822 // statshouse.sourceBucket2
	StatshouseTestConnection2                    = 0x4285ff58 // statshouse.testConnection2
	StatshouseTopElement                         = 0x9ffdea42 // statshouse.top_element
	StatshouseApiChunkResponse                   = 0x63928b42 // statshouseApi.chunkResponse
	StatshouseApiFilter                          = 0x511276a6 // statshouseApi.filter
	StatshouseApiFlagAuto                        = 0x2a6e4c14 // statshouseApi.flagAuto
	StatshouseApiFlagMapped                      = 0x670ab89c // statshouseApi.flagMapped
	StatshouseApiFlagRaw                         = 0x4ca979c0 // statshouseApi.flagRaw
	StatshouseApiFnAvg                           = 0x6323c2f6 // statshouseApi.fnAvg
	StatshouseApiFnCount                         = 0x89689775 // statshouseApi.fnCount
	StatshouseApiFnCountNorm                     = 0x60e68b5c // statshouseApi.fnCountNorm
	StatshouseApiFnCumulAvg                      = 0xf4d9ad09 // statshouseApi.fnCumulAvg
	StatshouseApiFnCumulCount                    = 0x871201c4 // statshouseApi.fnCumulCount
	StatshouseApiFnCumulSum                      = 0x42fc39b6 // statshouseApi.fnCumulSum
	StatshouseApiFnDerivativeAvg                 = 0x60d2b603 // statshouseApi.fnDerivativeAvg
	StatshouseApiFnDerivativeCount               = 0xe617771c // statshouseApi.fnDerivativeCount
	StatshouseApiFnDerivativeCountNorm           = 0xbfb5f7fc // statshouseApi.fnDerivativeCountNorm
	StatshouseApiFnDerivativeMax                 = 0x43eeb810 // statshouseApi.fnDerivativeMax
	StatshouseApiFnDerivativeMin                 = 0x4817df2b // statshouseApi.fnDerivativeMin
	StatshouseApiFnDerivativeSum                 = 0xa3a43781 // statshouseApi.fnDerivativeSum
	StatshouseApiFnDerivativeSumNorm             = 0x96683390 // statshouseApi.fnDerivativeSumNorm
	StatshouseApiFnDerivativeUnique              = 0x5745a0a3 // statshouseApi.fnDerivativeUnique
	StatshouseApiFnDerivativeUniqueNorm          = 0x4bd4f327 // statshouseApi.fnDerivativeUniqueNorm
	StatshouseApiFnMax                           = 0xf90de384 // statshouseApi.fnMax
	StatshouseApiFnMaxCountHost                  = 0x885e665b // statshouseApi.fnMaxCountHost
	StatshouseApiFnMaxHost                       = 0xb4790064 // statshouseApi.fnMaxHost
	StatshouseApiFnMin                           = 0xb4cb2644 // statshouseApi.fnMin
	StatshouseApiFnP25                           = 0xcf9ad7bf // statshouseApi.fnP25
	StatshouseApiFnP50                           = 0x77c5de5c // statshouseApi.fnP50
	StatshouseApiFnP75                           = 0x0e674272 // statshouseApi.fnP75
	StatshouseApiFnP90                           = 0xd4c8c793 // statshouseApi.fnP90
	StatshouseApiFnP95                           = 0x9a92b76f // statshouseApi.fnP95
	StatshouseApiFnP99                           = 0x71992e9a // statshouseApi.fnP99
	StatshouseApiFnP999                          = 0xa3434c26 // statshouseApi.fnP999
	StatshouseApiFnStddev                        = 0x2043e480 // statshouseApi.fnStddev
	StatshouseApiFnSum                           = 0x80ce3cf1 // statshouseApi.fnSum
	StatshouseApiFnSumNorm                       = 0x361963d5 // statshouseApi.fnSumNorm
	StatshouseApiFnUnique                        = 0xf20fb854 // statshouseApi.fnUnique
	StatshouseApiFnUniqueNorm                    = 0x9ceb6f68 // statshouseApi.fnUniqueNorm
	StatshouseApiGetChunk                        = 0x52721884 // statshouseApi.getChunk
	StatshouseApiGetQuery                        = 0x0c7349bb // statshouseApi.getQuery
	StatshouseApiGetQueryPoint                   = 0x0c7348bb // statshouseApi.getQueryPoint
	StatshouseApiPointMeta                       = 0x5c2bf296 // statshouseApi.pointMeta
	StatshouseApiQuery                           = 0xc9951bb9 // statshouseApi.query
	StatshouseApiQueryPoint                      = 0xc9951bbb // statshouseApi.queryPoint
	StatshouseApiQueryPointResponse              = 0x4487e41a // statshouseApi.queryPointResponse
	StatshouseApiQueryResponse                   = 0x4487e49a // statshouseApi.queryResponse
	StatshouseApiReleaseChunks                   = 0x62adc773 // statshouseApi.releaseChunks
	StatshouseApiReleaseChunksResponse           = 0xd12dc2bd // statshouseApi.releaseChunksResponse
	StatshouseApiSeries                          = 0x07a3e919 // statshouseApi.series
	StatshouseApiSeriesMeta                      = 0x5c2bf286 // statshouseApi.seriesMeta
	StatshouseApiTagValue                        = 0x43eeb763 // statshouseApi.tagValue
	True                                         = 0x3fedd339 // true
	Tuple                                        = 0x9770768a // tuple
	Vector                                       = 0x1cb5c415 // vector
	VectorTotal                                  = 0x10133f47 // vectorTotal
)
