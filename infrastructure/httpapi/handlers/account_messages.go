package handlers

import (
	"strings"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	account_message_view "github.com/crypto-com/chain-indexing/projection/account_message/view"
	"github.com/valyala/fasthttp"
)

type AccountMessages struct {
	logger applogger.Logger

	accountMessagesView *account_message_view.AccountMessages
}

func NewAccountMessages(logger applogger.Logger, rdbHandle *rdb.Handle) *AccountMessages {
	return &AccountMessages{
		logger.WithFields(applogger.LogFields{
			"module": "AccountMessagesHandler",
		}),

		account_message_view.NewAccountMessages(rdbHandle),
	}
}

func (handler *AccountMessages) Register(server *httpapi.Server, routePrefix string) {
	server.GET("/ListByAccount", handler.ListByAccount)
}

func (handler *AccountMessages) ListByAccount(ctx *fasthttp.RequestCtx) {
	var err error

	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	account := ctx.UserValue("account").(string)
	filter := account_message_view.AccountMessagesListFilter{
		Account:       account,
		MaybeMsgTypes: nil,
	}
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("filter.msgType") {
		filter.MaybeMsgTypes = strings.Split(string(queryArgs.Peek("filter.msgType")), ",")
	}

	idOrder := view.ORDER_ASC
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "height.desc" {
			idOrder = view.ORDER_DESC
		}
	}

	blocks, paginationResult, err := handler.accountMessagesView.List(
		filter, account_message_view.AccountMessagesListOrder{Id: idOrder}, pagination,
	)
	if err != nil {
		handler.logger.Errorf("error listing account messages: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, blocks, paginationResult)
}
