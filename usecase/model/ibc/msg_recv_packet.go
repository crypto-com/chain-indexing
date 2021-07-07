package ibc

type MsgRecvPacketParams struct {
	RawMsgRecvPacket

	Application                  string                                `json:"application"`
	MessageType                  string                                `json:"messageType"`
	MaybeFungibleTokenPacketData *MsgRecvPacketFungibleTokenPacketData `json:"maybeMsgTransfer"`

	PacketSequence  uint64                 `json:"packetSequence,string"`
	ChannelOrdering string                 `json:"channelOrdering"`
	ConnectionID    string                 `json:"connectionId"`
	PacketAck       MsgRecvPacketPacketAck `json:"packetAck"`
}

type RawMsgRecvPacket struct {
	Packet          RawMsgRecvPacketPacket `mapstructure:"packet" json:"packet"`
	ProofCommitment string                 `mapstructure:"proof_commitment" json:"proofCommitment"`
	ProofHeight     Height                 `mapstructure:"proof_height" json:"proofHeight"`
	Signer          string                 `mapstructure:"signer" json:"signer"`
}

type RawMsgRecvPacketPacket struct {
	Sequence           string `mapstructure:"sequence" json:"sequence"`
	SourcePort         string `mapstructure:"source_port" json:"sourcePort"`
	SourceChannel      string `mapstructure:"source_channel" json:"sourceChannel"`
	DestinationPort    string `mapstructure:"destination_port" json:"destinationPort"`
	DestinationChannel string `mapstructure:"destination_channel" json:"destinationChannel"`
	Data               string `mapstructure:"data" json:"data"`
	TimeoutHeight      Height `mapstructure:"timeout_height" json:"timeoutHeight"`
	TimeoutTimestamp   string `mapstructure:"timeout_timestamp" json:"timeoutTimestamp"`
}

type MsgRecvPacketFungibleTokenPacketData struct {
	RawMsgRecvPacketFungibleTokenPacketData

	Success                bool   `json:"success"`
	DenominationTraceHash  string `json:"denominationTraceHash"`
	DenominationTraceDenom string `json:"denominationTraceDenom"`
}

type RawMsgRecvPacketFungibleTokenPacketData struct {
	Sender   string `mapstructure:"sender" json:"sender"`
	Receiver string `mapstructure:"receiver" json:"receiver"`
	Denom    string `mapstructure:"denom" json:"denom"`
	Amount   string `mapstructure:"amount" json:"amount"`
}

type MsgRecvPacketPacketAck struct {
	Result []byte `mspstructure:"result" json:"result"`
}
