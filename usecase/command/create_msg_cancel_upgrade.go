package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	v1_model "github.com/crypto-com/chain-indexing/usecase/model/v1"
)

type CreateMsgCancelUpgrade struct {
	msgCommonParams event.MsgCommonParams
	params          v1_model.MsgCancelUpgradeParams
}

func NewCreateMsgCancelUpgrade(
	msgCommonParams event.MsgCommonParams,
	params v1_model.MsgCancelUpgradeParams,
) *CreateMsgCancelUpgrade {
	return &CreateMsgCancelUpgrade{
		msgCommonParams,
		params,
	}
}

// Name returns name of command
func (*CreateMsgCancelUpgrade) Name() string {
	return "/cosmos.upgrade.v1beta1.MsgCancelUpgrade.Create"
}

// Version returns version of command
func (*CreateMsgCancelUpgrade) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateMsgCancelUpgrade) Exec() (entity_event.Event, error) {
	event := event.NewMsgCancelUpgrade(cmd.msgCommonParams, cmd.params)
	return event, nil
}
