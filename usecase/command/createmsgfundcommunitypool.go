package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgFundCommunityPool struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgFundCommunityPoolParams
}

func NewCreateMsgFundCommunityPool(
	msgCommonParams event.MsgCommonParams,
	params model.MsgFundCommunityPoolParams,
) *CreateMsgFundCommunityPool {
	return &CreateMsgFundCommunityPool{
		msgCommonParams,
		params,
	}
}

func (_ *CreateMsgFundCommunityPool) Name() string {
	return "CreateMsgFundCommunityPool"
}

func (_ *CreateMsgFundCommunityPool) Version() int {
	return 1
}

func (cmd *CreateMsgFundCommunityPool) Exec() (entity_event.Event, error) {
	event := event.NewMsgFundCommunityPool(cmd.msgCommonParams, cmd.params)
	return event, nil
}
