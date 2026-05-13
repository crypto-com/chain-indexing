package model

type MsgTriggerExitFromTierParams struct {
	RawMsgTriggerExitFromTier

	ExitUnlockAt string `json:"exitUnlockAt"`
}

type RawMsgTriggerExitFromTier struct {
	Type       string `mapstructure:"@type" json:"@type"`
	Owner      string `mapstructure:"owner" json:"owner"`
	PositionId string `mapstructure:"position_id" json:"positionId"`
}
