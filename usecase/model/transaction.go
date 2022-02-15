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
	MaybeKeyInfo *TransactionSignerKeyInfo `json:"keyInfo"`

	Address         string `json:"address"`
	AccountSequence uint64 `json:"accountSequence"`
}

type TransactionSignerKeyInfo struct {
	Type           string   `json:"type"`
	IsMultiSig     bool     `json:"isMultiSig"`
	Pubkeys        []string `json:"pubkeys"`
	MaybeThreshold *int     `json:"threshold,omitempty"`
}
