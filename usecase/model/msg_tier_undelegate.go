package model

type MsgTierUndelegateParams struct {
	RawMsgTierUndelegate
}

type RawMsgTierUndelegate struct {
	Type       string `mapstructure:"@type" json:"@type"`
	Owner      string `mapstructure:"owner" json:"owner"`
	PositionId string `mapstructure:"position_id" json:"positionId"`
}
