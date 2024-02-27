// Copyright 2024 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/url"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/vkcom/statshouse/internal/vkgo/basictl"
	"github.com/vkcom/statshouse/internal/vkgo/rpc/internal/gen/constants"
	"github.com/vkcom/statshouse/internal/vkgo/rpc/internal/gen/tl"
	"github.com/vkcom/statshouse/internal/vkgo/rpc/internal/gen/tlengine"
	"github.com/vkcom/statshouse/internal/vkgo/rpc/internal/gen/tlgo"
	"github.com/vkcom/statshouse/internal/vkgo/srvfunc"
	"github.com/vkcom/statshouse/internal/vkgo/tlpprof" // TODO - modernize pprof package, it is very old
)

func (s *Server) collectStats(localAddr net.Addr) map[string]string {
	cmdline := readCommandLine()
	now := time.Now().Unix()
	uptime := now - int64(s.startTime)
	pid := prepareHandshakePIDServer(localAddr, s.startTime)

	rps := s.statRPS.Load()
	reqTotal := s.statRequestsTotal.Load()
	reqCurrent := s.statRequestsCurrent.Load()
	connTotal := s.statConnectionsTotal.Load()
	connCurrent := s.statConnectionsCurrent.Load()
	requestMem, _ := s.reqMemSem.Observe()
	responseMem, _ := s.respMemSem.Observe()

	workersTotal := s.workerPool.Created()

	gc := srvfunc.GetGCStats()
	gcPausesMs, _ := json.Marshal(gc.LastPausesMs)
	gcPausesMcs, _ := json.Marshal(gc.LastPausesMcs)

	maxCPUUsage := 100 * runtime.GOMAXPROCS(-1)
	avgUsage, curUsage := statCPUInfo.GetSelfCpuUsage()
	avgIdle := 100 * (maxCPUUsage - avgUsage) / maxCPUUsage
	curIdle := 100 * (maxCPUUsage - curUsage) / maxCPUUsage
	if avgIdle < 0 {
		avgIdle = 0
	}
	if curIdle < 0 {
		curIdle = 0
	}

	m := map[string]string{}
	s.opts.StatsHandler(m)

	m["version"] = s.opts.Version
	m["hostname"] = s.statHostname
	m["command_line"] = cmdline
	m["current_time"] = strconv.Itoa(int(now))
	m["start_time"] = strconv.Itoa(int(s.startTime))
	m["uptime"] = strconv.Itoa(int(uptime))
	m["pid"] = strconv.Itoa(int(pidFromPortPid(pid.PortPid)))
	m["PID"] = asTextStat(pid)

	m["qps"] = "0" // memcached protocol
	m["rpc_qps"] = strconv.FormatInt(rps, 10)
	m["queries_total"] = strconv.FormatInt(reqTotal, 10)
	m["queries_active"] = strconv.FormatInt(reqCurrent, 10)
	m["clients_total"] = strconv.FormatInt(connTotal, 10)
	m["clients_active"] = strconv.FormatInt(connCurrent, 10)
	m["request_memory"] = strconv.FormatInt(requestMem, 10)
	m["response_memory"] = strconv.FormatInt(responseMem, 10)
	m["workers_total"] = strconv.FormatInt(int64(workersTotal), 10)

	m["gc_ms"] = strconv.FormatUint(gc.PauseTotalMs, 10)
	m["gc_mcs"] = strconv.FormatUint(gc.PauseTotalMcs, 10)
	m["gc_cpu_usage"] = strconv.FormatFloat(gc.GCCPUFraction, 'f', 5, 32)
	m["gc_pauses"] = string(gcPausesMs)
	m["gc_pauses_mcs"] = string(gcPausesMcs)

	m["average_idle_percent"] = strconv.Itoa(avgIdle)
	m["recent_idle_percent"] = strconv.Itoa(curIdle)

	mem, _ := srvfunc.GetMemStat(0)
	if mem != nil {
		m["vmsize_bytes"] = strconv.FormatUint(mem.Size, 10)
		m["vmrss_bytes"] = strconv.FormatUint(mem.Res, 10)
		m["vmdata_bytes"] = strconv.FormatUint(mem.Data, 10)
	}

	return m
}

