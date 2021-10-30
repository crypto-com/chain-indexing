package view

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"

	jsoniter "github.com/json-iterator/go"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/usecase/coin"
)

// AccountTransactionData projection view implemented by relational database
type AccountTransactionData struct {
	rdb *rdb.Handle
}

func NewAccountTransactionData(handle *rdb.Handle) *AccountTransactionData {
	return &AccountTransactionData{
		handle,
	}
}

func (transactionsView *AccountTransactionData) InsertAll(transactions []TransactionRow) error {
	pendingRowCount := 0
	var stmtBuilder sq.InsertBuilder

	transactionCount := len(transactions)
	for i, transaction := range transactions {
		if pendingRowCount == 0 {
			stmtBuilder = transactionsView.rdb.StmtBuilder.Insert(
				"view_account_transaction_data",
			).Columns(
				"block_height",
				"block_hash",
				"block_time",
				"hash",
				"index",
				"success",
				"code",
				"log",
				"fee",
				"fee_payer",
				"fee_granter",
				"gas_wanted",
				"gas_used",
				"memo",
				"timeout_height",
				"messages",
			)
		}
		transactionMessagesJSON, err := jsoniter.MarshalToString(transaction.Messages)
		if err != nil {
			return fmt.Errorf("error JSON marshalling block transation messages for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
		}

		feeJSON, err := jsoniter.MarshalToString(transaction.Fee)
		if err != nil {
			return fmt.Errorf("error JSON marshalling block transation fee for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
		}

		stmtBuilder = stmtBuilder.Values(
			transaction.BlockHeight,
			transaction.BlockHash,
			transactionsView.rdb.Tton(&transaction.BlockTime),
			transaction.Hash,
			transaction.Index,
			transaction.Success,
			transaction.Code,
			transaction.Log,
			feeJSON,
			transaction.FeePayer,
			transaction.FeeGranter,
			transaction.GasWanted,
			transaction.GasUsed,
			transaction.Memo,
			transaction.TimeoutHeight,
			transactionMessagesJSON,
		)
		pendingRowCount += 1

		// Postgres has a limit of 65536 parameters.
		if pendingRowCount == 500 || i+1 == transactionCount {
			sql, sqlArgs, err := stmtBuilder.ToSql()
			if err != nil {
				return fmt.Errorf("error building account transactions data insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
			}
			result, err := transactionsView.rdb.Exec(sql, sqlArgs...)
			if err != nil {
				return fmt.Errorf("error inserting block transaction into the table: %v: %w", err, rdb.ErrWrite)
			}
			if result.RowsAffected() != int64(pendingRowCount) {
				return fmt.Errorf("error inserting account transaction data into the table: mismatch rows inserted: %w", rdb.ErrWrite)
			}
			pendingRowCount = 0
		}
	}

	return nil
}

func (transactionsView *AccountTransactionData) Insert(transaction *TransactionRow) error {
	var err error

	var sql string
	sql, _, err = transactionsView.rdb.StmtBuilder.Insert(
		"view_account_transaction_data",
	).Columns(
		"block_height",
		"block_hash",
		"block_time",
		"hash",
		"index",
		"success",
		"code",
		"log",
		"fee",
		"fee_payer",
		"fee_granter",
		"gas_wanted",
		"gas_used",
		"memo",
		"timeout_height",
		"messages",
	).Values("?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?").ToSql()
	if err != nil {
		return fmt.Errorf("error building block transactions insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	var transactionMessagesJSON string
	if transactionMessagesJSON, err = jsoniter.MarshalToString(transaction.Messages); err != nil {
		return fmt.Errorf("error JSON marshalling block transation messages for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	var feeJSON string
	if feeJSON, err = jsoniter.MarshalToString(transaction.Fee); err != nil {
		return fmt.Errorf("error JSON marshalling block transation fee for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := transactionsView.rdb.Exec(sql,
		transaction.BlockHeight,
		transaction.BlockHash,
		transactionsView.rdb.Tton(&transaction.BlockTime),
		transaction.Hash,
		transaction.Index,
		transaction.Success,
		transaction.Code,
		transaction.Log,
		feeJSON,
		transaction.FeePayer,
		transaction.FeeGranter,
		transaction.GasWanted,
		transaction.GasUsed,
		transaction.Memo,
		transaction.TimeoutHeight,
		transactionMessagesJSON,
	)
	if err != nil {
		return fmt.Errorf("error inserting block transaction into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting block transaction into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}

type TransactionRow struct {
	BlockHeight   int64                   `json:"blockHeight"`
	BlockHash     string                  `json:"blockHash"`
	BlockTime     utctime.UTCTime         `json:"blockTime"`
	Hash          string                  `json:"hash"`
	Index         int                     `json:"index"`
	Success       bool                    `json:"success"`
	Code          int                     `json:"code"`
	Log           string                  `json:"log"`
	Fee           coin.Coins              `json:"fee"`
	FeePayer      string                  `json:"feePayer"`
	FeeGranter    string                  `json:"feeGranter"`
	GasWanted     int                     `json:"gasWanted"`
	GasUsed       int                     `json:"gasUsed"`
	Memo          string                  `json:"memo"`
	TimeoutHeight int64                   `json:"timeoutHeight"`
	Messages      []TransactionRowMessage `json:"messages"`
}

type TransactionRowMessage struct {
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
}

type TransactionsListFilter struct {
	MaybeBlockHeight *int64
}

type TransactionsListOrder struct {
	Height view.ORDER
}
