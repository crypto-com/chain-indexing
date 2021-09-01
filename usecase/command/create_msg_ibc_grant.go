package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

type CreateMsgGrant struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgGrantParams
}

func NewCreateMsgGrant(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgGrantParams,
) *CreateMsgGrant {
	return &CreateMsgGrant{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgGrant) Name() string {
	return "CreateMsgGrant"
}

func (*CreateMsgGrant) Version() int {
	return 1
}

func (cmd *CreateMsgGrant) Exec() (entity_event.Event, error) {
	event := event.NewMsgIBCGrant(cmd.msgCommonParams, cmd.params)
	return event, nil
}
