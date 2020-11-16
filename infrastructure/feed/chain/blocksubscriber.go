package chain

import (
	"fmt"

	applogger "github.com/crypto-com/chainindex/internal/logger"

	appevent "github.com/crypto-com/chainindex/appinterface/event"
	"github.com/crypto-com/chainindex/infrastructure/notification"
	"github.com/crypto-com/chainindex/usecase/executor"
	"github.com/crypto-com/chainindex/usecase/parser"
)

type BlockSubscriber struct {
	moduleAccounts *parser.ModuleAccounts
}

func NewBlockSubscriber(moduleAccounts *parser.ModuleAccounts) *BlockSubscriber {
	return &BlockSubscriber{
		moduleAccounts,
	}
}

func (subscriber *BlockSubscriber) OnNotification(n *notification.BlockNotification, logger applogger.Logger, eventStore *appevent.RDbStore) error {
	// create an executor instance for current height
	executor := executor.NewBlockExecutor(logger, n.Height)
	commands, err := parser.ParseBlockToCommands(
		subscriber.moduleAccounts,
		n.Block,
		n.RawBlock,
		n.BlockResults,
	)
	if err != nil {
		return fmt.Errorf("error parsing block data to commands %v", err)
	}
	executor.AddAllCommands(commands)

	// generate all events, make them persistent
	if err := executor.ExecAllCommands(); err != nil {
		return fmt.Errorf("error generating all events%v", err)
	}
	if err := executor.StoreAllEvents(eventStore); err != nil {
		return fmt.Errorf("error storing all events%v", err)
	}

	return nil
}
