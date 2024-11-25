package view

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	sq "github.com/Masterminds/squirrel"

	"github.com/crypto-com/chain-indexing/appinterface/pagination"
	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/projection/validator/constants"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
)

type Validators interface {
	LastJoinedBlockHeight(
		operatorAddress string,
		consensusNodeAddress string,
	) (bool, int64, error)
	Upsert(validator *ValidatorRow) error
	Insert(validator *ValidatorRow) error
	totalPower() (*big.Float, error)
	Search(
		keyword string,
		maybeLowestBlockHeight *int64,
	) ([]ValidatorRow, error)
	FindBy(
		identity ValidatorIdentity,
		maybeLowestBlockHeight *int64,
	) (*ValidatorRow, error)
	Update(validator *ValidatorRow) error
	UpdateAllValidatorUpTime([]ValidatorRow) error
	ListAll(
		filter ValidatorsListFilter,
		order ValidatorsListOrder,
	) ([]ValidatorRow, error)

	List(
		filter ValidatorsListFilter,
		order ValidatorsListOrder,
		pagination *pagination_interface.Pagination) ([]ListValidatorRow, *pagination.PaginationResult, error)
	Count(filter CountFilter) (int64, error)
	LastHandledHeight() (int64, error)
}

type ValidatorsView struct {
	rdb *rdb.Handle
}

func NewValidatorsView(handle *rdb.Handle) Validators {
	return &ValidatorsView{
		handle,
	}
}

const (
	DEFAULT_LOWEST_ACTIVE_BLOCKS_BLOCK_HEIGHT = int64(0)
	VALIDATORS_TABLE_NAME                     = "view_validators"
)

func (validatorsView *ValidatorsView) LastHandledHeight() (int64, error) {
	var err error

	var sql string
	var sqlArgs []interface{}
	if sql, sqlArgs, err = validatorsView.rdb.StmtBuilder.Select(
		"last_handled_event_height",
	).From(
		"projections",
	).Where(
		"id = ?", "Validator",
	).ToSql(); err != nil {
		return int64(0), fmt.Errorf("error build last handled height from Validator projection sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	var lastHandledEventHeight int64
	if err = validatorsView.rdb.QueryRow(sql, sqlArgs...).Scan(&lastHandledEventHeight); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return int64(0), nil
		}
		return int64(0), fmt.Errorf("error query last handled height from Validator projection sql: %v", err)
	}

	return lastHandledEventHeight, nil
}

