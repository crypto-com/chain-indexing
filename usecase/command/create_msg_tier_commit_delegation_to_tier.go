package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgTierCommitDelegationToTier struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgCommitDelegationToTierParams
}

func NewCreateMsgTierCommitDelegationToTier(msgCommonParams event.MsgCommonParams, params model.MsgCommitDelegationToTierParams) *CreateMsgTierCommitDelegationToTier {
	return &CreateMsgTierCommitDelegationToTier{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgTierCommitDelegationToTier) Name() string {
	return "CreateMsgTierCommitDelegationToTier"
}

func (*CreateMsgTierCommitDelegationToTier) Version() int {
	return 1
}

func (cmd *CreateMsgTierCommitDelegationToTier) Exec() (entity_event.Event, error) {
	event := event.NewMsgTierCommitDelegationToTier(cmd.msgCommonParams, cmd.params)
	return event, nil
}
