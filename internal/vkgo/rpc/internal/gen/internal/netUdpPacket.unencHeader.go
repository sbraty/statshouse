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

type NetUdpPacketUnencHeader struct {
	Flags      uint32
	RemotePid  NetPid    // Conditional: item.Flags.0
	LocalPid   NetPid    // Conditional: item.Flags.0
	Generation int32     // Conditional: item.Flags.0
	PidHash    int64     // Conditional: item.Flags.2
	CryptoInit int32     // Conditional: item.Flags.3
	RandomKey  [8]uint32 // Conditional: item.Flags.5
}

func (NetUdpPacketUnencHeader) TLName() string { return "netUdpPacket.unencHeader" }
func (NetUdpPacketUnencHeader) TLTag() uint32  { return 0xa8e945 }

func (item *NetUdpPacketUnencHeader) SetRemotePid(v NetPid) {
	item.RemotePid = v
	item.Flags |= 1 << 0
}
func (item *NetUdpPacketUnencHeader) ClearRemotePid() {
	item.RemotePid.Reset()
	item.Flags &^= 1 << 0
}
func (item NetUdpPacketUnencHeader) IsSetRemotePid() bool { return item.Flags&(1<<0) != 0 }

func (item *NetUdpPacketUnencHeader) SetLocalPid(v NetPid) {
	item.LocalPid = v
	item.Flags |= 1 << 0
}
func (item *NetUdpPacketUnencHeader) ClearLocalPid() {
	item.LocalPid.Reset()
	item.Flags &^= 1 << 0
}
func (item NetUdpPacketUnencHeader) IsSetLocalPid() bool { return item.Flags&(1<<0) != 0 }

func (item *NetUdpPacketUnencHeader) SetGeneration(v int32) {
	item.Generation = v
	item.Flags |= 1 << 0
}
func (item *NetUdpPacketUnencHeader) ClearGeneration() {
	item.Generation = 0
	item.Flags &^= 1 << 0
}
func (item NetUdpPacketUnencHeader) IsSetGeneration() bool { return item.Flags&(1<<0) != 0 }

func (item *NetUdpPacketUnencHeader) SetPidHash(v int64) {
	item.PidHash = v
	item.Flags |= 1 << 2
}
func (item *NetUdpPacketUnencHeader) ClearPidHash() {
	item.PidHash = 0
	item.Flags &^= 1 << 2
}
func (item NetUdpPacketUnencHeader) IsSetPidHash() bool { return item.Flags&(1<<2) != 0 }

func (item *NetUdpPacketUnencHeader) SetCryptoInit(v int32) {
	item.CryptoInit = v
	item.Flags |= 1 << 3
}
func (item *NetUdpPacketUnencHeader) ClearCryptoInit() {
	item.CryptoInit = 0
	item.Flags &^= 1 << 3
}
func (item NetUdpPacketUnencHeader) IsSetCryptoInit() bool { return item.Flags&(1<<3) != 0 }

func (item *NetUdpPacketUnencHeader) SetRandomKey(v [8]uint32) {
	item.RandomKey = v
	item.Flags |= 1 << 5
}
func (item *NetUdpPacketUnencHeader) ClearRandomKey() {
	BuiltinTuple8Reset(&item.RandomKey)
	item.Flags &^= 1 << 5
}
func (item NetUdpPacketUnencHeader) IsSetRandomKey() bool { return item.Flags&(1<<5) != 0 }

func (item *NetUdpPacketUnencHeader) Reset() {
	item.Flags = 0
	item.RemotePid.Reset()
	item.LocalPid.Reset()
	item.Generation = 0
	item.PidHash = 0
	item.CryptoInit = 0
	BuiltinTuple8Reset(&item.RandomKey)
}

func (item *NetUdpPacketUnencHeader) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.Flags); err != nil {
		return w, err
	}
	if item.Flags&(1<<0) != 0 {
		if w, err = item.RemotePid.Read(w); err != nil {
			return w, err
		}
	} else {
		item.RemotePid.Reset()
	}
	if item.Flags&(1<<0) != 0 {
		if w, err = item.LocalPid.Read(w); err != nil {
			return w, err
		}
	} else {
		item.LocalPid.Reset()
	}
	if item.Flags&(1<<0) != 0 {
		if w, err = basictl.IntRead(w, &item.Generation); err != nil {
			return w, err
		}
	} else {
		item.Generation = 0
	}
	if item.Flags&(1<<2) != 0 {
		if w, err = basictl.LongRead(w, &item.PidHash); err != nil {
			return w, err
		}
	} else {
		item.PidHash = 0
	}
	if item.Flags&(1<<3) != 0 {
		if w, err = basictl.IntRead(w, &item.CryptoInit); err != nil {
			return w, err
		}
	} else {
		item.CryptoInit = 0
	}
	if item.Flags&(1<<5) != 0 {
		if w, err = BuiltinTuple8Read(w, &item.RandomKey); err != nil {
			return w, err
		}
	} else {
		BuiltinTuple8Reset(&item.RandomKey)
	}
	return w, nil
}

