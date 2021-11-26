package blockevent

import (
	"fmt"
	"strconv"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper"
	"github.com/crypto-com/chain-indexing/projection/blockevent/view"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

type BlockEvent struct {
	*rdbprojectionbase.Base

	rdbConn rdb.Conn
	logger  applogger.Logger

	migrationHelper migrationhelper.MigrationHelper
}

func NewBlockEvent(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	migrationHelper migrationhelper.MigrationHelper,
) *BlockEvent {
	return &BlockEvent{
		rdbprojectionbase.NewRDbBase(
			rdbConn.ToHandle(),
			"BlockEvent",
		),

		rdbConn,
		logger,

		migrationHelper,
	}
}

func (_ *BlockEvent) GetEventsToListen() []string {
	return []string{
		event_usecase.BLOCK_CREATED,

		event_usecase.BLOCK_PROPOSER_REWARDED,
		event_usecase.BLOCK_REWARDED,
		event_usecase.BLOCK_COMMISSIONED,
		event_usecase.MINTED,
		event_usecase.PROPOSAL_ENDED,
		event_usecase.PROPOSAL_INACTIVED,
		event_usecase.VALIDATOR_SLASHED,
		event_usecase.VALIDATOR_JAILED,
	}
}

// TODO: should change it to projection folder name to `block_event`, then we can remove it
const (
	MIGRATION_DIRECOTRY = "projection/blockevent/migrations"
)

func (projection *BlockEvent) OnInit() error {
	if projection.migrationHelper != nil {
		projection.migrationHelper.Migrate()
	}
	return nil
}

func (projection *BlockEvent) HandleEvents(height int64, events []event_entity.Event) error {
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
	eventsView := view.NewBlockEvents(rdbTxHandle)
	totalView := view.NewBlockEventsTotal(rdbTxHandle)

	totalMap := make(map[string]int64)

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

			heightEventTypeKey := fmt.Sprintf("%d:%s", height, event.Name())
			if _, ok := totalMap[heightEventTypeKey]; !ok {
				totalMap[heightEventTypeKey] = 0
			}
			totalMap[heightEventTypeKey] += 1

			eventTypeKey := fmt.Sprintf("-:%s", event.Name())
			if _, ok := totalMap[eventTypeKey]; !ok {
				totalMap[eventTypeKey] = 0
			}
			totalMap[eventTypeKey] += 1
		}
	}
	if err = totalView.Set(strconv.FormatInt(height, 10), int64(len(eventRows))); err != nil {
		return fmt.Errorf("error incrementing block event type total")
	}
	totalMap["-"] = int64(len(eventRows))
	for key, value := range totalMap {
		if err = totalView.Increment(key, value); err != nil {
			return fmt.Errorf("error incrementing block event type total")
		}
	}

	for i := range eventRows {
		mutEventRow := &eventRows[i]
		mutEventRow.BlockTime = blockTime
		mutEventRow.BlockHash = blockHash
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
