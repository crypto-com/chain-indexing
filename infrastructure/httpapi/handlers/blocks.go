package handlers

import (
	"errors"
	"strconv"

	"github.com/valyala/fasthttp"

	"github.com/crypto-com/chainindex/appinterface/projection/view"
	"github.com/crypto-com/chainindex/appinterface/rdb"
	"github.com/crypto-com/chainindex/infrastructure/httpapi"
	applogger "github.com/crypto-com/chainindex/internal/logger"
)

type Blocks struct {
	logger applogger.Logger

	blocksView       *view.Blocks
	transactionsView *view.Transactions
}

func NewBlocks(logger applogger.Logger, rdbHandle *rdb.Handle) *Blocks {
	return &Blocks{
		logger.WithFields(applogger.LogFields{
			"module": "BlocksHandler",
		}),

		view.NewBlocks(rdbHandle),
		view.NewTransactions(rdbHandle),
	}
}

func (handler *Blocks) List(ctx *fasthttp.RequestCtx) {
	var err error

	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	blocks, paginationResult, err := handler.blocksView.List(pagination)
	if err != nil {
		handler.logger.Errorf("error listing blocks: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, blocks, paginationResult)
}

func (handler *Blocks) ListTransactionsByHeight(ctx *fasthttp.RequestCtx) {
	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		httpapi.BadRequest(ctx, err)
		return
	}

	blockHeightParam := ctx.UserValue("height")
	blockHeight, err := strconv.ParseInt(blockHeightParam.(string), 10, 64)
	if err != nil {
		httpapi.BadRequest(ctx, errors.New("invalid block height"))
		return
	}

	blocks, paginationResult, err := handler.transactionsView.List(view.TransactionsListFilter{
		MaybeBlockHeight: &blockHeight,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing transactions: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, blocks, paginationResult)
}
