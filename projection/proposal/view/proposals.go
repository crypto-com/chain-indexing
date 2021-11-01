package view

import (
	"errors"
	"fmt"
	"math/big"

	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/json"
	"github.com/crypto-com/chain-indexing/internal/utctime"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	jsoniter "github.com/json-iterator/go"
)

const PROPOSALS_TABLE_NAME = "view_proposals"
const PARAMS_TABLE_NAME = "view_proposal_params"
const VALIDATORS_TABLE_NAME = "view_proposal_validators"

type Proposals interface {
	Insert(proposal *ProposalRow) error
	IncrementTotalVoteBy(proposalId uint64, voteToAdd *big.Int) error
	Update(row *ProposalRow) error
	FindById(proposalId string) (*ProposalWithMonikerRow, error)
	List(
		filter ProposalListFilter,
		order ProposalListOrder,
		pagination *pagination_interface.Pagination,
	) (
		[]ProposalWithMonikerRow,
		*pagination_interface.PaginationResult,
		error,
	)
}

type ProposalsView struct {
	rdb *rdb.Handle
}

func NewProposalsView(handle *rdb.Handle) Proposals {
	return &ProposalsView{
		handle,
	}
}

func (proposalView *ProposalsView) Insert(proposal *ProposalRow) error {
	var err error

	var dataJSON string
	if dataJSON, err = jsoniter.MarshalToString(proposal.Data); err != nil {
		return fmt.Errorf("error JSON marshalling proposal data for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	sql, sqlArgs, err := proposalView.rdb.StmtBuilder.Insert(
		PROPOSALS_TABLE_NAME,
	).Columns(
		"proposal_id",
		"title",
		"description",
		"type",
		"status",
		"proposer_address",
		"maybe_proposer_operator_address",
		"data",
		"initial_deposit",
		"total_deposit",
		"total_vote",
		"transaction_hash",
		"submit_block_height",
		"submit_time",
		"deposit_end_time",
		"maybe_voting_start_time",
		"maybe_voting_end_block_height",
		"maybe_voting_end_time",
	).Values(
		proposal.ProposalId,
		proposal.Title,
		proposal.Description,
		proposal.Type,
		proposal.Status,
		proposal.ProposerAddress,
		proposal.MaybeProposerOperatorAddress,
		dataJSON,
		json.MustMarshalToString(proposal.InitialDeposit),
		json.MustMarshalToString(proposal.TotalDeposit),
		proposalView.rdb.Bton(proposal.TotalVote),
		proposal.TransactionHash,
		proposal.SubmitBlockHeight,
		proposalView.rdb.Tton(&proposal.SubmitTime),
		proposalView.rdb.Tton(&proposal.DepositEndTime),
		proposalView.rdb.Tton(proposal.MaybeVotingStartTime),
		proposal.MaybeVotingEndBlockHeight,
		proposalView.rdb.Tton(proposal.MaybeVotingEndTime),
	).ToSql()
	if err != nil {
		return fmt.Errorf("error building proposal insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := proposalView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting proposal into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting proposal into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (proposalView *ProposalsView) IncrementTotalVoteBy(proposalId uint64, voteToAdd *big.Int) error {
	var err error

	selectStmtBuilder := proposalView.rdb.StmtBuilder.Select(
		PROPOSALS_TABLE_NAME,
	).From(
		PROPOSALS_TABLE_NAME,
	).Where(
		"proposal_id = ?", proposalId,
	)

	sql, sqlArgs, err := selectStmtBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("error building proposal selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	totalVoteReader := proposalView.rdb.NtobReader()

	if err = proposalView.rdb.QueryRow(sql, sqlArgs...).Scan(totalVoteReader.ScannableArg()); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return rdb.ErrNoRows
		}
		return fmt.Errorf("error scanning projection row: %v: %w", err, rdb.ErrQuery)
	}
	totalVote, parseErr := totalVoteReader.Parse()
	if parseErr != nil {
		return fmt.Errorf("error parsing proposal total vote: %v: %w", parseErr, rdb.ErrQuery)
	}

	totalVote = new(big.Int).Add(totalVote, voteToAdd)

	sql, sqlArgs, err = proposalView.rdb.StmtBuilder.Update(
		PROPOSALS_TABLE_NAME,
	).Set(
		"total_vote", proposalView.rdb.Bton(totalVote),
	).ToSql()
	if err != nil {
		return fmt.Errorf("error building proposal update sql: %v: %w", err, rdb.ErrPrepare)
	}
	result, err := proposalView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error updating proposal total vote: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("error updating proposal total vote: no rows updated: %w", rdb.ErrWrite)
	}

	return nil
}

func (proposalView *ProposalsView) Update(row *ProposalRow) error {
	sql, sqlArgs, err := proposalView.rdb.StmtBuilder.Update(
		PROPOSALS_TABLE_NAME,
	).SetMap(map[string]interface{}{
		"title":                           row.Title,
		"description":                     row.Description,
		"type":                            row.Type,
		"status":                          row.Status,
		"proposer_address":                row.ProposerAddress,
		"maybe_proposer_operator_address": row.MaybeProposerOperatorAddress,
		"data":                            json.MustMarshalToString(row.Data),
		"initial_deposit":                 json.MustMarshalToString(row.InitialDeposit),
		"total_deposit":                   json.MustMarshalToString(row.TotalDeposit),
		"total_vote":                      proposalView.rdb.Bton(row.TotalVote),
		"transaction_hash":                row.TransactionHash,
		"submit_block_height":             row.SubmitBlockHeight,
		"submit_time":                     proposalView.rdb.Tton(&row.SubmitTime),
		"deposit_end_time":                proposalView.rdb.Tton(&row.DepositEndTime),
		"maybe_voting_start_time":         proposalView.rdb.Tton(row.MaybeVotingStartTime),
		"maybe_voting_end_block_height":   row.MaybeVotingEndBlockHeight,
		"maybe_voting_end_time":           proposalView.rdb.Tton(row.MaybeVotingEndTime),
	}).Where("proposal_id = ?", row.ProposalId).ToSql()
	if err != nil {
		return fmt.Errorf("error building proposal update sql: %v: %w", err, rdb.ErrPrepare)
	}
	result, err := proposalView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error updating proposal: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("error updating proposal: no rows updated: %w", rdb.ErrWrite)
	}

	return nil
}

