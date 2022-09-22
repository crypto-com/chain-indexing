package handlers

import (
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/valyala/fasthttp"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	account_raw_transaction_view "github.com/crypto-com/chain-indexing/projection/account_raw_transaction/view"
)

type AccountRawTransactions struct {
	logger applogger.Logger

	accountRawTransactionsView account_raw_transaction_view.AccountRawTransactions
}

func NewAccountRawTransactions(logger applogger.Logger, rdbHandle *rdb.Handle) *AccountRawTransactions {
	return &AccountRawTransactions{
		logger.WithFields(applogger.LogFields{
			"module": "AccountRawTransactionsHandler",
		}),

		account_raw_transaction_view.NewAccountRawTransactionsView(rdbHandle),
	}
}

func (handler *AccountRawTransactions) ListByAccount(ctx *fasthttp.RequestCtx) {
	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	account, accountOk := URLValueGuard(ctx, handler.logger, "account")
	if !accountOk {
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

	filter := account_raw_transaction_view.AccountRawTransactionsListFilter{
		Account: account,
		Memo:    memo,
	}

	blocks, paginationResult, err := handler.accountRawTransactionsView.List(
		filter, account_raw_transaction_view.AccountRawTransactionsListOrder{Id: idOrder}, pagination,
	)
	if err != nil {
		handler.logger.Errorf("error listing account raw transactions: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, blocks, paginationResult)
}
