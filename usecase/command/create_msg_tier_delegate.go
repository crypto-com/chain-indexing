package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgTierDelegate struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgTierDelegateParams
}

func NewCreateMsgTierDelegate(msgCommonParams event.MsgCommonParams, params model.MsgTierDelegateParams) *CreateMsgTierDelegate {
	return &CreateMsgTierDelegate{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgTierDelegate) Name() string {
	return "CreateMsgTierDelegate"
}

func (*CreateMsgTierDelegate) Version() int {
	return 1
}

func (cmd *CreateMsgTierDelegate) Exec() (entity_event.Event, error) {
	event := event.NewMsgTierDelegate(cmd.msgCommonParams, cmd.params)
	return event, nil
}
