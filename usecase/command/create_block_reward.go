package command

import (
	entity_event "github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/usecase/event"
)

type CreateBlockReward struct {
	Nam         string
	blockHeight int64
	validator   string
	amount      string
}

func NewCreateBlockReward(blockHeight int64, validator string, amount string) *CreateBlockReward {
	return &CreateBlockReward{
		"CreateBlockReward",
		blockHeight,
		validator,
		amount,
	}
}

// Name returns name of command
func (*CreateBlockReward) Name() string {
	return "CreateBlockReward"
}

// Version returns version of command
func (*CreateBlockReward) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateBlockReward) Exec() (entity_event.Event, error) {
	event := event.NewBlockRewarded(cmd.blockHeight, cmd.validator, cmd.amount)
	return event, nil
}
