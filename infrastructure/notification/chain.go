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
const (
	NOTIF_NEW_BLOCK                 = "NEW_BLOCK"
	NOTIF_TENDERMINT_BLOCK_RECEIVED = "TENDERMINT_BLOCK_RECEIVED"
	NOTIF_CURRBLOCK_DONE            = "CURRBLOCK_DONE"
)

type BlockReceivedPayload struct {
	*model.Block
	*model.RawBlock
}

func NewNotifBlockReceived(block *model.Block, rawBlock *model.RawBlock) *feed.Notification {
	// Prepare notification payload and dispatch
	data := BlockReceivedPayload{
		block,
		rawBlock,
	}
	return NewNotif(NOTIF_TENDERMINT_BLOCK_RECEIVED, data)
}

type NewBlockPayload struct {
	Height int64
}

func NewNotifNewBlock(Height int64) *feed.Notification {
	return NewNotif(NOTIF_NEW_BLOCK, NewBlockPayload{Height})
}

func NewNotifCurrBlockDone() *feed.Notification {
	return NewNotif(NOTIF_CURRBLOCK_DONE, struct{}{})
}

func NewNotif(name string, data interface{}) *feed.Notification {
	return &feed.Notification{
		Name:    name,
		Payload: data,
	}
}
