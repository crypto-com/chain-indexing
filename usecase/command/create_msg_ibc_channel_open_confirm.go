package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

type CreateMsgIBCChannelOpenConfirm struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgChannelOpenConfirmParams
}

func NewCreateMsgIBCChannelOpenConfirm(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgChannelOpenConfirmParams,
) *CreateMsgIBCChannelOpenConfirm {
	return &CreateMsgIBCChannelOpenConfirm{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgIBCChannelOpenConfirm) Name() string {
	return "CreateMsgIBCChannelOpenConfirm"
}

func (*CreateMsgIBCChannelOpenConfirm) Version() int {
	return 1
}

func (cmd *CreateMsgIBCChannelOpenConfirm) Exec() (entity_event.Event, error) {
	event := event.NewMsgIBCChannelOpenConfirm(cmd.msgCommonParams, cmd.params)
	return event, nil
}
