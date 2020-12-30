package model

import "github.com/crypto-com/chain-indexing/usecase/coin"

type CreateTransactionParams struct {
	TxHash        string
	Code          int
	Log           string
	MsgCount      int
	Signers       []TransactionSigner
	Fee           coin.Coin
	FeePayer      string
	FeeGranter    string
	GasWanted     int
	GasUsed       int
	Memo          string
	TimeoutHeight int64
}

type TransactionSigner struct {
	Type            string
	Pubkeys         []string
	MaybeThreshold  *int
	AccountSequence uint64
}
