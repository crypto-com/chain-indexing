package model

type MsgAddTierParams struct {
	RawMsgAddTier
}

type RawMsgAddTier struct {
	Type      string `mapstructure:"@type" json:"@type"`
	Authority string `mapstructure:"authority" json:"authority"`
	Tier      string `mapstructure:"tier" json:"tier"`
}
