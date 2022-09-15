package blockevent

import (
	"fmt"
	"strconv"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper"
	"github.com/crypto-com/chain-indexing/projection/account_raw_event/view"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ projection.Projection = &AccountRawEvent{}

type AccountRawEvent struct {
	*rdbprojectionbase.Base

	rdbConn rdb.Conn
	logger  applogger.Logger

	migrationHelper migrationhelper.MigrationHelper
}

func NewAccountRawEvent(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	migrationHelper migrationhelper.MigrationHelper,
) *AccountRawEvent {
	return &AccountRawEvent{
		rdbprojectionbase.NewRDbBase(
			rdbConn.ToHandle(),
			"AccountRawEvent",
		),

		rdbConn,
		logger,

		migrationHelper,
	}
}

func (_ *AccountRawEvent) GetEventsToListen() []string {
	return append([]string{
		event_usecase.BLOCK_CREATED,
		event_usecase.TRANSACTION_CREATED,
		event_usecase.TRANSACTION_FAILED,
	}, event_usecase.MSG_EVENTS...)
}

func (projection *AccountRawEvent) OnInit() error {
	if projection.migrationHelper != nil {
		projection.migrationHelper.Migrate()
	}
	return nil
}

func (projection *AccountRawEvent) HandleEvents(height int64, events []event_entity.Event) error {
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
	eventsView := view.NewAccountRawEvents(rdbTxHandle)
	totalView := view.NewAccountRawEventsTotal(rdbTxHandle)

	totalMap := make(map[string]int64)

	var blockTime utctime.UTCTime
	var blockHash string
	eventRows := make([]view.AccountRawEventRow, 0)
	for _, event := range events {
		// TODO
		if blockCreatedEvent, ok := event.(*event_usecase.BlockCreated); ok {
			blockTime = blockCreatedEvent.Block.Time
			blockHash = blockCreatedEvent.Block.Hash
		} else {
			eventRows = append(eventRows, view.AccountRawEventRow{
				BlockHeight: height,
				BlockHash:   "",
				BlockTime:   utctime.UTCTime{},
				Data: view.AccountRawEventRowData{
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
		return fmt.Errorf("error incrementing account raw event type total")
	}
	totalMap["-"] = int64(len(eventRows))
	for key, value := range totalMap {
		if err = totalView.Increment(key, value); err != nil {
			return fmt.Errorf("error incrementing account raw event type total")
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
