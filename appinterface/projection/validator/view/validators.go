package view

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"

	sq "github.com/Masterminds/squirrel"

	"github.com/crypto-com/chain-indexing/appinterface/projection/validator/constants"

	"github.com/crypto-com/chain-indexing/appinterface/pagination"
	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/internal/utctime"
)

type Validators struct {
	rdb *rdb.Handle
}

func NewValidators(handle *rdb.Handle) *Validators {
	return &Validators{
		handle,
	}
}

func (validatorsView *Validators) Upsert(validator *ValidatorRow) error {
	var err error

	var sql string
	var sqlArgs []interface{}
	if sql, sqlArgs, err = validatorsView.rdb.StmtBuilder.Select(
		"id",
	).From(
		"view_validators",
	).Where(
		"operator_address = ? AND consensus_node_address = ?",
		validator.OperatorAddress,
		validator.ConsensusNodeAddress,
	).ToSql(); err != nil {
		return fmt.Errorf("error building validator existencen query sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	var existingValidatorId int64
	if err = validatorsView.rdb.QueryRow(sql, sqlArgs...).Scan(&existingValidatorId); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			if err = validatorsView.Insert(validator); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("error checking validator existence: %v", err)
		}
		return nil
	}
	// Do update
	mutValidator := *validator
	mutValidator.MaybeId = &existingValidatorId
	if updateErr := validatorsView.Update(&mutValidator); updateErr != nil {
		return updateErr
	}

	return nil
}

