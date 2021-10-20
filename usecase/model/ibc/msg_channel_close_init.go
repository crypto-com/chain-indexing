package ibc

type MsgChannelCloseInitParams struct {
	RawMsgChannelCloseInit

	CounterpartyPortID    string `json:"counterpartyPortId"`
	CounterpartyChannelID string `json:"counterpartyChannelId"`
	ConnectionID          string `json:"connectionId"`
}

type RawMsgChannelCloseInit struct {
	PortID    string `mapstructure:"port_id" json:"portId"`
	ChannelID string `mapstructure:"channel_id" json:"channelId"`
	Signer    string `mapstructure:"signer" json:"signer"`
}
