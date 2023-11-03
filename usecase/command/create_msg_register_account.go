package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	icaauthmodel "github.com/crypto-com/chain-indexing/usecase/model/icaauth"
)

type CreateMsgRegisterAccount struct {
	msgCommonParams event.MsgCommonParams
	params          icaauthmodel.MsgRegisterAccountParams
}

func NewCreateMsgRegisterAccount(
	msgCommonParams event.MsgCommonParams,
	params icaauthmodel.MsgRegisterAccountParams,
) *CreateMsgRegisterAccount {
	return &CreateMsgRegisterAccount{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgRegisterAccount) Name() string {
	return "/icaauth.v1.MsgRegisterAccount.Create"
}

func (*CreateMsgRegisterAccount) Version() int {
	return 1
}

func (cmd *CreateMsgRegisterAccount) Exec() (entity_event.Event, error) {
	event := event.NewMsgRegisterAccount(cmd.msgCommonParams, cmd.params)
	return event, nil
}
