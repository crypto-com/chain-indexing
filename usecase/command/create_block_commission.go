package command

import (
	entity_event "github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/usecase/event"
)

type CreateBlockCommission struct {
	blockHeight int64
	validator   string
	amount      string
}

func NewCreateBlockCommission(blockHeight int64, validator string, amount string) *CreateBlockCommission {
	return &CreateBlockCommission{
		blockHeight,
		validator,
		amount,
	}
}

// Name returns name of command
func (*CreateBlockCommission) Name() string {
	return "CreateBlockCommission"
}

// Version returns version of command
func (*CreateBlockCommission) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateBlockCommission) Exec() (entity_event.Event, error) {
	event := event.NewBlockCommissioned(cmd.blockHeight, cmd.validator, cmd.amount)
	return event, nil
}
