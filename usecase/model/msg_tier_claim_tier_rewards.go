package model

type MsgClaimTierRewardsParams struct {
	RawMsgClaimTierRewards

	BaseRewards  string `json:"baseRewards"`
	BonusRewards string `json:"bonusRewards"`
}

type RawMsgClaimTierRewards struct {
	Type        string   `mapstructure:"@type" json:"@type"`
	Owner       string   `mapstructure:"owner" json:"owner"`
	PositionIds []string `mapstructure:"position_ids" json:"positionIds"`
}
