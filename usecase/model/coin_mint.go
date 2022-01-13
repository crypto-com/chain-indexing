package model

import "github.com/crypto-com/chain-indexing/usecase/coin"

type CoinMintParams struct {
	Address string
	Amount  coin.Coins
}
