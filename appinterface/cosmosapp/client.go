package cosmosapp

import "github.com/crypto-com/chain-indexing/usecase/coin"

type Client interface {
	Account(accountAddress string) (*Account, error)
	Balances(accountAddress string) (coin.Coins, error)
	BalanceByDenom(accountAddress string, denom string) (*coin.Coin, error)
	Validator(validatorAddress string) (*Validator, error)
	Delegation(delegator string, validator string) (*DelegationResponse, error)
}

type Validator struct {
	OperatorAddress   string               `json:"operator_address"`
	ConsensusPubkey   string               `json:"consensus_pubkey"`
	Jailed            bool                 `json:"jailed"`
	Status            string               `json:"status"`
	Tokens            string               `json:"tokens"`
	DelegatorShares   string               `json:"delegator_shares"`
	Description       ValidatorDescription `json:"description"`
	UnbondingHeight   string               `json:"unbonding_height"`
	UnbondingTime     string               `json:"unbonding_time"`
	Commission        ValidatorCommission  `json:"commission"`
	MinSelfDelegation string               `json:"min_self_delegation"`
}

type ValidatorCommission struct {
	CommissionRates ValidatorCommissionRates `json:"commission_rates"`
	UpdateTime      string                   `json:"update_time"`
}

type ValidatorCommissionRates struct {
	Rate          string `json:"rate"`
	MaxRate       string `json:"max_rate"`
	MaxChangeRate string `json:"max_change_rate"`
}

type ValidatorDescription struct {
	Moniker         string `json:"moniker"`
	Identity        string `json:"identity"`
	Website         string `json:"website"`
	SecurityContact string `json:"security_contact"`
	Details         string `json:"details"`
}

type DelegationResponse struct {
	Delegation Delegation        `json:"delegation"`
	Balance    DelegationBalance `json:"balance"`
}

type DelegationBalance struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type Delegation struct {
	DelegatorAddress string `json:"delegator_address"`
	ValidatorAddress string `json:"validator_address"`
	Shares           string `json:"shares"`
}

type Pagination struct {
	MaybeNextKey *string `json:"next_key"`
	Total        string  `json:"total"`
}

type Account struct {
	AccountType    string `json:"account_type"`
	AccountAddress string `json:"account_address"`
	Pubkey         string `json:"account_pubkey"`
	AccountNumber  string `json:"account_number"`
	SequenceNumber string `json:"sequence_number"`
}

type BankBalance struct {
	Amount string `json:"amount"`
	Denom  string `json:"denom"`
}
