package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

type CreateMsgCreateClient struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgCreateClientParams
}

func NewCreateMsgCreateClient(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgCreateClientParams,
) *CreateMsgCreateClient {
	return &CreateMsgCreateClient{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgCreateClient) Name() string {
	return "CreateMsgCreateClient"
}

func (*CreateMsgCreateClient) Version() int {
	return 1
}

func (cmd *CreateMsgCreateClient) Exec() (entity_event.Event, error) {
	event := event.NewMsgCreateClient(cmd.msgCommonParams, cmd.params)
	return event, nil
}
