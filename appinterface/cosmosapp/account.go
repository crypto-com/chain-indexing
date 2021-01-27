package cosmosapp

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
