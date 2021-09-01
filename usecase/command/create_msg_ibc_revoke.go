package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

type CreateMsgRevoke struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgRevokeParams
}

func NewCreateMsgRevoke(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgRevokeParams,
) *CreateMsgRevoke {
	return &CreateMsgRevoke{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgRevoke) Name() string {
	return "CreateMsgRevoke"
}

func (*CreateMsgRevoke) Version() int {
	return 1
}

func (cmd *CreateMsgRevoke) Exec() (entity_event.Event, error) {
	event := event.NewMsgIBCRevoke(cmd.msgCommonParams, cmd.params)
	return event, nil
}
