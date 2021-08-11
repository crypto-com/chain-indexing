package model

import "github.com/crypto-com/chain-indexing/usecase/coin"

type MintParams struct {
	BondedRatio      string
	Inflation        string
	AnnualProvisions coin.DecCoin
	Amount           coin.Coins
}
