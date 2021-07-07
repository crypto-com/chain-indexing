package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

type CreateMsgIBCTimeoutOnClose struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgTimeoutOnCloseParams
}

func NewCreateMsgIBCTimeoutOnClose(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgTimeoutOnCloseParams,
) *CreateMsgIBCTimeoutOnClose {
	return &CreateMsgIBCTimeoutOnClose{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgIBCTimeoutOnClose) Name() string {
	return "CreateMsgIBCTimeoutOnClose"
}

func (*CreateMsgIBCTimeoutOnClose) Version() int {
	return 1
}

func (cmd *CreateMsgIBCTimeoutOnClose) Exec() (entity_event.Event, error) {
	event := event.NewMsgIBCTimeoutOnClose(cmd.msgCommonParams, cmd.params)
	return event, nil
}
