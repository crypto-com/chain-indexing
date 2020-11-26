package command

import (
	entity_event "github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/usecase/event"
	"github.com/crypto-com/chainindex/usecase/model/genesis"
)

type CreateGenesis struct {
	genesis genesis.Genesis
}

func NewCreateGenesis(genesis genesis.Genesis) *CreateGenesis {
	return &CreateGenesis{
		genesis,
	}
}

// Name returns name of command
func (*CreateGenesis) Name() string {
	return "CreateMint"
}

// Version returns version of command
func (*CreateGenesis) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateGenesis) Exec() (entity_event.Event, error) {
	event := event.NewGenesisCreated(cmd.genesis)
	return event, nil
}
