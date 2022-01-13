package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateCoinSpent struct {
	blockHeight int64
	params      model.CoinSpentParams
}

func NewCreateCoinSpent(
	blockHeight int64,
	params model.CoinSpentParams,
) *CreateCoinSpent {
	return &CreateCoinSpent{
		blockHeight,
		params,
	}
}

// Name returns name of command
func (*CreateCoinSpent) Name() string {
	return "CreateCoinSpent"
}

// Version returns version of command
func (*CreateCoinSpent) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateCoinSpent) Exec() (entity_event.Event, error) {
	event := event.NewCoinSpent(cmd.blockHeight, cmd.params)
	return event, nil
}
