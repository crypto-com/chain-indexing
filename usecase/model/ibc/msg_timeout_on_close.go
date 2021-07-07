package ibc

type MsgTimeoutOnCloseParams struct {
	RawMsgTimeoutOnClose

	Application      string                 `json:"application"`
	MessageType      string                 `json:"messageType"`
	MaybeMsgTransfer *MsgTimeoutMsgTransfer `json:"maybeMsgTransfer"`

	PacketTimeoutHeight    Height `json:"packetTimeoutHeight"`
	PacketTimeoutTimestamp uint64 `json:"packetTimeoutTimestamp,string"`
	PacketSequence         uint64 `json:"packetSequence,string"`

	ChannelOrdering string `json:"channelOrdering"`
}

type RawMsgTimeoutOnClose struct {
	Packet           Packet `mapstructure:"packet" json:"packet"`
	ProofUnreceived  []byte `mapstructure:"proof_unreceived" json:"proofUnreceived"`
	ProofClose       []byte `mapstructure:"proof_close" json:"proofClose"`
	ProofHeight      Height `mapstructure:"proof_height" json:"proofHeight"`
	NextSequenceRecv uint64 `mapstructure:"next_sequence_recv" json:"nextSequenceRecv,string"`
	Signer           string `mapstructure:"signer" json:"signer"`
}
