package view

import (
	"errors"
	"fmt"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

const TABLE_NAME = "view_crossfire_validators"

type CrossfireValidators struct {
	rdb *rdb.Handle
}

func NewCrossfireValidators(handle *rdb.Handle) *CrossfireValidators {
	return &CrossfireValidators{
		handle,
	}
}

func (validatorsView *CrossfireValidators) LastJoinedBlockHeight(
	operatorAddress string,
	consensusNodeAddress string,
) (bool, int64, error) {
	var err error

	var sql string
	var sqlArgs []interface{}
	if sql, sqlArgs, err = validatorsView.rdb.StmtBuilder.Select(
		"joined_at_block_height",
	).From(
		TABLE_NAME,
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

func (validatorsView *CrossfireValidators) Upsert(validator *CrossfireValidatorRow) error {
	sql, sqlArgs, err := validatorsView.rdb.StmtBuilder.Insert(
		TABLE_NAME,
	).Columns(
		"operator_address",
		"consensus_node_address",
		"initial_delegator_address",
		"status",
		"jailed",
		"joined_at_block_height",
		"moniker",
		"identity",
		"website",
		"security_contact",
		"details",
		"task_phase_1_node_setup",
		"task_phase_1_block_valid_commit",
		"task_phase_2_keep_node_active",
		"task_phase_2_proposal_vote",
		"task_phase_2_network_upgrade",
		"task_phase_2_network_upgrade_block_commit",
		"rank_task_phase_1_2_commitment_count",
		"rank_task_phase_3_commitment_count",
		"rank_task_highest_tx_sent",
	).Values(
		validator.OperatorAddress,
		validator.ConsensusNodeAddress,
		validator.InitialDelegatorAddress,
		validator.Status,
		validator.Jailed,
		validator.JoinedAtBlockHeight,
		validator.Moniker,
		validator.Identity,
		validator.Website,
		validator.SecurityContact,
		validator.Details,
		validator.TaskPhase1NodeSetup,
		validator.TaskPhase2KeepNodeActive,
		validator.TaskPhase2ProposalVote,
		validator.TaskPhase2NetworkUpgrade,
		validator.RankTaskPhase1n2CommitmentCount,
		validator.RankTaskPhase3CommitmentCount,
		validator.RankTaskHighestTxSent,
	).Suffix(`ON CONFLICT (operator_address, consensus_node_address) DO UPDATE SET
		initial_delegator_address = EXCLUDED.initial_delegator_address,
		status = EXCLUDED.status,
		jailed = EXCLUDED.jailed,
		joined_at_block_height = EXCLUDED.joined_at_block_height,
		moniker = EXCLUDED.moniker,
		identity = EXCLUDED.identity,
		website = EXCLUDED.website,
		security_contact = EXCLUDED.security_contact,
		details = EXCLUDED.details,
		task_phase_1_node_setup = EXCLUDED.task_phase_1_node_setup,
		task_phase_1_block_valid_commit = EXCLUDED.task_phase_1_block_valid_commit,
		task_phase_2_keep_node_active = EXCLUDED.task_phase_2_keep_node_active,
		task_phase_2_proposal_vote = EXCLUDED.task_phase_2_proposal_vote,
		task_phase_2_network_upgrade = EXCLUDED.task_phase_2_network_upgrade,
		task_phase_2_network_upgrade_block_commit = EXCLUDED.task_phase_2_network_upgrade_block_commit,
		rank_task_phase_1_2_commitment_count = EXCLUDED.rank_task_phase_1_2_commitment_count,
		rank_task_phase_3_commitment_count = EXCLUDED.rank_task_phase_3_commitment_count,
		rank_task_highest_tx_sent = EXCLUDED.rank_task_highest_tx_sent
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

func (validatorsView *CrossfireValidators) Count() (int64, error) {
	var count int64

	stmt := validatorsView.rdb.StmtBuilder.Select(
		"COUNT(*)",
	).From(
		TABLE_NAME,
	)

	sql, sqlArgs, err := stmt.ToSql()
	if err != nil {
		return int64(0), fmt.Errorf("error building validator count sql: %v: %w", err, rdb.ErrPrepare)
	}

	if err := validatorsView.rdb.QueryRow(sql, sqlArgs...).Scan(&count); err != nil {
		return int64(0), fmt.Errorf("error getting validators count: %v", err)
	}
	return count, nil
}

func (validatorsView *CrossfireValidators) List() ([]CrossfireValidatorRow, error) {
	stmtBuilder := validatorsView.rdb.StmtBuilder.Select(
		"id",
		"operator_address",
		"consensus_node_address",
		"initial_delegator_address",
		"status",
		"jailed",
		"joined_at_block_height",
		"moniker",
		"identity",
		"website",
		"security_contact",
		"details",
		"task_phase_1_node_setup",
		"task_phase_1_block_valid_commit",
		"task_phase_2_keep_node_active",
		"task_phase_2_proposal_vote",
		"task_phase_2_network_upgrade",
		"task_phase_2_network_upgrade_block_commit",
		"rank_task_phase_1_2_commitment_count",
		"rank_task_phase_3_commitment_count",
		"rank_task_highest_tx_sent",
	).From(
		TABLE_NAME,
	)

	sql, sqlArgs, err := stmtBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building blocks select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := validatorsView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, fmt.Errorf("error executing blocks select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	validators := make([]CrossfireValidatorRow, 0)
	for rowsResult.Next() {
		var validator CrossfireValidatorRow
		if err = rowsResult.Scan(
			&validator.MaybeId,
			&validator.OperatorAddress,
			&validator.ConsensusNodeAddress,
			&validator.InitialDelegatorAddress,
			&validator.Status,
			&validator.Jailed,
			&validator.JoinedAtBlockHeight,
			&validator.Moniker,
			&validator.Identity,
			&validator.Website,
			&validator.SecurityContact,
			&validator.Details,
			&validator.TaskPhase1NodeSetup,
			&validator.TaskPhase2KeepNodeActive,
			&validator.TaskPhase2ProposalVote,
			&validator.TaskPhase2NetworkUpgrade,
			&validator.RankTaskPhase1n2CommitmentCount,
			&validator.RankTaskPhase3CommitmentCount,
			&validator.RankTaskHighestTxSent,
		); err != nil {
			if errors.Is(err, rdb.ErrNoRows) {
				return nil, rdb.ErrNoRows
			}
			return nil, fmt.Errorf("error scanning crossfire validator row: %v: %w", err, rdb.ErrQuery)
		} else {
			validators = append(validators, validator)
		}
	}

	return validators, nil
}

type CrossfireValidatorRow struct {
	MaybeId                             *int64 `json:"-"`
	OperatorAddress                     string `json:"operatorAddress"`
	ConsensusNodeAddress                string `json:"consensusNodeAddress"`
	InitialDelegatorAddress             string `json:"initialDelegatorAddress"`
	Status                              string `json:"status"`
	Jailed                              bool   `json:"jailed"`
	JoinedAtBlockHeight                 int64  `json:"joinedAtBlockHeight"`
	Moniker                             string `json:"moniker"`
	Identity                            string `json:"identity"`
	Website                             string `json:"website"`
	SecurityContact                     string `json:"securityContact"`
	Details                             string `json:"details"`
	TaskPhase1NodeSetup                 string `json:"taskPhase1NodeSetup"`
	TaskPhase2KeepNodeActive            string `json:"taskPhase2KeepNodeActive"`
	TaskPhase2ProposalVote              string `json:"taskPhase2ProposalVote"`
	TaskPhase2NetworkUpgrade            string `json:"taskPhase2NetworkUpgrade"`
	RankTaskPhase1n2CommitmentCount     int64  `json:"taskPhase1n2CommitmentCountRank"`
	RankTaskPhase3CommitmentCount       int64  `json:"taskPhase3CommitmentCountRank"`
	RankTaskHighestTxSent               int64  `json:"taskHighestTxSentRank"`
}
