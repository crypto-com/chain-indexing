package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

type CreateMsgCreateVestingAccount struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgCreateVestingAccountParams
}

func NewCreateMsgCreateVestingAccount(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgCreateVestingAccountParams,
) *CreateMsgCreateVestingAccount {
	return &CreateMsgCreateVestingAccount{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgCreateVestingAccount) Name() string {
	return "CreateMsgCreateVestingAccount"
}

func (*CreateMsgCreateVestingAccount) Version() int {
	return 1
}

func (cmd *CreateMsgCreateVestingAccount) Exec() (entity_event.Event, error) {
	event := event.NewMsgCreateVestingAccount(cmd.msgCommonParams, cmd.params)
	return event, nil
}
