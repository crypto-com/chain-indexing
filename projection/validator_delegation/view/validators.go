package view

import (
	"errors"
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/projection/validator_delegation/types"
	"github.com/crypto-com/chain-indexing/usecase/coin"
)

type Validators interface {
	Clone(previousHeight int64, currentHeight int64) error
	Insert(row ValidatorRow) error
	Update(row ValidatorRow) error
	Delete(row ValidatorRow) error
	FindByOperatorAddr(operatorAddress string, height int64) (ValidatorRow, bool, error)
	FindByConsensusNodeAddr(consensusNodeAddress string, height int64) (ValidatorRow, bool, error)
	List(height int64, pagination *pagination.Pagination) ([]ValidatorRow, *pagination.PaginationResult, error)
}

type ValidatorsView struct {
	rdb *rdb.Handle
}

func NewValidatorsView(handle *rdb.Handle) Validators {
	return &ValidatorsView{
		handle,
	}
}

func (view *ValidatorsView) Clone(previousHeight, currentHeight int64) error {

	sql, sqlErr := view.rdb.StmtBuilder.ReplacePlaceholders(`INSERT INTO view_vd_validators(
		height,
		operator_address,
		consensus_node_address,
		tendermint_address,
		status,
		jailed,
		power,
		unbonding_height,
		unbonding_time,
		tokens,
		shares,
		min_self_delegation
	) SELECT
		?,
		operator_address,
		consensus_node_address,
		tendermint_address,
		status,
		jailed,
		power,
		unbonding_height,
		unbonding_time,
		tokens,
		shares,
		min_self_delegation
	FROM view_vd_validators WHERE height = ?
	`)
	if sqlErr != nil {
		return fmt.Errorf("error building validator clone sql: %v: %w", sqlErr, rdb.ErrBuildSQLStmt)
	}
	sqlArgs := []interface{}{
		currentHeight,
		previousHeight,
	}

	if _, execErr := view.rdb.Exec(sql, sqlArgs...); execErr != nil {
		return fmt.Errorf("error cloning validator into the table: %v: %w", execErr, rdb.ErrWrite)
	}

	return nil
}

func (view *ValidatorsView) Insert(row ValidatorRow) error {

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Insert(
			"view_vd_validators",
		).
		Columns(
			"height",
			"operator_address",
			"consensus_node_address",
			"tendermint_address",
			"status",
			"jailed",
			"power",
			"unbonding_height",
			"unbonding_time",
			"tokens",
			"shares",
			"min_self_delegation",
		).
		Values(
			row.Height,
			row.OperatorAddress,
			row.ConsensusNodeAddress,
			row.TendermintAddress,
			row.Status,
			row.Jailed,
			row.Power,
			row.UnbondingHeight,
			view.rdb.Tton(&row.UnbondingTime),
			row.Tokens.String(),
			row.Shares.String(),
			row.MinSelfDelegation.String(),
		).
		ToSql()

	if err != nil {
		return fmt.Errorf("error building validator insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := view.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting validator into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting validator into the table: rows inserted: %v: %w", result.RowsAffected(), rdb.ErrWrite)
	}

	return nil
}

func (view *ValidatorsView) Update(row ValidatorRow) error {

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Update(
			"view_vd_validators",
		).
		SetMap(map[string]interface{}{
			"status":              row.Status,
			"jailed":              row.Jailed,
			"power":               row.Power,
			"unbonding_height":    row.UnbondingHeight,
			"unbonding_time":      view.rdb.Tton(&row.UnbondingTime),
			"tokens":              row.Tokens.String(),
			"shares":              row.Shares.String(),
			"min_self_delegation": row.MinSelfDelegation.String(),
		}).
		Where(
			"operator_address = ? AND height = ?",
			row.OperatorAddress,
			row.Height,
		).
		ToSql()

	if err != nil {
		return fmt.Errorf("error building validator update sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := view.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error updating validator into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error updating validator into the table: row updated: %v: %w", result.RowsAffected(), rdb.ErrWrite)
	}

	return nil
}

