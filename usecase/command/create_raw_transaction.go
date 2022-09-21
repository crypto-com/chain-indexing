package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateRawTransaction struct {
	blockHeight int64
	params      model.CreateRawTransactionParams
}

func NewCreateRawTransaction(blockHeight int64, params model.CreateRawTransactionParams) *CreateRawTransaction {
	return &CreateRawTransaction{
		blockHeight,
		params,
	}
}

func (_ *CreateRawTransaction) Name() string {
	return "CreateRawTransaction"
}

func (_ *CreateRawTransaction) Version() int {
	return 1
}

func (cmd *CreateRawTransaction) Exec() (entity_event.Event, error) {
	if cmd.params.Code != 0 {
		return event.NewRawTransactionFailed(cmd.blockHeight, cmd.params), nil
	}
	return event.NewRawTransactionCreated(cmd.blockHeight, cmd.params), nil
}
