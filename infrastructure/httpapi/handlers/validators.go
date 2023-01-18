package handlers

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/projection/chainstats"

	"github.com/valyala/fasthttp"

	"github.com/crypto-com/chain-indexing/appinterface/cosmosapp"
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/appinterface/tendermint"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	block_view "github.com/crypto-com/chain-indexing/projection/block/view"
	chainstats_view "github.com/crypto-com/chain-indexing/projection/chainstats/view"
	"github.com/crypto-com/chain-indexing/projection/validator/constants"
	validator_view "github.com/crypto-com/chain-indexing/projection/validator/view"
)

// When we have a large number of blocks, we would like only take the recent N most blocks (blocks in 7 recent days),
// in order to calculate the averageBlockTime.
//
// Assume average block generation time is 6 seconds per block.
// Then in recent 7 days, number of estimated generated block will be:
//
// nRecentBlocks: n (block) = 7(day) * 24(hour/day) * 3600(sec/hour) / 6(sec/block)
//
const nRecentBlocksInInt = 100800

type Validators struct {
	logger applogger.Logger

	validatorAddressPrefix     string
	consNodeAddressPrefix      string
	maxActiveBlocksPeriodLimit int64

	cosmosAppClient         cosmosapp.Client
	tendermintClient        tendermint.Client
	validatorsView          validator_view.Validators
	validatorActivitiesView validator_view.ValidatorActivities
	chainStatsView          *chainstats_view.ChainStats
	blockView               *block_view.Blocks

	globalAPY              *big.Float
	globalAPYLastUpdatedAt time.Time
}

func NewValidators(
	logger applogger.Logger,
	validatorAddressPrefix string,
	consNodeAddressPrefix string,
	maxActiveBlocksPeriodLimit int64,
	cosmosAppClient cosmosapp.Client,
	tendermintClient tendermint.Client,
	rdbHandle *rdb.Handle,
) *Validators {
	return &Validators{
		logger.WithFields(applogger.LogFields{
			"module": "ValidatorsHandler",
		}),

		validatorAddressPrefix,
		consNodeAddressPrefix,
		maxActiveBlocksPeriodLimit,

		cosmosAppClient,
		tendermintClient,
		validator_view.NewValidatorsView(rdbHandle),
		validator_view.NewValidatorActivitiesView(rdbHandle),
		chainstats_view.NewChainStats(rdbHandle),
		block_view.NewBlocks(rdbHandle),

		nil,
		time.Unix(int64(0), int64(0)),
	}
}

