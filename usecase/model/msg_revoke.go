package model

type MsgRevokeParams struct {
	RawMsgRevoke
}

type RawMsgRevoke struct {
	Type       string `mapstructure:"@type" json:"@type"`
	Granter    string `mapstructure:"granter" json:"granter"`
	Grantee    string `mapstructure:"grantee" json:"grantee"`
	MsgTypeURL string `mapstructure:"msg_type_url" json:"msgTypeUrl"`
}
