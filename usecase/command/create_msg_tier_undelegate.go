package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgTierUndelegate struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgTierUndelegateParams
}

func NewCreateMsgTierUndelegate(msgCommonParams event.MsgCommonParams, params model.MsgTierUndelegateParams) *CreateMsgTierUndelegate {
	return &CreateMsgTierUndelegate{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgTierUndelegate) Name() string {
	return "CreateMsgTierUndelegate"
}

func (*CreateMsgTierUndelegate) Version() int {
	return 1
}

func (cmd *CreateMsgTierUndelegate) Exec() (entity_event.Event, error) {
	event := event.NewMsgTierUndelegate(cmd.msgCommonParams, cmd.params)
	return event, nil
}
