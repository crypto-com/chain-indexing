package model

import "github.com/crypto-com/chain-indexing/usecase/coin"

type AccountTransferParams struct {
	Recipient string
	Sender    string
	Amount    coin.Coin
}
