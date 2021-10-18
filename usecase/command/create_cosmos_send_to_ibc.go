package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateCosmosSendToIBC struct {
	blockHeight int64
	params      model.CosmosSendToIBCParams
}

func NewCreateCosmosSendToIBC(
	blockHeight int64,
	params model.CosmosSendToIBCParams,
) *CreateCosmosSendToIBC {
	return &CreateCosmosSendToIBC{
		blockHeight,
		params,
	}
}

// Name returns name of command
func (*CreateCosmosSendToIBC) Name() string {
	return "CreateCosmosSendToIBC"
}

// Version returns version of command
func (*CreateCosmosSendToIBC) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateCosmosSendToIBC) Exec() (entity_event.Event, error) {
	return event.NewCronosSendToIBCCreated(cmd.blockHeight, cmd.params), nil
}
