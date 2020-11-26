package projection

import (
	"fmt"

	"github.com/crypto-com/chainindex/internal/utctime"

	"github.com/crypto-com/chainindex/appinterface/projection/rdbbase"
	"github.com/crypto-com/chainindex/appinterface/projection/view"
	"github.com/crypto-com/chainindex/appinterface/rdb"
	event_entity "github.com/crypto-com/chainindex/entity/event"
	applogger "github.com/crypto-com/chainindex/internal/logger"
	event_usecase "github.com/crypto-com/chainindex/usecase/event"
)

type BlockEvent struct {
	*rdbbase.RDbBase

	rdbConn rdb.Conn
	logger  applogger.Logger
}

func NewBlockEvent(logger applogger.Logger, rdbConn rdb.Conn) *BlockEvent {
	return &BlockEvent{
		rdbbase.NewRDbBase(rdbConn.ToHandle(), "BlockEvent"),

		rdbConn,
		logger,
	}
}

func (_ *BlockEvent) GetEventsToListen() []string {
	return []string{
		event_usecase.BLOCK_CREATED,
		event_usecase.TRANSACTION_CREATED,
		event_usecase.TRANSACTION_FAILED,

		event_usecase.BLOCK_PROPOSER_REWARDED,
		event_usecase.BLOCK_REWARDED,
		event_usecase.BLOCK_COMMISSIONED,
		event_usecase.MINTED,
		event_usecase.PROPOSAL_ENDED,
		event_usecase.PROPOSAL_INACTIVED,
	}
}

func (projection *BlockEvent) OnInit() error {
	return nil
}

func (projection *BlockEvent) HandleEvents(height int64, events []event_entity.Event) error {
	var err error

	var rdbTx rdb.Tx
	rdbTx, err = projection.rdbConn.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction: %v", err)
	}
	rdbTxHandle := rdbTx.ToHandle()
	eventsView := view.NewBlockEvents(rdbTxHandle)

	var blockTime utctime.UTCTime
	var blockHash string
	eventRows := make([]view.BlockEventRow, 0)
	for _, event := range events {
		if blockCreatedEvent, ok := event.(*event_usecase.BlockCreated); ok {
			blockTime = blockCreatedEvent.Block.Time
			blockHash = blockCreatedEvent.Block.Hash
		} else {
			eventRows = append(eventRows, view.BlockEventRow{
				BlockHeight: height,
				BlockHash:   "",
				BlockTime:   utctime.UTCTime{},
				Data: view.BlockEventRowData{
					Type:    event.Name(),
					Content: event,
				},
			})
		}
	}

	for i := range eventRows {
		mutEventRow := eventRows[i]
		mutEventRow.BlockTime = blockTime
		mutEventRow.BlockHash = blockHash

	}
	if insertErr := eventsView.InsertAll(eventRows); insertErr != nil {
		_ = rdbTx.Rollback()
		return fmt.Errorf("error batch inserting events into view: %v", insertErr)
	}

	if err = projection.UpdateLastHandledEventHeight(rdbTxHandle, height); err != nil {
		_ = rdbTx.Rollback()
		return fmt.Errorf("error updating last handled event height: %v", err)
	}

	if err = rdbTx.Commit(); err != nil {
		_ = rdbTx.Rollback()
		return fmt.Errorf("error committing changes: %v", err)
	}
	return nil
}
