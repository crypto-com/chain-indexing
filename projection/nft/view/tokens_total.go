package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

// nolint:gosec
const TOKENS_TOTAL_TABLE_NAME = "view_nft_tokens_total"

type TokensTotal interface {
	IncrementAll(identities []string, total int64) error
	DecrementAll(identities []string, total int64) error
	FindBy(identity string) (int64, error)
}

type TokensTotalView struct {
	*view.Total
}

func NewTokensTotalView(rdbHandle *rdb.Handle) TokensTotal {
	return &TokensTotalView{
		view.NewTotal(rdbHandle, TOKENS_TOTAL_TABLE_NAME),
	}
}
