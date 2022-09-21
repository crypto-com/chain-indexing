package account_raw_transaction

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper"
	"github.com/crypto-com/chain-indexing/projection/account_raw_transaction/view"
	account_raw_transaction_view "github.com/crypto-com/chain-indexing/projection/account_raw_transaction/view"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ projection_entity.Projection = &AccountRawTransaction{}

var (
	NewAccountRawTransactions      = account_raw_transaction_view.NewAccountRawTransactionsView
	NewAccountRawTransactionsTotal = account_raw_transaction_view.NewAccountRawTransactionsTotalView
	UpdateLastHandledEventHeight   = (*AccountRawTransaction).UpdateLastHandledEventHeight
)

type AccountRawTransaction struct {
	*rdbprojectionbase.Base

	accountAddressPrefix string

	rdbConn rdb.Conn
	logger  applogger.Logger

	migrationHelper migrationhelper.MigrationHelper
}

func NewAccountRawTransaction(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	accountAddressPrefix string,
	migrationHelper migrationhelper.MigrationHelper,
) *AccountRawTransaction {
	return &AccountRawTransaction{
		rdbprojectionbase.NewRDbBase(
			rdbConn.ToHandle(),
			"AccountRawTransaction",
		),

		accountAddressPrefix,

		rdbConn,
		logger,

		migrationHelper,
	}
}

func (_ *AccountRawTransaction) GetEventsToListen() []string {
	return append([]string{
		event_usecase.BLOCK_CREATED,
		event_usecase.RAW_TRANSACTION_CREATED,
		event_usecase.RAW_TRANSACTION_FAILED,
	}, event_usecase.MSG_EVENTS...)
}

func (projection *AccountRawTransaction) OnInit() error {
	if projection.migrationHelper != nil {
		projection.migrationHelper.Migrate()
	}
	return nil
}

