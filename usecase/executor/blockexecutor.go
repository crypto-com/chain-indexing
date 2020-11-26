package executor

import (
	"fmt"

	event_interface "github.com/crypto-com/chainindex/appinterface/event"
	"github.com/crypto-com/chainindex/entity/command"
	event_entity "github.com/crypto-com/chainindex/entity/event"
	applogger "github.com/crypto-com/chainindex/internal/logger"
)

type BlockExecutor struct {
	logger applogger.Logger

	Height   int64
	Commands []command.Command
	Events   []event_entity.Event
}

func NewBlockExecutor(logger applogger.Logger, height int64) *BlockExecutor {
	commands := make([]command.Command, 0)
	events := make([]event_entity.Event, 0)

	return &BlockExecutor{
		logger.WithFields(applogger.LogFields{
			"module":      "BlockExecutor",
			"blockHeight": height,
		}),

		height,
		commands,
		events,
	}
}

func (exec *BlockExecutor) AddAllCommands(commands []command.Command) {
	exec.Commands = append(exec.Commands, commands...)
}

func (exec *BlockExecutor) ExecAllCommands() error {
	for _, command := range exec.Commands {
		event, err := command.Exec()
		if err != nil {
			exec.logger.Errorf("Error generating event", err)
			return err
		}
		exec.Events = append(exec.Events, event)
	}
	return nil
}

func (exec *BlockExecutor) StoreAllEvents(eventStore *event_interface.RDbStore) error {
	if err := eventStore.InsertAll(exec.Events); err != nil {
		return fmt.Errorf("executor error storing all events for height %d %v", exec.Height, err)
	}
	return nil
}
