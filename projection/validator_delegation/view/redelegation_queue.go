package view

import (
	"errors"
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
	jsoniter "github.com/json-iterator/go"
)

type RedelegationQueue interface {
	Upsert(row RedelegationQueueRow) error
	FindBy(completionTime utctime.UTCTime) (RedelegationQueueRow, bool, error)
	DequeueAllMatureRedelegationQueue(blockTime utctime.UTCTime) ([]DVVTriplet, error)
}

type RedelegationQueueView struct {
	rdb *rdb.Handle
}

func NewRedelegationQueueView(handle *rdb.Handle) RedelegationQueue {
	return &RedelegationQueueView{
		handle,
	}
}

func (view *RedelegationQueueView) Upsert(row RedelegationQueueRow) error {

	dvvTripletsJSON, err := jsoniter.MarshalToString(row.DVVTriplets)
	if err != nil {
		return fmt.Errorf("error JSON marshalling DVVTriplets for upsertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Insert(
			"view_vd_redelegation_queue",
		).
		Columns(
			"completion_time",
			"dvv_triplets",
		).
		Values(
			view.rdb.Tton(&row.CompletionTime),
			dvvTripletsJSON,
		).
		Suffix("ON CONFLICT(completion_time) DO UPDATE SET dvv_triplets = EXCLUDED.dvv_triplets").
		ToSql()

	if err != nil {
		return fmt.Errorf("error building RedelegationQueueRow upsertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := view.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error upserting RedelegationQueueRow into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error upserting RedelegationQueueRow into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (view *RedelegationQueueView) FindBy(completionTime utctime.UTCTime) (RedelegationQueueRow, bool, error) {

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Select(
			"dvv_triplets",
		).
		From("view_vd_redelegation_queue").
		Where(
			"completion_time = ?",
			view.rdb.Tton(&completionTime),
		).
		ToSql()
	if err != nil {
		return RedelegationQueueRow{}, false, fmt.Errorf("error building select RedelegationQueueRow sql: %v: %w", err, rdb.ErrPrepare)
	}

	var row RedelegationQueueRow
	row.CompletionTime = completionTime

	var dvvTripletsJSON string
	if err = view.rdb.QueryRow(sql, sqlArgs...).Scan(
		&dvvTripletsJSON,
	); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			// When the row is not found, do not return error
			return RedelegationQueueRow{}, false, nil
		}
		return RedelegationQueueRow{}, false, fmt.Errorf("error scanning RedelegationQueueRow: %v: %w", err, rdb.ErrQuery)
	}

	if err = jsoniter.UnmarshalFromString(dvvTripletsJSON, &row.DVVTriplets); err != nil {
		return RedelegationQueueRow{}, false, fmt.Errorf("error unmarshalling DVVTriplets JSON: %v: %w", err, rdb.ErrQuery)
	}

	return row, true, nil
}

func (view *RedelegationQueueView) DequeueAllMatureRedelegationQueue(blockTime utctime.UTCTime) ([]DVVTriplet, error) {

	// Find all mature RedelegationQueueRow, then concate their DVVTriplets

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Select(
			"dvv_triplets",
		).
		From(
			"view_vd_redelegation_queue",
		).
		Where(
			"completion_time <= ?",
			view.rdb.Tton(&blockTime),
		).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building mature RedelegationQueueRow select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := view.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, fmt.Errorf("error executing mature RedelegationQueueRow select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	matureDVVTriplets := make([]DVVTriplet, 0)
	for rowsResult.Next() {

		var dvvTripletsJSON string
		if err = rowsResult.Scan(
			&dvvTripletsJSON,
		); err != nil {
			return nil, fmt.Errorf("error scanning RedelegationQueueRow: %v: %w", err, rdb.ErrQuery)
		}

		var dvvTriplets []DVVTriplet
		if err = jsoniter.UnmarshalFromString(dvvTripletsJSON, &dvvTriplets); err != nil {
			return nil, fmt.Errorf("error unmarshalling DVVTriplets JSON: %v: %w", err, rdb.ErrQuery)
		}

		matureDVVTriplets = append(matureDVVTriplets, dvvTriplets...)
	}

	// Optional TODO: De-duplicate the return slice. A same DVVTriplet could appear multiple times, we could avoid it here.
	//                At the moment, we are following CosmosSDK implementation, so not doing the de-duplicate.

	// Delete the mature rows

	sql, sqlArgs, err = view.rdb.StmtBuilder.
		Delete(
			"view_vd_redelegation_queue",
		).
		Where(
			"completion_time <= ?",
			view.rdb.Tton(&blockTime),
		).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building RedelegationQueueRow deletion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	_, err = view.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return nil, fmt.Errorf("error deleting RedelegationQueueRow from the table: %v: %w", err, rdb.ErrWrite)
	}

	return matureDVVTriplets, nil
}

// NOTES:
// - UNIQUE CompletionTime
type RedelegationQueueRow struct {
	CompletionTime utctime.UTCTime `json:"completionTime"`
	DVVTriplets    []DVVTriplet    `json:"dvvTriplets"`
}

type DVVTriplet struct {
	DelegatorAddress    string `json:"delegatorAddress"`
	ValidatorSrcAddress string `json:"validatorSrcAddress"`
	ValidatorDstAddress string `json:"validatorDstAddress"`
}
