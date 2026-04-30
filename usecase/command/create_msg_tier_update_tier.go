package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgTierUpdateTier struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgUpdateTierParams
}

func NewCreateMsgTierUpdateTier(msgCommonParams event.MsgCommonParams, params model.MsgUpdateTierParams) *CreateMsgTierUpdateTier {
	return &CreateMsgTierUpdateTier{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgTierUpdateTier) Name() string {
	return "CreateMsgTierUpdateTier"
}

func (*CreateMsgTierUpdateTier) Version() int {
	return 1
}

func (cmd *CreateMsgTierUpdateTier) Exec() (entity_event.Event, error) {
	event := event.NewMsgTierUpdateTier(cmd.msgCommonParams, cmd.params)
	return event, nil
}
