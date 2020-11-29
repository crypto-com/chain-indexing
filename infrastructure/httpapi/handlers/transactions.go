package handlers

import (
	"errors"

	"github.com/valyala/fasthttp"

	transaction_view "github.com/crypto-com/chain-indexing/appinterface/projection/transasaction/view"
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
)

type Transactions struct {
	logger applogger.Logger

	transactionsView *transaction_view.BlockTransactions
}

func NewTransactions(logger applogger.Logger, rdbHandle *rdb.Handle) *Transactions {
	return &Transactions{
		logger.WithFields(applogger.LogFields{
			"module": "TransactionsHandler",
		}),

		transaction_view.NewTransactions(rdbHandle),
	}
}

func (handler *Transactions) FindByHash(ctx *fasthttp.RequestCtx) {
	var err error

	transaction, err := handler.transactionsView.FindByHash(
		ctx.UserValue("hash").(string),
	)
	if err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			httpapi.NotFound(ctx)
			return
		}
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

	heightOrder := view.ORDER_ASC
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "height.desc" {
			heightOrder = view.ORDER_DESC
		}
	}

	blocks, paginationResult, err := handler.transactionsView.List(transaction_view.TransactionsListFilter{
		MaybeBlockHeight: nil,
	}, transaction_view.TransactionsListOrder{
		Height: heightOrder,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing transactions: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, blocks, paginationResult)
}
