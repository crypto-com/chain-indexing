package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

type CreateMsgIBCCreateClient struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgCreateClientParams
}

func NewCreateMsgIBCCreateClient(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgCreateClientParams,
) *CreateMsgIBCCreateClient {
	return &CreateMsgIBCCreateClient{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgIBCCreateClient) Name() string {
	return "/ibc.core.client.v1.MsgCreateClient.Create"
}

func (*CreateMsgIBCCreateClient) Version() int {
	return 1
}

func (cmd *CreateMsgIBCCreateClient) Exec() (entity_event.Event, error) {
	event := event.NewMsgIBCCreateClient(cmd.msgCommonParams, cmd.params)
	return event, nil
}
