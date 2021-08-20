package handlers

import (
	"github.com/valyala/fasthttp"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	channel_view "github.com/crypto-com/chain-indexing/projection/channel/view"
)

type Channels struct {
	logger applogger.Logger

	channelView *channel_view.Channels
}

func NewChannels(logger applogger.Logger, rdbHandle *rdb.Handle) *Channels {
	return &Channels{
		logger.WithFields(applogger.LogFields{
			"module": "ChannelsHandler",
		}),

		channel_view.NewChannels(rdbHandle),
	}
}

func (handler *Channels) List(ctx *fasthttp.RequestCtx) {
	var err error

	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	channelIdOrder := view.ORDER_ASC
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "channel_id.desc" {
			channelIdOrder = view.ORDER_DESC
		}
	}

	channels, paginationResult, err := handler.channelView.List(channel_view.ChannelsListOrder{
		ChannelID: channelIdOrder,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing channels: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, channels, paginationResult)
}
