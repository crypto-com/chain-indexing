package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	v1_model "github.com/crypto-com/chain-indexing/usecase/model/v1"
)

type CreateMsgSoftwareUpgrade struct {
	msgCommonParams event.MsgCommonParams
	params          v1_model.MsgSoftwareUpgradeParams
}

func NewCreateMsgSoftwareUpgrade(
	msgCommonParams event.MsgCommonParams,
	params v1_model.MsgSoftwareUpgradeParams,
) *CreateMsgSoftwareUpgrade {
	return &CreateMsgSoftwareUpgrade{
		msgCommonParams,
		params,
	}
}

// Name returns name of command
func (*CreateMsgSoftwareUpgrade) Name() string {
	return "/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade.Create"
}

// Version returns version of command
func (*CreateMsgSoftwareUpgrade) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateMsgSoftwareUpgrade) Exec() (entity_event.Event, error) {
	event := event.NewMsgSoftwareUpgrade(cmd.msgCommonParams, cmd.params)
	return event, nil
}
