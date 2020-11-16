package model

import (
	"github.com/crypto-com/chainindex/usecase/coin"
)

type MsgMultiSendParams struct {
	Inputs  []MsgMultiSendInput  `json:"inputs"`
	Outputs []MsgMultiSendOutput `json:"outputs"`
}

type MsgMultiSendInput struct {
	Address string    `json:"address"`
	Amount  coin.Coin `json:"amount"`
}

type MsgMultiSendOutput struct {
	Address string    `json:"address"`
	Amount  coin.Coin `json:"amount"`
}
