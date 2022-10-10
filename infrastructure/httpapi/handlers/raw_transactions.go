package handlers

import (
	"errors"

	applogger "github.com/crypto-com/chain-indexing/external/logger"
	raw_transaction_view "github.com/crypto-com/chain-indexing/projection/raw_transaction/view"
	"github.com/valyala/fasthttp"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
)

type RawTransactions struct {
	logger applogger.Logger

	rawTransactionsView raw_transaction_view.BlockRawTransactions
}

func NewRawTransactions(logger applogger.Logger, rdbHandle *rdb.Handle) *RawTransactions {
	return &RawTransactions{
		logger.WithFields(applogger.LogFields{
			"module": "RawTransactionsHandler",
		}),

		raw_transaction_view.NewRawTransactionsView(rdbHandle),
	}
}

func (handler *RawTransactions) FindByHash(ctx *fasthttp.RequestCtx) {
	hashParam, hashParamOk := URLValueGuard(ctx, handler.logger, "hash")
	if !hashParamOk {
		return
	}

	rawTransaction, err := handler.rawTransactionsView.FindByHash(hashParam)
	if err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			httpapi.NotFound(ctx)
			return
		}
		handler.logger.Errorf("error finding transactions by hash: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.Success(ctx, rawTransaction)
}
