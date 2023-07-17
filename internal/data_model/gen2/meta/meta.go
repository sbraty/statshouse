// Copyright 2023 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package meta

import (
	"fmt"

	"github.com/vkcom/statshouse/internal/data_model/gen2/internal"
	"github.com/vkcom/statshouse/internal/vkgo/basictl"
)

// We can create only types which have zero type arguments and zero nat arguments
type Object interface {
	TLName() string // returns type's TL name. For union, returns constructor name depending on actual union value
	TLTag() uint32  // returns type's TL tag. For union, returns constructor tag depending on actual union value
	String() string // returns type's representation for debugging (JSON for now)

	Read(w []byte) ([]byte, error)       // reads type's bare TL representation by consuming bytes from the start of w and returns remaining bytes, plus error
	Write(w []byte) ([]byte, error)      // appends bytes of type's bare TL representation to the end of w and returns it, plus error
	ReadBoxed(w []byte) ([]byte, error)  // same as Read, but reads/checks TLTag first
	WriteBoxed(w []byte) ([]byte, error) // same as Write, but writes TLTag first

	MarshalJSON() ([]byte, error)       // returns type's JSON representation, plus error
	UnmarshalJSON([]byte) error         // reads type's JSON representation
	WriteJSON(w []byte) ([]byte, error) // like MarshalJSON, but appends to w and returns it
}

type Function interface {
	Object

	ReadResultWriteResultJSON(r []byte, w []byte) ([]byte, []byte, error) // combination of ReadResult(r) + WriteResultJSON(w). Returns new r, new w, plus error
	ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) // combination of ReadResultJSON(r) + WriteResult(w). Returns new r, new w, plus error
}

// for quick one-liners
func GetTLName(tag uint32, notFoundName string) string {
	if item := FactoryItemByTLTag(tag); item != nil {
		return item.TLName()
	}
	return notFoundName
}

func CreateFunction(tag uint32) Function {
	if item := FactoryItemByTLTag(tag); item != nil && item.createFunction != nil {
		return item.createFunction()
	}
	return nil
}

func CreateObject(tag uint32) Object {
	if item := FactoryItemByTLTag(tag); item != nil && item.createObject != nil {
		return item.createObject()
	}
	return nil
}

// name can be in any of 3 forms "ch_proxy.insert#7cf362ba", "ch_proxy.insert" or "#7cf362ba"
func CreateFunctionFromName(name string) Function {
	if item := FactoryItemByTLName(name); item != nil && item.createFunction != nil {
		return item.createFunction()
	}
	return nil
}

// name can be in any of 3 forms "ch_proxy.insert#7cf362ba", "ch_proxy.insert" or "#7cf362ba"
func CreateObjectFromName(name string) Object {
	if item := FactoryItemByTLName(name); item != nil && item.createObject != nil {
		return item.createObject()
	}
	return nil
}

type TLItem struct {
	tag            uint32
	tlName         string
	createFunction func() Function
	createObject   func() Object
	// TODO - annotations, etc
}

func (item TLItem) IsFunction() bool         { return item.createFunction != nil }
func (item TLItem) TLTag() uint32            { return item.tag }
func (item TLItem) TLName() string           { return item.tlName }
func (item TLItem) CreateFunction() Function { return item.createFunction() }
func (item TLItem) CreateObject() Object     { return item.createObject() }

