package ibc

type MsgTimeoutParams struct {
	RawMsgTimeout

	Application      string                 `json:"application"`
	MessageType      string                 `json:"messageType"`
	MaybeMsgTransfer *MsgTimeoutMsgTransfer `json:"maybeMsgTransfer"`

	PacketTimeoutHeight    Height `json:"packetTimeoutHeight"`
	PacketTimeoutTimestamp uint64 `json:"packetTimeoutTimestamp,string"`
	PacketSequence         uint64 `json:"packetSequence,string"`

	ChannelOrdering string `json:"channelOrdering"`
}

type RawMsgTimeout struct {
	Packet           Packet `mapstructure:"packet" json:"packet"`
	ProofUnreceived  []byte `mapstructure:"proof_unreceived" json:"proofUnreceived"`
	ProofHeight      Height `mapstructure:"proof_height" json:"proofHeight"`
	NextSequenceRecv uint64 `mapstructure:"next_sequence_recv" json:"nextSequenceRecv,string"`
	Signer           string `mapstructure:"signer" json:"signer"`
}

type MsgTimeoutMsgTransfer struct {
	RefundReceiver string `json:"refundReceiver"`
	RefundDenom    string `json:"refundDenom"`
	RefundAmount   uint64 `json:"refundAmount,string"`
}
