package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

type CreateMsgConnectionOpenInit struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgConnectionOpenInitParams
}

func NewCreateMsgConnectionOpenInit(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgConnectionOpenInitParams,
) *CreateMsgConnectionOpenInit {
	return &CreateMsgConnectionOpenInit{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgConnectionOpenInit) Name() string {
	return "CreateMsgConnectionOpenInit"
}

func (*CreateMsgConnectionOpenInit) Version() int {
	return 1
}

func (cmd *CreateMsgConnectionOpenInit) Exec() (entity_event.Event, error) {
	event := event.NewMsgConnectionOpenInit(cmd.msgCommonParams, cmd.params)
	return event, nil
}