func (view *ValidatorsView) Delete(row ValidatorRow) error {

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Delete(
			"view_vd_validators",
		).
		Where(
			"operator_address = ? AND height = ?",
			row.OperatorAddress,
			row.Height,
		).
		ToSql()
	if err != nil {
		return fmt.Errorf("error building validator deletion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := view.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error deleting validator from the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error deleting validator into the table: row deleted: %v: %w", result.RowsAffected(), rdb.ErrWrite)
	}

	return nil
}

func (view *ValidatorsView) FindByOperatorAddr(operatorAddress string, height int64) (ValidatorRow, bool, error) {

	row, found, err := view.findBy(&operatorAddress, nil, height)
	if err != nil {
		return ValidatorRow{}, false, fmt.Errorf("error finding validator by OperatorAddr: %v", err)
	}

	return row, found, nil
}

func (view *ValidatorsView) FindByConsensusNodeAddr(consensusNodeAddress string, height int64) (ValidatorRow, bool, error) {

	row, found, err := view.findBy(nil, &consensusNodeAddress, height)
	if err != nil {
		return ValidatorRow{}, false, fmt.Errorf("error finding validator by OperatorAddr: %v", err)
	}

	return row, found, nil
}

func (view *ValidatorsView) findBy(
	maybeOperatorAddress *string,
	maybeConsensusNodeAddress *string,
	height int64,
) (ValidatorRow, bool, error) {

	stmtBuilder := view.rdb.StmtBuilder.
		Select(
			"operator_address",
			"consensus_node_address",
			"tendermint_address",
			"status",
			"jailed",
			"power",
			"unbonding_height",
			"unbonding_time",
			"tokens",
			"shares",
			"min_self_delegation",
		).
		From("view_vd_validators")

	if maybeOperatorAddress != nil {
		stmtBuilder = stmtBuilder.Where(
			"operator_address = ? AND height = ?",
			*maybeOperatorAddress,
			height,
		)
	} else if maybeConsensusNodeAddress != nil {
		stmtBuilder = stmtBuilder.Where(
			"consensus_node_address = ? AND height = ?",
			*maybeConsensusNodeAddress,
			height,
		)
	} else {
		return ValidatorRow{}, false, errors.New("no address provided for finding validator")
	}

	sql, sqlArgs, err := stmtBuilder.ToSql()
	if err != nil {
		return ValidatorRow{}, false, fmt.Errorf("error building select validator sql: %v: %w", err, rdb.ErrPrepare)
	}

	var validator ValidatorRow
	validator.Height = height

	unbondingTimeReader := view.rdb.NtotReader()
	var tokensInString string
	var sharesInString string
	var minSelfDelegationInString string
	if err = view.rdb.QueryRow(sql, sqlArgs...).Scan(
		&validator.OperatorAddress,
		&validator.ConsensusNodeAddress,
		&validator.TendermintAddress,
		&validator.Status,
		&validator.Jailed,
		&validator.Power,
		&validator.UnbondingHeight,
		unbondingTimeReader.ScannableArg(),
		&tokensInString,
		&sharesInString,
		&minSelfDelegationInString,
	); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			// When the row is not found, do not return error
			return ValidatorRow{}, false, nil
		}
		return ValidatorRow{}, false, fmt.Errorf("error scanning validator: %v: %w", err, rdb.ErrQuery)
	}

	unbondingTime, parseErr := unbondingTimeReader.Parse()
	if parseErr != nil {
		return ValidatorRow{}, false, fmt.Errorf("error parsing unbondingTime: %v: %w", parseErr, rdb.ErrQuery)
	}
	validator.UnbondingTime = *unbondingTime

	var ok bool
	validator.Tokens, ok = coin.NewIntFromString(tokensInString)
	if !ok {
		return ValidatorRow{}, false, fmt.Errorf("error parsing tokens to coin.Int: %w", rdb.ErrQuery)
	}

	validator.Shares, err = coin.NewDecFromStr(sharesInString)
	if err != nil {
		return ValidatorRow{}, false, fmt.Errorf("error parsing shares to coin.Dec: %v: %w", err, rdb.ErrQuery)
	}

	validator.MinSelfDelegation, ok = coin.NewIntFromString(minSelfDelegationInString)
	if !ok {
		return ValidatorRow{}, false, fmt.Errorf("error parsing min_self_delegation to coin.Int: %w", rdb.ErrQuery)
	}

	return validator, true, nil
}