func (item *NetUdpPacketUnencHeader) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.Flags)
	if item.Flags&(1<<0) != 0 {
		if w, err = item.RemotePid.Write(w); err != nil {
			return w, err
		}
	}
	if item.Flags&(1<<0) != 0 {
		if w, err = item.LocalPid.Write(w); err != nil {
			return w, err
		}
	}
	if item.Flags&(1<<0) != 0 {
		w = basictl.IntWrite(w, item.Generation)
	}
	if item.Flags&(1<<2) != 0 {
		w = basictl.LongWrite(w, item.PidHash)
	}
	if item.Flags&(1<<3) != 0 {
		w = basictl.IntWrite(w, item.CryptoInit)
	}
	if item.Flags&(1<<5) != 0 {
		if w, err = BuiltinTuple8Write(w, &item.RandomKey); err != nil {
			return w, err
		}
	}
	return w, nil
}

func (item *NetUdpPacketUnencHeader) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xa8e945); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *NetUdpPacketUnencHeader) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xa8e945)
	return item.Write(w)
}

func (item NetUdpPacketUnencHeader) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item *NetUdpPacketUnencHeader) ReadJSONLegacy(legacyTypeNames bool, j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("netUdpPacket.unencHeader", "expected json object")
	}
	_jFlags := _jm["flags"]
	delete(_jm, "flags")
	if err := JsonReadUint32(_jFlags, &item.Flags); err != nil {
		return err
	}
	_jRemotePid := _jm["remote_pid"]
	delete(_jm, "remote_pid")
	_jLocalPid := _jm["local_pid"]
	delete(_jm, "local_pid")
	_jGeneration := _jm["generation"]
	delete(_jm, "generation")
	_jPidHash := _jm["pid_hash"]
	delete(_jm, "pid_hash")
	_jCryptoInit := _jm["crypto_init"]
	delete(_jm, "crypto_init")
	_jRandomKey := _jm["random_key"]
	delete(_jm, "random_key")
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("netUdpPacket.unencHeader", k)
	}
	if _jRemotePid != nil {
		item.Flags |= 1 << 0
	}
	if _jLocalPid != nil {
		item.Flags |= 1 << 0
	}
	if _jGeneration != nil {
		item.Flags |= 1 << 0
	}
	if _jPidHash != nil {
		item.Flags |= 1 << 2
	}
	if _jCryptoInit != nil {
		item.Flags |= 1 << 3
	}
	if _jRandomKey != nil {
		item.Flags |= 1 << 5
	}
	if _jRemotePid != nil {
		if err := item.RemotePid.ReadJSONLegacy(legacyTypeNames, _jRemotePid); err != nil {
			return err
		}
	} else {
		item.RemotePid.Reset()
	}
	if _jLocalPid != nil {
		if err := item.LocalPid.ReadJSONLegacy(legacyTypeNames, _jLocalPid); err != nil {
			return err
		}
	} else {
		item.LocalPid.Reset()
	}
	if _jGeneration != nil {
		if err := JsonReadInt32(_jGeneration, &item.Generation); err != nil {
			return err
		}
	} else {
		item.Generation = 0
	}
	if _jPidHash != nil {
		if err := JsonReadInt64(_jPidHash, &item.PidHash); err != nil {
			return err
		}
	} else {
		item.PidHash = 0
	}
	if _jCryptoInit != nil {
		if err := JsonReadInt32(_jCryptoInit, &item.CryptoInit); err != nil {
			return err
		}
	} else {
		item.CryptoInit = 0
	}
	if _jRandomKey != nil {
		if err := BuiltinTuple8ReadJSONLegacy(legacyTypeNames, _jRandomKey, &item.RandomKey); err != nil {
			return err
		}
	} else {
		BuiltinTuple8Reset(&item.RandomKey)
	}
	return nil
}

