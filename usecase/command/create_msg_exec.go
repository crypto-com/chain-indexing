package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgExec struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgExecParams
}

func NewCreateMsgExec(
	msgCommonParams event.MsgCommonParams,
	params model.MsgExecParams,
) *CreateMsgExec {
	return &CreateMsgExec{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgExec) Name() string {
	return "/cosmos.authz.v1beta1.MsgExec.Create"
}

func (*CreateMsgExec) Version() int {
	return 1
}

func (cmd *CreateMsgExec) Exec() (entity_event.Event, error) {
	event := event.NewMsgExec(cmd.msgCommonParams, cmd.params)
	return event, nil
}