func (validatorsView *Validators) Insert(validator *ValidatorRow) error {
	var err error

	var sql string
	sql, _, err = validatorsView.rdb.StmtBuilder.Insert(
		"view_validators",
	).Columns(
		"operator_address",
		"consensus_node_address",
		"initial_delegator_address",
		"status",
		"jailed",
		"joined_at_block_height",
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
	).Values("?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?").ToSql()
	if err != nil {
		return fmt.Errorf("error building validator insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	var unbondingCompletionTime interface{} = nil
	if validator.MaybeUnbondingCompletionTime != nil {
		unbondingCompletionTime = validatorsView.rdb.Tton(validator.MaybeUnbondingCompletionTime)
	}
	result, err := validatorsView.rdb.Exec(sql,
		validator.OperatorAddress,
		validator.ConsensusNodeAddress,
		validator.InitialDelegatorAddress,
		validator.Status,
		validator.Jailed,
		validator.JoinedAtBlockHeight,
		validator.Power,
		validator.MaybeUnbondingHeight,
		unbondingCompletionTime,
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

func (validatorsView *Validators) Update(validator *ValidatorRow) error {
	var unbondingCompletionTime interface{} = nil
	if validator.MaybeUnbondingCompletionTime != nil {
		unbondingCompletionTime = validatorsView.rdb.Tton(validator.MaybeUnbondingCompletionTime)
	}
	sql, sqlArgs, err := validatorsView.rdb.StmtBuilder.Update(
		"view_validators",
	).SetMap(map[string]interface{}{
		"initial_delegator_address":  validator.InitialDelegatorAddress,
		"status":                     validator.Status,
		"jailed":                     validator.Jailed,
		"power":                      validator.Power,
		"unbonding_height":           validator.MaybeUnbondingHeight,
		"unbonding_completion_time":  unbondingCompletionTime,
		"moniker":                    validator.Moniker,
		"identity":                   validator.Identity,
		"website":                    validator.Website,
		"security_contact":           validator.SecurityContact,
		"details":                    validator.Details,
		"commission_rate":            validator.CommissionRate,
		"commission_max_rate":        validator.CommissionMaxRate,
		"commission_max_change_rate": validator.CommissionMaxChangeRate,
		"min_self_delegation":        validator.MinSelfDelegation,
	}).Where(
		"id = ?", validator.MaybeId,
	).ToSql()
	if err != nil {
		return fmt.Errorf("error building validator update sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := validatorsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error updating validator into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error updating validator: no rows updated: %w", rdb.ErrWrite)
	}

	return nil
}

type ValidatorsListFilter struct {
	MaybeStatuses []constants.Status
}

type ValidatorsListOrder struct {
	MaybePower *view.ORDER
}

func (validatorsView *Validators) List(
	filter ValidatorsListFilter,
	order ValidatorsListOrder,
	pagination *pagination_interface.Pagination,
) ([]ListValidatorRow, *pagination.PaginationResult, error) {
	if pagination.Type() != pagination_interface.PAGINATION_OFFSET {
		panic(fmt.Sprintf("unsupported pagination type: %s", pagination.Type()))
	}
	cumulativePowerStmtBuilder := validatorsView.rdb.StmtBuilder.Select(
		"power",
	).From(
		"view_validators",
	).Offset(0).Limit(
		uint64(pagination.OffsetParams().Offset()),
	)
	if order.MaybePower == nil {
		cumulativePowerStmtBuilder = cumulativePowerStmtBuilder.OrderBy("id")
	} else if *order.MaybePower == view.ORDER_ASC {
		cumulativePowerStmtBuilder = cumulativePowerStmtBuilder.OrderBy("power")
	} else {
		// DESC order
		cumulativePowerStmtBuilder = cumulativePowerStmtBuilder.OrderBy("power DESC")
	}

	var statusOrCondition sq.Or
	if filter.MaybeStatuses != nil {
		statusOrCondition = make(sq.Or, 0)
		for _, status := range filter.MaybeStatuses {
			statusOrCondition = append(statusOrCondition, sq.Eq{
				"status": status,
			})
		}
		cumulativePowerStmtBuilder = cumulativePowerStmtBuilder.Where(statusOrCondition)
	}
	cumulativePowerSql, cumulativePowerSqlArgs, err := cumulativePowerStmtBuilder.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf(
			"error building validators cumulative power SQL: %v, %w", err, rdb.ErrBuildSQLStmt,
		)
	}
	rowsResult, err := validatorsView.rdb.Query(cumulativePowerSql, cumulativePowerSqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing cumulative power SQL: %v", err)
	}

	cumulativePower := new(big.Float).SetInt64(int64(0))
	for rowsResult.Next() {
		var powerStr string
		if scanErr := rowsResult.Scan(&powerStr); scanErr != nil {
			if errors.Is(scanErr, rdb.ErrNoRows) {
				return nil, nil, fmt.Errorf(
					"error executing validators cumulative power SQL: %v, %w", scanErr, rdb.ErrQuery,
				)
			}
			powerStr = "0"
		}
		power, ok := new(big.Float).SetString(powerStr)
		if !ok {
			return nil, nil, fmt.Errorf(
				"error parsing validators power: %v, %w", err, rdb.ErrTypeConv,
			)
		}
		cumulativePower = new(big.Float).Add(cumulativePower, power)
	}
	totalPower, err := validatorsView.totalPower()
	if err != nil {
		return nil, nil, fmt.Errorf("error getting total power: %v", err)
	}

	stmtBuilder := validatorsView.rdb.StmtBuilder.Select(
		"id",
		"operator_address",
		"consensus_node_address",
		"initial_delegator_address",
		"status",
		"jailed",
		"joined_at_block_height",
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
	)
	if order.MaybePower == nil {
		stmtBuilder = stmtBuilder.OrderBy("id")
	} else if *order.MaybePower == view.ORDER_ASC {
		stmtBuilder = stmtBuilder.OrderBy("power")
	} else {
		// DESC order
		stmtBuilder = stmtBuilder.OrderBy("power DESC")
	}

	if statusOrCondition != nil {
		stmtBuilder = stmtBuilder.Where(statusOrCondition)
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		validatorsView.rdb,
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building blocks select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err = validatorsView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing blocks select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	validators := make([]ListValidatorRow, 0)
	for rowsResult.Next() {
		var validator ValidatorRow
		unbondingCompletionTimeReader := validatorsView.rdb.NtotReader()
		if err = rowsResult.Scan(
			&validator.MaybeId,
			&validator.OperatorAddress,
			&validator.ConsensusNodeAddress,
			&validator.InitialDelegatorAddress,
			&validator.Status,
			&validator.Jailed,
			&validator.JoinedAtBlockHeight,
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
			if errors.Is(err, rdb.ErrNoRows) {
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

		power, ok := new(big.Float).SetString(validator.Power)
		if !ok {
			return nil, nil, errors.New("error parsing validator power as float")
		}

		if totalPower.Cmp(new(big.Float).SetInt64(int64(0))) == 0 {
			validators = append(validators, ListValidatorRow{
				validator,

				"0",
				"0",
			})
		} else {
			cumulativePower = new(big.Float).Add(cumulativePower, power)
			cumulativePowerPercentage := new(big.Float).Quo(cumulativePower, totalPower)
			powerPercentage := new(big.Float).Quo(power, totalPower)
			validators = append(validators, ListValidatorRow{
				validator,

				powerPercentage.String(),
				cumulativePowerPercentage.String(),
			})
		}
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return validators, paginationResult, nil
}

func (validatorsView *Validators) totalPower() (*big.Float, error) {
	sql, _, _ := validatorsView.rdb.StmtBuilder.Select("power").From("view_validators").ToSql()
	rowsResult, err := validatorsView.rdb.Query(sql)
	if err != nil {
		return nil, fmt.Errorf("error getting validators from table: %v", err)
	}
	defer rowsResult.Close()

	totalPower := new(big.Float).SetInt64(int64(0))
	for rowsResult.Next() {
		var powerStr string
		if scanErr := rowsResult.Scan(&powerStr); scanErr != nil {
			return nil, fmt.Errorf("error scanning power from validators: %v", scanErr)
		}
		power, ok := new(big.Float).SetString(powerStr)
		if !ok {
			return nil, fmt.Errorf("error creating big.Float from power retrieved: %v", err)
		}
		totalPower = new(big.Float).Add(totalPower, power)
	}

	return totalPower, nil
}

func (validatorsView *Validators) Search(keyword string) ([]ValidatorRow, error) {
	keyword = strings.ToLower(keyword)
	sql, sqlArgs, err := validatorsView.rdb.StmtBuilder.Select(
		"id",
		"operator_address",
		"consensus_node_address",
		"initial_delegator_address",
		"status",
		"jailed",
		"joined_at_block_height",
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
	).Where(
		"operator_address = ? OR consensus_node_address = ? OR LOWER(moniker) LIKE ?",
		keyword, keyword, fmt.Sprintf("%%%s%%", keyword),
	).OrderBy("id").ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building blocks select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := validatorsView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, fmt.Errorf("error executing blocks select SQL: %v: %w", err, rdb.ErrQuery)
	}

	validators := make([]ValidatorRow, 0)
	for rowsResult.Next() {
		var validator ValidatorRow
		unbondingCompletionTimeReader := validatorsView.rdb.NtotReader()
		if err = rowsResult.Scan(
			&validator.MaybeId,
			&validator.OperatorAddress,
			&validator.ConsensusNodeAddress,
			&validator.InitialDelegatorAddress,
			&validator.Status,
			&validator.Jailed,
			&validator.JoinedAtBlockHeight,
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
			if errors.Is(err, rdb.ErrNoRows) {
				return nil, rdb.ErrNoRows
			}
			return nil, fmt.Errorf("error scanning validator row: %v: %w", err, rdb.ErrQuery)
		}
		unbondingCompletionTime, parseErr := unbondingCompletionTimeReader.Parse()
		if parseErr != nil {
			return nil, fmt.Errorf("error parsing validator time: %v: %w", parseErr, rdb.ErrQuery)
		}
		if unbondingCompletionTime != nil {
			validator.MaybeUnbondingCompletionTime = unbondingCompletionTime
		}

		validators = append(validators, validator)
	}
	rowsResult.Close()

	return validators, nil
}

func (validatorsView *Validators) FindBy(identity ValidatorIdentity) (*ValidatorRow, error) {
	var err error

	selectStmtBuilder := validatorsView.rdb.StmtBuilder.Select(
		"id",
		"operator_address",
		"consensus_node_address",
		"initial_delegator_address",
		"status",
		"jailed",
		"joined_at_block_height",
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
	unbondingCompletionTimeReader := validatorsView.rdb.NtotReader()
	if err = validatorsView.rdb.QueryRow(sql, sqlArgs...).Scan(
		&validator.MaybeId,
		&validator.OperatorAddress,
		&validator.ConsensusNodeAddress,
		&validator.InitialDelegatorAddress,
		&validator.Status,
		&validator.Jailed,
		&validator.JoinedAtBlockHeight,
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
		if errors.Is(err, rdb.ErrNoRows) {
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

func (validatorsView *Validators) Count() (int64, error) {
	var count int64
	if err := validatorsView.rdb.QueryRow("SELECT COUNT(*) FROM view_validators").Scan(&count); err != nil {
		return int64(0), fmt.Errorf("error getting validators count: %v", err)
	}
	return count, nil
}

type ValidatorIdentity struct {
	MaybeConsensusNodeAddress *string
	MaybeOperatorAddress      *string
}

type ValidatorRow struct {
	MaybeId                      *int64           `json:"-"`
	OperatorAddress              string           `json:"operatorAddress"`
	ConsensusNodeAddress         string           `json:"consensusNodeAddress"`
	InitialDelegatorAddress      string           `json:"initialDelegatorAddress"`
	Status                       string           `json:"status"`
	Jailed                       bool             `json:"jailed"`
	JoinedAtBlockHeight          int64            `json:"joinedAtBlockHeight"`
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

type ListValidatorRow struct {
	ValidatorRow

	PowerPercentage           string `json:"powerPercentage"`
	CumulativePowerPercentage string `json:"cumulativePowerPercentage"`
}
