package model

import "github.com/crypto-com/chainindex/usecase/coin"

type CompleteBondingParams struct {
	Delegator string
	Validator string
	Amount    coin.Coin
}
