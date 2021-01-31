package view

import (
	"errors"
	"fmt"

	"github.com/crypto-com/chain-indexing/internal/json"

	"github.com/crypto-com/chain-indexing/usecase/coin"

	sq "github.com/Masterminds/squirrel"

	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"

	jsoniter "github.com/json-iterator/go"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/internal/utctime"
)

// BlockTransactions projection view implemented by relational database
type AccountTransactions struct {
	rdb *rdb.Handle
}

func NewAccountTransactions(handle *rdb.Handle) *AccountTransactions {
	return &AccountTransactions{
		handle,
	}
}

func (accountMessagesView *AccountTransactions) InsertAll(
	rows []AccountTransactionBaseRow,
) error {
	pendingRowCount := 0
	var stmtBuilder sq.InsertBuilder

	rowCount := len(rows)
	for i, row := range rows {
		if pendingRowCount == 0 {
			stmtBuilder = accountMessagesView.rdb.StmtBuilder.Insert(
				"view_account_transactions",
			).Columns(
				"block_height",
				"block_hash",
				"block_time",
				"account",
				"transaction_hash",
				"success",
				"message_types",
			)
		}

		blockTime := accountMessagesView.rdb.Tton(&row.BlockTime)
		stmtBuilder = stmtBuilder.Values(
			row.BlockHeight,
			row.BlockHash,
			blockTime,
			row.Account,
			row.Hash,
			row.Success,
			json.MustMarshalToString(row.MessageTypes),
		)
		pendingRowCount += 1

		// Postgres has a limit of 65536 parameters.
		if pendingRowCount == 500 || i+1 == rowCount {
			sql, sqlArgs, err := stmtBuilder.ToSql()
			if err != nil {
				return fmt.Errorf("error building account transaction id insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
			}

			result, err := accountMessagesView.rdb.Exec(sql, sqlArgs...)
			if err != nil {
				return fmt.Errorf("error inserting account transactions into the table: %v: %w", err, rdb.ErrWrite)
			}
			if result.RowsAffected() != int64(pendingRowCount) {
				return fmt.Errorf(
					"error inserting account transactions into the table: mismatch rows inserted: %w", rdb.ErrWrite,
				)
			}
			pendingRowCount = 0
		}
	}

	return nil
}

func (accountMessagesView *AccountTransactions) List(
	filter AccountTransactionsListFilter,
	order AccountTransactionsListOrder,
	pagination *pagination_interface.Pagination,
) ([]AccountTransactionReadRow, *pagination_interface.PaginationResult, error) {
	stmtBuilder := accountMessagesView.rdb.StmtBuilder.Select(
		"view_account_transactions.account",
		"view_account_transactions.block_height",
		"view_account_transactions.block_hash",
		"view_account_transactions.block_time",
		"view_account_transactions.transaction_hash",
		"view_account_transactions.success",
		"view_account_transaction_data.code",
		"view_account_transaction_data.log",
		"view_account_transaction_data.fee",
		"view_account_transaction_data.fee_payer",
		"view_account_transaction_data.fee_granter",
		"view_account_transaction_data.gas_wanted",
		"view_account_transaction_data.gas_used",
		"view_account_transaction_data.memo",
		"view_account_transaction_data.timeout_height",
		"view_account_transactions.message_types",
		"view_account_transaction_data.messages",
	).From(
		"view_account_transactions",
	).InnerJoin(
		"view_account_transaction_data ON view_account_transactions.block_height = view_account_transaction_data.block_height AND view_account_transactions.transaction_hash = view_account_transaction_data.hash",
	).Where(
		"view_account_transactions.account = ?", filter.Account,
	)

	if order.Id == view.ORDER_DESC {
		stmtBuilder = stmtBuilder.OrderBy("view_account_transactions.id DESC")
	} else {
		stmtBuilder = stmtBuilder.OrderBy("view_account_transactions.id")
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		accountMessagesView.rdb,
	).WithCustomTotalQueryFn(
		func(rdbHandle *rdb.Handle, _ sq.SelectBuilder) (int64, error) {
			totalView := NewAccountTransactionsTotal(rdbHandle)
			identity := fmt.Sprintf("%s:-", filter.Account)
			total, err := totalView.FindBy(identity)
			if err != nil {
				return int64(0), err
			}
			return total, nil
		},
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf(
			"error building account transactions select SQL: %v, %w", err, rdb.ErrBuildSQLStmt,
		)
	}

	rowsResult, err := accountMessagesView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing account transactions select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	accountMessages := make([]AccountTransactionReadRow, 0)
	for rowsResult.Next() {
		var accountMessage AccountTransactionReadRow
		var feeJSON *string
		var messagesJSON *string
		var messageTypesJSON *string
		blockTimeReader := accountMessagesView.rdb.NtotReader()

		if err = rowsResult.Scan(
			&accountMessage.Account,
			&accountMessage.BlockHeight,
			&accountMessage.BlockHash,
			blockTimeReader.ScannableArg(),
			&accountMessage.Hash,
			&accountMessage.Success,

			&accountMessage.Code,
			&accountMessage.Log,
			&feeJSON,
			&accountMessage.FeePayer,
			&accountMessage.FeeGranter,
			&accountMessage.GasWanted,
			&accountMessage.GasUsed,
			&accountMessage.Memo,
			&accountMessage.TimeoutHeight,
			&messageTypesJSON,
			&messagesJSON,
		); err != nil {
			if errors.Is(err, rdb.ErrNoRows) {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning account transaction row: %v: %w", err, rdb.ErrQuery)
		}
		blockTime, parseErr := blockTimeReader.Parse()
		if parseErr != nil {
			return nil, nil, fmt.Errorf(
				"error parsing account transaction block time: %v: %w", parseErr, rdb.ErrQuery,
			)
		}
		accountMessage.BlockTime = *blockTime

		var fee coin.Coins
		if unmarshalErr := jsoniter.UnmarshalFromString(*feeJSON, &fee); unmarshalErr != nil {
			return nil, nil, fmt.Errorf("error unmarshalling account transaction fee JSON: %v: %w", unmarshalErr, rdb.ErrQuery)
		}
		accountMessage.Fee = fee

		var messageTypes []string
		if unmarshalErr := jsoniter.UnmarshalFromString(*messageTypesJSON, &messageTypes); unmarshalErr != nil {
			return nil, nil, fmt.Errorf("error unmarshalling account transaction message types JSON: %v: %w", unmarshalErr, rdb.ErrQuery)
		}
		accountMessage.MessageTypes = messageTypes

		var messages []TransactionRowMessage
		if unmarshalErr := jsoniter.UnmarshalFromString(*messagesJSON, &messages); unmarshalErr != nil {
			return nil, nil, fmt.Errorf("error unmarshalling account transaction messages JSON: %v: %w", unmarshalErr, rdb.ErrQuery)
		}
		accountMessage.Messages = messages

		accountMessages = append(accountMessages, accountMessage)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return accountMessages, paginationResult, nil
}

type AccountTransactionRecord struct {
	Row      AccountTransactionBaseRow
	Accounts []string
}

type AccountTransactionBaseRow struct {
	Account      string          `json:"account,omitempty"`
	BlockHeight  int64           `json:"blockHeight"`
	BlockHash    string          `json:"blockHash"`
	BlockTime    utctime.UTCTime `json:"blockTime"`
	Hash         string          `json:"hash"`
	MessageTypes []string        `json:"messageTypes"`
	Success      bool            `json:"success"`
}

type AccountTransactionReadRow struct {
	AccountTransactionBaseRow

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

type AccountTransactionsListFilter struct {
	// Required account filter
	Account string
}

type AccountTransactionsListOrder struct {
	Id view.ORDER
}
