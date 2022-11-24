// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package agent

import (
	"context"
	"time"

	"github.com/vkcom/statshouse/internal/data_model"
	"github.com/vkcom/statshouse/internal/data_model/gen2/tlstatshouse"
)

func (s *Shard) IsAlive() bool {
	return s.alive.Load()
}

func (s *Shard) appendlastSendSuccessfulLocked(success bool) int { // returns how many successes in the list
	if len(s.lastSendSuccessful) >= s.config.LivenessResponsesWindowLength {
		s.lastSendSuccessful = append(s.lastSendSuccessful[:0], s.lastSendSuccessful[1:s.config.LivenessResponsesWindowLength]...)
	}
	s.lastSendSuccessful = append(s.lastSendSuccessful, success)
	succ := 0
	for _, ss := range s.lastSendSuccessful {
		if ss {
			succ++
		}
	}
	return succ
}

func (s *Shard) recordSendResult(success bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.alive.Load() {
		return
	}
	succ := s.appendlastSendSuccessfulLocked(success)
	if len(s.lastSendSuccessful) == s.config.LivenessResponsesWindowLength && succ < s.config.LivenessResponsesWindowSuccesses {
		s.alive.Store(false)
		s.lastSendSuccessful = s.lastSendSuccessful[:0]
		s.client.Client.Logf("Aggregator Dead: # of successful recent sends dropped below %d out of %d for shard %d", s.config.LivenessResponsesWindowSuccesses, s.config.LivenessResponsesWindowLength, s.ShardReplicaNum)
	}
}

func (s *Shard) recordKeepLiveResult(success bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	succ := s.appendlastSendSuccessfulLocked(success)
	if succ == s.config.LivenessResponsesWindowLength {
		s.client.Client.Logf("Aggregator Alive: # of successful keepalive sends reached %d out of %d for shard %d", s.config.LivenessResponsesWindowLength, s.config.LivenessResponsesWindowLength, s.ShardReplicaNum)
		s.alive.Store(true)
	}
}

func (s *Shard) sendKeepLive() error {
	now := time.Now()
	ctx, cancel := context.WithDeadline(context.Background(), now.Add(time.Second*60)) // Relatively large timeout here
	defer cancel()

	args := tlstatshouse.SendKeepAlive2Bytes{}
	s.fillProxyHeaderBytes(&args.FieldsMask, &args.Header)
	// It is important that keep alive will not be successful when shards are not configured correctly

	var ret []byte
	err := s.client.SendKeepAlive2Bytes(ctx, args, nil, &ret)
	dur := time.Since(now)
	s.recordKeepLiveResult(err == nil && dur < s.config.KeepAliveSuccessTimeout) // we require strict response time here
	return err
}

func (s *Shard) goLiveChecker() {
	// We have separate loops instead of using flushBuckets agent loop, because sendKeepLive can block on connect for
	// very long time, and it is convenient if it blocks there
	now := time.Now()
	backoffTimeout := time.Duration(0)
	for { // TODO - quit
		tick := time.After(data_model.TillStartOfNextSecond(now))
		now = <-tick // We synchronize with calendar second boundary
		if s.alive.Load() {
			continue
		}
		if err := s.sendKeepLive(); err == nil {
			backoffTimeout = 0
			continue
		}
		backoffTimeout = data_model.NextBackoffDuration(backoffTimeout)
		time.Sleep(backoffTimeout)
	}
}
