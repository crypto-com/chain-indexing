package example

import (
	"fmt"
	"log"

	example_view "github.com/crypto-com/chain-indexing/example/projection/example/view"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

type AdditionalExampleProjection struct {
	*rdbprojectionbase.Base

	rdbConn rdb.Conn
	logger  applogger.Logger

	migrationHelper migrationhelper.MigrationHelper
}

func NewAdditionalProjection(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	migrationHelper migrationhelper.MigrationHelper,
) *AdditionalExampleProjection {
	return &AdditionalExampleProjection{
		rdbprojectionbase.NewRDbBase(rdbConn.ToHandle(), "Example"),

		rdbConn,
		logger,

		migrationHelper,
	}
}

var (
	NewExamplesView              = example_view.NewExamplesView
	UpdateLastHandledEventHeight = (*AdditionalExampleProjection).UpdateLastHandledEventHeight
)

func (_ *AdditionalExampleProjection) GetEventsToListen() []string {
	return event_usecase.MSG_EVENTS
}

func (projection *AdditionalExampleProjection) OnInit() error {
	if projection.migrationHelper != nil {
		projection.migrationHelper.Migrate()
	}

	return nil
}

func (projection *AdditionalExampleProjection) HandleEvents(height int64, events []event_entity.Event) error {
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

	examplesView := NewExamplesView(rdbTxHandle)

	for _, event := range events {
		log.Println(event)
		if typedEvent, ok := event.(*event_usecase.MsgSend); ok {
			row := &example_view.ExampleRow{
				Address: typedEvent.ToAddress,
				Balance: typedEvent.Amount,
			}
			if handleErr := projection.handleSomeEvent(examplesView, row); handleErr != nil {
				return fmt.Errorf("error handling MsgSend: %v", handleErr)
			}
		}
	}

	if err = UpdateLastHandledEventHeight(projection, rdbTxHandle, height); err != nil {
		return fmt.Errorf("error updating last handled event height: %v", err)
	}

	if err = rdbTx.Commit(); err != nil {
		return fmt.Errorf("error committing changes: %v", err)
	}
	committed = true

	return nil
}

func (projection *AdditionalExampleProjection) handleSomeEvent(examplesView example_view.Examples, row *example_view.ExampleRow) error {
	return examplesView.Insert(row)
}
