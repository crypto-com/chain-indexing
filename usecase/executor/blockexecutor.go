package executor

import (
	"fmt"
	"os"

	"github.com/crypto-com/chainindex/entity/command"
	"github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/infrastructure"
	applogger "github.com/crypto-com/chainindex/internal/logger"
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

func (ce *BlockExecutor) AddAllCommands(commands []command.Command) {
	ce.Commands = append(ce.Commands, commands...)
}

func (ce *BlockExecutor) ExecAllComands() error {
	for _, command := range ce.Commands {
		event, error := command.Exec()
		if error != nil {
			ce.logger.Error("Error generating event")
			return error
		}
		ce.Events = append(ce.Events, event)
	}
	return nil
}

func (ce *BlockExecutor) StoreAllEvents() error {
	// TODO: store batch events
	fmt.Println("\t\tTODO: store all events number", len(ce.Events))

	return nil
}
