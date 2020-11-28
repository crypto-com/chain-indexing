package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateCompleteBonding struct {
	blockHeight int64
	params      model.CompleteBondingParams
}

func NewCreateCompleteBonding(blockHeight int64, params model.CompleteBondingParams) *CreateCompleteBonding {
	return &CreateCompleteBonding{
		blockHeight,

		params,
	}
}

// Name returns name of command
func (*CreateCompleteBonding) Name() string {
	return "CreateBlockReward"
}

// Version returns version of command
func (*CreateCompleteBonding) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateCompleteBonding) Exec() (entity_event.Event, error) {
	event := event.NewBondingCompleted(cmd.blockHeight, cmd.params)
	return event, nil
}
