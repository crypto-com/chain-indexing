package handlers

import (
	"errors"
	"fmt"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	"github.com/crypto-com/chain-indexing/projection/bridge_activity/types"
	bridge_activitiy_view "github.com/crypto-com/chain-indexing/projection/bridge_activity/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
)

type Bridges struct {
	logger applogger.Logger

	bridgeActivitiesView bridge_activitiy_view.BridgeActivities
	accountAddressPrefix string
	config               BridgesConfig
}

func NewBridges(
	logger applogger.Logger,
	rdbHandle *rdb.Handle,
	cosmosAddressPrefix string,
	config BridgesConfig,
) *Bridges {
	return &Bridges{
		logger.WithFields(applogger.LogFields{
			"module": "BridgesHandler",
		}),

		bridge_activitiy_view.NewBridgeActivitiesView(rdbHandle),
		cosmosAddressPrefix,
		config,
	}
}

type BridgesConfig struct {
	Networks []BridgeNetworkConfig `mapstructure:"networks"`
	Chains   []BridgeChainConfig   `mapstructure:"chains"`
}

type BridgeNetworkConfig struct {
	ChainName string `mapstructure:"chain_name"`
	// Chain network unique abbreviation, used in URL query params
	Abbreviation     NetworkAbbreviation `mapstructure:"abbreviation"`
	MaybeAddressHook *AddressHook
}

type AddressHook = func(string) (string, error)

type BridgeChainConfig struct {
	Name       string                `mapstructure:"name"`
	Currencies []BridgeChainCurrency `mapstructure:"currencies"`
}

type BridgeChainCurrency struct {
	MinimalCoinDenom string `mapstructure:"minimal_coin_denom"`
	CoinDecimals     uint64 `mapstructure:"coin_decimals"`
	DisplayCoinDenom string `mapstructure:"display_coin_denom"`
}

type NetworkAbbreviation = string

func (handler *Bridges) FindByTransactionHash(ctx *fasthttp.RequestCtx) {
	hashParam, _ := ctx.UserValue("hash").(string)
	activity, activityErr := handler.bridgeActivitiesView.FindBy(bridge_activitiy_view.BridgeActivitiesFindByFilter{
		MaybeLinkId:        nil,
		MaybeTransactionId: &hashParam,
	})
	if activityErr != nil {
		if errors.Is(activityErr, rdb.ErrNoRows) {
			httpapi.NotFound(ctx)
			return
		}
		handler.logger.Errorf("error finding activity by transaction hash: %v", activityErr)
		httpapi.InternalServerError(ctx)
		return
	}

	displayAmount, maybeDisplayDenom := handler.toDisplayCoin(&activity)
	httpapi.Success(ctx, BridgeActivityResponseRow{
		BridgeActivityReadRow: activity,
		DisplayAmount:         displayAmount,
		MaybeDisplayDenom:     maybeDisplayDenom,
	})
}

