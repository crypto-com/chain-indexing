package model

type MsgCommitDelegationToTierParams struct {
	RawMsgCommitDelegationToTier
}

type RawMsgCommitDelegationToTier struct {
	Type                   string `mapstructure:"@type" json:"@type"`
	DelegatorAddress       string `mapstructure:"delegator_address" json:"delegator_address"`
	ValidatorAddress       string `mapstructure:"validator_address" json:"validator_address"`
	Amount                 string `mapstructure:"amount" json:"amount"`
	Id                     uint32 `mapstructure:"id" json:"id"`
	TriggerExitImmediately bool   `mapstructure:"trigger_exit_immediately" json:"trigger_exit_immediately"`
}
