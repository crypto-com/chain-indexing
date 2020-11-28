package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgMultiSend struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgMultiSendParams
}

func NewCreateMsgMultiSend(msgCommonParams event.MsgCommonParams, params model.MsgMultiSendParams) *CreateMsgMultiSend {
	return &CreateMsgMultiSend{
		msgCommonParams,
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
	event := event.NewMsgMultiSend(cmd.msgCommonParams, cmd.params)
	return event, nil
}
