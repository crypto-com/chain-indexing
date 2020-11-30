package handlers

import (
	"strings"

	"github.com/valyala/fasthttp"

	validator_view "github.com/crypto-com/chain-indexing/appinterface/projection/validator/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
)

type Validators struct {
	logger applogger.Logger

	validatorAddressPrefix string
	consNodeAddressPrefix  string

	validatorsView          *validator_view.Validators
	validatorActivitiesView *validator_view.ValidatorActivities
}

func NewValidators(
	logger applogger.Logger,
	validatorAddressPrefix string,
	consNodeAddressPrefix string,
	rdbHandle *rdb.Handle,
) *Validators {
	return &Validators{
		logger.WithFields(applogger.LogFields{
			"module": "ValidatorsHandler",
		}),

		validatorAddressPrefix,
		consNodeAddressPrefix,

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

	validator, err := handler.validatorsView.FindBy(identity)
	if err != nil {
		if err == rdb.ErrNoRows {
			httpapi.NotFound(ctx)
			return
		}
		handler.logger.Errorf("error finding validator by operator address: %v", err)
		httpapi.InternalServerError(ctx)
		return
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

	validators, paginationResult, err := handler.validatorsView.List(pagination)
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

	blocks, paginationResult, err := handler.validatorActivitiesView.List(filter, pagination)
	if err != nil {
		handler.logger.Errorf("error listing activities: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, blocks, paginationResult)
}
