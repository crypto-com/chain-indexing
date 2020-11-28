package model

import "github.com/crypto-com/chain-indexing/usecase/coin"

type MsgWithdrawDelegatorRewardParams struct {
	DelegatorAddress string    `json:"delegatorAddress"`
	ValidatorAddress string    `json:"validatorAddress"`
	RecipientAddress string    `json:"recipientAddress"`
	Amount           coin.Coin `json:"amount"`
}
