package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	icaauthmodel "github.com/crypto-com/chain-indexing/usecase/model/icaauth"
)

type CreateMsgSubmitTx struct {
	msgCommonParams event.MsgCommonParams
	params          icaauthmodel.MsgSubmitTxParams
}

func NewCreateMsgSubmitTx(
	msgCommonParams event.MsgCommonParams,
	params icaauthmodel.MsgSubmitTxParams,
) *CreateMsgSubmitTx {
	return &CreateMsgSubmitTx{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgSubmitTx) Name() string {
	return "/chainmain.icaauth.v1.MsgSubmitTx.Create"
}

func (*CreateMsgSubmitTx) Version() int {
	return 1
}

func (cmd *CreateMsgSubmitTx) Exec() (entity_event.Event, error) {
	event := event.NewMsgSubmitTx(cmd.msgCommonParams, cmd.params)
	return event, nil
}
