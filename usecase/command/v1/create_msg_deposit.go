package v1_command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	v1_event "github.com/crypto-com/chain-indexing/usecase/event/v1"
	v1_model "github.com/crypto-com/chain-indexing/usecase/model/v1"
)

type CreateMsgDeposit struct {
	msgCommonParams event.MsgCommonParams
	params          v1_model.MsgDepositParams
}

func NewCreateMsgDeposit(
	msgCommonParams event.MsgCommonParams,
	params v1_model.MsgDepositParams,
) *CreateMsgDeposit {
	return &CreateMsgDeposit{
		msgCommonParams,
		params,
	}
}

// Name returns name of command
func (*CreateMsgDeposit) Name() string {
	return "/cosmos.gov.v1.MsgDeposit.Create"
}

// Version returns version of command
func (*CreateMsgDeposit) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateMsgDeposit) Exec() (entity_event.Event, error) {
	event := v1_event.NewMsgDeposit(cmd.msgCommonParams, cmd.params)
	return event, nil
}
