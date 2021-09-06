package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

type CreateMsgRevokeAllowance struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgRevokeAllowanceParams
}

func NewCreateMsgRevokeAllowance(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgRevokeAllowanceParams,
) *CreateMsgRevokeAllowance {
	return &CreateMsgRevokeAllowance{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgRevokeAllowance) Name() string {
	return "CreateMsgRevokeAllowance"
}

func (*CreateMsgRevokeAllowance) Version() int {
	return 1
}

func (cmd *CreateMsgRevokeAllowance) Exec() (entity_event.Event, error) {
	event := event.NewMsgRevokeAllowance(cmd.msgCommonParams, cmd.params)
	return event, nil
}
