package chain

import (
	"fmt"
	"github.com/crypto-com/chainindex/infrastructure/feed/chain/parser"
	"github.com/crypto-com/chainindex/infrastructure/notification"
	"github.com/crypto-com/chainindex/usecase/executor"
)

type BlockSubscriber struct {
	Id int64
}

func NewBlockSubscriber(id int64) *BlockSubscriber {
	return &BlockSubscriber{
		Id: id,
	}
}

func (bs *BlockSubscriber) OnNotification(n *notification.BlockNotification) error {
	fmt.Println("processor", bs.Id, "got notification", n.Height, n.Block)

	// create an executor instance for current height
	executor := executor.NewBlockExecutor(n.Height)
	cmds, err := parser.ParseBlockToCommands(n.Block, n.RawBlock)
	if err != nil {
		return fmt.Errorf("error parsing block data to commands %v", err)
	}
	executor.AddAll(cmds)

	// generate all events, make them persistent
	if err := executor.GenerateAll(); err != nil {
		return fmt.Errorf("error generating all events%v", err)
	}
	if err := executor.StoreAll(); err != nil {
		return fmt.Errorf("error storing all events%v", err)
	}

	return nil
}
