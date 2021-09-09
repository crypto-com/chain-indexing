package model

type MsgRevokeAllowanceParams struct {
	RawMsgRevokeAllowance
}

type RawMsgRevokeAllowance struct {
	Type       string `mapstructure:"@type" json:"@type"`
	Granter    string `mapstructure:"granter" json:"granter"`
	Grantee    string `mapstructure:"grantee" json:"grantee"`
}
