package handlers

import (
	"errors"

	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/valyala/fasthttp"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	ibc_channel_types "github.com/crypto-com/chain-indexing/projection/ibc_channel/types"
	ibc_channel_view "github.com/crypto-com/chain-indexing/projection/ibc_channel/view"
)

type IBCChannel struct {
	logger applogger.Logger

	ibcChannelsView         ibc_channel_view.IBCChannels
	ibcDenomHashMappingView ibc_channel_view.IBCDenomHashMapping
}

func NewIBCChannel(logger applogger.Logger, rdbHandle *rdb.Handle) *IBCChannel {
	return &IBCChannel{
		logger.WithFields(applogger.LogFields{
			"module": "IBCChannelHandler",
		}),

		ibc_channel_view.NewIBCChannelsView(rdbHandle),
		ibc_channel_view.NewIBCDenomHashMappingView(rdbHandle),
	}
}

func (handler *IBCChannel) ListChannels(ctx *fasthttp.RequestCtx) {
	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	queryArgs := ctx.QueryArgs()

	var listOrder ibc_channel_view.IBCChannelsListOrder
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "createdAtBlockTime.desc" {
			listOrder.MaybeCreatedAtBlockTime = primptr.String(view.ORDER_DESC)
		} else if string(queryArgs.Peek("order")) == "createdAtBlockTime.asc" {
			listOrder.MaybeCreatedAtBlockTime = primptr.String(view.ORDER_ASC)
		} else if string(queryArgs.Peek("order")) == "lastActivityBlockTime.desc" {
			listOrder.MaybeLastActivityBlockTime = primptr.String(view.ORDER_DESC)
		} else if string(queryArgs.Peek("order")) == "lastActivityBlockTime.asc" {
			listOrder.MaybeLastActivityBlockTime = primptr.String(view.ORDER_ASC)
		}
	}

	var listFilter ibc_channel_view.IBCChannelsListFilter
	if queryArgs.Has("status") {
		if string(queryArgs.Peek("status")) == ibc_channel_types.STATUS_NOT_ESTABLISHED {
			listFilter.MaybeStatus = primptr.String(ibc_channel_types.STATUS_NOT_ESTABLISHED)
		} else if string(queryArgs.Peek("status")) == ibc_channel_types.STATUS_OPENED {
			listFilter.MaybeStatus = primptr.String(ibc_channel_types.STATUS_OPENED)
		} else if string(queryArgs.Peek("status")) == ibc_channel_types.STATUS_CLOSED {
			listFilter.MaybeStatus = primptr.String(ibc_channel_types.STATUS_CLOSED)
		}
	}

	if queryArgs.Has("groupBy") {
		if string(queryArgs.Peek("groupBy")) == "chainId" {
			ibcChannelsGroupByChainId, paginationResult, listChannelsErr := handler.ibcChannelsView.ListChannelsGroupByChainId(listOrder, listFilter, pagination)
			if listChannelsErr != nil {
				handler.logger.Errorf("error listing IBC Channels grouped by chainId: %v", err)
				httpapi.InternalServerError(ctx)
				return
			}

			httpapi.SuccessWithPagination(ctx, ibcChannelsGroupByChainId, paginationResult)
			return
		}
	}

	ibcChannels, paginationResult, err := handler.ibcChannelsView.List(listOrder, listFilter, pagination)
	if err != nil {
		handler.logger.Errorf("error listing IBC channels: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, ibcChannels, paginationResult)
}
func (handler *IBCChannel) FindChannelById(ctx *fasthttp.RequestCtx) {
	channelId, channelIdOk := URLValueGuard(ctx, handler.logger, "channelId")
	if !channelIdOk {
		return
	}
	ibcChannel, err := handler.ibcChannelsView.FindBy(channelId)
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
