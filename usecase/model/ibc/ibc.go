package ibc

type Height struct {
	RevisionNumber uint64 `mapstructure:"revision_number" json:"revisionNumber,string"`
	RevisionHeight uint64 `mapstructure:"revision_height" json:"revisionHeight,string"`
}

type Counterparty struct {
	ClientID     string             `mapstructure:"client_id" json:"clientId"`
	ConnectionID string             `mapstructure:"connection_id" json:"connectionId"`
	Prefix       CounterpartyPrefix `mapstructure:"prefix" json:"prefix"`
}

type CounterpartyPrefix struct {
	KeyPrefix []byte `mapstructure:"key_prefix" json:"keyPrefix"`
}

type CounterpartyVersion struct {
	Identifier string   `json:"identifier"`
	Features   []string `json:"features"`
}
