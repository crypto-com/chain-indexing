package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

type CreateMsgIBCTimeout struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgTimeoutParams
}

func NewCreateMsgIBCTimeout(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgTimeoutParams,
) *CreateMsgIBCTimeout {
	return &CreateMsgIBCTimeout{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgIBCTimeout) Name() string {
	return "CreateMsgIBCTimeout"
}

func (*CreateMsgIBCTimeout) Version() int {
	return 1
}

func (cmd *CreateMsgIBCTimeout) Exec() (entity_event.Event, error) {
	event := event.NewMsgIBCTimeout(cmd.msgCommonParams, cmd.params)
	return event, nil
}
