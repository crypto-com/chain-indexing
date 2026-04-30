package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgTierDeleteTier struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgDeleteTierParams
}

func NewCreateMsgTierDeleteTier(msgCommonParams event.MsgCommonParams, params model.MsgDeleteTierParams) *CreateMsgTierDeleteTier {
	return &CreateMsgTierDeleteTier{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgTierDeleteTier) Name() string {
	return "CreateMsgTierDeleteTier"
}

func (*CreateMsgTierDeleteTier) Version() int {
	return 1
}

func (cmd *CreateMsgTierDeleteTier) Exec() (entity_event.Event, error) {
	event := event.NewMsgTierDeleteTier(cmd.msgCommonParams, cmd.params)
	return event, nil
}
