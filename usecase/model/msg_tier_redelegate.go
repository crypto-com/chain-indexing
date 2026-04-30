package model

type MsgTierRedelegateParams struct {
	RawMsgTierRedelegate
}

type RawMsgTierRedelegate struct {
	Type         string `mapstructure:"@type" json:"@type"`
	Owner        string `mapstructure:"owner" json:"owner"`
	PositionId   string `mapstructure:"position_id" json:"position_id"`
	DstValidator string `mapstructure:"dst_validator" json:"dst_validator"`
}
