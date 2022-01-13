package model

import "github.com/crypto-com/chain-indexing/usecase/coin"

type CoinReceivedParams struct {
	Address string
	Amount  coin.Coins
}
