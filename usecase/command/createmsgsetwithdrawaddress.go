package command

import (
	entity_event "github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/usecase/event"
	"github.com/crypto-com/chainindex/usecase/model"
)

type CreateMsgSetWithdrawAddress struct {
	msgCommonParams event.MsgCommonParams

	params model.MsgSetWithdrawAddressParams
}

func NewCreateMsgSetWithdrawAddress(
	msgCommonParams event.MsgCommonParams,
	params model.MsgSetWithdrawAddressParams,
) *CreateMsgSetWithdrawAddress {
	return &CreateMsgSetWithdrawAddress{
		msgCommonParams,

		params,
	}
}

func (_ *CreateMsgSetWithdrawAddress) Name() string {
	return "CreateMsgSetWithdrawAddress"
}

func (_ *CreateMsgSetWithdrawAddress) Version() int {
	return 1
}

func (cmd *CreateMsgSetWithdrawAddress) Exec() (entity_event.Event, error) {
	event := event.NewMsgSetWithdrawAddress(
		cmd.msgCommonParams,
		cmd.params,
	)
	return event, nil
}
