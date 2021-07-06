package ibc

type MsgConnectionOpenInitParams struct {
	RawMsgConnectionOpenInit

	ConnectionID string `json:"connectionId"`
}

type RawMsgConnectionOpenInit struct {
	ClientID     string                        `mapstructure:"client_id" json:"clientId"`
	Counterparty ConnectionCounterparty        `mapstructure:"counterparty" json:"counterparty"`
	Version      ConnectionCounterpartyVersion `mapstructure:"version" json:"version"`
	DelayPeriod  uint64                        `mapstructure:"delay_period" json:"delayPeriod,string"`
	Signer       string                        `mapstructure:"signer" json:"signer"`
}
