package command

import "github.com/crypto-com/chainindex/ddd"

// Executor interface defines command executor signature
type Executor interface {
	// Exec turns commands to events
	Exec(command Command) error
	ExecAll(commands []Command) error
	// Store makes the events persistent
	Store(event ddd.Event) error
	StoreAll(events []ddd.Event) error
}
