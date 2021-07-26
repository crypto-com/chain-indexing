package model

import "github.com/crypto-com/chain-indexing/usecase/coin"

type CreateTransactionParams struct {
	TxHash        string
	Index         int
	Code          int
	Log           string
	MsgCount      int
	Signers       []TransactionSigner
	Fee           coin.Coins
	FeePayer      string
	FeeGranter    string
	GasWanted     int
	GasUsed       int
	Memo          string
	TimeoutHeight int64
}

type TransactionSigner struct {
	TransactionSignerInfo

	Address string `json:"address"`
}

type TransactionSignerInfo struct {
	Type            string   `json:"type"`
	IsMultiSig      bool     `json:"isMultiSig"`
	Pubkeys         []string `json:"pubkeys"`
	MaybeThreshold  *int     `json:"threshold,omitempty"`
	AccountSequence uint64   `json:"accountSequence"`
}
