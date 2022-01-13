package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateCoinBurn struct {
	blockHeight int64
	params      model.CoinBurnParams
}

func NewCreateCoinBurn(
	blockHeight int64,
	params model.CoinBurnParams,
) *CreateCoinBurn {
	return &CreateCoinBurn{
		blockHeight,
		params,
	}
}

// Name returns name of command
func (*CreateCoinBurn) Name() string {
	return "CreateCoinBurn"
}

// Version returns version of command
func (*CreateCoinBurn) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateCoinBurn) Exec() (entity_event.Event, error) {
	event := event.NewCoinBurn(cmd.blockHeight, cmd.params)
	return event, nil
}
