package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
)

type CreateMsgSend struct {
	msgCommonParams event.MsgCommonParams
	params          event.MsgSendCreatedParams
}

func NewCreateMsgSend(msgCommonParams event.MsgCommonParams, params event.MsgSendCreatedParams) *CreateMsgSend {
	return &CreateMsgSend{
		msgCommonParams,
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
	event := event.NewMsgSend(cmd.msgCommonParams, cmd.params)
	return event, nil
}
