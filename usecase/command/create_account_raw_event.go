package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateAccountRawEvent struct {
	blockHeight int64
	event       model.BlockEvent
}

func NewCreateAccountRawEvent(
	blockHeight int64,
	event model.BlockEvent,
) *CreateAccountRawEvent {
	return &CreateAccountRawEvent{
		blockHeight,
		event,
	}
}

// Name returns name of command
func (*CreateAccountRawEvent) Name() string {
	return "CreateAccountTransfer"
}

// Version returns version of command
func (*CreateAccountRawEvent) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateAccountRawEvent) Exec() (entity_event.Event, error) {
	event := event.NewAccountRawEventCreated(cmd.blockHeight, cmd.event)
	return event, nil
}
