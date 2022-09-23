package account_raw_event

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/external/tmcosmosutils"
	"github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper"
	"github.com/crypto-com/chain-indexing/projection/account_raw_event/view"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
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
	eventsView := view.NewAccountRawEventsView(rdbTxHandle)
	totalView := view.NewAccountRawEventsTotalView(rdbTxHandle)

	eventRows := make([]view.AccountRawEventRow, 0)

	for _, event := range events {
		if rawBlockCreatedEvent, ok := event.(*event_usecase.BlockRawEventCreated); ok {
			accountRawEvent := utils.NewParsedTxsResultLogEvent(&rawBlockCreatedEvent.Data.Content)
			if rawBlockCreatedEvent.Data.Type == "coin_spent" {
				eventRows = append(eventRows, view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   rawBlockCreatedEvent.BlockHash,
					BlockTime:   rawBlockCreatedEvent.BlockTime,
					Account:     accountRawEvent.MustGetAttributeByKey("spender"),
					Data: view.AccountRawEventRowData{
						Type:    rawBlockCreatedEvent.Data.Content.Type,
						Content: rawBlockCreatedEvent.Data.Content,
					},
				})
			} else if rawBlockCreatedEvent.Data.Type == "coin_received" {
				eventRows = append(eventRows, view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   rawBlockCreatedEvent.BlockHash,
					BlockTime:   rawBlockCreatedEvent.BlockTime,
					Account:     accountRawEvent.MustGetAttributeByKey("receiver"),
					Data: view.AccountRawEventRowData{
						Type:    rawBlockCreatedEvent.Data.Content.Type,
						Content: rawBlockCreatedEvent.Data.Content,
					},
				})

			} else if rawBlockCreatedEvent.Data.Type == "transfer" {
				eventRows = append(eventRows, view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   rawBlockCreatedEvent.BlockHash,
					BlockTime:   rawBlockCreatedEvent.BlockTime,
					Account:     accountRawEvent.MustGetAttributeByKey("recipient"),
					Data: view.AccountRawEventRowData{
						Type:    rawBlockCreatedEvent.Data.Content.Type,
						Content: rawBlockCreatedEvent.Data.Content,
					},
				})
				eventRows = append(eventRows, view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   rawBlockCreatedEvent.BlockHash,
					BlockTime:   rawBlockCreatedEvent.BlockTime,
					Account:     accountRawEvent.MustGetAttributeByKey("sender"),
					Data: view.AccountRawEventRowData{
						Type:    rawBlockCreatedEvent.Data.Content.Type,
						Content: rawBlockCreatedEvent.Data.Content,
					},
				})

			} else if rawBlockCreatedEvent.Data.Type == "coinbase" {
				eventRows = append(eventRows, view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   rawBlockCreatedEvent.BlockHash,
					BlockTime:   rawBlockCreatedEvent.BlockTime,
					Account:     accountRawEvent.MustGetAttributeByKey("minter"),
					Data: view.AccountRawEventRowData{
						Type:    rawBlockCreatedEvent.Data.Content.Type,
						Content: rawBlockCreatedEvent.Data.Content,
					},
				})

			} else if rawBlockCreatedEvent.Data.Type == "burn" {
				eventRows = append(eventRows, view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   rawBlockCreatedEvent.BlockHash,
					BlockTime:   rawBlockCreatedEvent.BlockTime,
					Account:     accountRawEvent.MustGetAttributeByKey("burner"),
					Data: view.AccountRawEventRowData{
						Type:    rawBlockCreatedEvent.Data.Content.Type,
						Content: rawBlockCreatedEvent.Data.Content,
					},
				})

			} else if rawBlockCreatedEvent.Data.Type == "proposer_reward" {
				accountAddress := tmcosmosutils.MustAccountAddressFromValidatorAddress(
					projection.accountAddressPrefix,
					accountRawEvent.MustGetAttributeByKey("validator"),
				)

				eventRows = append(eventRows, view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   rawBlockCreatedEvent.BlockHash,
					BlockTime:   rawBlockCreatedEvent.BlockTime,
					Account:     accountAddress,
					Data: view.AccountRawEventRowData{
						Type:    rawBlockCreatedEvent.Data.Content.Type,
						Content: rawBlockCreatedEvent.Data.Content,
					},
				})

			} else if rawBlockCreatedEvent.Data.Type == "rewards" {
				accountAddress := tmcosmosutils.MustAccountAddressFromValidatorAddress(
					projection.accountAddressPrefix,
					accountRawEvent.MustGetAttributeByKey("validator"),
				)

				eventRows = append(eventRows, view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   rawBlockCreatedEvent.BlockHash,
					BlockTime:   rawBlockCreatedEvent.BlockTime,
					Account:     accountAddress,
					Data: view.AccountRawEventRowData{
						Type:    rawBlockCreatedEvent.Data.Content.Type,
						Content: rawBlockCreatedEvent.Data.Content,
					},
				})

			} else if rawBlockCreatedEvent.Data.Type == "commission" {
				accountAddress := tmcosmosutils.MustAccountAddressFromValidatorAddress(
					projection.accountAddressPrefix,
					accountRawEvent.MustGetAttributeByKey("validator"),
				)

				eventRows = append(eventRows, view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   rawBlockCreatedEvent.BlockHash,
					BlockTime:   rawBlockCreatedEvent.BlockTime,
					Account:     accountAddress,
					Data: view.AccountRawEventRowData{
						Type:    rawBlockCreatedEvent.Data.Content.Type,
						Content: rawBlockCreatedEvent.Data.Content,
					},
				})

			} else if rawBlockCreatedEvent.Data.Type == "slash" {
				if accountRawEvent.HasAttribute("address") {
					eventRows = append(eventRows, view.AccountRawEventRow{
						BlockHeight: height,
						BlockHash:   rawBlockCreatedEvent.BlockHash,
						BlockTime:   rawBlockCreatedEvent.BlockTime,
						Account:     accountRawEvent.MustGetAttributeByKey("address"),
						Data: view.AccountRawEventRowData{
							Type:    rawBlockCreatedEvent.Data.Content.Type,
							Content: rawBlockCreatedEvent.Data.Content,
						},
					})
				}

			} else if rawBlockCreatedEvent.Data.Type == "complete_unbonding" {
				eventRows = append(eventRows, view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   rawBlockCreatedEvent.BlockHash,
					BlockTime:   rawBlockCreatedEvent.BlockTime,
					Account:     accountRawEvent.MustGetAttributeByKey("delegator"),
					Data: view.AccountRawEventRowData{
						Type:    rawBlockCreatedEvent.Data.Content.Type,
						Content: rawBlockCreatedEvent.Data.Content,
					},
				})

				accountAddress := tmcosmosutils.MustAccountAddressFromValidatorAddress(
					projection.accountAddressPrefix,
					accountRawEvent.MustGetAttributeByKey("validator"),
				)
				eventRows = append(eventRows, view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   rawBlockCreatedEvent.BlockHash,
					BlockTime:   rawBlockCreatedEvent.BlockTime,
					Account:     accountAddress,
					Data: view.AccountRawEventRowData{
						Type:    rawBlockCreatedEvent.Data.Content.Type,
						Content: rawBlockCreatedEvent.Data.Content,
					},
				})

			} else if rawBlockCreatedEvent.Data.Type == "ethereum_send_to_cosmos_handled" {
				eventRows = append(eventRows, view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   rawBlockCreatedEvent.BlockHash,
					BlockTime:   rawBlockCreatedEvent.BlockTime,
					Account:     accountRawEvent.MustGetAttributeByKey("sender"),
					Data: view.AccountRawEventRowData{
						Type:    rawBlockCreatedEvent.Data.Content.Type,
						Content: rawBlockCreatedEvent.Data.Content,
					},
				})
				eventRows = append(eventRows, view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   rawBlockCreatedEvent.BlockHash,
					BlockTime:   rawBlockCreatedEvent.BlockTime,
					Account:     accountRawEvent.MustGetAttributeByKey("receiver"),
					Data: view.AccountRawEventRowData{
						Type:    rawBlockCreatedEvent.Data.Content.Type,
						Content: rawBlockCreatedEvent.Data.Content,
					},
				})

			} else if rawBlockCreatedEvent.Data.Type == "message" {
				if accountRawEvent.HasAttribute("sender") {
					eventRows = append(eventRows, view.AccountRawEventRow{
						BlockHeight: height,
						BlockHash:   rawBlockCreatedEvent.BlockHash,
						BlockTime:   rawBlockCreatedEvent.BlockTime,
						Account:     accountRawEvent.MustGetAttributeByKey("sender"),
						Data: view.AccountRawEventRowData{
							Type:    rawBlockCreatedEvent.Data.Content.Type,
							Content: rawBlockCreatedEvent.Data.Content,
						},
					})
				} else if accountRawEvent.HasAttribute("granter") {
					eventRows = append(eventRows, view.AccountRawEventRow{
						BlockHeight: height,
						BlockHash:   rawBlockCreatedEvent.BlockHash,
						BlockTime:   rawBlockCreatedEvent.BlockTime,
						Account:     accountRawEvent.MustGetAttributeByKey("granter"),
						Data: view.AccountRawEventRowData{
							Type:    rawBlockCreatedEvent.Data.Content.Type,
							Content: rawBlockCreatedEvent.Data.Content,
						},
					})
				} else if accountRawEvent.HasAttribute("grantee") {
					eventRows = append(eventRows, view.AccountRawEventRow{
						BlockHeight: height,
						BlockHash:   rawBlockCreatedEvent.BlockHash,
						BlockTime:   rawBlockCreatedEvent.BlockTime,
						Account:     accountRawEvent.MustGetAttributeByKey("grantee"),
						Data: view.AccountRawEventRowData{
							Type:    rawBlockCreatedEvent.Data.Content.Type,
							Content: rawBlockCreatedEvent.Data.Content,
						},
					})
				}

			} else if rawBlockCreatedEvent.Data.Type == "withdraw_rewards" {
				accountAddress := tmcosmosutils.MustAccountAddressFromValidatorAddress(
					projection.accountAddressPrefix,
					accountRawEvent.MustGetAttributeByKey("validator"),
				)

				eventRows = append(eventRows, view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   rawBlockCreatedEvent.BlockHash,
					BlockTime:   rawBlockCreatedEvent.BlockTime,
					Account:     accountAddress,
					Data: view.AccountRawEventRowData{
						Type:    rawBlockCreatedEvent.Data.Content.Type,
						Content: rawBlockCreatedEvent.Data.Content,
					},
				})

			} else if rawBlockCreatedEvent.Data.Type == "set_withdraw_address" {
				eventRows = append(eventRows, view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   rawBlockCreatedEvent.BlockHash,
					BlockTime:   rawBlockCreatedEvent.BlockTime,
					Account:     accountRawEvent.MustGetAttributeByKey("withdraw_address"),
					Data: view.AccountRawEventRowData{
						Type:    rawBlockCreatedEvent.Data.Content.Type,
						Content: rawBlockCreatedEvent.Data.Content,
					},
				})

			} else if rawBlockCreatedEvent.Data.Type == "complete_redelegation" {
				sourceValidator := tmcosmosutils.MustAccountAddressFromValidatorAddress(
					projection.accountAddressPrefix,
					accountRawEvent.MustGetAttributeByKey("source_validator"),
				)

				destinationValidator := tmcosmosutils.MustAccountAddressFromValidatorAddress(
					projection.accountAddressPrefix,
					accountRawEvent.MustGetAttributeByKey("destination_validator"),
				)

				eventRows = append(eventRows, view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   rawBlockCreatedEvent.BlockHash,
					BlockTime:   rawBlockCreatedEvent.BlockTime,
					Account:     sourceValidator,
					Data: view.AccountRawEventRowData{
						Type:    rawBlockCreatedEvent.Data.Content.Type,
						Content: rawBlockCreatedEvent.Data.Content,
					},
				})
				eventRows = append(eventRows, view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   rawBlockCreatedEvent.BlockHash,
					BlockTime:   rawBlockCreatedEvent.BlockTime,
					Account:     destinationValidator,
					Data: view.AccountRawEventRowData{
						Type:    rawBlockCreatedEvent.Data.Content.Type,
						Content: rawBlockCreatedEvent.Data.Content,
					},
				})
				eventRows = append(eventRows, view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   rawBlockCreatedEvent.BlockHash,
					BlockTime:   rawBlockCreatedEvent.BlockTime,
					Account:     accountRawEvent.MustGetAttributeByKey("delegator"),
					Data: view.AccountRawEventRowData{
						Type:    rawBlockCreatedEvent.Data.Content.Type,
						Content: rawBlockCreatedEvent.Data.Content,
					},
				})

			} else if rawBlockCreatedEvent.Data.Type == "create_validator" {
				accountAddress := tmcosmosutils.MustAccountAddressFromValidatorAddress(
					projection.accountAddressPrefix,
					accountRawEvent.MustGetAttributeByKey("validator"),
				)

				eventRows = append(eventRows, view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   rawBlockCreatedEvent.BlockHash,
					BlockTime:   rawBlockCreatedEvent.BlockTime,
					Account:     accountAddress,
					Data: view.AccountRawEventRowData{
						Type:    rawBlockCreatedEvent.Data.Content.Type,
						Content: rawBlockCreatedEvent.Data.Content,
					},
				})

			} else if rawBlockCreatedEvent.Data.Type == "delegate" {
				accountAddress := tmcosmosutils.MustAccountAddressFromValidatorAddress(
					projection.accountAddressPrefix,
					accountRawEvent.MustGetAttributeByKey("validator"),
				)

				eventRows = append(eventRows, view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   rawBlockCreatedEvent.BlockHash,
					BlockTime:   rawBlockCreatedEvent.BlockTime,
					Account:     accountAddress,
					Data: view.AccountRawEventRowData{
						Type:    rawBlockCreatedEvent.Data.Content.Type,
						Content: rawBlockCreatedEvent.Data.Content,
					},
				})

			} else if rawBlockCreatedEvent.Data.Type == "unbond" {
				accountAddress := tmcosmosutils.MustAccountAddressFromValidatorAddress(
					projection.accountAddressPrefix,
					accountRawEvent.MustGetAttributeByKey("destination_validator"),
				)

				eventRows = append(eventRows, view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   rawBlockCreatedEvent.BlockHash,
					BlockTime:   rawBlockCreatedEvent.BlockTime,
					Account:     accountAddress,
					Data: view.AccountRawEventRowData{
						Type:    rawBlockCreatedEvent.Data.Content.Type,
						Content: rawBlockCreatedEvent.Data.Content,
					},
				})

			} else if rawBlockCreatedEvent.Data.Type == "redelegate" {
				sourceValidator := tmcosmosutils.MustAccountAddressFromValidatorAddress(
					projection.accountAddressPrefix,
					accountRawEvent.MustGetAttributeByKey("source_validator"),
				)

				destinationValidator := tmcosmosutils.MustAccountAddressFromValidatorAddress(
					projection.accountAddressPrefix,
					accountRawEvent.MustGetAttributeByKey("destination_validator"),
				)

				eventRows = append(eventRows, view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   rawBlockCreatedEvent.BlockHash,
					BlockTime:   rawBlockCreatedEvent.BlockTime,
					Account:     sourceValidator,
					Data: view.AccountRawEventRowData{
						Type:    rawBlockCreatedEvent.Data.Content.Type,
						Content: rawBlockCreatedEvent.Data.Content,
					},
				})
				eventRows = append(eventRows, view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   rawBlockCreatedEvent.BlockHash,
					BlockTime:   rawBlockCreatedEvent.BlockTime,
					Account:     destinationValidator,
					Data: view.AccountRawEventRowData{
						Type:    rawBlockCreatedEvent.Data.Content.Type,
						Content: rawBlockCreatedEvent.Data.Content,
					},
				})

			} else if rawBlockCreatedEvent.Data.Type == "ibc_transfer" {
				eventRows = append(eventRows, view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   rawBlockCreatedEvent.BlockHash,
					BlockTime:   rawBlockCreatedEvent.BlockTime,
					Account:     accountRawEvent.MustGetAttributeByKey("sender"),
					Data: view.AccountRawEventRowData{
						Type:    rawBlockCreatedEvent.Data.Content.Type,
						Content: rawBlockCreatedEvent.Data.Content,
					},
				})
				eventRows = append(eventRows, view.AccountRawEventRow{
					BlockHeight: height,
					BlockHash:   rawBlockCreatedEvent.BlockHash,
					BlockTime:   rawBlockCreatedEvent.BlockTime,
					Account:     accountRawEvent.MustGetAttributeByKey("receiver"),
					Data: view.AccountRawEventRowData{
						Type:    rawBlockCreatedEvent.Data.Content.Type,
						Content: rawBlockCreatedEvent.Data.Content,
					},
				})
			}

		}
	}

	for _, eventRow := range eventRows {
		// Calculate account raw event total
		if err := totalView.Increment(fmt.Sprintf("%s:-", eventRow.Account), 1); err != nil {
			return fmt.Errorf("error incrementing total account transaction of account: %w", err)
		}

	}

	if err := eventsView.InsertAll(eventRows); err != nil {
		return fmt.Errorf("error inserting account message: %w", err)
	}

	if err := projection.UpdateLastHandledEventHeight(rdbTxHandle, height); err != nil {
		return fmt.Errorf("error updating last handled event height: %v", err)
	}

	if err := rdbTx.Commit(); err != nil {
		return fmt.Errorf("error committing changes: %v", err)
	}
	committed = true
	return nil
}
