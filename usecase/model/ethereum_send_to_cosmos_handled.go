package model

import "github.com/crypto-com/chain-indexing/usecase/coin"

type GravityEthereumSendToCosmosHandledEventParams struct {
	Module                    string     `json:"module"`
	Sender                    string     `json:"sender"`
	Receiver                  string     `json:"receiver"`
	Amount                    coin.Coins `json:"amount"`
	BridgeChainId             uint64     `json:"bridgeChainId"`
	EthereumTokenContract     string     `json:"tokenContract"`
	Nonce                     uint64     `json:"nonce"`
	EthereumEventVoteRecordId []byte     `json:"ethereumEventVoteRecordId"`
}
