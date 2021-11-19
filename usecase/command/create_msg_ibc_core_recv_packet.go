package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

type CreateMsgIBCRecvPacket struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgRecvPacketParams
}

func NewCreateMsgIBCRecvPacket(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgRecvPacketParams,
) *CreateMsgIBCRecvPacket {
	return &CreateMsgIBCRecvPacket{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgIBCRecvPacket) Name() string {
	return "CreateMsgIBCRecvPacket"
}

func (*CreateMsgIBCRecvPacket) Version() int {
	return 1
}

func (cmd *CreateMsgIBCRecvPacket) Exec() (entity_event.Event, error) {
	event := event.NewMsgIBCRecvPacket(cmd.msgCommonParams, cmd.params)
	return event, nil
}

type CreateMsgAlreadyRelayedIBCRecvPacket struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgRecvPacketParams
}

func NewCreateMsgAlreadyRelayedIBCRecvPacket(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgRecvPacketParams,
) *CreateMsgAlreadyRelayedIBCRecvPacket {
	return &CreateMsgAlreadyRelayedIBCRecvPacket{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgAlreadyRelayedIBCRecvPacket) Name() string {
	return "CreateMsgAlreadyRelayedIBCRecvPacket"
}

func (*CreateMsgAlreadyRelayedIBCRecvPacket) Version() int {
	return 1
}

func (cmd *CreateMsgAlreadyRelayedIBCRecvPacket) Exec() (entity_event.Event, error) {
	event := event.NewMsgAlreadyRelayedIBCRecvPacket(cmd.msgCommonParams, cmd.params)
	return event, nil
}
