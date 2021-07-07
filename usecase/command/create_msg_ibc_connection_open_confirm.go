package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

type CreateMsgIBCConnectionOpenConfirm struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgConnectionOpenConfirmParams
}

func NewCreateMsgIBCConnectionOpenConfirm(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgConnectionOpenConfirmParams,
) *CreateMsgIBCConnectionOpenConfirm {
	return &CreateMsgIBCConnectionOpenConfirm{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgIBCConnectionOpenConfirm) Name() string {
	return "CreateMsgIBCConnectionOpenConfirm"
}

func (*CreateMsgIBCConnectionOpenConfirm) Version() int {
	return 1
}

func (cmd *CreateMsgIBCConnectionOpenConfirm) Exec() (entity_event.Event, error) {
	event := event.NewMsgIBCConnectionOpenConfirm(cmd.msgCommonParams, cmd.params)
	return event, nil
}
