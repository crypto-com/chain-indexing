package ibc

type MsgTransferParams struct {
	RawMsgTransfer

	PacketSequence     uint64 `json:"packetSequence,string"`
	DestinationPort    string `json:"destinationPort"`
	DestinationChannel string `json:"destinationChannel"`
	ChannelOrdering    string `json:"channelOrdering"`
	ConnectionID       string `json:"connectionId"`
}

type RawMsgTransfer struct {
	Type             string        `mapstructure:"@type" json:"-"`
	SourcePort       string        `mapstructure:"source_port" json:"sourcePort"`
	SourceChannel    string        `mapstructure:"source_channel" json:"sourceChannel"`
	Token            TransferToken `mapstructure:"token" json:"token"`
	Sender           string        `mapstructure:"sender" json:"sender"`
	Receiver         string        `mapstructure:"receiver" json:"receiver"`
	TimeoutHeight    Height        `mapstructure:"timeout_height" json:"timeoutHeight"`
	TimeoutTimestamp string        `mapstructure:"timeout_timestamp" json:"timeoutTimestamp"`
}

type TransferToken struct {
	Denom  string `mapstructure:"denom" json:"denom"`
	Amount string `mapstructure:"amount" json:"amount"`
}
