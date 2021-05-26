package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type MessagesTotal struct {
	*view.Total
}

func NewMessagesTotal(rdbHandle *rdb.Handle) *MessagesTotal {
	return &MessagesTotal{
		view.NewTotal(rdbHandle, "view_nft_messages_total"),
	}
}
