package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

type CreateMsgIBCConnectionOpenAck struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgConnectionOpenAckParams
}

func NewCreateMsgIBCConnectionOpenAck(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgConnectionOpenAckParams,
) *CreateMsgIBCConnectionOpenAck {
	return &CreateMsgIBCConnectionOpenAck{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgIBCConnectionOpenAck) Name() string {
	return "CreateMsgIBCConnectionOpenAck"
}

func (*CreateMsgIBCConnectionOpenAck) Version() int {
	return 1
}

func (cmd *CreateMsgIBCConnectionOpenAck) Exec() (entity_event.Event, error) {
	event := event.NewMsgIBCConnectionOpenAck(cmd.msgCommonParams, cmd.params)
	return event, nil
}
