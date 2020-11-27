package view

import (
	"fmt"

	"github.com/crypto-com/chainindex/appinterface/pagination"

	"github.com/crypto-com/chainindex/appinterface/rdb"
	"github.com/crypto-com/chainindex/internal/utctime"
)

type Validators struct {
	rdb *rdb.Handle
}

func NewValidators(handle *rdb.Handle) *Validators {
	return &Validators{
		handle,
	}
}

func (view *Validators) Insert(validator *ValidatorRow) error {
	var err error

	var sql string
	sql, _, err = view.rdb.StmtBuilder.Insert(
		"view_validators",
	).Columns(
		"operator_address",
		"consensus_node_address",
		"initial_delegator_address",
		"status",
		"jailed",
		"power",
		"unbonding_height",
		"unbonding_completion_time",
		"moniker",
		"identity",
		"website",
		"security_contact",
		"details",
		"commission_rate",
		"commission_max_rate",
		"commission_max_change_rate",
		"min_self_delegation",
	).Values("?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?").ToSql()
	if err != nil {
		return fmt.Errorf("error building blocks insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := view.rdb.Exec(sql,
		validator.OperatorAddress,
		validator.ConsensusNodeAddress,
		validator.InitialDelegatorAddress,
		validator.Status,
		validator.Jailed,
		validator.Power,
		validator.MaybeUnbondingHeight,
		validator.MaybeUnbondingCompletionTime,
		validator.Moniker,
		validator.Identity,
		validator.Website,
		validator.SecurityContact,
		validator.Details,
		validator.CommissionRate,
		validator.CommissionMaxRate,
		validator.CommissionMaxChangeRate,
		validator.MinSelfDelegation,
	)
	if err != nil {
		return fmt.Errorf("error inserting validator into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting validator into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (view *Validators) Update(validator *ValidatorRow) error {
	//row, _, err := view.List(pagination.NewOffsetPagination(1, 20))
	//fmt.Println(row, err)

	var unbondingCompletionTime interface{} = nil
	if validator.MaybeUnbondingCompletionTime != nil {
		unbondingCompletionTime = view.rdb.Tton(validator.MaybeUnbondingCompletionTime)
	}
	sql, sqlArgs, err := view.rdb.StmtBuilder.Update(
		"view_validators",
	).SetMap(map[string]interface{}{
		"initial_delegator_address": validator.InitialDelegatorAddress,
		"status":                    validator.Status,
		"jailed":                    validator.Jailed,
		"power":                     validator.Power,
		"unbonding_height":          validator.MaybeUnbondingHeight,
		"unbonding_completion_time": unbondingCompletionTime,
		"moniker":                   validator.Moniker,
		"identity":                  validator.Identity,
		"website":                   validator.Website,
		"security_contact":          validator.SecurityContact,
		"details":                   validator.Details,
		"commission_rate":           validator.CommissionRate,
	}).Where(
		"id = ?", validator.Id,
	).ToSql()
	if err != nil {
		return fmt.Errorf("error building validator update sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := view.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error updating validator into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error updating validator: no rows updated: %w", rdb.ErrWrite)
	}

	return nil
}

func (view *Validators) List(pagination *pagination.Pagination) ([]ValidatorRow, *pagination.PaginationResult, error) {
	stmtBuilder := view.rdb.StmtBuilder.Select(
		"id",
		"operator_address",
		"consensus_node_address",
		"initial_delegator_address",
		"status",
		"jailed",
		"power",
		"unbonding_height",
		"unbonding_completion_time",
		"moniker",
		"identity",
		"website",
		"security_contact",
		"details",
		"commission_rate",
		"commission_max_rate",
		"commission_max_change_rate",
		"min_self_delegation",
	).From(
		"view_validators",
	).OrderBy("id")

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		view.rdb.Runner,
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building blocks select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	fmt.Println(sql, sqlArgs)
	rowsResult, err := view.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing blocks select SQL: %v: %w", err, rdb.ErrQuery)
	}

	validators := make([]ValidatorRow, 0)
	for rowsResult.Next() {
		var validator ValidatorRow
		unbondingCompletionTimeReader := view.rdb.NtotReader()
		if err = rowsResult.Scan(
			&validator.Id,
			&validator.OperatorAddress,
			&validator.ConsensusNodeAddress,
			&validator.InitialDelegatorAddress,
			&validator.Status,
			&validator.Jailed,
			&validator.Power,
			&validator.MaybeUnbondingHeight,
			unbondingCompletionTimeReader.ScannableArg(),
			&validator.Moniker,
			&validator.Identity,
			&validator.Website,
			&validator.SecurityContact,
			&validator.Details,
			&validator.CommissionRate,
			&validator.CommissionMaxRate,
			&validator.CommissionMaxChangeRate,
			&validator.MinSelfDelegation,
		); err != nil {
			if err == rdb.ErrNoRows {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning validator row: %v: %w", err, rdb.ErrQuery)
		}
		unbondingCompletionTime, parseErr := unbondingCompletionTimeReader.Parse()
		if parseErr != nil {
			return nil, nil, fmt.Errorf("error parsing validator time: %v: %w", parseErr, rdb.ErrQuery)
		}
		if unbondingCompletionTime != nil {
			validator.MaybeUnbondingCompletionTime = unbondingCompletionTime
		}

		validators = append(validators, validator)
	}
	rowsResult.Close()

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return validators, paginationResult, nil
}

func (view *Validators) FindBy(identity ValidatorIdentity) (*ValidatorRow, error) {
	var err error

	selectStmtBuilder := view.rdb.StmtBuilder.Select(
		"id",
		"operator_address",
		"consensus_node_address",
		"initial_delegator_address",
		"status",
		"jailed",
		"power",
		"unbonding_height",
		"unbonding_completion_time",
		"moniker",
		"identity",
		"website",
		"security_contact",
		"details",
		"commission_rate",
		"commission_max_rate",
		"commission_max_change_rate",
		"min_self_delegation",
	).From(
		"view_validators",
	).OrderBy("id DESC")
	if identity.MaybeConsensusNodeAddress != nil {
		selectStmtBuilder = selectStmtBuilder.Where(
			"consensus_node_address = ?", *identity.MaybeConsensusNodeAddress,
		)
	}
	if identity.MaybeOperatorAddress != nil {
		selectStmtBuilder = selectStmtBuilder.Where("operator_address = ?", *identity.MaybeOperatorAddress)
	}

	sql, sqlArgs, err := selectStmtBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building validator selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	var validator ValidatorRow
	unbondingCompletionTimeReader := view.rdb.NtotReader()
	if err = view.rdb.QueryRow(sql, sqlArgs...).Scan(
		&validator.Id,
		&validator.OperatorAddress,
		&validator.ConsensusNodeAddress,
		&validator.InitialDelegatorAddress,
		&validator.Status,
		&validator.Jailed,
		&validator.Power,
		&validator.MaybeUnbondingHeight,
		unbondingCompletionTimeReader.ScannableArg(),
		&validator.Moniker,
		&validator.Identity,
		&validator.Website,
		&validator.SecurityContact,
		&validator.Details,
		&validator.CommissionRate,
		&validator.CommissionMaxRate,
		&validator.CommissionMaxChangeRate,
		&validator.MinSelfDelegation,
	); err != nil {
		if err == rdb.ErrNoRows {
			return nil, rdb.ErrNoRows
		}
		return nil, fmt.Errorf("error scanning validator row: %v: %w", err, rdb.ErrQuery)
	}
	unbondingCompletionTime, err := unbondingCompletionTimeReader.Parse()
	if err != nil {
		return nil, fmt.Errorf("error parsing validator time: %v: %w", err, rdb.ErrQuery)
	}
	if unbondingCompletionTime != nil {
		validator.MaybeUnbondingCompletionTime = unbondingCompletionTime
	}

	return &validator, nil
}

type ValidatorIdentity struct {
	MaybeConsensusNodeAddress *string
	MaybeOperatorAddress      *string
}

type ValidatorRow struct {
	Id                           int64            `json:"-"`
	OperatorAddress              string           `json:"operatorAddress"`
	ConsensusNodeAddress         string           `json:"consensusNodeAddress"`
	InitialDelegatorAddress      string           `json:"initialDelegatorAddress"`
	Status                       string           `json:"status"`
	Jailed                       bool             `json:"jailed"`
	Power                        string           `json:"power"`
	MaybeUnbondingHeight         *int64           `json:"unbondingHeight"`
	MaybeUnbondingCompletionTime *utctime.UTCTime `json:"unbondingCompletionTime"`
	Moniker                      string           `json:"moniker"`
	Identity                     string           `json:"identity"`
	Website                      string           `json:"website"`
	SecurityContact              string           `json:"securityContact"`
	Details                      string           `json:"details"`
	CommissionRate               string           `json:"commissionRate"`
	CommissionMaxRate            string           `json:"commissionMaxRate"`
	CommissionMaxChangeRate      string           `json:"commissionMaxChangeRate"`
	MinSelfDelegation            string           `json:"minSelfDelegation"`
}
