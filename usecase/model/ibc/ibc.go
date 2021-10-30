package ibc

import (
	"github.com/crypto-com/chain-indexing/external/json"
)

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

type Packet struct {
	Sequence           string `mapstructure:"sequence" json:"sequence"`
	SourcePort         string `mapstructure:"source_port" json:"sourcePort"`
	SourceChannel      string `mapstructure:"source_channel" json:"sourceChannel"`
	DestinationPort    string `mapstructure:"destination_port" json:"destinationPort"`
	DestinationChannel string `mapstructure:"destination_channel" json:"destinationChannel"`
	Data               string `mapstructure:"data" json:"data"`
	TimeoutHeight      Height `mapstructure:"timeout_height" json:"timeoutHeight"`
	TimeoutTimestamp   string `mapstructure:"timeout_timestamp" json:"timeoutTimestamp"`
}

type FungibleTokenPacketData struct {
	Sender   string              `mapstructure:"sender" json:"sender"`
	Receiver string              `mapstructure:"receiver" json:"receiver"`
	Denom    string              `mapstructure:"denom" json:"denom"`
	Amount   *json.NumericString `mapstructure:"amount" json:"amount"`
}
