package view

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"

	sq "github.com/Masterminds/squirrel"

	"github.com/crypto-com/chain-indexing/projection/validator/constants"

	"github.com/crypto-com/chain-indexing/appinterface/pagination"
	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type Validators struct {
	rdb *rdb.Handle
}

func NewValidators(handle *rdb.Handle) *Validators {
	return &Validators{
		handle,
	}
}

func (validatorsView *Validators) LastJoinedBlockHeight(
	operatorAddress string,
	consensusNodeAddress string,
) (bool, int64, error) {
	var err error

	var sql string
	var sqlArgs []interface{}
	if sql, sqlArgs, err = validatorsView.rdb.StmtBuilder.Select(
		"joined_at_block_height",
	).From(
		"view_validators",
	).Where(
		"operator_address = ? AND consensus_node_address = ?", operatorAddress, consensusNodeAddress,
	).ToSql(); err != nil {
		return false, int64(0), fmt.Errorf("error building validator existencen query sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	var joinedAtBlockHeight int64
	if err = validatorsView.rdb.QueryRow(sql, sqlArgs...).Scan(&joinedAtBlockHeight); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return false, int64(0), nil
		}
		return false, int64(0), fmt.Errorf("error query validator existence: %v", err)
	}

	return true, joinedAtBlockHeight, nil
}

