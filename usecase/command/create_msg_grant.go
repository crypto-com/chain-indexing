package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgGrant struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgGrantParams
}

func NewCreateMsgGrant(
	msgCommonParams event.MsgCommonParams,
	params model.MsgGrantParams,
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
	event := event.NewMsgGrant(cmd.msgCommonParams, cmd.params)
	return event, nil
}
