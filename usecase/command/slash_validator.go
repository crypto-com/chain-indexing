package command

import (
	entity_event "github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/usecase/event"
	"github.com/crypto-com/chainindex/usecase/model"
)

type SlashValidator struct {
	blockHeight int64

	params model.SlashValidatorParams
}

func NewSlashValidator(blockHeight int64, params model.SlashValidatorParams) *SlashValidator {
	return &SlashValidator{
		blockHeight,

		params,
	}
}

// Name returns name of command
func (*SlashValidator) Name() string {
	return "SlashValidator"
}

// Version returns version of command
func (*SlashValidator) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *SlashValidator) Exec() (entity_event.Event, error) {
	event := event.NewValidatorSlashed(cmd.blockHeight, cmd.params)
	return event, nil
}
