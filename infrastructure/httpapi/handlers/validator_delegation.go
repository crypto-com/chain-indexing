package handlers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/valyala/fasthttp"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	validator_delegation_view "github.com/crypto-com/chain-indexing/projection/validator_delegation/view"
)

type ValidatorDelegation struct {
	logger applogger.Logger

	accountAddressPrefix   string
	validatorAddressPrefix string
	consNodeAddressPrefix  string

	validatorsView           validator_delegation_view.Validators
	delegationsView          validator_delegation_view.Delegations
	unbondingDelegationsView validator_delegation_view.UnbondingDelegations
	redelegationsView        validator_delegation_view.Redelegations
}

func NewValidatorDelegation(
	logger applogger.Logger,
	accountAddressPrefix string,
	validatorAddressPrefix string,
	consNodeAddressPrefix string,
	rdbHandle *rdb.Handle,
) *ValidatorDelegation {
	return &ValidatorDelegation{
		logger.WithFields(applogger.LogFields{
			"module": "ValidatorDelegationHandler",
		}),

		accountAddressPrefix,
		validatorAddressPrefix,
		consNodeAddressPrefix,

		validator_delegation_view.NewValidatorsView(rdbHandle),
		validator_delegation_view.NewDelegationsView(rdbHandle),
		validator_delegation_view.NewUnbondingDelegationsView(rdbHandle),
		validator_delegation_view.NewRedelegationsView(rdbHandle),
	}
}

func (handler *ValidatorDelegation) FindValidatorByAddress(ctx *fasthttp.RequestCtx) {

	height, err := parseInputHeight(ctx)
	if err != nil {
		httpapi.BadRequest(ctx, fmt.Errorf("error parsing input height: %v", err))
		return
	}

	validatorAddress, ok := ctx.UserValue("address").(string)
	if !ok {
		httpapi.BadRequest(ctx, errors.New("error parsing input address"))
		return
	}

	var validator validator_delegation_view.ValidatorRow
	var found bool

	if strings.HasPrefix(validatorAddress, handler.consNodeAddressPrefix) {

		validator, found, err = handler.validatorsView.FindByConsensusNodeAddr(validatorAddress, height)
		if err != nil {
			handler.logger.Errorf("error finding validator by ConsensusNodeAddress: %v", err)
			httpapi.InternalServerError(ctx)
			return
		}
		if !found {
			httpapi.NotFound(ctx)
			return
		}

	} else if strings.HasPrefix(validatorAddress, handler.validatorAddressPrefix) {

		validator, found, err = handler.validatorsView.FindByOperatorAddr(validatorAddress, height)
		if err != nil {
			handler.logger.Errorf("error finding validaotr by OperatorAddress: %v", err)
			httpapi.InternalServerError(ctx)
			return
		}
		if !found {
			httpapi.NotFound(ctx)
			return
		}

	} else {
		httpapi.BadRequest(ctx, errors.New("invalid address"))
		return
	}

	httpapi.Success(ctx, validator)
}

func (handler *ValidatorDelegation) ListValidator(ctx *fasthttp.RequestCtx) {

	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	height, err := parseInputHeight(ctx)
	if err != nil {
		httpapi.BadRequest(ctx, fmt.Errorf("error parsing input height: %v", err))
		return
	}

	validators, paginationResult, err := handler.validatorsView.List(height, pagination)
	if err != nil {
		handler.logger.Errorf("error listing validators: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, validators, paginationResult)
}

func (handler *ValidatorDelegation) ListDelegationByValidator(ctx *fasthttp.RequestCtx) {

	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	height, err := parseInputHeight(ctx)
	if err != nil {
		httpapi.BadRequest(ctx, fmt.Errorf("error parsing input height: %v", err))
		return
	}

	validatorAddress, ok := ctx.UserValue("address").(string)
	if !ok {
		httpapi.BadRequest(ctx, errors.New("error parsing input address"))
		return
	}

	delegations, paginationResult, err := handler.delegationsView.ListByValidatorAddr(validatorAddress, height, pagination)
	if err != nil {
		handler.logger.Errorf("error listing delegations by ValidatorAddress: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, delegations, paginationResult)
}

func (handler *ValidatorDelegation) ListDelegationByDelegator(ctx *fasthttp.RequestCtx) {

	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	height, err := parseInputHeight(ctx)
	if err != nil {
		httpapi.BadRequest(ctx, fmt.Errorf("error parsing input height: %v", err))
		return
	}

	delegatorAddress, ok := ctx.UserValue("address").(string)
	if !ok {
		httpapi.BadRequest(ctx, errors.New("error parsing input address"))
		return
	}

	delegations, paginationResult, err := handler.delegationsView.ListByDelegatorAddr(delegatorAddress, height, pagination)
	if err != nil {
		handler.logger.Errorf("error listing delegations by DelegatorAddress: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, delegations, paginationResult)
}

func (handler *ValidatorDelegation) ListUnbondingDelegationByValidator(ctx *fasthttp.RequestCtx) {

	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	height, err := parseInputHeight(ctx)
	if err != nil {
		httpapi.BadRequest(ctx, fmt.Errorf("error parsing input height: %v", err))
		return
	}

	validatorAddress, ok := ctx.UserValue("address").(string)
	if !ok {
		httpapi.BadRequest(ctx, errors.New("error parsing input address"))
		return
	}

	delegations, paginationResult, err := handler.unbondingDelegationsView.ListByValidatorWithPagination(validatorAddress, height, pagination)
	if err != nil {
		handler.logger.Errorf("error listing unbonding delegations by ValidatorAddress: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, delegations, paginationResult)
}

func (handler *ValidatorDelegation) ListRedelegationBySrcValidator(ctx *fasthttp.RequestCtx) {

	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	height, err := parseInputHeight(ctx)
	if err != nil {
		httpapi.BadRequest(ctx, fmt.Errorf("error parsing input height: %v", err))
		return
	}

	srcValidatorAddress, ok := ctx.UserValue("srcValAddress").(string)
	if !ok {
		httpapi.BadRequest(ctx, errors.New("error parsing input address"))
		return
	}

	delegations, paginationResult, err := handler.redelegationsView.ListBySrcValidatorWithPagination(srcValidatorAddress, height, pagination)
	if err != nil {
		handler.logger.Errorf("error listing redelegations by SrcValidatorAddress: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, delegations, paginationResult)
}

func parseInputHeight(ctx *fasthttp.RequestCtx) (int64, error) {

	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("height") {

		heightInString := string(queryArgs.Peek("height"))
		return strconv.ParseInt(heightInString, 10, 64)

	} else {
		// TODO: If no input height, then use latest height as height
		// return validator_delegation_view.LATEST_HEIGHT, nil
		return 0, errors.New("No input height")
	}

}
