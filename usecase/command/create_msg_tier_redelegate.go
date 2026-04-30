package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgTierRedelegate struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgTierRedelegateParams
}

func NewCreateMsgTierRedelegate(msgCommonParams event.MsgCommonParams, params model.MsgTierRedelegateParams) *CreateMsgTierRedelegate {
	return &CreateMsgTierRedelegate{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgTierRedelegate) Name() string {
	return "CreateMsgTierRedelegate"
}

func (*CreateMsgTierRedelegate) Version() int {
	return 1
}

func (cmd *CreateMsgTierRedelegate) Exec() (entity_event.Event, error) {
	event := event.NewMsgTierRedelegate(cmd.msgCommonParams, cmd.params)
	return event, nil
}
