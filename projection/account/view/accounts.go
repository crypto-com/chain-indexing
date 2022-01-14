package view

import (
	"errors"
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/external/json"
	jsoniter "github.com/json-iterator/go"

	"github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	_ "github.com/crypto-com/chain-indexing/test/factory"
)

type Accounts interface {
	FindBy(*AccountIdentity) (*AccountRow, error)
	List(AccountsListOrder, *pagination.Pagination) ([]AccountRow, *pagination.PaginationResult, error)
	IncrementUsableBalance(string, string, int64) error
	DecrementUsableBalance(string, string, int64) error
	InsertAccountEvent(AccountEvent) error
}

type AccountsView struct {
	rdb *rdb.Handle
}

type AccountsListOrder struct {
	AccountAddress view.ORDER
}

type AccountIdentity struct {
	Address string
}

func NewAccountsView(handle *rdb.Handle) Accounts {
	return &AccountsView{
		handle,
	}
}

func (accountsView *AccountsView) InsertAccountEvent(accountEvent AccountEvent) error {
	var accountEventDataJSON string
	var err error
	if accountEventDataJSON, err = jsoniter.MarshalToString(accountEvent.Data); err != nil {
		return fmt.Errorf("error JSON marshalling block event data for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	sql, sqlArgs, err := accountsView.rdb.StmtBuilder.
		Insert(
			"view_account_events",
		).
		Columns(
			"address",
			"data",
			"type",
			"block_height",
		).
		Values(
			accountEvent.Address,
			accountEventDataJSON,
			accountEvent.Type,
			accountEvent.BlockHeight,
		).
		ToSql()

	if err != nil {
		return fmt.Errorf("error building account events insert sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := accountsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting account events into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting account events into the table: no rows updated: %w", rdb.ErrWrite)
	}

	return nil
}

func (accountsView *AccountsView) IncrementUsableBalance(address string, denom string, increment int64) error {
	initUsableBalance := make(map[string]int64)
	initUsableBalance[denom] = increment

	var initUsableBalanceJSON string
	var err error
	if initUsableBalanceJSON, err = jsoniter.MarshalToString(initUsableBalance); err != nil {
		return fmt.Errorf("error JSON marshalling block event data for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	sql, sqlArgs, err := accountsView.rdb.StmtBuilder.
		Insert(
			"view_accounts",
		).
		Columns(
			"address",
			"usable_balance",
		).
		Values(
			address,
			initUsableBalanceJSON,
		).
		Suffix(fmt.Sprintf("ON CONFLICT(address) DO UPDATE SET usable_balance = jsonb_set(view_accounts.usable_balance, '{%s}', (COALESCE(view_accounts.usable_balance->>'%s','0')::bigint + %d)::text::jsonb)", denom, denom, increment)).
		ToSql()

	if err != nil {
		return fmt.Errorf("error building accounts update sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := accountsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error updating account into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error updating account into the table: no rows updated: %w", rdb.ErrWrite)
	}

	return nil
}

func (accountsView *AccountsView) DecrementUsableBalance(address string, denom string, decrement int64) error {
	initUsableBalance := make(map[string]int64)
	initUsableBalance[denom] = -decrement

	var initUsableBalanceJSON string
	var err error
	if initUsableBalanceJSON, err = jsoniter.MarshalToString(initUsableBalance); err != nil {
		return fmt.Errorf("error JSON marshalling block event data for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	sql, sqlArgs, err := accountsView.rdb.StmtBuilder.
		Insert(
			"view_accounts",
		).
		Columns(
			"address",
			"usable_balance",
		).
		Values(
			address,
			initUsableBalanceJSON,
		).
		Suffix(fmt.Sprintf("ON CONFLICT(address) DO UPDATE SET usable_balance = jsonb_set(view_accounts.usable_balance, '{%s}', (COALESCE(view_accounts.usable_balance->>'%s','0')::bigint - %d)::text::jsonb)", denom, denom, decrement)).
		ToSql()

	if err != nil {
		return fmt.Errorf("error building accounts insert sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := accountsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting account into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error updating account into the table: no rows updated: %w", rdb.ErrWrite)
	}

	return nil
}

func (accountsView *AccountsView) FindBy(identity *AccountIdentity) (*AccountRow, error) {
	var err error

	selectStmtBuilder := accountsView.rdb.StmtBuilder.Select(
		"address",
		"account_number",
		"sequence_number",
		"usable_balance",
	).From("view_accounts")

	selectStmtBuilder = selectStmtBuilder.Where("address = ?", identity.Address)

	sql, sqlArgs, err := selectStmtBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building account selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	var account AccountRow
	var usableBalance string
	if err = accountsView.rdb.QueryRow(sql, sqlArgs...).Scan(
		&account.Address,
		&account.AccountNumber,
		&account.SequenceNumber,
		&usableBalance,
	); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return nil, rdb.ErrNoRows
		}
		return nil, fmt.Errorf("error scanning account row: %v: %w", err, rdb.ErrQuery)
	}

	json.MustUnmarshalFromString(usableBalance, &account.UsableBalance)
	return &account, nil
}

func (accountsView *AccountsView) List(
	order AccountsListOrder,
	pagination *pagination.Pagination,
) ([]AccountRow, *pagination.PaginationResult, error) {
	stmtBuilder := accountsView.rdb.StmtBuilder.Select(
		"address",
		"account_number",
		"sequence_number",
		"usable_balance",
	).From(
		"view_accounts",
	)

	if order.AccountAddress == view.ORDER_DESC {
		stmtBuilder = stmtBuilder.OrderBy("address DESC")
	} else {
		stmtBuilder = stmtBuilder.OrderBy("address")
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		accountsView.rdb,
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building accounts select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := accountsView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing accounts select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	accounts := make([]AccountRow, 0)
	for rowsResult.Next() {
		var account AccountRow
		var usableBalance string
		if err = rowsResult.Scan(
			&account.Address,
			&account.AccountNumber,
			&account.SequenceNumber,
			&usableBalance,
		); err != nil {
			if errors.Is(err, rdb.ErrNoRows) {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning account row: %v: %w", err, rdb.ErrQuery)
		}

		json.MustUnmarshalFromString(usableBalance, &account.UsableBalance)
		accounts = append(accounts, account)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return accounts, paginationResult, nil
}

type AccountRow struct {
	Address        string         `json:"address"`
	AccountNumber  string         `json:"accountNumber"`
	SequenceNumber string         `json:"sequenceNumber"`
	UsableBalance  AccountBalance `json:"usableBalance"`
}

type AccountBalanceDenom string

type AccountBalance map[AccountBalanceDenom]int64

type AccountEvent struct {
	Address     string      `json:"address"`
	Type        string      `json:"type"`
	Data        interface{} `json:"data"`
	BlockHeight int64       `json:"block_height"`
}
