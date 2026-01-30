package cosmosapp

import (
	"errors"
	"math/big"

	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type Client interface {
	Account(accountAddress string, cosmosAPIVersion string) (*Account, error)
	Balances(accountAddress string, cosmosAPIVersion string) (coin.Coins, error)
	BalanceByDenom(accountAddresss string, denom string, cosmosAPIVersion string) (coin.Coin, error)
	BondedBalance(accountAddress string, cosmosAPIVersion string) (coin.Coins, error)
	RedelegatingBalance(accountAddress string, cosmosAPIVersion string) (coin.Coins, error)
	UnbondingBalance(accountAddress string, cosmosAPIVersion string) (coin.Coins, error)

	SupplyByDenom(denom string, cosmosAPIVersion string) (coin.Coin, error)

	TotalRewards(accountAddress string, cosmosAPIVersion string) (coin.DecCoins, error)
	Commission(validatorAddress string, cosmosAPIVersion string) (coin.DecCoins, error)
	CommunityPool(cosmosAPIVersion string) (coin.DecCoins, error)

	Validator(validatorAddress string, cosmosAPIVersion string) (*Validator, error)
	Delegation(delegator string, validator string, cosmosAPIVersion string) (*DelegationResponse, error)
	TotalBondedBalance(cosmosAPIVersion string) (coin.Coin, error)
	CommunityTax(cosmosAPIVersion string) (*big.Float, error)

	AnnualProvisions(cosmosAPIVersion string) (coin.DecCoin, error)

	Proposals(cosmosAPIVersion string) ([]Proposal, error)
	ProposalById(id string, cosmosAPIVersion string) (Proposal, error)
	ProposalTally(id string, cosmosAPIVersion string) (Tally, error)

	Tx(txHash string, cosmosAPIVersion string) (*model.Tx, error)
}

var ErrAccountNotFound = errors.New("account not found")
var ErrAccountNoDelegation = errors.New("account has no delegation")
var ErrProposalNotFound = errors.New("proposal not found")