// TLItem serves as a single type for all enum values
func (item *TLItem) Reset()                              {}
func (item *TLItem) Read(w []byte) ([]byte, error)       { return w, nil }
func (item *TLItem) Write(w []byte) ([]byte, error)      { return w, nil }
func (item *TLItem) ReadBoxed(w []byte) ([]byte, error)  { return basictl.NatReadExactTag(w, item.tag) }
func (item *TLItem) WriteBoxed(w []byte) ([]byte, error) { return basictl.NatWrite(w, item.tag), nil }
func (item TLItem) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}
func (item *TLItem) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return internal.ErrorInvalidJSON(item.tlName, "expected json object")
	}
	for k := range _jm {
		return internal.ErrorInvalidJSONExcessElement(item.tlName, k)
	}
	return nil
}
func (item *TLItem) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	return append(w, '}'), nil
}
func (item *TLItem) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}
func (item *TLItem) UnmarshalJSON(b []byte) error {
	j, err := internal.JsonBytesToInterface(b)
	if err != nil {
		return internal.ErrorInvalidJSON(item.tlName, err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return internal.ErrorInvalidJSON(item.tlName, err.Error())
	}
	return nil
}

func FactoryItemByTLTag(tag uint32) *TLItem {
	return itemsByTag[tag]
}

func FactoryItemByTLName(name string) *TLItem {
	return itemsByName[name]
}

var itemsByTag = map[uint32]*TLItem{}

var itemsByName = map[string]*TLItem{}

func SetGlobalFactoryCreateForFunction(itemTag uint32, createObject func() Object, createFunction func() Function) {
	item := itemsByTag[itemTag]
	if item == nil {
		panic(fmt.Sprintf("factory cannot find function tag #%08x to set", itemTag))
	}
	item.createObject = createObject
	item.createFunction = createFunction
}

func SetGlobalFactoryCreateForObject(itemTag uint32, createObject func() Object) {
	item := itemsByTag[itemTag]
	if item == nil {
		panic(fmt.Sprintf("factory cannot find item tag #%08x to set", itemTag))
	}
	item.createObject = createObject
}

func SetGlobalFactoryCreateForEnumElement(itemTag uint32) {
	item := itemsByTag[itemTag]
	if item == nil {
		panic(fmt.Sprintf("factory cannot find enum tag #%08x to set", itemTag))
	}
	item.createObject = func() Object { return item }
}

func pleaseImportFactoryObject() Object {
	panic("factory functions are not linked to reduce code bloat, please import 'gen/factory' instead of 'gen/meta'.")
}

func pleaseImportFactoryFunction() Function {
	panic("factory functions are not linked to reduce code bloat, please import 'gen/factory' instead of 'gen/meta'.")
}

func fillObject(n1 string, n2 string, item *TLItem) {
	itemsByTag[item.tag] = item
	itemsByName[item.tlName] = item
	itemsByName[n1] = item
	itemsByName[n2] = item
	item.createObject = pleaseImportFactoryObject
	// code below is as fast, but allocates some extra strings which are already in binary const segment due to JSON code
	// itemsByName[fmt.Sprintf("%s#%08x", item.tlName, item.tag)] = item
	// itemsByName[fmt.Sprintf("#%08x", item.tag)] = item
}

func fillFunction(n1 string, n2 string, item *TLItem) {
	fillObject(n1, n2, item)
	item.createFunction = pleaseImportFactoryFunction
}

func init() {
	fillObject("boolStat#92cbcbfa", "#92cbcbfa", &TLItem{tag: 0x92cbcbfa, tlName: "boolStat"})
	fillObject("engine.alreadyInMasterMode#402409cb", "#402409cb", &TLItem{tag: 0x402409cb, tlName: "engine.alreadyInMasterMode"})
	fillObject("engine.alreadyInReplicaMode#ebd80142", "#ebd80142", &TLItem{tag: 0xebd80142, tlName: "engine.alreadyInReplicaMode"})
	fillFunction("engine.asyncSleep#60e50d3d", "#60e50d3d", &TLItem{tag: 0x60e50d3d, tlName: "engine.asyncSleep"})
	fillObject("engine.binlogPrefix#4c09c894", "#4c09c894", &TLItem{tag: 0x4c09c894, tlName: "engine.binlogPrefix"})
	fillFunction("engine.count#19d0f020", "#19d0f020", &TLItem{tag: 0x19d0f020, tlName: "engine.count"})
	fillFunction("engine.dumpForceQueries#f1f90880", "#f1f90880", &TLItem{tag: 0xf1f90880, tlName: "engine.dumpForceQueries"})
	fillFunction("engine.dumpLastQueries#c060a29f", "#c060a29f", &TLItem{tag: 0xc060a29f, tlName: "engine.dumpLastQueries"})
	fillFunction("engine.dumpNextQueries#e65872ad", "#e65872ad", &TLItem{tag: 0xe65872ad, tlName: "engine.dumpNextQueries"})
	fillFunction("engine.enableMetafilesAnalyzer#88836cdc", "#88836cdc", &TLItem{tag: 0x88836cdc, tlName: "engine.enableMetafilesAnalyzer"})
	fillObject("engine.failedToSwitchMode#17418662", "#17418662", &TLItem{tag: 0x17418662, tlName: "engine.failedToSwitchMode"})
	fillFunction("engine.filteredStat#594870d6", "#594870d6", &TLItem{tag: 0x594870d6, tlName: "engine.filteredStat"})
	fillFunction("engine.getBinlogPrefixes#ef14db93", "#ef14db93", &TLItem{tag: 0xef14db93, tlName: "engine.getBinlogPrefixes"})
	fillFunction("engine.getExpectedMetafilesStats#0342f391", "#0342f391", &TLItem{tag: 0x342f391, tlName: "engine.getExpectedMetafilesStats"})
	fillFunction("engine.getReadWriteMode#61b3f593", "#61b3f593", &TLItem{tag: 0x61b3f593, tlName: "engine.getReadWriteMode"})
	fillFunction("engine.getReindexStatus#f492042e", "#f492042e", &TLItem{tag: 0xf492042e, tlName: "engine.getReindexStatus"})
	fillObject("engine.httpQuery#58300321", "#58300321", &TLItem{tag: 0x58300321, tlName: "engine.httpQuery"})
	fillObject("engine.httpQueryResponse#284852fc", "#284852fc", &TLItem{tag: 0x284852fc, tlName: "engine.httpQueryResponse"})
	fillFunction("engine.invokeHttpQuery#f4c73c0b", "#f4c73c0b", &TLItem{tag: 0xf4c73c0b, tlName: "engine.invokeHttpQuery"})
	fillFunction("engine.isProduction#ccdea0ac", "#ccdea0ac", &TLItem{tag: 0xccdea0ac, tlName: "engine.isProduction"})
	fillObject("engine.metafilesOneMemoryStat#c292e4a6", "#c292e4a6", &TLItem{tag: 0xc292e4a6, tlName: "engine.metafilesOneMemoryStat"})
	fillObject("engine.metafilesStatData#b673669b", "#b673669b", &TLItem{tag: 0xb673669b, tlName: "engine.metafilesStatData"})
	fillFunction("engine.nop#166bb7c6", "#166bb7c6", &TLItem{tag: 0x166bb7c6, tlName: "engine.nop"})
	fillFunction("engine.pid#559d6e36", "#559d6e36", &TLItem{tag: 0x559d6e36, tlName: "engine.pid"})
	fillFunction("engine.pushStat#f4b19fa2", "#f4b19fa2", &TLItem{tag: 0xf4b19fa2, tlName: "engine.pushStat"})
	fillObject("engine.queryResult#ac4d6fe9", "#ac4d6fe9", &TLItem{tag: 0xac4d6fe9, tlName: "engine.queryResult"})
	fillObject("engine.queryResultAio#ee2879b0", "#ee2879b0", &TLItem{tag: 0xee2879b0, tlName: "engine.queryResultAio"})
	fillObject("engine.queryResultError#2b4dd0ba", "#2b4dd0ba", &TLItem{tag: 0x2b4dd0ba, tlName: "engine.queryResultError"})
	fillFunction("engine.readNop#9d2b841f", "#9d2b841f", &TLItem{tag: 0x9d2b841f, tlName: "engine.readNop"})
	fillFunction("engine.recordNextQueries#0001e9d6", "#0001e9d6", &TLItem{tag: 0x1e9d6, tlName: "engine.recordNextQueries"})
	fillFunction("engine.registerDynamicLib#2f86f276", "#2f86f276", &TLItem{tag: 0x2f86f276, tlName: "engine.registerDynamicLib"})
	fillObject("engine.reindexStatusDone#0f67569a", "#0f67569a", &TLItem{tag: 0xf67569a, tlName: "engine.reindexStatusDone"})
	fillObject("engine.reindexStatusDoneOld#afdbd505", "#afdbd505", &TLItem{tag: 0xafdbd505, tlName: "engine.reindexStatusDoneOld"})
	fillObject("engine.reindexStatusFailed#10533721", "#10533721", &TLItem{tag: 0x10533721, tlName: "engine.reindexStatusFailed"})
	fillObject("engine.reindexStatusNever#7f6a89b9", "#7f6a89b9", &TLItem{tag: 0x7f6a89b9, tlName: "engine.reindexStatusNever"})
	fillObject("engine.reindexStatusRunning#fa198b59", "#fa198b59", &TLItem{tag: 0xfa198b59, tlName: "engine.reindexStatusRunning"})
	fillObject("engine.reindexStatusRunningOld#ac530b46", "#ac530b46", &TLItem{tag: 0xac530b46, tlName: "engine.reindexStatusRunningOld"})
	fillObject("engine.reindexStatusSignaled#756e878b", "#756e878b", &TLItem{tag: 0x756e878b, tlName: "engine.reindexStatusSignaled"})
	fillFunction("engine.reloadDynamicLib#602d62c1", "#602d62c1", &TLItem{tag: 0x602d62c1, tlName: "engine.reloadDynamicLib"})
	fillObject("engine.reloadDynamicLibOptions#0f3d0fb1", "#0f3d0fb1", &TLItem{tag: 0xf3d0fb1, tlName: "engine.reloadDynamicLibOptions"})
	fillFunction("engine.replaceConfigServer#5fcd8e77", "#5fcd8e77", &TLItem{tag: 0x5fcd8e77, tlName: "engine.replaceConfigServer"})
	fillFunction("engine.sendSignal#1a7708a3", "#1a7708a3", &TLItem{tag: 0x1a7708a3, tlName: "engine.sendSignal"})
	fillFunction("engine.setFsyncInterval#665d2ab7", "#665d2ab7", &TLItem{tag: 0x665d2ab7, tlName: "engine.setFsyncInterval"})
	fillFunction("engine.setMetafileMemory#7bdcf404", "#7bdcf404", &TLItem{tag: 0x7bdcf404, tlName: "engine.setMetafileMemory"})
	fillFunction("engine.setNoPersistentConfigArray#5806a520", "#5806a520", &TLItem{tag: 0x5806a520, tlName: "engine.setNoPersistentConfigArray"})
	fillFunction("engine.setNoPersistentConfigValue#92aaa5b9", "#92aaa5b9", &TLItem{tag: 0x92aaa5b9, tlName: "engine.setNoPersistentConfigValue"})
	fillFunction("engine.setPersistentConfigArray#fc99af0b", "#fc99af0b", &TLItem{tag: 0xfc99af0b, tlName: "engine.setPersistentConfigArray"})
	fillFunction("engine.setPersistentConfigValue#4cc8953f", "#4cc8953f", &TLItem{tag: 0x4cc8953f, tlName: "engine.setPersistentConfigValue"})
	fillFunction("engine.setVerbosity#9d980926", "#9d980926", &TLItem{tag: 0x9d980926, tlName: "engine.setVerbosity"})
	fillFunction("engine.setVerbosityType#5388c0ae", "#5388c0ae", &TLItem{tag: 0x5388c0ae, tlName: "engine.setVerbosityType"})
	fillFunction("engine.sleep#3d3bcd48", "#3d3bcd48", &TLItem{tag: 0x3d3bcd48, tlName: "engine.sleep"})
	fillFunction("engine.stat#efb3c36b", "#efb3c36b", &TLItem{tag: 0xefb3c36b, tlName: "engine.stat"})
	fillFunction("engine.switchToMasterMode#8cdcb5f9", "#8cdcb5f9", &TLItem{tag: 0x8cdcb5f9, tlName: "engine.switchToMasterMode"})
	fillFunction("engine.switchToMasterModeForcefully#1973fa8f", "#1973fa8f", &TLItem{tag: 0x1973fa8f, tlName: "engine.switchToMasterModeForcefully"})
	fillFunction("engine.switchToReplicaMode#23c3a87e", "#23c3a87e", &TLItem{tag: 0x23c3a87e, tlName: "engine.switchToReplicaMode"})
	fillObject("engine.switchedToMasterMode#95b13964", "#95b13964", &TLItem{tag: 0x95b13964, tlName: "engine.switchedToMasterMode"})
	fillObject("engine.switchedToMasterModeForcefully#ec61b4be", "#ec61b4be", &TLItem{tag: 0xec61b4be, tlName: "engine.switchedToMasterModeForcefully"})
	fillObject("engine.switchedToReplicaMode#ad642a0b", "#ad642a0b", &TLItem{tag: 0xad642a0b, tlName: "engine.switchedToReplicaMode"})
	fillFunction("engine.unregisterDynamicLib#84d5fcb9", "#84d5fcb9", &TLItem{tag: 0x84d5fcb9, tlName: "engine.unregisterDynamicLib"})
	fillFunction("engine.version#1a2e06fa", "#1a2e06fa", &TLItem{tag: 0x1a2e06fa, tlName: "engine.version"})
	fillFunction("engine.writeNop#58160af4", "#58160af4", &TLItem{tag: 0x58160af4, tlName: "engine.writeNop"})
	fillObject("int#a8509bda", "#a8509bda", &TLItem{tag: 0xa8509bda, tlName: "int"})
	fillObject("metadata.createEntityEvent#1a345674", "#1a345674", &TLItem{tag: 0x1a345674, tlName: "metadata.createEntityEvent"})
	fillObject("metadata.createMappingEvent#12345678", "#12345678", &TLItem{tag: 0x12345678, tlName: "metadata.createMappingEvent"})
	fillObject("metadata.createMetricEvent#12345674", "#12345674", &TLItem{tag: 0x12345674, tlName: "metadata.createMetricEvent"})
	fillObject("metadata.editEntityEvent#1234b677", "#1234b677", &TLItem{tag: 0x1234b677, tlName: "metadata.editEntityEvent"})
	fillFunction("metadata.editEntitynew#86df475f", "#86df475f", &TLItem{tag: 0x86df475f, tlName: "metadata.editEntitynew"})
	fillObject("metadata.editMetricEvent#12345677", "#12345677", &TLItem{tag: 0x12345677, tlName: "metadata.editMetricEvent"})
	fillObject("metadata.event#9286affa", "#9286affa", &TLItem{tag: 0x9286affa, tlName: "metadata.event"})
	fillFunction("metadata.getInvertMapping#9faf5280", "#9faf5280", &TLItem{tag: 0x9faf5280, tlName: "metadata.getInvertMapping"})
	fillFunction("metadata.getJournalnew#93ba92f8", "#93ba92f8", &TLItem{tag: 0x93ba92f8, tlName: "metadata.getJournalnew"})
	fillFunction("metadata.getMapping#9dfa7a83", "#9dfa7a83", &TLItem{tag: 0x9dfa7a83, tlName: "metadata.getMapping"})
	fillFunction("metadata.getMetrics#93ba92f5", "#93ba92f5", &TLItem{tag: 0x93ba92f5, tlName: "metadata.getMetrics"})
	fillFunction("metadata.getTagMappingBootstrap#5fc81a9b", "#5fc81a9b", &TLItem{tag: 0x5fc81a9b, tlName: "metadata.getTagMappingBootstrap"})
	fillObject("metadata.putBootstrapEvent#5854dfaf", "#5854dfaf", &TLItem{tag: 0x5854dfaf, tlName: "metadata.putBootstrapEvent"})
	fillFunction("metadata.putMapping#9faf5281", "#9faf5281", &TLItem{tag: 0x9faf5281, tlName: "metadata.putMapping"})
	fillObject("metadata.putMappingEvent#12345676", "#12345676", &TLItem{tag: 0x12345676, tlName: "metadata.putMappingEvent"})
	fillObject("metadata.putMappingResponse#9286abfe", "#9286abfe", &TLItem{tag: 0x9286abfe, tlName: "metadata.putMappingResponse"})
	fillFunction("metadata.putTagMappingBootstrap#5fc8ab9b", "#5fc8ab9b", &TLItem{tag: 0x5fc8ab9b, tlName: "metadata.putTagMappingBootstrap"})
	fillFunction("metadata.resetFlood#9faf5282", "#9faf5282", &TLItem{tag: 0x9faf5282, tlName: "metadata.resetFlood"})
	fillFunction("metadata.resetFlood2#88d0fd5e", "#88d0fd5e", &TLItem{tag: 0x88d0fd5e, tlName: "metadata.resetFlood2"})
	fillObject("metadata.resetFloodResponse#9286abee", "#9286abee", &TLItem{tag: 0x9286abee, tlName: "metadata.resetFloodResponse"})
	fillObject("metadata.resetFloodResponse2#9286abef", "#9286abef", &TLItem{tag: 0x9286abef, tlName: "metadata.resetFloodResponse2"})
	fillObject("net.pid#46409ccf", "#46409ccf", &TLItem{tag: 0x46409ccf, tlName: "net.pid"})
	fillObject("stat#9d56e6b2", "#9d56e6b2", &TLItem{tag: 0x9d56e6b2, tlName: "stat"})
	fillFunction("statshouse.addMetricsBatch#56580239", "#56580239", &TLItem{tag: 0x56580239, tlName: "statshouse.addMetricsBatch"})
	fillObject("statshouseApi.filter#511276a6", "#511276a6", &TLItem{tag: 0x511276a6, tlName: "statshouseApi.filter"})
	fillObject("statshouseApi.flagAuto#2a6e4c14", "#2a6e4c14", &TLItem{tag: 0x2a6e4c14, tlName: "statshouseApi.flagAuto"})
	fillObject("statshouseApi.flagMapped#670ab89c", "#670ab89c", &TLItem{tag: 0x670ab89c, tlName: "statshouseApi.flagMapped"})
	fillObject("statshouseApi.flagRaw#4ca979c0", "#4ca979c0", &TLItem{tag: 0x4ca979c0, tlName: "statshouseApi.flagRaw"})
	fillObject("statshouseApi.fnAvg#6323c2f6", "#6323c2f6", &TLItem{tag: 0x6323c2f6, tlName: "statshouseApi.fnAvg"})
	fillObject("statshouseApi.fnCount#89689775", "#89689775", &TLItem{tag: 0x89689775, tlName: "statshouseApi.fnCount"})
	fillObject("statshouseApi.fnCountNorm#60e68b5c", "#60e68b5c", &TLItem{tag: 0x60e68b5c, tlName: "statshouseApi.fnCountNorm"})
	fillObject("statshouseApi.fnCumulAvg#f4d9ad09", "#f4d9ad09", &TLItem{tag: 0xf4d9ad09, tlName: "statshouseApi.fnCumulAvg"})
	fillObject("statshouseApi.fnCumulCount#871201c4", "#871201c4", &TLItem{tag: 0x871201c4, tlName: "statshouseApi.fnCumulCount"})
	fillObject("statshouseApi.fnCumulSum#42fc39b6", "#42fc39b6", &TLItem{tag: 0x42fc39b6, tlName: "statshouseApi.fnCumulSum"})
	fillObject("statshouseApi.fnDerivativeAvg#60d2b603", "#60d2b603", &TLItem{tag: 0x60d2b603, tlName: "statshouseApi.fnDerivativeAvg"})
	fillObject("statshouseApi.fnDerivativeCount#e617771c", "#e617771c", &TLItem{tag: 0xe617771c, tlName: "statshouseApi.fnDerivativeCount"})
	fillObject("statshouseApi.fnDerivativeCountNorm#bfb5f7fc", "#bfb5f7fc", &TLItem{tag: 0xbfb5f7fc, tlName: "statshouseApi.fnDerivativeCountNorm"})
	fillObject("statshouseApi.fnDerivativeMax#43eeb810", "#43eeb810", &TLItem{tag: 0x43eeb810, tlName: "statshouseApi.fnDerivativeMax"})
	fillObject("statshouseApi.fnDerivativeMin#4817df2b", "#4817df2b", &TLItem{tag: 0x4817df2b, tlName: "statshouseApi.fnDerivativeMin"})
	fillObject("statshouseApi.fnDerivativeSum#a3a43781", "#a3a43781", &TLItem{tag: 0xa3a43781, tlName: "statshouseApi.fnDerivativeSum"})
	fillObject("statshouseApi.fnDerivativeSumNorm#96683390", "#96683390", &TLItem{tag: 0x96683390, tlName: "statshouseApi.fnDerivativeSumNorm"})
	fillObject("statshouseApi.fnDerivativeUnique#5745a0a3", "#5745a0a3", &TLItem{tag: 0x5745a0a3, tlName: "statshouseApi.fnDerivativeUnique"})
	fillObject("statshouseApi.fnDerivativeUniqueNorm#4bd4f327", "#4bd4f327", &TLItem{tag: 0x4bd4f327, tlName: "statshouseApi.fnDerivativeUniqueNorm"})
	fillObject("statshouseApi.fnMax#f90de384", "#f90de384", &TLItem{tag: 0xf90de384, tlName: "statshouseApi.fnMax"})
	fillObject("statshouseApi.fnMaxCountHost#885e665b", "#885e665b", &TLItem{tag: 0x885e665b, tlName: "statshouseApi.fnMaxCountHost"})
	fillObject("statshouseApi.fnMaxHost#b4790064", "#b4790064", &TLItem{tag: 0xb4790064, tlName: "statshouseApi.fnMaxHost"})
	fillObject("statshouseApi.fnMin#b4cb2644", "#b4cb2644", &TLItem{tag: 0xb4cb2644, tlName: "statshouseApi.fnMin"})
	fillObject("statshouseApi.fnP01#381b1cee", "#381b1cee", &TLItem{tag: 0x381b1cee, tlName: "statshouseApi.fnP01"})
	fillObject("statshouseApi.fnP1#bbb36a23", "#bbb36a23", &TLItem{tag: 0xbbb36a23, tlName: "statshouseApi.fnP1"})
	fillObject("statshouseApi.fnP10#56071d39", "#56071d39", &TLItem{tag: 0x56071d39, tlName: "statshouseApi.fnP10"})
	fillObject("statshouseApi.fnP25#cf9ad7bf", "#cf9ad7bf", &TLItem{tag: 0xcf9ad7bf, tlName: "statshouseApi.fnP25"})
	fillObject("statshouseApi.fnP5#bcdeae3a", "#bcdeae3a", &TLItem{tag: 0xbcdeae3a, tlName: "statshouseApi.fnP5"})
	fillObject("statshouseApi.fnP50#77c5de5c", "#77c5de5c", &TLItem{tag: 0x77c5de5c, tlName: "statshouseApi.fnP50"})
	fillObject("statshouseApi.fnP75#0e674272", "#0e674272", &TLItem{tag: 0xe674272, tlName: "statshouseApi.fnP75"})
	fillObject("statshouseApi.fnP90#d4c8c793", "#d4c8c793", &TLItem{tag: 0xd4c8c793, tlName: "statshouseApi.fnP90"})
	fillObject("statshouseApi.fnP95#9a92b76f", "#9a92b76f", &TLItem{tag: 0x9a92b76f, tlName: "statshouseApi.fnP95"})
	fillObject("statshouseApi.fnP99#71992e9a", "#71992e9a", &TLItem{tag: 0x71992e9a, tlName: "statshouseApi.fnP99"})
	fillObject("statshouseApi.fnP999#a3434c26", "#a3434c26", &TLItem{tag: 0xa3434c26, tlName: "statshouseApi.fnP999"})
	fillObject("statshouseApi.fnStddev#2043e480", "#2043e480", &TLItem{tag: 0x2043e480, tlName: "statshouseApi.fnStddev"})
	fillObject("statshouseApi.fnSum#80ce3cf1", "#80ce3cf1", &TLItem{tag: 0x80ce3cf1, tlName: "statshouseApi.fnSum"})
	fillObject("statshouseApi.fnSumNorm#361963d5", "#361963d5", &TLItem{tag: 0x361963d5, tlName: "statshouseApi.fnSumNorm"})
	fillObject("statshouseApi.fnUnique#f20fb854", "#f20fb854", &TLItem{tag: 0xf20fb854, tlName: "statshouseApi.fnUnique"})
	fillObject("statshouseApi.fnUniqueNorm#9ceb6f68", "#9ceb6f68", &TLItem{tag: 0x9ceb6f68, tlName: "statshouseApi.fnUniqueNorm"})
	fillFunction("statshouseApi.getChunk#52721884", "#52721884", &TLItem{tag: 0x52721884, tlName: "statshouseApi.getChunk"})
	fillObject("statshouseApi.chunkResponse#63928b42", "#63928b42", &TLItem{tag: 0x63928b42, tlName: "statshouseApi.chunkResponse"})
	fillFunction("statshouseApi.getQuery#0c7349bb", "#0c7349bb", &TLItem{tag: 0xc7349bb, tlName: "statshouseApi.getQuery"})
	fillFunction("statshouseApi.getQueryPoint#0c7348bb", "#0c7348bb", &TLItem{tag: 0xc7348bb, tlName: "statshouseApi.getQueryPoint"})
	fillObject("statshouseApi.queryPointResponse#4487e41a", "#4487e41a", &TLItem{tag: 0x4487e41a, tlName: "statshouseApi.queryPointResponse"})
	fillObject("statshouseApi.pointMeta#5c2bf296", "#5c2bf296", &TLItem{tag: 0x5c2bf296, tlName: "statshouseApi.pointMeta"})
	fillObject("statshouseApi.query#c9951bb9", "#c9951bb9", &TLItem{tag: 0xc9951bb9, tlName: "statshouseApi.query"})
	fillObject("statshouseApi.queryPoint#c9951bbb", "#c9951bbb", &TLItem{tag: 0xc9951bbb, tlName: "statshouseApi.queryPoint"})
	fillFunction("statshouseApi.releaseChunks#62adc773", "#62adc773", &TLItem{tag: 0x62adc773, tlName: "statshouseApi.releaseChunks"})
	fillObject("statshouseApi.releaseChunksResponse#d12dc2bd", "#d12dc2bd", &TLItem{tag: 0xd12dc2bd, tlName: "statshouseApi.releaseChunksResponse"})
	fillObject("statshouseApi.series#07a3e919", "#07a3e919", &TLItem{tag: 0x7a3e919, tlName: "statshouseApi.series"})
	fillObject("statshouseApi.tagValue#43eeb763", "#43eeb763", &TLItem{tag: 0x43eeb763, tlName: "statshouseApi.tagValue"})
	fillFunction("statshouse.autoCreate#28bea524", "#28bea524", &TLItem{tag: 0x28bea524, tlName: "statshouse.autoCreate"})
	fillObject("statshouse.centroid#73fd01e0", "#73fd01e0", &TLItem{tag: 0x73fd01e0, tlName: "statshouse.centroid"})
	fillFunction("statshouse.getConfig2#4285ff57", "#4285ff57", &TLItem{tag: 0x4285ff57, tlName: "statshouse.getConfig2"})
	fillFunction("statshouse.getMetrics3#42855554", "#42855554", &TLItem{tag: 0x42855554, tlName: "statshouse.getMetrics3"})
	fillObject("statshouse.getMetricsResult#0c803d05", "#0c803d05", &TLItem{tag: 0xc803d05, tlName: "statshouse.getMetricsResult"})
	fillFunction("statshouse.getTagMapping2#4285ff56", "#4285ff56", &TLItem{tag: 0x4285ff56, tlName: "statshouse.getTagMapping2"})
	fillFunction("statshouse.getTagMappingBootstrap#75a7f68e", "#75a7f68e", &TLItem{tag: 0x75a7f68e, tlName: "statshouse.getTagMappingBootstrap"})
	fillObject("statshouse.getTagMappingBootstrapResult#486a40de", "#486a40de", &TLItem{tag: 0x486a40de, tlName: "statshouse.getTagMappingBootstrapResult"})
	fillObject("statshouse.getTagMappingResult#1a7d91fd", "#1a7d91fd", &TLItem{tag: 0x1a7d91fd, tlName: "statshouse.getTagMappingResult"})
	fillFunction("statshouse.getTargets2#41df72a3", "#41df72a3", &TLItem{tag: 0x41df72a3, tlName: "statshouse.getTargets2"})
	fillObject("statshouse.ingestion_status2#2e17a6d3", "#2e17a6d3", &TLItem{tag: 0x2e17a6d3, tlName: "statshouse.ingestion_status2"})
	fillObject("statshouse.mapping#bf401d4b", "#bf401d4b", &TLItem{tag: 0xbf401d4b, tlName: "statshouse.mapping"})
	fillObject("statshouse.metric#3325d884", "#3325d884", &TLItem{tag: 0x3325d884, tlName: "statshouse.metric"})
	fillObject("statshouse.multi_item#0c803e07", "#0c803e07", &TLItem{tag: 0xc803e07, tlName: "statshouse.multi_item"})
	fillObject("statshouse.promTarget#ac5296df", "#ac5296df", &TLItem{tag: 0xac5296df, tlName: "statshouse.promTarget"})
	fillObject("statshouse.putTagMappingBootstrapResult#486affde", "#486affde", &TLItem{tag: 0x486affde, tlName: "statshouse.putTagMappingBootstrapResult"})
	fillObject("statshouse.sample_factor#4f7b7822", "#4f7b7822", &TLItem{tag: 0x4f7b7822, tlName: "statshouse.sample_factor"})
	fillFunction("statshouse.sendKeepAlive2#4285ff53", "#4285ff53", &TLItem{tag: 0x4285ff53, tlName: "statshouse.sendKeepAlive2"})
	fillFunction("statshouse.sendSourceBucket2#44575940", "#44575940", &TLItem{tag: 0x44575940, tlName: "statshouse.sendSourceBucket2"})
	fillObject("statshouse.sourceBucket2#3af6e822", "#3af6e822", &TLItem{tag: 0x3af6e822, tlName: "statshouse.sourceBucket2"})
	fillFunction("statshouse.testConnection2#4285ff58", "#4285ff58", &TLItem{tag: 0x4285ff58, tlName: "statshouse.testConnection2"})
	fillObject("statshouse.top_element#9ffdea42", "#9ffdea42", &TLItem{tag: 0x9ffdea42, tlName: "statshouse.top_element"})
	fillObject("string#b5286e24", "#b5286e24", &TLItem{tag: 0xb5286e24, tlName: "string"})
	fillObject("true#3fedd339", "#3fedd339", &TLItem{tag: 0x3fedd339, tlName: "true"})
}
