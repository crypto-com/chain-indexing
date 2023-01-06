package view

import (
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/crypto-com/chain-indexing/external/json"
	jsoniter "github.com/json-iterator/go"

	"github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
)

const VOTES_TABLE_NAME = "view_proposal_votes"

type Votes interface {
	Insert(row *VoteRow) error
	Update(row *VoteRow) error
	FindByProposalIdVoter(
		proposalId string,
		voterAddress string,
	) (*[]VoteWithMonikerRow, error)
	ListByProposalId(
		proposalId string,
		order VoteListOrder,
		pagination *pagination.Pagination,
	) ([]VoteWithMonikerRow, *pagination.PaginationResult, error)
	DeleteByProposalIdVoter(
		proposalId string,
		voterAddress string,
	) (int64, error)
}

type VotesView struct {
	rdb *rdb.Handle
}

func NewVotesView(handle *rdb.Handle) Votes {
	return &VotesView{
		handle,
	}
}

func (votesView *VotesView) Insert(row *VoteRow) error {
	var err error

	var historiesJSON string
	if historiesJSON, err = jsoniter.MarshalToString(row.Histories); err != nil {
		return fmt.Errorf("error JSON marshalling vote histories for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	var answerJSON string
	if answerJSON, err = jsoniter.MarshalToString(row.Answer); err != nil {
		return fmt.Errorf("error JSON marshalling vote answer for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	sql, sqlArgs, err := votesView.rdb.StmtBuilder.Insert(
		VOTES_TABLE_NAME,
	).Columns(
		"proposal_id",
		"voter_address",
		"maybe_voter_operator_address",
		"transaction_hash",
		"vote_at_block_height",
		"vote_at_block_time",
		"answer",
		"histories",
		"metadata",
		"weight",
	).Values(
		row.ProposalId,
		row.VoterAddress,
		row.MaybeVoterOperatorAddress,
		row.TransactionHash,
		row.VoteAtBlockHeight,
		votesView.rdb.TypeConv.Tton(&row.VoteAtBlockTime),
		answerJSON,
		historiesJSON,
		row.Metadata,
		row.Weight,
	).ToSql()
	if err != nil {
		return fmt.Errorf("error building vote insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := votesView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting vote into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting vote into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (votesView *VotesView) Update(row *VoteRow) error {
	historiesJSON, err := jsoniter.MarshalToString(row.Histories)
	if err != nil {
		return fmt.Errorf("error JSON marshalling vote histories for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	var answerJSON string
	if answerJSON, err = jsoniter.MarshalToString(row.Answer); err != nil {
		return fmt.Errorf("error JSON marshalling vote answer for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	sql, sqlArgs, err := votesView.rdb.StmtBuilder.Update(
		VOTES_TABLE_NAME,
	).SetMap(map[string]interface{}{
		"transaction_hash":     row.TransactionHash,
		"vote_at_block_height": row.VoteAtBlockHeight,
		"vote_at_block_time":   votesView.rdb.TypeConv.Tton(&row.VoteAtBlockTime),
		"answer":               answerJSON,
		"histories":            historiesJSON,
	}).Where("voter_address = ?", row.VoterAddress).ToSql()
	if err != nil {
		return fmt.Errorf("error building vote update sql: %v: %w", err, rdb.ErrPrepare)
	}
	result, err := votesView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error updating vote: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("error updating vote: no rows updated: %w", rdb.ErrWrite)
	}

	return nil
}

func (votesView *VotesView) FindByProposalIdVoter(proposalId string, voterAddress string) (*[]VoteWithMonikerRow, error) {
	sql, sqlArgs, err := votesView.rdb.StmtBuilder.Select(
		fmt.Sprintf("%s.proposal_id", VOTES_TABLE_NAME),
		fmt.Sprintf("%s.voter_address", VOTES_TABLE_NAME),
		fmt.Sprintf("%s.maybe_voter_operator_address", VOTES_TABLE_NAME),
		fmt.Sprintf("%s.transaction_hash", VOTES_TABLE_NAME),
		fmt.Sprintf("%s.vote_at_block_height", VOTES_TABLE_NAME),
		fmt.Sprintf("%s.vote_at_block_time", VOTES_TABLE_NAME),
		fmt.Sprintf("%s.weight", VOTES_TABLE_NAME),
		fmt.Sprintf("%s.answer", VOTES_TABLE_NAME),
		fmt.Sprintf("%s.histories", VOTES_TABLE_NAME),
		fmt.Sprintf("%s.moniker", VALIDATORS_TABLE_NAME),
	).From(
		VOTES_TABLE_NAME,
	).LeftJoin(
		fmt.Sprintf(
			"%s ON %s.maybe_voter_operator_address=%s.operator_address",
			VALIDATORS_TABLE_NAME, VOTES_TABLE_NAME, VALIDATORS_TABLE_NAME,
		),
	).Where(
		fmt.Sprintf("%s.proposal_id = ?", VOTES_TABLE_NAME), proposalId,
	).Where(
		fmt.Sprintf("%s.voter_address = ?", VOTES_TABLE_NAME), voterAddress,
	).ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building vote selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	var historiesJSON *string
	var answerJSON *string
	voteAtBlockTimeReader := votesView.rdb.NtotReader()

	rowsResult, err := votesView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, fmt.Errorf("error executing vote select SQL: %v: %w", err, rdb.ErrQuery)
	}

	rows := make([]VoteWithMonikerRow, 0)
	for rowsResult.Next() {
		var row VoteWithMonikerRow
		if err := rowsResult.Scan(
			&row.ProposalId,
			&row.VoterAddress,
			&row.MaybeVoterOperatorAddress,
			&row.TransactionHash,
			&row.VoteAtBlockHeight,
			voteAtBlockTimeReader.ScannableArg(),
			&row.Weight,
			&answerJSON,
			&historiesJSON,
			&row.MaybeVoterMoniker,
		); err != nil {
			if errors.Is(err, rdb.ErrNoRows) {
				return nil, rdb.ErrNoRows
			}
			return nil, fmt.Errorf("error scanning proposal row: %v: %w", err, rdb.ErrQuery)
		}
		json.MustUnmarshalFromString(*answerJSON, &row.Answer)
		json.MustUnmarshalFromString(*historiesJSON, &row.Histories)

		voteAtBlockTime, parseErr := voteAtBlockTimeReader.Parse()
		if parseErr != nil {
			return nil, fmt.Errorf("error parsing vote at block time: %v: %w", parseErr, rdb.ErrQuery)
		}
		row.VoteAtBlockTime = *voteAtBlockTime
		rows = append(rows, row)
	}

	return &rows, nil
}

func (votesView *VotesView) ListByProposalId(
	proposalId string,
	order VoteListOrder,
	pagination *pagination.Pagination,
) ([]VoteWithMonikerRow, *pagination.PaginationResult, error) {
	stmtBuilder := votesView.rdb.StmtBuilder.Select(
		fmt.Sprintf("%s.proposal_id", VOTES_TABLE_NAME),
		fmt.Sprintf("%s.voter_address", VOTES_TABLE_NAME),
		fmt.Sprintf("%s.maybe_voter_operator_address", VOTES_TABLE_NAME),
		fmt.Sprintf("%s.transaction_hash", VOTES_TABLE_NAME),
		fmt.Sprintf("%s.vote_at_block_height", VOTES_TABLE_NAME),
		fmt.Sprintf("%s.vote_at_block_time", VOTES_TABLE_NAME),
		fmt.Sprintf("%s.answer", VOTES_TABLE_NAME),
		fmt.Sprintf("%s.histories", VOTES_TABLE_NAME),
		fmt.Sprintf("%s.moniker", VALIDATORS_TABLE_NAME),
	).From(
		VOTES_TABLE_NAME,
	).LeftJoin(
		fmt.Sprintf(
			"%s ON %s.maybe_voter_operator_address=%s.operator_address",
			VALIDATORS_TABLE_NAME, VOTES_TABLE_NAME, VALIDATORS_TABLE_NAME,
		),
	).Where(
		fmt.Sprintf("%s.proposal_id = ?", VOTES_TABLE_NAME), proposalId,
	)

	if order.VoteAtBlockHeight == view.ORDER_DESC {
		stmtBuilder = stmtBuilder.OrderBy(fmt.Sprintf("%s.vote_at_block_height DESC", VOTES_TABLE_NAME))
	} else {
		stmtBuilder = stmtBuilder.OrderBy(fmt.Sprintf("%s.vote_at_block_height", VOTES_TABLE_NAME))
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		votesView.rdb,
	).WithCustomTotalQueryFn(
		func(rdbHandle *rdb.Handle, _ sq.SelectBuilder) (int64, error) {
			totalView := NewVotesTotalView(rdbHandle)
			total, err := totalView.FindBy(proposalId)
			if err != nil {
				return int64(0), err
			}
			return total, nil
		},
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building vote selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	var historiesJSON *string
	var answerJSON *string
	voteAtBlockTimeReader := votesView.rdb.NtotReader()

	rowsResult, err := votesView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing vote select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	rows := make([]VoteWithMonikerRow, 0)
	for rowsResult.Next() {
		var row VoteWithMonikerRow
		if scanErr := rowsResult.Scan(
			&row.ProposalId,
			&row.VoterAddress,
			&row.MaybeVoterOperatorAddress,
			&row.TransactionHash,
			&row.VoteAtBlockHeight,
			voteAtBlockTimeReader.ScannableArg(),
			&answerJSON,
			&historiesJSON,
			&row.MaybeVoterMoniker,
		); scanErr != nil {
			if errors.Is(scanErr, rdb.ErrNoRows) {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning proposal row: %v: %w", scanErr, rdb.ErrQuery)
		}

		json.MustUnmarshalFromString(*answerJSON, &row.Answer)
		json.MustUnmarshalFromString(*historiesJSON, &row.Histories)

		voteAtBlockTime, parseErr := voteAtBlockTimeReader.Parse()
		if parseErr != nil {
			return nil, nil, fmt.Errorf("error parsing vote at block time: %v: %w", parseErr, rdb.ErrQuery)
		}
		row.VoteAtBlockTime = *voteAtBlockTime

		rows = append(rows, row)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return rows, paginationResult, nil
}

func (votesView *VotesView) DeleteByProposalIdVoter(proposalId string, voterAddress string) (int64, error) {
	sql, sqlArgs, err := votesView.rdb.StmtBuilder.Delete(
		VOTES_TABLE_NAME,
	).Where("proposal_id = ? AND voter_address = ?", proposalId, voterAddress).ToSql()
	if err != nil {
		return 0, fmt.Errorf("error building proposal votes deletion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := votesView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return 0, fmt.Errorf("error deleting proposal votes from the table: %v: %w", err, rdb.ErrWrite)
	}

	return result.RowsAffected(), nil
}

type VoteListOrder struct {
	VoteAtBlockHeight view.ORDER
}

type VoteWithMonikerRow struct {
	VoteRow

	MaybeVoterMoniker *string `json:"maybe_voter_moniker"`
}

type VoteRow struct {
	ProposalId                string          `json:"proposalId"`
	VoterAddress              string          `json:"voterAddress"`
	MaybeVoterOperatorAddress *string         `json:"maybeVoterOperatorAddress"`
	TransactionHash           string          `json:"transactionHash"`
	VoteAtBlockHeight         int64           `json:"voteAtBlockHeight"`
	VoteAtBlockTime           utctime.UTCTime `json:"voteAtBlockTime"`
	Answer                    string          `json:"answer"`
	Histories                 []VoteHistory   `json:"histories"`
	Metadata                  string          `json:"metadata"`
	Weight                    string          `json:"weight"`
}

type VoteHistory struct {
	TransactionHash   string          `json:"transactionHash"`
	VoteAtBlockHeight int64           `json:"voteAtBlockHeight"`
	VoteAtBlockTime   utctime.UTCTime `json:"voteAtBlockTime"`
	Answer            string          `json:"answer"`
	Weight            string          `json:"weight"`
}