func (proposalView *ProposalsView) FindById(proposalId string) (*ProposalWithMonikerRow, error) {
	selectStmtBuilder := proposalView.rdb.StmtBuilder.Select(
		fmt.Sprintf("%s.proposal_id", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.title", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.description", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.type", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.status", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.proposer_address", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.maybe_proposer_operator_address", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.data", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.initial_deposit", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.total_deposit", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.total_vote", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.transaction_hash", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.submit_block_height", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.submit_time", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.deposit_end_time", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.maybe_voting_start_time", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.maybe_voting_end_block_height", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.maybe_voting_end_time", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.moniker", VALIDATORS_TABLE_NAME),
	).From(
		PROPOSALS_TABLE_NAME,
	).LeftJoin(
		fmt.Sprintf(
			"%s ON %s.maybe_proposer_operator_address=%s.operator_address",
			VALIDATORS_TABLE_NAME, PROPOSALS_TABLE_NAME, VALIDATORS_TABLE_NAME,
		),
	).Where(
		fmt.Sprintf("%s.proposal_id = ?", PROPOSALS_TABLE_NAME), proposalId,
	)

	sql, sqlArgs, err := selectStmtBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building proposal selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	var row ProposalWithMonikerRow
	var dataJSON *string
	var initialDepositJSON *string
	var totalDepositJSON *string
	totalVoteJSON := proposalView.rdb.NtobReader()
	submitBlockTimeReader := proposalView.rdb.NtotReader()
	depositEndTimeReader := proposalView.rdb.NtotReader()
	votingStartTimeReader := proposalView.rdb.NtotReader()
	votingEndTimeReader := proposalView.rdb.NtotReader()

	if err = proposalView.rdb.QueryRow(sql, sqlArgs...).Scan(
		&row.ProposalId,
		&row.Title,
		&row.Description,
		&row.Type,
		&row.Status,
		&row.ProposerAddress,
		&row.MaybeProposerOperatorAddress,
		&dataJSON,
		&initialDepositJSON,
		&totalDepositJSON,
		totalVoteJSON.ScannableArg(),
		&row.TransactionHash,
		&row.SubmitBlockHeight,
		submitBlockTimeReader.ScannableArg(),
		depositEndTimeReader.ScannableArg(),
		votingStartTimeReader.ScannableArg(),
		&row.MaybeVotingEndBlockHeight,
		votingEndTimeReader.ScannableArg(),
		&row.MaybeProposerMoniker,
	); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return nil, rdb.ErrNoRows
		}
		return nil, fmt.Errorf("error scanning proposal row: %v: %w", err, rdb.ErrQuery)
	}

	json.MustUnmarshalFromString(*dataJSON, &row.Data)
	json.MustUnmarshalFromString(*initialDepositJSON, &row.InitialDeposit)
	json.MustUnmarshalFromString(*totalDepositJSON, &row.TotalDeposit)

	var parseErr error
	row.TotalVote, parseErr = totalVoteJSON.Parse()
	if parseErr != nil {
		return nil, fmt.Errorf("error parsing proposal total vote: %v: %w", parseErr, rdb.ErrQuery)
	}

	submitBlockTime, parseErr := submitBlockTimeReader.Parse()
	if parseErr != nil {
		return nil, fmt.Errorf("error parsing proposal submit block time: %v: %w", parseErr, rdb.ErrQuery)
	}
	row.SubmitTime = *submitBlockTime

	depositEndTime, parseErr := depositEndTimeReader.Parse()
	if parseErr != nil {
		return nil, fmt.Errorf("error parsing proposal deposit end time: %v: %w", parseErr, rdb.ErrQuery)
	}
	row.DepositEndTime = *depositEndTime

	row.MaybeVotingStartTime, parseErr = votingStartTimeReader.Parse()
	if parseErr != nil {
		return nil, fmt.Errorf("error parsing proposal voting start time: %v: %w", parseErr, rdb.ErrQuery)
	}

	row.MaybeVotingEndTime, parseErr = votingEndTimeReader.Parse()
	if parseErr != nil {
		return nil, fmt.Errorf("error parsing proposal voting end time: %v: %w", parseErr, rdb.ErrQuery)
	}

	return &row, nil
}

