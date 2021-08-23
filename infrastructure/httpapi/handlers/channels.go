package handlers

import (
	"errors"

	"github.com/valyala/fasthttp"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	channel_view "github.com/crypto-com/chain-indexing/projection/channel/view"
)

type Channels struct {
	logger applogger.Logger

	ibcChannelView *channel_view.IBCChannels
}

func NewChannels(logger applogger.Logger, rdbHandle *rdb.Handle) *Channels {
	return &Channels{
		logger.WithFields(applogger.LogFields{
			"module": "ChannelsHandler",
		}),

		channel_view.NewIBCChannels(rdbHandle),
	}
}

func (handler *Channels) List(ctx *fasthttp.RequestCtx) {
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

	channels, paginationResult, err := handler.ibcChannelView.List(channel_view.ChannelsListOrder{
		ChannelID: channelIdOrder,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing channels: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, channels, paginationResult)
}

func (handler *Channels) FindById(ctx *fasthttp.RequestCtx) {
	channel, err := handler.ibcChannelView.FindBy(
		ctx.UserValue("channelId").(string),
	)
	if err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			httpapi.NotFound(ctx)
			return
		}
		handler.logger.Errorf("error finding channel by id: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.Success(ctx, channel)
}
