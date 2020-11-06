package executor

import (
	"github.com/crypto-com/chainindex/entity/command"
	"github.com/crypto-com/chainindex/entity/event"
)

// Executor interface defines command executor signature
type Executor interface {
	AddCommand(command command.Command)
	AddCommands(commands []command.Command)

	GenerateEvent(command command.Command) error
	GenerateEvents(commands []command.Command) error

	Store(event event.Event) error
	StoreAll(events []event.Event) error
}