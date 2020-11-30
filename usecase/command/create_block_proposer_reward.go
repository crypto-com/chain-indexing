package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
)

type CreateBlockProposerReward struct {
	blockHeight int64
	validator   string
	amount      string
}

func NewCreateBlockProposerReward(blockHeight int64, validator string, amount string) *CreateBlockProposerReward {
	return &CreateBlockProposerReward{
		blockHeight,
		validator,
		amount,
	}
}

// Name returns name of command
func (*CreateBlockProposerReward) Name() string {
	return "CreateBlockProposerReward"
}

// Version returns version of command
func (*CreateBlockProposerReward) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateBlockProposerReward) Exec() (entity_event.Event, error) {
	event := event.NewProposerRewarded(cmd.blockHeight, cmd.validator, cmd.amount)
	return event, nil
}
