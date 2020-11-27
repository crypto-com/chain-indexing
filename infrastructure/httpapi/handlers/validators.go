package handlers

import (
	"github.com/valyala/fasthttp"

	"github.com/crypto-com/chainindex/appinterface/projection/view"
	"github.com/crypto-com/chainindex/appinterface/rdb"
	"github.com/crypto-com/chainindex/infrastructure/httpapi"
	applogger "github.com/crypto-com/chainindex/internal/logger"
)

type Validators struct {
	logger applogger.Logger

	validatorsView *view.Validators
}

func NewValidators(logger applogger.Logger, rdbHandle *rdb.Handle) *Validators {
	return &Validators{
		logger.WithFields(applogger.LogFields{
			"module": "ValidatorsHandler",
		}),

		view.NewValidators(rdbHandle),
	}
}

func (handler *Validators) FindBy(ctx *fasthttp.RequestCtx) {
	// TODO: support find by consensus node address
	operatorAddress, _ := ctx.UserValue("operator_address").(string)

	identity := view.ValidatorIdentity{
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

func (handler *Validators) ListActivities(_ *fasthttp.RequestCtx) {
	//pagination, err := httpapi.ParsePagination(ctx)
	//if err != nil {
	//    httpapi.BadRequest(ctx, err)
	//    return
	//}
	//
	//blockHeightParam := ctx.UserValue("height")
	//blockHeight, err := strconv.ParseInt(blockHeightParam.(string), 10, 64)
	//if err != nil {
	//    httpapi.BadRequest(ctx, errors.New("invalid block height"))
	//    return
	//}
	//
	//blocks, paginationResult, err := handler.transactionsView.List(view.TransactionsListFilter{
	//    MaybeBlockHeight: &blockHeight,
	//}, pagination)
	//if err != nil {
	//    handler.logger.Errorf("error listing transactions: %v", err)
	//    httpapi.InternalServerError(ctx)
	//    return
	//}
	//
	//httpapi.SuccessWithPagination(ctx, blocks, paginationResult)
}