func (handler *Bridges) ListActivities(ctx *fasthttp.RequestCtx) {
	var err error

	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	sourceBlockTimeOrder := view.ORDER_ASC
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("order") {
		orderStr := string(queryArgs.Peek("order"))
		if orderStr == "sourceBlockTime.desc" || orderStr == "sourceBlockHeight.desc" {
			sourceBlockTimeOrder = view.ORDER_DESC
		}
	}

	filter := bridge_activitiy_view.BridgeActivitiesListFilter{
		MaybeStatus:      nil,
		MaybeIdGt:        nil,
		MaybeCreatedAtLt: nil,
		MaybeCreatedAtGt: nil,
		MaybeUpdatedAtGt: nil,
	}
	if queryArgs.Has("status") {
		maybeStatus, maybeStatusErr := parseStatus(string(queryArgs.Peek("status")))
		if maybeStatusErr != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}
		filter.MaybeStatus = maybeStatus
	}
	if queryArgs.Has("id.gt") {
		filter.MaybeIdGt = primptr.String(string(queryArgs.Peek("id.gt")))
	}
	if queryArgs.Has("createdAt.lt") {
		maybeCreatedAtLt, createdAtLtParseErr := parseTimeFilter(string(queryArgs.Peek("createdAt.lt")))
		if createdAtLtParseErr != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}
		filter.MaybeCreatedAtLt = maybeCreatedAtLt
	}
	if queryArgs.Has("createdAt.ago") {
		ago, agoErr := time.ParseDuration(string(queryArgs.Peek("createdAt.ago")))
		if agoErr != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}
		filter.MaybeCreatedAtLt = primptr.UTCTime(utctime.Now().Add(-ago))
	}
	if queryArgs.Has("createdAt.gt") {
		maybeCreatedAtGt, createdAtGtParseErr := parseTimeFilter(string(queryArgs.Peek("createdAt.gt")))
		if createdAtGtParseErr != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}
		filter.MaybeCreatedAtGt = maybeCreatedAtGt
	}
	if queryArgs.Has("updatedAt.gt") {
		maybeUpdatedAtGt, updatedAtGtParseErr := parseTimeFilter(string(queryArgs.Peek("updatedAt.gt")))
		if updatedAtGtParseErr != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}
		filter.MaybeUpdatedAtGt = maybeUpdatedAtGt
	}

	addressFilter := make([]bridge_activitiy_view.BridgeActivitiesListAddressFilterCond, 0)
	for _, network := range handler.config.Networks {
		paramName := fmt.Sprintf("%sAddress", network.Abbreviation)
		if queryArgs.Has(paramName) {
			address := string(queryArgs.Peek(paramName))
			if network.MaybeAddressHook != nil {
				if parsedAddr, addrErr := network.MaybeAddressHook(address); addrErr != nil {
					handler.logger.Errorf("error converting address: %v", addrErr)
					httpapi.InternalServerError(ctx)
					return
				} else {
					address = parsedAddr
				}
			}

			addressFilter = append(addressFilter, bridge_activitiy_view.BridgeActivitiesListAddressFilterCond{
				Chain:   network.ChainName,
				Address: address,
			})
		}
	}

	order := bridge_activitiy_view.BridgeActivitiesListOrder{
		MaybeSourceBlockTime: &sourceBlockTimeOrder,
	}

	activities, paginationResult, listErr := handler.bridgeActivitiesView.List(
		addressFilter,
		filter,
		order,
		pagination,
	)
	if listErr != nil {
		handler.logger.Errorf("error listing bridge activities: %v", listErr)
		httpapi.InternalServerError(ctx)
		return
	}

	activityRespRows := make([]BridgeActivityResponseRow, 0)
	for i := range activities {
		displayAmount, maybeDisplayDenom := handler.toDisplayCoin(&activities[i])

		activityRespRows = append(activityRespRows, BridgeActivityResponseRow{
			BridgeActivityReadRow: activities[i],
			DisplayAmount:         displayAmount,
			MaybeDisplayDenom:     maybeDisplayDenom,
		})
	}

	httpapi.SuccessWithPagination(ctx, activityRespRows, paginationResult)
}

