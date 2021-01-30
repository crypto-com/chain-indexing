package view

import (
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"

	jsoniter "github.com/json-iterator/go"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/internal/utctime"
)

// BlockTransactions projection view implemented by relational database
type AccountMessages struct {
	rdb *rdb.Handle
}

func NewAccountMessages(handle *rdb.Handle) *AccountMessages {
	return &AccountMessages{
		handle,
	}
}

func (accountMessagesView *AccountMessages) Insert(messageRow *AccountMessageRow, accounts []string) error {
	accountMessageDataJSON, err := jsoniter.MarshalToString(messageRow.Data)
	if err != nil {
		return fmt.Errorf("error JSON marshalling account message for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	stmtBuilder := accountMessagesView.rdb.StmtBuilder.Insert(
		"view_account_messages",
	).Columns(
		"block_height",
		"block_hash",
		"block_time",
		"account",
		"transaction_hash",
		"success",
		"message_index",
		"message_type",
		"data",
	)
	blockTime := accountMessagesView.rdb.Tton(&messageRow.BlockTime)
	for _, account := range accounts {
		stmtBuilder = stmtBuilder.Values(
			messageRow.BlockHeight,
			messageRow.BlockHash,
			blockTime,
			account,
			messageRow.TransactionHash,
			messageRow.Success,
			messageRow.MessageIndex,
			messageRow.MessageType,
			accountMessageDataJSON,
		)
	}
	sql, sqlArgs, err := stmtBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("error building account message id insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}
	result, err := accountMessagesView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting account messages into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != int64(len(accounts)) {
		return fmt.Errorf(
			"error inserting account messages into the table: mismatched rows inserted: %w", rdb.ErrWrite,
		)
	}

	return nil
}

func (accountMessagesView *AccountMessages) List(
	filter AccountMessagesListFilter,
	order AccountMessagesListOrder,
	pagination *pagination_interface.Pagination,
) ([]AccountMessageRow, *pagination_interface.PaginationResult, error) {
	stmtBuilder := accountMessagesView.rdb.StmtBuilder.Select(
		"view_account_messages.account",
		"view_account_messages.block_height",
		"view_account_messages.block_hash",
		"view_account_messages.block_time",
		"view_account_messages.transaction_hash",
		"view_account_messages.success",
		"view_account_messages.message_index",
		"view_account_messages.message_type",
		"view_account_messages.data",
	).From(
		"view_account_messages",
	).Where(
		"view_account_messages.account = ?", filter.Account,
	)

	var totalIdentities []string
	if filter.MaybeMsgTypes == nil {
		totalIdentities = []string{fmt.Sprintf("%s:-", filter.Account)}
	} else {
		totalIdentities = make([]string, 0)
		stmtBuilder = stmtBuilder.Where(sq.Eq{"view_account_messages.message_type": filter.MaybeMsgTypes})
		for _, msgType := range filter.MaybeMsgTypes {
			totalIdentities = append(totalIdentities, fmt.Sprintf("%s:%s", filter.Account, msgType))
		}
	}

	if order.Id == view.ORDER_DESC {
		stmtBuilder = stmtBuilder.OrderBy("id DESC")
	} else {
		stmtBuilder = stmtBuilder.OrderBy("id")
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		accountMessagesView.rdb,
	).WithCustomTotalQueryFn(
		func(rdbHandle *rdb.Handle, _ sq.SelectBuilder) (int64, error) {
			totalView := NewAccountMessagesTotal(rdbHandle)
			total, err := totalView.SumBy(totalIdentities)
			if err != nil {
				return int64(0), err
			}
			return total, nil
		},
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf(
			"error building account messages select SQL: %v, %w", err, rdb.ErrBuildSQLStmt,
		)
	}

	rowsResult, err := accountMessagesView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing account messages select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	accountMessages := make([]AccountMessageRow, 0)
	for rowsResult.Next() {
		var accountMessage AccountMessageRow
		var accountMessageDataJSON *string
		blockTimeReader := accountMessagesView.rdb.NtotReader()

		if err = rowsResult.Scan(
			&accountMessage.MaybeAccount,
			&accountMessage.BlockHeight,
			&accountMessage.BlockHash,
			blockTimeReader.ScannableArg(),
			&accountMessage.TransactionHash,
			&accountMessage.Success,
			&accountMessage.MessageIndex,
			&accountMessage.MessageType,
			&accountMessageDataJSON,
		); err != nil {
			if errors.Is(err, rdb.ErrNoRows) {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning account message row: %v: %w", err, rdb.ErrQuery)
		}
		blockTime, parseErr := blockTimeReader.Parse()
		if parseErr != nil {
			return nil, nil, fmt.Errorf(
				"error parsing account message block time: %v: %w", parseErr, rdb.ErrQuery,
			)
		}
		accountMessage.BlockTime = *blockTime

		var data interface{}
		if unmarshalErr := jsoniter.Unmarshal([]byte(*accountMessageDataJSON), &data); unmarshalErr != nil {
			return nil, nil, fmt.Errorf(
				"error unmarshalling account message data JSON: %v: %w", unmarshalErr, rdb.ErrQuery,
			)
		}
		accountMessage.Data = data

		accountMessages = append(accountMessages, accountMessage)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return accountMessages, paginationResult, nil
}

type AccountMessageRecord struct {
	Row      AccountMessageRow
	Accounts []string
}

type AccountMessageRow struct {
	MaybeAccount    *string         `json:"account,omitempty"`
	BlockHeight     int64           `json:"blockHeight"`
	BlockHash       string          `json:"blockHash"`
	BlockTime       utctime.UTCTime `json:"blockTime"`
	TransactionHash string          `json:"transactionHash"`
	Success         bool            `json:"success"`
	MessageIndex    int             `json:"messageIndex"`
	MessageType     string          `json:"messageType"`
	Data            interface{}     `json:"data"`
}

type AccountMessagesListFilter struct {
	// Required account filter
	Account string

	// Optional filtering
	MaybeMsgTypes []string
}

type AccountMessagesListOrder struct {
	Id view.ORDER
}
