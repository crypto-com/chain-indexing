package view

import (
	"errors"
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	_ "github.com/crypto-com/chain-indexing/test/factory"
)

type Account struct {
	AccountAddress string `json:"accountAddress"`
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
			"AccountAddress",
			"AccountBalance",
			"AccountDenom",
		).
		Values("?", "?", "?").
		Suffix("ON CONFLICT(AccountAddress) DO UPDATE SET AccountBalance = EXCLUDED.AccountBalance").
		ToSql()

	if err != nil {
		return fmt.Errorf("error building accounts insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := accountsView.rdb.Exec(sql,
		account.AccountAddress,
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
		"AccountAddress", "AccountBalance", "AccountDenom",
	).From("view_accounts")

	selectStmtBuilder = selectStmtBuilder.Where("AccountAddress = ?", identity.MaybeAddress)

	sql, sqlArgs, err := selectStmtBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building account selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	var account Account
	if err = accountsView.rdb.QueryRow(sql, sqlArgs...).Scan(
		&account.AccountAddress,
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
