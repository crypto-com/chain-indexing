package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

type CreateMsgIBCUpdateClient struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgUpdateClientParams
}

func NewCreateMsgIBCUpdateClient(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgUpdateClientParams,
) *CreateMsgIBCUpdateClient {
	return &CreateMsgIBCUpdateClient{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgIBCUpdateClient) Name() string {
	return "CreateMsgIBCUpdateClient"
}

func (*CreateMsgIBCUpdateClient) Version() int {
	return 1
}

func (cmd *CreateMsgIBCUpdateClient) Exec() (entity_event.Event, error) {
	event := event.NewMsgIBCUpdateClient(cmd.msgCommonParams, cmd.params)
	return event, nil
}
