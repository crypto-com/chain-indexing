package projection

import (
	"fmt"

	"github.com/crypto-com/chainindex/internal/utctime"

	"github.com/crypto-com/chainindex/appinterface/projection/rdbbase"
	"github.com/crypto-com/chainindex/appinterface/projection/view"
	"github.com/crypto-com/chainindex/appinterface/rdb"
	event_entity "github.com/crypto-com/chainindex/entity/event"
	applogger "github.com/crypto-com/chainindex/internal/logger"
	event_usecase "github.com/crypto-com/chainindex/usecase/event"
)

type BlockTransaction struct {
	*rdbbase.RDbBase

	rdbConn rdb.Conn
	logger  applogger.Logger
}

func NewTransaction(logger applogger.Logger, rdbConn rdb.Conn) *BlockTransaction {
	return &BlockTransaction{
		rdbbase.NewRDbBase(rdbConn.ToHandle(), "Transaction"),

		rdbConn,
		logger,
	}
}

func (_ *BlockTransaction) GetEventsToListen() []string {
	return append([]string{
		event_usecase.BLOCK_CREATED,
		event_usecase.TRANSACTION_CREATED,
		event_usecase.TRANSACTION_FAILED,
	}, event_usecase.MSG_EVENTS...)
}

func (projection *BlockTransaction) OnInit() error {
	return nil
}

func (projection *BlockTransaction) HandleEvents(height int64, events []event_entity.Event) error {
	var err error

	var rdbTx rdb.Tx
	rdbTx, err = projection.rdbConn.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction: %v", err)
	}
	rdbTxHandle := rdbTx.ToHandle()
	transactionsView := view.NewTransactions(rdbTxHandle)

	var blockTime utctime.UTCTime
	var blockHash string
	txs := make([]view.Transaction, 0)
	txMsgs := make(map[string][]event_usecase.MsgEvent)
	for _, event := range events {
		if blockCreatedEvent, ok := event.(*event_usecase.BlockCreated); ok {
			blockTime = blockCreatedEvent.Block.Time
			blockHash = blockCreatedEvent.Block.Hash
		} else if transactionCreatedEvent, ok := event.(*event_usecase.TransactionCreated); ok {
			txs = append(txs, view.Transaction{
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
				Messages:      make([]view.TransactionMessage, 0),
			})
		} else if transactionFailedEvent, ok := event.(*event_usecase.TransactionFailed); ok {
			txs = append(txs, view.Transaction{
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
				Messages:      make([]view.TransactionMessage, 0),
			})
		} else if msgEvent, ok := event.(event_usecase.MsgEvent); ok {
			if _, exist := txMsgs[msgEvent.TxHash()]; !exist {
				txMsgs[msgEvent.TxHash()] = make([]event_usecase.MsgEvent, 0)
			}
			txMsgs[msgEvent.TxHash()] = append(txMsgs[msgEvent.TxHash()], msgEvent)
		}
	}

	for i, tx := range txs {
		tx.BlockTime = blockTime
		tx.BlockHash = blockHash

		for _, msg := range txMsgs[tx.Hash] {
			tx.Messages = append(tx.Messages, view.TransactionMessage{
				Type:    msg.MsgType(),
				Content: msg,
			})
		}

		if insertErr := transactionsView.Insert(&txs[i]); insertErr != nil {
			_ = rdbTx.Rollback()
			return fmt.Errorf("error inserting parsed transaction into view: %v", err)
		}
	}

	if err = projection.UpdateLastHandledEventHeight(rdbTxHandle, height); err != nil {
		_ = rdbTx.Rollback()
		return fmt.Errorf("error updating last handled event height: %v", err)
	}

	if err = rdbTx.Commit(); err != nil {
		_ = rdbTx.Rollback()
		return fmt.Errorf("error committing changes: %v", err)
	}
	return nil
}
