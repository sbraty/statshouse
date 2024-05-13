// Copyright 2024 V Kontakte LLC
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

type NetUdpPacketEncHeader struct {
	Flags           uint32
	Time            int32   // Conditional: item.Flags.9
	Version         int32   // Conditional: item.Flags.10
	PacketAckPrefix int32   // Conditional: item.Flags.13
	PacketAckFrom   int32   // Conditional: item.Flags.14
	PacketAckTo     int32   // Conditional: item.Flags.14
	PacketAckSet    []int32 // Conditional: item.Flags.15
	PacketNum       int32   // Conditional: item.Flags.20
	PacketsFrom     int32   // Conditional: item.Flags.21
	PacketsTo       int32   // Conditional: item.Flags.21
	PrevParts       int32   // Conditional: item.Flags.22
	NextParts       int32   // Conditional: item.Flags.23
}

func (NetUdpPacketEncHeader) TLName() string { return "netUdpPacket.encHeader" }
func (NetUdpPacketEncHeader) TLTag() uint32  { return 0x251a7bfd }

func (item *NetUdpPacketEncHeader) SetTime(v int32) {
	item.Time = v
	item.Flags |= 1 << 9
}
func (item *NetUdpPacketEncHeader) ClearTime() {
	item.Time = 0
	item.Flags &^= 1 << 9
}
func (item NetUdpPacketEncHeader) IsSetTime() bool { return item.Flags&(1<<9) != 0 }

func (item *NetUdpPacketEncHeader) SetVersion(v int32) {
	item.Version = v
	item.Flags |= 1 << 10
}
func (item *NetUdpPacketEncHeader) ClearVersion() {
	item.Version = 0
	item.Flags &^= 1 << 10
}
func (item NetUdpPacketEncHeader) IsSetVersion() bool { return item.Flags&(1<<10) != 0 }

func (item *NetUdpPacketEncHeader) SetPacketAckPrefix(v int32) {
	item.PacketAckPrefix = v
	item.Flags |= 1 << 13
}
func (item *NetUdpPacketEncHeader) ClearPacketAckPrefix() {
	item.PacketAckPrefix = 0
	item.Flags &^= 1 << 13
}
func (item NetUdpPacketEncHeader) IsSetPacketAckPrefix() bool { return item.Flags&(1<<13) != 0 }

func (item *NetUdpPacketEncHeader) SetPacketAckFrom(v int32) {
	item.PacketAckFrom = v
	item.Flags |= 1 << 14
}
func (item *NetUdpPacketEncHeader) ClearPacketAckFrom() {
	item.PacketAckFrom = 0
	item.Flags &^= 1 << 14
}
func (item NetUdpPacketEncHeader) IsSetPacketAckFrom() bool { return item.Flags&(1<<14) != 0 }

func (item *NetUdpPacketEncHeader) SetPacketAckTo(v int32) {
	item.PacketAckTo = v
	item.Flags |= 1 << 14
}
func (item *NetUdpPacketEncHeader) ClearPacketAckTo() {
	item.PacketAckTo = 0
	item.Flags &^= 1 << 14
}
func (item NetUdpPacketEncHeader) IsSetPacketAckTo() bool { return item.Flags&(1<<14) != 0 }

func (item *NetUdpPacketEncHeader) SetPacketAckSet(v []int32) {
	item.PacketAckSet = v
	item.Flags |= 1 << 15
}
func (item *NetUdpPacketEncHeader) ClearPacketAckSet() {
	item.PacketAckSet = item.PacketAckSet[:0]
	item.Flags &^= 1 << 15
}
func (item NetUdpPacketEncHeader) IsSetPacketAckSet() bool { return item.Flags&(1<<15) != 0 }

func (item *NetUdpPacketEncHeader) SetPacketNum(v int32) {
	item.PacketNum = v
	item.Flags |= 1 << 20
}
func (item *NetUdpPacketEncHeader) ClearPacketNum() {
	item.PacketNum = 0
	item.Flags &^= 1 << 20
}
func (item NetUdpPacketEncHeader) IsSetPacketNum() bool { return item.Flags&(1<<20) != 0 }

func (item *NetUdpPacketEncHeader) SetPacketsFrom(v int32) {
	item.PacketsFrom = v
	item.Flags |= 1 << 21
}
func (item *NetUdpPacketEncHeader) ClearPacketsFrom() {
	item.PacketsFrom = 0
	item.Flags &^= 1 << 21
}
func (item NetUdpPacketEncHeader) IsSetPacketsFrom() bool { return item.Flags&(1<<21) != 0 }

