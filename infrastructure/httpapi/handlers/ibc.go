package handlers

import (
	"errors"

	"github.com/valyala/fasthttp"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	ibc_view "github.com/crypto-com/chain-indexing/projection/ibc/view"
)

type IBC struct {
	logger applogger.Logger

	ibcChannelsView *ibc_view.IBCChannels
}

func NewIBC(logger applogger.Logger, rdbHandle *rdb.Handle) *IBC {
	return &IBC{
		logger.WithFields(applogger.LogFields{
			"module": "IBCHandler",
		}),

		ibc_view.NewIBCChannels(rdbHandle),
	}
}

func (handler *IBC) ListChannels(ctx *fasthttp.RequestCtx) {
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

	ibcChannels, paginationResult, err := handler.ibcChannelsView.List(ibc_view.IBCChannelsListOrder{
		ChannelID: channelIdOrder,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing IBC channels: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, ibcChannels, paginationResult)
}

func (handler *IBC) FindChannelById(ctx *fasthttp.RequestCtx) {
	ibcChannel, err := handler.ibcChannelsView.FindBy(
		ctx.UserValue("channelId").(string),
	)
	if err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			httpapi.NotFound(ctx)
			return
		}
		handler.logger.Errorf("error finding IBC channel by id: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.Success(ctx, ibcChannel)
}
