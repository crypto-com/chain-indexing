package eventhandler

import "github.com/crypto-com/chain-indexing/entity/event"

type Handler interface {
	GetLastHandledEventHeight() (*int64, error)

	HandleEvents(blockHeight int64, events []event.Event) error
}
