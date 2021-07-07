package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

type CreateMsgIBCAcknowledgement struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgAcknowledgementParams
}

func NewCreateMsgIBCAcknowledgement(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgAcknowledgementParams,
) *CreateMsgIBCAcknowledgement {
	return &CreateMsgIBCAcknowledgement{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgIBCAcknowledgement) Name() string {
	return "CreateMsgIBCAcknowledgement"
}

func (*CreateMsgIBCAcknowledgement) Version() int {
	return 1
}

func (cmd *CreateMsgIBCAcknowledgement) Exec() (entity_event.Event, error) {
	event := event.NewMsgIBCAcknowledgement(cmd.msgCommonParams, cmd.params)
	return event, nil
}
