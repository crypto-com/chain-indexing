package ibc

type MsgChannelOpenConfirmParams struct {
	RawMsgChannelOpenConfirm

	CounterpartyChannelID string `json:"counterpartyChannelId"`
	CounterpartyPortID    string `json:"counterpartyPortId"`
	ConnectionID          string `json:"connectionId"`
}

type RawMsgChannelOpenConfirm struct {
	Type        string `mapstructure:"@type" json:"-"`
	PortID      string `mapstructure:"port_id" json:"portId"`
	ChannelID   string `mapstructure:"channel_id" json:"channelId"`
	ProofACK    []byte `mapstructure:"proof_ack" json:"proofAck"`
	ProofHeight Height `mapstructure:"proof_height" json:"proofHeight"`
	Signer      string `mapstructure:"signer" json:"signer"`
}
