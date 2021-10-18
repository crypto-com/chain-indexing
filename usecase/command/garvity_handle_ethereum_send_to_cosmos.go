package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type GravityHandleEthereumSendToCosmos struct {
	blockHeight int64

	params model.GravityEthereumSendToCosmosHandledEventParams
}

// FIXME: issue the command in parser after gathering real life example
func NewGravityHandleEthereumSendToCosmos(
	blockHeight int64,
	params model.GravityEthereumSendToCosmosHandledEventParams,
) *GravityHandleEthereumSendToCosmos {
	return &GravityHandleEthereumSendToCosmos{
		blockHeight,

		params,
	}
}

// Name returns name of command
func (*GravityHandleEthereumSendToCosmos) Name() string {
	return "GravityHandleEthereumSendToCosmos"
}

// Version returns version of command
func (*GravityHandleEthereumSendToCosmos) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *GravityHandleEthereumSendToCosmos) Exec() (entity_event.Event, error) {
	event := event.NewGravityEthereumSendToCosmosHandled(cmd.blockHeight, cmd.params)
	return event, nil
}
