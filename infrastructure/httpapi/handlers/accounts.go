package handlers

import (
	"errors"

	"github.com/valyala/fasthttp"

	account_view "github.com/crypto-com/chain-indexing/appinterface/projection/account/view"
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
)

type Accounts struct {
	logger applogger.Logger

	accountsView *account_view.Accounts
}

func NewAccounts(logger applogger.Logger, rdbHandle *rdb.Handle) *Accounts {
	return &Accounts{
		logger.WithFields(applogger.LogFields{
			"module": "AccountsHandler",
		}),

		account_view.NewAccounts(rdbHandle),
	}
}

func (handler *Accounts) FindBy(ctx *fasthttp.RequestCtx) {
	accountparam, _ := ctx.UserValue("address").(string)
	var identity account_view.AccountIdentity
	identity.MaybeAddress = accountparam

	account, err := handler.accountsView.FindBy(&identity)
	if err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			httpapi.NotFound(ctx)
			return
		}
		handler.logger.Errorf("error finding account by address: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.Success(ctx, account)
}

func (handler *Accounts) List(ctx *fasthttp.RequestCtx) {
	var err error

	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	addressOrder := view.ORDER_ASC
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "height.desc" {
			addressOrder = view.ORDER_DESC
		}
	}

	accounts, paginationResult, err := handler.accountsView.List(account_view.AccountsListOrder{
		AccountAddress: addressOrder,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing blocks: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, accounts, paginationResult)
}
