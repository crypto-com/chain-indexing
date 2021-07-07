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
	Packet          Packet `mapstructure:"packet" json:"packet"`
	ProofCommitment string `mapstructure:"proof_commitment" json:"proofCommitment"`
	ProofHeight     Height `mapstructure:"proof_height" json:"proofHeight"`
	Signer          string `mapstructure:"signer" json:"signer"`
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
	Amount   uint64 `mapstructure:"amount" json:"amount,string"`
}

type MsgRecvPacketPacketAck struct {
	Result []byte `mspstructure:"result" json:"result"`
}
