package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgSetSendEnabled struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgSetSendEnabled
}

func NewCreateMsgSetSendEnabled(msgCommonParams event.MsgCommonParams, params model.MsgSetSendEnabled) *CreateMsgSetSendEnabled {
	return &CreateMsgSetSendEnabled{
		msgCommonParams,
		params,
	}
}

func (_ *CreateMsgSetSendEnabled) Name() string {
	return "/cosmos.bank.v1beta1.MsgSetSendEnabled.Create"
}

func (_ *CreateMsgSetSendEnabled) Version() int {
	return 1
}

func (cmd *CreateMsgSetSendEnabled) Exec() (entity_event.Event, error) {
	event := event.NewMsgSetSendEnabled(cmd.msgCommonParams, cmd.params)
	return event, nil
}
