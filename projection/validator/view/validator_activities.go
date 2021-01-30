package view

import (
	"errors"
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"

	sq "github.com/Masterminds/squirrel"

	"github.com/crypto-com/chain-indexing/internal/utctime"

	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

// BlockEvents projection view implemented by relational database
type ValidatorActivities struct {
	rdb *rdb.Handle
}

func NewValidatorActivities(handle *rdb.Handle) *ValidatorActivities {
	return &ValidatorActivities{
		handle,
	}
}

func (validatorActivitiesView *ValidatorActivities) Insert(validatorActivity *ValidatorActivityRow) error {
	var err error

	var sql string
	sql, _, err = validatorActivitiesView.rdb.StmtBuilder.Insert(
		"view_validator_activities",
	).Columns(
		"block_height",
		"block_hash",
		"block_time",
		"transaction_hash",
		"operator_address",
		"success",
		"data",
	).Values("?", "?", "?", "?", "?", "?", "?").ToSql()
	if err != nil {
		return fmt.Errorf("error building valiator activity insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := validatorActivitiesView.rdb.Exec(sql,
		validatorActivity.BlockHeight,
		validatorActivity.BlockHash,
		validatorActivitiesView.rdb.Tton(&validatorActivity.BlockTime),
		validatorActivity.MaybeTransactionHash,
		validatorActivity.OperatorAddress,
		validatorActivity.Success,
		validatorActivity.Data,
	)
	if err != nil {
		return fmt.Errorf("error inserting validator activity into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting validator activity into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (validatorActivitiesView *ValidatorActivities) InsertAll(validatorActivities []ValidatorActivityRow) error {
	if len(validatorActivities) == 0 {
		return nil
	}

	pendingRowCount := 0
	var stmtBuilder sq.InsertBuilder

	validatorActivityCount := len(validatorActivities)
	for i, validatorActivity := range validatorActivities {
		if pendingRowCount == 0 {
			stmtBuilder = validatorActivitiesView.rdb.StmtBuilder.Insert(
				"view_validator_activities",
			).Columns(
				"block_height",
				"block_hash",
				"block_time",
				"transaction_hash",
				"operator_address",
				"success",
				"data",
			)
		}

		stmtBuilder = stmtBuilder.Values(
			validatorActivity.BlockHeight,
			validatorActivity.BlockHash,
			validatorActivitiesView.rdb.Tton(&validatorActivity.BlockTime),
			validatorActivity.MaybeTransactionHash,
			validatorActivity.OperatorAddress,
			validatorActivity.Success,
			validatorActivity.Data,
		)
		pendingRowCount += 1

		if pendingRowCount == 500 || i+1 == validatorActivityCount {
			sql, sqlArgs, err := stmtBuilder.ToSql()
			if err != nil {
				return fmt.Errorf("error building valiator activity insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
			}

			result, err := validatorActivitiesView.rdb.Exec(sql, sqlArgs...)
			if err != nil {
				return fmt.Errorf("error inserting validator activity into the table: %v: %w", err, rdb.ErrWrite)
			}
			if result.RowsAffected() != int64(pendingRowCount) {
				return fmt.Errorf("error inserting validator activities into the table: mismatch rows inserted: %w", rdb.ErrWrite)
			}

			pendingRowCount = 0
		}
	}

	return nil
}

func (validatorActivitiesView *ValidatorActivities) List(
	filter ValidatorActivitiesListFilter,
	order ValidatorActivitiesListOrder,
	pagination *pagination_interface.Pagination,
) ([]ValidatorActivityRow, *pagination_interface.PaginationResult, error) {
	stmtBuilder := validatorActivitiesView.rdb.StmtBuilder.Select(
		"block_height",
		"block_hash",
		"block_time",
		"transaction_hash",
		"operator_address",
		"success",
		"data",
	).From(
		"view_validator_activities",
	)

	if order.MaybeBlockHeight == nil {
		stmtBuilder = stmtBuilder.OrderBy("id")
	} else if *order.MaybeBlockHeight == view.ORDER_ASC {
		stmtBuilder = stmtBuilder.OrderBy("block_height")
	} else if *order.MaybeBlockHeight == view.ORDER_DESC {
		stmtBuilder = stmtBuilder.OrderBy("block_height DESC")
	}

	if filter.MaybeOperatorAddress != nil {
		stmtBuilder = stmtBuilder.Where("operator_address = ?", *filter.MaybeOperatorAddress)
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		validatorActivitiesView.rdb,
	).WithCustomTotalQueryFn(
		func(rdbHandle *rdb.Handle, _ sq.SelectBuilder) (int64, error) {
			identity := "-"
			if filter.MaybeOperatorAddress != nil {
				identity = *filter.MaybeOperatorAddress
			} else if filter.MaybeConsensusNodeAddress != nil {
				validatorsView := NewValidators(rdbHandle)
				validator, err := validatorsView.FindBy(ValidatorIdentity{
					MaybeConsensusNodeAddress: filter.MaybeConsensusNodeAddress,
				})
				if err != nil {
					return int64(0), err
				}
				identity = validator.OperatorAddress
			}
			totalView := NewValidatorActivitiesTotal(rdbHandle)
			total, err := totalView.FindBy(identity)
			if err != nil {
				return int64(0), err
			}
			return total, nil
		},
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building transactions select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := validatorActivitiesView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing transactions select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	validatorActivities := make([]ValidatorActivityRow, 0)
	for rowsResult.Next() {
		var validatorActivity ValidatorActivityRow

		blockTimeParser := validatorActivitiesView.rdb.NtotReader()
		if scanErr := rowsResult.Scan(
			&validatorActivity.BlockHeight,
			&validatorActivity.BlockHash,
			blockTimeParser.ScannableArg(),
			&validatorActivity.MaybeTransactionHash,
			&validatorActivity.OperatorAddress,
			&validatorActivity.Success,
			&validatorActivity.Data,
		); scanErr != nil {
			if errors.Is(scanErr, rdb.ErrNoRows) {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning transaction row: %v: %w", scanErr, rdb.ErrQuery)
		}

		blockTime, parseErr := blockTimeParser.Parse()
		if parseErr != nil {
			return nil, nil, fmt.Errorf("error parsing block time: %v", parseErr)
		}
		validatorActivity.BlockTime = *blockTime

		validatorActivities = append(validatorActivities, validatorActivity)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return validatorActivities, paginationResult, nil
}

type ValidatorActivitiesListFilter struct {
	MaybeOperatorAddress      *string
	MaybeConsensusNodeAddress *string
}

type ValidatorActivitiesListOrder struct {
	MaybeBlockHeight *view.ORDER
}

type ValidatorActivityRow struct {
	BlockHeight          int64                    `json:"blockHeight"`
	BlockTime            utctime.UTCTime          `json:"blockTime"`
	BlockHash            string                   `json:"blockHash"`
	MaybeTransactionHash *string                  `json:"transactionHash"`
	OperatorAddress      string                   `json:"operatorAddress"`
	Success              bool                     `json:"success"`
	Data                 ValidatorActivityRowData `json:"activity"`
}

type ValidatorActivityRowData struct {
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
}
