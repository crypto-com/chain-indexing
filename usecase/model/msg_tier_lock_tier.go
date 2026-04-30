package model

type MsgLockTierParams struct {
	RawMsgLockTier
}

type RawMsgLockTier struct {
	Type                   string `mapstructure:"@type" json:"@type"`
	Owner                  string `mapstructure:"owner" json:"owner"`
	Id                     uint32 `mapstructure:"id" json:"id"`
	Amount                 string `mapstructure:"amount" json:"amount"`
	ValidatorAddress       string `mapstructure:"validator_address" json:"validatorAddress"`
	TriggerExitImmediately bool   `mapstructure:"trigger_exit_immediately" json:"triggerExitImmediately"`
}
