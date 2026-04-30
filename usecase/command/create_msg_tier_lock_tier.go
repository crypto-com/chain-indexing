package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgTierLockTier struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgLockTierParams
}

func NewCreateMsgTierLockTier(msgCommonParams event.MsgCommonParams, params model.MsgLockTierParams) *CreateMsgTierLockTier {
	return &CreateMsgTierLockTier{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgTierLockTier) Name() string {
	return "CreateMsgTierLockTier"
}

func (*CreateMsgTierLockTier) Version() int {
	return 1
}

func (cmd *CreateMsgTierLockTier) Exec() (entity_event.Event, error) {
	event := event.NewMsgTierLockTier(cmd.msgCommonParams, cmd.params)
	return event, nil
}
