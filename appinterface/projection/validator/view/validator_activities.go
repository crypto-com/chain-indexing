package view

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/crypto-com/chainindex/internal/utctime"

	pagination_interface "github.com/crypto-com/chainindex/appinterface/pagination"
	"github.com/crypto-com/chainindex/appinterface/rdb"
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

func (view *ValidatorActivities) Insert(validatorActivity *ValidatorActivityRow) error {
	var err error

	var sql string
	sql, _, err = view.rdb.StmtBuilder.Insert(
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

	result, err := view.rdb.Exec(sql,
		validatorActivity.BlockHeight,
		validatorActivity.BlockHash,
		view.rdb.Tton(&validatorActivity.BlockTime),
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

func (view *ValidatorActivities) List(
	filter ValidatorActivitiesListFilter,
	pagination *pagination_interface.Pagination,
) ([]ValidatorActivityRow, *pagination_interface.PaginationResult, error) {
	stmtBuilder := view.rdb.StmtBuilder.Select(
		"block_height",
		"block_hash",
		"block_time",
		"transaction_hash",
		"operator_address",
		"success",
		"data",
	).From(
		"view_validator_activities",
	).OrderBy(
		"id",
	)

	if filter.MaybeOperatorAddress != nil {
		stmtBuilder = stmtBuilder.Where("operator_address = ?", *filter.MaybeOperatorAddress)
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		view.rdb,
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

	rowsResult, err := view.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing transactions select SQL: %v: %w", err, rdb.ErrQuery)
	}

	validatorActivities := make([]ValidatorActivityRow, 0)
	for rowsResult.Next() {
		var validatorActivity ValidatorActivityRow

		blockTimeParser := view.rdb.NtotReader()
		if scanErr := rowsResult.Scan(
			&validatorActivity.BlockHeight,
			&validatorActivity.BlockHash,
			blockTimeParser.ScannableArg(),
			&validatorActivity.MaybeTransactionHash,
			&validatorActivity.OperatorAddress,
			&validatorActivity.Success,
			&validatorActivity.Data,
		); scanErr != nil {
			if scanErr == rdb.ErrNoRows {
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
