package view

import (
	"fmt"

	pagination_interface "github.com/crypto-com/chainindex/appinterface/pagination"

	jsoniter "github.com/json-iterator/go"

	"github.com/crypto-com/chainindex/appinterface/rdb"
	"github.com/crypto-com/chainindex/internal/utctime"
	"github.com/crypto-com/chainindex/usecase/coin"
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

type TransactionsListFilter struct {
	MaybeBlockHeight *int64
}

func (view *BlockTransactions) Insert(transaction *TransactionRow) error {
	var err error

	var sql string
	sql, _, err = view.rdb.StmtBuilder.Insert(
		"view_transactions",
	).Columns(
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
	).Values("?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?").ToSql()
	if err != nil {
		return fmt.Errorf("error building block transactions insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	var transactionMessagesJSON string
	if transactionMessagesJSON, err = jsoniter.MarshalToString(transaction.Messages); err != nil {
		return fmt.Errorf("error JSON marshalling block transation messages for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := view.rdb.Exec(sql,
		transaction.BlockHeight,
		transaction.BlockHash,
		view.rdb.Tton(&transaction.BlockTime),
		transaction.Hash,
		transaction.Success,
		transaction.Code,
		transaction.Log,
		transaction.Fee.String(),
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

func (view *BlockTransactions) FindByHash(txHash string) (*TransactionRow, error) {
	var err error

	selectStmtBuilder := view.rdb.StmtBuilder.Select(
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
	)

	sql, sqlArgs, err := selectStmtBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building transactions selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	var transaction TransactionRow
	var messagesJSON *string
	blockTimeReader := view.rdb.NtotReader()
	var fee string

	if err = view.rdb.QueryRow(sql, sqlArgs...).Scan(
		&transaction.BlockHeight,
		&transaction.BlockHash,
		blockTimeReader.ScannableArg(),
		&transaction.Hash,
		&transaction.Success,
		&transaction.Code,
		&transaction.Log,
		&fee,
		&transaction.FeePayer,
		&transaction.FeeGranter,
		&transaction.GasWanted,
		&transaction.GasUsed,
		&transaction.Memo,
		&transaction.TimeoutHeight,
		&messagesJSON,
	); err != nil {
		if err == rdb.ErrNoRows {
			return nil, rdb.ErrNoRows
		}
		return nil, fmt.Errorf("error scanning transaction row: %v: %w", err, rdb.ErrQuery)
	}
	blockTime, parseErr := blockTimeReader.Parse()
	if parseErr != nil {
		return nil, fmt.Errorf("error parsing transaction block time: %v: %w", parseErr, rdb.ErrQuery)
	}
	transaction.BlockTime = *blockTime
	transaction.Fee = coin.MustNewCoinFromString(fee)

	var messages []TransactionRowMessage
	if unmarshalErr := jsoniter.Unmarshal([]byte(*messagesJSON), &messages); unmarshalErr != nil {
		return nil, fmt.Errorf("error unmarshalling transaction messages JSON: %v: %w", unmarshalErr, rdb.ErrQuery)
	}
	transaction.Messages = messages

	return &transaction, nil
}

func (view *BlockTransactions) List(
	filter TransactionsListFilter,
	pagination *pagination_interface.Pagination,
) ([]TransactionRow, *pagination_interface.PaginationResult, error) {
	stmtBuilder := view.rdb.StmtBuilder.Select(
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
	).OrderBy(
		"height",
	)

	if filter.MaybeBlockHeight != nil {
		stmtBuilder = stmtBuilder.Where("block_height = ?", *filter.MaybeBlockHeight)
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		view.rdb.Runner,
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building transactions select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := view.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing transactions select SQL: %v: %w", err, rdb.ErrQuery)
	}

	transactions := make([]TransactionRow, 0)
	for rowsResult.Next() {
		var transaction TransactionRow
		var messagesJSON *string
		blockTimeReader := view.rdb.NtotReader()
		var fee string

		if err = rowsResult.Scan(
			&transaction.BlockHeight,
			&transaction.BlockHash,
			blockTimeReader.ScannableArg(),
			&transaction.Hash,
			&transaction.Success,
			&transaction.Code,
			&transaction.Log,
			&fee,
			&transaction.FeePayer,
			&transaction.FeeGranter,
			&transaction.GasWanted,
			&transaction.GasUsed,
			&transaction.Memo,
			&transaction.TimeoutHeight,
			&messagesJSON,
		); err != nil {
			if err == rdb.ErrNoRows {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning transaction row: %v: %w", err, rdb.ErrQuery)
		}
		blockTime, parseErr := blockTimeReader.Parse()
		if parseErr != nil {
			return nil, nil, fmt.Errorf("error parsing transaction block time: %v: %w", parseErr, rdb.ErrQuery)
		}
		transaction.BlockTime = *blockTime
		transaction.Fee = coin.MustNewCoinFromString(fee)

		var messages []TransactionRowMessage
		if unmarshalErr := jsoniter.Unmarshal([]byte(*messagesJSON), &messages); unmarshalErr != nil {
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

type TransactionRow struct {
	BlockHeight   int64                   `json:"blockHeight"`
	BlockHash     string                  `json:"blockHash"`
	BlockTime     utctime.UTCTime         `json:"blockTime"`
	Hash          string                  `json:"hash"`
	Success       bool                    `json:"success"`
	Code          int                     `json:"code"`
	Log           string                  `json:"log"`
	Fee           coin.Coin               `json:"fee"`
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
