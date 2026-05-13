package model

type MsgCommitDelegationToTierParams struct {
	RawMsgCommitDelegationToTier

	PositionId string `json:"positionId"`
}

type RawMsgCommitDelegationToTier struct {
	Type                   string `mapstructure:"@type" json:"@type"`
	DelegatorAddress       string `mapstructure:"delegator_address" json:"delegatorAddress"`
	ValidatorAddress       string `mapstructure:"validator_address" json:"validatorAddress"`
	Amount                 string `mapstructure:"amount" json:"amount"`
	Id                     uint32 `mapstructure:"id" json:"id"`
	TriggerExitImmediately bool   `mapstructure:"trigger_exit_immediately" json:"triggerExitImmediately"`
}
