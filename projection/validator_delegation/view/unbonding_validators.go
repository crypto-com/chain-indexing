package view

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
)

type UnbondingValidators interface {
	Insert(operatorAddress string, unbondingTime utctime.UTCTime) error
	RemoveIfExist(operatorAddress string) error
	GetMatureEntries(blockTime utctime.UTCTime) ([]UnbondingValidatorRow, error)
	DeleteMatureEntries(blockTime utctime.UTCTime) error
}

type UnbondingValidatorsView struct {
	rdb *rdb.Handle
}

func NewUnbondingValidatorsView(handle *rdb.Handle) UnbondingValidators {
	return &UnbondingValidatorsView{
		handle,
	}
}

func (view *UnbondingValidatorsView) Insert(operatorAddress string, unbondingTime utctime.UTCTime) error {

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Insert(
			"view_vd_unbonding_validators",
		).
		Columns(
			"operator_address",
			"unbonding_time",
		).
		Values(
			operatorAddress,
			view.rdb.Tton(&unbondingTime),
		).
		ToSql()

	if err != nil {
		return fmt.Errorf("error building UnbondingValidator insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := view.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting UnbondingValidator into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting UnbondingValidator into the table: rows inserted: %v: %w", result.RowsAffected(), rdb.ErrWrite)
	}

	return nil
}

func (view *UnbondingValidatorsView) RemoveIfExist(operatorAddress string) error {

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Delete(
			"view_vd_unbonding_validators",
		).
		Where(
			"operator_address = ?",
			operatorAddress,
		).
		ToSql()
	if err != nil {
		return fmt.Errorf("error building UnbondingValidator deletion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	// Not checking the number of rows deleted here, as it could be 0 or 1.
	_, err = view.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error deleting UnbondingValidator from the table: %v: %w", err, rdb.ErrWrite)
	}

	return nil
}

func (view *UnbondingValidatorsView) GetMatureEntries(blockTime utctime.UTCTime) ([]UnbondingValidatorRow, error) {

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Select(
			"operator_address",
			"unbonding_time",
		).
		From(
			"view_vd_unbonding_validators",
		).
		Where(
			"unbonding_time <= ?",
			view.rdb.Tton(&blockTime),
		).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building mature UnbondingValidatorRow select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := view.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, fmt.Errorf("error executing mature UnbondingValidatorRow select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	matureRows := make([]UnbondingValidatorRow, 0)
	for rowsResult.Next() {

		var row UnbondingValidatorRow
		unbondingTimeReader := view.rdb.NtotReader()

		if err = rowsResult.Scan(
			&row.OperatorAddress,
			unbondingTimeReader.ScannableArg(),
		); err != nil {
			return nil, fmt.Errorf("error scanning UnbondingValidatorRow: %v: %w", err, rdb.ErrQuery)
		}

		unbondingTime, parseErr := unbondingTimeReader.Parse()
		if parseErr != nil {
			return nil, fmt.Errorf("error parsing unbondingTime: %v: %w", parseErr, rdb.ErrQuery)
		}
		row.UnbondingTime = *unbondingTime

		matureRows = append(matureRows, row)
	}

	return matureRows, nil
}

func (view *UnbondingValidatorsView) DeleteMatureEntries(blockTime utctime.UTCTime) error {

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Delete(
			"view_vd_unbonding_validators",
		).
		Where(
			"unbonding_time <= ?",
			view.rdb.Tton(&blockTime),
		).
		ToSql()
	if err != nil {
		return fmt.Errorf("error building UnbondingValidatorRow deletion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	_, err = view.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error deleting UnbondingValidatorRow from the table: %v: %w", err, rdb.ErrWrite)
	}

	return nil
}

// NOTES:
// - UNIQUE OperatorAddress
type UnbondingValidatorRow struct {
	OperatorAddress string `json:"operatorAddress"`
	// UnbondingTime is the time when Unbonding is finished
	UnbondingTime utctime.UTCTime `json:"UnbondingTime"`
}
