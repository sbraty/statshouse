// Copyright 2024 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rpc

import (
	"fmt"
	"time"

	"github.com/vkcom/statshouse/internal/vkgo/basictl"
	"github.com/vkcom/statshouse/internal/vkgo/rpc/internal/gen/tl"
)

// Lifecycle of call context:
// After setupCall, context is owned both by 'calls' and 'writeQ': sent=false, stale=false
//    If cancelCall is made in this state, than context is owned by 'writeQ' only: sent=false, stale=true
//    In this case, context is freed during the next acquireBuffer, otherwise
// After acquireBuffer, context is owned by 'calls' only: sent=true, stale=false
//    If cancelCall is made in this case, than context is freed and removed from 'calls'.
// After finishCall, context is (shortly) owned by receiver goroutine, then owned by result channel
//    cancelCall is NOP in this state
//    If receiver reads context from channel, it must pass it to putCallContext
//    If receiver does not read from channel, context is GC-ed with the channel

func preparePacket(req *Request, deadline time.Time) error {
	if !deadline.IsZero() && !req.Extra.IsSetCustomTimeoutMs() {
		// Not as close to actual writing as possible (we want that due to time.Until)
		// If negative already, so be it.
		req.Extra.SetCustomTimeoutMs(int32(time.Until(deadline).Milliseconds()))
	}

	headerBuf := req.Body // move to local var, then back for speed
	req.extraStart = len(headerBuf)

	reqHeader := tl.RpcInvokeReqHeader{QueryId: req.queryID}
	headerBuf = reqHeader.Write(headerBuf)
	switch {
	case req.ActorID != 0 && req.Extra.Flags != 0:
		// extra := tl.RpcDestActorFlags{ActorId: req.ActorID, Extra: req.Extra.RpcInvokeReqExtra}
		// if headerBuf, err = extra.WriteBoxed(headerBuf); err != nil {
		// here we optimize copy of large extra
		headerBuf = basictl.NatWrite(headerBuf, tl.RpcDestActorFlags{}.TLTag())
		headerBuf = basictl.LongWrite(headerBuf, req.ActorID)
		headerBuf = req.Extra.Write(headerBuf)
	case req.Extra.Flags != 0:
		// extra := tl.RpcDestFlags{Extra: req.Extra.RpcInvokeReqExtra}
		// if headerBuf, err = extra.WriteBoxed(headerBuf); err != nil {
		// here we optimize copy of large extra
		headerBuf = basictl.NatWrite(headerBuf, tl.RpcDestFlags{}.TLTag())
		headerBuf = req.Extra.Write(headerBuf)
	case req.ActorID != 0:
		extra := tl.RpcDestActor{ActorId: req.ActorID}
		headerBuf = extra.WriteBoxed(headerBuf)
	}
	if err := validBodyLen(len(headerBuf)); err != nil { // exact
		return err
	}
	req.Body = headerBuf
	return nil
}

func parseResponseExtra(extra *ReqResultExtra, respBody []byte) (_ []byte, err error) {
	var tag uint32
	var afterTag []byte
	extraSet := 0
	for {
		if afterTag, err = basictl.NatRead(respBody, &tag); err != nil {
			return respBody, err
		}
		if (tag != tl.ReqResultHeader{}.TLTag()) {
			break
		}
		// var extra tl.ReqResultHeader
		// if resp.Body, err = extra.Read(afterTag); err != nil {
		// we optimize copy of large extra here
		if respBody, err = extra.Read(afterTag); err != nil {
			return respBody, err
		}
		extraSet++
	}
	if extraSet > 1 {
		return respBody, fmt.Errorf("rpc: ResultExtra set more than once (%d) for result tag #%08d; please report to infrastructure team", extraSet, tag)
	}

	switch tag {
	case tl.ReqError{}.TLTag():
		var rpcErr tl.ReqError
		if respBody, err = rpcErr.Read(afterTag); err != nil {
			return respBody, err
		}
		return respBody, Error{Code: rpcErr.ErrorCode, Description: rpcErr.Error}
	case tl.RpcReqResultError{}.TLTag():
		var rpcErr tl.RpcReqResultError // ignore query ID
		if respBody, err = rpcErr.Read(afterTag); err != nil {
			return respBody, err
		}
		return respBody, Error{Code: rpcErr.ErrorCode, Description: rpcErr.Error}
	case tl.RpcReqResultErrorWrapped{}.TLTag():
		var rpcErr tl.RpcReqResultErrorWrapped
		if respBody, err = rpcErr.Read(afterTag); err != nil {
			return respBody, err
		}
		return respBody, Error{Code: rpcErr.ErrorCode, Description: rpcErr.Error}
	}
	return respBody, nil
}

func (resp *Response) deliverResult(c *Client) {
	cb := resp.cb
	if cb == nil {
		resp.result <- resp // cctx owned by channel
		return
	}
	// TODO - catch panic in callback, write to rare log
	cb(c, resp, resp.err)
}
