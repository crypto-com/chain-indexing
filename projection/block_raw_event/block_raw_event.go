package block_raw_event

import (
	"fmt"
	"strconv"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper"
	"github.com/crypto-com/chain-indexing/projection/block_raw_event/view"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ projection.Projection = &BlockRawEvent{}

type BlockRawEvent struct {
	*rdbprojectionbase.Base

	rdbConn rdb.Conn
	logger  applogger.Logger

	migrationHelper migrationhelper.MigrationHelper
}

func NewBlockRawEvent(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	migrationHelper migrationhelper.MigrationHelper,
) *BlockRawEvent {
	return &BlockRawEvent{
		rdbprojectionbase.NewRDbBase(
			rdbConn.ToHandle(),
			"BlockRawEvent",
		),

		rdbConn,
		logger,

		migrationHelper,
	}
}

func (_ *BlockRawEvent) GetEventsToListen() []string {
	return []string{
		event_usecase.BLOCK_RAW_EVENT_CREATED,
	}
}

func (projection *BlockRawEvent) OnInit() error {
	if projection.migrationHelper != nil {
		projection.migrationHelper.Migrate()
	}
	return nil
}

func (projection *BlockRawEvent) HandleEvents(height int64, events []event_entity.Event) error {
	var err error

	var rdbTx rdb.Tx
	rdbTx, err = projection.rdbConn.Begin()
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
	eventsView := view.NewBlockRawEvents(rdbTxHandle)
	totalView := view.NewBlockRawEventsTotal(rdbTxHandle)

	totalMap := make(map[string]int64)

	eventRows := make([]view.BlockRawEventRow, 0)
	for _, event := range events {
		if rawBlockCreatedEvent, ok := event.(*event_usecase.BlockRawEventCreated); ok {
			eventRows = append(eventRows, view.BlockRawEventRow{
				BlockHeight: rawBlockCreatedEvent.BlockHeight,
				BlockHash:   rawBlockCreatedEvent.BlockHash,
				BlockTime:   rawBlockCreatedEvent.BlockTime,
				FromResult:  rawBlockCreatedEvent.FromResult,
				Data: view.BlockRawEventRowData{
					Type:    rawBlockCreatedEvent.Data.Type,
					Content: rawBlockCreatedEvent.Data.Content,
				}})

		}
	}
	if err = totalView.Set(strconv.FormatInt(height, 10), int64(len(eventRows))); err != nil {
		return fmt.Errorf("error incrementing block raw event type total")
	}
	totalMap["-"] = int64(len(eventRows))
	for key, value := range totalMap {
		if err = totalView.Increment(key, value); err != nil {
			return fmt.Errorf("error incrementing block raw event type total")
		}
	}
	if err = eventsView.InsertAll(eventRows); err != nil {
		return fmt.Errorf("error batch inserting events into view: %v", err)
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
