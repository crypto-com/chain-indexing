package icaauth

type MsgRegisterAccountParams struct {
	MsgRegisterAccount

	PortID                string `json:"portId"`
	ChannelID             string `json:"channelId"`
	CounterpartyPortID    string `json:"counterpartyPortId"`
	CounterpartyChannelID string `json:"counterpartyChannelId"`
}

type MsgRegisterAccount struct {
	Owner        string `mapstructure:"owner" json:"owner"`
	ConnectionID string `mapstructure:"connectionId" json:"connectionId"`
	Version      string `mapstructure:"version" json:"version"`
}

type RawChainmainMsgRegisterAccount struct {
	Owner        string `mapstructure:"owner" json:"owner"`
	ConnectionID string `mapstructure:"connectionId" json:"connectionId"`
	Version      string `mapstructure:"version" json:"version"`
}

type RawMsgRegisterAccount struct {
	Owner        string `mapstructure:"owner" json:"owner"`
	ConnectionID string `mapstructure:"connection_id" json:"connection_id"`
	Version      string `mapstructure:"version" json:"version"`
}
