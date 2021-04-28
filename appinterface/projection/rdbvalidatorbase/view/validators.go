package view

import (
	"errors"
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

// A "Validators" compatible table should have the follow table schema
// | Column                    | Type      | Constraint           |
// | ------------------------- | --------- | -------------------- |
// | id                        | BIGSERIAL | PRIMARY KEY          |
// | consensus_node_address    | VARCHAR   | UNIQUE1              |
// | operator_address          | VARCHAR   | NOT NULL, UNIQUE1    |
// | initial_delegator_address | VARCHAR   | NOT NULL             |
// | tendermint_pubkey         | VARCHAR   | NOT NULL             |
// | tendermint_address        | VARCHAR   | NOT NULL             |
// | moniker                   | VARCHAR   | NOT NULL             |

// A generic validator view
type Validators struct {
	rdb *rdb.Handle

	tableName string
}

func NewValidators(handle *rdb.Handle, tableName string) *Validators {
	return &Validators{
		handle,

		tableName,
	}
}

func (validatorsView *Validators) Upsert(validator *ValidatorRow) error {
	sql, sqlArgs, err := validatorsView.rdb.StmtBuilder.Insert(
		validatorsView.tableName,
	).Columns(
		"consensus_node_address",
		"operator_address",
		"initial_delegator_address",
		"tendermint_pubkey",
		"tendermint_address",
		"moniker",
	).Values(
		validator.ConsensusNodeAddress,
		validator.OperatorAddress,
		validator.InitialDelegatorAddress,
		validator.TendermintPubkey,
		validator.TendermintAddress,
		validator.Moniker,
	).Suffix(`ON CONFLICT (operator_address, consensus_node_address) DO UPDATE SET
		moniker = EXCLUDED.moniker
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
	sql, sqlArgs, err := validatorsView.rdb.StmtBuilder.Insert(
		validatorsView.tableName,
	).Columns(
		"consensus_node_address",
		"operator_address",
		"initial_delegator_address",
		"tendermint_pubkey",
		"tendermint_address",
		"moniker",
	).Values(
		validator.ConsensusNodeAddress,
		validator.OperatorAddress,
		validator.InitialDelegatorAddress,
		validator.TendermintPubkey,
		validator.TendermintAddress,
		validator.Moniker,
	).ToSql()
	if err != nil {
		return fmt.Errorf("error building validator insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := validatorsView.rdb.Exec(sql, sqlArgs...)
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
		validatorsView.tableName,
	).SetMap(map[string]interface{}{
		"moniker": validator.Moniker,
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

func (validatorsView *Validators) ListAll() ([]ValidatorRow, error) {
	orderClauses := make([]string, 0)

	stmtBuilder := validatorsView.rdb.StmtBuilder.Select(
		"id",
		"consensus_node_address",
		"operator_address",
		"initial_delegator_address",
		"tendermint_pubkey",
		"tendermint_address",
		"moniker",
	).From(
		validatorsView.tableName,
	)
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
		if err = rowsResult.Scan(
			&validator.MaybeId,
			&validator.ConsensusNodeAddress,
			&validator.OperatorAddress,
			&validator.InitialDelegatorAddress,
			&validator.TendermintPubkey,
			&validator.TendermintAddress,
			&validator.Moniker,
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

func (validatorsView *Validators) FindLastBy(identity ValidatorIdentity) (*ValidatorRow, error) {
	var err error

	selectStmtBuilder := validatorsView.rdb.StmtBuilder.Select(
		"id",
		"consensus_node_address",
		"operator_address",
		"initial_delegator_address",
		"tendermint_pubkey",
		"tendermint_address",
		"moniker",
	).From(
		validatorsView.tableName,
	).OrderBy("id DESC")
	if identity.MaybeConsensusNodeAddress != nil {
		selectStmtBuilder = selectStmtBuilder.Where(
			"consensus_node_address = ?", *identity.MaybeConsensusNodeAddress,
		)
	}
	if identity.MaybeOperatorAddress != nil {
		selectStmtBuilder = selectStmtBuilder.Where("operator_address = ?", *identity.MaybeOperatorAddress)
	}
	if identity.MaybeInititalDelegatorAddress != nil {
		selectStmtBuilder = selectStmtBuilder.Where("initial_delegator_address = ?", *identity.MaybeInititalDelegatorAddress)
	}

	sql, sqlArgs, err := selectStmtBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building validator selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	var validator ValidatorRow
	if err = validatorsView.rdb.QueryRow(sql, sqlArgs...).Scan(
		&validator.MaybeId,
		&validator.ConsensusNodeAddress,
		&validator.OperatorAddress,
		&validator.InitialDelegatorAddress,
		&validator.TendermintPubkey,
		&validator.TendermintAddress,
		&validator.Moniker,
	); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return nil, rdb.ErrNoRows
		}
		return nil, fmt.Errorf("error scanning validator row: %v: %w", err, rdb.ErrQuery)
	}

	return &validator, nil
}

func (validatorsView *Validators) Count() (int64, error) {
	var count int64

	sql, sqlArgs, err := validatorsView.rdb.StmtBuilder.Select(
		"COUNT(*)",
	).From(
		validatorsView.tableName,
	).ToSql()
	if err != nil {
		return int64(0), fmt.Errorf("error building validator count sql: %v: %w", err, rdb.ErrPrepare)
	}

	if err := validatorsView.rdb.QueryRow(sql, sqlArgs...).Scan(&count); err != nil {
		return int64(0), fmt.Errorf("error getting validators count: %v", err)
	}
	return count, nil
}

type ValidatorIdentity struct {
	MaybeConsensusNodeAddress     *string
	MaybeOperatorAddress          *string
	MaybeInititalDelegatorAddress *string
}

type ValidatorRow struct {
	MaybeId                 *int64 `json:"-"`
	ConsensusNodeAddress    string `json:"consensusNodeAddress"`
	OperatorAddress         string `json:"operatorAddress"`
	InitialDelegatorAddress string `json:"initialDelegatorAddress"`
	TendermintPubkey        string `json:"tendermintPubkey"`
	TendermintAddress       string `json:"tendermintAddress"`
	Moniker                 string `json:"moniker"`
}
