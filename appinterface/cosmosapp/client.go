package cosmosapp

import (
	"errors"
	"math/big"

	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type Client interface {
	Account(accountAddress string) (*Account, error)
	Balances(accountAddress string) (coin.Coins, error)
	BalanceByDenom(accountAddresss string, denom string) (coin.Coin, error)
	BondedBalance(accountAddress string) (coin.Coins, error)
	RedelegatingBalance(accountAddress string) (coin.Coins, error)
	UnbondingBalance(accountAddress string) (coin.Coins, error)

	SupplyByDenom(denom string) (coin.Coin, error)

	TotalRewards(accountAddress string) (coin.DecCoins, error)
	Commission(validatorAddress string) (coin.DecCoins, error)
	CommunityPool() (coin.DecCoins, error)

	Validator(validatorAddress string) (*Validator, error)
	Delegation(delegator string, validator string) (*DelegationResponse, error)
	TotalBondedBalance() (coin.Coin, error)
	CommunityTax() (*big.Float, error)

	AnnualProvisions() (coin.DecCoin, error)

	Proposals() ([]Proposal, error)
	ProposalById(id string) (Proposal, error)
	ProposalTally(id string) (Tally, error)

	Tx(txHash string) (*model.Tx, error)
}

var ErrAccountNotFound = errors.New("account not found")
var ErrAccountNoDelegation = errors.New("account has no delegation")
var ErrProposalNotFound = errors.New("proposal not found")
