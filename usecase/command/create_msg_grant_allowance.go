package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgGrantAllowance struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgGrantAllowanceParams
}

func NewCreateMsgGrantAllowance(
	msgCommonParams event.MsgCommonParams,
	params model.MsgGrantAllowanceParams,
) *CreateMsgGrantAllowance {
	return &CreateMsgGrantAllowance{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgGrantAllowance) Name() string {
	return "/cosmos.feegrant.v1beta1.MsgGrantAllowance.Create"
}

func (*CreateMsgGrantAllowance) Version() int {
	return 1
}

func (cmd *CreateMsgGrantAllowance) Exec() (entity_event.Event, error) {
	event := event.NewMsgGrantAllowance(cmd.msgCommonParams, cmd.params)
	return event, nil
}
