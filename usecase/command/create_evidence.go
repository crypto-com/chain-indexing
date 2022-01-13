package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateEvidence struct {
	blockHeight int64
	params      model.EvidenceParams
}

func NewCreateEvidence(blockHeight int64, params model.EvidenceParams) *CreateEvidence {
	return &CreateEvidence{
		blockHeight,
		params,
	}
}

// Name returns name of command
func (*CreateEvidence) Name() string {
	return "CreateEvidence"
}

// Version returns version of command
func (*CreateEvidence) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateEvidence) Exec() (entity_event.Event, error) {
	event := event.NewEvidenceSubmitted(cmd.blockHeight, cmd.params)
	return event, nil
}
