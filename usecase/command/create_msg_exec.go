package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

type CreateMsgExec struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgExecParams
}

func NewCreateMsgExec(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgExecParams,
) *CreateMsgExec {
	return &CreateMsgExec{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgExec) Name() string {
	return "CreateMsgExec"
}

func (*CreateMsgExec) Version() int {
	return 1
}

func (cmd *CreateMsgExec) Exec() (entity_event.Event, error) {
	event := event.NewMsgExec(cmd.msgCommonParams, cmd.params)
	return event, nil
}