func (item *NetUdpPacketEncHeader) SetPacketsTo(v int32) {
	item.PacketsTo = v
	item.Flags |= 1 << 21
}
func (item *NetUdpPacketEncHeader) ClearPacketsTo() {
	item.PacketsTo = 0
	item.Flags &^= 1 << 21
}
func (item NetUdpPacketEncHeader) IsSetPacketsTo() bool { return item.Flags&(1<<21) != 0 }

func (item *NetUdpPacketEncHeader) SetPrevParts(v int32) {
	item.PrevParts = v
	item.Flags |= 1 << 22
}
func (item *NetUdpPacketEncHeader) ClearPrevParts() {
	item.PrevParts = 0
	item.Flags &^= 1 << 22
}
func (item NetUdpPacketEncHeader) IsSetPrevParts() bool { return item.Flags&(1<<22) != 0 }

func (item *NetUdpPacketEncHeader) SetNextParts(v int32) {
	item.NextParts = v
	item.Flags |= 1 << 23
}
func (item *NetUdpPacketEncHeader) ClearNextParts() {
	item.NextParts = 0
	item.Flags &^= 1 << 23
}
func (item NetUdpPacketEncHeader) IsSetNextParts() bool { return item.Flags&(1<<23) != 0 }

func (item *NetUdpPacketEncHeader) Reset() {
	item.Flags = 0
	item.Time = 0
	item.Version = 0
	item.PacketAckPrefix = 0
	item.PacketAckFrom = 0
	item.PacketAckTo = 0
	item.PacketAckSet = item.PacketAckSet[:0]
	item.PacketNum = 0
	item.PacketsFrom = 0
	item.PacketsTo = 0
	item.PrevParts = 0
	item.NextParts = 0
}

func (item *NetUdpPacketEncHeader) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.Flags); err != nil {
		return w, err
	}
	if item.Flags&(1<<9) != 0 {
		if w, err = basictl.IntRead(w, &item.Time); err != nil {
			return w, err
		}
	} else {
		item.Time = 0
	}
	if item.Flags&(1<<10) != 0 {
		if w, err = basictl.IntRead(w, &item.Version); err != nil {
			return w, err
		}
	} else {
		item.Version = 0
	}
	if item.Flags&(1<<13) != 0 {
		if w, err = basictl.IntRead(w, &item.PacketAckPrefix); err != nil {
			return w, err
		}
	} else {
		item.PacketAckPrefix = 0
	}
	if item.Flags&(1<<14) != 0 {
		if w, err = basictl.IntRead(w, &item.PacketAckFrom); err != nil {
			return w, err
		}
	} else {
		item.PacketAckFrom = 0
	}
	if item.Flags&(1<<14) != 0 {
		if w, err = basictl.IntRead(w, &item.PacketAckTo); err != nil {
			return w, err
		}
	} else {
		item.PacketAckTo = 0
	}
	if item.Flags&(1<<15) != 0 {
		if w, err = BuiltinVectorIntRead(w, &item.PacketAckSet); err != nil {
			return w, err
		}
	} else {
		item.PacketAckSet = item.PacketAckSet[:0]
	}
	if item.Flags&(1<<20) != 0 {
		if w, err = basictl.IntRead(w, &item.PacketNum); err != nil {
			return w, err
		}
	} else {
		item.PacketNum = 0
	}
	if item.Flags&(1<<21) != 0 {
		if w, err = basictl.IntRead(w, &item.PacketsFrom); err != nil {
			return w, err
		}
	} else {
		item.PacketsFrom = 0
	}
	if item.Flags&(1<<21) != 0 {
		if w, err = basictl.IntRead(w, &item.PacketsTo); err != nil {
			return w, err
		}
	} else {
		item.PacketsTo = 0
	}
	if item.Flags&(1<<22) != 0 {
		if w, err = basictl.IntRead(w, &item.PrevParts); err != nil {
			return w, err
		}
	} else {
		item.PrevParts = 0
	}
	if item.Flags&(1<<23) != 0 {
		if w, err = basictl.IntRead(w, &item.NextParts); err != nil {
			return w, err
		}
	} else {
		item.NextParts = 0
	}
	return w, nil
}

