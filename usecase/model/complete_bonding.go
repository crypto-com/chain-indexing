package model

import "github.com/crypto-com/chain-indexing/usecase/coin"

type CompleteBondingParams struct {
	Delegator string
	Validator string
	Amount    coin.Coins
}
