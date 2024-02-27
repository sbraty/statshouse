// Copyright 2023 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlstatshouse

import (
	"context"

	"github.com/vkcom/statshouse/internal/data_model/gen2/internal"
	"github.com/vkcom/statshouse/internal/vkgo/basictl"
	"github.com/vkcom/statshouse/internal/vkgo/rpc"
)

type (
	AddMetricsBatch                   = internal.StatshouseAddMetricsBatch
	AddMetricsBatchBytes              = internal.StatshouseAddMetricsBatchBytes
	AutoCreate                        = internal.StatshouseAutoCreate
	AutoCreateBytes                   = internal.StatshouseAutoCreateBytes
	Centroid                          = internal.StatshouseCentroid
	CommonProxyHeader                 = internal.StatshouseCommonProxyHeader
	CommonProxyHeaderBytes            = internal.StatshouseCommonProxyHeaderBytes
	GetConfig2                        = internal.StatshouseGetConfig2
	GetConfig2Bytes                   = internal.StatshouseGetConfig2Bytes
	GetConfigResult                   = internal.StatshouseGetConfigResult
	GetConfigResultBytes              = internal.StatshouseGetConfigResultBytes
	GetMetrics3                       = internal.StatshouseGetMetrics3
	GetMetrics3Bytes                  = internal.StatshouseGetMetrics3Bytes
	GetMetricsResult                  = internal.StatshouseGetMetricsResult
	GetMetricsResultBytes             = internal.StatshouseGetMetricsResultBytes
	GetTagMapping2                    = internal.StatshouseGetTagMapping2
	GetTagMapping2Bytes               = internal.StatshouseGetTagMapping2Bytes
	GetTagMappingBootstrap            = internal.StatshouseGetTagMappingBootstrap
	GetTagMappingBootstrapBytes       = internal.StatshouseGetTagMappingBootstrapBytes
	GetTagMappingBootstrapResult      = internal.StatshouseGetTagMappingBootstrapResult
	GetTagMappingBootstrapResultBytes = internal.StatshouseGetTagMappingBootstrapResultBytes
	GetTagMappingResult               = internal.StatshouseGetTagMappingResult
	GetTargets2                       = internal.StatshouseGetTargets2
	GetTargets2Bytes                  = internal.StatshouseGetTargets2Bytes
	GetTargetsResult                  = internal.StatshouseGetTargetsResult
	GetTargetsResultBytes             = internal.StatshouseGetTargetsResultBytes
	IngestionStatus2                  = internal.StatshouseIngestionStatus2
	Mapping                           = internal.StatshouseMapping
	MappingBytes                      = internal.StatshouseMappingBytes
	Metric                            = internal.StatshouseMetric
	MetricBytes                       = internal.StatshouseMetricBytes
	MultiItem                         = internal.StatshouseMultiItem
	MultiItemBytes                    = internal.StatshouseMultiItemBytes
	MultiValue                        = internal.StatshouseMultiValue
	MultiValueBytes                   = internal.StatshouseMultiValueBytes
	PromTarget                        = internal.StatshousePromTarget
	PromTargetBytes                   = internal.StatshousePromTargetBytes
	PutTagMappingBootstrapResult      = internal.StatshousePutTagMappingBootstrapResult
	SampleFactor                      = internal.StatshouseSampleFactor
	SendKeepAlive2                    = internal.StatshouseSendKeepAlive2
	SendKeepAlive2Bytes               = internal.StatshouseSendKeepAlive2Bytes
	SendSourceBucket2                 = internal.StatshouseSendSourceBucket2
	SendSourceBucket2Bytes            = internal.StatshouseSendSourceBucket2Bytes
	SourceBucket2                     = internal.StatshouseSourceBucket2
	SourceBucket2Bytes                = internal.StatshouseSourceBucket2Bytes
	TestConnection2                   = internal.StatshouseTestConnection2
	TestConnection2Bytes              = internal.StatshouseTestConnection2Bytes
	TopElement                        = internal.StatshouseTopElement
	TopElementBytes                   = internal.StatshouseTopElementBytes
)

type Client struct {
	Client  *rpc.Client
	Network string // should be either "tcp4" or "unix"
	Address string
	ActorID int64 // should be non-zero when using rpc-proxy
}

