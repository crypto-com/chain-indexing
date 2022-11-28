package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgExtensionOptionDynamicFeeTx struct {
	msgCommonParams event.MsgCommonParams
	params          model.EthermintDynamicFeeTxParams
}

func NewCreateMsgExtensionOptionDynamicFeeTx(
	msgCommonParams event.MsgCommonParams,
	params model.EthermintDynamicFeeTxParams,
) *CreateMsgExtensionOptionDynamicFeeTx {
	return &CreateMsgExtensionOptionDynamicFeeTx{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgExtensionOptionDynamicFeeTx) Name() string {
	return "/ethermint.evm.v1.DynamicFeeTx.Create"
}

func (*CreateMsgExtensionOptionDynamicFeeTx) Version() int {
	return 1
}

func (cmd *CreateMsgExtensionOptionDynamicFeeTx) Exec() (entity_event.Event, error) {
	event := event.NewEthermintExtensionOptionDynamicFeeTx(cmd.msgCommonParams, cmd.params)
	return event, nil
}
