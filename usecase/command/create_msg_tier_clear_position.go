package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgTierClearPosition struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgClearPositionParams
}

func NewCreateMsgTierClearPosition(msgCommonParams event.MsgCommonParams, params model.MsgClearPositionParams) *CreateMsgTierClearPosition {
	return &CreateMsgTierClearPosition{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgTierClearPosition) Name() string {
	return "CreateMsgTierClearPosition"
}

func (*CreateMsgTierClearPosition) Version() int {
	return 1
}

func (cmd *CreateMsgTierClearPosition) Exec() (entity_event.Event, error) {
	event := event.NewMsgTierClearPosition(cmd.msgCommonParams, cmd.params)
	return event, nil
}
