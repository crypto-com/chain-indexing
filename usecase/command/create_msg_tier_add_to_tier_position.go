package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgTierAddToTierPosition struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgAddToTierPositionParams
}

func NewCreateMsgTierAddToTierPosition(msgCommonParams event.MsgCommonParams, params model.MsgAddToTierPositionParams) *CreateMsgTierAddToTierPosition {
	return &CreateMsgTierAddToTierPosition{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgTierAddToTierPosition) Name() string {
	return "CreateMsgTierAddToTierPosition"
}

func (*CreateMsgTierAddToTierPosition) Version() int {
	return 1
}

func (cmd *CreateMsgTierAddToTierPosition) Exec() (entity_event.Event, error) {
	event := event.NewMsgTierAddToTierPosition(cmd.msgCommonParams, cmd.params)
	return event, nil
}
