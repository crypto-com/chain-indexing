package model

import (
	"github.com/crypto-com/chainindex/usecase/coin"
)

type MsgCreateValidatorParams struct {
	Description       MsgValidatorDescription `json:"description"`
	Commission        MsgValidatorCommission  `json:"commission"`
	MinSelfDelegation string                  `json:"minSelfDelegation"`
	DelegatorAddress  string                  `json:"delegatorAddress"`
	ValidatorAddress  string                  `json:"validatorAddress"`
	Pubkey            string                  `json:"pubkey"`
	Amount            coin.Coin               `json:"amount"`
}
