package view

import (
	"errors"
	"fmt"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

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
		"view_crossfire_validators",
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

func (validatorsView *CrossfireValidators) Upsert(validator *ValidatorRow) error {
	sql, sqlArgs, err := validatorsView.rdb.StmtBuilder.Insert(
		"view_crossfire_validators",
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
		"phase_0_task_registration",
		"phase_1_task_node_setup",
		"phase_1_task_block_valid_commit",
		"phase_2_task_keep_node_active",
		"phase_2_task_proposal_vote",
		"phase_2_task_network_upgrade",
		"phase_2_task_network_upgrade_block_commit",
		"phase_1_2_task_commitment_count_rank",
		"phase_3_task_commitment_count_rank",
		"task_highest_sequence_rank",
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
		validator.Phase0TaskRegistration,
		validator.Phase1TaskNodeSetup,
		validator.Phase1TaskBlockValidCommit,
		validator.Phase2TaskKeepNodeActive,
		validator.Phase2TaskProposalVote,
		validator.Phase2TaskNetworkUpgrade,
		validator.Phase2TaskNetworkUpgradeBlockCommit,
		validator.Phase1n2TaskCommitmentCountRank,
		validator.Phase3TaskCommitmentCountRank,
		validator.TaskHighestSequenceRank,
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
		phase_0_task_registration = EXCLUDED.phase_0_task_registration,
		phase_1_task_node_setup = EXCLUDED.phase_1_task_node_setup,
		phase_1_task_block_valid_commit = EXCLUDED.phase_1_task_block_valid_commit,
		phase_2_task_keep_node_active = EXCLUDED.phase_2_task_keep_node_active,
		phase_2_task_proposal_vote = EXCLUDED.phase_2_task_proposal_vote,
		phase_2_task_network_upgrade = EXCLUDED.phase_2_task_network_upgrade,
		phase_2_task_network_upgrade_block_commit = EXCLUDED.phase_2_task_network_upgrade_block_commit,
		phase_1_2_task_commitment_count_rank = EXCLUDED.phase_1_2_task_commitment_count_rank,
		phase_3_task_commitment_count_rank = EXCLUDED.phase_3_task_commitment_count_rank,
		task_highest_sequence_rank = EXCLUDED.task_highest_sequence_rank
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
		"view_crossfire_validators",
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

type ValidatorRow struct {
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
	Phase0TaskRegistration              string `json:"phase0TaskRegistration"`
	Phase1TaskNodeSetup                 string `json:"phase1TaskNodeSetup"`
	Phase1TaskBlockValidCommit          string `json:"phase1TaskBlockValidCommit"`
	Phase2TaskKeepNodeActive            string `json:"phase2TaskKeepNodeActive"`
	Phase2TaskProposalVote              string `json:"phase2TaskProposalVote"`
	Phase2TaskNetworkUpgrade            string `json:"phase2TaskNetworkUpgrade"`
	Phase2TaskNetworkUpgradeBlockCommit string `json:"phase2TaskNetworkUpgradeBlockCommit"`
	Phase1n2TaskCommitmentCountRank     int64  `json:"phase1n2TaskCommitmentCountRank"`
	Phase3TaskCommitmentCountRank       int64  `json:"phase3TaskCommitmentCountRank"`
	TaskHighestSequenceRank             int64  `json:"taskHighestSequenceRank"`
}
