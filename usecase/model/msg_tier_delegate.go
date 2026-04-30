package model

type MsgTierDelegateParams struct {
	RawMsgTierDelegate
}

type RawMsgTierDelegate struct {
	Type       string `mapstructure:"@type" json:"@type"`
	Owner      string `mapstructure:"owner" json:"owner"`
	PositionId string `mapstructure:"position_id" json:"position_id"`
	Validator  string `mapstructure:"validator" json:"validator"`
}
