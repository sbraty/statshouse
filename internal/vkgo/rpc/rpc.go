// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rpc

import (
	"fmt"
	"time"
)

const (
	flagCRC32C    = uint32(0x00000800) // #define RPC_CRYPTO_USE_CRC32C         0x00000800
	flagCancelReq = uint32(0x00001000) // #define RPC_CRYPTO_RPC_CANCEL_REQ     0x00001000
	FlagP2PHijack = uint32(0x40000000) // #define RPC_CRYPTO_P2P_HIJACK         0x40000000

	packetTypeRPCNonce          = uint32(0x7acb87aa)
	packetTypeRPCHandshake      = uint32(0x7682eef5)
	packetTypeRPCInvokeReq      = uint32(0x2374df3d)
	packetTypeRPCReqResult      = uint32(0x63aeda4e)
	packetTypeRPCReqError       = uint32(0x7ae432f5)
	packetTypeRPCCancelReq      = uint32(0x193f1b22)
	packetTypeRPCServerWantsFin = uint32(0xa8ddbc46)

	destActorTag      = uint32(0x7568aabd) // copy of vktl.MagicTlRpcDestActor
	destFlagsTag      = uint32(0xe352035e) // copy of vktl.MagicTlRpcDestFlags
	destActorFlagsTag = uint32(0xf0a5acf7) // copy of vktl.MagicTlRpcDestActorFlags

	reqResultHeaderTag       = uint32(0x8cc84ce1) // copy of  vktl.MagicTlReqResultHeader
	reqResultErrorTag        = packetTypeRPCReqError
	reqResultErrorWrappedTag = packetTypeRPCReqError + 1 // RPC_REQ_ERROR_WRAPPED

	enginePIDTag          = uint32(0x559d6e36) // copy of vktl.MagicTlEnginePid
	engineStatTag         = uint32(0xefb3c36b) // copy of vktl.MagicTlEngineStat
	engineVersionTag      = uint32(0x1a2e06fa) // copy of vktl.MagicTlEngineVersion
	engineSetVerbosityTag = uint32(0x9d980926) // copy of vktl.MagicTlEngineSetVerbosity
	goPProfTag            = uint32(0xea2876a6) // copy of vktl.MagicTlGoPprof

	// rpc-error-codes.h
	TLErrorSyntax        = -1000 // TL_ERROR_SYNTAX
	TlErrorNoHandler     = -2000 // TL_ERROR_UNKNOWN_FUNCTION_ID
	TlErrorTimeout       = -3000 // TL_ERROR_QUERY_TIMEOUT
	TLErrorNoConnections = -3002 // TL_ERROR_NO_CONNECTIONS
	TlErrorInternal      = -3003 // TL_ERROR_INTERNAL
	TlErrorUnknown       = -4000 // TL_ERROR_UNKNOWN

	clientPingInterval       = 5 * time.Second
	DefaultClientPongTimeout = 10 * time.Second

	DefaultConnTimeoutAccuracy = 100 * time.Millisecond

	// keeping this above 10 seconds helps to avoid disconnecting engines with default 10 seconds ping interval
	maxIdleDuration = 30 * time.Second

	// keeping this less than DefaultConnTimeoutAccuracy away from maxIdleDuration allows to skip setting
	// almost all deadlines (see PacketConn.timeoutAccuracy)
	maxPacketRWTime = maxIdleDuration
)

// NoopLogf is a do-nothing log function
func NoopLogf(string, ...any) {}

type Error struct {
	Code        int32
	Description string
}

func (err Error) Error() string {
	return fmt.Sprintf("RPC error %v: %v", err.Code, err.Description)
}

type NetAddr struct {
	Network string
	Address string
}

func (na NetAddr) String() string {
	return na.Network + "://" + na.Address
}

type packetHeader struct {
	length uint32
	seqNum uint32
	tip    uint32
	crc    uint32
}

func humanByteCountIEC(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}

// rpc-error-codes.h
func IsUserError(errCode int32) bool {
	return errCode > -3000 && errCode <= -1000
}
