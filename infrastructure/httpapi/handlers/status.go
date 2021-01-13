package handlers

import (
	"strconv"

	status_polling "github.com/crypto-com/chain-indexing/appinterface/polling"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	block_view "github.com/crypto-com/chain-indexing/projection/block/view"
	transaction_view "github.com/crypto-com/chain-indexing/projection/transaction/view"
	"github.com/crypto-com/chain-indexing/projection/validator/constants"
	validator_view "github.com/crypto-com/chain-indexing/projection/validator/view"
	"github.com/crypto-com/chain-indexing/projection/validatorstats"
	validatorstats_view "github.com/crypto-com/chain-indexing/projection/validatorstats/view"
	"github.com/valyala/fasthttp"
)

type StatusHandler struct {
	logger applogger.Logger

	blocksView            *block_view.Blocks
	transactionsTotalView *transaction_view.TransactionsTotal
	validatorsView        *validator_view.Validators
	validatorStatsView    *validatorstats_view.ValidatorStats
	statusView            *status_polling.Status
}

func NewStatusHandler(logger applogger.Logger, rdbHandle *rdb.Handle) *StatusHandler {
	return &StatusHandler{
		logger.WithFields(applogger.LogFields{
			"module": "StatusHandler",
		}),

		block_view.NewBlocks(rdbHandle),
		transaction_view.NewTransactionsTotal(rdbHandle),
		validator_view.NewValidators(rdbHandle),
		validatorstats_view.NewValidatorStats(rdbHandle),
		status_polling.NewStatus(rdbHandle),
	}
}

func (handler *StatusHandler) GetStatus(ctx *fasthttp.RequestCtx) {
	blockCount, err := handler.blocksView.Count()
	if err != nil {
		handler.logger.Errorf("error fetching block count: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	transactionCount, err := handler.transactionsTotalView.FindBy("-")
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

	validatorCount, err := handler.validatorsView.Count(validator_view.CountFilter{
		MaybeStatus: nil,
	})
	if err != nil {
		handler.logger.Errorf("error fetching validators count: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	activeValidatorCount, err := handler.validatorsView.Count(validator_view.CountFilter{
		MaybeStatus: []string{constants.BONDED, constants.UNBONDING},
	})
	if err != nil {
		handler.logger.Errorf("error fetching active validators count: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	latestHeight, err := handler.statusView.FindBy("LatestHeight")
	if err != nil {
		handler.logger.Errorf("error fetching latest height: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	var latestHeightValue int64 = 0
	if n, err := strconv.ParseInt(latestHeight, 10, 64); err == nil {
		latestHeightValue = n
	} else {
		handler.logger.Errorf("error convert latest height from string to int64: %v", err)
		httpapi.InternalServerError(ctx)
	}

	status := Status{
		BlockCount:           blockCount,
		TransactionCount:     transactionCount,
		TotalDelegated:       totalDelegated,
		TotalReward:          totalReward,
		ValidatorCount:       validatorCount,
		ActiveValidatorCount: activeValidatorCount,
		LatestHeight:         latestHeightValue,
	}

	httpapi.Success(ctx, status)
}

type Status struct {
	BlockCount           int64  `json:"blockCount"`
	TransactionCount     int64  `json:"transactionCount"`
	TotalDelegated       string `json:"totalDelegated"`
	TotalReward          string `json:"totalReward"`
	ValidatorCount       int64  `json:"validatorCount"`
	ActiveValidatorCount int64  `json:"activeValidatorCount"`
	LatestHeight         int64  `json:"latestHeight"`
}
