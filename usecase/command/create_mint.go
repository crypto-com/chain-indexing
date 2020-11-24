package command

import (
	entity_event "github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/usecase/event"
	"github.com/crypto-com/chainindex/usecase/model"
)

type CreateMint struct {
	blockHeight int64
	params      model.MintParams
}

func NewCreateMint(blockHeight int64, params model.MintParams) *CreateMint {
	return &CreateMint{
		blockHeight,
		params,
	}
}

// Name returns name of command
func (*CreateMint) Name() string {
	return "CreateMint"
}

// Version returns version of command
func (*CreateMint) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateMint) Exec() (entity_event.Event, error) {
	event := event.NewMinted(cmd.blockHeight, cmd.params)
	return event, nil
}