func (item *NetUdpPacketEncHeader) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.Flags)
	if item.Flags&(1<<9) != 0 {
		w = basictl.IntWrite(w, item.Time)
	}
	if item.Flags&(1<<10) != 0 {
		w = basictl.IntWrite(w, item.Version)
	}
	if item.Flags&(1<<13) != 0 {
		w = basictl.IntWrite(w, item.PacketAckPrefix)
	}
	if item.Flags&(1<<14) != 0 {
		w = basictl.IntWrite(w, item.PacketAckFrom)
	}
	if item.Flags&(1<<14) != 0 {
		w = basictl.IntWrite(w, item.PacketAckTo)
	}
	if item.Flags&(1<<15) != 0 {
		if w, err = BuiltinVectorIntWrite(w, item.PacketAckSet); err != nil {
			return w, err
		}
	}
	if item.Flags&(1<<20) != 0 {
		w = basictl.IntWrite(w, item.PacketNum)
	}
	if item.Flags&(1<<21) != 0 {
		w = basictl.IntWrite(w, item.PacketsFrom)
	}
	if item.Flags&(1<<21) != 0 {
		w = basictl.IntWrite(w, item.PacketsTo)
	}
	if item.Flags&(1<<22) != 0 {
		w = basictl.IntWrite(w, item.PrevParts)
	}
	if item.Flags&(1<<23) != 0 {
		w = basictl.IntWrite(w, item.NextParts)
	}
	return w, nil
}

func (item *NetUdpPacketEncHeader) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x251a7bfd); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *NetUdpPacketEncHeader) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x251a7bfd)
	return item.Write(w)
}

func (item NetUdpPacketEncHeader) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *NetUdpPacketEncHeader) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFlagsPresented bool
	var propTimePresented bool
	var propVersionPresented bool
	var propPacketAckPrefixPresented bool
	var propPacketAckFromPresented bool
	var propPacketAckToPresented bool
	var propPacketAckSetPresented bool
	var propPacketNumPresented bool
	var propPacketsFromPresented bool
	var propPacketsToPresented bool
	var propPrevPartsPresented bool
	var propNextPartsPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "flags":
				if propFlagsPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("netUdpPacket.encHeader", "flags")
				}
				if err := Json2ReadUint32(in, &item.Flags); err != nil {
					return err
				}
				propFlagsPresented = true
			case "time":
				if propTimePresented {
					return ErrorInvalidJSONWithDuplicatingKeys("netUdpPacket.encHeader", "time")
				}
				if err := Json2ReadInt32(in, &item.Time); err != nil {
					return err
				}
				propTimePresented = true
			case "version":
				if propVersionPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("netUdpPacket.encHeader", "version")
				}
				if err := Json2ReadInt32(in, &item.Version); err != nil {
					return err
				}
				propVersionPresented = true
			case "packet_ack_prefix":
				if propPacketAckPrefixPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("netUdpPacket.encHeader", "packet_ack_prefix")
				}
				if err := Json2ReadInt32(in, &item.PacketAckPrefix); err != nil {
					return err
				}
				propPacketAckPrefixPresented = true
			case "packet_ack_from":
				if propPacketAckFromPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("netUdpPacket.encHeader", "packet_ack_from")
				}
				if err := Json2ReadInt32(in, &item.PacketAckFrom); err != nil {
					return err
				}
				propPacketAckFromPresented = true
			case "packet_ack_to":
				if propPacketAckToPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("netUdpPacket.encHeader", "packet_ack_to")
				}
				if err := Json2ReadInt32(in, &item.PacketAckTo); err != nil {
					return err
				}
				propPacketAckToPresented = true
			case "packet_ack_set":
				if propPacketAckSetPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("netUdpPacket.encHeader", "packet_ack_set")
				}
				if err := BuiltinVectorIntReadJSON(legacyTypeNames, in, &item.PacketAckSet); err != nil {
					return err
				}
				propPacketAckSetPresented = true
			case "packet_num":
				if propPacketNumPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("netUdpPacket.encHeader", "packet_num")
				}
				if err := Json2ReadInt32(in, &item.PacketNum); err != nil {
					return err
				}
				propPacketNumPresented = true
			case "packets_from":
				if propPacketsFromPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("netUdpPacket.encHeader", "packets_from")
				}
				if err := Json2ReadInt32(in, &item.PacketsFrom); err != nil {
					return err
				}
				propPacketsFromPresented = true
			case "packets_to":
				if propPacketsToPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("netUdpPacket.encHeader", "packets_to")
				}
				if err := Json2ReadInt32(in, &item.PacketsTo); err != nil {
					return err
				}
				propPacketsToPresented = true
			case "prev_parts":
				if propPrevPartsPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("netUdpPacket.encHeader", "prev_parts")
				}
				if err := Json2ReadInt32(in, &item.PrevParts); err != nil {
					return err
				}
				propPrevPartsPresented = true
			case "next_parts":
				if propNextPartsPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("netUdpPacket.encHeader", "next_parts")
				}
				if err := Json2ReadInt32(in, &item.NextParts); err != nil {
					return err
				}
				propNextPartsPresented = true
			default:
				return ErrorInvalidJSONExcessElement("netUdpPacket.encHeader", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propFlagsPresented {
		item.Flags = 0
	}
	if !propTimePresented {
		item.Time = 0
	}
	if !propVersionPresented {
		item.Version = 0
	}
	if !propPacketAckPrefixPresented {
		item.PacketAckPrefix = 0
	}
	if !propPacketAckFromPresented {
		item.PacketAckFrom = 0
	}
	if !propPacketAckToPresented {
		item.PacketAckTo = 0
	}
	if !propPacketAckSetPresented {
		item.PacketAckSet = item.PacketAckSet[:0]
	}
	if !propPacketNumPresented {
		item.PacketNum = 0
	}
	if !propPacketsFromPresented {
		item.PacketsFrom = 0
	}
	if !propPacketsToPresented {
		item.PacketsTo = 0
	}
	if !propPrevPartsPresented {
		item.PrevParts = 0
	}
	if !propNextPartsPresented {
		item.NextParts = 0
	}
	if propTimePresented {
		item.Flags |= 1 << 9
	}
	if propVersionPresented {
		item.Flags |= 1 << 10
	}
	if propPacketAckPrefixPresented {
		item.Flags |= 1 << 13
	}
	if propPacketAckFromPresented {
		item.Flags |= 1 << 14
	}
	if propPacketAckToPresented {
		item.Flags |= 1 << 14
	}
	if propPacketAckSetPresented {
		item.Flags |= 1 << 15
	}
	if propPacketNumPresented {
		item.Flags |= 1 << 20
	}
	if propPacketsFromPresented {
		item.Flags |= 1 << 21
	}
	if propPacketsToPresented {
		item.Flags |= 1 << 21
	}
	if propPrevPartsPresented {
		item.Flags |= 1 << 22
	}
	if propNextPartsPresented {
		item.Flags |= 1 << 23
	}
	return nil
}

