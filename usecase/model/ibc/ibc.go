package ibc

type Height struct {
	RevisionNumber uint64 `mapstructure:"revision_number" json:"revisionNumber,string"`
	RevisionHeight uint64 `mapstructure:"revision_height" json:"revisionHeight,string"`
}

type ConnectionCounterparty struct {
	ClientID     string                       `mapstructure:"client_id" json:"clientId"`
	ConnectionID string                       `mapstructure:"connection_id" json:"connectionId"`
	Prefix       ConnectionCounterpartyPrefix `mapstructure:"prefix" json:"prefix"`
}

type ConnectionCounterpartyPrefix struct {
	KeyPrefix []byte `mapstructure:"key_prefix" json:"keyPrefix"`
}

type ConnectionCounterpartyVersion struct {
	Identifier string   `mapstructure:"identifier" json:"identifier"`
	Features   []string `mapstructure:"features" json:"features"`
}

type ChannelCounterparty struct {
	PortID    string `mapstructure:"port_id" json:"portId"`
	ChannelID string `mapstructure:"channel_id" json:"channelId"`
}

type Channel struct {
	State          string              `mapstructure:"state" json:"state"`
	Ordering       string              `mapstructure:"ordering" json:"ordering"`
	Counterparty   ChannelCounterparty `mapstructure:"counterparty" json:"counterparty"`
	ConnectionHops []string            `mapstructure:"connection_hops" json:"connectionHops"`
	Version        string              `mapstructure:"version" json:"version"`
}
