package model

type MsgDeleteTierParams struct {
	RawMsgDeleteTier
}

type RawMsgDeleteTier struct {
	Type      string `mapstructure:"@type" json:"@type"`
	Authority string `mapstructure:"authority" json:"authority"`
	Id        uint32 `mapstructure:"id" json:"id"`
}
