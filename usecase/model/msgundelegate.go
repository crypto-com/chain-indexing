package model

import (
	"github.com/crypto-com/chainindex/usecase/coin"
)

type MsgUndelegateParams struct {
	DelegatorAddress string    `json:"delegatorAddress"`
	ValidatorAddress string    `json:"validatorAddress"`
	Amount           coin.Coin `json:"amount"`
}