func (handler *Bridges) ListActivitiesByNetwork(ctx *fasthttp.RequestCtx) {
	var err error

	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	networkParam := ctx.UserValue("network").(string)
	accountParam := ctx.UserValue("account").(string)

	sourceBlockTimeOrder := view.ORDER_ASC
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("order") {
		orderStr := string(queryArgs.Peek("order"))
		if orderStr == "sourceBlockTime.desc" || orderStr == "sourceBlockHeight.desc" {
			sourceBlockTimeOrder = view.ORDER_DESC
		}
	}

	filter := bridge_activitiy_view.BridgeActivitiesListFilter{
		MaybeStatus:      nil,
		MaybeIdGt:        nil,
		MaybeCreatedAtLt: nil,
		MaybeCreatedAtGt: nil,
		MaybeUpdatedAtGt: nil,
	}
	if queryArgs.Has("filter.status") {
		maybeStatus, maybeStatusErr := parseStatus(string(queryArgs.Peek("filter.status")))
		if maybeStatusErr != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}
		filter.MaybeStatus = maybeStatus
	}
	if queryArgs.Has("id.gt") {
		filter.MaybeIdGt = primptr.String(string(queryArgs.Peek("id.gt")))
	}
	if queryArgs.Has("createdAt.lt") {
		maybeCreatedAtLt, createdAtLtParseErr := parseTimeFilter(string(queryArgs.Peek("createdAt.lt")))
		if createdAtLtParseErr != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}
		filter.MaybeCreatedAtLt = maybeCreatedAtLt
	}
	if queryArgs.Has("createdAt.ago") {
		ago, agoErr := time.ParseDuration(string(queryArgs.Peek("createdAt.ago")))
		if agoErr != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}
		filter.MaybeCreatedAtLt = primptr.UTCTime(utctime.Now().Add(-ago))
	}
	if queryArgs.Has("createdAt.gt") {
		maybeCreatedAtGt, createdAtGtParseErr := parseTimeFilter(string(queryArgs.Peek("createdAt.gt")))
		if createdAtGtParseErr != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}
		filter.MaybeCreatedAtGt = maybeCreatedAtGt
	}
	if queryArgs.Has("updatedAt.gt") {
		maybeUpdatedAtGt, updatedAtGtParseErr := parseTimeFilter(string(queryArgs.Peek("updatedAt.gt")))
		if updatedAtGtParseErr != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}
		filter.MaybeUpdatedAtGt = maybeUpdatedAtGt
	}

	var addressFilter []bridge_activitiy_view.BridgeActivitiesListAddressFilterCond
	for _, network := range handler.config.Networks {
		if network.Abbreviation == networkParam {
			address := accountParam
			if network.MaybeAddressHook != nil {
				if parsedAddr, addrErr := network.MaybeAddressHook(accountParam); addrErr != nil {
					handler.logger.Errorf("error running address hook: %v", addrErr)
					httpapi.InternalServerError(ctx)
					return
				} else {
					address = parsedAddr
				}
			}
			addressFilter = []bridge_activitiy_view.BridgeActivitiesListAddressFilterCond{
				{
					Chain:   network.ChainName,
					Address: address,
				},
			}
		}
	}

	order := bridge_activitiy_view.BridgeActivitiesListOrder{
		MaybeSourceBlockTime: &sourceBlockTimeOrder,
	}

	activities, paginationResult, listErr := handler.bridgeActivitiesView.List(
		addressFilter,
		filter,
		order,
		pagination,
	)
	if listErr != nil {
		handler.logger.Errorf("error listing bridge activities: %v", listErr)
		httpapi.InternalServerError(ctx)
		return
	}

	activityRespRows := make([]BridgeActivityResponseRow, 0)
	for i := range activities {
		displayAmount, maybeDisplayDenom := handler.toDisplayCoin(&activities[i])

		activityRespRows = append(activityRespRows, BridgeActivityResponseRow{
			BridgeActivityReadRow: activities[i],
			DisplayAmount:         displayAmount,
			MaybeDisplayDenom:     maybeDisplayDenom,
		})
	}

	httpapi.SuccessWithPagination(ctx, activityRespRows, paginationResult)
}

func (handler *Bridges) toDisplayCoin(
	activity *bridge_activitiy_view.BridgeActivityReadRow,
) (string, *string) {
	displayAmount := activity.Amount.String()
	displayDenom := activity.MaybeDenom
	if activity.MaybeDenom == nil {
		return displayAmount, displayDenom
	}

	for _, chain := range handler.config.Chains {
		if activity.SourceChain == chain.Name {
			for _, currency := range chain.Currencies {
				if currency.MinimalCoinDenom == *activity.MaybeDenom {
					displayAmount = activity.Amount.ToDec().Quo(
						coin.NewDec(10).Power(currency.CoinDecimals),
					).String()
					displayDenom = primptr.String(currency.DisplayCoinDenom)
				}
			}
		}
	}

	return displayAmount, displayDenom
}

type BridgeActivityResponseRow struct {
	bridge_activitiy_view.BridgeActivityReadRow

	DisplayAmount     string  `json:"displayAmount"`
	MaybeDisplayDenom *string `json:"displayDenom"`
}

func parseTimeFilter(value string) (*utctime.UTCTime, error) {
	if ut, utParseErr := utctime.Parse(time.RFC3339, value); utParseErr != nil {
		return nil, utParseErr
	} else {
		return &ut, nil
	}
}

func parseStatus(value string) (*types.Status, error) {
	switch value {
	case types.STATUS_PENDING:
		return primptr.String(types.STATUS_PENDING), nil
	case types.STATUS_CANCELLED:
		return primptr.String(types.STATUS_CANCELLED), nil
	case types.STATUS_COUNTERPARTY_CONFIRMED:
		return primptr.String(types.STATUS_COUNTERPARTY_CONFIRMED), nil
	case types.STATUS_FAILED_ON_CHAIN:
		return primptr.String(types.STATUS_FAILED_ON_CHAIN), nil
	case types.STATUS_COUNTERPARTY_REJECTED:
		return primptr.String(types.STATUS_COUNTERPARTY_REJECTED), nil
	default:
		return nil, errors.New("unrecognized status")
	}
}