func (proposalView *ProposalsView) List(
	filter ProposalListFilter,
	order ProposalListOrder,
	pagination *pagination_interface.Pagination,
) ([]ProposalWithMonikerRow, *pagination_interface.PaginationResult, error) {
	stmtBuilder := proposalView.rdb.StmtBuilder.Select(
		fmt.Sprintf("%s.proposal_id", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.title", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.description", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.type", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.status", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.proposer_address", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.maybe_proposer_operator_address", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.data", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.initial_deposit", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.total_deposit", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.total_vote", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.transaction_hash", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.submit_block_height", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.submit_time", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.deposit_end_time", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.maybe_voting_start_time", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.maybe_voting_end_block_height", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.maybe_voting_end_time", PROPOSALS_TABLE_NAME),
		fmt.Sprintf("%s.moniker", VALIDATORS_TABLE_NAME),
	).From(
		PROPOSALS_TABLE_NAME,
	).LeftJoin(
		fmt.Sprintf(
			"%s ON %s.maybe_proposer_operator_address=%s.operator_address",
			VALIDATORS_TABLE_NAME, PROPOSALS_TABLE_NAME, VALIDATORS_TABLE_NAME,
		),
	)

	if filter.MaybeStatus != nil {
		stmtBuilder = stmtBuilder.Where(fmt.Sprintf("%s.status = ?", PROPOSALS_TABLE_NAME), *filter.MaybeStatus)
	}

	if order.Id == view.ORDER_DESC {
		stmtBuilder = stmtBuilder.OrderBy(fmt.Sprintf("%s.proposal_id::bigserial DESC", PROPOSALS_TABLE_NAME))
	} else {
		stmtBuilder = stmtBuilder.OrderBy(fmt.Sprintf("%s.proposal_id::bigserial", PROPOSALS_TABLE_NAME))
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		proposalView.rdb,
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building proposal list SQL: %v: %w", err, rdb.ErrPrepare)
	}

	rowsResult, err := proposalView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing proposal list query: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	rows := make([]ProposalWithMonikerRow, 0)
	for rowsResult.Next() {
		var row ProposalWithMonikerRow
		var dataJSON *string
		var initialDepositJSON *string
		var totalDepositJSON *string
		totalVoteJSON := proposalView.rdb.NtobReader()
		submitBlockTimeReader := proposalView.rdb.NtotReader()
		depositEndTimeReader := proposalView.rdb.NtotReader()
		votingStartTimeReader := proposalView.rdb.NtotReader()
		votingEndTimeReader := proposalView.rdb.NtotReader()

		if scanErr := rowsResult.Scan(
			&row.ProposalId,
			&row.Title,
			&row.Description,
			&row.Type,
			&row.Status,
			&row.ProposerAddress,
			&row.MaybeProposerOperatorAddress,
			&dataJSON,
			&initialDepositJSON,
			&totalDepositJSON,
			totalVoteJSON.ScannableArg(),
			&row.TransactionHash,
			&row.SubmitBlockHeight,
			submitBlockTimeReader.ScannableArg(),
			depositEndTimeReader.ScannableArg(),
			votingStartTimeReader.ScannableArg(),
			&row.MaybeVotingEndBlockHeight,
			votingEndTimeReader.ScannableArg(),
			&row.MaybeProposerMoniker,
		); scanErr != nil {
			if errors.Is(scanErr, rdb.ErrNoRows) {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning proposal row: %v: %w", scanErr, rdb.ErrQuery)
		}

		json.MustUnmarshalFromString(*dataJSON, &row.Data)
		json.MustUnmarshalFromString(*initialDepositJSON, &row.InitialDeposit)
		json.MustUnmarshalFromString(*totalDepositJSON, &row.TotalDeposit)

		var parseErr error
		row.TotalVote, parseErr = totalVoteJSON.Parse()
		if parseErr != nil {
			return nil, nil, fmt.Errorf("error parsing proposal total vote: %v: %w", parseErr, rdb.ErrQuery)
		}

		submitBlockTime, parseErr := submitBlockTimeReader.Parse()
		if parseErr != nil {
			return nil, nil, fmt.Errorf("error parsing proposal submit block time: %v: %w", parseErr, rdb.ErrQuery)
		}
		row.SubmitTime = *submitBlockTime

		depositEndTime, parseErr := depositEndTimeReader.Parse()
		if parseErr != nil {
			return nil, nil, fmt.Errorf("error parsing proposal deposit end time: %v: %w", parseErr, rdb.ErrQuery)
		}
		row.DepositEndTime = *depositEndTime

		row.MaybeVotingStartTime, parseErr = votingStartTimeReader.Parse()
		if parseErr != nil {
			return nil, nil, fmt.Errorf("error parsing proposal voting start time: %v: %w", parseErr, rdb.ErrQuery)
		}

		row.MaybeVotingEndTime, parseErr = votingEndTimeReader.Parse()
		if parseErr != nil {
			return nil, nil, fmt.Errorf("error parsing proposal voting end time: %v: %w", parseErr, rdb.ErrQuery)
		}

		rows = append(rows, row)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return rows, paginationResult, nil
}

type ProposalListFilter struct {
	MaybeStatus *string
}

type ProposalListOrder struct {
	Id view.ORDER
}

const PROPOSAL_STATUS_DEPOSIT_PERIOD = "DEPOSIT_PERIOD"
const PROPOSAL_STATUS_INACTIVE = "INACTIVE"
const PROPOSAL_STATUS_VOTING_PERIOD = "VOTING_PERIOD"
const PROPOSAL_STATUS_PASSED = "PASSED"
const PROPOSAL_STATUS_FAILED = "FAILED"
const PROPOSAL_STATUS_REJECTED = "REJECTED"

type ProposalWithMonikerRow struct {
	ProposalRow

	MaybeProposerMoniker *string `json:"maybeProposerMoniker"`
}

type ProposalRow struct {
	ProposalId                   string           `json:"id"`
	Title                        string           `json:"title"`
	Description                  string           `json:"description"`
	Type                         string           `json:"type"`
	Status                       string           `json:"status"`
	ProposerAddress              string           `json:"proposerAddress"`
	MaybeProposerOperatorAddress *string          `json:"maybeProposerOperatorAddress"`
	Data                         interface{}      `json:"data"`
	InitialDeposit               coin.Coins       `json:"initialDeposit"`
	TotalDeposit                 coin.Coins       `json:"totalDeposit"`
	TotalVote                    *big.Int         `json:"totalVote"`
	TransactionHash              string           `json:"transactionHash"`
	SubmitBlockHeight            int64            `json:"submitBlockHeight"`
	SubmitTime                   utctime.UTCTime  `json:"submitTime"`
	DepositEndTime               utctime.UTCTime  `json:"depositEndTime"`
	MaybeVotingStartTime         *utctime.UTCTime `json:"maybeVotingStartTime"`
	MaybeVotingEndBlockHeight    *int64           `json:"maybeVotingEndBlockHeight"`
	MaybeVotingEndTime           *utctime.UTCTime `json:"maybeVotingEndTime"`
}
