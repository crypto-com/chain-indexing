package executor

import (
	"github.com/crypto-com/chainindex/entity/command"
	"github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/infrastructure"
	"os"
)

type Base struct {
	Commands []command.Command
	Events   []event.Event
	logger   *infrastructure.ZerologLogger
}

func New() *Base {
	commands := make([]command.Command, 0)
	events := make([]event.Event, 0)
	logger := infrastructure.NewZerologLoggerWithColor(os.Stdout)

	return &Base{
		commands,
		events,
		logger,
	}
}

func (exec *Base) AddCommand(command command.Command) {
	exec.Commands = append(exec.Commands, command)
}

func (exec *Base) AddCommands(commands []command.Command) {
	exec.Commands = append(exec.Commands, commands...)
}

func (exec *Base) GenerateEvent(command command.Command) error {
	event, error := command.Exec()
	if error != nil {
		exec.logger.Error("Error generating event")
		return error
	}
	exec.Events = append(exec.Events, event)
	return nil
}

func (exec *Base) GenerateEvents(commands []command.Command) error {
	for _, command := range commands {
		event, error := command.Exec()
		if error != nil {
			exec.logger.Error("Error generating event")
			return error
		}
		exec.Events = append(exec.Events, event)
	}
	return nil
}

func (exec *Base) Store(event event.Event) error {
	// TODO: store single event
	return nil
}

func (exec *Base) StoreAll(events []event.Event) error {
	// TODO: store batch events
	return nil
}
