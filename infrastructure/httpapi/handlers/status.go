package handlers

import (
	view2 "github.com/crypto-com/chainindex/appinterface/projection/block/view"
	view3 "github.com/crypto-com/chainindex/appinterface/projection/transasaction/view"
	"github.com/crypto-com/chainindex/appinterface/rdb"
	"github.com/crypto-com/chainindex/infrastructure/httpapi"
	applogger "github.com/crypto-com/chainindex/internal/logger"
	"github.com/valyala/fasthttp"
)

type StatusHandler struct {
	logger applogger.Logger

	blocksView       *view2.Blocks
	transactionsView *view3.BlockTransactions
}

func NewStatusHandler(logger applogger.Logger, rdbHandle *rdb.Handle) *StatusHandler {
	return &StatusHandler{
		logger.WithFields(applogger.LogFields{
			"module": "StatusHandler",
		}),

		view2.NewBlocks(rdbHandle),
		view3.NewTransactions(rdbHandle),
	}
}

func (handler *StatusHandler) GetStatus(ctx *fasthttp.RequestCtx) {
	blockCount, err := handler.blocksView.Count()

	if err != nil {
		handler.logger.Errorf("error fetching block count: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	transactionCount, err := handler.transactionsView.Count()

	if err != nil {
		handler.logger.Errorf("error fetching transaction count: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	status := Status{
		BlockCount:       blockCount,
		TransactionCount: transactionCount,
	}

	httpapi.Success(ctx, status)
}

type Status struct {
	BlockCount       int64 `json:"blockCount"`
	TransactionCount int64 `json:"transactionCount"`
	// TODO: Add more items when available in projections
}
