package handlers

import (
	"github.com/crypto-com/chain-indexing/appinterface/cosmosapp"

	"github.com/crypto-com/chain-indexing/usecase/coin"

	"github.com/valyala/fasthttp"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	account_view "github.com/crypto-com/chain-indexing/projection/account/view"
)

type Accounts struct {
	logger applogger.Logger

	accountsView *account_view.Accounts
	cosmosClient cosmosapp.Client
}

func NewAccounts(logger applogger.Logger, rdbHandle *rdb.Handle, cosmosClient cosmosapp.Client) *Accounts {
	return &Accounts{
		logger.WithFields(applogger.LogFields{
			"module": "AccountsHandler",
		}),

		account_view.NewAccounts(rdbHandle),
		cosmosClient,
	}
}

func (handler *Accounts) FindBy(ctx *fasthttp.RequestCtx) {
	accountParam, _ := ctx.UserValue("account").(string)

	var info AccountInfo
	account, err := handler.cosmosClient.Account(accountParam)
	if err != nil {
		httpapi.NotFound(ctx)
		return
	}

	info.Type = account.Type
	info.Address = account.Address
	if info.Type == cosmosapp.ACCOUNT_MODULE {
		info.Name = account.MaybeModuleAccount.Name
	}

	bondedBalance, err := handler.cosmosClient.BondedBalance(accountParam)
	if err == nil {
		info.BondedBalance = bondedBalance
	}
	info.RedelegatingBalance, err = handler.cosmosClient.RedelegatingBalance(accountParam)
	if err != nil {
		handler.logger.Errorf("error fetching account redelegating balance: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}
	info.UnbondingBalance, err = handler.cosmosClient.UnbondingBalance(accountParam)
	if err != nil {
		handler.logger.Errorf("error fetching account unbonding balance: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}
	info.TotalRewards, err = handler.cosmosClient.TotalRewards(accountParam)
	if err != nil {
		handler.logger.Errorf("error fetching account total rewards: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.Success(ctx, info)
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
		if string(queryArgs.Peek("order")) == "address.desc" {
			addressOrder = view.ORDER_DESC
		}
	}

	accounts, paginationResult, err := handler.accountsView.List(account_view.AccountsListOrder{
		AccountAddress: addressOrder,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing addresses: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, accounts, paginationResult)
}

type AccountInfo struct {
	Type                string        `json:"type"`
	Name                string        `json:"name"`
	Address             string        `json:"address"`
	BondedBalance       coin.Coins    `json:"bondedBalance"`
	RedelegatingBalance coin.Coins    `json:"redelegatingBalance"`
	UnbondingBalance    coin.Coins    `json:"unbondingBalance"`
	TotalRewards        coin.DecCoins `json:"totalReward"`
}
