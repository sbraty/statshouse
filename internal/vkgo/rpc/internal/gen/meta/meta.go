// Copyright 2024 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package meta

import (
	"fmt"

	"github.com/vkcom/statshouse/internal/vkgo/basictl"
	"github.com/vkcom/statshouse/internal/vkgo/rpc/internal/gen/internal"
)

// We can create only types which have zero type arguments and zero nat arguments
type Object interface {
	TLName() string // returns type's TL name. For union, returns constructor name depending on actual union value
	TLTag() uint32  // returns type's TL tag. For union, returns constructor tag depending on actual union value
	String() string // returns type's representation for debugging (JSON for now)

	Read(w []byte) ([]byte, error)              // reads type's bare TL representation by consuming bytes from the start of w and returns remaining bytes, plus error
	ReadBoxed(w []byte) ([]byte, error)         // same as Read, but reads/checks TLTag first (this method is general version of Write, use it only when you are working with interface)
	WriteGeneral(w []byte) ([]byte, error)      // appends bytes of type's bare TL representation to the end of w and returns it, plus error
	WriteBoxedGeneral(w []byte) ([]byte, error) // same as Write, but writes TLTag first (this method is general version of WriteBoxed, use it only when you are working with interface)

	MarshalJSON() ([]byte, error) // returns type's JSON representation, plus error
	UnmarshalJSON([]byte) error   // reads type's JSON representation

	ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error
	WriteJSONGeneral(w []byte) ([]byte, error) // like MarshalJSON, but appends to w and returns it (this method is general version of WriteBoxed, use it only when you are working with interface)
}

type Function interface {
	Object

	ReadResultWriteResultJSON(r []byte, w []byte) ([]byte, []byte, error) // combination of ReadResult(r) + WriteResultJSON(w). Returns new r, new w, plus error
	ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) // combination of ReadResultJSON(r) + WriteResult(w). Returns new r, new w, plus error

	// For transcoding short-long version during Long ID and newTypeNames transition
	ReadResultWriteResultJSONOpt(newTypeNames bool, short bool, r []byte, w []byte) ([]byte, []byte, error)
}

func GetAllTLItems() []TLItem {
	var allItems []TLItem
	for _, item := range itemsByName {
		if item != nil {
			allItems = append(allItems, *item)
		}
	}
	return allItems
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
	tag                uint32
	annotations        uint32
	tlName             string
	createFunction     func() Function
	createFunctionLong func() Function
	createObject       func() Object
}

func (item TLItem) TLTag() uint32            { return item.tag }
func (item TLItem) TLName() string           { return item.tlName }
func (item TLItem) CreateObject() Object     { return item.createObject() }
func (item TLItem) IsFunction() bool         { return item.createFunction != nil }
func (item TLItem) CreateFunction() Function { return item.createFunction() }

// For transcoding short-long version during Long ID transition
func (item TLItem) HasFunctionLong() bool        { return item.createFunctionLong != nil }
func (item TLItem) CreateFunctionLong() Function { return item.createFunctionLong() }

// Annotations
func (item TLItem) AnnotationAny() bool       { return item.annotations&0x1 != 0 }
func (item TLItem) AnnotationInternal() bool  { return item.annotations&0x2 != 0 }
func (item TLItem) AnnotationReadwrite() bool { return item.annotations&0x4 != 0 }

// TLItem serves as a single type for all enum values
func (item *TLItem) Reset()                                {}
func (item *TLItem) Read(w []byte) ([]byte, error)         { return w, nil }
func (item *TLItem) WriteGeneral(w []byte) ([]byte, error) { return w, nil }
func (item *TLItem) Write(w []byte) []byte                 { return w }
func (item *TLItem) ReadBoxed(w []byte) ([]byte, error)    { return basictl.NatReadExactTag(w, item.tag) }
func (item *TLItem) WriteBoxedGeneral(w []byte) ([]byte, error) {
	return basictl.NatWrite(w, item.tag), nil
}
func (item *TLItem) WriteBoxed(w []byte) []byte { return basictl.NatWrite(w, item.tag) }
func (item TLItem) String() string {
	return string(item.WriteJSON(nil))
}
func (item *TLItem) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	in.Delim('{')
	if !in.Ok() {
		return in.Error()
	}
	for !in.IsDelim('}') {
		return internal.ErrorInvalidJSONExcessElement(item.tlName, in.UnsafeFieldName(true))
	}
	in.Delim('}')
	if !in.Ok() {
		return in.Error()
	}
	return nil
}
func (item *TLItem) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSON(w), nil
}
func (item *TLItem) WriteJSON(w []byte) []byte {
	w = append(w, '{')
	return append(w, '}')
}
func (item *TLItem) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}
func (item *TLItem) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
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

