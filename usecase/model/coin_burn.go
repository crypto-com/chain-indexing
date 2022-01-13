package model

import "github.com/crypto-com/chain-indexing/usecase/coin"

type CoinBurnParams struct {
	Address string
	Amount  coin.Coins
}
