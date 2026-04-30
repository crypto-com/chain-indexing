package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgTierAddTier struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgAddTierParams
}

func NewCreateMsgTierAddTier(msgCommonParams event.MsgCommonParams, params model.MsgAddTierParams) *CreateMsgTierAddTier {
	return &CreateMsgTierAddTier{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgTierAddTier) Name() string {
	return "CreateMsgTierAddTier"
}

func (*CreateMsgTierAddTier) Version() int {
	return 1
}

func (cmd *CreateMsgTierAddTier) Exec() (entity_event.Event, error) {
	event := event.NewMsgTierAddTier(cmd.msgCommonParams, cmd.params)
	return event, nil
}
