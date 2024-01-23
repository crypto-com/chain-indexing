package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	icaauthmodel "github.com/crypto-com/chain-indexing/usecase/model/icaauth"
)

type CreateChainmainMsgRegisterAccount struct {
	msgCommonParams event.MsgCommonParams
	params          icaauthmodel.ChainmainMsgRegisterAccountParams
}

func NewCreateChainmainMsgRegisterAccount(
	msgCommonParams event.MsgCommonParams,
	params icaauthmodel.ChainmainMsgRegisterAccountParams,
) *CreateChainmainMsgRegisterAccount {
	return &CreateChainmainMsgRegisterAccount{
		msgCommonParams,
		params,
	}
}

func (*CreateChainmainMsgRegisterAccount) Name() string {
	return "/chainmain.icaauth.v1.MsgRegisterAccount.Create"
}

func (*CreateChainmainMsgRegisterAccount) Version() int {
	return 1
}

func (cmd *CreateChainmainMsgRegisterAccount) Exec() (entity_event.Event, error) {
	event := event.NewChainmainMsgRegisterAccount(cmd.msgCommonParams, cmd.params)
	return event, nil
}
