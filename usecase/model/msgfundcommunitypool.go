package model

import "github.com/crypto-com/chainindex/usecase/coin"

type MsgFundCommunityPoolParams struct {
	Depositor string    `json:"depositor"`
	Amount    coin.Coin `json:"amount"`
}
