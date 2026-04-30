package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgTierTriggerExitFromTier struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgTriggerExitFromTierParams
}

func NewCreateMsgTierTriggerExitFromTier(msgCommonParams event.MsgCommonParams, params model.MsgTriggerExitFromTierParams) *CreateMsgTierTriggerExitFromTier {
	return &CreateMsgTierTriggerExitFromTier{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgTierTriggerExitFromTier) Name() string {
	return "CreateMsgTierTriggerExitFromTier"
}

func (*CreateMsgTierTriggerExitFromTier) Version() int {
	return 1
}

func (cmd *CreateMsgTierTriggerExitFromTier) Exec() (entity_event.Event, error) {
	event := event.NewMsgTierTriggerExitFromTier(cmd.msgCommonParams, cmd.params)
	return event, nil
}
