package handlers

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	account_raw_event_view "github.com/crypto-com/chain-indexing/projection/account_raw_event/view"
	"github.com/valyala/fasthttp"
)

type AccountRawEvents struct {
	logger applogger.Logger

	accountRawEventsView account_raw_event_view.AccountRawEvents
}

func NewAccountRawEvents(logger applogger.Logger, rdbHandle *rdb.Handle) *AccountRawEvents {
	return &AccountRawEvents{
		logger.WithFields(applogger.LogFields{
			"module": "AccountMessagesHandler",
		}),
		account_raw_event_view.NewAccountRawEventsView(rdbHandle),
	}
}

func (handler *AccountRawEvents) ListByAccount(ctx *fasthttp.RequestCtx) {
	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	account, accountOk := URLValueGuard(ctx, handler.logger, "account")
	if !accountOk {
		return
	}
	filter := account_raw_event_view.AccountRawEventsListFilter{
		Account:       account,
		MaybeMsgTypes: nil,
	}
	queryArgs := ctx.QueryArgs()
	idOrder := view.ORDER_ASC
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "height.desc" {
			idOrder = view.ORDER_DESC
		}
	}

	blocks, paginationResult, err := handler.accountRawEventsView.List(
		filter, account_raw_event_view.AccountRawEventsListOrder{Id: idOrder}, pagination,
	)
	if err != nil {
		handler.logger.Errorf("error listing account raw events: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, blocks, paginationResult)
}
