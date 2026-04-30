package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgTierExitTierWithDelegation struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgExitTierWithDelegationParams
}

func NewCreateMsgTierExitTierWithDelegation(msgCommonParams event.MsgCommonParams, params model.MsgExitTierWithDelegationParams) *CreateMsgTierExitTierWithDelegation {
	return &CreateMsgTierExitTierWithDelegation{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgTierExitTierWithDelegation) Name() string {
	return "CreateMsgTierExitTierWithDelegation"
}

func (*CreateMsgTierExitTierWithDelegation) Version() int {
	return 1
}

func (cmd *CreateMsgTierExitTierWithDelegation) Exec() (entity_event.Event, error) {
	event := event.NewMsgTierExitTierWithDelegation(cmd.msgCommonParams, cmd.params)
	return event, nil
}
