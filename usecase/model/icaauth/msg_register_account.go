package icaauth

type MsgRegisterAccountParams struct {
	RawMsgRegisterAccount

	PortID                string `json:"portId"`
	ChannelID             string `json:"channelId"`
	CounterpartyPortID    string `json:"counterpartyPortId"`
	CounterpartyChannelID string `json:"counterpartyChannelId"`
}

type RawMsgRegisterAccount struct {
	Owner        string `mapstructure:"owner" json:"owner"`
	ConnectionID string `mapstructure:"connection_id" json:"connection_id"`
	Version      string `mapstructure:"version" json:"version"`
}