func (projection *AccountRawTransaction) HandleEvents(height int64, events []event_entity.Event) error {
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

	accountTransactionsView := NewAccountRawTransactions(rdbTxHandle)
	accountTransactionsTotalView := NewAccountRawTransactionsTotal(rdbTxHandle)

	var blockTime utctime.UTCTime
	var blockHash string

	// Handle and insert a single copy of transaction data
	txs := make([]view.AccountRawTransactionRow, 0)
	for _, event := range events {
		if blockCreatedEvent, ok := event.(*event_usecase.BlockCreated); ok {
			blockTime = blockCreatedEvent.Block.Time
			blockHash = blockCreatedEvent.Block.Hash
		} else if transactionCreatedEvent, ok := event.(*event_usecase.RawTransactionCreated); ok {
			MessageBytes, err := json.Marshal(transactionCreatedEvent.Messages)
			if err != nil {
				return fmt.Errorf("error marshal messages: %v", err)
			}
			addresses := extractAddressesFromMessageJSON(
				string(MessageBytes),
				projection.accountAddressPrefix,
			)
			for i := range addresses {
				tx := view.AccountRawTransactionRow{
					Account:       addresses[i],
					BlockHeight:   height,
					BlockTime:     utctime.UTCTime{}, // placeholder
					Hash:          transactionCreatedEvent.TxHash,
					Success:       true,
					Code:          transactionCreatedEvent.Code,
					Log:           transactionCreatedEvent.Log,
					Fee:           transactionCreatedEvent.Fee,
					FeePayer:      transactionCreatedEvent.FeePayer,
					FeeGranter:    transactionCreatedEvent.FeeGranter,
					GasWanted:     transactionCreatedEvent.GasWanted,
					GasUsed:       transactionCreatedEvent.GasUsed,
					Memo:          transactionCreatedEvent.Memo,
					TimeoutHeight: transactionCreatedEvent.TimeoutHeight,
					Messages:      transactionCreatedEvent.Messages,
				}

				signers := make([]view.AccountRawTransactionRowSigner, 0)
				for _, transactionCreatedEventSigner := range transactionCreatedEvent.Signers {
					transactionRowSigner := view.AccountRawTransactionRowSigner{
						AccountSequence: transactionCreatedEventSigner.AccountSequence,
						Address:         transactionCreatedEventSigner.Address,
					}
					if transactionCreatedEventSigner.MaybeKeyInfo != nil {
						transactionRowSigner.MaybeKeyInfo = &view.AccountRawTransactionRowSignerKeyInfo{
							Type:           transactionCreatedEventSigner.MaybeKeyInfo.Type,
							IsMultiSig:     transactionCreatedEventSigner.MaybeKeyInfo.IsMultiSig,
							Pubkeys:        transactionCreatedEventSigner.MaybeKeyInfo.Pubkeys,
							MaybeThreshold: transactionCreatedEventSigner.MaybeKeyInfo.MaybeThreshold,
						}
					}
					signers = append(signers, transactionRowSigner)
				}
				tx.Signers = signers
				txs = append(txs, tx)
			}
		} else if transactionFailedEvent, ok := event.(*event_usecase.RawTransactionFailed); ok {
			MessageBytes, err := json.Marshal(transactionFailedEvent.Messages)
			if err != nil {
				return fmt.Errorf("error marshal messages: %v", err)
			}
			addresses := extractAddressesFromMessageJSON(
				string(MessageBytes),
				projection.accountAddressPrefix,
			)
			for i := range addresses {
				tx := view.AccountRawTransactionRow{
					Account:       addresses[i],
					BlockHeight:   height,
					BlockTime:     utctime.UTCTime{}, // placeholder
					Hash:          transactionFailedEvent.TxHash,
					Success:       false,
					Code:          transactionFailedEvent.Code,
					Log:           transactionFailedEvent.Log,
					Fee:           transactionFailedEvent.Fee,
					FeePayer:      transactionFailedEvent.FeePayer,
					FeeGranter:    transactionFailedEvent.FeeGranter,
					GasWanted:     transactionFailedEvent.GasWanted,
					GasUsed:       transactionFailedEvent.GasUsed,
					Memo:          transactionFailedEvent.Memo,
					TimeoutHeight: transactionFailedEvent.TimeoutHeight,
					Messages:      transactionFailedEvent.Messages,
				}

				signers := make([]view.AccountRawTransactionRowSigner, 0)
				for _, transactionFailedEventSigner := range transactionFailedEvent.Signers {
					transactionRowSigner := view.AccountRawTransactionRowSigner{
						AccountSequence: transactionFailedEventSigner.AccountSequence,
						Address:         transactionFailedEventSigner.Address,
					}
					if transactionFailedEventSigner.MaybeKeyInfo != nil {
						transactionRowSigner.MaybeKeyInfo = &view.AccountRawTransactionRowSignerKeyInfo{
							Type:           transactionFailedEventSigner.MaybeKeyInfo.Type,
							IsMultiSig:     transactionFailedEventSigner.MaybeKeyInfo.IsMultiSig,
							Pubkeys:        transactionFailedEventSigner.MaybeKeyInfo.Pubkeys,
							MaybeThreshold: transactionFailedEventSigner.MaybeKeyInfo.MaybeThreshold,
						}
					}
					signers = append(signers, transactionRowSigner)
				}
				tx.Signers = signers
				txs = append(txs, tx)
			}
		}
	}

	if len(txs) == 0 {
		if err := UpdateLastHandledEventHeight(projection, rdbTxHandle, height); err != nil {
			return fmt.Errorf("error updating last handled event height: %v", err)
		}

		if err := rdbTx.Commit(); err != nil {
			return fmt.Errorf("error committing changes: %v", err)
		}
		committed = true
		return nil
	}

	for i, tx := range txs {
		txs[i].BlockTime = blockTime
		txs[i].BlockHash = blockHash

		if err := accountTransactionsTotalView.Increment(tx.Account, 1); err != nil {
			return fmt.Errorf("error incrementing total account raw transaction of account: %w", err)
		}
	}

	if err := accountTransactionsView.InsertAll(txs); err != nil {
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

func extractAddressesFromMessageJSON(message string, accountAddressPrefix string) []string {
	splits := strings.Split(message, "\""+accountAddressPrefix)
	addresses := make([]string, 0)
	if len(splits) > 1 {
		for i := 1; i < len(splits); i++ {
			for j := range splits[i] {
				if splits[i][j] == '"' {
					addresses = append(addresses, accountAddressPrefix+splits[i][:j])
					break
				}
			}
		}
	}

	return unique(addresses)
}

func unique(stringSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
