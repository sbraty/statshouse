// Copyright 2024 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package factory

import (
	"github.com/vkcom/statshouse/internal/vkgo/binlog/fsbinlog/internal/gen/internal"
	"github.com/vkcom/statshouse/internal/vkgo/binlog/fsbinlog/internal/gen/meta"
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
	meta.SetGlobalFactoryCreateForObject(0x044c644b, func() meta.Object { var ret internal.FsbinlogLevStart; return &ret })
	meta.SetGlobalFactoryCreateForObject(0xb75009a0, func() meta.Object { var ret internal.FsbinlogLevUpgradeToGms; return &ret })
	meta.SetGlobalFactoryCreateForObject(0x6b49d850, func() meta.Object { var ret internal.FsbinlogSnapshotMeta; return &ret })
}
