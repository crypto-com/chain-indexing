package command

import (
	entity_event "github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/usecase/event"
	"github.com/crypto-com/chainindex/usecase/model"
)

type CreateMsgMultiSend struct {
	blockHeight int64
	txHash      string
	msgIndex    int

	params model.MsgMultiSendParams
}

func NewCreateMsgMultiSend(blockHeight int64, txHash string, msgIndex int, params model.MsgMultiSendParams) *CreateMsgMultiSend {
	return &CreateMsgMultiSend{
		blockHeight,
		txHash,
		msgIndex,

		params,
	}
}

func (_ *CreateMsgMultiSend) Name() string {
	return "CreateMsgMultiSend"
}

func (_ *CreateMsgMultiSend) Version() int {
	return 1
}

func (cmd *CreateMsgMultiSend) Exec() (entity_event.Event, error) {
	event := event.NewMsgMultiSendCreated(cmd.blockHeight, cmd.txHash, cmd.msgIndex, cmd.params)
	return event, nil
}