func (item *NetUdpPacketEncHeader) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *NetUdpPacketEncHeader) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	backupIndexFlags := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"flags":`...)
	w = basictl.JSONWriteUint32(w, item.Flags)
	if (item.Flags != 0) == false {
		w = w[:backupIndexFlags]
	}
	if item.Flags&(1<<9) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"time":`...)
		w = basictl.JSONWriteInt32(w, item.Time)
	}
	if item.Flags&(1<<10) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"version":`...)
		w = basictl.JSONWriteInt32(w, item.Version)
	}
	if item.Flags&(1<<13) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"packet_ack_prefix":`...)
		w = basictl.JSONWriteInt32(w, item.PacketAckPrefix)
	}
	if item.Flags&(1<<14) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"packet_ack_from":`...)
		w = basictl.JSONWriteInt32(w, item.PacketAckFrom)
	}
	if item.Flags&(1<<14) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"packet_ack_to":`...)
		w = basictl.JSONWriteInt32(w, item.PacketAckTo)
	}
	if item.Flags&(1<<15) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"packet_ack_set":`...)
		if w, err = BuiltinVectorIntWriteJSONOpt(newTypeNames, short, w, item.PacketAckSet); err != nil {
			return w, err
		}
	}
	if item.Flags&(1<<20) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"packet_num":`...)
		w = basictl.JSONWriteInt32(w, item.PacketNum)
	}
	if item.Flags&(1<<21) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"packets_from":`...)
		w = basictl.JSONWriteInt32(w, item.PacketsFrom)
	}
	if item.Flags&(1<<21) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"packets_to":`...)
		w = basictl.JSONWriteInt32(w, item.PacketsTo)
	}
	if item.Flags&(1<<22) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"prev_parts":`...)
		w = basictl.JSONWriteInt32(w, item.PrevParts)
	}
	if item.Flags&(1<<23) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"next_parts":`...)
		w = basictl.JSONWriteInt32(w, item.NextParts)
	}
	return append(w, '}'), nil
}

func (item *NetUdpPacketEncHeader) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *NetUdpPacketEncHeader) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("netUdpPacket.encHeader", err.Error())
	}
	return nil
}
