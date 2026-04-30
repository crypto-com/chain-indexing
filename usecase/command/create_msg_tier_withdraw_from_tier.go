package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgTierWithdrawFromTier struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgWithdrawFromTierParams
}

func NewCreateMsgTierWithdrawFromTier(msgCommonParams event.MsgCommonParams, params model.MsgWithdrawFromTierParams) *CreateMsgTierWithdrawFromTier {
	return &CreateMsgTierWithdrawFromTier{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgTierWithdrawFromTier) Name() string {
	return "CreateMsgTierWithdrawFromTier"
}

func (*CreateMsgTierWithdrawFromTier) Version() int {
	return 1
}

func (cmd *CreateMsgTierWithdrawFromTier) Exec() (entity_event.Event, error) {
	event := event.NewMsgTierWithdrawFromTier(cmd.msgCommonParams, cmd.params)
	return event, nil
}
