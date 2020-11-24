package model

import "github.com/crypto-com/chainindex/usecase/coin"

type CreateTransactionParams struct {
	TxHash        string
	Code          int
	Log           string
	MsgCount      int
	Fee           coin.Coin
	FeePayer      string
	FeeGranter    string
	GasWanted     string
	GasUsed       string
	Memo          string
	TimeoutHeight int64
}
