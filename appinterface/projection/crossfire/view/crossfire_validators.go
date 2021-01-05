package view

import (
	"errors"
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/internal/utctime"
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

func (view *CrossfireValidators) LastJoinedBlockHeight(
	operatorAddress string,
	consensusNodeAddress string,
) (bool, int64, utctime.UTCTime, error) {
	var err error

	var sql string
	var sqlArgs []interface{}
	if sql, sqlArgs, err = view.rdb.StmtBuilder.Select(
		"joined_at_block_height",
		"joined_at_block_time",
	).From(
		TABLE_NAME,
	).Where(
		"operator_address = ? AND consensus_node_address = ?", operatorAddress, consensusNodeAddress,
	).ToSql(); err != nil {
		return false, int64(0), utctime.Now(), fmt.Errorf("error building validator existencen query sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	var joinedAtBlockHeight int64
	timeReader := view.rdb.NtotReader()
	if err = view.rdb.QueryRow(sql, sqlArgs...).Scan(
		&joinedAtBlockHeight,
		timeReader.ScannableArg(),
	); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return false, int64(0), utctime.Now(), nil
		}
		return false, int64(0), utctime.Now(), fmt.Errorf("error query validator existence: %v", err)
	}

	joinedAtBlockTime, parseErr := timeReader.Parse()
	if parseErr != nil {
		return false, int64(0), utctime.Now(), fmt.Errorf("error parsing block time: %v: %w", parseErr, rdb.ErrQuery)
	}

	return true, joinedAtBlockHeight, *joinedAtBlockTime, nil
}

func (view *CrossfireValidators) Upsert(validator *CrossfireValidatorRow) error {
	sql, sqlArgs, err := view.rdb.StmtBuilder.Insert(
		TABLE_NAME,
	).Columns(
		"operator_address",
		"consensus_node_address",
		"initial_delegator_address",
		"tendermint_pubkey",
		"tendermint_address",
		"status",
		"jailed",
		"joined_at_block_height",
		"joined_at_block_time",
		"moniker",
		"identity",
		"website",
		"security_contact",
		"details",
		"task_phase_1_node_setup",
		"task_phase_2_keep_node_active",
		"task_phase_2_proposal_vote",
		"task_phase_2_network_upgrade",
		"rank_task_phase_1_2_commitment_count",
		"rank_task_phase_3_commitment_count",
		"rank_task_highest_tx_sent",
	).Values(
		validator.OperatorAddress,
		validator.ConsensusNodeAddress,
		validator.InitialDelegatorAddress,
		validator.TendermintPubkey,
		validator.TendermintAddress,
		validator.Status,
		validator.Jailed,
		validator.JoinedAtBlockHeight,
		view.rdb.Tton(&validator.JoinedAtBlockTime),
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
		tendermint_pubkey = EXCLUDED.tendermint_pubkey,
		tendermint_address = EXCLUDED.tendermint_address,
		status = EXCLUDED.status,
		jailed = EXCLUDED.jailed,
		joined_at_block_height = EXCLUDED.joined_at_block_height,
		joined_at_block_time = EXCLUDED.joined_at_block_time,
		moniker = EXCLUDED.moniker,
		identity = EXCLUDED.identity,
		website = EXCLUDED.website,
		security_contact = EXCLUDED.security_contact,
		details = EXCLUDED.details,
		task_phase_1_node_setup = EXCLUDED.task_phase_1_node_setup,
		task_phase_2_keep_node_active = EXCLUDED.task_phase_2_keep_node_active,
		task_phase_2_proposal_vote = EXCLUDED.task_phase_2_proposal_vote,
		task_phase_2_network_upgrade = EXCLUDED.task_phase_2_network_upgrade,
		rank_task_phase_1_2_commitment_count = EXCLUDED.rank_task_phase_1_2_commitment_count,
		rank_task_phase_3_commitment_count = EXCLUDED.rank_task_phase_3_commitment_count,
		rank_task_highest_tx_sent = EXCLUDED.rank_task_highest_tx_sent
	`).ToSql()
	if err != nil {
		return fmt.Errorf("error building validator upsertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}
	result, err := view.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error upserting validator into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error upserting validator into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (view *CrossfireValidators) UpdateTask(
	taskColumnName string,
	status string,
	operatorAddress string,
	consensusNodeAddress string,
) error {
	sql, sqlArgs, err := view.rdb.StmtBuilder.Update(
		TABLE_NAME,
	).Set(
		taskColumnName, status,
	).Where(
		"operator_address = ? AND consensus_node_address = ?", operatorAddress, consensusNodeAddress,
	).ToSql()
	if err != nil {
		return fmt.Errorf("error building metrics update sql: %v", err)
	}

	var execResult rdb.ExecResult
	if execResult, err = view.rdb.Exec(sql, sqlArgs...); err != nil {
		return fmt.Errorf("error updating task: %v", err)
	}
	if execResult.RowsAffected() != 1 {
		return errors.New("error updating task: no rows updated")
	}

	return nil
}

func (view *CrossfireValidators) UpdateTaskForOperatorAddress(
	taskColumnName string,
	status string,
	operatorAddress string,
) error {
	sql, sqlArgs, err := view.rdb.StmtBuilder.Update(
		TABLE_NAME,
	).Set(
		taskColumnName, status,
	).Where(
		"operator_address = ?", operatorAddress,
	).ToSql()
	if err != nil {
		return fmt.Errorf("error building metrics update sql: %v", err)
	}

	var execResult rdb.ExecResult
	if execResult, err = view.rdb.Exec(sql, sqlArgs...); err != nil {
		return fmt.Errorf("error updating task: %v", err)
	}
	if execResult.RowsAffected() != 1 {
		return errors.New("error updating task: no rows updated")
	}

	return nil
}

func (view *CrossfireValidators) Count() (int64, error) {
	var count int64

	stmt := view.rdb.StmtBuilder.Select(
		"COUNT(*)",
	).From(
		TABLE_NAME,
	)

	sql, sqlArgs, err := stmt.ToSql()
	if err != nil {
		return int64(0), fmt.Errorf("error building validator count sql: %v: %w", err, rdb.ErrPrepare)
	}

	if err := view.rdb.QueryRow(sql, sqlArgs...).Scan(&count); err != nil {
		return int64(0), fmt.Errorf("error getting validators count: %v", err)
	}
	return count, nil
}

func (view *CrossfireValidators) List() ([]CrossfireValidatorRow, error) {
	stmtBuilder := view.rdb.StmtBuilder.Select(
		"id",
		"operator_address",
		"consensus_node_address",
		"initial_delegator_address",
		"tendermint_pubkey",
		"tendermint_address",
		"status",
		"jailed",
		"joined_at_block_height",
		"joined_at_block_time",
		"moniker",
		"identity",
		"website",
		"security_contact",
		"details",
		"task_phase_1_node_setup",
		"task_phase_2_keep_node_active",
		"task_phase_2_proposal_vote",
		"task_phase_2_network_upgrade",
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

	rowsResult, err := view.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, fmt.Errorf("error executing blocks select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	validators := make([]CrossfireValidatorRow, 0)
	for rowsResult.Next() {
		var validator CrossfireValidatorRow
		timeReader := view.rdb.NtotReader()
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
			timeReader.ScannableArg(),
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
		}

		blockTime, parseErr := timeReader.Parse()
		if parseErr != nil {
			return nil, fmt.Errorf("error parsing block time: %v: %w", parseErr, rdb.ErrQuery)
		}
		validator.JoinedAtBlockTime = *blockTime
		validators = append(validators, validator)
	}

	return validators, nil
}

func (view *CrossfireValidators) FindBy(identity CrossfireValidatorIdentity) (*CrossfireValidatorRow, error) {
	var err error

	selectStmtBuilder := view.rdb.StmtBuilder.Select(
		"id",
		"operator_address",
		"consensus_node_address",
		"initial_delegator_address",
		"tendermint_pubkey",
		"tendermint_address",
		"status",
		"jailed",
		"joined_at_block_height",
		"joined_at_block_time",
		"moniker",
		"identity",
		"website",
		"security_contact",
		"details",
		"task_phase_1_node_setup",
		"task_phase_2_keep_node_active",
		"task_phase_2_proposal_vote",
		"task_phase_2_network_upgrade",
		"rank_task_phase_1_2_commitment_count",
		"rank_task_phase_3_commitment_count",
		"rank_task_highest_tx_sent",
	).From(
		TABLE_NAME,
	).OrderBy("id DESC")
	if identity.MaybeConsensusNodeAddress != nil {
		selectStmtBuilder = selectStmtBuilder.Where(
			"consensus_node_address = ?", *identity.MaybeConsensusNodeAddress,
		)
	}
	if identity.MaybeOperatorAddress != nil {
		selectStmtBuilder = selectStmtBuilder.Where("operator_address = ?", *identity.MaybeOperatorAddress)
	}
	if identity.MaybeTendermintAddress != nil {
		selectStmtBuilder = selectStmtBuilder.Where("tendermint_address = ?", *identity.MaybeTendermintAddress)
	}

	sql, sqlArgs, err := selectStmtBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building validator selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	var validator CrossfireValidatorRow
	timeReader := view.rdb.NtotReader()
	if err = view.rdb.QueryRow(sql, sqlArgs...).Scan(
		&validator.MaybeId,
		&validator.OperatorAddress,
		&validator.ConsensusNodeAddress,
		&validator.InitialDelegatorAddress,
		&validator.TendermintPubkey,
		&validator.TendermintAddress,
		&validator.Status,
		&validator.Jailed,
		&validator.JoinedAtBlockHeight,
		timeReader.ScannableArg(),
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
	}

	joinedAtBlockTime, parseErr := timeReader.Parse()
	if parseErr != nil {
		return nil, fmt.Errorf("error parsing block time: %v: %w", parseErr, rdb.ErrQuery)
	}
	validator.JoinedAtBlockTime = *joinedAtBlockTime

	return &validator, nil
}

/**
UpdateTxSentRank : Updates the participating validator rank
**/
func (view *CrossfireValidators) UpdateTxSentRank(
	rank int,
	primaryAddress string,
	operatorAddress string,
) error {
	sql, sqlArgs, err := view.rdb.StmtBuilder.Update(
		TABLE_NAME,
	).Set(
		"rank_task_highest_tx_sent", rank,
	).Where(
		"operator_address = ? AND initial_delegator_address = ?",
		operatorAddress, primaryAddress,
	).ToSql()
	if err != nil {
		return fmt.Errorf("error building UpdateTxSentRank update sql: %v", err)
	}

	var execResult rdb.ExecResult
	if execResult, err = view.rdb.Exec(sql, sqlArgs...); err != nil {
		return fmt.Errorf("error updating Tx Sent task Rank: %v", err)
	}
	if execResult.RowsAffected() != 1 {
		return errors.New("error updating task: no rows updated")
	}

	return nil
}

type CrossfireValidatorIdentity struct {
	MaybeConsensusNodeAddress *string
	MaybeOperatorAddress      *string
	MaybeTendermintAddress    *string
}

type CrossfireValidatorRow struct {
	MaybeId                         *int64          `json:"-"`
	OperatorAddress                 string          `json:"operatorAddress"`
	ConsensusNodeAddress            string          `json:"consensusNodeAddress"`
	InitialDelegatorAddress         string          `json:"initialDelegatorAddress"`
	TendermintPubkey                string          `json:"tendermintPubkey"`
	TendermintAddress               string          `json:"tendermintAddress"`
	Status                          string          `json:"status"`
	Jailed                          bool            `json:"jailed"`
	JoinedAtBlockHeight             int64           `json:"joinedAtBlockHeight"`
	JoinedAtBlockTime               utctime.UTCTime `json:"joinedAtBlockTime"`
	Moniker                         string          `json:"moniker"`
	Identity                        string          `json:"identity"`
	Website                         string          `json:"website"`
	SecurityContact                 string          `json:"securityContact"`
	Details                         string          `json:"details"`
	TaskPhase1NodeSetup             string          `json:"taskPhase1NodeSetup"`
	TaskPhase2KeepNodeActive        string          `json:"taskPhase2KeepNodeActive"`
	TaskPhase2ProposalVote          string          `json:"taskPhase2ProposalVote"`
	TaskPhase2NetworkUpgrade        string          `json:"taskPhase2NetworkUpgrade"`
	RankTaskPhase1n2CommitmentCount int64           `json:"taskPhase1n2CommitmentCountRank"`
	RankTaskPhase3CommitmentCount   int64           `json:"taskPhase3CommitmentCountRank"`
	RankTaskHighestTxSent           int64           `json:"taskHighestTxSentRank"`
}
