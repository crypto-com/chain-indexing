package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateCoinMint struct {
	blockHeight int64
	params      model.CoinMintParams
}

func NewCreateCoinMint(
	blockHeight int64,
	params model.CoinMintParams,
) *CreateCoinMint {
	return &CreateCoinMint{
		blockHeight,
		params,
	}
}

// Name returns name of command
func (*CreateCoinMint) Name() string {
	return "CreateCoinMint"
}

// Version returns version of command
func (*CreateCoinMint) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateCoinMint) Exec() (entity_event.Event, error) {
	event := event.NewCoinMint(cmd.blockHeight, cmd.params)
	return event, nil
}
