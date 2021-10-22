package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

type CreateMsgIBCChannelCloseConfirm struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgChannelCloseConfirmParams
}

func NewCreateMsgIBCChannelCloseConfirm(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgChannelCloseConfirmParams,
) *CreateMsgIBCChannelCloseConfirm {
	return &CreateMsgIBCChannelCloseConfirm{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgIBCChannelCloseConfirm) Name() string {
	return "CreateMsgIBCChannelCloseConfirm"
}

func (*CreateMsgIBCChannelCloseConfirm) Version() int {
	return 1
}

func (cmd *CreateMsgIBCChannelCloseConfirm) Exec() (entity_event.Event, error) {
	event := event.NewMsgIBCChannelCloseConfirm(cmd.msgCommonParams, cmd.params)
	return event, nil
}
