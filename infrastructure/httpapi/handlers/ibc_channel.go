package handlers

import (
	"errors"

	"github.com/valyala/fasthttp"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	ibc_channel_view "github.com/crypto-com/chain-indexing/projection/ibc_channel/view"
)

type IBCChannel struct {
	logger applogger.Logger

	ibcChannelsView         *ibc_channel_view.IBCChannels
	ibcDenomHashMappingView *ibc_channel_view.IBCDenomHashMapping
}

func NewIBCChannel(logger applogger.Logger, rdbHandle *rdb.Handle) *IBCChannel {
	return &IBCChannel{
		logger.WithFields(applogger.LogFields{
			"module": "IBCChannelHandler",
		}),

		ibc_channel_view.NewIBCChannels(rdbHandle),
		ibc_channel_view.NewIBCDenomHashMapping(rdbHandle),
	}
}

func (handler *IBCChannel) ListChannels(ctx *fasthttp.RequestCtx) {
	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	var listOrder ibc_channel_view.IBCChannelsListOrder
	view_order_desc := view.ORDER_DESC
	view_order_asc := view.ORDER_ASC

	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "createdAtBlockTime.desc" {
			listOrder.MaybeCreatedAtBlockTime = &view_order_desc
		} else if string(queryArgs.Peek("order")) == "createdAtBlockTime.asc" {
			listOrder.MaybeCreatedAtBlockTime = &view_order_asc
		} else if string(queryArgs.Peek("order")) == "lastActivityBlockTime.desc" {
			listOrder.MaybeLastActivityBlockTime = &view_order_desc
		} else if string(queryArgs.Peek("order")) == "lastActivityBlockTime.asc" {
			listOrder.MaybeLastActivityBlockTime = &view_order_asc
		}
	}

	var listFilter ibc_channel_view.IBCChannelsListFilter
	view_filter_status_true := true
	view_filter_status_false := false

	if queryArgs.Has("status") {
		if string(queryArgs.Peek("status")) == "true" {
			listFilter.MaybeStatus = &view_filter_status_true
		} else if string(queryArgs.Peek("status")) == "false" {
			listFilter.MaybeStatus = &view_filter_status_false
		}
	}

	ibcChannels, paginationResult, err := handler.ibcChannelsView.List(listOrder, listFilter, pagination)
	if err != nil {
		handler.logger.Errorf("error listing IBCChannel channels: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, ibcChannels, paginationResult)
}

func (handler *IBCChannel) FindChannelById(ctx *fasthttp.RequestCtx) {
	ibcChannel, err := handler.ibcChannelsView.FindBy(
		ctx.UserValue("channelId").(string),
	)
	if err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			httpapi.NotFound(ctx)
			return
		}
		handler.logger.Errorf("error finding IBCChannel channel by id: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.Success(ctx, ibcChannel)
}

func (handler *IBCChannel) ListAllDenomHashMapping(ctx *fasthttp.RequestCtx) {
	ibcDenomHashMappings, err := handler.ibcDenomHashMappingView.ListAll()
	if err != nil {
		handler.logger.Errorf("error listing IBCDenomHashMppings: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.Success(ctx, ibcDenomHashMappings)
}
