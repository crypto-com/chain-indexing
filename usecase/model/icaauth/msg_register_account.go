package icaauth

type ChainmainMsgRegisterAccountParams struct {
	ChainmainMsgRegisterAccount

	PortID                string `json:"portId"`
	ChannelID             string `json:"channelId"`
	CounterpartyPortID    string `json:"counterpartyPortId"`
	CounterpartyChannelID string `json:"counterpartyChannelId"`
}

type ChainmainMsgRegisterAccount struct {
	Owner        string `mapstructure:"owner" json:"owner"`
	ConnectionID string `mapstructure:"connectionId" json:"connectionId"`
	Version      string `mapstructure:"version" json:"version"`
}

type MsgRegisterAccountParams struct {
	MsgRegisterAccount

	PortID                string `json:"portId"`
	ChannelID             string `json:"channelId"`
	CounterpartyPortID    string `json:"counterpartyPortId"`
	CounterpartyChannelID string `json:"counterpartyChannelId"`
}

type RawChainmainMsgRegisterAccount struct {
	Owner        string `mapstructure:"owner" json:"owner"`
	ConnectionID string `mapstructure:"connectionId" json:"connectionId"`
	Version      string `mapstructure:"version" json:"version"`
}

type MsgRegisterAccount struct {
	Owner        string `mapstructure:"owner" json:"owner"`
	ConnectionID string `mapstructure:"connection_id" json:"connection_id"`
	Version      string `mapstructure:"version" json:"version"`
}

type RawMsgRegisterAccount struct {
	Owner        string `mapstructure:"owner" json:"owner"`
	ConnectionID string `mapstructure:"connection_id" json:"connection_id"`
	Version      string `mapstructure:"version" json:"version"`
}
