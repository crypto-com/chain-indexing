package model

type MsgUpdateTierParams struct {
	RawMsgUpdateTier
}

type RawMsgUpdateTier struct {
	Type      string `mapstructure:"@type" json:"@type"`
	Authority string `mapstructure:"authority" json:"authority"`
	Tier      string `mapstructure:"tier" json:"tier"`
}
