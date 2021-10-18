package view

import (
	"errors"
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type IBCClients interface {
	Insert(*IBCClientRow) error
	FindCounterpartyChainIDBy(string) (string, error)
}

type IBCClientsView struct {
	rdb *rdb.Handle
}

func NewIBCClientsView(handle *rdb.Handle) IBCClients {
	return &IBCClientsView{
		handle,
	}
}

func (ibcClientsView *IBCClientsView) Insert(ibcClient *IBCClientRow) error {
	sql, sqlArgs, err := ibcClientsView.rdb.StmtBuilder.
		Insert("view_ibc_clients").
		Columns(
			"client_id",
			"counterparty_chain_id",
		).
		Values(
			ibcClient.ClientID,
			ibcClient.CounterpartyChainID,
		).
		ToSql()

	if err != nil {
		return fmt.Errorf("error building ibc_client insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := ibcClientsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting ibc_client into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting ibc_client into the table: no row inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (ibcClientsView *IBCClientsView) FindCounterpartyChainIDBy(clientID string) (string, error) {
	sql, sqlArgs, err := ibcClientsView.rdb.StmtBuilder.
		Select("counterparty_chain_id").
		From("view_ibc_clients").
		Where("client_id = ?", clientID).
		ToSql()
	if err != nil {
		return "", fmt.Errorf("error building selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	var chainID string
	if err = ibcClientsView.rdb.QueryRow(sql, sqlArgs...).Scan(&chainID); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return "", rdb.ErrNoRows
		}
		return "", fmt.Errorf("error scanning ibc_client row: %v: %w", err, rdb.ErrQuery)
	}

	return chainID, nil
}

type IBCClientRow struct {
	ClientID            string `json:"clientId"`
	CounterpartyChainID string `json:"counterpartyChainId"`
}
