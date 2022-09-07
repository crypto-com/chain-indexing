package raw_transaction

import (
	"fmt"
	"strconv"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper"
	raw_transaction_view "github.com/crypto-com/chain-indexing/projection/raw_transaction/view"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ projection_entity.Projection = &RawTransaction{}

var (
	NewRawTransactions           = raw_transaction_view.NewRawTransactionsView
	NewRawTransactionsTotal      = raw_transaction_view.NewRawTransactionsTotalView
	UpdateLastHandledEventHeight = (*RawTransaction).UpdateLastHandledEventHeight
)

type RawTransaction struct {
	*rdbprojectionbase.Base

	rdbConn rdb.Conn
	logger  applogger.Logger

	migrationHelper migrationhelper.MigrationHelper
}

func NewRawTransaction(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	migrationHelper migrationhelper.MigrationHelper,
) *RawTransaction {
	return &RawTransaction{
		rdbprojectionbase.NewRDbBase(
			rdbConn.ToHandle(),
			"RawTransaction",
		),

		rdbConn,
		logger,
		migrationHelper,
	}
}

func (_ *RawTransaction) GetEventsToListen() []string {
	return append([]string{
		event_usecase.BLOCK_CREATED,
		event_usecase.RAW_TRANSACTION_CREATED,
		event_usecase.RAW_TRANSACTION_FAILED,
	})
}

func (projection *RawTransaction) OnInit() error {
	if projection.migrationHelper != nil {
		projection.migrationHelper.Migrate()
	}

	return nil
}

func (projection *RawTransaction) HandleEvents(height int64, events []event_entity.Event) error {
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
	rawTransactionsView := NewRawTransactions(rdbTxHandle)
	rawTransactionsTotalView := NewRawTransactionsTotal(rdbTxHandle)

	var blockTime utctime.UTCTime
	var blockHash string
	txs := make([]raw_transaction_view.RawTransactionRow, 0)
	for _, event := range events {
		if blockCreatedEvent, ok := event.(*event_usecase.BlockCreated); ok {
			blockTime = blockCreatedEvent.Block.Time
			blockHash = blockCreatedEvent.Block.Hash
		} else if transactionCreatedEvent, ok := event.(*event_usecase.RawTransactionCreated); ok {
			tx := raw_transaction_view.RawTransactionRow{
				BlockHeight:   height,
				BlockTime:     utctime.UTCTime{}, // placeholder
				Hash:          transactionCreatedEvent.TxHash,
				Index:         transactionCreatedEvent.Index,
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

			signers := make([]raw_transaction_view.RawTransactionRowSigner, 0)
			for _, transactionCreatedEventSigner := range transactionCreatedEvent.Signers {
				transactionRowSigner := raw_transaction_view.RawTransactionRowSigner{
					AccountSequence: transactionCreatedEventSigner.AccountSequence,
					Address:         transactionCreatedEventSigner.Address,
				}
				if transactionCreatedEventSigner.MaybeKeyInfo != nil {
					transactionRowSigner.MaybeKeyInfo = &raw_transaction_view.RawTransactionRowSignerKeyInfo{
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
		} else if transactionFailedEvent, ok := event.(*event_usecase.RawTransactionFailed); ok {
			tx := raw_transaction_view.RawTransactionRow{
				BlockHeight:   height,
				BlockTime:     utctime.UTCTime{}, // placeholder
				Hash:          transactionFailedEvent.TxHash,
				Index:         transactionFailedEvent.Index,
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

			signers := make([]raw_transaction_view.RawTransactionRowSigner, 0)
			for _, transactionFailedEventSigner := range transactionFailedEvent.Signers {
				transactionRowSigner := raw_transaction_view.RawTransactionRowSigner{
					AccountSequence: transactionFailedEventSigner.AccountSequence,
					Address:         transactionFailedEventSigner.Address,
				}
				if transactionFailedEventSigner.MaybeKeyInfo != nil {
					transactionRowSigner.MaybeKeyInfo = &raw_transaction_view.RawTransactionRowSignerKeyInfo{
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

	if len(txs) == 0 {
		if err := projection.UpdateLastHandledEventHeight(rdbTxHandle, height); err != nil {
			return fmt.Errorf("error updating last handled event height: %v", err)
		}

		if err := rdbTx.Commit(); err != nil {
			return fmt.Errorf("error committing changes: %v", err)
		}
		committed = true
		return nil
	}

	for i := range txs {
		txs[i].BlockTime = blockTime
		txs[i].BlockHash = blockHash
	}

	if insertErr := rawTransactionsView.InsertAll(txs); insertErr != nil {
		return fmt.Errorf("error inserting transaction into view: %v", insertErr)
	}

	totalTxs := int64(len(txs))
	if err := rawTransactionsTotalView.Increment("-", totalTxs); err != nil {
		return fmt.Errorf("error incremnting total transactions: %w", err)
	}
	if err := rawTransactionsTotalView.Set(strconv.FormatInt(height, 10), totalTxs); err != nil {
		return fmt.Errorf("error setting total blcok transactions: %w", err)
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
