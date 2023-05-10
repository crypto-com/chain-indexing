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
	ConnectionID string `mapstructure:"connectionId" json:"connectionId"`
	Version      string `mapstructure:"version" json:"version"`
}
