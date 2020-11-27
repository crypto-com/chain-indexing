package handlers

import (
	view2 "github.com/crypto-com/chainindex/appinterface/projection/validator/view"
	"github.com/valyala/fasthttp"

	"github.com/crypto-com/chainindex/appinterface/rdb"
	"github.com/crypto-com/chainindex/infrastructure/httpapi"
	applogger "github.com/crypto-com/chainindex/internal/logger"
)

type Validators struct {
	logger applogger.Logger

	validatorsView          *view2.Validators
	validatorActivitiesView *view2.ValidatorActivities
}

func NewValidators(logger applogger.Logger, rdbHandle *rdb.Handle) *Validators {
	return &Validators{
		logger.WithFields(applogger.LogFields{
			"module": "ValidatorsHandler",
		}),

		view2.NewValidators(rdbHandle),
		view2.NewValidatorActivities(rdbHandle),
	}
}

func (handler *Validators) FindBy(ctx *fasthttp.RequestCtx) {
	// TODO: support find by consensus node address
	operatorAddress, _ := ctx.UserValue("operator_address").(string)

	identity := view2.ValidatorIdentity{
		MaybeOperatorAddress: &operatorAddress,
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

	operatorAddress, _ := ctx.UserValue("operator_address").(string)

	blocks, paginationResult, err := handler.validatorActivitiesView.List(view2.ValidatorActivitiesListFilter{
		MaybeOperatorAddress: &operatorAddress,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing transactions: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, blocks, paginationResult)
}
