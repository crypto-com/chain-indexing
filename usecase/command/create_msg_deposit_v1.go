package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	model_gov_v1 "github.com/crypto-com/chain-indexing/usecase/model/gov/v1"
)

type CreateMsgDepositV1 struct {
	msgCommonParams event.MsgCommonParams
	params          model_gov_v1.MsgDepositParams
}

func NewCreateMsgDepositV1(
	msgCommonParams event.MsgCommonParams,
	params model_gov_v1.MsgDepositParams,
) *CreateMsgDepositV1 {
	return &CreateMsgDepositV1{
		msgCommonParams,
		params,
	}
}

// Name returns name of command
func (*CreateMsgDepositV1) Name() string {
	return "/cosmos.gov.v1.MsgDeposit.Create"
}

// Version returns version of command
func (*CreateMsgDepositV1) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateMsgDepositV1) Exec() (entity_event.Event, error) {
	event := event.NewMsgDepositV1(cmd.msgCommonParams, cmd.params)
	return event, nil
}
