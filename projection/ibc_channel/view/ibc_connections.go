package view

import (
	"errors"
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type IBCConnections interface {
	Insert(*IBCConnectionRow) error
	FindCounterpartyChainIDBy(string) (string, error)
}

type IBCConnectionsView struct {
	rdb *rdb.Handle
}

func NewIBCConnectionsView(handle *rdb.Handle) IBCConnections {
	return &IBCConnectionsView{
		handle,
	}
}

func (ibcConnectionsView *IBCConnectionsView) Insert(ibcConnection *IBCConnectionRow) error {
	sql, sqlArgs, err := ibcConnectionsView.rdb.StmtBuilder.
		Insert("view_ibc_connections").
		Columns(
			"connection_id",
			"client_id",
			"counterparty_connection_id",
			"counterparty_client_id",
			"counterparty_chain_id",
		).
		Values(
			ibcConnection.ConnectionID,
			ibcConnection.ClientID,
			ibcConnection.CounterpartyConnectionID,
			ibcConnection.CounterpartyClientID,
			ibcConnection.CounterpartyChainID,
		).
		ToSql()

	if err != nil {
		return fmt.Errorf("error building ibc_connection insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := ibcConnectionsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting ibc_connection into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting ibc_connection into the table: no row inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (ibcConnectionsView *IBCConnectionsView) FindCounterpartyChainIDBy(connectionID string) (string, error) {
	sql, sqlArgs, err := ibcConnectionsView.rdb.StmtBuilder.
		Select("counterparty_chain_id").
		From("view_ibc_connections").
		Where("connection_id = ?", connectionID).
		ToSql()
	if err != nil {
		return "", fmt.Errorf("error building selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	var counterpartyChainID string
	if err = ibcConnectionsView.rdb.QueryRow(sql, sqlArgs...).Scan(&counterpartyChainID); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return "", rdb.ErrNoRows
		}
		return "", fmt.Errorf("error scanning ibc_connection row: %v: %w", err, rdb.ErrQuery)
	}

	return counterpartyChainID, nil
}

type IBCConnectionRow struct {
	ConnectionID             string `json:"connectionId"`
	ClientID                 string `json:"clientId"`
	CounterpartyConnectionID string `json:"counterpartyConnectionId"`
	CounterpartyClientID     string `json:"counterpartyClientId"`
	CounterpartyChainID      string `json:"counterpartyChainId"`
}
