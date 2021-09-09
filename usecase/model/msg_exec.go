package model

type MsgExecParams struct {
	RawMsgExec
}

type RawMsgExec struct {
	Type    string       `mapstructure:"@type" json:"@type"`
	Grantee string       `mapstructure:"grantee" json:"grantee"`
	Msgs    []MsgExecMsg `mapstructure:"msgs" json:"msgs"`
}

type MsgExecMsg struct {
	Type        string          `mapstructure:"@type" json:"@type"`
}
