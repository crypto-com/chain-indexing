package model

import "github.com/crypto-com/chain-indexing/usecase/coin"

type MsgFundCommunityPoolParams struct {
	Depositor string    `json:"depositor"`
	Amount    coin.Coin `json:"amount"`
}
