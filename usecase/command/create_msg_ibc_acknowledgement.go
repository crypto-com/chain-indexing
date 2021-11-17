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


type CreateMsgAlreadyRelayedIBCAcknowledgement struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgAcknowledgementParams
}

func NewCreateMsgAlreadyRelayedIBCAcknowledgement(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgAcknowledgementParams,
) *CreateMsgAlreadyRelayedIBCAcknowledgement {
	return &CreateMsgAlreadyRelayedIBCAcknowledgement{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgAlreadyRelayedIBCAcknowledgement) Name() string {
	return "CreateMsgAlreadyRelayedIBCAcknowledgement"
}

func (*CreateMsgAlreadyRelayedIBCAcknowledgement) Version() int {
	return 1
}

func (cmd *CreateMsgAlreadyRelayedIBCAcknowledgement) Exec() (entity_event.Event, error) {
	event := event.NewMsgAlreadyRelayedIBCAcknowledgement(cmd.msgCommonParams, cmd.params)
	return event, nil
}
