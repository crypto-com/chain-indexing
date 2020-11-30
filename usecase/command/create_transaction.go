package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateTransaction struct {
	blockHeight int64
	params      model.CreateTransactionParams
}

func NewCreateTransaction(blockHeight int64, params model.CreateTransactionParams) *CreateTransaction {
	return &CreateTransaction{
		blockHeight,
		params,
	}
}

func (_ *CreateTransaction) Name() string {
	return "CreateTransaction"
}

func (_ *CreateTransaction) Version() int {
	return 1
}

func (cmd *CreateTransaction) Exec() (entity_event.Event, error) {
	if cmd.params.Code != 0 {
		return event.NewTransactionFailed(cmd.blockHeight, cmd.params), nil
	}
	return event.NewTransactionCreated(cmd.blockHeight, cmd.params), nil
}
