package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

type CreateMsgIBCChannelOpenTry struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgChannelOpenTryParams
}

func NewCreateMsgIBCChannelOpenTry(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgChannelOpenTryParams,
) *CreateMsgIBCChannelOpenTry {
	return &CreateMsgIBCChannelOpenTry{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgIBCChannelOpenTry) Name() string {
	return "CreateMsgIBCChannelOpenTry"
}

func (*CreateMsgIBCChannelOpenTry) Version() int {
	return 1
}

func (cmd *CreateMsgIBCChannelOpenTry) Exec() (entity_event.Event, error) {
	event := event.NewMsgIBCChannelOpenTry(cmd.msgCommonParams, cmd.params)
	return event, nil
}
