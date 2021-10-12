package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type IBCChannelMessagesTotal struct {
	*view.Total
}

type IBCChannelMessagesTotalI interface {
	Increment(identity string, total int64) error
	SumBy(identities []string) (int64, error)
}

func NewIBCChannelMessagesTotal(rdbHandle *rdb.Handle) IBCChannelMessagesTotalI {
	return &IBCChannelMessagesTotal{
		view.NewTotal(rdbHandle, "view_ibc_channel_messages_total"),
	}
}
