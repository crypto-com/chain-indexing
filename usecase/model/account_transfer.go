package model

import "github.com/crypto-com/chainindex/usecase/coin"

type AccountTransferParams struct {
	Recipient string
	Sender    string
	Amount    coin.Coin
}
