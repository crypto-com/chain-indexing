package ibc

type MsgChannelCloseConfirmParams struct {
	RawMsgChannelCloseConfirm

	CounterpartyPortID    string `json:"counterpartyPortId"`
	CounterpartyChannelID string `json:"counterpartyChannelId"`
	ConnectionID          string `json:"connectionId"`
}

type RawMsgChannelCloseConfirm struct {
	PortID      string `mapstructure:"port_id" json:"portId"`
	ChannelID   string `mapstructure:"channel_id" json:"channelId"`
	ProofInit   []byte `mapstructure:"proof_init" json:"proofInit"`
	ProofHeight Height `mapstructure:"proof_height" json:"proofHeight"`
	Signer      string `mapstructure:"signer" json:"signer"`
}
