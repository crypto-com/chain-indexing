package executor

import (
	"fmt"
	appevent "github.com/crypto-com/chainindex/appinterface/event"
	"github.com/crypto-com/chainindex/appinterface/rdb"
	"github.com/crypto-com/chainindex/entity/command"
	"github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/infrastructure"
	applogger "github.com/crypto-com/chainindex/internal/logger"
	"os"
)

type BlockExecutor struct {
	Height   int64
	Commands []command.Command
	Events   []event.Event
	logger   applogger.Logger
}

func NewBlockExecutor(Hegiht int64) *BlockExecutor {
	commands := make([]command.Command, 0)
	events := make([]event.Event, 0)
	logger := infrastructure.NewZerologLoggerWithColor(os.Stdout)

	return &BlockExecutor{
		Hegiht,
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

func (exec *BlockExecutor) StoreAllEvents(handle *rdb.Handle) error {
	// TODO: tx rollback when has error
	eventStore := appevent.NewRDbStore(handle)
	if err := eventStore.InsertAll(exec.Events); err != nil {
		return fmt.Errorf("executor error storing all events for height %d %v", exec.Height, err)
	}
	return nil
}
