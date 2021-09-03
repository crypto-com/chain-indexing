package view

import (
	"errors"
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type IBCDenomHashMapping struct {
	rdb *rdb.Handle
}

func NewIBCDenomHashMapping(handle *rdb.Handle) *IBCDenomHashMapping {
	return &IBCDenomHashMapping{
		handle,
	}
}

func (ibcDenomHashMappingView *IBCDenomHashMapping) IfDenomExist(denom string) (bool, error) {
	//	SELECT EXISTS (
	//		SELECT * FROM some_table where some_field = $1
	//	)
	sql, sqlArgs, err := ibcDenomHashMappingView.rdb.StmtBuilder.
		Select("*").
		Prefix("SELECT EXISTS (").
		From("view_ibc_denom_hash_mapping").
		Where("denom = ?", denom).
		Suffix(")").
		ToSql()
	if err != nil {
		return false, fmt.Errorf("error building SELECT EXISTS view_ibc_denom_hash_mapping sql: %v: %w", err, rdb.ErrPrepare)
	}

	var exist bool
	if err = ibcDenomHashMappingView.rdb.QueryRow(sql, sqlArgs...).Scan(&exist); err != nil {
		return false, fmt.Errorf("error executing SELECT EXISTS on view_ibc_denom_hash_mapping: %v: %w", err, rdb.ErrQuery)
	}

	return exist, nil
}

func (ibcDenomHashMappingView *IBCDenomHashMapping) Insert(ibcDenomHash *IBCDenomHashMappingRow) error {
	sql, sqlArgs, err := ibcDenomHashMappingView.rdb.StmtBuilder.
		Insert("view_ibc_denom_hash_mapping").
		Columns(
			"denom",
			"hash",
		).
		Values(
			ibcDenomHash.Denom,
			ibcDenomHash.Hash,
		).
		ToSql()

	if err != nil {
		return fmt.Errorf("error building view_ibc_denom_hash_mapping insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := ibcDenomHashMappingView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting view_ibc_denom_hash_mapping into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting view_ibc_denom_hash_mapping into the table: no row inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (ibcDenomHashMappingView *IBCDenomHashMapping) ListAll() (*[]IBCDenomHashMappingRow, error) {
	sql, sqlArgs, err := ibcDenomHashMappingView.rdb.StmtBuilder.Select(
		"denom",
		"hash",
	).
		From("view_ibc_denom_hash_mapping").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building view_ibc_denom_hash_mapping select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := ibcDenomHashMappingView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, fmt.Errorf("error executing view_ibc_denom_hash_mapping select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	denomHashMappings := make([]IBCDenomHashMappingRow, 0)
	for rowsResult.Next() {
		var denomHashRow IBCDenomHashMappingRow

		if err = rowsResult.Scan(
			&denomHashRow.Denom,
			&denomHashRow.Hash,
		); err != nil {
			if errors.Is(err, rdb.ErrNoRows) {
				return nil, rdb.ErrNoRows
			}
			return nil, fmt.Errorf("error scanning view_ibc_denom_hash_mapping row: %v: %w", err, rdb.ErrQuery)
		}

		denomHashMappings = append(denomHashMappings, denomHashRow)
	}

	return &denomHashMappings, nil
}

type IBCDenomHashMappingRow struct {
	Denom string `json:"denom"`
	// The hash should be a sha256 hex encoded string
	Hash string `json:"hash"`
}
