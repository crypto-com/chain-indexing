package model

import (
	"github.com/crypto-com/chain-indexing/internal/utctime"
	"github.com/crypto-com/chain-indexing/usecase/coin"
)

type MsgUndelegateParams struct {
	DelegatorAddress      string           `json:"delegatorAddress"`
	ValidatorAddress      string           `json:"validatorAddress"`
	Amount                coin.Coin        `json:"amount"`
	MaybeUnbondCompleteAt *utctime.UTCTime `json:"unbondCompleteAt"`
	AutoClaimedReward     coin.Coin        `json:"autoClaimedAmount"`
}
