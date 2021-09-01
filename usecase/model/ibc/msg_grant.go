package ibc

type MsgGrantParams struct {
	RawMsgGrant
}

type RawMsgGrant struct {
	Type    string        `mapstructure:"@type" json:"@type"`
	Granter string        `mapstructure:"granter" json:"granter"`
	Grantee string        `mapstructure:"grantee" json:"grantee"`
	Grant   MsgGrantGrant `mapstructure:"grant" json:"grant"`
}

type MsgGrantGrant struct {
	Authorization MsgGrantAuthorization `mapstructure:"authorization" json:"authorization"`
	Expiration    string                `mapstructure:"expiration" json:"expiration"`
}

type MsgGrantAuthorization struct {
	Type       string               `mapstructure:"@type" json:"@type"`
	SpendLimit []MsgGrantSpendLimit `mapstructure:"spend_limit" json:"spendLimit"`
}

type MsgGrantSpendLimit struct {
	Denom  string `mapstructure:"denom" json:"denom"`
	Amount string `mapstructure:"amount" json:"amount"`
}
