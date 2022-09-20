package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateLegacyTx struct {
	msgCommonParams event.MsgCommonParams
	params          model.EthermintLegacyTxParams
}

func NewCreateLegacyTx(
	msgCommonParams event.MsgCommonParams,
	params model.EthermintLegacyTxParams,
) *CreateLegacyTx {
	return &CreateLegacyTx{
		msgCommonParams,
		params,
	}
}

func (*CreateLegacyTx) Name() string {
	return "/ethermint.evm.v1.LegacyTx.Create"
}

func (*CreateLegacyTx) Version() int {
	return 1
}

func (cmd *CreateLegacyTx) Exec() (entity_event.Event, error) {
	event := event.NewEthermintLegacyx(cmd.msgCommonParams, cmd.params)
	return event, nil
}
