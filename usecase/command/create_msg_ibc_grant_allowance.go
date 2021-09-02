package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

type CreateMsgGrantAllowance struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgGrantAllowanceParams
}

func NewCreateMsgGrantAllowance(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgGrantAllowanceParams,
) *CreateMsgGrantAllowance {
	return &CreateMsgGrantAllowance{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgGrantAllowance) Name() string {
	return "CreateMsgGrantAllowance"
}

func (*CreateMsgGrantAllowance) Version() int {
	return 1
}

func (cmd *CreateMsgGrantAllowance) Exec() (entity_event.Event, error) {
	event := event.NewMsgIBCGrantAllowance(cmd.msgCommonParams, cmd.params)
	return event, nil
}
