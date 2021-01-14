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

const ACCOUNT_MODULE = "/cosmos.auth.v1beta1.ModuleAccount"
const ACCOUNT_BASE = "/cosmos.auth.v1beta1.BaseAccount"
const ACCOUNT_VESTING_DELAYED = "/cosmos.vesting.v1beta1.DelayedVestingAccount"
const ACCOUNT_VESTING_CONTINUOUS = "/cosmos.vesting.v1beta1.ContinuousVestingAccount"
const ACCOUNT_VESTING_PERIODIC = "/cosmos.vesting.v1beta1.PeriodicVestingAccount"

type Account struct {
	Type          string  `json:"type"`
	Address       string  `json:"address"`
	MaybePubkey   *PubKey `json:"pubkey"`
	AccountNumber string  `json:"account_number"`
	Sequence      string  `json:"sequence"`

	MaybeModuleAccount            *ModuleAccount            `json:"module_account"`
	MaybeDelayedVestingAccount    *DelayedVestingAccount    `json:"delayed_vesting_account"`
	MaybeContinuousVestingAccount *ContinuousVestingAccount `json:"continuous_vesting_account"`
	MaybePeriodicVestingAccount   *PeriodicVestingAccount   `json:"periodic_vesting_account"`
}

type PubKey struct {
	Type string `json:"@type"`
	Key  string `json:"key"`
}

type ModuleAccount struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

type DelayedVestingAccount struct {
	OriginalVesting  []VestingBalance `json:"original_vesting"`
	DelegatedFree    []VestingBalance `json:"delegated_free"`
	DelegatedVesting []VestingBalance `json:"delegated_vesting"`
	EndTime          string           `json:"end_time"`
}

type ContinuousVestingAccount struct {
	OriginalVesting  []VestingBalance `json:"original_vesting"`
	DelegatedFree    []VestingBalance `json:"delegated_free"`
	DelegatedVesting []VestingBalance `json:"delegated_vesting"`
	StartTime        string           `json:"start_time"`
	EndTime          string           `json:"end_time"`
}

type PeriodicVestingAccount struct {
	OriginalVesting  []VestingBalance `json:"original_vesting"`
	DelegatedFree    []VestingBalance `json:"delegated_free"`
	DelegatedVesting []VestingBalance `json:"delegated_vesting"`
	StartTime        string           `json:"start_time"`
	EndTime          string           `json:"end_time"`
	VestingPeriods   []VestingPeriod  `json:"vesting_periods"`
}

type VestingPeriod struct {
	Amount []VestingBalance `json:"amount"`
	Length string           `json:"length"`
}

type VestingBalance struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type BankBalance struct {
	Amount string `json:"amount"`
	Denom  string `json:"denom"`
}