func readCommandLine() string {
	buf, _ := os.ReadFile("/proc/self/cmdline")
	return strings.ReplaceAll(string(buf), "\x00", " ")
}

func (s *Server) handleEnginePID(hctx *HandlerContext) (err error) {
	req := tlengine.Pid{}
	if _, err := req.ReadBoxed(hctx.Request); err != nil {
		return err
	}
	pid := prepareHandshakePIDServer(hctx.localAddr, s.startTime)
	hctx.Response, err = req.WriteResult(hctx.Response, pid)
	return err
}

func (s *Server) handleEngineStat(hctx *HandlerContext) error {
	req := tlengine.Stat{}
	if _, err := req.ReadBoxed(hctx.Request); err != nil {
		return err
	}
	stats := s.collectStats(hctx.localAddr)
	keys := sortedStatKeys(stats)
	// hctx.Response, err = req.WriteResult(hctx.Response, pid) - TODO - generate code to write sorted stats

	hctx.Response = basictl.NatWrite(hctx.Response, constants.Stat)
	hctx.Response = basictl.NatWrite(hctx.Response, uint32(len(keys)))
	for _, k := range keys {
		hctx.Response = basictl.StringWriteTruncated(hctx.Response, k)
		hctx.Response = basictl.StringWriteTruncated(hctx.Response, stats[k])
	}
	return nil
}

func (s *Server) handleEngineFilteredStat(hctx *HandlerContext) error {
	req := tlengine.FilteredStat{}
	if _, err := req.ReadBoxed(hctx.Request); err != nil {
		return err
	}
	stats := s.collectStats(hctx.localAddr)
	filterStats(stats, req.StatNames)
	keys := sortedStatKeys(stats)
	// hctx.Response, err = req.WriteResult(hctx.Response, pid) - TODO - generate code to write sorted stats

	hctx.Response = basictl.NatWrite(hctx.Response, constants.Stat)
	hctx.Response = basictl.NatWrite(hctx.Response, uint32(len(keys)))
	for _, k := range keys {
		hctx.Response = basictl.StringWriteTruncated(hctx.Response, k)
		hctx.Response = basictl.StringWriteTruncated(hctx.Response, stats[k])
	}

	return nil
}

func (s *Server) handleEngineVersion(hctx *HandlerContext) (err error) {
	req := tlengine.Version{}
	if _, err := req.ReadBoxed(hctx.Request); err != nil {
		return err
	}
	hctx.Response, err = req.WriteResult(hctx.Response, s.opts.Version)
	return err
}

func (s *Server) handleEngineSetVerbosity(hctx *HandlerContext) (err error) {
	req := tlengine.SetVerbosity{}
	if _, err := req.ReadBoxed(hctx.Request); err != nil {
		return err
	}
	if s.opts.VerbosityHandler == nil {
		return nil
	}
	if err = s.opts.VerbosityHandler(int(req.Verbosity)); err != nil {
		return err
	}
	hctx.Response, err = req.WriteResult(hctx.Response, tl.True{})
	return err
}