func (validatorsView *Validators) Upsert(validator *ValidatorRow) error {
	sql, sqlArgs, err := validatorsView.rdb.StmtBuilder.Insert(
		"view_validators",
	).Columns(
		"operator_address",
		"consensus_node_address",
		"initial_delegator_address",
		"tendermint_pubkey",
		"tendermint_address",
		"status",
		"jailed",
		"joined_at_block_height",
		"power",
		"moniker",
		"identity",
		"website",
		"security_contact",
		"details",
		"commission_rate",
		"commission_max_rate",
		"commission_max_change_rate",
		"min_self_delegation",
	).Values(
		validator.OperatorAddress,
		validator.ConsensusNodeAddress,
		validator.InitialDelegatorAddress,
		validator.TendermintPubkey,
		validator.TendermintAddress,
		validator.Status,
		validator.Jailed,
		validator.JoinedAtBlockHeight,
		validator.Power,
		validator.Moniker,
		validator.Identity,
		validator.Website,
		validator.SecurityContact,
		validator.Details,
		validator.CommissionRate,
		validator.CommissionMaxRate,
		validator.CommissionMaxChangeRate,
		validator.MinSelfDelegation,
	).Suffix(`ON CONFLICT (operator_address, consensus_node_address) DO UPDATE SET
		initial_delegator_address = EXCLUDED.initial_delegator_address,
		status = EXCLUDED.status,
		jailed = EXCLUDED.jailed,
		joined_at_block_height = EXCLUDED.joined_at_block_height,
		power = EXCLUDED.power,
		moniker = EXCLUDED.moniker,
		identity = EXCLUDED.identity,
		website = EXCLUDED.website,
		security_contact = EXCLUDED.security_contact,
		details = EXCLUDED.details,
		commission_rate = EXCLUDED.commission_rate,
		commission_max_rate = EXCLUDED.commission_max_rate,
		commission_max_change_rate = EXCLUDED.commission_max_change_rate,
		min_self_delegation = EXCLUDED.min_self_delegation
	`).ToSql()
	if err != nil {
		return fmt.Errorf("error building validator upsertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}
	result, err := validatorsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error upserting validator into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error upserting validator into the table: no rows inserted: %w", rdb.ErrWrite)
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
		"tendermint_pubkey",
		"tendermint_address",
		"status",
		"jailed",
		"joined_at_block_height",
		"power",
		"moniker",
		"identity",
		"website",
		"security_contact",
		"details",
		"commission_rate",
		"commission_max_rate",
		"commission_max_change_rate",
		"min_self_delegation",
	).Values("?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?", "?").ToSql()
	if err != nil {
		return fmt.Errorf("error building validator insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := validatorsView.rdb.Exec(sql,
		validator.OperatorAddress,
		validator.ConsensusNodeAddress,
		validator.InitialDelegatorAddress,
		validator.TendermintPubkey,
		validator.TendermintAddress,
		validator.Status,
		validator.Jailed,
		validator.JoinedAtBlockHeight,
		validator.Power,
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
	sql, sqlArgs, err := validatorsView.rdb.StmtBuilder.Update(
		"view_validators",
	).SetMap(map[string]interface{}{
		"initial_delegator_address":  validator.InitialDelegatorAddress,
		"status":                     validator.Status,
		"jailed":                     validator.Jailed,
		"power":                      validator.Power,
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
	MaybeStatus              *view.ORDER
	MaybePower               *view.ORDER
	MaybeJoinedAtBlockHeight *view.ORDER
}

func (validatorsView *Validators) ListAll(
	filter ValidatorsListFilter,
	order ValidatorsListOrder,
) ([]ValidatorRow, error) {
	orderClauses := make([]string, 0)

	if order.MaybeJoinedAtBlockHeight != nil {
		if *order.MaybeStatus == view.ORDER_ASC {
			statusOrder := "CASE UPPER(status) " +
				"WHEN 'BONDED' THEN 1" +
				"WHEN 'UNBONDING' THEN 2" +
				"WHEN 'JAILED' THEN 3" +
				"WHEN 'UNBONDED' THEN 4" +
				"ELSE 5 END"
			orderClauses = append(orderClauses, statusOrder)
		} else if *order.MaybeStatus == view.ORDER_DESC {
			statusOrder := "CASE UPPER(status) " +
				"WHEN 'UNBONDED' THEN 1" +
				"WHEN 'JAILED' THEN 2" +
				"WHEN 'UNBONDING' THEN 3" +
				"WHEN 'BONDED' THEN 4" +
				"ELSE 5 END"
			orderClauses = append(orderClauses, statusOrder)
		}
	}

	if order.MaybePower != nil {
		if *order.MaybePower == view.ORDER_ASC {
			orderClauses = append(orderClauses, "CAST(power AS BIGINT)")
		} else if *order.MaybePower == view.ORDER_DESC {
			// DESC order
			orderClauses = append(orderClauses, "CAST(power AS BIGINT) DESC")
		}
	}

	if order.MaybeJoinedAtBlockHeight != nil {
		if *order.MaybeJoinedAtBlockHeight == view.ORDER_ASC {
			orderClauses = append(orderClauses, "joined_at_block_height")
		} else if *order.MaybeJoinedAtBlockHeight == view.ORDER_DESC {
			// DESC order
			orderClauses = append(orderClauses, "joined_at_block_height DESC")
		}
	}

	var statusOrCondition sq.Or
	if filter.MaybeStatuses != nil {
		statusOrCondition = make(sq.Or, 0)
		for _, status := range filter.MaybeStatuses {
			statusOrCondition = append(statusOrCondition, sq.Eq{
				"status": status,
			})
		}
	}

	stmtBuilder := validatorsView.rdb.StmtBuilder.Select(
		"id",
		"operator_address",
		"consensus_node_address",
		"initial_delegator_address",
		"tendermint_pubkey",
		"tendermint_address",
		"status",
		"jailed",
		"joined_at_block_height",
		"power",
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
	stmtBuilder = stmtBuilder.OrderBy(orderClauses...)

	if statusOrCondition != nil {
		stmtBuilder = stmtBuilder.Where(statusOrCondition)
	}
	sql, sqlArgs, err := stmtBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building validators select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := validatorsView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, fmt.Errorf("error executing validators select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	validators := make([]ValidatorRow, 0)
	for rowsResult.Next() {
		var validator ValidatorRow
		if err = rowsResult.Scan(
			&validator.MaybeId,
			&validator.OperatorAddress,
			&validator.ConsensusNodeAddress,
			&validator.InitialDelegatorAddress,
			&validator.TendermintPubkey,
			&validator.TendermintAddress,
			&validator.Status,
			&validator.Jailed,
			&validator.JoinedAtBlockHeight,
			&validator.Power,
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

		validators = append(validators, validator)
	}

	return validators, nil
}

func (validatorsView *Validators) List(
	filter ValidatorsListFilter,
	order ValidatorsListOrder,
	pagination *pagination_interface.Pagination,
) ([]ListValidatorRow, *pagination.PaginationResult, error) {
	if pagination.Type() != pagination_interface.PAGINATION_OFFSET {
		panic(fmt.Sprintf("unsupported pagination type: %s", pagination.Type()))
	}

	orderClauses := make([]string, 0)

	if order.MaybeJoinedAtBlockHeight != nil {
		if *order.MaybeStatus == view.ORDER_ASC {
			statusOrder := "CASE UPPER(status) " +
				"WHEN 'BONDED' THEN 1" +
				"WHEN 'UNBONDING' THEN 2" +
				"WHEN 'JAILED' THEN 3" +
				"WHEN 'UNBONDED' THEN 4" +
				"ELSE 5 END"
			orderClauses = append(orderClauses, statusOrder)
		} else if *order.MaybeStatus == view.ORDER_DESC {
			statusOrder := "CASE UPPER(status) " +
				"WHEN 'UNBONDED' THEN 1" +
				"WHEN 'JAILED' THEN 2" +
				"WHEN 'UNBONDING' THEN 3" +
				"WHEN 'BONDED' THEN 4" +
				"ELSE 5 END"
			orderClauses = append(orderClauses, statusOrder)
		}
	}

	if order.MaybePower != nil {
		if *order.MaybePower == view.ORDER_ASC {
			orderClauses = append(orderClauses, "CAST(power AS BIGINT)")
		} else if *order.MaybePower == view.ORDER_DESC {
			// DESC order
			orderClauses = append(orderClauses, "CAST(power AS BIGINT) DESC")
		}
	}

	if order.MaybeJoinedAtBlockHeight != nil {
		if *order.MaybeJoinedAtBlockHeight == view.ORDER_ASC {
			orderClauses = append(orderClauses, "joined_at_block_height")
		} else if *order.MaybeJoinedAtBlockHeight == view.ORDER_DESC {
			// DESC order
			orderClauses = append(orderClauses, "joined_at_block_height DESC")
		}
	}

	cumulativePowerStmtBuilder := validatorsView.rdb.StmtBuilder.Select(
		"power",
	).From(
		"view_validators",
	).Offset(0).Limit(
		uint64(pagination.OffsetParams().Offset()),
	)
	cumulativePowerStmtBuilder = cumulativePowerStmtBuilder.OrderBy(orderClauses...)

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
	defer rowsResult.Close()

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
		"tendermint_pubkey",
		"tendermint_address",
		"status",
		"jailed",
		"joined_at_block_height",
		"power",
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
	stmtBuilder = stmtBuilder.OrderBy(orderClauses...)

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
		if err = rowsResult.Scan(
			&validator.MaybeId,
			&validator.OperatorAddress,
			&validator.ConsensusNodeAddress,
			&validator.InitialDelegatorAddress,
			&validator.TendermintPubkey,
			&validator.TendermintAddress,
			&validator.Status,
			&validator.Jailed,
			&validator.JoinedAtBlockHeight,
			&validator.Power,
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
			return nil, fmt.Errorf("error creating big.Float from power retrieved: %s", powerStr)
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
		"tendermint_pubkey",
		"tendermint_address",
		"status",
		"jailed",
		"joined_at_block_height",
		"power",
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
	defer rowsResult.Close()

	validators := make([]ValidatorRow, 0)
	for rowsResult.Next() {
		var validator ValidatorRow
		if err = rowsResult.Scan(
			&validator.MaybeId,
			&validator.OperatorAddress,
			&validator.ConsensusNodeAddress,
			&validator.InitialDelegatorAddress,
			&validator.TendermintPubkey,
			&validator.TendermintAddress,
			&validator.Status,
			&validator.Jailed,
			&validator.JoinedAtBlockHeight,
			&validator.Power,
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
		"tendermint_pubkey",
		"tendermint_address",
		"status",
		"jailed",
		"joined_at_block_height",
		"power",
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
	if err = validatorsView.rdb.QueryRow(sql, sqlArgs...).Scan(
		&validator.MaybeId,
		&validator.OperatorAddress,
		&validator.ConsensusNodeAddress,
		&validator.InitialDelegatorAddress,
		&validator.TendermintPubkey,
		&validator.TendermintAddress,
		&validator.Status,
		&validator.Jailed,
		&validator.JoinedAtBlockHeight,
		&validator.Power,
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

	return &validator, nil
}

func (validatorsView *Validators) Count(filter CountFilter) (int64, error) {
	var count int64

	stmt := validatorsView.rdb.StmtBuilder.Select(
		"COUNT(*)",
	).From(
		"view_validators",
	)

	if filter.MaybeStatus != nil {
		orClause := make(sq.Or, 0, len(filter.MaybeStatus))
		for _, status := range filter.MaybeStatus {
			orClause = append(orClause, sq.Eq{"status": status})
		}
		stmt = stmt.Where(orClause)
	}

	sql, sqlArgs, err := stmt.ToSql()
	if err != nil {
		return int64(0), fmt.Errorf("error building validator count sql: %v: %w", err, rdb.ErrPrepare)
	}

	if err := validatorsView.rdb.QueryRow(sql, sqlArgs...).Scan(&count); err != nil {
		return int64(0), fmt.Errorf("error getting validators count: %v", err)
	}
	return count, nil
}

type CountFilter struct {
	MaybeStatus []string
}

type ValidatorIdentity struct {
	MaybeConsensusNodeAddress *string
	MaybeOperatorAddress      *string
}

type ValidatorRow struct {
	MaybeId                 *int64 `json:"-"`
	OperatorAddress         string `json:"operatorAddress"`
	ConsensusNodeAddress    string `json:"consensusNodeAddress"`
	InitialDelegatorAddress string `json:"initialDelegatorAddress"`
	TendermintPubkey        string `json:"tendermintPubkey"`
	TendermintAddress       string `json:"tendermintAddress"`
	Status                  string `json:"status"`
	Jailed                  bool   `json:"jailed"`
	JoinedAtBlockHeight     int64  `json:"joinedAtBlockHeight"`
	Power                   string `json:"power"`
	Moniker                 string `json:"moniker"`
	Identity                string `json:"identity"`
	Website                 string `json:"website"`
	SecurityContact         string `json:"securityContact"`
	Details                 string `json:"details"`
	CommissionRate          string `json:"commissionRate"`
	CommissionMaxRate       string `json:"commissionMaxRate"`
	CommissionMaxChangeRate string `json:"commissionMaxChangeRate"`
	MinSelfDelegation       string `json:"minSelfDelegation"`
}

type ListValidatorRow struct {
	ValidatorRow

	PowerPercentage           string `json:"powerPercentage"`
	CumulativePowerPercentage string `json:"cumulativePowerPercentage"`
}
