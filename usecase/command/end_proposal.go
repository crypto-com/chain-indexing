package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
)

type EndProposal struct {
	blockHeight int64

	proposalId string
	result     string
}

func NewEndProposal(blockHeight int64, proposalId string, result string) *EndProposal {
	return &EndProposal{
		blockHeight,
		proposalId,
		result,
	}
}

// Name returns name of command
func (*EndProposal) Name() string {
	return "CreateBlockReward"
}

// Version returns version of command
func (*EndProposal) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *EndProposal) Exec() (entity_event.Event, error) {
	event := event.NewProposalEnded(cmd.blockHeight, cmd.proposalId, cmd.result)
	return event, nil
}
