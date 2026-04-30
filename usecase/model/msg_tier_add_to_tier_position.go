package model

type MsgAddToTierPositionParams struct {
	RawMsgAddToTierPosition
}

type RawMsgAddToTierPosition struct {
	Type       string `mapstructure:"@type" json:"@type"`
	Owner      string `mapstructure:"owner" json:"owner"`
	PositionId string `mapstructure:"position_id" json:"positionId"`
	Amount     string `mapstructure:"amount" json:"amount"`
}
