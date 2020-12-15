package transasaction

import (
	"fmt"
	"strconv"

	transaction_view "github.com/crypto-com/chain-indexing/appinterface/projection/transaction/view"

	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	"github.com/crypto-com/chain-indexing/internal/utctime"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ projection_entity.Projection = &Transaction{}

type Transaction struct {
	*rdbprojectionbase.Base

	rdbConn rdb.Conn
	logger  applogger.Logger
}

func NewTransaction(logger applogger.Logger, rdbConn rdb.Conn) *Transaction {
	return &Transaction{
		rdbprojectionbase.NewRDbBase(rdbConn.ToHandle(), "Transaction"),

		rdbConn,
		logger,
	}
}

func (_ *Transaction) GetEventsToListen() []string {
	return append([]string{
		event_usecase.BLOCK_CREATED,
		event_usecase.TRANSACTION_CREATED,
		event_usecase.TRANSACTION_FAILED,
	}, event_usecase.MSG_EVENTS...)
}

func (projection *Transaction) OnInit() error {
	return nil
}

func (projection *Transaction) HandleEvents(height int64, events []event_entity.Event) error {
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
	transactionsView := transaction_view.NewTransactions(rdbTxHandle)
	transactionsTotalView := transaction_view.NewTransactionsTotal(rdbTxHandle)

	var blockTime utctime.UTCTime
	var blockHash string
	txs := make([]transaction_view.TransactionRow, 0)
	txMsgs := make(map[string][]event_usecase.MsgEvent)
	for _, event := range events {
		if blockCreatedEvent, ok := event.(*event_usecase.BlockCreated); ok {
			blockTime = blockCreatedEvent.Block.Time
			blockHash = blockCreatedEvent.Block.Hash
		} else if transactionCreatedEvent, ok := event.(*event_usecase.TransactionCreated); ok {
			txs = append(txs, transaction_view.TransactionRow{
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
				Messages:      make([]transaction_view.TransactionRowMessage, 0),
			})
		} else if transactionFailedEvent, ok := event.(*event_usecase.TransactionFailed); ok {
			txs = append(txs, transaction_view.TransactionRow{
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
				Messages:      make([]transaction_view.TransactionRowMessage, 0),
			})
		} else if msgEvent, ok := event.(event_usecase.MsgEvent); ok {
			if _, exist := txMsgs[msgEvent.TxHash()]; !exist {
				txMsgs[msgEvent.TxHash()] = make([]event_usecase.MsgEvent, 0)
			}
			txMsgs[msgEvent.TxHash()] = append(txMsgs[msgEvent.TxHash()], msgEvent)
		}
	}

	for _, tx := range txs {
		mutTx := tx
		mutTx.BlockTime = blockTime
		mutTx.BlockHash = blockHash

		for _, msg := range txMsgs[mutTx.Hash] {
			mutTx.Messages = append(mutTx.Messages, transaction_view.TransactionRowMessage{
				Type:    msg.MsgType(),
				Content: msg,
			})
		}

		if insertErr := transactionsView.Insert(&mutTx); insertErr != nil {
			return fmt.Errorf("error inserting transaction into view: %v", insertErr)
		}
	}

	totalTxs := int64(len(txs))
	if err := transactionsTotalView.Increment("-", totalTxs); err != nil {
		return fmt.Errorf("error incremnting total transactions: %w", err)
	}
	if err := transactionsTotalView.Set(strconv.FormatInt(height, 10), totalTxs); err != nil {
		return fmt.Errorf("error setting total blcok transactions: %w", err)
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
