package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

type CreateMsgIBCConnectionOpenInit struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgConnectionOpenInitParams
}

func NewCreateMsgIBCConnectionOpenInit(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgConnectionOpenInitParams,
) *CreateMsgIBCConnectionOpenInit {
	return &CreateMsgIBCConnectionOpenInit{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgIBCConnectionOpenInit) Name() string {
	return "CreateMsgIBCConnectionOpenInit"
}

func (*CreateMsgIBCConnectionOpenInit) Version() int {
	return 1
}

func (cmd *CreateMsgIBCConnectionOpenInit) Exec() (entity_event.Event, error) {
	event := event.NewMsgIBCConnectionOpenInit(cmd.msgCommonParams, cmd.params)
	return event, nil
}
