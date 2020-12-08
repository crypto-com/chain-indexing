package eventhandler

import (
	"fmt"

	event_interface "github.com/crypto-com/chain-indexing/appinterface/event"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/appinterface/rdbstatusstore"
	"github.com/crypto-com/chain-indexing/entity/event"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
)

var _ Handler = &RDbEventStoreHandler{}

// RDbEventStoreHandler is an event handler which persist the event to event store
type RDbEventStoreHandler struct {
	logger  applogger.Logger
	rdbConn rdb.Conn

	eventStore  *event_interface.RDbStore
	statusStore *rdbstatusstore.RDbStatusStore
}

func NewRDbEventStoreHandler(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	eventRegistry *event.Registry,
) *RDbEventStoreHandler {
	rdbHandle := rdbConn.ToHandle()
	return &RDbEventStoreHandler{
		logger: logger.WithFields(applogger.LogFields{
			"module": "RDbEventStoreHandler",
		}),
		rdbConn: rdbConn,

		eventStore:  initEventStore(rdbHandle, eventRegistry),
		statusStore: initStatusStore(rdbHandle),
	}
}

func (handler *RDbEventStoreHandler) GetLastHandledEventHeight() (*int64, error) {
	return handler.statusStore.GetLastIndexedBlockHeight()
}

func (handler *RDbEventStoreHandler) HandleEvents(blockHeight int64, events []event.Event) error {
	handler.logger.Debug("start persisting blocks events")
	tx, err := handler.rdbConn.Begin()
	if err != nil {
		return fmt.Errorf("error when beginning transaction: %v", err)
	}
	txHandle := tx.ToHandle()

	if err := handler.eventStore.InsertAllWithRDbHandle(txHandle, events); err != nil {
		return fmt.Errorf("error storing all events for height %d: %v", blockHeight, err)
	}

	if err := handler.statusStore.UpdateLastIndexedBlockHeightWithRDbHandle(txHandle, blockHeight); err != nil {
		return fmt.Errorf("error updating last indexed block height to %d: %v", blockHeight, err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing block synchronization outcomes: %v", err)
	}
	return nil
}

func initEventStore(rdbHandle *rdb.Handle, registry *event.Registry) *event_interface.RDbStore {
	return event_interface.NewRDbStore(rdbHandle, registry)
}

func initStatusStore(rdbHandle *rdb.Handle) *rdbstatusstore.RDbStatusStore {
	return rdbstatusstore.NewRDbStatusStore(rdbHandle)
}
