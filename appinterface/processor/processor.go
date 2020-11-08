package processor

import (
	"github.com/crypto-com/chainindex/appinterface/feed"
	"github.com/crypto-com/chainindex/entity/command"
)

// Processor subscribe to the datafeed and process notification to commands,
// also generates and stores the events
type Processor interface {
	AddAll(commands []command.Command)
	GenerateAll() error
	StoreAll() error

	NotifyCallback(n *feed.Notification) error
}
