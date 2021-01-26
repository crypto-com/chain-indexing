package handlers

import (
	"errors"
	"strings"

	"github.com/crypto-com/chain-indexing/appinterface/cosmosapp"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/internal/primptr"

	"github.com/crypto-com/chain-indexing/projection/validator/constants"

	"github.com/valyala/fasthttp"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	validator_view "github.com/crypto-com/chain-indexing/projection/validator/view"
)

type Validators struct {
	logger applogger.Logger

	validatorAddressPrefix string
	consNodeAddressPrefix  string

	cosmosAppClient         cosmosapp.Client
	validatorsView          *validator_view.Validators
	validatorActivitiesView *validator_view.ValidatorActivities
}

func NewValidators(
	logger applogger.Logger,
	validatorAddressPrefix string,
	consNodeAddressPrefix string,
	cosmosAppClient cosmosapp.Client,
	rdbHandle *rdb.Handle,
) *Validators {
	return &Validators{
		logger.WithFields(applogger.LogFields{
			"module": "ValidatorsHandler",
		}),

		validatorAddressPrefix,
		consNodeAddressPrefix,

		cosmosAppClient,
		validator_view.NewValidators(rdbHandle),
		validator_view.NewValidatorActivities(rdbHandle),
	}
}

func (handler *Validators) FindBy(ctx *fasthttp.RequestCtx) {
	addressParams, _ := ctx.UserValue("address").(string)
	var identity validator_view.ValidatorIdentity
	if strings.HasPrefix(addressParams, handler.validatorAddressPrefix) {
		identity = validator_view.ValidatorIdentity{
			MaybeOperatorAddress: &addressParams,
		}
	} else if strings.HasPrefix(addressParams, handler.consNodeAddressPrefix) {
		identity = validator_view.ValidatorIdentity{
			MaybeConsensusNodeAddress: &addressParams,
		}
	}

	rawValidator, err := handler.validatorsView.FindBy(identity)
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
	} else {
		validator.SelfDelegation = delegation.Balance.Amount
	}

	httpapi.Success(ctx, validator)
}

func (handler *Validators) List(ctx *fasthttp.RequestCtx) {
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
		orderArg := string(queryArgs.Peek("order"))
		if orderArg == "power" {
			order.MaybePower = primptr.String(view.ORDER_ASC)
		} else if orderArg == "power.desc" {
			order.MaybePower = primptr.String(view.ORDER_DESC)
		} else {
			httpapi.BadRequest(ctx, errors.New("invalid order"))
			return
		}
	}

	validators, paginationResult, err := handler.validatorsView.List(
		validator_view.ValidatorsListFilter{}, order, pagination,
	)
	if err != nil {
		handler.logger.Errorf("error listing validators: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, validators, paginationResult)
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
		orderArg := string(queryArgs.Peek("order"))
		if orderArg == "power" {
			order.MaybePower = primptr.String(view.ORDER_ASC)
		} else if orderArg == "power.desc" {
			order.MaybePower = primptr.String(view.ORDER_DESC)
		} else {
			httpapi.BadRequest(ctx, errors.New("invalid order"))
			return
		}
	}

	validators, paginationResult, err := handler.validatorsView.List(validator_view.ValidatorsListFilter{
		MaybeStatuses: []constants.Status{
			constants.BONDED,
			constants.JAILED,
			constants.UNBONDING,
		},
	}, order, pagination)
	if err != nil {
		handler.logger.Errorf("error listing validators: %v", err)
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

	addressParams, _ := ctx.UserValue("address").(string)
	var filter validator_view.ValidatorActivitiesListFilter
	if strings.HasPrefix(addressParams, handler.validatorAddressPrefix) {
		filter = validator_view.ValidatorActivitiesListFilter{
			MaybeOperatorAddress: &addressParams,
		}
	} else if strings.HasPrefix(addressParams, handler.consNodeAddressPrefix) {
		filter = validator_view.ValidatorActivitiesListFilter{
			MaybeConsensusNodeAddress: &addressParams,
		}
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
