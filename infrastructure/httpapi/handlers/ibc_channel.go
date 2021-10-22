package handlers

import (
	"errors"
	"fmt"

	"github.com/valyala/fasthttp"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	"github.com/crypto-com/chain-indexing/internal/primptr"
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

func (handler *IBCChannel) Register(server *httpapi.Server, routePrefix string) {
	server.GET(fmt.Sprintf("%s/api/v1/ibc/channels", routePrefix), handler.ListChannels)
	server.GET(fmt.Sprintf("%s/api/v1/ibc/channels/{channelId}", routePrefix), handler.FindChannelById)
	server.GET(fmt.Sprintf("%s/api/v1/ibc/denom-hash-mappings", routePrefix), handler.ListAllDenomHashMapping)
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
		if string(queryArgs.Peek("status")) == "true" {
			listFilter.MaybeStatus = primptr.Bool(true)
		} else if string(queryArgs.Peek("status")) == "false" {
			listFilter.MaybeStatus = primptr.Bool(false)
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
