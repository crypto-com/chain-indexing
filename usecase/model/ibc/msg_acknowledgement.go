package ibc

type MsgAcknowledgementParams struct {
	RawMsgAcknowledgement

	Application                  string                                     `json:"application"`
	MessageType                  string                                     `json:"messageType"`
	MaybeFungibleTokenPacketData *MsgAcknowledgementFungibleTokenPacketData `json:"maybeMsgTransfer"`

	PacketSequence  uint64 `json:"packetSequence,string"`
	ChannelOrdering string `json:"channelOrdering"`
	ConnectionID    string `json:"connectionId"`
}

type RawMsgAcknowledgement struct {
	Packet          Packet `mapstructure:"packet" json:"packet"`
	Acknowledgement string `mapstructure:"acknowledgement" json:"acknowledgement"`
	ProofAcked      string `mapstructure:"proof_acked" json:"proofAcked"`
	ProofHeight     Height `mapstructure:"proof_height" json:"proofHeight"`
	Signer          string `mapstructure:"signer" json:"signer"`
}

type MsgAcknowledgementFungibleTokenPacketData struct {
	FungibleTokenPacketData

	Success         bool   `json:"success"`
	Acknowledgement string `json:"acknowledgement"`
}
