package handlers

import (
	"errors"

	"github.com/crypto-com/chain-indexing/appinterface/cosmosapp"
	"github.com/crypto-com/chain-indexing/internal/primptr"
	"github.com/crypto-com/chain-indexing/internal/tmcosmosutils"
	validator_view "github.com/crypto-com/chain-indexing/projection/validator/view"

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

	accountsView   *account_view.Accounts
	validatorsView *validator_view.Validators
	cosmosClient   cosmosapp.Client

	validatorAddressPrefix string
}

func NewAccounts(
	logger applogger.Logger,
	rdbHandle *rdb.Handle,
	cosmosClient cosmosapp.Client,
	validatorAddressPrefix string,
) *Accounts {
	return &Accounts{
		logger.WithFields(applogger.LogFields{
			"module": "AccountsHandler",
		}),

		account_view.NewAccounts(rdbHandle),
		validator_view.NewValidators(rdbHandle),
		cosmosClient,

		validatorAddressPrefix,
	}
}

func (handler *Accounts) FindBy(ctx *fasthttp.RequestCtx) {
	accountParam, _ := ctx.UserValue("account").(string)

	info := AccountInfo{
		Balance:             coin.NewEmptyCoins(),
		BondedBalance:       coin.NewEmptyCoins(),
		RedelegatingBalance: coin.NewEmptyCoins(),
		UnbondingBalance:    coin.NewEmptyCoins(),
		TotalRewards:        coin.NewEmptyDecCoins(),
		Commissions:         coin.NewEmptyDecCoins(),
		TotalBalance:        coin.NewEmptyDecCoins(),
	}
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

	if balance, queryErr := handler.cosmosClient.Balances(accountParam); queryErr != nil {
		handler.logger.Errorf("error fetching account balance: %v", queryErr)
		httpapi.InternalServerError(ctx)
		return
	} else {
		info.Balance = balance
	}

	if bondedBalance, queryErr := handler.cosmosClient.BondedBalance(accountParam); queryErr != nil {
		if !errors.Is(queryErr, cosmosapp.ErrAccountNotFound) {
			handler.logger.Errorf("error fetching account bonded balance: %v", queryErr)
			httpapi.InternalServerError(ctx)
			return
		}
	} else {
		info.BondedBalance = bondedBalance
	}

	if redelegatingBalance, queryErr := handler.cosmosClient.RedelegatingBalance(accountParam); queryErr != nil {
		handler.logger.Errorf("error fetching account redelegating balance: %v", queryErr)
		httpapi.InternalServerError(ctx)
		return
	} else {
		info.RedelegatingBalance = redelegatingBalance
	}

	if unbondingBalance, queryErr := handler.cosmosClient.UnbondingBalance(accountParam); queryErr != nil {
		handler.logger.Errorf("error fetching account unbonding balance: %v", queryErr)
		httpapi.InternalServerError(ctx)
		return
	} else {
		info.UnbondingBalance = unbondingBalance
	}

	if totalRewards, queryErr := handler.cosmosClient.TotalRewards(accountParam); queryErr != nil {
		handler.logger.Errorf("error fetching account total rewards: %v", queryErr)
		httpapi.InternalServerError(ctx)
		return
	} else {
		info.TotalRewards = totalRewards
	}

	validator, err := handler.validatorsView.FindBy(validator_view.ValidatorIdentity{
		MaybeOperatorAddress: primptr.String(tmcosmosutils.MustValidatorAddressFromAccountAddress(
			handler.validatorAddressPrefix, accountParam,
		)),
	})
	if err != nil {
		if !errors.Is(err, rdb.ErrNoRows) {
			handler.logger.Errorf("error fetching account's validator: %v", err)
			httpapi.InternalServerError(ctx)
			return
		}
		// account does not have validator
		info.Commissions = coin.NewEmptyDecCoins()
	} else {
		// account has validator
		commissions, commissionErr := handler.cosmosClient.Commission(validator.OperatorAddress)
		if commissionErr != nil {
			handler.logger.Errorf("error fetching account commissions: %v", commissionErr)
			httpapi.InternalServerError(ctx)
			return
		}
		info.Commissions = commissions
	}

	totalBalance := coin.NewEmptyDecCoins()
	totalBalance = totalBalance.Add(coin.NewDecCoinsFromCoins(info.Balance...)...)
	totalBalance = totalBalance.Add(coin.NewDecCoinsFromCoins(info.BondedBalance...)...)
	totalBalance = totalBalance.Add(coin.NewDecCoinsFromCoins(info.RedelegatingBalance...)...)
	totalBalance = totalBalance.Add(coin.NewDecCoinsFromCoins(info.UnbondingBalance...)...)
	totalBalance = totalBalance.Add(info.TotalRewards...)
	totalBalance = totalBalance.Add(info.Commissions...)
	info.TotalBalance = totalBalance

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
	Balance             coin.Coins    `json:"balance"`
	BondedBalance       coin.Coins    `json:"bondedBalance"`
	RedelegatingBalance coin.Coins    `json:"redelegatingBalance"`
	UnbondingBalance    coin.Coins    `json:"unbondingBalance"`
	TotalRewards        coin.DecCoins `json:"totalRewards"`
	Commissions         coin.DecCoins `json:"commissions"`
	TotalBalance        coin.DecCoins `json:"totalBalance"`
}
