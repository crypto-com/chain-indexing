package model

import (
	"github.com/crypto-com/chain-indexing/external/json"
)

type RawCronosSendToIBCParams struct {
	PacketChannelOrdering  string
	PacketConnection       string
	PacketData             string
	PacketDataHex          string
	PacketDstChannel       string
	PacketDstPort          string
	PacketSequence         string
	PacketSrcChannel       string
	PacketSrcPort          string
	PacketTimeoutHeight    string
	PacketTimeoutTimestamp string
}

type FungibleTokenPacketData struct {
	Sender   string              `json:"sender"`
	Receiver string              `json:"receiver"`
	Denom    string              `json:"denom"`
	Amount   *json.NumericString `json:"amount"`
}

type CronosSendToIBCParams struct {
	TxHash             string                `json:"txHash"`
	EthereumTxHash     string                `json:"ethereumTxHash"`
	SourcePort         string                `json:"sourcePort"`
	SourceChannel      string                `json:"sourceChannel"`
	Token              CronosSendToIBCToken  `json:"token"`
	Sender             string                `json:"sender"`
	Receiver           string                `json:"receiver"`
	TimeoutHeight      CronosSendToIBCHeight `json:"timeoutHeight"`
	TimeoutTimestamp   string                `json:"timeoutTimestamp"`
	PacketDataHex      string                `json:"packetDataHex"`
	PacketSequence     uint64                `json:"packetSequence,string"`
	DestinationPort    string                `json:"destinationPort"`
	DestinationChannel string                `json:"destinationChannel"`
	ChannelOrdering    string                `json:"channelOrdering"`
	ConnectionID       string                `json:"connectionId"`
}

type CronosSendToIBCToken struct {
	Denom  string              `json:"denom"`
	Amount *json.NumericString `json:"amount"`
}

type CronosSendToIBCHeight struct {
	RevisionNumber uint64 `json:"revisionNumber,string"`
	RevisionHeight uint64 `json:"revisionHeight,string"`
}
