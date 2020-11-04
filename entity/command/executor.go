package command

import (
	"github.com/crypto-com/chainindex/ddd"
	"github.com/crypto-com/chainindex/ddd/command"
	"github.com/crypto-com/chainindex/infrastructure"
	"os"
)

type BaseExecutor struct {
	Commands []command.Command
	Events   []ddd.Event
	logger   *infrastructure.ZerologLogger
}

func NewExecutor() *BaseExecutor {
	commands := make([]command.Command, 0)
	events := make([]ddd.Event, 0)
	logger := infrastructure.NewZerologLoggerWithColor(os.Stdout)

	return &BaseExecutor{
		commands,
		events,
		logger,
	}
}

func (exec *BaseExecutor) AddCommand(command command.Command) {
	exec.Commands = append(exec.Commands, command)
}

func (exec *BaseExecutor) AddCommands(commands []command.Command) {
	exec.Commands = append(exec.Commands, commands...)
}

func (exec *BaseExecutor) Exec(command command.Command)  error {
	event, error := command.Exec()
	if error != nil {
		exec.logger.Error("Error generating event")
		return error
	}
	exec.Events = append(exec.Events, event)
	return nil
}

func (exec *BaseExecutor) ExecAll(commands []command.Command) error {
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

func (exec *BaseExecutor) Store(event ddd.Event) error {
	// TODO: store single event
	return nil
}

func (exec *BaseExecutor) StoreAll(events []ddd.Event) error {
	// TODO: store batch events
	return nil
}
