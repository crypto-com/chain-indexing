package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateCoinReceived struct {
	blockHeight int64
	params      model.CoinReceivedParams
}

func NewCreateCoinReceived(
	blockHeight int64,
	params model.CoinReceivedParams,
) *CreateCoinReceived {
	return &CreateCoinReceived{
		blockHeight,
		params,
	}
}

// Name returns name of command
func (*CreateCoinReceived) Name() string {
	return "CreateCoinReceived"
}

// Version returns version of command
func (*CreateCoinReceived) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateCoinReceived) Exec() (entity_event.Event, error) {
	event := event.NewCoinReceived(cmd.blockHeight, cmd.params)
	return event, nil
}
