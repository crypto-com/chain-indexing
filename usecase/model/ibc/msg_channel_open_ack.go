package ibc

type MsgChannelOpenAckParams struct {
	RawMsgChannelOpenAck

	CounterpartyPortID string `json:"counterpartyPortId"`
	ConnectionID       string `json:"connectionId"`
}

type RawMsgChannelOpenAck struct {
	PortID                string `mapstructure:"port_id" json:"portId"`
	ChannelID             string `mapstructure:"channel_id" json:"channelId"`
	CounterpartyChannelID string `mapstructure:"counterparty_channel_id" json:"counterpartyChannelId"`
	CounterpartyVersion   string `mapstructure:"counterparty_version" json:"counterpartyVersion"`
	ProofTry              []byte `mapstructure:"proof_try" json:"proofTry"`
	ProofHeight           Height `mapstructure:"proof_height" json:"proofHeight"`
	Signer                string `mapstructure:"signer" json:"signer"`
}
