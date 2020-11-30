package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
)

type JailValidator struct {
	blockHeight int64

	consensusNodeAddress string
	reason               string
}

func NewJailValidator(blockHeight int64, consensusNodeAddress string, reason string) *JailValidator {
	return &JailValidator{
		blockHeight,

		consensusNodeAddress,
		reason,
	}
}

// Name returns name of command
func (*JailValidator) Name() string {
	return "JailValidator"
}

// Version returns version of command
func (*JailValidator) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *JailValidator) Exec() (entity_event.Event, error) {
	event := event.NewValidatorJailed(cmd.blockHeight, cmd.consensusNodeAddress, cmd.reason)
	return event, nil
}
