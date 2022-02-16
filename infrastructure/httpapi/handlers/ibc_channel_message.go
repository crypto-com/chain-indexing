package handlers

import (
	"strings"

	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/valyala/fasthttp"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	ibc_channel_message_view "github.com/crypto-com/chain-indexing/projection/ibc_channel_message/view"
)

type IBCChannelMessage struct {
	logger applogger.Logger

	ibcChannelMessagesView ibc_channel_message_view.IBCChannelMessages
}

func NewIBCChannelMessage(logger applogger.Logger, rdbHandle *rdb.Handle) *IBCChannelMessage {
	return &IBCChannelMessage{
		logger.WithFields(applogger.LogFields{
			"module": "IBCChannelMessageHandler",
		}),

		ibc_channel_message_view.NewIBCChannelMessagesView(rdbHandle),
	}
}

func (handler *IBCChannelMessage) ListByChannelID(ctx *fasthttp.RequestCtx) {
	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	channelID, channelIDOk := URLValueGuard(ctx, handler.logger, "channelId")
	if !channelIDOk {
		return
	}

	queryArgs := ctx.QueryArgs()

	listOrder := ibc_channel_message_view.IBCChannelMessagesListOrder{
		BlockTime: view.ORDER_DESC,
	}
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "blockTime.asc" {
			listOrder.BlockTime = view.ORDER_ASC
		}
	}

	listFilter := ibc_channel_message_view.IBCChannelMessagesListFilter{
		MaybeMsgTypes: nil,
	}
	if queryArgs.Has("filter.msgType") {
		listFilter.MaybeMsgTypes = strings.Split(string(queryArgs.Peek("filter.msgType")), ",")
	}

	ibcChannelMessages, paginationResult, err := handler.ibcChannelMessagesView.ListByChannelID(
		channelID,
		listOrder,
		listFilter,
		pagination,
	)
	if err != nil {
		handler.logger.Errorf("error listing IBCChannelMessage: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, ibcChannelMessages, paginationResult)
}
