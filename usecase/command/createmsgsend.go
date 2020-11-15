package command

import (
	entity_event "github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/usecase/event"
)

type CreateMsgSend struct {
	blockHeight int64
	txHash      string
	msgIndex    int
	params      event.MsgSendCreatedParams
}

func NewCreateMsgSend(blockHeight int64, txHash string, msgIndex int, params event.MsgSendCreatedParams) *CreateMsgSend {
	return &CreateMsgSend{
		blockHeight,
		txHash,
		msgIndex,
		params,
	}
}

func (_ *CreateMsgSend) Name() string {
	return "CreateMsgSend"
}

func (_ *CreateMsgSend) Version() int {
	return 1
}

func (cmd *CreateMsgSend) Exec() (entity_event.Event, error) {
	event := event.NewMsgSendCreated(cmd.blockHeight, cmd.txHash, cmd.msgIndex, cmd.params)
	return event, nil
}
