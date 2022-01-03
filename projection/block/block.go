package block

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	entity_projection "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper"
	"github.com/crypto-com/chain-indexing/projection/block/view"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ entity_projection.Projection = &Block{}

// TODO: Listen to council node related events and project council node
type Block struct {
	*rdbprojectionbase.Base

	rdbConn rdb.Conn
	logger  applogger.Logger

	migrationHelper migrationhelper.MigrationHelper
}

func NewBlock(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	migrationHelper migrationhelper.MigrationHelper,
) *Block {
	return &Block{
		rdbprojectionbase.NewRDbBase(
			rdbConn.ToHandle(),
			"Block",
		),

		rdbConn,
		logger,

		migrationHelper,
	}
}

func (_ *Block) GetEventsToListen() []string {
	return []string{event_usecase.BLOCK_CREATED}
}

func (projection *Block) OnInit() error {
	if projection.migrationHelper != nil {
		projection.migrationHelper.Migrate()
	}
	return nil
}

func (projection *Block) HandleEvents(height int64, events []event_entity.Event) error {
	rdbTx, err := projection.rdbConn.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction: %v", err)
	}

	committed := false
	defer func() {
		if !committed {
			_ = rdbTx.Rollback()
		}
	}()

	rdbTxHandle := rdbTx.ToHandle()
	blocksView := view.NewBlocks(rdbTxHandle)

	for _, event := range events {
		if blockCreatedEvent, ok := event.(*event_usecase.BlockCreated); ok {
			if handleErr := projection.handleBlockCreatedEvent(blocksView, blockCreatedEvent); handleErr != nil {
				return fmt.Errorf("error handling BlockCreatedEvent: %v", handleErr)
			}
		}
	}
	if err = projection.UpdateLastHandledEventHeight(rdbTxHandle, height); err != nil {
		return fmt.Errorf("error updating last handled event height: %v", err)
	}

	if err = rdbTx.Commit(); err != nil {
		return fmt.Errorf("error committing changes: %v", err)
	}
	committed = true
	return nil
}

func (projection *Block) handleBlockCreatedEvent(blocksView *view.Blocks, event *event_usecase.BlockCreated) error {
	committedCouncilNodes := make([]view.BlockCommittedCouncilNode, 0)
	for _, signature := range event.Block.Signatures {
		committedCouncilNodes = append(committedCouncilNodes, view.BlockCommittedCouncilNode{
			Address:    signature.ValidatorAddress,
			Time:       signature.Timestamp,
			Signature:  signature.Signature,
			IsProposer: event.Block.ProposerAddress == signature.ValidatorAddress,
		})
	}

	if err := blocksView.Insert(&view.Block{
		Height:                event.Block.Height,
		Hash:                  event.Block.Hash,
		Time:                  event.Block.Time,
		AppHash:               event.Block.AppHash,
		TransactionCount:      len(event.Block.Txs),
		CommittedCouncilNodes: committedCouncilNodes,
	}); err != nil {
		return err
	}

	return nil
}
