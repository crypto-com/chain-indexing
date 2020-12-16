package handlers

import (
	block_view "github.com/crypto-com/chain-indexing/appinterface/projection/block/view"
	transaction_view "github.com/crypto-com/chain-indexing/appinterface/projection/transaction/view"
	validator_view "github.com/crypto-com/chain-indexing/appinterface/projection/validator/view"
	"github.com/crypto-com/chain-indexing/appinterface/projection/validatorstats"
	validatorstats_view "github.com/crypto-com/chain-indexing/appinterface/projection/validatorstats/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	"github.com/valyala/fasthttp"
)

type StatusHandler struct {
	logger applogger.Logger

	blocksView         *block_view.Blocks
	transactionsView   *transaction_view.BlockTransactions
	validatorsView     *validator_view.Validators
	validatorStatsView *validatorstats_view.ValidatorStats
}

func NewStatusHandler(logger applogger.Logger, rdbHandle *rdb.Handle) *StatusHandler {
	return &StatusHandler{
		logger.WithFields(applogger.LogFields{
			"module": "StatusHandler",
		}),

		block_view.NewBlocks(rdbHandle),
		transaction_view.NewTransactions(rdbHandle),
		validator_view.NewValidators(rdbHandle),
		validatorstats_view.NewValidatorStats(rdbHandle),
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

	totalDelegated, err := handler.validatorStatsView.FindBy(validatorstats.TOTAL_DELEGATE)
	if err != nil {
		handler.logger.Errorf("error fetching total delegate: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	totalReward, err := handler.validatorStatsView.FindBy(validatorstats.TOTAL_REWARD)
	if err != nil {
		handler.logger.Errorf("error fetching total reward: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	validatorCount, err := handler.validatorsView.Count()
	if err != nil {
		handler.logger.Errorf("error fetching validators count: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	status := Status{
		BlockCount:       blockCount,
		TransactionCount: transactionCount,
		TotalDelegated:   totalDelegated,
		TotalReward:      totalReward,
		ValiatorCount:    validatorCount,
	}

	httpapi.Success(ctx, status)
}

type Status struct {
	BlockCount       int64  `json:"blockCount"`
	TransactionCount int64  `json:"transactionCount"`
	TotalDelegated   string `json:"totalDelegated"`
	TotalReward      string `json:"totalReward"`
	ValiatorCount    int64  `json:"validatorCount"`
}
