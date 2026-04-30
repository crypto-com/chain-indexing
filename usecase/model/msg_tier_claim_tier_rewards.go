package model

type MsgClaimTierRewardsParams struct {
	RawMsgClaimTierRewards
}

type RawMsgClaimTierRewards struct {
	Type        string   `mapstructure:"@type" json:"@type"`
	Owner       string   `mapstructure:"owner" json:"owner"`
	PositionIds []string `mapstructure:"position_ids" json:"position_ids"`
}
