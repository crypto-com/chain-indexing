package view

import (
	"errors"
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
	jsoniter "github.com/json-iterator/go"
)

type UnbondingDelegationQueue interface {
	Upsert(row UnbondingDelegationQueueRow) error
	FindBy(completionTime utctime.UTCTime) (UnbondingDelegationQueueRow, bool, error)
	DequeueAllMatureUnbondingDelegationQueue(blockTime utctime.UTCTime) ([]DVPair, error)
}

type UnbondingDelegationQueueView struct {
	rdb *rdb.Handle
}

func NewUnbondingDelegationQueueView(handle *rdb.Handle) UnbondingDelegationQueue {
	return &UnbondingDelegationQueueView{
		handle,
	}
}

func (view *UnbondingDelegationQueueView) Upsert(row UnbondingDelegationQueueRow) error {

	dvPairsJSON, err := jsoniter.MarshalToString(row.DVPairs)
	if err != nil {
		return fmt.Errorf("error JSON marshalling DVPairs for upsertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Insert(
			"view_vd_unbonding_delegation_queue",
		).
		Columns(
			"completion_time",
			"dv_pairs",
		).
		Values(
			view.rdb.Tton(&row.CompletionTime),
			dvPairsJSON,
		).
		Suffix("ON CONFLICT(completion_time) DO UPDATE SET dv_pairs = EXCLUDED.dv_pairs").
		ToSql()

	if err != nil {
		return fmt.Errorf("error building UnbondingDelegationQueueRow upsertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := view.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error upserting UnbondingDelegationQueueRow into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error upserting UnbondingDelegationQueueRow into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (view *UnbondingDelegationQueueView) FindBy(completionTime utctime.UTCTime) (UnbondingDelegationQueueRow, bool, error) {

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Select(
			"dv_pairs",
		).
		From("view_vd_unbonding_delegation_queue").
		Where(
			"completion_time = ?",
			view.rdb.Tton(&completionTime),
		).
		ToSql()
	if err != nil {
		return UnbondingDelegationQueueRow{}, false, fmt.Errorf("error building select UnbondingDelegationQueueRow sql: %v: %w", err, rdb.ErrPrepare)
	}

	var row UnbondingDelegationQueueRow
	row.CompletionTime = completionTime

	var dvPairsJSON string
	if err = view.rdb.QueryRow(sql, sqlArgs...).Scan(
		&dvPairsJSON,
	); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			// When the row is not found, do not return error
			return UnbondingDelegationQueueRow{}, false, nil
		}
		return UnbondingDelegationQueueRow{}, false, fmt.Errorf("error scanning RedelegationQueueRow: %v: %w", err, rdb.ErrQuery)
	}

	if err = jsoniter.UnmarshalFromString(dvPairsJSON, &row.DVPairs); err != nil {
		return UnbondingDelegationQueueRow{}, false, fmt.Errorf("error unmarshalling DVPairs JSON: %v: %w", err, rdb.ErrQuery)
	}

	return row, true, nil
}

func (view *UnbondingDelegationQueueView) DequeueAllMatureUnbondingDelegationQueue(blockTime utctime.UTCTime) ([]DVPair, error) {

	// Find all mature UnbondingDelegationQueueRow, then concate their DVPairs

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Select(
			"dv_pairs",
		).
		From(
			"view_vd_unbonding_delegation_queue",
		).
		Where(
			"completion_time <= ?",
			view.rdb.Tton(&blockTime),
		).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building mature UnbondingDelegationQueueRow select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := view.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, fmt.Errorf("error executing mature UnbondingDelegationQueueRow select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	matureDVPairs := make([]DVPair, 0)
	for rowsResult.Next() {

		var dvPairsJSON string
		if err = rowsResult.Scan(
			&dvPairsJSON,
		); err != nil {
			return nil, fmt.Errorf("error scanning UnbondingDelegationQueueRow: %v: %w", err, rdb.ErrQuery)
		}

		var dvPairs []DVPair
		if err = jsoniter.UnmarshalFromString(dvPairsJSON, &dvPairs); err != nil {
			return nil, fmt.Errorf("error unmarshalling DVPairs JSON: %v: %w", err, rdb.ErrQuery)
		}

		matureDVPairs = append(matureDVPairs, dvPairs...)
	}

	// Optional TODO: De-duplicate the return slice. A same DVPair could appear multiple times, we could avoid it here.
	//                At the moment, we are following CosmosSDK implementation, so not doing the de-duplicate.

	// Delete the mature rows

	sql, sqlArgs, err = view.rdb.StmtBuilder.
		Delete(
			"view_vd_unbonding_delegation_queue",
		).
		Where(
			"completion_time <= ?",
			view.rdb.Tton(&blockTime),
		).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building UnbondingDelegationQueueRow deletion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	_, err = view.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return nil, fmt.Errorf("error deleting UnbondingDelegationQueueRow from the table: %v: %w", err, rdb.ErrWrite)
	}

	return matureDVPairs, nil
}

// NOTES:
// - UNIQUE CompletionTime
type UnbondingDelegationQueueRow struct {
	CompletionTime utctime.UTCTime `json:"completionTime"`
	DVPairs        []DVPair        `json:"dvPairs"`
}

type DVPair struct {
	DelegatorAddress string `json:"delegatorAddress"`
	ValidatorAddress string `json:"validatorAddress"`
}