func (view *ValidatorsView) List(
	height int64,
	pagination *pagination.Pagination,
) (
	[]ValidatorRow,
	*pagination.PaginationResult,
	error,
) {

	stmtBuilder := view.rdb.StmtBuilder.
		Select(
			"operator_address",
			"consensus_node_address",
			"tendermint_address",
			"status",
			"jailed",
			"power",
			"unbonding_height",
			"unbonding_time",
			"tokens",
			"shares",
			"min_self_delegation",
		).
		From(
			"view_vd_validators",
		).
		Where(
			"height = ?",
			height,
		)

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		view.rdb,
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building validators select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := view.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing validators select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	validators := make([]ValidatorRow, 0)
	for rowsResult.Next() {
		var validator ValidatorRow
		validator.Height = height

		unbondingTimeReader := view.rdb.NtotReader()
		var tokensInString string
		var sharesInString string
		var minSelfDelegationInString string
		if err = rowsResult.Scan(
			&validator.OperatorAddress,
			&validator.ConsensusNodeAddress,
			&validator.TendermintAddress,
			&validator.Status,
			&validator.Jailed,
			&validator.Power,
			&validator.UnbondingHeight,
			unbondingTimeReader.ScannableArg(),
			&tokensInString,
			&sharesInString,
			&minSelfDelegationInString,
		); err != nil {
			return nil, nil, fmt.Errorf("error scanning delegation row: %v: %w", err, rdb.ErrQuery)
		}

		unbondingTime, parseErr := unbondingTimeReader.Parse()
		if parseErr != nil {
			return nil, nil, fmt.Errorf("error parsing unbondingTime: %v: %w", parseErr, rdb.ErrQuery)
		}
		validator.UnbondingTime = *unbondingTime

		var ok bool
		validator.Tokens, ok = coin.NewIntFromString(tokensInString)
		if !ok {
			return nil, nil, fmt.Errorf("error parsing tokens to coin.Int: %w", rdb.ErrQuery)
		}

		validator.Shares, err = coin.NewDecFromStr(sharesInString)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing shares to coin.Dec: %v: %w", err, rdb.ErrQuery)
		}

		validator.MinSelfDelegation, ok = coin.NewIntFromString(minSelfDelegationInString)
		if !ok {
			return nil, nil, fmt.Errorf("error parsing min_self_delegation to coin.Int: %w", rdb.ErrQuery)
		}

		validators = append(validators, validator)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return validators, paginationResult, nil
}

// Notes on `Validator.Status`:
//
// - MsgCreateValidator, CreateGenesisValidator: Validator.Status = Unbonded
// - PowerChanged event:
//   - power > 0: Validator.Status = Bonded
//   - power = 0: Validator.Status = Unbonding, UnbondingValidator entry created
// - UnbondingValidator complete Unbonding period: Validator.Status = Unbonded

// NOTES:
// - UNIQUE(operatorAddress, height)
// - UNIQUE(consensusNodeAddress, height)
// - UNIQUE(tendermintAddress, height)
// - INDEX(height)
type ValidatorRow struct {
	Height int64 `json:"height"`

	OperatorAddress      string `json:"operatorAddress"`
	ConsensusNodeAddress string `json:"consensusNodeAddress"`
	TendermintAddress    string `json:"tendermintAddress"`

	Status types.ValidatorStatus `json:"status"`
	Jailed bool                  `json:"jailed"`
	Power  string                `json:"power"`

	// `UnbondingHeight` and `UnbondingTime` only useful when `Status` is `Unbonding`
	// The height start the Unbonding
	UnbondingHeight int64 `json:"UnbondingHeight"`
	// The time when Unbonding is finished
	UnbondingTime utctime.UTCTime `json:"UnbondingTime"`

	Tokens            coin.Int `json:"tokens"`
	Shares            coin.Dec `json:"shares"`
	MinSelfDelegation coin.Int `json:"minSelfDelegation"`
}

func (v *ValidatorRow) IsBonded() bool {
	return v.Status == types.BONDED
}

func (v *ValidatorRow) IsUnbonded() bool {
	return v.Status == types.UNBONDED
}

func (v *ValidatorRow) IsUnbonding() bool {
	return v.Status == types.UNBONDING
}