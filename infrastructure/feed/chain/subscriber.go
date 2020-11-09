package chain

import (
	"errors"
	"fmt"
	"github.com/crypto-com/chainindex/appinterface/feed"
	"github.com/crypto-com/chainindex/infrastructure/feed/chain/parser"
	"github.com/crypto-com/chainindex/infrastructure/notification"
	"github.com/crypto-com/chainindex/usecase/executor"
)

type BlockSubscriber struct {
	Id       int64
	Executor *executor.BlockExecutor
}

func NewBlockSubscriber(Id int64) *BlockSubscriber {
	return &BlockSubscriber{
		Id:       Id,
		Executor: nil,
	}
}

func (bs *BlockSubscriber) OnNotification(n *feed.Notification) error {
	fmt.Println("processor", bs.Id, "got notification", n.Name)

	if n.Name == notification.NOTIF_NEW_BLOCK {
		h := n.Payload.(notification.NewBlockPayload).Height
		if bs.Executor == nil {
			bs.Executor = executor.NewBlockExecutor(h)
		}
	}

	if n.Name == notification.NOTIF_TENDERMINT_BLOCK_RECEIVED {
		if bs.Executor == nil {
			return errors.New("block subscriber needs to init an executor for current block, pass NOTIF_NEW_BLOCK to resolve the problem")
		}

		// TODO: handling type conversion error
		block := n.Payload.(notification.BlockReceivedPayload).Block
		rawBlock := n.Payload.(notification.BlockReceivedPayload).RawBlock

		cmds, err := parser.ParseBlockToCommands(block, rawBlock)
		if err != nil {
			return fmt.Errorf("error parsing block data to commands %v", err)
		}

		bs.Executor.AddAll(cmds)
	}

	// Gathered all commands for current block, generate and store the events
	if n.Name == notification.NOTIF_CURRBLOCK_DONE {
		if bs.Executor == nil {
			return errors.New("block subscriber needs to init an executor for current block, pass NOTIF_NEW_BLOCK to resolve the problem")
		}

		// generate all events, make them persistent
		if err := bs.Executor.GenerateAll(); err != nil {
			return fmt.Errorf("error generating all events%v", err)
		}
		if err := bs.Executor.StoreAll(); err != nil {
			return fmt.Errorf("error storing all events%v", err)
		}

		// clear executor
		bs.Executor = nil
	}

	return nil
}
