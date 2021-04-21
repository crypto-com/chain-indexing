package handlers

import (
	"math/big"
	"strconv"
	"time"

	"github.com/crypto-com/chain-indexing/appinterface/cosmosapp"

	"github.com/crypto-com/chain-indexing/projection/chainstats"

	"github.com/crypto-com/chain-indexing/internal/json"
	"github.com/crypto-com/chain-indexing/usecase/coin"

	status_polling "github.com/crypto-com/chain-indexing/appinterface/polling"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	block_view "github.com/crypto-com/chain-indexing/projection/block/view"
	chainstats_view "github.com/crypto-com/chain-indexing/projection/chainstats/view"
	transaction_view "github.com/crypto-com/chain-indexing/projection/transaction/view"
	"github.com/crypto-com/chain-indexing/projection/validator/constants"
	validator_view "github.com/crypto-com/chain-indexing/projection/validator/view"
	"github.com/crypto-com/chain-indexing/projection/validatorstats"
	validatorstats_view "github.com/crypto-com/chain-indexing/projection/validatorstats/view"
	"github.com/valyala/fasthttp"
)

type StatusHandler struct {
	logger applogger.Logger

	cosmosAppClient       cosmosapp.Client
	blocksView            *block_view.Blocks
	chainStatsView        *chainstats_view.ChainStats
	transactionsTotalView *transaction_view.TransactionsTotal
	validatorsView        *validator_view.Validators
	validatorStatsView    *validatorstats_view.ValidatorStats
	statusView            *status_polling.Status

	totalDelegated              coin.Coins
	totalDelegatedLastUpdatedAt time.Time
}

func NewStatusHandler(logger applogger.Logger, cosmosAppClient cosmosapp.Client, rdbHandle *rdb.Handle) *StatusHandler {
	return &StatusHandler{
		logger.WithFields(applogger.LogFields{
			"module": "StatusHandler",
		}),

		cosmosAppClient,
		block_view.NewBlocks(rdbHandle),
		chainstats_view.NewChainStats(rdbHandle),
		transaction_view.NewTransactionsTotal(rdbHandle),
		validator_view.NewValidators(rdbHandle),
		validatorstats_view.NewValidatorStats(rdbHandle),
		status_polling.NewStatus(rdbHandle),

		coin.NewEmptyCoins(),
		time.Unix(int64(0), int64(0)),
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

	// TODO: https://github.com/crypto-com/chain-indexing/issues/386
	//rawTotalDelegated, err := handler.validatorStatsView.FindBy(validatorstats.TOTAL_DELEGATE)
	//if err != nil {
	//	handler.logger.Errorf("error fetching total delegate: %v", err)
	//	httpapi.InternalServerError(ctx)
	//	return
	//}
	//var totalDelegated coin.Coins
	//if rawTotalDelegated != "" {
	//	json.MustUnmarshalFromString(rawTotalDelegated, &totalDelegated)
	//}
	if handler.totalDelegatedLastUpdatedAt.Add(15 * time.Minute).Before(time.Now()) {
		totalBondedBalance, totalBondedBalanceErr := handler.cosmosAppClient.TotalBondedBalance()
		if totalBondedBalanceErr != nil {
			handler.logger.Errorf("error fetching total delegate: %v", totalBondedBalanceErr)
			httpapi.InternalServerError(ctx)
			return
		}
		handler.totalDelegated = coin.NewCoins(totalBondedBalance)
	}

	rawTotalReward, err := handler.validatorStatsView.FindBy(validatorstats.TOTAL_REWARD)
	if err != nil {
		handler.logger.Errorf("error fetching total reward: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}
	var totalReward coin.DecCoins
	if rawTotalReward != "" {
		json.MustUnmarshalFromString(rawTotalReward, &totalReward)
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

	rawLatestHeight, err := handler.statusView.FindBy("LatestHeight")
	if err != nil {
		handler.logger.Errorf("error fetching latest height: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	var latestHeight int64 = 0
	if rawLatestHeight != "" {
		// TODO: Use big.Int
		if n, err := strconv.ParseInt(rawLatestHeight, 10, 64); err != nil {
			handler.logger.Errorf("error converting latest height from string to int64: %v", err)
			httpapi.InternalServerError(ctx)
			return
		} else {
			latestHeight = n
		}
	}

	var totalBlockTime = new(big.Int).SetInt64(int64(0))
	var totalBlockCount = new(big.Int).SetInt64(int64(1))
	if rawTotalBlockTime, err := handler.chainStatsView.FindBy(chainstats.TOTAL_BLOCK_TIME); err != nil {
		handler.logger.Errorf("error fetching total block time: %v", err)
		httpapi.InternalServerError(ctx)
		return
	} else {
		if rawTotalBlockTime != "" {
			var ok bool
			if totalBlockTime, ok = new(big.Int).SetString(rawTotalBlockTime, 10); !ok {
				handler.logger.Error("error converting total block time from string to big.Int")
				httpapi.InternalServerError(ctx)
				return
			}
		}

		if rawTotalBlockCount, err := handler.chainStatsView.FindBy(chainstats.TOTAL_BLOCK_COUNT); err != nil {
			handler.logger.Errorf("error fetching total block time: %v", err)
			httpapi.InternalServerError(ctx)
			return
		} else {
			if rawTotalBlockCount != "" {
				var ok bool
				if totalBlockCount, ok = new(big.Int).SetString(rawTotalBlockCount, 10); !ok {
					handler.logger.Error("error converting total block count from string to big.Int")
					httpapi.InternalServerError(ctx)
					return
				}
			}
		}
	}

	totalBlockTimeMilliSecond := new(big.Float).Quo(
		new(big.Float).SetInt(totalBlockTime),
		new(big.Float).SetInt64(int64(1000000)),
	)
	averageBlockTime := new(big.Float).Quo(
		totalBlockTimeMilliSecond,
		new(big.Float).SetInt(totalBlockCount),
	)

	status := Status{
		BlockCount:                  blockCount,
		TransactionCount:            transactionCount,
		TotalDelegated:              handler.totalDelegated,
		TotalReward:                 totalReward,
		ValidatorCount:              validatorCount,
		ActiveValidatorCount:        activeValidatorCount,
		LatestHeight:                latestHeight,
		AverageBlockTimeMillisecond: averageBlockTime.Text('f', 0),
	}

	httpapi.Success(ctx, status)
}

type Status struct {
	BlockCount                  int64         `json:"blockCount"`
	TransactionCount            int64         `json:"transactionCount"`
	TotalDelegated              coin.Coins    `json:"totalDelegated"`
	TotalReward                 coin.DecCoins `json:"totalReward"`
	ValidatorCount              int64         `json:"validatorCount"`
	ActiveValidatorCount        int64         `json:"activeValidatorCount"`
	LatestHeight                int64         `json:"latestHeight"`
	AverageBlockTimeMillisecond string        `json:"averageBlockTimeMillisecond"`
}
