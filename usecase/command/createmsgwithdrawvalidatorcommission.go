package command

import (
	entity_event "github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/usecase/event"
	"github.com/crypto-com/chainindex/usecase/model"
)

type CreateMsgWithdrawValidatorCommission struct {
	msgCommonParams event.MsgCommonParams

	params model.MsgWithdrawValidatorCommissionParams
}

func NewCreateMsgWithdrawValidatorCommission(
	msgCommonParams event.MsgCommonParams,
	params model.MsgWithdrawValidatorCommissionParams,
) *CreateMsgWithdrawValidatorCommission {
	return &CreateMsgWithdrawValidatorCommission{
		msgCommonParams,

		params,
	}
}

func (_ *CreateMsgWithdrawValidatorCommission) Name() string {
	return "CreateMsgWithdrawDelegatorReward"
}

func (_ *CreateMsgWithdrawValidatorCommission) Version() int {
	return 1
}

func (cmd *CreateMsgWithdrawValidatorCommission) Exec() (entity_event.Event, error) {
	event := event.NewMsgWithdrawValidatorCommission(
		cmd.msgCommonParams,
		cmd.params,
	)
	return event, nil
}
