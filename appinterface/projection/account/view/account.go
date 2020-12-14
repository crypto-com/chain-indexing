package view

import (
	"errors"
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"

	"github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	_ "github.com/crypto-com/chain-indexing/test/factory"
)

type Account struct {
	AccountType    string `json:"accountType"`
	AccountAddress string `json:"accountAddress"`
	Pubkey         string `json:"pubkey"`
	AccountNumber  int64  `json:"accountNumber"`
	SequenceNumber int64  `json:"sequenceNumber"`

	AccountBalance int64  `json:"accountBalance"`
	AccountDenom   string `json:"accountDenom"`
}

type Accounts struct {
	rdb *rdb.Handle
}

type AccountsListOrder struct {
	AccountAddress view.ORDER
}

type AccountIdentity struct {
	MaybeAddress string
}

func NewAccounts(handle *rdb.Handle) *Accounts {
	return &Accounts{
		handle,
	}
}

func (accountsView *Accounts) Upsert(account *Account) error {

	var err error

	var sql string
	sql, _, err = accountsView.rdb.StmtBuilder.
		Insert(
			"view_accounts",
		).
		Columns(
			"AccountType",
			"AccountAddress",
			"Pubkey",
			"AccountNumber",
			"SequenceNumber",

			"AccountBalance",
			"AccountDenom",
		).
		Values("?", "?", "?", "?", "?", "?", "?").
		Suffix("ON CONFLICT(AccountAddress) DO UPDATE SET AccountBalance = EXCLUDED.AccountBalance").
		ToSql()

	if err != nil {
		return fmt.Errorf("error building accounts insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := accountsView.rdb.Exec(sql,
		account.AccountType,
		account.AccountAddress,
		account.Pubkey,
		account.AccountNumber,
		account.SequenceNumber,

		account.AccountBalance,
		account.AccountDenom,
	)
	if err != nil {
		return fmt.Errorf("error inserting block into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting block into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (accountsView *Accounts) FindBy(identity *AccountIdentity) (*Account, error) {
	var err error

	selectStmtBuilder := accountsView.rdb.StmtBuilder.Select(
		"AccountType",
		"AccountAddress",
		"Pubkey",
		"AccountNumber",
		"SequenceNumber",

		"AccountBalance", "AccountDenom",
	).From("view_accounts")

	selectStmtBuilder = selectStmtBuilder.Where("AccountAddress = ?", identity.MaybeAddress)

	sql, sqlArgs, err := selectStmtBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building account selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	var account Account
	if err = accountsView.rdb.QueryRow(sql, sqlArgs...).Scan(
		&account.AccountType,
		&account.AccountAddress,
		&account.Pubkey,
		&account.AccountNumber,
		&account.SequenceNumber,

		&account.AccountBalance,
		&account.AccountDenom,
	); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return nil, rdb.ErrNoRows
		}
		return nil, fmt.Errorf("error scanning account row: %v: %w", err, rdb.ErrQuery)
	}
	return &account, nil
}

func (accountsView *Accounts) List(order AccountsListOrder, pagination *pagination.Pagination) ([]Account, *pagination.PaginationResult, error) {
	stmtBuilder := accountsView.rdb.StmtBuilder.Select(

		"AccountType",
		"AccountAddress",
		"Pubkey",
		"AccountNumber",
		"SequenceNumber",

		"AccountBalance",
		"AccountDenom",
	).From(
		"view_accounts",
	)

	if order.AccountAddress == view.ORDER_DESC {
		stmtBuilder = stmtBuilder.OrderBy("AccountAddress DESC")
	} else {
		stmtBuilder = stmtBuilder.OrderBy("AccountAddress")
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		accountsView.rdb,
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building blocks select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := accountsView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing blocks select SQL: %v: %w", err, rdb.ErrQuery)
	}

	accounts := make([]Account, 0)
	for rowsResult.Next() {
		var account Account
		if err = rowsResult.Scan(
			&account.AccountType,
			&account.AccountAddress,
			&account.Pubkey,
			&account.AccountNumber,
			&account.SequenceNumber,

			&account.AccountBalance,
			&account.AccountDenom,
		); err != nil {
			if errors.Is(err, rdb.ErrNoRows) {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning account row: %v: %w", err, rdb.ErrQuery)
		}

		accounts = append(accounts, account)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return accounts, paginationResult, nil
}
