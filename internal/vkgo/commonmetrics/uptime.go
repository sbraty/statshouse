// Copyright 2024 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package commonmetrics

import (
	"runtime"
	"time"

	"github.com/vkcom/statshouse-go"

	"github.com/vkcom/statshouse/internal/vkgo/commonmetrics/internal/env"
)

type uptime struct {
	heartbeat  string
	commit     string
	commitTime string
	branch     string
	buildID    string
	buildName  string
	buildTime  string
	buildArch  string
	configHash string
}

func (u uptime) toRawTags() statshouse.Tags {
	tags := AttachBase(statshouse.Tags{})
	tags[4] = u.heartbeat
	tags[5] = u.commit
	tags[6] = u.commitTime
	tags[7] = u.branch
	tags[8] = u.buildID
	tags[9] = u.buildName
	tags[10] = u.buildTime
	tags[11] = u.buildArch
	tags[12] = u.configHash

	return tags
}

type uptimeMetricsKeeper struct {
	startTime time.Time
	name      string
}

var (
	uptimeMetrics = &uptimeMetricsKeeper{
		startTime: time.Now(),
		name:      env.FullMetricName("common_uptime"),
	}
)

func (u *uptimeMetricsKeeper) rawTags(name string) statshouse.Tags {
	return uptime{
		heartbeat:  name,
		commit:     parsers.get(Commit)(),
		commitTime: parsers.get(CommitTime)(),
		branch:     parsers.get(Branch)(),
		buildID:    parsers.get(BuildID)(),
		buildName:  parsers.get(BuildName)(),
		buildTime:  parsers.get(BuildTime)(),
		buildArch:  runtime.GOARCH,
		configHash: getConfigHash(),
	}.toRawTags()
}

func (u *uptimeMetricsKeeper) start() {
	statshouse.Metric(u.name, u.rawTags(HeartbeatStart)).Value(time.Since(u.startTime).Seconds())
}

func (u *uptimeMetricsKeeper) Warmup() {
	statshouse.Metric(u.name, u.rawTags(HeartbeatWarmup)).Value(time.Since(u.startTime).Seconds())
}

func (u *uptimeMetricsKeeper) heartbeat(client *statshouse.Client) {
	client.Metric(u.name, u.rawTags(HeartbeatRunning)).Value(time.Since(u.startTime).Seconds())
}

func (u *uptimeMetricsKeeper) shutdown() {
	statshouse.Metric(u.name, u.rawTags(HeartbeatShutdown)).Value(time.Since(u.startTime).Seconds())
}
