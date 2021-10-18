package ibc

type MsgTransferParams struct {
	RawMsgTransfer

	PacketSequence     uint64 `json:"packetSequence,string"`
	DestinationPort    string `json:"destinationPort"`
	DestinationChannel string `json:"destinationChannel"`
	ChannelOrdering    string `json:"channelOrdering"`
	ConnectionID       string `json:"connectionId"`
	// `PacketData.Denom` is always in the format of <port>/<channel>/<basedenom> or <basedenom>
	// While `RawMsgTransfer.Token.Denom` is possible in the format of ibc/sha256(xxx)
	PacketData FungibleTokenPacketData `json:"packetData"`
}

type RawMsgTransfer struct {
	SourcePort       string           `mapstructure:"source_port" json:"sourcePort"`
	SourceChannel    string           `mapstructure:"source_channel" json:"sourceChannel"`
	Token            MsgTransferToken `mapstructure:"token" json:"token"`
	Sender           string           `mapstructure:"sender" json:"sender"`
	Receiver         string           `mapstructure:"receiver" json:"receiver"`
	TimeoutHeight    Height           `mapstructure:"timeout_height" json:"timeoutHeight"`
	TimeoutTimestamp string           `mapstructure:"timeout_timestamp" json:"timeoutTimestamp"`
}

type MsgTransferToken struct {
	Denom  string `mapstructure:"denom" json:"denom"`
	Amount uint64 `mapstructure:"amount" json:"amount,string"`
}
