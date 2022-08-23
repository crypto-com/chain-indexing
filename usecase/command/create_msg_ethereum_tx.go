package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgEthereumTx struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgEthereumTxParams
}

func NewCreateMsgEthereumTx(
	msgCommonParams event.MsgCommonParams,
	params model.MsgEthereumTxParams,
) *CreateMsgEthereumTx {
	return &CreateMsgEthereumTx{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgEthereumTx) Name() string {
	return "CreateMsgEthereumTx"
}

func (*CreateMsgEthereumTx) Version() int {
	return 1
}

func (cmd *CreateMsgEthereumTx) Exec() (entity_event.Event, error) {
	event := event.NewMsgEthereumTx(cmd.msgCommonParams, cmd.params)
	return event, nil
}
