package model

import "github.com/crypto-com/chainindex/usecase/coin"

type CreateTransactionParams struct {
	TxHash    string
	Code      int
	Log       string
	MsgCount  int
	Fee       coin.Coin
	GasWanted string
	GasUsed   string
}
