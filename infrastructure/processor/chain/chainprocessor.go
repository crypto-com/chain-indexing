package chain

import (
	"fmt"
	"github.com/crypto-com/chainindex/appinterface/feed"
	"github.com/crypto-com/chainindex/entity/command"
	"github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/infrastructure"
	"github.com/crypto-com/chainindex/infrastructure/notification"
	"github.com/crypto-com/chainindex/infrastructure/processor/chain/parser"
	"os"
)

type ChainProcessor struct {
	Id       int64
	Commands []command.Command
	Events   []event.Event
	logger   *infrastructure.ZerologLogger
}

func NewChainProcessor(Id int64) *ChainProcessor {
	commands := make([]command.Command, 0)
	events := make([]event.Event, 0)
	logger := infrastructure.NewZerologLoggerWithColor(os.Stdout)

	return &ChainProcessor{
		Id,
		commands,
		events,
		logger,
	}
}

func (cp *ChainProcessor) AddAll(commands []command.Command) {
	cp.Commands = append(cp.Commands, commands...)
}

func (cp *ChainProcessor) GenerateAll() error {
	for _, command := range cp.Commands {
		event, error := command.Exec()
		if error != nil {
			cp.logger.Error("Error generating event")
			return error
		}
		cp.Events = append(cp.Events, event)
	}
	return nil
}

func (cp *ChainProcessor) StoreAll() error {
	// TODO: store batch events
	fmt.Println("\t\tTODO: store all events number", len(cp.Events))

	// clear current events and commands
	cp.Commands = make([]command.Command, 0)
	cp.Events = make([]event.Event, 0)
	return nil
}

func (cp *ChainProcessor) NotifyCallback(n *feed.Notification) error {
	fmt.Println("processor", cp.Id, "got notification", n.Name)

	if n.Name == notification.NOTIF_TENDERMINT_BLOCK {
		// TODO: handling type conversion error
		block := n.Payload.(notification.BlockNotifPayload).Block
		rawBlock := n.Payload.(notification.BlockNotifPayload).RawBlock

		cmds, err := parser.ParseBlockToCommands(block, rawBlock)
		if err != nil {
			return fmt.Errorf("error parsing block data to commands %v", err)
		}

		cp.AddAll(cmds)
	}
	
	// Gathered all commands for current block, generate and store the events
	if n.Name == notification.NOTIF_CURRBLOCK_DONE {
		if err := cp.GenerateAll(); err != nil {
			return fmt.Errorf("error generating all events%v", err)
		}
		if err := cp.StoreAll(); err != nil {
			return fmt.Errorf("error storing all events%v", err)
		}
	}

	return nil
}
