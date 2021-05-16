package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

// nolint:gosec
const TOKENS_TOTAL_TABLE_NAME = "view_nft_tokens_total"

type TokensTotal struct {
	*view.Total
}

func NewTokensTotal(rdbHandle *rdb.Handle) *TokensTotal {
	return &TokensTotal{
		view.NewTotal(rdbHandle, TOKENS_TOTAL_TABLE_NAME),
	}
}
