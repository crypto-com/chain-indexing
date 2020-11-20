package model

import (
	"github.com/crypto-com/chainindex/usecase/coin"
)

type MsgCreateValidatorParams struct {
	Description      Description     `json:"description"`
	CommissionRates  CommissionRates `json:"commission"`
	DelegatorAddress string          `json:"delegatorAddress"`
	ValidatorAddress string          `json:"validatorAddress"`
	Pubkey           string          `json:"pubkey"`
	Value            coin.Coin       `json:"value"`
}