func (handler *Validators) FindBy(ctx *fasthttp.RequestCtx) {
	addressParams, addressParamsOk := URLValueGuard(ctx, handler.logger, "address")
	if !addressParamsOk {
		return
	}
	addressParams = strings.ToLower(addressParams)

	recentBlocks := handler.maxActiveBlocksPeriodLimit
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("recentBlocks") {
		recentBlocksParam, err := queryArgs.GetUint("recentBlocks")
		if err != nil {
			httpapi.BadRequest(ctx, errors.New("invalid recentBlocks"))
			return
		}

		if int64(recentBlocksParam) > handler.maxActiveBlocksPeriodLimit {
			httpapi.BadRequest(ctx, errors.New("recentBlocks excess limit"))
			return
		}

		recentBlocks = int64(recentBlocksParam)
	}

	var identity validator_view.ValidatorIdentity
	if strings.HasPrefix(addressParams, handler.validatorAddressPrefix) {
		identity = validator_view.ValidatorIdentity{
			MaybeOperatorAddress: &addressParams,
		}
	} else if strings.HasPrefix(addressParams, handler.consNodeAddressPrefix) {
		identity = validator_view.ValidatorIdentity{
			MaybeConsensusNodeAddress: &addressParams,
		}
	} else {
		httpapi.BadRequest(ctx, errors.New("invalid address"))
		return
	}

	lastHandledHeight, err := handler.validatorsView.LastHandledHeight()
	if err != nil {
		handler.logger.Errorf("error finding validator by operator address: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	lowestActiveBlockHeight := lastHandledHeight - recentBlocks

	rawValidator, err := handler.validatorsView.FindBy(identity, &lowestActiveBlockHeight)
	if err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			httpapi.NotFound(ctx)
			return
		}
		handler.logger.Errorf("error finding validator by operator address: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	validator := ValidatorDetails{
		ValidatorRow: rawValidator,

		Tokens:         "0",
		SelfDelegation: "0",
	}

	validatorData, err := handler.cosmosAppClient.Validator(validator.OperatorAddress)
	if err != nil {
		handler.logger.Errorf("error getting validator details: %v", err)
	} else {
		validator.Tokens = validatorData.Tokens
	}

	delegation, err := handler.cosmosAppClient.Delegation(validator.InitialDelegatorAddress, validator.OperatorAddress)
	if err != nil {
		handler.logger.Errorf("error getting self delegation record: %v", err)
	} else if delegation != nil {
		validator.SelfDelegation = delegation.Balance.Amount
	}

	httpapi.Success(ctx, validator)
}

func (handler *Validators) List(ctx *fasthttp.RequestCtx) {
	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	queryArgs := ctx.QueryArgs()
	order := validator_view.ValidatorsListOrder{
		MaybeStatus:              primptr.String(view.ORDER_ASC),
		MaybeJoinedAtBlockHeight: primptr.String(view.ORDER_ASC),
	}
	if queryArgs.Has("order") {
		rawOrderArgs := queryArgs.PeekMulti("order")
		for _, rawOrderArg := range rawOrderArgs {
			orderArg := string(rawOrderArg)
			if orderArg == "power" {
				order.MaybePower = primptr.String(view.ORDER_ASC)
			} else if orderArg == "power.desc" {
				order.MaybePower = primptr.String(view.ORDER_DESC)
			} else if orderArg == "commission" {
				order.MaybeCommission = primptr.String(view.ORDER_ASC)
			} else if orderArg == "commission.desc" {
				order.MaybeCommission = primptr.String(view.ORDER_DESC)
			} else {
				handler.logger.Errorf("error listing validators: invalid order: %s", orderArg)
				httpapi.BadRequest(ctx, errors.New("invalid order"))
				return
			}
		}
	}

	recentBlocks := int64(0)
	var recentBlocksParam int
	if queryArgs.Has("recentBlocks") {
		recentBlocksParam, err = queryArgs.GetUint("recentBlocks")
		if err != nil {
			httpapi.BadRequest(ctx, errors.New("invalid recentBlocks"))
			return
		}

		if int64(recentBlocksParam) > handler.maxActiveBlocksPeriodLimit {
			httpapi.BadRequest(ctx, fmt.Errorf("recentBlocks excess limit(%d)", handler.maxActiveBlocksPeriodLimit))
			return
		}

		recentBlocks = int64(recentBlocksParam)
	}

	lastHandledHeight, err := handler.validatorsView.LastHandledHeight()
	if err != nil {
		handler.logger.Errorf("error finding validator by operator address: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	filter := validator_view.ValidatorsListFilter{}
	var lowestActiveBlockHeight int64

	if recentBlocks > 0 {
		lowestActiveBlockHeight = lastHandledHeight - recentBlocks
		filter = validator_view.ValidatorsListFilter{
			MaybeRecentActiveBlock: &validator_view.ValidatorsListRecentActiveBlockFilter{
				MaybeLowestActiveBlockHeight: lowestActiveBlockHeight,
			},
		}
	}

	validators, paginationResult, err := handler.validatorsView.List(
		filter, order, pagination,
	)
	if err != nil {
		handler.logger.Errorf("error listing validators: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	if handler.globalAPYLastUpdatedAt.Add(1 * time.Hour).Before(time.Now()) {
		handler.logger.Info("going to fetch latest global APY")
		handler.globalAPY, err = handler.getGlobalAPY()
		if err != nil {
			handler.logger.Errorf("error getting global APY: %v", err)
			httpapi.InternalServerError(ctx)
			return
		}
		handler.globalAPYLastUpdatedAt = time.Now()
	}
	validatorsWithAPY := make([]validatorRowWithAPY, 0, len(validators))
	for _, validator := range validators {
		if validator.Status != constants.BONDED {
			validatorsWithAPY = append(validatorsWithAPY, validatorRowWithAPY{
				validator,
				"0",
			})
			continue
		}
		commissionRate, commissionRateOk := new(big.Float).SetString(validator.CommissionRate)
		if !commissionRateOk {
			handler.logger.Errorf("error parsing validator commission rate: %s", validator.CommissionRate)
			httpapi.InternalServerError(ctx)
			return
		}
		afterCommission := new(big.Float).Sub(
			new(big.Float).SetInt64(int64(1)),
			commissionRate,
		)
		apy := new(big.Float).Mul(handler.globalAPY, afterCommission)
		validatorsWithAPY = append(validatorsWithAPY, validatorRowWithAPY{
			validator,
			apy.Text('f', -1),
		})
	}

	httpapi.SuccessWithPagination(ctx, validatorsWithAPY, paginationResult)
}

type validatorRowWithAPY struct {
	validator_view.ListValidatorRow

	APY string `json:"apy"`
}

func (handler *Validators) getGlobalAPY() (*big.Float, error) {
	var err error

	// TODO: should use annual provisions and total bonded from validator and validatorstats
	annualProvisions, err := handler.cosmosAppClient.AnnualProvisions()
	if err != nil {
		return nil, fmt.Errorf("error fetching annual provisions: %v", err)
	}

	totalBonded, err := handler.cosmosAppClient.TotalBondedBalance()
	if err != nil {
		return nil, fmt.Errorf("error fetching total bonded: %v", err)
	}

	// estimated APY = expected APY * estimated block count / actual block count
	genesis, err := handler.tendermintClient.GenesisChunked()
	if err != nil {
		return nil, fmt.Errorf("error fetching genesis: %v", err)
	}
	blockPerYearParam, blockPerYearParamOk := new(big.Float).SetString(genesis.AppState.Mint.Params.BlocksPerYear)
	if !blockPerYearParamOk {
		return nil, fmt.Errorf("error parsing block per year param: %s", genesis.AppState.Mint.Params.BlocksPerYear)
	}

	averageBlockTime, err := handler.getAverageBlockTime()
	if err != nil {
		return nil, fmt.Errorf("error fetching average block time: %v", err)
	}

	totalSecondsPerYear := big.NewFloat(31556952)
	estimatedBlockPerYear := new(big.Float).Quo(
		totalSecondsPerYear,
		averageBlockTime,
	)

	annualProvisionsBigFloat, annualProvisionsOk := new(big.Float).SetString(annualProvisions.Amount.String())
	if !annualProvisionsOk {
		return nil, fmt.Errorf("error parsing annual provisions: %s", annualProvisions.Amount.String())
	}
	expectedAPY := new(big.Float).Quo(
		annualProvisionsBigFloat,
		new(big.Float).SetInt(totalBonded.Amount.BigInt()),
	)

	estimatedAPY := new(big.Float).Mul(
		expectedAPY,
		new(big.Float).Quo(
			estimatedBlockPerYear,
			blockPerYearParam,
		),
	)

	return estimatedAPY, nil
}

func (handler *Validators) getAverageBlockTime() (*big.Float, error) {
	// Average block time calculation
	//
	// Case A: totalBlockCount <= nRecentBlocks, calculate with blocks from Genesis block to Latest block
	// Case B: totalBlockCount > nRecentBlocks, calculate with n recent blocks
	var totalBlockTime = big.NewInt(0)
	var totalBlockCount = big.NewInt(1)

	if rawTotalBlockTime, err := handler.chainStatsView.FindBy(chainstats.TOTAL_BLOCK_TIME); err != nil {
		return nil, fmt.Errorf("error fetching total block time: %v", err)
	} else {
		if rawTotalBlockTime != "" {
			var ok bool
			if totalBlockTime, ok = new(big.Int).SetString(rawTotalBlockTime, 10); !ok {
				return nil, errors.New("error converting total block time from string to big.Int")
			}
		}

		if rawTotalBlockCount, err := handler.chainStatsView.FindBy(chainstats.TOTAL_BLOCK_COUNT); err != nil {
			return nil, fmt.Errorf("error fetching total block time: %v", err)
		} else {
			if rawTotalBlockCount != "" {
				var ok bool
				if totalBlockCount, ok = new(big.Int).SetString(rawTotalBlockCount, 10); !ok {
					return nil, fmt.Errorf("error converting total block count from string to big.Int")
				}
			}
		}
	}

	nRecentBlocks := big.NewInt(nRecentBlocksInInt)

	hasNBlocksSinceGenesis := (totalBlockCount.Cmp(nRecentBlocks) == 1)

	var averageBlockTimeMilliSecond *big.Float
	// Determine case A or case B
	if hasNBlocksSinceGenesis {
		// Case B
		latestBlockHeight, err := handler.blockView.Count()
		if err != nil {
			return nil, fmt.Errorf("error fetching latest block height: %v", err)
		}
		latestBlockIdentity := block_view.BlockIdentity{MaybeHeight: &latestBlockHeight}
		latestBlock, err := handler.blockView.FindBy(&latestBlockIdentity)
		if err != nil {
			return nil, fmt.Errorf("error fetching latest block: %v", err)
		}

		// Find the nth block before latest block
		startBlockHeight := latestBlockHeight - nRecentBlocks.Int64()
		startBlockIdentity := block_view.BlockIdentity{MaybeHeight: &startBlockHeight}
		startBlock, err := handler.blockView.FindBy(&startBlockIdentity)
		if err != nil {
			return nil, fmt.Errorf("error fetching the start block: %v", err)
		}

		// Calculate total time in generating n recent blocks
		nRecentBlocksTotalTime := latestBlock.Time.UnixNano() - startBlock.Time.UnixNano()

		nRecentBlocksTotalTimeMilliSecond := new(big.Float).Quo(
			new(big.Float).SetInt64(nRecentBlocksTotalTime),
			new(big.Float).SetInt64(int64(1000000)),
		)
		averageBlockTimeMilliSecond = new(big.Float).Quo(
			nRecentBlocksTotalTimeMilliSecond,
			new(big.Float).SetInt(nRecentBlocks),
		)

	} else {
		// Case A
		totalBlockTimeMilliSecond := new(big.Float).Quo(
			new(big.Float).SetInt(totalBlockTime),
			new(big.Float).SetInt64(int64(1000000)),
		)
		averageBlockTimeMilliSecond = new(big.Float).Quo(
			totalBlockTimeMilliSecond,
			new(big.Float).SetInt(totalBlockCount),
		)
	}

	averageBlockTime := new(big.Float).Quo(
		averageBlockTimeMilliSecond,
		big.NewFloat(1000),
	)

	return averageBlockTime, nil
}

func (handler *Validators) ListActive(ctx *fasthttp.RequestCtx) {
	var err error

	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	queryArgs := ctx.QueryArgs()
	order := validator_view.ValidatorsListOrder{
		MaybeStatus:              primptr.String(view.ORDER_ASC),
		MaybeJoinedAtBlockHeight: primptr.String(view.ORDER_ASC),
	}
	if queryArgs.Has("order") {
		rawOrderArgs := queryArgs.PeekMulti("order")
		for _, rawOrderArg := range rawOrderArgs {
			orderArg := string(rawOrderArg)
			if orderArg == "power" {
				order.MaybePower = primptr.String(view.ORDER_ASC)
			} else if orderArg == "power.desc" {
				order.MaybePower = primptr.String(view.ORDER_DESC)
			} else if orderArg == "commission" {
				order.MaybeCommission = primptr.String(view.ORDER_ASC)
			} else if orderArg == "commission.desc" {
				order.MaybeCommission = primptr.String(view.ORDER_DESC)
			} else {
				handler.logger.Errorf("error listing active validators: invalid order: %s", orderArg)
				httpapi.BadRequest(ctx, errors.New("invalid order"))
				return
			}
		}
	}

	recentBlocks := handler.maxActiveBlocksPeriodLimit
	var recentBlocksParam int
	if queryArgs.Has("recentBlocks") {
		recentBlocksParam, err = queryArgs.GetUint("recentBlocks")
		if err != nil {
			httpapi.BadRequest(ctx, errors.New("invalid recentBlocks"))
			return
		}

		if int64(recentBlocksParam) > handler.maxActiveBlocksPeriodLimit {
			httpapi.BadRequest(ctx, errors.New("recentBlocks excess limit"))
			return
		}

		recentBlocks = int64(recentBlocksParam)
	}

	lastHandledHeight, err := handler.validatorsView.LastHandledHeight()
	if err != nil {
		handler.logger.Errorf("error finding validator by operator address: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	lowestActiveBlockHeight := lastHandledHeight - recentBlocks

	validators, paginationResult, err := handler.validatorsView.List(
		validator_view.ValidatorsListFilter{
			MaybeStatuses: []constants.Status{
				constants.BONDED,
				constants.JAILED,
				constants.UNBONDING,
			},
			MaybeRecentActiveBlock: &validator_view.ValidatorsListRecentActiveBlockFilter{
				MaybeLowestActiveBlockHeight: lowestActiveBlockHeight,
			},
		},
		order,
		pagination,
	)
	if err != nil {
		handler.logger.Errorf("error listing active validators: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, validators, paginationResult)
}

func (handler *Validators) ListActivities(ctx *fasthttp.RequestCtx) {
	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		httpapi.BadRequest(ctx, err)
		return
	}

	addressParams, addressParamsOk := URLValueGuard(ctx, handler.logger, "address")
	if !addressParamsOk {
		return
	}

	addressParams = strings.ToLower(addressParams)

	var filter validator_view.ValidatorActivitiesListFilter
	if strings.HasPrefix(addressParams, handler.validatorAddressPrefix) {
		filter = validator_view.ValidatorActivitiesListFilter{
			MaybeOperatorAddress: &addressParams,
		}
	} else if strings.HasPrefix(addressParams, handler.consNodeAddressPrefix) {
		filter = validator_view.ValidatorActivitiesListFilter{
			MaybeConsensusNodeAddress: &addressParams,
		}
	} else {
		httpapi.BadRequest(ctx, errors.New("invalid address"))
		return
	}

	order := validator_view.ValidatorActivitiesListOrder{}
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("order") {
		orderArg := string(queryArgs.Peek("order"))
		if orderArg == "height" {
			order.MaybeBlockHeight = primptr.String(view.ORDER_ASC)
		} else if orderArg == "height.desc" {
			order.MaybeBlockHeight = primptr.String(view.ORDER_DESC)
		} else {
			httpapi.BadRequest(ctx, errors.New("invalid order"))
			return
		}
	}

	blocks, paginationResult, err := handler.validatorActivitiesView.List(filter, order, pagination)
	if err != nil {
		handler.logger.Errorf("error listing activities: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, blocks, paginationResult)
}

type ValidatorDetails struct {
	*validator_view.ValidatorRow

	Tokens         string `json:"tokens"`
	SelfDelegation string `json:"selfDelegation"`
}
