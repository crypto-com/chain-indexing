package handlers

import (
	"fmt"

	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/valyala/fasthttp"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	account_transaction_view "github.com/crypto-com/chain-indexing/projection/account_transaction/view"
)

type AccountTransactions struct {
	logger applogger.Logger

	accountTransactionsView *account_transaction_view.AccountTransactions
}

func NewAccountTransactions(logger applogger.Logger, rdbHandle *rdb.Handle) *AccountTransactions {
	return &AccountTransactions{
		logger.WithFields(applogger.LogFields{
			"module": "AccountTransactionsHandler",
		}),

		account_transaction_view.NewAccountTransactions(rdbHandle),
	}
}

func (handler *AccountTransactions) Register(server *httpapi.Server, routePrefix string) {
	server.GET(fmt.Sprintf("%s/api/v1/accounts/{account}/transactions", routePrefix), handler.ListByAccount)
}

func (handler *AccountTransactions) ListByAccount(ctx *fasthttp.RequestCtx) {
	var err error

	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	queryArgs := ctx.QueryArgs()

	idOrder := view.ORDER_ASC
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "height.desc" {
			idOrder = view.ORDER_DESC
		}
	}
	memo := ""
	if queryArgs.Has("memo") {
		memo = string(queryArgs.Peek("memo"))
	}

	account := ctx.UserValue("account").(string)
	filter := account_transaction_view.AccountTransactionsListFilter{
		Account: account,
		Memo:    memo,
	}

	blocks, paginationResult, err := handler.accountTransactionsView.List(
		filter, account_transaction_view.AccountTransactionsListOrder{Id: idOrder}, pagination,
	)
	if err != nil {
		handler.logger.Errorf("error listing account transactions: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, blocks, paginationResult)
}
