package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	icaauthmodel "github.com/crypto-com/chain-indexing/usecase/model/icaauth"
)

type CreateChainmainMsgSubmitTx struct {
	msgCommonParams event.MsgCommonParams
	params          icaauthmodel.MsgSubmitTxParams
}

func NewCreateChainmainMsgSubmitTx(
	msgCommonParams event.MsgCommonParams,
	params icaauthmodel.MsgSubmitTxParams,
) *CreateChainmainMsgSubmitTx {
	return &CreateChainmainMsgSubmitTx{
		msgCommonParams,
		params,
	}
}

func (*CreateChainmainMsgSubmitTx) Name() string {
	return "/chainmain.icaauth.v1.MsgSubmitTx.Create"
}

func (*CreateChainmainMsgSubmitTx) Version() int {
	return 1
}

func (cmd *CreateChainmainMsgSubmitTx) Exec() (entity_event.Event, error) {
	event := event.NewChainmainMsgSubmitTx(cmd.msgCommonParams, cmd.params)
	return event, nil
}
