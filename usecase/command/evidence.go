package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type Evidence struct {
	blockHeight int64
	params      model.EvidenceParams
}

func NewEvidence(blockHeight int64, params model.EvidenceParams) *Evidence {
	return &Evidence{
		blockHeight,
		params,
	}
}

// Name returns name of command
func (*Evidence) Name() string {
	return "Evidence"
}

// Version returns version of command
func (*Evidence) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *Evidence) Exec() (entity_event.Event, error) {
	event := event.NewEvidence(cmd.blockHeight, cmd.params)
	return event, nil
}
