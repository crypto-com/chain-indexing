package command

import "github.com/crypto-com/chainindex/entity/event"

type Command interface {
	Name() string
	Version() int

	// Exec process the command data and return the event accordingly
	// Currently one command will generates only one event
	Exec() (event.Event, error)
}