func (c *Client) AddMetricsBatchBytes(ctx context.Context, args AddMetricsBatchBytes, extra *rpc.InvokeReqExtra, ret *internal.True) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "statshouse.addMetricsBatch"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxed(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("statshouse.addMetricsBatch", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("statshouse.addMetricsBatch", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("statshouse.addMetricsBatch", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) AddMetricsBatch(ctx context.Context, args AddMetricsBatch, extra *rpc.InvokeReqExtra, ret *internal.True) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "statshouse.addMetricsBatch"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxed(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("statshouse.addMetricsBatch", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("statshouse.addMetricsBatch", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("statshouse.addMetricsBatch", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) AutoCreateBytes(ctx context.Context, args AutoCreateBytes, extra *rpc.InvokeReqExtra, ret *internal.True) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "statshouse.autoCreate"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxed(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("statshouse.autoCreate", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("statshouse.autoCreate", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("statshouse.autoCreate", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) AutoCreate(ctx context.Context, args AutoCreate, extra *rpc.InvokeReqExtra, ret *internal.True) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "statshouse.autoCreate"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxed(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("statshouse.autoCreate", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("statshouse.autoCreate", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("statshouse.autoCreate", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) GetConfig2Bytes(ctx context.Context, args GetConfig2Bytes, extra *rpc.InvokeReqExtra, ret *GetConfigResultBytes) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "statshouse.getConfig2"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxed(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("statshouse.getConfig2", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("statshouse.getConfig2", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("statshouse.getConfig2", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) GetConfig2(ctx context.Context, args GetConfig2, extra *rpc.InvokeReqExtra, ret *GetConfigResult) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "statshouse.getConfig2"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxed(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("statshouse.getConfig2", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("statshouse.getConfig2", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("statshouse.getConfig2", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) GetMetrics3Bytes(ctx context.Context, args GetMetrics3Bytes, extra *rpc.InvokeReqExtra, ret *internal.MetadataGetJournalResponsenewBytes) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "statshouse.getMetrics3"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxed(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("statshouse.getMetrics3", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("statshouse.getMetrics3", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("statshouse.getMetrics3", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) GetMetrics3(ctx context.Context, args GetMetrics3, extra *rpc.InvokeReqExtra, ret *internal.MetadataGetJournalResponsenew) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "statshouse.getMetrics3"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxed(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("statshouse.getMetrics3", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("statshouse.getMetrics3", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("statshouse.getMetrics3", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) GetTagMapping2Bytes(ctx context.Context, args GetTagMapping2Bytes, extra *rpc.InvokeReqExtra, ret *GetTagMappingResult) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "statshouse.getTagMapping2"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxed(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("statshouse.getTagMapping2", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("statshouse.getTagMapping2", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("statshouse.getTagMapping2", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) GetTagMapping2(ctx context.Context, args GetTagMapping2, extra *rpc.InvokeReqExtra, ret *GetTagMappingResult) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "statshouse.getTagMapping2"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxed(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("statshouse.getTagMapping2", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("statshouse.getTagMapping2", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("statshouse.getTagMapping2", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) GetTagMappingBootstrapBytes(ctx context.Context, args GetTagMappingBootstrapBytes, extra *rpc.InvokeReqExtra, ret *GetTagMappingBootstrapResultBytes) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "statshouse.getTagMappingBootstrap"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxed(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("statshouse.getTagMappingBootstrap", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("statshouse.getTagMappingBootstrap", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("statshouse.getTagMappingBootstrap", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) GetTagMappingBootstrap(ctx context.Context, args GetTagMappingBootstrap, extra *rpc.InvokeReqExtra, ret *GetTagMappingBootstrapResult) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "statshouse.getTagMappingBootstrap"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxed(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("statshouse.getTagMappingBootstrap", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("statshouse.getTagMappingBootstrap", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("statshouse.getTagMappingBootstrap", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) GetTargets2Bytes(ctx context.Context, args GetTargets2Bytes, extra *rpc.InvokeReqExtra, ret *GetTargetsResultBytes) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "statshouse.getTargets2"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxed(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("statshouse.getTargets2", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("statshouse.getTargets2", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("statshouse.getTargets2", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) GetTargets2(ctx context.Context, args GetTargets2, extra *rpc.InvokeReqExtra, ret *GetTargetsResult) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "statshouse.getTargets2"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxed(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("statshouse.getTargets2", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("statshouse.getTargets2", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("statshouse.getTargets2", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) SendKeepAlive2Bytes(ctx context.Context, args SendKeepAlive2Bytes, extra *rpc.InvokeReqExtra, ret *[]byte) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "statshouse.sendKeepAlive2"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxed(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("statshouse.sendKeepAlive2", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("statshouse.sendKeepAlive2", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("statshouse.sendKeepAlive2", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) SendKeepAlive2(ctx context.Context, args SendKeepAlive2, extra *rpc.InvokeReqExtra, ret *string) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "statshouse.sendKeepAlive2"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxed(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("statshouse.sendKeepAlive2", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("statshouse.sendKeepAlive2", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("statshouse.sendKeepAlive2", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) SendSourceBucket2Bytes(ctx context.Context, args SendSourceBucket2Bytes, extra *rpc.InvokeReqExtra, ret *[]byte) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "statshouse.sendSourceBucket2"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxed(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("statshouse.sendSourceBucket2", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("statshouse.sendSourceBucket2", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("statshouse.sendSourceBucket2", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) SendSourceBucket2(ctx context.Context, args SendSourceBucket2, extra *rpc.InvokeReqExtra, ret *string) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "statshouse.sendSourceBucket2"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxed(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("statshouse.sendSourceBucket2", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("statshouse.sendSourceBucket2", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("statshouse.sendSourceBucket2", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) TestConnection2Bytes(ctx context.Context, args TestConnection2Bytes, extra *rpc.InvokeReqExtra, ret *[]byte) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "statshouse.testConnection2"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxed(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("statshouse.testConnection2", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("statshouse.testConnection2", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("statshouse.testConnection2", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) TestConnection2(ctx context.Context, args TestConnection2, extra *rpc.InvokeReqExtra, ret *string) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "statshouse.testConnection2"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxed(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("statshouse.testConnection2", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("statshouse.testConnection2", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("statshouse.testConnection2", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

type Handler struct {
	AddMetricsBatch        func(ctx context.Context, args AddMetricsBatch) (internal.True, error)                       // statshouse.addMetricsBatch
	AutoCreate             func(ctx context.Context, args AutoCreate) (internal.True, error)                            // statshouse.autoCreate
	GetConfig2             func(ctx context.Context, args GetConfig2) (GetConfigResult, error)                          // statshouse.getConfig2
	GetMetrics3            func(ctx context.Context, args GetMetrics3) (internal.MetadataGetJournalResponsenew, error)  // statshouse.getMetrics3
	GetTagMapping2         func(ctx context.Context, args GetTagMapping2) (GetTagMappingResult, error)                  // statshouse.getTagMapping2
	GetTagMappingBootstrap func(ctx context.Context, args GetTagMappingBootstrap) (GetTagMappingBootstrapResult, error) // statshouse.getTagMappingBootstrap
	GetTargets2            func(ctx context.Context, args GetTargets2) (GetTargetsResult, error)                        // statshouse.getTargets2
	SendKeepAlive2         func(ctx context.Context, args SendKeepAlive2) (string, error)                               // statshouse.sendKeepAlive2
	SendSourceBucket2      func(ctx context.Context, args SendSourceBucket2) (string, error)                            // statshouse.sendSourceBucket2
	TestConnection2        func(ctx context.Context, args TestConnection2) (string, error)                              // statshouse.testConnection2

	RawAddMetricsBatch        func(ctx context.Context, hctx *rpc.HandlerContext) error // statshouse.addMetricsBatch
	RawAutoCreate             func(ctx context.Context, hctx *rpc.HandlerContext) error // statshouse.autoCreate
	RawGetConfig2             func(ctx context.Context, hctx *rpc.HandlerContext) error // statshouse.getConfig2
	RawGetMetrics3            func(ctx context.Context, hctx *rpc.HandlerContext) error // statshouse.getMetrics3
	RawGetTagMapping2         func(ctx context.Context, hctx *rpc.HandlerContext) error // statshouse.getTagMapping2
	RawGetTagMappingBootstrap func(ctx context.Context, hctx *rpc.HandlerContext) error // statshouse.getTagMappingBootstrap
	RawGetTargets2            func(ctx context.Context, hctx *rpc.HandlerContext) error // statshouse.getTargets2
	RawSendKeepAlive2         func(ctx context.Context, hctx *rpc.HandlerContext) error // statshouse.sendKeepAlive2
	RawSendSourceBucket2      func(ctx context.Context, hctx *rpc.HandlerContext) error // statshouse.sendSourceBucket2
	RawTestConnection2        func(ctx context.Context, hctx *rpc.HandlerContext) error // statshouse.testConnection2
}

func (h *Handler) Handle(ctx context.Context, hctx *rpc.HandlerContext) (err error) {
	tag, r, _ := basictl.NatReadTag(hctx.Request) // keep hctx.Request intact for handler chaining
	switch tag {
	case 0x56580239: // statshouse.addMetricsBatch
		hctx.RequestFunctionName = "statshouse.addMetricsBatch"
		if h.RawAddMetricsBatch != nil {
			hctx.Request = r
			err = h.RawAddMetricsBatch(ctx, hctx)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("statshouse.addMetricsBatch", err)
			}
			return nil
		}
		if h.AddMetricsBatch != nil {
			var args AddMetricsBatch
			if _, err = args.Read(r); err != nil {
				return internal.ErrorServerRead("statshouse.addMetricsBatch", err)
			}
			ctx = hctx.WithContext(ctx)
			ret, err := h.AddMetricsBatch(ctx, args)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("statshouse.addMetricsBatch", err)
			}
			if hctx.Response, err = args.WriteResult(hctx.Response, ret); err != nil {
				return internal.ErrorServerWriteResult("statshouse.addMetricsBatch", err)
			}
			return nil
		}
	case 0x28bea524: // statshouse.autoCreate
		hctx.RequestFunctionName = "statshouse.autoCreate"
		if h.RawAutoCreate != nil {
			hctx.Request = r
			err = h.RawAutoCreate(ctx, hctx)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("statshouse.autoCreate", err)
			}
			return nil
		}
		if h.AutoCreate != nil {
			var args AutoCreate
			if _, err = args.Read(r); err != nil {
				return internal.ErrorServerRead("statshouse.autoCreate", err)
			}
			ctx = hctx.WithContext(ctx)
			ret, err := h.AutoCreate(ctx, args)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("statshouse.autoCreate", err)
			}
			if hctx.Response, err = args.WriteResult(hctx.Response, ret); err != nil {
				return internal.ErrorServerWriteResult("statshouse.autoCreate", err)
			}
			return nil
		}
	case 0x4285ff57: // statshouse.getConfig2
		hctx.RequestFunctionName = "statshouse.getConfig2"
		if h.RawGetConfig2 != nil {
			hctx.Request = r
			err = h.RawGetConfig2(ctx, hctx)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("statshouse.getConfig2", err)
			}
			return nil
		}
		if h.GetConfig2 != nil {
			var args GetConfig2
			if _, err = args.Read(r); err != nil {
				return internal.ErrorServerRead("statshouse.getConfig2", err)
			}
			ctx = hctx.WithContext(ctx)
			ret, err := h.GetConfig2(ctx, args)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("statshouse.getConfig2", err)
			}
			if hctx.Response, err = args.WriteResult(hctx.Response, ret); err != nil {
				return internal.ErrorServerWriteResult("statshouse.getConfig2", err)
			}
			return nil
		}
	case 0x42855554: // statshouse.getMetrics3
		hctx.RequestFunctionName = "statshouse.getMetrics3"
		if h.RawGetMetrics3 != nil {
			hctx.Request = r
			err = h.RawGetMetrics3(ctx, hctx)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("statshouse.getMetrics3", err)
			}
			return nil
		}
		if h.GetMetrics3 != nil {
			var args GetMetrics3
			if _, err = args.Read(r); err != nil {
				return internal.ErrorServerRead("statshouse.getMetrics3", err)
			}
			ctx = hctx.WithContext(ctx)
			ret, err := h.GetMetrics3(ctx, args)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("statshouse.getMetrics3", err)
			}
			if hctx.Response, err = args.WriteResult(hctx.Response, ret); err != nil {
				return internal.ErrorServerWriteResult("statshouse.getMetrics3", err)
			}
			return nil
		}
	case 0x4285ff56: // statshouse.getTagMapping2
		hctx.RequestFunctionName = "statshouse.getTagMapping2"
		if h.RawGetTagMapping2 != nil {
			hctx.Request = r
			err = h.RawGetTagMapping2(ctx, hctx)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("statshouse.getTagMapping2", err)
			}
			return nil
		}
		if h.GetTagMapping2 != nil {
			var args GetTagMapping2
			if _, err = args.Read(r); err != nil {
				return internal.ErrorServerRead("statshouse.getTagMapping2", err)
			}
			ctx = hctx.WithContext(ctx)
			ret, err := h.GetTagMapping2(ctx, args)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("statshouse.getTagMapping2", err)
			}
			if hctx.Response, err = args.WriteResult(hctx.Response, ret); err != nil {
				return internal.ErrorServerWriteResult("statshouse.getTagMapping2", err)
			}
			return nil
		}
	case 0x75a7f68e: // statshouse.getTagMappingBootstrap
		hctx.RequestFunctionName = "statshouse.getTagMappingBootstrap"
		if h.RawGetTagMappingBootstrap != nil {
			hctx.Request = r
			err = h.RawGetTagMappingBootstrap(ctx, hctx)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("statshouse.getTagMappingBootstrap", err)
			}
			return nil
		}
		if h.GetTagMappingBootstrap != nil {
			var args GetTagMappingBootstrap
			if _, err = args.Read(r); err != nil {
				return internal.ErrorServerRead("statshouse.getTagMappingBootstrap", err)
			}
			ctx = hctx.WithContext(ctx)
			ret, err := h.GetTagMappingBootstrap(ctx, args)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("statshouse.getTagMappingBootstrap", err)
			}
			if hctx.Response, err = args.WriteResult(hctx.Response, ret); err != nil {
				return internal.ErrorServerWriteResult("statshouse.getTagMappingBootstrap", err)
			}
			return nil
		}
	case 0x41df72a3: // statshouse.getTargets2
		hctx.RequestFunctionName = "statshouse.getTargets2"
		if h.RawGetTargets2 != nil {
			hctx.Request = r
			err = h.RawGetTargets2(ctx, hctx)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("statshouse.getTargets2", err)
			}
			return nil
		}
		if h.GetTargets2 != nil {
			var args GetTargets2
			if _, err = args.Read(r); err != nil {
				return internal.ErrorServerRead("statshouse.getTargets2", err)
			}
			ctx = hctx.WithContext(ctx)
			ret, err := h.GetTargets2(ctx, args)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("statshouse.getTargets2", err)
			}
			if hctx.Response, err = args.WriteResult(hctx.Response, ret); err != nil {
				return internal.ErrorServerWriteResult("statshouse.getTargets2", err)
			}
			return nil
		}
	case 0x4285ff53: // statshouse.sendKeepAlive2
		hctx.RequestFunctionName = "statshouse.sendKeepAlive2"
		if h.RawSendKeepAlive2 != nil {
			hctx.Request = r
			err = h.RawSendKeepAlive2(ctx, hctx)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("statshouse.sendKeepAlive2", err)
			}
			return nil
		}
		if h.SendKeepAlive2 != nil {
			var args SendKeepAlive2
			if _, err = args.Read(r); err != nil {
				return internal.ErrorServerRead("statshouse.sendKeepAlive2", err)
			}
			ctx = hctx.WithContext(ctx)
			ret, err := h.SendKeepAlive2(ctx, args)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("statshouse.sendKeepAlive2", err)
			}
			if hctx.Response, err = args.WriteResult(hctx.Response, ret); err != nil {
				return internal.ErrorServerWriteResult("statshouse.sendKeepAlive2", err)
			}
			return nil
		}
	case 0x44575940: // statshouse.sendSourceBucket2
		hctx.RequestFunctionName = "statshouse.sendSourceBucket2"
		if h.RawSendSourceBucket2 != nil {
			hctx.Request = r
			err = h.RawSendSourceBucket2(ctx, hctx)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("statshouse.sendSourceBucket2", err)
			}
			return nil
		}
		if h.SendSourceBucket2 != nil {
			var args SendSourceBucket2
			if _, err = args.Read(r); err != nil {
				return internal.ErrorServerRead("statshouse.sendSourceBucket2", err)
			}
			ctx = hctx.WithContext(ctx)
			ret, err := h.SendSourceBucket2(ctx, args)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("statshouse.sendSourceBucket2", err)
			}
			if hctx.Response, err = args.WriteResult(hctx.Response, ret); err != nil {
				return internal.ErrorServerWriteResult("statshouse.sendSourceBucket2", err)
			}
			return nil
		}
	case 0x4285ff58: // statshouse.testConnection2
		hctx.RequestFunctionName = "statshouse.testConnection2"
		if h.RawTestConnection2 != nil {
			hctx.Request = r
			err = h.RawTestConnection2(ctx, hctx)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("statshouse.testConnection2", err)
			}
			return nil
		}
		if h.TestConnection2 != nil {
			var args TestConnection2
			if _, err = args.Read(r); err != nil {
				return internal.ErrorServerRead("statshouse.testConnection2", err)
			}
			ctx = hctx.WithContext(ctx)
			ret, err := h.TestConnection2(ctx, args)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("statshouse.testConnection2", err)
			}
			if hctx.Response, err = args.WriteResult(hctx.Response, ret); err != nil {
				return internal.ErrorServerWriteResult("statshouse.testConnection2", err)
			}
			return nil
		}
	}
	return rpc.ErrNoHandler
}
