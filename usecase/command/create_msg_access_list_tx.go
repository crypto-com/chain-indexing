package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgAccessListTx struct {
	msgCommonParams event.MsgCommonParams
	params          model.EthermintAccessListTxParams
}

func NewCreateMsgAccessListTx(
	msgCommonParams event.MsgCommonParams,
	params model.EthermintAccessListTxParams,
) *CreateMsgAccessListTx {
	return &CreateMsgAccessListTx{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgAccessListTx) Name() string {
	return "/ethermint.evm.v1.DynamicFeeTx.Create"
}

func (*CreateMsgAccessListTx) Version() int {
	return 1
}

func (cmd *CreateMsgAccessListTx) Exec() (entity_event.Event, error) {
	event := event.NewEthermintAccessListTx(cmd.msgCommonParams, cmd.params)
	return event, nil
}
