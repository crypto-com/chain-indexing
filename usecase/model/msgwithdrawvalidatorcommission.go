package model

import "github.com/crypto-com/chainindex/usecase/coin"

type MsgWithdrawValidatorCommissionParams struct {
	ValidatorAddress string    `json:"validatorAddress"`
	RecipientAddress string    `json:"recipientAddress"`
	Amount           coin.Coin `json:"amount"`
}
