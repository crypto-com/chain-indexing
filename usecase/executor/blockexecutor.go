package executor

import (
	"fmt"
	"os"

	event_interface "github.com/crypto-com/chainindex/appinterface/event"
	"github.com/crypto-com/chainindex/entity/command"
	event_entity "github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/infrastructure"
	applogger "github.com/crypto-com/chainindex/internal/logger"
)

type BlockExecutor struct {
	Height   int64
	Commands []command.Command
	Events   []event_entity.Event
	logger   applogger.Logger
}

func NewBlockExecutor(height int64) *BlockExecutor {
	commands := make([]command.Command, 0)
	events := make([]event_entity.Event, 0)
	logger := infrastructure.NewZerologLoggerWithColor(os.Stdout)

	return &BlockExecutor{
		height,
		commands,
		events,
		logger,
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
	// TODO: tx rollback when has error
	if err := eventStore.InsertAll(exec.Events); err != nil {
		return fmt.Errorf("executor error storing all events for height %d %v", exec.Height, err)
	}
	return nil
}
