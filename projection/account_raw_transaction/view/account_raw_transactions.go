package view

import (
	"errors"
	"fmt"

	"github.com/crypto-com/chain-indexing/external/json"
	"github.com/crypto-com/chain-indexing/usecase/coin"

	sq "github.com/Masterminds/squirrel"

	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"

	jsoniter "github.com/json-iterator/go"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
)

type AccountRawTransactions interface {
	InsertAll(rows []AccountRawTransactionRow) error
	List(
		filter AccountRawTransactionsListFilter,
		order AccountRawTransactionsListOrder,
		pagination *pagination_interface.Pagination,
	) ([]AccountRawTransactionRow, *pagination_interface.PaginationResult, error)
}

// AccountRawTransactionsView projection view implemented by relational database
type AccountRawTransactionsView struct {
	rdb *rdb.Handle
}

func NewAccountRawTransactionsView(handle *rdb.Handle) AccountRawTransactions {
	return &AccountRawTransactionsView{
		handle,
	}
}

func (accountMessagesView *AccountRawTransactionsView) InsertAll(
	rows []AccountRawTransactionRow,
) error {
	pendingRowCount := 0
	var stmtBuilder sq.InsertBuilder

	rowCount := len(rows)
	for i, row := range rows {
		if pendingRowCount == 0 {
			stmtBuilder = accountMessagesView.rdb.StmtBuilder.Insert(
				"view_account_raw_transactions",
			).Columns(
				"block_height",
				"block_hash",
				"block_time",
				"account",
				"transaction_hash",
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
				"signers",
			)
		}

		var transactionMessagesJSON string
		var marshalErr error
		if transactionMessagesJSON, marshalErr = json.MarshalToString(row.Messages); marshalErr != nil {
			return fmt.Errorf(
				"error JSON marshalling block transation messages for insertion: %v: %w", marshalErr, rdb.ErrBuildSQLStmt,
			)
		}

		var feeJSON string
		if feeJSON, marshalErr = json.MarshalToString(row.Fee); marshalErr != nil {
			return fmt.Errorf(
				"error JSON marshalling block transation fee for insertion: %v: %w", marshalErr, rdb.ErrBuildSQLStmt,
			)
		}

		var signersJSON string
		if signersJSON, marshalErr = json.MarshalToString(row.Signers); marshalErr != nil {
			return fmt.Errorf(
				"error JSON marshalling block transation signers for insertion: %v: %w", marshalErr, rdb.ErrBuildSQLStmt,
			)
		}

		stmtBuilder = stmtBuilder.Values(
			row.BlockHeight,
			row.BlockHash,
			accountMessagesView.rdb.Tton(&row.BlockTime),
			row.Account,
			row.Hash,
			row.Success,
			row.Code,
			row.Log,
			feeJSON,
			row.FeePayer,
			row.FeeGranter,
			row.GasWanted,
			row.GasUsed,
			row.Memo,
			row.TimeoutHeight,
			transactionMessagesJSON,
			signersJSON,
		)
		pendingRowCount += 1

		// Postgres has a limit of 65536 parameters.
		if pendingRowCount == 500 || i+1 == rowCount {
			sql, sqlArgs, err := stmtBuilder.ToSql()
			if err != nil {
				return fmt.Errorf("error building account raw transactions insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
			}

			result, err := accountMessagesView.rdb.Exec(sql, sqlArgs...)
			if err != nil {
				return fmt.Errorf("error inserting account raw transaction into the table: %v: %w", err, rdb.ErrWrite)
			}
			if result.RowsAffected() != int64(pendingRowCount) {
				return fmt.Errorf("error inserting account raw transaction into the table: no rows inserted: %w", rdb.ErrWrite)
			}
			pendingRowCount = 0
		}
	}

	return nil
}

func (accountMessagesView *AccountRawTransactionsView) List(
	filter AccountRawTransactionsListFilter,
	order AccountRawTransactionsListOrder,
	pagination *pagination_interface.Pagination,
) ([]AccountRawTransactionRow, *pagination_interface.PaginationResult, error) {
	stmtBuilder := accountMessagesView.rdb.StmtBuilder.Select(
		"account",
		"block_height",
		"block_hash",
		"block_time",
		"transaction_hash",
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
		"signers",
	).From(
		"view_account_raw_transactions",
	)

	if filter.Memo != "" {
		stmtBuilder = stmtBuilder.Where(
			"view_account_raw_transactions.account = ? AND view_account_raw_transactions.memo = ?", filter.Account, filter.Memo,
		)
	} else {
		stmtBuilder = stmtBuilder.Where(
			"view_account_raw_transactions.account = ?", filter.Account,
		)
	}

	if order.Id == view.ORDER_DESC {
		stmtBuilder = stmtBuilder.OrderBy("view_account_raw_transactions.id DESC")
	} else {
		stmtBuilder = stmtBuilder.OrderBy("view_account_raw_transactions.id")
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		accountMessagesView.rdb,
	).WithCustomTotalQueryFn(
		func(rdbHandle *rdb.Handle, _ sq.SelectBuilder) (int64, error) {
			totalView := NewAccountRawTransactionsTotalView(rdbHandle)

			identity := ""
			if filter.Memo != "" {
				identity = fmt.Sprintf("%s/%s:-", filter.Account, filter.Memo)
			} else {
				identity = fmt.Sprintf("%s:-", filter.Account)
			}

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
			"error building account raw transactions select SQL: %v, %w", err, rdb.ErrBuildSQLStmt,
		)
	}

	rowsResult, err := accountMessagesView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing account raw transactions select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	accountMessages := make([]AccountRawTransactionRow, 0)
	for rowsResult.Next() {
		var accountMessage AccountRawTransactionRow
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
			return nil, nil, fmt.Errorf("error scanning account raw transaction row: %v: %w", err, rdb.ErrQuery)
		}
		blockTime, parseErr := blockTimeReader.Parse()
		if parseErr != nil {
			return nil, nil, fmt.Errorf(
				"error parsing account raw transaction block time: %v: %w", parseErr, rdb.ErrQuery,
			)
		}
		accountMessage.BlockTime = *blockTime

		var fee coin.Coins
		if unmarshalErr := jsoniter.UnmarshalFromString(*feeJSON, &fee); unmarshalErr != nil {
			return nil, nil, fmt.Errorf("error unmarshalling account raw transaction fee JSON: %v: %w", unmarshalErr, rdb.ErrQuery)
		}
		accountMessage.Fee = fee

		var messages []map[string]interface{}
		if unmarshalErr := jsoniter.UnmarshalFromString(*messagesJSON, &messages); unmarshalErr != nil {
			return nil, nil, fmt.Errorf("error unmarshalling account raw transaction messages JSON: %v: %w", unmarshalErr, rdb.ErrQuery)
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

type AccountRawTransactionRow struct {
	Account       string                           `json:"account,omitempty"`
	BlockHeight   int64                            `json:"blockHeight"`
	BlockHash     string                           `json:"blockHash"`
	BlockTime     utctime.UTCTime                  `json:"blockTime"`
	Hash          string                           `json:"hash"`
	Success       bool                             `json:"success"`
	Code          int                              `json:"code"`
	Log           string                           `json:"log"`
	Fee           coin.Coins                       `json:"fee"`
	FeePayer      string                           `json:"feePayer"`
	FeeGranter    string                           `json:"feeGranter"`
	GasWanted     int                              `json:"gasWanted"`
	GasUsed       int                              `json:"gasUsed"`
	Memo          string                           `json:"memo"`
	TimeoutHeight int64                            `json:"timeoutHeight"`
	Messages      []map[string]interface{}         `json:"messages"`
	Signers       []AccountRawTransactionRowSigner `json:"signers"`
}

type AccountRawTransactionsListFilter struct {
	// Required account filter
	Account string
	// Optional memo filter
	Memo string
}

type AccountRawTransactionsListOrder struct {
	Id view.ORDER
}

type AccountRawTransactionRowSigner struct {
	MaybeKeyInfo *AccountRawTransactionRowSignerKeyInfo `json:"keyInfo"`

	Address         string `json:"address"`
	AccountSequence uint64 `json:"accountSequence"`
}

type AccountRawTransactionRowSignerKeyInfo struct {
	Type           string   `json:"type"`
	IsMultiSig     bool     `json:"isMultiSig"`
	Pubkeys        []string `json:"pubkeys"`
	MaybeThreshold *int     `json:"threshold,omitempty"`
}