func SetGlobalFactoryCreateForFunction(itemTag uint32, createObject func() Object, createFunction func() Function, createFunctionLong func() Function) {
	item := itemsByTag[itemTag]
	if item == nil {
		panic(fmt.Sprintf("factory cannot find function tag #%08x to set", itemTag))
	}
	item.createObject = createObject
	item.createFunction = createFunction
	item.createFunctionLong = createFunctionLong
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
	fillFunction("engine.asyncSleep#60e50d3d", "#60e50d3d", &TLItem{tag: 0x60e50d3d, annotations: 0x3, tlName: "engine.asyncSleep"})
	fillFunction("engine.filteredStat#594870d6", "#594870d6", &TLItem{tag: 0x594870d6, annotations: 0x1, tlName: "engine.filteredStat"})
	fillFunction("engine.pid#559d6e36", "#559d6e36", &TLItem{tag: 0x559d6e36, annotations: 0x1, tlName: "engine.pid"})
	fillFunction("engine.setVerbosity#9d980926", "#9d980926", &TLItem{tag: 0x9d980926, annotations: 0x3, tlName: "engine.setVerbosity"})
	fillFunction("engine.sleep#3d3bcd48", "#3d3bcd48", &TLItem{tag: 0x3d3bcd48, annotations: 0x3, tlName: "engine.sleep"})
	fillFunction("engine.stat#efb3c36b", "#efb3c36b", &TLItem{tag: 0xefb3c36b, annotations: 0x1, tlName: "engine.stat"})
	fillFunction("engine.version#1a2e06fa", "#1a2e06fa", &TLItem{tag: 0x1a2e06fa, annotations: 0x1, tlName: "engine.version"})
	fillFunction("go.pprof#ea2876a6", "#ea2876a6", &TLItem{tag: 0xea2876a6, annotations: 0x4, tlName: "go.pprof"})
	fillObject("net.pid#46409ccf", "#46409ccf", &TLItem{tag: 0x46409ccf, annotations: 0x0, tlName: "net.pid"})
	fillObject("netUdpPacket.encHeader#251a7bfd", "#251a7bfd", &TLItem{tag: 0x251a7bfd, annotations: 0x0, tlName: "netUdpPacket.encHeader"})
	fillObject("netUdpPacket.unencHeader#00a8e945", "#00a8e945", &TLItem{tag: 0xa8e945, annotations: 0x0, tlName: "netUdpPacket.unencHeader"})
	fillObject("reqError#b527877d", "#b527877d", &TLItem{tag: 0xb527877d, annotations: 0x0, tlName: "reqError"})
	fillObject("reqResultHeader#8cc84ce1", "#8cc84ce1", &TLItem{tag: 0x8cc84ce1, annotations: 0x0, tlName: "reqResultHeader"})
	fillObject("rpcCancelReq#193f1b22", "#193f1b22", &TLItem{tag: 0x193f1b22, annotations: 0x0, tlName: "rpcCancelReq"})
	fillObject("rpcDestActor#7568aabd", "#7568aabd", &TLItem{tag: 0x7568aabd, annotations: 0x0, tlName: "rpcDestActor"})
	fillObject("rpcDestActorFlags#f0a5acf7", "#f0a5acf7", &TLItem{tag: 0xf0a5acf7, annotations: 0x0, tlName: "rpcDestActorFlags"})
	fillObject("rpcDestFlags#e352035e", "#e352035e", &TLItem{tag: 0xe352035e, annotations: 0x0, tlName: "rpcDestFlags"})
	fillObject("rpcInvokeReqExtra#f3ef81a9", "#f3ef81a9", &TLItem{tag: 0xf3ef81a9, annotations: 0x0, tlName: "rpcInvokeReqExtra"})
	fillObject("rpcInvokeReqHeader#2374df3d", "#2374df3d", &TLItem{tag: 0x2374df3d, annotations: 0x0, tlName: "rpcInvokeReqHeader"})
	fillObject("rpcPing#5730a2df", "#5730a2df", &TLItem{tag: 0x5730a2df, annotations: 0x0, tlName: "rpcPing"})
	fillObject("rpcPong#8430eaa7", "#8430eaa7", &TLItem{tag: 0x8430eaa7, annotations: 0x0, tlName: "rpcPong"})
	fillObject("rpcReqResultError#7ae432f5", "#7ae432f5", &TLItem{tag: 0x7ae432f5, annotations: 0x0, tlName: "rpcReqResultError"})
	fillObject("rpcReqResultErrorWrapped#7ae432f6", "#7ae432f6", &TLItem{tag: 0x7ae432f6, annotations: 0x0, tlName: "rpcReqResultErrorWrapped"})
	fillObject("rpcReqResultExtra#c5011709", "#c5011709", &TLItem{tag: 0xc5011709, annotations: 0x0, tlName: "rpcReqResultExtra"})
	fillObject("rpcReqResultHeader#63aeda4e", "#63aeda4e", &TLItem{tag: 0x63aeda4e, annotations: 0x0, tlName: "rpcReqResultHeader"})
	fillObject("rpcServerWantsFin#a8ddbc46", "#a8ddbc46", &TLItem{tag: 0xa8ddbc46, annotations: 0x0, tlName: "rpcServerWantsFin"})
	fillObject("stat#9d56e6b2", "#9d56e6b2", &TLItem{tag: 0x9d56e6b2, annotations: 0x0, tlName: "stat"})
	fillObject("string#b5286e24", "#b5286e24", &TLItem{tag: 0xb5286e24, annotations: 0x0, tlName: "string"})
	fillObject("true#3fedd339", "#3fedd339", &TLItem{tag: 0x3fedd339, annotations: 0x0, tlName: "true"})
}
