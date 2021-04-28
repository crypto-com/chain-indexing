package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateCompleteUnbonding struct {
	blockHeight int64
	params      model.CompleteBondingParams
}

func NewCreateCompleteBonding(blockHeight int64, params model.CompleteBondingParams) *CreateCompleteUnbonding {
	return &CreateCompleteUnbonding{
		blockHeight,

		params,
	}
}

// Name returns name of command
func (*CreateCompleteUnbonding) Name() string {
	return "CreateCompleteUnbonding"
}

// Version returns version of command
func (*CreateCompleteUnbonding) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateCompleteUnbonding) Exec() (entity_event.Event, error) {
	event := event.NewUnbondingCompleted(cmd.blockHeight, cmd.params)
	return event, nil
}
