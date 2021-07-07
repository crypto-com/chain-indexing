package ibc

type MsgChannelOpenTryParams struct {
	RawMsgChannelOpenTry

	ChannelID    string `json:"channelId"`
	ConnectionID string `json:"connectionId"`
}

type RawMsgChannelOpenTry struct {
	PortID              string  `mapstructure:"port_id" json:"portId"`
	PreviousChannelID   string  `mapstructure:"previous_channel_id" json:"previousChannelId"`
	Channel             Channel `mapstructure:"channel" json:"channel"`
	CounterpartyVersion string  `mapstructure:"counterparty_version" json:"counterpartyVersion"`
	ProofInit           []byte  `mapstructure:"proof_init" json:"proofInit"`
	ProofHeight         Height  `mapstructure:"proof_height" json:"proofHeight"`
	Signer              string  `mapstructure:"signer" json:"signer"`
}
