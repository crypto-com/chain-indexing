package ibc

type MsgConnectionOpenInitParams struct {
	RawMsgConnectionOpenInit

	ConnectionID string `json:"connectionId"`
}

type RawMsgConnectionOpenInit struct {
	ClientID     string                            `mapstructure:"client_id" json:"clientId"`
	Counterparty MsgConnectionOpenInitCounterparty `mapstructure:"counterparty" json:"counterparty"`
	Version      MsgConnectionOpenInitVersion      `mapstructure:"version" json:"version"`
	DelayPeriod  uint64                            `mapstructure:"delay_period" json:"delayPeriod,string"`
	Signer       string                            `mapstructure:"signer" json:"signer"`
}

type MsgConnectionOpenInitCounterparty struct {
	ClientID     string                      `mapstructure:"client_id" json:"clientId"`
	ConnectionID string                      `mapstructure:"connection_id" json:"connectionId"`
	Prefix       MsgConnectionOpenInitPrefix `mapstructure:"prefix" json:"prefix"`
}

type MsgConnectionOpenInitPrefix struct {
	KeyPrefix []byte `mapstructure:"key_prefix" json:"keyPrefix"`
}

type MsgConnectionOpenInitVersion struct {
	Identifier string   `json:"identifier"`
	Features   []string `json:"features"`
}
