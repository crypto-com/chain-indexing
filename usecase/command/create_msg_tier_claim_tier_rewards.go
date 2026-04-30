package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgTierClaimTierRewards struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgClaimTierRewardsParams
}

func NewCreateMsgTierClaimTierRewards(msgCommonParams event.MsgCommonParams, params model.MsgClaimTierRewardsParams) *CreateMsgTierClaimTierRewards {
	return &CreateMsgTierClaimTierRewards{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgTierClaimTierRewards) Name() string {
	return "CreateMsgTierClaimTierRewards"
}

func (*CreateMsgTierClaimTierRewards) Version() int {
	return 1
}

func (cmd *CreateMsgTierClaimTierRewards) Exec() (entity_event.Event, error) {
	event := event.NewMsgTierClaimTierRewards(cmd.msgCommonParams, cmd.params)
	return event, nil
}
