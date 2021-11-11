package handlers

import (
	"errors"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/valyala/fasthttp"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/external/tmcosmosutils"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	"github.com/crypto-com/chain-indexing/projection/bridge_activity/types"
	bridge_activitiy_view "github.com/crypto-com/chain-indexing/projection/bridge_activity/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
)

type Bridges struct {
	logger applogger.Logger

	bridgeActivitiesView *bridge_activitiy_view.BridgeActivities
	accountAddressPrefix string
}

func NewBridges(logger applogger.Logger, rdbHandle *rdb.Handle, cosmosAddressPrefix string) *Bridges {
	return &Bridges{
		logger.WithFields(applogger.LogFields{
			"module": "BridgesHandler",
		}),

		bridge_activitiy_view.NewBridgeActivities(rdbHandle),
		cosmosAddressPrefix,
	}
}

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

	displayAmount, maybeDisplayDenom := toDisplayCoin(activity.Amount, activity.MaybeDenom)
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

	addressFilter := bridge_activitiy_view.BridgeActivitiesListAddressFilter{
		MaybeCronosAddress:         nil,
		MaybeCryptoOrgChainAddress: nil,
	}
	if queryArgs.Has("cronosAddress") {
		addressFilter.MaybeCronosAddress = primptr.String(string(queryArgs.Peek("cronosAddress")))
	}
	if queryArgs.Has("cronosevmAddress") {
		hexAddress := string(queryArgs.Peek("cronosevmAddress"))
		var addr []byte
		if strings.HasPrefix(hexAddress, "0x") {
			addr = common.HexToAddress(hexAddress[2:]).Bytes()
		} else {
			addr = common.HexToAddress(hexAddress).Bytes()
		}
		accountAddr, accountAddrErr := tmcosmosutils.AccountAddressFromBytes(handler.accountAddressPrefix, addr)
		if accountAddrErr != nil {
			handler.logger.Errorf("error converting cronosevm address to account address: %v", accountAddrErr)
			httpapi.InternalServerError(ctx)
			return
		}
		addressFilter.MaybeCronosAddress = primptr.String(accountAddr)
	}
	if queryArgs.Has("cryptoorgchainAddress") {
		addressFilter.MaybeCryptoOrgChainAddress = primptr.String(string(queryArgs.Peek("cryptoorgchainAddress")))
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
	for i, activity := range activities {
		displayAmount, maybeDisplayDenom := toDisplayCoin(activity.Amount, activity.MaybeDenom)

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

	addressFilter := bridge_activitiy_view.BridgeActivitiesListAddressFilter{
		MaybeCronosAddress:         nil,
		MaybeCryptoOrgChainAddress: nil,
	}
	if networkParam == "cronos" {
		addressFilter.MaybeCronosAddress = primptr.String(accountParam)
	} else if networkParam == "cronosevm" {
		hexAddress := accountParam
		var addr []byte
		if strings.HasPrefix(hexAddress, "0x") {
			addr = common.HexToAddress(hexAddress[2:]).Bytes()
		} else {
			addr = common.HexToAddress(hexAddress).Bytes()
		}
		accountAddr, accountAddrErr := tmcosmosutils.AccountAddressFromBytes(handler.accountAddressPrefix, addr)
		if accountAddrErr != nil {
			handler.logger.Errorf("error converting cronosevm address to account address: %v", accountAddrErr)
			httpapi.InternalServerError(ctx)
			return
		}
		addressFilter.MaybeCronosAddress = primptr.String(accountAddr)
	} else if networkParam == "cryptoorgchain" {
		addressFilter.MaybeCryptoOrgChainAddress = primptr.String(accountParam)
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
	for i, activity := range activities {
		displayAmount, maybeDisplayDenom := toDisplayCoin(activity.Amount, activity.MaybeDenom)

		activityRespRows = append(activityRespRows, BridgeActivityResponseRow{
			BridgeActivityReadRow: activities[i],
			DisplayAmount:         displayAmount,
			MaybeDisplayDenom:     maybeDisplayDenom,
		})
	}

	httpapi.SuccessWithPagination(ctx, activityRespRows, paginationResult)
}

func toDisplayCoin(amount coin.Int, maybeDenom *string) (string, *string) {
	displayAmount := amount.String()
	displayDenom := maybeDenom
	if maybeDenom == nil {
		return displayAmount, displayDenom
	}

	parts := strings.Split(*maybeDenom, "/")
	lastPart := parts[len(parts)-1]
	if lastPart == "basecro" {
		displayAmount = amount.ToDec().Quo(coin.NewDec(100000000)).String()
		displayDenom = primptr.String("CRO")
	} else if lastPart == "basetcro" {
		displayAmount = amount.ToDec().Quo(coin.NewDec(100000000)).String()
		displayDenom = primptr.String("TCRO")
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
	case types.STATUS_NO_OPERATION:
		return primptr.String(types.STATUS_NO_OPERATION), nil
	default:
		return nil, errors.New("unrecognized status")
	}
}
