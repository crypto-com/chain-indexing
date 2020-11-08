package notification

import (
	"github.com/crypto-com/chainindex/appinterface/feed"
	"github.com/crypto-com/chainindex/usecase/model"
)

type Notification struct {
	Name    string
	Payload interface{}
}

// TODO: add block results notification
// Defines the notification name and payload here
const NOTIF_TENDERMINT_BLOCK = "TENDERMINT_BLOCK"
const NOTIF_CURRBLOCK_DONE = "CURRBLOCK_DONE"

type BlockNotifPayload struct {
	*model.Block
	*model.RawBlock
}
func NewBlockNotif(block *model.Block, rawBlock *model.RawBlock) *feed.Notification {
	// Prepare notification payload and dispatch
	data := BlockNotifPayload{
		block,
		rawBlock,
	}
	return NewNotif(NOTIF_TENDERMINT_BLOCK, data)
}

func NewCurrBlockDoneNotif() *feed.Notification {
	return NewNotif(NOTIF_CURRBLOCK_DONE, struct {}{})
}

func NewNotif(name string, data interface{}) *feed.Notification {
	return &feed.Notification{
		Name: name,
		Payload: data,
	}
}
