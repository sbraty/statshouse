// Copyright 2023 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlmetadata

import (
	"context"

	"github.com/vkcom/statshouse/internal/data_model/gen2/internal"
	"github.com/vkcom/statshouse/internal/vkgo/basictl"
	"github.com/vkcom/statshouse/internal/vkgo/rpc"
)

type (
	CreateEntityEvent                    = internal.MetadataCreateEntityEvent
	CreateMappingEvent                   = internal.MetadataCreateMappingEvent
	CreateMetricEvent                    = internal.MetadataCreateMetricEvent
	EditEntityEvent                      = internal.MetadataEditEntityEvent
	EditEntitynew                        = internal.MetadataEditEntitynew
	EditMetricEvent                      = internal.MetadataEditMetricEvent
	Event                                = internal.MetadataEvent
	EventBytes                           = internal.MetadataEventBytes
	GetEntity                            = internal.MetadataGetEntity
	GetHistoryShortInfo                  = internal.MetadataGetHistoryShortInfo
	GetInvertMapping                     = internal.MetadataGetInvertMapping
	GetInvertMappingResponse             = internal.MetadataGetInvertMappingResponse
	GetInvertMappingResponse0            = internal.MetadataGetInvertMappingResponse0
	GetInvertMappingResponseKeyNotExists = internal.MetadataGetInvertMappingResponseKeyNotExists
	GetJournalResponsenew                = internal.MetadataGetJournalResponsenew
	GetJournalResponsenewBytes           = internal.MetadataGetJournalResponsenewBytes
	GetJournalnew                        = internal.MetadataGetJournalnew
	GetMapping                           = internal.MetadataGetMapping
	GetMappingResponse                   = internal.MetadataGetMappingResponse
	GetMappingResponse0                  = internal.MetadataGetMappingResponse0
	GetMappingResponseCreated            = internal.MetadataGetMappingResponseCreated
	GetMappingResponseFloodLimitError    = internal.MetadataGetMappingResponseFloodLimitError
	GetMappingResponseKeyNotExists       = internal.MetadataGetMappingResponseKeyNotExists
	GetMetrics                           = internal.MetadataGetMetrics
	GetMetricsResponse                   = internal.MetadataGetMetricsResponse
	GetTagMappingBootstrap               = internal.MetadataGetTagMappingBootstrap
	HistoryShortResponse                 = internal.MetadataHistoryShortResponse
	HistoryShortResponseEvent            = internal.MetadataHistoryShortResponseEvent
	MetricOld                            = internal.MetadataMetricOld
	PutBootstrapEvent                    = internal.MetadataPutBootstrapEvent
	PutMapping                           = internal.MetadataPutMapping
	PutMappingEvent                      = internal.MetadataPutMappingEvent
	PutMappingResponse                   = internal.MetadataPutMappingResponse
	PutTagMappingBootstrap               = internal.MetadataPutTagMappingBootstrap
	ResetFlood                           = internal.MetadataResetFlood
	ResetFlood2                          = internal.MetadataResetFlood2
	ResetFloodResponse                   = internal.MetadataResetFloodResponse
	ResetFloodResponse2                  = internal.MetadataResetFloodResponse2
)

type Client struct {
	Client  *rpc.Client
	Network string // should be either "tcp4" or "unix"
	Address string
	ActorID int64 // should be non-zero when using rpc-proxy
}

