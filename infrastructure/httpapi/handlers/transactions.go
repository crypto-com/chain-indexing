package handlers

import (
	"github.com/valyala/fasthttp"

	"github.com/crypto-com/chainindex/appinterface/projection/view"
	"github.com/crypto-com/chainindex/appinterface/rdb"
	"github.com/crypto-com/chainindex/infrastructure/httpapi"
	applogger "github.com/crypto-com/chainindex/internal/logger"
)

type Transactions struct {
	logger applogger.Logger

	transactionsView *view.BlockTransactions
}

func NewTransactions(logger applogger.Logger, rdbHandle *rdb.Handle) *Transactions {
	return &Transactions{
		logger.WithFields(applogger.LogFields{
			"module": "TransactionsHandler",
		}),

		view.NewTransactions(rdbHandle),
	}
}

func (handler *Transactions) FindByHash(ctx *fasthttp.RequestCtx) {
	var err error

	transaction, err := handler.transactionsView.FindByHash(
		ctx.UserValue("hash").(string),
	)
	if err != nil {
		handler.logger.Errorf("error finding transactions by hash: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.Success(ctx, transaction)
}

func (handler *Transactions) List(ctx *fasthttp.RequestCtx) {
	var err error

	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	blocks, paginationResult, err := handler.transactionsView.List(view.TransactionsListFilter{
		MaybeBlockHeight: nil,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing transactions: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, blocks, paginationResult)
}
