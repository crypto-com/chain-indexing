package view

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	sq "github.com/Masterminds/squirrel"

	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"

	jsoniter "github.com/json-iterator/go"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/internal/utctime"
	"github.com/crypto-com/chain-indexing/usecase/coin"
)

// BlockTransactions projection view implemented by relational database
type BlockTransactions struct {
	rdb *rdb.Handle
}

func NewTransactions(handle *rdb.Handle) *BlockTransactions {
	return &BlockTransactions{
		handle,
	}
}

func (transactionsView *BlockTransactions) InsertAll(transactions []TransactionRow) error {
	pendingRowCount := 0
	var stmtBuilder sq.InsertBuilder

	transactionCount := len(transactions)
	for i, transaction := range transactions {
		if pendingRowCount == 0 {
			stmtBuilder = transactionsView.rdb.StmtBuilder.Insert(
				"view_transactions",
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
		var transactionMessagesJSON string
		var marshalErr error
		if transactionMessagesJSON, marshalErr = jsoniter.MarshalToString(transaction.Messages); marshalErr != nil {
			return fmt.Errorf(
				"error JSON marshalling block transation messages for insertion: %v: %w", marshalErr, rdb.ErrBuildSQLStmt,
			)
		}

		var feeJSON string
		if feeJSON, marshalErr = jsoniter.MarshalToString(transaction.Fee); marshalErr != nil {
			return fmt.Errorf(
				"error JSON marshalling block transation fee for insertion: %v: %w", marshalErr, rdb.ErrBuildSQLStmt,
			)
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
				return fmt.Errorf("error building block transactions insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
			}

			result, err := transactionsView.rdb.Exec(sql, sqlArgs...)
			if err != nil {
				return fmt.Errorf("error inserting block transaction into the table: %v: %w", err, rdb.ErrWrite)
			}
			if result.RowsAffected() != int64(pendingRowCount) {
				return fmt.Errorf("error inserting block transaction into the table: no rows inserted: %w", rdb.ErrWrite)
			}
			pendingRowCount = 0
		}
	}

	return nil
}

func (transactionsView *BlockTransactions) Insert(transaction *TransactionRow) error {
	var err error

	var sql string
	sql, _, err = transactionsView.rdb.StmtBuilder.Insert(
		"view_transactions",
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

func (transactionsView *BlockTransactions) FindByHash(txHash string) (*TransactionRow, error) {
	var err error

	selectStmtBuilder := transactionsView.rdb.StmtBuilder.Select(
		"block_height",
		"block_hash",
		"block_time",
		"hash",
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
	).From(
		"view_transactions",
	).Where(
		"hash = ?", txHash,
	).OrderBy("id DESC")

	sql, sqlArgs, err := selectStmtBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building transactions selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	var transaction TransactionRow
	var feeJSON *string
	var messagesJSON *string
	blockTimeReader := transactionsView.rdb.NtotReader()

	if err = transactionsView.rdb.QueryRow(sql, sqlArgs...).Scan(
		&transaction.BlockHeight,
		&transaction.BlockHash,
		blockTimeReader.ScannableArg(),
		&transaction.Hash,
		&transaction.Success,
		&transaction.Code,
		&transaction.Log,
		&feeJSON,
		&transaction.FeePayer,
		&transaction.FeeGranter,
		&transaction.GasWanted,
		&transaction.GasUsed,
		&transaction.Memo,
		&transaction.TimeoutHeight,
		&messagesJSON,
	); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return nil, rdb.ErrNoRows
		}
		return nil, fmt.Errorf("error scanning transaction row: %v: %w", err, rdb.ErrQuery)
	}
	blockTime, parseErr := blockTimeReader.Parse()
	if parseErr != nil {
		return nil, fmt.Errorf("error parsing transaction block time: %v: %w", parseErr, rdb.ErrQuery)
	}
	transaction.BlockTime = *blockTime

	var fee coin.Coins
	if unmarshalErr := jsoniter.UnmarshalFromString(*feeJSON, &fee); unmarshalErr != nil {
		return nil, fmt.Errorf("error unmarshalling transaction fee JSON: %v: %w", unmarshalErr, rdb.ErrQuery)
	}
	transaction.Fee = fee

	var messages []TransactionRowMessage
	if unmarshalErr := jsoniter.UnmarshalFromString(*messagesJSON, &messages); unmarshalErr != nil {
		return nil, fmt.Errorf("error unmarshalling transaction messages JSON: %v: %w", unmarshalErr, rdb.ErrQuery)
	}
	transaction.Messages = messages

	return &transaction, nil
}

func (transactionsView *BlockTransactions) List(
	filter TransactionsListFilter,
	order TransactionsListOrder,
	pagination *pagination_interface.Pagination,
) ([]TransactionRow, *pagination_interface.PaginationResult, error) {
	stmtBuilder := transactionsView.rdb.StmtBuilder.Select(
		"block_height",
		"block_hash",
		"block_time",
		"hash",
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
	).From(
		"view_transactions",
	)

	if order.Height == view.ORDER_DESC {
		stmtBuilder = stmtBuilder.OrderBy("block_height DESC, id")
	} else {
		stmtBuilder = stmtBuilder.OrderBy("block_height, id")
	}

	if filter.MaybeBlockHeight != nil {
		stmtBuilder = stmtBuilder.Where("block_height = ?", *filter.MaybeBlockHeight)
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		transactionsView.rdb,
	).WithCustomTotalQueryFn(
		func(rdbHandle *rdb.Handle, _ sq.SelectBuilder) (int64, error) {
			identity := "-"
			if filter.MaybeBlockHeight != nil {
				identity = strconv.FormatInt(*filter.MaybeBlockHeight, 10)
			}
			totalView := NewTransactionsTotal(rdbHandle)
			total, err := totalView.FindBy(identity)
			if err != nil {
				return int64(0), err
			}
			return total, nil
		},
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building transactions select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := transactionsView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing transactions select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	transactions := make([]TransactionRow, 0)
	for rowsResult.Next() {
		var transaction TransactionRow
		var feeJSON *string
		var messagesJSON *string
		blockTimeReader := transactionsView.rdb.NtotReader()

		if err = rowsResult.Scan(
			&transaction.BlockHeight,
			&transaction.BlockHash,
			blockTimeReader.ScannableArg(),
			&transaction.Hash,
			&transaction.Success,
			&transaction.Code,
			&transaction.Log,
			&feeJSON,
			&transaction.FeePayer,
			&transaction.FeeGranter,
			&transaction.GasWanted,
			&transaction.GasUsed,
			&transaction.Memo,
			&transaction.TimeoutHeight,
			&messagesJSON,
		); err != nil {
			if errors.Is(err, rdb.ErrNoRows) {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning transaction row: %v: %w", err, rdb.ErrQuery)
		}
		blockTime, parseErr := blockTimeReader.Parse()
		if parseErr != nil {
			return nil, nil, fmt.Errorf("error parsing transaction block time: %v: %w", parseErr, rdb.ErrQuery)
		}
		transaction.BlockTime = *blockTime

		var fee coin.Coins
		if unmarshalErr := jsoniter.UnmarshalFromString(*feeJSON, &fee); unmarshalErr != nil {
			return nil, nil, fmt.Errorf("error unmarshalling transaction fee JSON: %v: %w", unmarshalErr, rdb.ErrQuery)
		}
		transaction.Fee = fee

		var messages []TransactionRowMessage
		if unmarshalErr := jsoniter.UnmarshalFromString(*messagesJSON, &messages); unmarshalErr != nil {
			return nil, nil, fmt.Errorf("error unmarshalling transaction messages JSON: %v: %w", unmarshalErr, rdb.ErrQuery)
		}
		transaction.Messages = messages

		transactions = append(transactions, transaction)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return transactions, paginationResult, nil
}

func (transactionsView *BlockTransactions) Search(keyword string) ([]TransactionRow, error) {
	keyword = strings.ToUpper(keyword)
	sql, sqlArgs, err := transactionsView.rdb.StmtBuilder.Select(
		"block_height",
		"block_hash",
		"block_time",
		"hash",
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
	).From(
		"view_transactions",
	).Where(
		"hash = ?", keyword,
	).OrderBy(
		"block_height",
	).Limit(5).ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building transactions select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := transactionsView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, fmt.Errorf("error executing transactions select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	transactions := make([]TransactionRow, 0)
	for rowsResult.Next() {
		var transaction TransactionRow
		var feeJSON *string
		var messagesJSON *string
		blockTimeReader := transactionsView.rdb.NtotReader()

		if err = rowsResult.Scan(
			&transaction.BlockHeight,
			&transaction.BlockHash,
			blockTimeReader.ScannableArg(),
			&transaction.Hash,
			&transaction.Success,
			&transaction.Code,
			&transaction.Log,
			&feeJSON,
			&transaction.FeePayer,
			&transaction.FeeGranter,
			&transaction.GasWanted,
			&transaction.GasUsed,
			&transaction.Memo,
			&transaction.TimeoutHeight,
			&messagesJSON,
		); err != nil {
			if errors.Is(err, rdb.ErrNoRows) {
				return nil, rdb.ErrNoRows
			}
			return nil, fmt.Errorf("error scanning transaction row: %v: %w", err, rdb.ErrQuery)
		}
		blockTime, parseErr := blockTimeReader.Parse()
		if parseErr != nil {
			return nil, fmt.Errorf("error parsing transaction block time: %v: %w", parseErr, rdb.ErrQuery)
		}
		transaction.BlockTime = *blockTime

		var fee coin.Coins
		if unmarshalErr := jsoniter.UnmarshalFromString(*feeJSON, &fee); unmarshalErr != nil {
			return nil, fmt.Errorf("error unmarshalling transaction fee JSON: %v: %w", unmarshalErr, rdb.ErrQuery)
		}
		transaction.Fee = fee

		var messages []TransactionRowMessage
		if unmarshalErr := jsoniter.UnmarshalFromString(*messagesJSON, &messages); unmarshalErr != nil {
			return nil, fmt.Errorf("error unmarshalling transaction messages JSON: %v: %w", unmarshalErr, rdb.ErrQuery)
		}
		transaction.Messages = messages

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (transactionsView *BlockTransactions) Count() (int64, error) {
	sql, _, err := transactionsView.rdb.StmtBuilder.Select("COUNT(1)").From(
		"view_transactions",
	).ToSql()
	if err != nil {
		return 0, fmt.Errorf("error building transactions count selection sql: %v", err)
	}

	result := transactionsView.rdb.QueryRow(sql)
	var count int64
	if err := result.Scan(&count); err != nil {
		return 0, fmt.Errorf("error scanning transactions count selection query: %v", err)
	}

	return count, nil
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
