package account_raw_event

import (
	"fmt"
	"strings"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/external/tmcosmosutils"
	"github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper"
	"github.com/crypto-com/chain-indexing/projection/account_raw_event/view"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ projection.Projection = &AccountRawEvent{}

var (
	NewAccountRawEvents          = view.NewAccountRawEventsView
	NewAccountRawEventsTotal     = view.NewAccountRawEventsTotalView
	UpdateLastHandledEventHeight = (*AccountRawEvent).UpdateLastHandledEventHeight
)

type AccountRawEvent struct {
	*rdbprojectionbase.Base

	accountAddressPrefix string

	rdbConn rdb.Conn
	logger  applogger.Logger

	migrationHelper migrationhelper.MigrationHelper
}

func NewAccountRawEvent(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	accountAddressPrefix string,
	migrationHelper migrationhelper.MigrationHelper,
) *AccountRawEvent {
	return &AccountRawEvent{
		rdbprojectionbase.NewRDbBase(
			rdbConn.ToHandle(),
			"AccountRawEvent",
		),

		accountAddressPrefix,

		rdbConn,
		logger,

		migrationHelper,
	}
}

func (_ *AccountRawEvent) GetEventsToListen() []string {
	return []string{
		event_usecase.BLOCK_RAW_EVENT_CREATED,
	}
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

	// TODO: Handle genesis transaction
	if height == int64(0) {
		if err := projection.UpdateLastHandledEventHeight(rdbTxHandle, height); err != nil {
			return fmt.Errorf("error updating last handled event height: %v", err)
		}

		if err := rdbTx.Commit(); err != nil {
			return fmt.Errorf("error committing changes: %v", err)
		}
		committed = true
		return nil
	}

	eventsView := NewAccountRawEvents(rdbTxHandle)
	totalView := NewAccountRawEventsTotal(rdbTxHandle)

	eventRows := make([]view.AccountRawEventRow, 0)

	accountToAccountRawEventRowMap := make(map[string]view.AccountRawEventRow)
	for _, event := range events {
		if rawBlockCreatedEvent, ok := event.(*event_usecase.BlockRawEventCreated); ok {
			accounts := []string{}
			// append related account
			for _, attribute := range rawBlockCreatedEvent.Data.Content.Attributes {
				hasPrefix := strings.HasPrefix(attribute.Value, projection.accountAddressPrefix)
				IsValidCosmosAddress := tmcosmosutils.IsValidCosmosAddress(attribute.Value)

				if hasPrefix && IsValidCosmosAddress {
					accounts = append(accounts, attribute.Value)
				}
			}

			// create event for every account
			for _, account := range accounts {
				if accountEvents, ok := accountToAccountRawEventRowMap[account]; ok {
					accountEvents.Data = append(accountEvents.Data, view.AccountRawEventRowData{
						Type:    rawBlockCreatedEvent.Data.Content.Type,
						Content: rawBlockCreatedEvent.Data.Content,
					})
				} else {
					accountToAccountRawEventRowMap[account] = view.AccountRawEventRow{
						BlockHeight: height,
						BlockHash:   rawBlockCreatedEvent.BlockHash,
						BlockTime:   rawBlockCreatedEvent.BlockTime,
						Account:     account,
						Data: []view.AccountRawEventRowData{
							{
								Type:    rawBlockCreatedEvent.Data.Content.Type,
								Content: rawBlockCreatedEvent.Data.Content,
							},
						},
					}
				}
			}
		}
	}

	for account := range accountToAccountRawEventRowMap {
		eventRows = append(eventRows, accountToAccountRawEventRowMap[account])
	}

	for _, eventRow := range eventRows {
		// Calculate account raw event total
		if err := totalView.Increment(eventRow.Account, 1); err != nil {
			return fmt.Errorf("error incrementing total account transaction of account: %w", err)
		}
	}

	if err := eventsView.InsertAll(eventRows); err != nil {
		return fmt.Errorf("error inserting account message: %w", err)
	}

	if err := UpdateLastHandledEventHeight(projection, rdbTxHandle, height); err != nil {
		return fmt.Errorf("error updating last handled event height: %v", err)
	}

	if err := rdbTx.Commit(); err != nil {
		return fmt.Errorf("error committing changes: %v", err)
	}
	committed = true
	return nil
}