func (validatorsView *ValidatorsView) LastJoinedBlockHeight(
	operatorAddress string,
	consensusNodeAddress string,
) (bool, int64, error) {
	var err error

	var sql string
	var sqlArgs []interface{}
	if sql, sqlArgs, err = validatorsView.rdb.StmtBuilder.Select(
		"joined_at_block_height",
	).From(
		VALIDATORS_TABLE_NAME,
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

func (validatorsView *ValidatorsView) Upsert(validator *ValidatorRow) error {
	sql, sqlArgs, err := validatorsView.rdb.StmtBuilder.Insert(
		VALIDATORS_TABLE_NAME,
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
		"total_signed_block",
		"total_active_block",
		"imprecise_up_time",
		"voted_gov_proposal",
		"attention",
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
		validator.TotalSignedBlock,
		validator.TotalActiveBlock,
		validatorsView.rdb.BFton(validator.ImpreciseUpTime),
		validatorsView.rdb.Bton(validator.VotedGovProposal),
		validator.Attention,
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
		min_self_delegation = EXCLUDED.min_self_delegation,
		total_signed_block = EXCLUDED.total_signed_block,
		total_active_block = EXCLUDED.total_active_block,
		imprecise_up_time = EXCLUDED.imprecise_up_time,
		voted_gov_proposal = EXCLUDED.voted_gov_proposal,
		attention = EXCLUDED.attention
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

func (validatorsView *ValidatorsView) Insert(validator *ValidatorRow) error {
	sql, sqlArgs, buildStmtErr := validatorsView.rdb.StmtBuilder.Insert(
		VALIDATORS_TABLE_NAME,
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
		"total_signed_block",
		"total_active_block",
		"imprecise_up_time",
		"voted_gov_proposal",
		"attention",
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
		validator.TotalSignedBlock,
		validator.TotalActiveBlock,
		validatorsView.rdb.TypeConv.BFton(validator.ImpreciseUpTime),
		validatorsView.rdb.TypeConv.Bton(validator.VotedGovProposal),
		validator.Attention,
	).ToSql()
	if buildStmtErr != nil {
		return fmt.Errorf("error building validator insertion sql: %v: %w", buildStmtErr, rdb.ErrBuildSQLStmt)
	}

	result, execErr := validatorsView.rdb.Exec(sql, sqlArgs...)
	if execErr != nil {
		return fmt.Errorf("error inserting validator into the table: %v: %w", execErr, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting validator into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (validatorsView *ValidatorsView) Update(validator *ValidatorRow) error {
	sql, sqlArgs, err := validatorsView.rdb.StmtBuilder.Update(
		VALIDATORS_TABLE_NAME,
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
		"total_signed_block":         validator.TotalSignedBlock,
		"total_active_block":         validator.TotalActiveBlock,
		"imprecise_up_time":          validatorsView.rdb.TypeConv.BFton(validator.ImpreciseUpTime),
		"voted_gov_proposal":         validatorsView.rdb.TypeConv.Bton(validator.VotedGovProposal),
		"attention":                  validator.Attention,
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

func (validatorsView *ValidatorsView) UpdateAllValidatorUpTime(validators []ValidatorRow) error {

	pendingRowCount := 0
	totalRowCount := len(validators)

	sql := ""

	for i, validator := range validators {

		if pendingRowCount == 0 {

			sql = `UPDATE view_validators AS view SET
							total_signed_block = row.total_signed_block,
							total_active_block = row.total_active_block,
							imprecise_up_time = row.imprecise_up_time
						FROM (VALUES
						`
		}

		sql += fmt.Sprintf(
			"(%d, %d, %d, %s),\n",
			*validator.MaybeId,
			validator.TotalSignedBlock,
			validator.TotalActiveBlock,
			validator.ImpreciseUpTime.String(),
		)

		pendingRowCount += 1

		if pendingRowCount == 500 || i+1 == totalRowCount {

			sql = strings.TrimSuffix(sql, ",\n")

			sql += `) AS row(
								id, 
								total_signed_block,
								total_active_block, 
								imprecise_up_time
							)
							WHERE row.id = view.id;`

			result, err := validatorsView.rdb.Exec(sql)
			if err != nil {
				return fmt.Errorf("error updating validators up time into the table: %v: %w", err, rdb.ErrWrite)
			}
			if result.RowsAffected() != int64(pendingRowCount) {
				return fmt.Errorf("error updating validators up time into the table: wrong number of affected rows %d: %w", result.RowsAffected(), rdb.ErrWrite)
			}

			pendingRowCount = 0

		}

	}

	return nil
}

type ValidatorsListFilter struct {
	MaybeStatuses          []constants.Status
	MaybeRecentActiveBlock *ValidatorsListRecentActiveBlockFilter
}

type ValidatorsListRecentActiveBlockFilter struct {
	MaybeLowestActiveBlockHeight int64
}

type ValidatorsListOrder struct {
	MaybeStatus              *view.ORDER
	MaybePower               *view.ORDER
	MaybeCommission          *view.ORDER
	MaybeJoinedAtBlockHeight *view.ORDER
	MaybeImpreciseUpTime     *view.ORDER
}

func (validatorsView *ValidatorsView) ListAll(
	filter ValidatorsListFilter,
	order ValidatorsListOrder,
) ([]ValidatorRow, error) {
	orderClauses := make([]string, 0)

	if order.MaybeJoinedAtBlockHeight != nil {
		if *order.MaybeStatus == view.ORDER_ASC {
			statusOrder := "CASE UPPER(status) " +
				"WHEN 'BONDED' THEN 1 " +
				"WHEN 'UNBONDING' THEN 2 " +
				"WHEN 'JAILED' THEN 3 " +
				"WHEN 'INACTIVE' THEN 4 " +
				"WHEN 'UNBONDED' THEN 5 " +
				"ELSE 6 END"
			orderClauses = append(orderClauses, statusOrder)
		} else if *order.MaybeStatus == view.ORDER_DESC {
			statusOrder := "CASE UPPER(status) " +
				"WHEN 'UNBONDED' THEN 1 " +
				"WHEN 'JAILED' THEN 2 " +
				"WHEN 'UNBONDING' THEN 3 " +
				"WHEN 'INACTIVE' THEN 4 " +
				"WHEN 'UNBONDED' THEN 5 " +
				"ELSE 6 END"
			orderClauses = append(orderClauses, statusOrder)
		}
	}

	if order.MaybeImpreciseUpTime != nil {
		if *order.MaybeImpreciseUpTime == view.ORDER_ASC {
			orderClauses = append(orderClauses, "imprecise_up_time")
		} else if *order.MaybeImpreciseUpTime == view.ORDER_DESC {
			// DESC order
			orderClauses = append(orderClauses, "imprecise_up_time DESC")
		}
	}

	if order.MaybeCommission != nil {
		if *order.MaybeCommission == view.ORDER_ASC {
			orderClauses = append(orderClauses, "CAST(commission_rate AS DECIMAL)")
		} else if *order.MaybeCommission == view.ORDER_DESC {
			// DESC order
			orderClauses = append(orderClauses, "CAST(commission_rate AS DECIMAL) DESC")
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

	var whereClause sq.And

	var statusOrCondition sq.Or
	if filter.MaybeStatuses != nil {
		statusOrCondition = make(sq.Or, 0)
		for _, status := range filter.MaybeStatuses {
			statusOrCondition = append(statusOrCondition, sq.Eq{
				"status": status,
			})
		}
		whereClause = append(whereClause, statusOrCondition)
	}

	var stmtBuilder sq.SelectBuilder

	if filter.MaybeRecentActiveBlock != nil {
		if filter.MaybeRecentActiveBlock.MaybeLowestActiveBlockHeight < 0 {
			lowestBlockHeight := DEFAULT_LOWEST_ACTIVE_BLOCKS_BLOCK_HEIGHT
			filter.MaybeRecentActiveBlock.MaybeLowestActiveBlockHeight = lowestBlockHeight
		}

		stmtBuilder = validatorsView.rdb.StmtBuilder.Select(
			"id",
			fmt.Sprintf(
				"%s.operator_address",
				VALIDATORS_TABLE_NAME,
			),
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
			"total_signed_block",
			"total_active_block",
			"imprecise_up_time",
			"voted_gov_proposal",
			"attention",
			fmt.Sprintf(
				"COUNT(CASE WHEN %s.signed THEN 1 END) as recent_signed_blocks",
				VALIDATOR_ACTIVE_BLOCKS_TABLE_NAME,
			),
			fmt.Sprintf(
				"COUNT(%s.signed) as recent_active_blocks",
				VALIDATOR_ACTIVE_BLOCKS_TABLE_NAME,
			),
		).From(
			VALIDATORS_TABLE_NAME,
		).LeftJoin(
			fmt.Sprintf(
				"%s ON %s.operator_address=%s.operator_address AND %s.block_height > %d",
				VALIDATOR_ACTIVE_BLOCKS_TABLE_NAME,
				VALIDATORS_TABLE_NAME,
				VALIDATOR_ACTIVE_BLOCKS_TABLE_NAME,
				VALIDATOR_ACTIVE_BLOCKS_TABLE_NAME,
				filter.MaybeRecentActiveBlock.MaybeLowestActiveBlockHeight,
			),
		).GroupBy(
			fmt.Sprintf(
				"%s.operator_address, %s.id",
				VALIDATORS_TABLE_NAME,
				VALIDATORS_TABLE_NAME,
			),
		)
	} else {
		stmtBuilder = validatorsView.rdb.StmtBuilder.Select(
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
			"total_signed_block",
			"total_active_block",
			"imprecise_up_time",
			"voted_gov_proposal",
			"attention",
			"0 as recent_signed_blocks",
			"0 as recent_active_blocks",
		).From(
			VALIDATORS_TABLE_NAME,
		)
	}

	if whereClause != nil {
		stmtBuilder = stmtBuilder.Where(whereClause)
	}

	stmtBuilder = stmtBuilder.OrderBy(orderClauses...)

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

		impreciseUpTimeReader := validatorsView.rdb.TypeConv.NtobfReader()
		votedGovProposalReader := validatorsView.rdb.TypeConv.NtobReader()
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
			&validator.TotalSignedBlock,
			&validator.TotalActiveBlock,
			impreciseUpTimeReader.ScannableArg(),
			votedGovProposalReader.ScannableArg(),
			&validator.Attention,
			&validator.TotalRecentSignedBlocks,
			&validator.TotalRecentActiveBlocks,
		); err != nil {
			if errors.Is(err, rdb.ErrNoRows) {
				return nil, rdb.ErrNoRows
			}
			return nil, fmt.Errorf("error scanning validator row: %v: %w", err, rdb.ErrQuery)
		}

		impreciseUpTime, parseImpreciseUpTimeErr := impreciseUpTimeReader.Parse()
		if parseImpreciseUpTimeErr != nil {
			return nil, fmt.Errorf("error parsing imprecise up time: %v", parseImpreciseUpTimeErr)
		}
		validator.ImpreciseUpTime = impreciseUpTime

		votedGovProposal, votedGovProposalErr := votedGovProposalReader.Parse()
		if votedGovProposalErr != nil {
			return nil, fmt.Errorf("error parsing voted gov proposal: %v", votedGovProposalErr)
		}
		validator.VotedGovProposal = votedGovProposal

		validators = append(validators, validator)
	}

	return validators, nil
}

func (validatorsView *ValidatorsView) List(
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
				"WHEN 'BONDED' THEN 1 " +
				"WHEN 'UNBONDING' THEN 2 " +
				"WHEN 'JAILED' THEN 3 " +
				"WHEN 'INACTIVE' THEN 4 " +
				"WHEN 'UNBONDED' THEN 5 " +
				"ELSE 6 END"
			orderClauses = append(orderClauses, statusOrder)
		} else if *order.MaybeStatus == view.ORDER_DESC {
			statusOrder := "CASE UPPER(status) " +
				"WHEN 'UNBONDED' THEN 1 " +
				"WHEN 'JAILED' THEN 2 " +
				"WHEN 'UNBONDING' THEN 3 " +
				"WHEN 'INACTIVE' THEN 4 " +
				"WHEN 'BONDED' THEN 5 " +
				"ELSE 6 END"
			orderClauses = append(orderClauses, statusOrder)
		}
	}

	if order.MaybeImpreciseUpTime != nil {
		if *order.MaybeImpreciseUpTime == view.ORDER_ASC {
			orderClauses = append(orderClauses, "imprecise_up_time")
		} else if *order.MaybeImpreciseUpTime == view.ORDER_DESC {
			// DESC order
			orderClauses = append(orderClauses, "imprecise_up_time DESC")
		}
	}

	if order.MaybeCommission != nil {
		if *order.MaybeCommission == view.ORDER_ASC {
			orderClauses = append(orderClauses, "CAST(commission_rate AS DECIMAL)")
		} else if *order.MaybeCommission == view.ORDER_DESC {
			// DESC order
			orderClauses = append(orderClauses, "CAST(commission_rate AS DECIMAL) DESC")
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
		VALIDATORS_TABLE_NAME,
	).Offset(0).Limit(
		// disable G115
		uint64(pagination.OffsetParams().Offset()),
	)
	cumulativePowerStmtBuilder = cumulativePowerStmtBuilder.OrderBy(orderClauses...)

	var whereClause sq.And

	var statusOrCondition sq.Or
	if filter.MaybeStatuses != nil {
		statusOrCondition = make(sq.Or, 0)
		for _, status := range filter.MaybeStatuses {
			statusOrCondition = append(statusOrCondition, sq.Eq{
				"status": status,
			})
		}
		whereClause = append(whereClause, statusOrCondition)
	}

	if whereClause != nil {
		cumulativePowerStmtBuilder = cumulativePowerStmtBuilder.Where(whereClause)
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

	var stmtBuilder sq.SelectBuilder

	if filter.MaybeRecentActiveBlock != nil {
		if filter.MaybeRecentActiveBlock.MaybeLowestActiveBlockHeight < 0 {
			lowestBlockHeight := DEFAULT_LOWEST_ACTIVE_BLOCKS_BLOCK_HEIGHT
			filter.MaybeRecentActiveBlock.MaybeLowestActiveBlockHeight = lowestBlockHeight
		}
		stmtBuilder = validatorsView.rdb.StmtBuilder.Select(
			"id",
			fmt.Sprintf(
				"%s.operator_address",
				VALIDATORS_TABLE_NAME,
			),
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
			"total_signed_block",
			"total_active_block",
			"imprecise_up_time",
			"voted_gov_proposal",
			"attention",
			fmt.Sprintf(
				"COUNT(CASE WHEN %s.signed THEN 1 END) as recent_signed_blocks",
				VALIDATOR_ACTIVE_BLOCKS_TABLE_NAME,
			),
			"COUNT(id) as recent_active_blocks",
		).From(
			VALIDATORS_TABLE_NAME,
		).LeftJoin(
			fmt.Sprintf(
				"%s ON %s.operator_address=%s.operator_address AND %s.block_height > %d",
				VALIDATOR_ACTIVE_BLOCKS_TABLE_NAME,
				VALIDATORS_TABLE_NAME,
				VALIDATOR_ACTIVE_BLOCKS_TABLE_NAME,
				VALIDATOR_ACTIVE_BLOCKS_TABLE_NAME,
				filter.MaybeRecentActiveBlock.MaybeLowestActiveBlockHeight,
			),
		).GroupBy(
			fmt.Sprintf(
				"%s.id",
				VALIDATORS_TABLE_NAME,
			),
		)
	} else {
		stmtBuilder = validatorsView.rdb.StmtBuilder.Select(
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
			"total_signed_block",
			"total_active_block",
			"imprecise_up_time",
			"voted_gov_proposal",
			"attention",
			"0 as recent_signed_blocks",
			"0 as recent_active_blocks",
		).From(
			VALIDATORS_TABLE_NAME,
		)
	}

	if whereClause != nil {
		stmtBuilder = stmtBuilder.Where(whereClause)
	}

	stmtBuilder = stmtBuilder.OrderBy(orderClauses...)

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		validatorsView.rdb,
	).WithCustomTotalQueryFn(
		func(rdbHandle *rdb.Handle, _ sq.SelectBuilder) (int64, error) {
			totalStmtBuilder := validatorsView.rdb.StmtBuilder.Select(
				"COUNT(*)",
			).From(
				VALIDATORS_TABLE_NAME,
			)

			if whereClause != nil {
				totalStmtBuilder = totalStmtBuilder.Where(whereClause)
			}

			var total int64
			totalStmtSql, totalStmtSqlArgs, totalStmtBuilderErr := totalStmtBuilder.ToSql()
			if totalStmtBuilderErr != nil {
				return -1, fmt.Errorf("error building total select SQL: %v", err)
			}

			if err = rdbHandle.QueryRow(totalStmtSql, totalStmtSqlArgs...).Scan(&total); err != nil {
				return -1, fmt.Errorf("error executing total count select SQL: %v", err)
			}
			return total, nil
		},
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

		impreciseUpTimeReader := validatorsView.rdb.TypeConv.NtobfReader()
		votedGovProposalReader := validatorsView.rdb.TypeConv.NtobReader()
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
			&validator.TotalSignedBlock,
			&validator.TotalActiveBlock,
			impreciseUpTimeReader.ScannableArg(),
			votedGovProposalReader.ScannableArg(),
			&validator.Attention,
			&validator.TotalRecentSignedBlocks,
			&validator.TotalRecentActiveBlocks,
		); err != nil {
			if errors.Is(err, rdb.ErrNoRows) {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning validator row: %v: %w", err, rdb.ErrQuery)
		}

		impreciseUpTime, parseImpreciseUpTimeErr := impreciseUpTimeReader.Parse()
		if parseImpreciseUpTimeErr != nil {
			return nil, nil, fmt.Errorf("error parsing imprecise up time: %v", parseImpreciseUpTimeErr)
		}
		validator.ImpreciseUpTime = impreciseUpTime

		votedGovProposal, votedGovProposalErr := votedGovProposalReader.Parse()
		if votedGovProposalErr != nil {
			return nil, nil, fmt.Errorf("error parsing voted gov proposal: %v", votedGovProposalErr)
		}
		validator.VotedGovProposal = votedGovProposal

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

				powerPercentage.Text('f', 10),
				cumulativePowerPercentage.Text('f', 10),
			})
		}
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return validators, paginationResult, nil
}

func (validatorsView *ValidatorsView) totalPower() (*big.Float, error) {
	sql, _, _ := validatorsView.rdb.StmtBuilder.Select("power").From(VALIDATORS_TABLE_NAME).ToSql()
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

func (validatorsView *ValidatorsView) Search(
	keyword string,
	maybeLowestBlockHeight *int64,
) ([]ValidatorRow, error) {
	if maybeLowestBlockHeight == nil || *maybeLowestBlockHeight < 0 {
		lowestBlockHeight := DEFAULT_LOWEST_ACTIVE_BLOCKS_BLOCK_HEIGHT
		maybeLowestBlockHeight = &lowestBlockHeight
	}

	keyword = utils.AddressParse(keyword)
	sql, sqlArgs, err := validatorsView.rdb.StmtBuilder.Select(
		"id",
		fmt.Sprintf(
			"%s.operator_address",
			VALIDATORS_TABLE_NAME,
		),
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
		"total_signed_block",
		"total_active_block",
		"imprecise_up_time",
		"voted_gov_proposal",
		"attention",
		fmt.Sprintf(
			"COUNT(CASE WHEN %s.signed THEN 1 END) as recent_signed_blocks",
			VALIDATOR_ACTIVE_BLOCKS_TABLE_NAME,
		),
		fmt.Sprintf(
			"COUNT(%s.signed) as recent_active_blocks",
			VALIDATOR_ACTIVE_BLOCKS_TABLE_NAME,
		),
	).From(
		VALIDATORS_TABLE_NAME,
	).Where(
		fmt.Sprintf(
			"%s.operator_address = ? OR consensus_node_address = ? OR LOWER(moniker) LIKE ?",
			VALIDATORS_TABLE_NAME,
		),
		keyword, keyword, fmt.Sprintf("%%%s%%", keyword),
	).LeftJoin(
		fmt.Sprintf(
			"%s ON %s.operator_address=%s.operator_address AND %s.block_height > %d",
			VALIDATOR_ACTIVE_BLOCKS_TABLE_NAME,
			VALIDATORS_TABLE_NAME,
			VALIDATOR_ACTIVE_BLOCKS_TABLE_NAME,
			VALIDATOR_ACTIVE_BLOCKS_TABLE_NAME,
			*maybeLowestBlockHeight,
		),
	).GroupBy(
		fmt.Sprintf(
			"%s.operator_address, %s.id",
			VALIDATORS_TABLE_NAME,
			VALIDATORS_TABLE_NAME,
		),
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

		impreciseUpTimeReader := validatorsView.rdb.TypeConv.NtobfReader()
		votedGovProposalReader := validatorsView.rdb.TypeConv.NtobReader()
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
			&validator.TotalSignedBlock,
			&validator.TotalActiveBlock,
			impreciseUpTimeReader.ScannableArg(),
			votedGovProposalReader.ScannableArg(),
			&validator.Attention,
			&validator.TotalRecentSignedBlocks,
			&validator.TotalRecentActiveBlocks,
		); err != nil {
			if errors.Is(err, rdb.ErrNoRows) {
				return nil, rdb.ErrNoRows
			}
			return nil, fmt.Errorf("error scanning validator row: %v: %w", err, rdb.ErrQuery)
		}

		impreciseUpTime, parseImpreciseUpTimeErr := impreciseUpTimeReader.Parse()
		if parseImpreciseUpTimeErr != nil {
			return nil, fmt.Errorf("error parsing imprecise up time: %v", parseImpreciseUpTimeErr)
		}
		validator.ImpreciseUpTime = impreciseUpTime

		votedGovProposal, votedGovProposalErr := votedGovProposalReader.Parse()
		if votedGovProposalErr != nil {
			return nil, fmt.Errorf("error parsing voted gov proposal: %v", votedGovProposalErr)
		}
		validator.VotedGovProposal = votedGovProposal

		validators = append(validators, validator)
	}
	rowsResult.Close()

	return validators, nil
}

func (validatorsView *ValidatorsView) FindBy(
	identity ValidatorIdentity,
	maybeLowestBlockHeight *int64,
) (*ValidatorRow, error) {
	if maybeLowestBlockHeight == nil || *maybeLowestBlockHeight < 0 {
		lowestBlockHeight := DEFAULT_LOWEST_ACTIVE_BLOCKS_BLOCK_HEIGHT
		maybeLowestBlockHeight = &lowestBlockHeight
	}

	var err error

	selectStmtBuilder := validatorsView.rdb.StmtBuilder.Select(
		"id",
		fmt.Sprintf(
			"%s.operator_address",
			VALIDATORS_TABLE_NAME,
		),
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
		"total_signed_block",
		"total_active_block",
		"imprecise_up_time",
		"voted_gov_proposal",
		"attention",
		fmt.Sprintf(
			"COUNT(CASE WHEN %s.signed THEN 1 END) as recent_signed_blocks",
			VALIDATOR_ACTIVE_BLOCKS_TABLE_NAME,
		),
		fmt.Sprintf(
			"COUNT(%s.signed) as recent_active_blocks",
			VALIDATOR_ACTIVE_BLOCKS_TABLE_NAME,
		),
	).From(
		VALIDATORS_TABLE_NAME,
	)

	if identity.MaybeConsensusNodeAddress != nil {
		selectStmtBuilder = selectStmtBuilder.Where(
			"consensus_node_address = ?", *identity.MaybeConsensusNodeAddress,
		)
	}
	if identity.MaybeOperatorAddress != nil {
		selectStmtBuilder = selectStmtBuilder.Where(
			fmt.Sprintf(
				"%s.operator_address = ?",
				VALIDATORS_TABLE_NAME,
			),
			*identity.MaybeOperatorAddress,
		)
	}
	if identity.MaybeInitialDelegatorAddress != nil {
		selectStmtBuilder = selectStmtBuilder.Where("initial_delegator_address = ?", *identity.MaybeInitialDelegatorAddress)
	}

	selectStmtBuilder = selectStmtBuilder.LeftJoin(
		fmt.Sprintf(
			"%s ON %s.operator_address=%s.operator_address AND %s.block_height > %d",
			VALIDATOR_ACTIVE_BLOCKS_TABLE_NAME,
			VALIDATORS_TABLE_NAME,
			VALIDATOR_ACTIVE_BLOCKS_TABLE_NAME,
			VALIDATOR_ACTIVE_BLOCKS_TABLE_NAME,
			*maybeLowestBlockHeight,
		),
	).GroupBy(
		fmt.Sprintf(
			"%s.operator_address, %s.id",
			VALIDATORS_TABLE_NAME,
			VALIDATORS_TABLE_NAME,
		),
	).OrderBy("id DESC")

	sql, sqlArgs, err := selectStmtBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building validator selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	var validator ValidatorRow

	impreciseUpTimeReader := validatorsView.rdb.TypeConv.NtobfReader()
	votedGovProposalReader := validatorsView.rdb.TypeConv.NtobReader()
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
		&validator.TotalSignedBlock,
		&validator.TotalActiveBlock,
		impreciseUpTimeReader.ScannableArg(),
		votedGovProposalReader.ScannableArg(),
		&validator.Attention,
		&validator.TotalRecentSignedBlocks,
		&validator.TotalRecentActiveBlocks,
	); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return nil, rdb.ErrNoRows
		}
		return nil, fmt.Errorf("error scanning validator row: %v: %w", err, rdb.ErrQuery)
	}

	impreciseUpTime, parseImpreciseUpTimeErr := impreciseUpTimeReader.Parse()
	if parseImpreciseUpTimeErr != nil {
		return nil, fmt.Errorf("error parsing imprecise up time: %v", parseImpreciseUpTimeErr)
	}
	validator.ImpreciseUpTime = impreciseUpTime

	votedGovProposal, votedGovProposalErr := votedGovProposalReader.Parse()
	if votedGovProposalErr != nil {
		return nil, fmt.Errorf("error parsing voted gov proposal: %v", votedGovProposalErr)
	}
	validator.VotedGovProposal = votedGovProposal

	return &validator, nil
}

func (validatorsView *ValidatorsView) Count(filter CountFilter) (int64, error) {
	var count int64

	stmt := validatorsView.rdb.StmtBuilder.Select(
		"COUNT(*)",
	).From(
		VALIDATORS_TABLE_NAME,
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
	MaybeConsensusNodeAddress    *string
	MaybeOperatorAddress         *string
	MaybeInitialDelegatorAddress *string
}

type ValidatorRow struct {
	MaybeId                 *int64     `json:"-"`
	OperatorAddress         string     `json:"operatorAddress"`
	ConsensusNodeAddress    string     `json:"consensusNodeAddress"`
	InitialDelegatorAddress string     `json:"initialDelegatorAddress"`
	TendermintPubkey        string     `json:"tendermintPubkey"`
	TendermintAddress       string     `json:"tendermintAddress"`
	Status                  string     `json:"status"`
	Jailed                  bool       `json:"jailed"`
	JoinedAtBlockHeight     int64      `json:"joinedAtBlockHeight"`
	Power                   string     `json:"power"`
	Moniker                 string     `json:"moniker"`
	Identity                string     `json:"identity"`
	Website                 string     `json:"website"`
	SecurityContact         string     `json:"securityContact"`
	Details                 string     `json:"details"`
	CommissionRate          string     `json:"commissionRate"`
	CommissionMaxRate       string     `json:"commissionMaxRate"`
	CommissionMaxChangeRate string     `json:"commissionMaxChangeRate"`
	MinSelfDelegation       string     `json:"minSelfDelegation"`
	TotalSignedBlock        int64      `json:"totalSignedBlock"`
	TotalActiveBlock        int64      `json:"totalActiveBlock"`
	ImpreciseUpTime         *big.Float `json:"impreciseUpTime"`
	VotedGovProposal        *big.Int   `json:"votedGovProposal"`
	TotalRecentSignedBlocks int64      `json:"totalRecentSignedBlocks"`
	TotalRecentActiveBlocks int64      `json:"totalRecentActiveBlocks"`
	Attention               bool       `json:"attention"`
}

type ListValidatorRow struct {
	ValidatorRow

	PowerPercentage           string `json:"powerPercentage"`
	CumulativePowerPercentage string `json:"cumulativePowerPercentage"`
}
