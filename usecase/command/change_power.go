package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type ChangePower struct {
	blockHeight int64
	params      model.PowerChangeParams
}

func NewChangePower(blockHeight int64, params model.PowerChangeParams) *ChangePower {
	return &ChangePower{
		blockHeight,
		params,
	}
}

// Name returns name of command
func (*ChangePower) Name() string {
	return "ChangePower"
}

// Version returns version of command
func (*ChangePower) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *ChangePower) Exec() (entity_event.Event, error) {
	event := event.NewPowerChanged(cmd.blockHeight, cmd.params)
	return event, nil
}
