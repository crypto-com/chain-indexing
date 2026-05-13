package model

type MsgWithdrawFromTierParams struct {
	RawMsgWithdrawFromTier

	Amount string `json:"amount"`
}

type RawMsgWithdrawFromTier struct {
	Type       string `mapstructure:"@type" json:"@type"`
	Owner      string `mapstructure:"owner" json:"owner"`
	PositionId string `mapstructure:"position_id" json:"positionId"`
}
