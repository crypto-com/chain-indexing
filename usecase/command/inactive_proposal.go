package command

import (
	entity_event "github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/usecase/event"
)

type InactiveProposal struct {
	blockHeight int64

	proposalId string
	result     string
}

func NewInactiveProposal(blockHeight int64, proposalId string, result string) *InactiveProposal {
	return &InactiveProposal{
		blockHeight,
		proposalId,
		result,
	}
}

// Name returns name of command
func (*InactiveProposal) Name() string {
	return "CreateBlockReward"
}

// Version returns version of command
func (*InactiveProposal) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *InactiveProposal) Exec() (entity_event.Event, error) {
	event := event.NewProposalInactived(cmd.blockHeight, cmd.proposalId, cmd.result)
	return event, nil
}