func handleEngineSleep(ctx context.Context, hctx *HandlerContext, timeMs int32) (err error) {
	if hctx.timeout != 0 {
		deadline := hctx.RequestTime.Add(hctx.timeout)
		dt := time.Since(deadline)
		if dt >= 0 {
			return Error{
				Code:        TlErrorTimeout,
				Description: fmt.Sprintf("RPC query timeout (%v after deadline)", dt),
			}
		}

		var cancel context.CancelFunc
		ctx, cancel = context.WithDeadline(ctx, deadline)
		defer cancel()
	}

	select {
	case <-time.After(time.Duration(timeMs) * time.Millisecond):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s *Server) handleEngineSleep(ctx context.Context, hctx *HandlerContext) (err error) {
	req := tlengine.Sleep{}
	if _, err := req.ReadBoxed(hctx.Request); err != nil {
		return err
	}
	if err := handleEngineSleep(ctx, hctx, req.TimeMs); err != nil {
		return err
	}
	hctx.Response, err = req.WriteResult(hctx.Response, true)
	return err
}

func (s *Server) handleEngineAsyncSleep(ctx context.Context, hctx *HandlerContext) (err error) {
	req := tlengine.AsyncSleep{}
	if _, err := req.ReadBoxed(hctx.Request); err != nil {
		return err
	}
	if err := handleEngineSleep(ctx, hctx, req.TimeMs); err != nil {
		return err
	}
	hctx.Response, err = req.WriteResult(hctx.Response, true)
	return err
}

func (s *Server) respondWithMemcachedVersion(conn *PacketConn) {
	resp := fmt.Sprintf("VERSION %v\r\n", s.opts.Version)

	err := s.sendMemcachedResponse(conn, "version", []byte(resp))
	if err != nil {
		s.rareLog(&s.lastOtherLog, err.Error())
	}
}

func (s *Server) respondWithMemcachedStats(conn *PacketConn) {
	stats := s.collectStats(conn.conn.LocalAddr())
	resp := marshalMemcachedStats(stats)

	err := s.sendMemcachedResponse(conn, "stats", resp)
	if err != nil {
		s.rareLog(&s.lastOtherLog, "%v", err)
	}
}

func (s *Server) sendMemcachedResponse(c *PacketConn, name string, resp []byte) error {
	err := c.conn.SetWriteDeadline(time.Now().Add(DefaultPacketTimeout))
	if err != nil {
		return fmt.Errorf("rpc: failed to set write deadline for memcached %v for %v: %w", name, c.remoteAddr, err)
	}

	_, err = c.conn.Write(resp)
	if err != nil {
		return fmt.Errorf("rpc: failed to write memcached %v for %v: %w", name, c.remoteAddr, err)
	}

	return nil
}

func sortedStatKeys(stats map[string]string) []string {
	var keys []string
	for k, v := range stats {
		if v != "" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	return keys
}

func filterStats(stats map[string]string, names []string) {
	m := map[string]struct{}{}
	for _, name := range names {
		m[name] = struct{}{}
	}

	for key := range stats {
		if _, ok := m[key]; !ok {
			delete(stats, key)
		}
	}
}

func marshalMemcachedStats(stats map[string]string) []byte {
	keys := sortedStatKeys(stats)

	var buf bytes.Buffer
	for _, k := range keys {
		buf.WriteString(k)
		buf.WriteString(" ")
		buf.WriteString(stats[k])
		buf.WriteString("\r\n")
	}
	buf.WriteString("END\r\n")

	return buf.Bytes()
}

func (s *Server) handleGoPProf(hctx *HandlerContext) (err error) {
	req := tlgo.Pprof{}
	if _, err := req.ReadBoxed(hctx.Request); err != nil {
		return err
	}
	u, err := url.Parse(req.Params)
	if err != nil {
		return fmt.Errorf("failed to parse pprof URL: %w", err)
	}

	var buf bytes.Buffer
	switch u.Path {
	case "/debug/pprof/cmdline":
		tlpprof.Cmdline(&buf)
	case "/debug/pprof/profile":
		err = tlpprof.Profile(&buf, u)
	case "/debug/pprof/trace":
		err = tlpprof.Trace(&buf, u)
	case "/debug/pprof/allocs",
		"/debug/pprof/block",
		"/debug/pprof/goroutine",
		"/debug/pprof/heap",
		"/debug/pprof/mutex",
		"/debug/pprof/threadcreate":
		name := strings.TrimPrefix(u.Path, "/debug/pprof/")
		err = tlpprof.Index(&buf, name, u)
	default:
		err = fmt.Errorf("unsupported pprof URL %q", req.Params)
	}
	if err != nil {
		return err
	}
	hctx.Response, err = req.WriteResult(hctx.Response, buf.String())
	return err
}