func (item *NetUdpPacketUnencHeader) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFlagsPresented bool
	var propRemotePidPresented bool
	var propLocalPidPresented bool
	var propGenerationPresented bool
	var propPidHashPresented bool
	var propCryptoInitPresented bool
	var propRandomKeyPresented bool

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
					return ErrorInvalidJSONWithDuplicatingKeys("netUdpPacket.unencHeader", "flags")
				}
				if err := Json2ReadUint32(in, &item.Flags); err != nil {
					return err
				}
				propFlagsPresented = true
			case "remote_pid":
				if propRemotePidPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("netUdpPacket.unencHeader", "remote_pid")
				}
				if err := item.RemotePid.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propRemotePidPresented = true
			case "local_pid":
				if propLocalPidPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("netUdpPacket.unencHeader", "local_pid")
				}
				if err := item.LocalPid.ReadJSON(legacyTypeNames, in); err != nil {
					return err
				}
				propLocalPidPresented = true
			case "generation":
				if propGenerationPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("netUdpPacket.unencHeader", "generation")
				}
				if err := Json2ReadInt32(in, &item.Generation); err != nil {
					return err
				}
				propGenerationPresented = true
			case "pid_hash":
				if propPidHashPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("netUdpPacket.unencHeader", "pid_hash")
				}
				if err := Json2ReadInt64(in, &item.PidHash); err != nil {
					return err
				}
				propPidHashPresented = true
			case "crypto_init":
				if propCryptoInitPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("netUdpPacket.unencHeader", "crypto_init")
				}
				if err := Json2ReadInt32(in, &item.CryptoInit); err != nil {
					return err
				}
				propCryptoInitPresented = true
			case "random_key":
				if propRandomKeyPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("netUdpPacket.unencHeader", "random_key")
				}
				if err := BuiltinTuple8ReadJSON(legacyTypeNames, in, &item.RandomKey); err != nil {
					return err
				}
				propRandomKeyPresented = true
			default:
				return ErrorInvalidJSONExcessElement("netUdpPacket.unencHeader", key)
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
	if !propRemotePidPresented {
		item.RemotePid.Reset()
	}
	if !propLocalPidPresented {
		item.LocalPid.Reset()
	}
	if !propGenerationPresented {
		item.Generation = 0
	}
	if !propPidHashPresented {
		item.PidHash = 0
	}
	if !propCryptoInitPresented {
		item.CryptoInit = 0
	}
	if !propRandomKeyPresented {
		BuiltinTuple8Reset(&item.RandomKey)
	}
	if propRemotePidPresented {
		item.Flags |= 1 << 0
	}
	if propLocalPidPresented {
		item.Flags |= 1 << 0
	}
	if propGenerationPresented {
		item.Flags |= 1 << 0
	}
	if propPidHashPresented {
		item.Flags |= 1 << 2
	}
	if propCryptoInitPresented {
		item.Flags |= 1 << 3
	}
	if propRandomKeyPresented {
		item.Flags |= 1 << 5
	}
	return nil
}

func (item *NetUdpPacketUnencHeader) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w)
}
func (item *NetUdpPacketUnencHeader) WriteJSONOpt(newTypeNames bool, short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.Flags != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"flags":`...)
		w = basictl.JSONWriteUint32(w, item.Flags)
	}
	if item.Flags&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"remote_pid":`...)
		if w, err = item.RemotePid.WriteJSONOpt(newTypeNames, short, w); err != nil {
			return w, err
		}
	}
	if item.Flags&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"local_pid":`...)
		if w, err = item.LocalPid.WriteJSONOpt(newTypeNames, short, w); err != nil {
			return w, err
		}
	}
	if item.Flags&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"generation":`...)
		w = basictl.JSONWriteInt32(w, item.Generation)
	}
	if item.Flags&(1<<2) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"pid_hash":`...)
		w = basictl.JSONWriteInt64(w, item.PidHash)
	}
	if item.Flags&(1<<3) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"crypto_init":`...)
		w = basictl.JSONWriteInt32(w, item.CryptoInit)
	}
	if item.Flags&(1<<5) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"random_key":`...)
		if w, err = BuiltinTuple8WriteJSONOpt(newTypeNames, short, w, &item.RandomKey); err != nil {
			return w, err
		}
	}
	return append(w, '}'), nil
}

func (item *NetUdpPacketUnencHeader) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *NetUdpPacketUnencHeader) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("netUdpPacket.unencHeader", err.Error())
	}
	if err = item.ReadJSONLegacy(true, j); err != nil {
		return ErrorInvalidJSON("netUdpPacket.unencHeader", err.Error())
	}
	return nil
}
