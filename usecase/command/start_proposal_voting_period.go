package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
)

type StartProposalVotingPeriod struct {
	blockHeight int64
	proposalId  string
}

func NewStartProposalVotingPeriod(
	blockHeight int64,
	proposalId string,
) *StartProposalVotingPeriod {
	return &StartProposalVotingPeriod{
		blockHeight,
		proposalId,
	}
}

// Name returns name of command
func (*StartProposalVotingPeriod) Name() string {
	return "StartProposalVotingPeriod"
}

// Version returns version of command
func (*StartProposalVotingPeriod) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *StartProposalVotingPeriod) Exec() (entity_event.Event, error) {
	event := event.NewProposalVotingPeriodStarted(cmd.blockHeight, cmd.proposalId)
	return event, nil
}
