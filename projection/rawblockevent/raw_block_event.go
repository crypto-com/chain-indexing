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
	"github.com/crypto-com/chain-indexing/projection/rawblockevent/view"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ projection.Projection = &RawBlockEvent{}

type RawBlockEvent struct {
	*rdbprojectionbase.Base

	rdbConn rdb.Conn
	logger  applogger.Logger

	migrationHelper migrationhelper.MigrationHelper
}

func NewRawBlockEvent(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	migrationHelper migrationhelper.MigrationHelper,
) *RawBlockEvent {
	return &RawBlockEvent{
		rdbprojectionbase.NewRDbBase(
			rdbConn.ToHandle(),
			"RawBlockEvent",
		),

		rdbConn,
		logger,

		migrationHelper,
	}
}

func (_ *RawBlockEvent) GetEventsToListen() []string {
	return []string{
		event_usecase.RAW_BLOCK_EVENT_CREATED,
	}
}

// TODO: should change it to projection folder name to `block_event`, then we can remove it
const (
	MIGRATION_DIRECOTRY = "projection/rawblockevent/migrations"
)

func (projection *RawBlockEvent) OnInit() error {
	if projection.migrationHelper != nil {
		projection.migrationHelper.Migrate()
	}
	return nil
}

func (projection *RawBlockEvent) HandleEvents(height int64, events []event_entity.Event) error {
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
	eventsView := view.NewRawBlockEvents(rdbTxHandle)
	totalView := view.NewRawBlockEventsTotal(rdbTxHandle)

	totalMap := make(map[string]int64)

	eventRows := make([]view.RawBlockEventRow, 0)
	for _, event := range events {
		if rawBlockCreatedEvent, ok := event.(*event_usecase.RawBlockEventCreated); ok {
			fmt.Println("===> final rawBlockCreatedEvent: ", rawBlockCreatedEvent)
			eventRows = append(eventRows, view.RawBlockEventRow{
				BlockHeight: rawBlockCreatedEvent.BlockHeight,
				BlockHash:   rawBlockCreatedEvent.BlockHash,
				BlockTime:   rawBlockCreatedEvent.BlockTime,
				FromResult:  rawBlockCreatedEvent.FromResult,
				RawData: view.RawBlockEventRowData{
					Type:       rawBlockCreatedEvent.RawData.Type,
					Attributes: rawBlockCreatedEvent.RawData.Attributes,
				}})

		} else {
			eventRows = append(eventRows, view.RawBlockEventRow{
				BlockHeight: height,
				BlockHash:   "",
				BlockTime:   utctime.UTCTime{},
				FromResult:  "",
				RawData:     view.RawBlockEventRowData{}})

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
		return fmt.Errorf("error incrementing raw block event type total")
	}
	totalMap["-"] = int64(len(eventRows))
	for key, value := range totalMap {
		if err = totalView.Increment(key, value); err != nil {
			return fmt.Errorf("error incrementing raw block event type total")
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
