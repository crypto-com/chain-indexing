package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type IBCChannelMessagesTotal struct {
	*view.Total
}

func NewIBCChannelMessagesTotal(rdbHandle *rdb.Handle) *IBCChannelMessagesTotal {
	return &IBCChannelMessagesTotal{
		view.NewTotal(rdbHandle, "view_ibc_channel_messages_total"),
	}
}
