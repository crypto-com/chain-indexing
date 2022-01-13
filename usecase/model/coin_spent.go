package model

import "github.com/crypto-com/chain-indexing/usecase/coin"

type CoinSpentParams struct {
	Address string
	Amount  coin.Coins
}