func (c *Client) EditEntitynew(ctx context.Context, args EditEntitynew, extra *rpc.InvokeReqExtra, ret *Event) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "metadata.editEntitynew"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxedGeneral(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("metadata.editEntitynew", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("metadata.editEntitynew", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("metadata.editEntitynew", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) GetEntity(ctx context.Context, args GetEntity, extra *rpc.InvokeReqExtra, ret *Event) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.ReadOnly = true
	req.FunctionName = "metadata.getEntity"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxedGeneral(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("metadata.getEntity", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("metadata.getEntity", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("metadata.getEntity", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) GetHistoryShortInfo(ctx context.Context, args GetHistoryShortInfo, extra *rpc.InvokeReqExtra, ret *HistoryShortResponse) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "metadata.getHistoryShortInfo"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxedGeneral(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("metadata.getHistoryShortInfo", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("metadata.getHistoryShortInfo", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("metadata.getHistoryShortInfo", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) GetInvertMapping(ctx context.Context, args GetInvertMapping, extra *rpc.InvokeReqExtra, ret *GetInvertMappingResponse) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.ReadOnly = true
	req.FunctionName = "metadata.getInvertMapping"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxedGeneral(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("metadata.getInvertMapping", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("metadata.getInvertMapping", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("metadata.getInvertMapping", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) GetJournalnew(ctx context.Context, args GetJournalnew, extra *rpc.InvokeReqExtra, ret *GetJournalResponsenew) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.ReadOnly = true
	req.FunctionName = "metadata.getJournalnew"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxedGeneral(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("metadata.getJournalnew", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("metadata.getJournalnew", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("metadata.getJournalnew", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) GetMapping(ctx context.Context, args GetMapping, extra *rpc.InvokeReqExtra, ret *GetMappingResponse) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "metadata.getMapping"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxedGeneral(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("metadata.getMapping", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("metadata.getMapping", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("metadata.getMapping", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) GetMetrics(ctx context.Context, args GetMetrics, extra *rpc.InvokeReqExtra, ret *GetMetricsResponse) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.ReadOnly = true
	req.FunctionName = "metadata.getMetrics"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxedGeneral(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("metadata.getMetrics", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("metadata.getMetrics", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("metadata.getMetrics", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) GetTagMappingBootstrap(ctx context.Context, args GetTagMappingBootstrap, extra *rpc.InvokeReqExtra, ret *internal.StatshouseGetTagMappingBootstrapResult) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "metadata.getTagMappingBootstrap"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxedGeneral(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("metadata.getTagMappingBootstrap", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("metadata.getTagMappingBootstrap", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("metadata.getTagMappingBootstrap", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) PutMapping(ctx context.Context, args PutMapping, extra *rpc.InvokeReqExtra, ret *PutMappingResponse) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "metadata.putMapping"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxedGeneral(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("metadata.putMapping", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("metadata.putMapping", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("metadata.putMapping", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) PutTagMappingBootstrap(ctx context.Context, args PutTagMappingBootstrap, extra *rpc.InvokeReqExtra, ret *internal.StatshousePutTagMappingBootstrapResult) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "metadata.putTagMappingBootstrap"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxedGeneral(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("metadata.putTagMappingBootstrap", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("metadata.putTagMappingBootstrap", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("metadata.putTagMappingBootstrap", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) ResetFlood(ctx context.Context, args ResetFlood, extra *rpc.InvokeReqExtra, ret *ResetFloodResponse) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "metadata.resetFlood"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxedGeneral(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("metadata.resetFlood", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("metadata.resetFlood", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("metadata.resetFlood", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

func (c *Client) ResetFlood2(ctx context.Context, args ResetFlood2, extra *rpc.InvokeReqExtra, ret *ResetFloodResponse2) (err error) {
	req := c.Client.GetRequest()
	req.ActorID = c.ActorID
	req.FunctionName = "metadata.resetFlood2"
	if extra != nil {
		req.Extra = *extra
	}
	req.Body, err = args.WriteBoxedGeneral(req.Body)
	if err != nil {
		return internal.ErrorClientWrite("metadata.resetFlood2", err)
	}
	resp, err := c.Client.Do(ctx, c.Network, c.Address, req)
	defer c.Client.PutResponse(resp)
	if err != nil {
		return internal.ErrorClientDo("metadata.resetFlood2", c.Network, c.ActorID, c.Address, err)
	}
	if ret != nil {
		if _, err = args.ReadResult(resp.Body, ret); err != nil {
			return internal.ErrorClientReadResult("metadata.resetFlood2", c.Network, c.ActorID, c.Address, err)
		}
	}
	return nil
}

type Handler struct {
	EditEntitynew          func(ctx context.Context, args EditEntitynew) (Event, error)                                                    // metadata.editEntitynew
	GetEntity              func(ctx context.Context, args GetEntity) (Event, error)                                                        // metadata.getEntity
	GetHistoryShortInfo    func(ctx context.Context, args GetHistoryShortInfo) (HistoryShortResponse, error)                               // metadata.getHistoryShortInfo
	GetInvertMapping       func(ctx context.Context, args GetInvertMapping) (GetInvertMappingResponse, error)                              // metadata.getInvertMapping
	GetJournalnew          func(ctx context.Context, args GetJournalnew) (GetJournalResponsenew, error)                                    // metadata.getJournalnew
	GetMapping             func(ctx context.Context, args GetMapping) (GetMappingResponse, error)                                          // metadata.getMapping
	GetMetrics             func(ctx context.Context, args GetMetrics) (GetMetricsResponse, error)                                          // metadata.getMetrics
	GetTagMappingBootstrap func(ctx context.Context, args GetTagMappingBootstrap) (internal.StatshouseGetTagMappingBootstrapResult, error) // metadata.getTagMappingBootstrap
	PutMapping             func(ctx context.Context, args PutMapping) (PutMappingResponse, error)                                          // metadata.putMapping
	PutTagMappingBootstrap func(ctx context.Context, args PutTagMappingBootstrap) (internal.StatshousePutTagMappingBootstrapResult, error) // metadata.putTagMappingBootstrap
	ResetFlood             func(ctx context.Context, args ResetFlood) (ResetFloodResponse, error)                                          // metadata.resetFlood
	ResetFlood2            func(ctx context.Context, args ResetFlood2) (ResetFloodResponse2, error)                                        // metadata.resetFlood2

	RawEditEntitynew          func(ctx context.Context, hctx *rpc.HandlerContext) error // metadata.editEntitynew
	RawGetEntity              func(ctx context.Context, hctx *rpc.HandlerContext) error // metadata.getEntity
	RawGetHistoryShortInfo    func(ctx context.Context, hctx *rpc.HandlerContext) error // metadata.getHistoryShortInfo
	RawGetInvertMapping       func(ctx context.Context, hctx *rpc.HandlerContext) error // metadata.getInvertMapping
	RawGetJournalnew          func(ctx context.Context, hctx *rpc.HandlerContext) error // metadata.getJournalnew
	RawGetMapping             func(ctx context.Context, hctx *rpc.HandlerContext) error // metadata.getMapping
	RawGetMetrics             func(ctx context.Context, hctx *rpc.HandlerContext) error // metadata.getMetrics
	RawGetTagMappingBootstrap func(ctx context.Context, hctx *rpc.HandlerContext) error // metadata.getTagMappingBootstrap
	RawPutMapping             func(ctx context.Context, hctx *rpc.HandlerContext) error // metadata.putMapping
	RawPutTagMappingBootstrap func(ctx context.Context, hctx *rpc.HandlerContext) error // metadata.putTagMappingBootstrap
	RawResetFlood             func(ctx context.Context, hctx *rpc.HandlerContext) error // metadata.resetFlood
	RawResetFlood2            func(ctx context.Context, hctx *rpc.HandlerContext) error // metadata.resetFlood2
}

func (h *Handler) Handle(ctx context.Context, hctx *rpc.HandlerContext) (err error) {
	tag, r, _ := basictl.NatReadTag(hctx.Request) // keep hctx.Request intact for handler chaining
	switch tag {
	case 0x86df475f: // metadata.editEntitynew
		hctx.RequestFunctionName = "metadata.editEntitynew"
		if h.RawEditEntitynew != nil {
			hctx.Request = r
			err = h.RawEditEntitynew(ctx, hctx)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("metadata.editEntitynew", err)
			}
			return nil
		}
		if h.EditEntitynew != nil {
			var args EditEntitynew
			if _, err = args.Read(r); err != nil {
				return internal.ErrorServerRead("metadata.editEntitynew", err)
			}
			ctx = hctx.WithContext(ctx)
			ret, err := h.EditEntitynew(ctx, args)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("metadata.editEntitynew", err)
			}
			if hctx.Response, err = args.WriteResult(hctx.Response, ret); err != nil {
				return internal.ErrorServerWriteResult("metadata.editEntitynew", err)
			}
			return nil
		}
	case 0x72b132f8: // metadata.getEntity
		hctx.RequestFunctionName = "metadata.getEntity"
		if h.RawGetEntity != nil {
			hctx.Request = r
			err = h.RawGetEntity(ctx, hctx)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("metadata.getEntity", err)
			}
			return nil
		}
		if h.GetEntity != nil {
			var args GetEntity
			if _, err = args.Read(r); err != nil {
				return internal.ErrorServerRead("metadata.getEntity", err)
			}
			ctx = hctx.WithContext(ctx)
			ret, err := h.GetEntity(ctx, args)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("metadata.getEntity", err)
			}
			if hctx.Response, err = args.WriteResult(hctx.Response, ret); err != nil {
				return internal.ErrorServerWriteResult("metadata.getEntity", err)
			}
			return nil
		}
	case 0x22ff6a79: // metadata.getHistoryShortInfo
		hctx.RequestFunctionName = "metadata.getHistoryShortInfo"
		if h.RawGetHistoryShortInfo != nil {
			hctx.Request = r
			err = h.RawGetHistoryShortInfo(ctx, hctx)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("metadata.getHistoryShortInfo", err)
			}
			return nil
		}
		if h.GetHistoryShortInfo != nil {
			var args GetHistoryShortInfo
			if _, err = args.Read(r); err != nil {
				return internal.ErrorServerRead("metadata.getHistoryShortInfo", err)
			}
			ctx = hctx.WithContext(ctx)
			ret, err := h.GetHistoryShortInfo(ctx, args)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("metadata.getHistoryShortInfo", err)
			}
			if hctx.Response, err = args.WriteResult(hctx.Response, ret); err != nil {
				return internal.ErrorServerWriteResult("metadata.getHistoryShortInfo", err)
			}
			return nil
		}
	case 0x9faf5280: // metadata.getInvertMapping
		hctx.RequestFunctionName = "metadata.getInvertMapping"
		if h.RawGetInvertMapping != nil {
			hctx.Request = r
			err = h.RawGetInvertMapping(ctx, hctx)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("metadata.getInvertMapping", err)
			}
			return nil
		}
		if h.GetInvertMapping != nil {
			var args GetInvertMapping
			if _, err = args.Read(r); err != nil {
				return internal.ErrorServerRead("metadata.getInvertMapping", err)
			}
			ctx = hctx.WithContext(ctx)
			ret, err := h.GetInvertMapping(ctx, args)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("metadata.getInvertMapping", err)
			}
			if hctx.Response, err = args.WriteResult(hctx.Response, ret); err != nil {
				return internal.ErrorServerWriteResult("metadata.getInvertMapping", err)
			}
			return nil
		}
	case 0x93ba92f8: // metadata.getJournalnew
		hctx.RequestFunctionName = "metadata.getJournalnew"
		if h.RawGetJournalnew != nil {
			hctx.Request = r
			err = h.RawGetJournalnew(ctx, hctx)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("metadata.getJournalnew", err)
			}
			return nil
		}
		if h.GetJournalnew != nil {
			var args GetJournalnew
			if _, err = args.Read(r); err != nil {
				return internal.ErrorServerRead("metadata.getJournalnew", err)
			}
			ctx = hctx.WithContext(ctx)
			ret, err := h.GetJournalnew(ctx, args)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("metadata.getJournalnew", err)
			}
			if hctx.Response, err = args.WriteResult(hctx.Response, ret); err != nil {
				return internal.ErrorServerWriteResult("metadata.getJournalnew", err)
			}
			return nil
		}
	case 0x9dfa7a83: // metadata.getMapping
		hctx.RequestFunctionName = "metadata.getMapping"
		if h.RawGetMapping != nil {
			hctx.Request = r
			err = h.RawGetMapping(ctx, hctx)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("metadata.getMapping", err)
			}
			return nil
		}
		if h.GetMapping != nil {
			var args GetMapping
			if _, err = args.Read(r); err != nil {
				return internal.ErrorServerRead("metadata.getMapping", err)
			}
			ctx = hctx.WithContext(ctx)
			ret, err := h.GetMapping(ctx, args)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("metadata.getMapping", err)
			}
			if hctx.Response, err = args.WriteResult(hctx.Response, ret); err != nil {
				return internal.ErrorServerWriteResult("metadata.getMapping", err)
			}
			return nil
		}
	case 0x93ba92f5: // metadata.getMetrics
		hctx.RequestFunctionName = "metadata.getMetrics"
		if h.RawGetMetrics != nil {
			hctx.Request = r
			err = h.RawGetMetrics(ctx, hctx)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("metadata.getMetrics", err)
			}
			return nil
		}
		if h.GetMetrics != nil {
			var args GetMetrics
			if _, err = args.Read(r); err != nil {
				return internal.ErrorServerRead("metadata.getMetrics", err)
			}
			ctx = hctx.WithContext(ctx)
			ret, err := h.GetMetrics(ctx, args)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("metadata.getMetrics", err)
			}
			if hctx.Response, err = args.WriteResult(hctx.Response, ret); err != nil {
				return internal.ErrorServerWriteResult("metadata.getMetrics", err)
			}
			return nil
		}
	case 0x5fc81a9b: // metadata.getTagMappingBootstrap
		hctx.RequestFunctionName = "metadata.getTagMappingBootstrap"
		if h.RawGetTagMappingBootstrap != nil {
			hctx.Request = r
			err = h.RawGetTagMappingBootstrap(ctx, hctx)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("metadata.getTagMappingBootstrap", err)
			}
			return nil
		}
		if h.GetTagMappingBootstrap != nil {
			var args GetTagMappingBootstrap
			if _, err = args.Read(r); err != nil {
				return internal.ErrorServerRead("metadata.getTagMappingBootstrap", err)
			}
			ctx = hctx.WithContext(ctx)
			ret, err := h.GetTagMappingBootstrap(ctx, args)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("metadata.getTagMappingBootstrap", err)
			}
			if hctx.Response, err = args.WriteResult(hctx.Response, ret); err != nil {
				return internal.ErrorServerWriteResult("metadata.getTagMappingBootstrap", err)
			}
			return nil
		}
	case 0x9faf5281: // metadata.putMapping
		hctx.RequestFunctionName = "metadata.putMapping"
		if h.RawPutMapping != nil {
			hctx.Request = r
			err = h.RawPutMapping(ctx, hctx)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("metadata.putMapping", err)
			}
			return nil
		}
		if h.PutMapping != nil {
			var args PutMapping
			if _, err = args.Read(r); err != nil {
				return internal.ErrorServerRead("metadata.putMapping", err)
			}
			ctx = hctx.WithContext(ctx)
			ret, err := h.PutMapping(ctx, args)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("metadata.putMapping", err)
			}
			if hctx.Response, err = args.WriteResult(hctx.Response, ret); err != nil {
				return internal.ErrorServerWriteResult("metadata.putMapping", err)
			}
			return nil
		}
	case 0x5fc8ab9b: // metadata.putTagMappingBootstrap
		hctx.RequestFunctionName = "metadata.putTagMappingBootstrap"
		if h.RawPutTagMappingBootstrap != nil {
			hctx.Request = r
			err = h.RawPutTagMappingBootstrap(ctx, hctx)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("metadata.putTagMappingBootstrap", err)
			}
			return nil
		}
		if h.PutTagMappingBootstrap != nil {
			var args PutTagMappingBootstrap
			if _, err = args.Read(r); err != nil {
				return internal.ErrorServerRead("metadata.putTagMappingBootstrap", err)
			}
			ctx = hctx.WithContext(ctx)
			ret, err := h.PutTagMappingBootstrap(ctx, args)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("metadata.putTagMappingBootstrap", err)
			}
			if hctx.Response, err = args.WriteResult(hctx.Response, ret); err != nil {
				return internal.ErrorServerWriteResult("metadata.putTagMappingBootstrap", err)
			}
			return nil
		}
	case 0x9faf5282: // metadata.resetFlood
		hctx.RequestFunctionName = "metadata.resetFlood"
		if h.RawResetFlood != nil {
			hctx.Request = r
			err = h.RawResetFlood(ctx, hctx)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("metadata.resetFlood", err)
			}
			return nil
		}
		if h.ResetFlood != nil {
			var args ResetFlood
			if _, err = args.Read(r); err != nil {
				return internal.ErrorServerRead("metadata.resetFlood", err)
			}
			ctx = hctx.WithContext(ctx)
			ret, err := h.ResetFlood(ctx, args)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("metadata.resetFlood", err)
			}
			if hctx.Response, err = args.WriteResult(hctx.Response, ret); err != nil {
				return internal.ErrorServerWriteResult("metadata.resetFlood", err)
			}
			return nil
		}
	case 0x88d0fd5e: // metadata.resetFlood2
		hctx.RequestFunctionName = "metadata.resetFlood2"
		if h.RawResetFlood2 != nil {
			hctx.Request = r
			err = h.RawResetFlood2(ctx, hctx)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("metadata.resetFlood2", err)
			}
			return nil
		}
		if h.ResetFlood2 != nil {
			var args ResetFlood2
			if _, err = args.Read(r); err != nil {
				return internal.ErrorServerRead("metadata.resetFlood2", err)
			}
			ctx = hctx.WithContext(ctx)
			ret, err := h.ResetFlood2(ctx, args)
			if rpc.IsHijackedResponse(err) {
				return err
			}
			if err != nil {
				return internal.ErrorServerHandle("metadata.resetFlood2", err)
			}
			if hctx.Response, err = args.WriteResult(hctx.Response, ret); err != nil {
				return internal.ErrorServerWriteResult("metadata.resetFlood2", err)
			}
			return nil
		}
	}
	return rpc.ErrNoHandler
}
