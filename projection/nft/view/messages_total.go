package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type MessagesTotal interface {
	IncrementAll(identities []string, total int64) error
	SumBy(identities []string) (int64, error)
}

type MessagesTotalView struct {
	*view.Total
}

func NewMessagesTotalView(rdbHandle *rdb.Handle) MessagesTotal {
	return &MessagesTotalView{
		view.NewTotal(rdbHandle, "view_nft_messages_total"),
	}
}
