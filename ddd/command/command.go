package command

import "github.com/crypto-com/chainindex/ddd/event"

type Command interface {
	Name() string
	Version() int

	Exec() []event.Event
}
