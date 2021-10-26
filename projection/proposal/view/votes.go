package view

import (
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/crypto-com/chain-indexing/external/json"

	"github.com/crypto-com/chain-indexing/appinterface/pagination"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/internal/utctime"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	jsoniter "github.com/json-iterator/go"
)

const VOTES_TABLE_NAME = "view_proposal_votes"

type Votes struct {
	rdb *rdb.Handle
}

func NewVotes(handle *rdb.Handle) *Votes {
	return &Votes{
		handle,
	}
}

func (proposalView *Votes) Insert(row *VoteRow) error {
	var err error

	var historiesJSON string
	if historiesJSON, err = jsoniter.MarshalToString(row.Histories); err != nil {
		return fmt.Errorf("error JSON marshalling vote histories for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	sql, sqlArgs, err := proposalView.rdb.StmtBuilder.Insert(
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
	).Values(
		row.ProposalId,
		row.VoterAddress,
		row.MaybeVoterOperatorAddress,
		row.TransactionHash,
		row.VoteAtBlockHeight,
		proposalView.rdb.TypeConv.Tton(&row.VoteAtBlockTime),
		row.Answer,
		historiesJSON,
	).ToSql()
	if err != nil {
		return fmt.Errorf("error building vote insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := proposalView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting vote into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting vote into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (proposalView *Votes) Update(row *VoteRow) error {
	historiesJSON, err := jsoniter.MarshalToString(row.Histories)
	if err != nil {
		return fmt.Errorf("error JSON marshalling vote histories for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	sql, sqlArgs, err := proposalView.rdb.StmtBuilder.Update(
		VOTES_TABLE_NAME,
	).SetMap(map[string]interface{}{
		"transaction_hash":     row.TransactionHash,
		"vote_at_block_height": row.VoteAtBlockHeight,
		"vote_at_block_time":   proposalView.rdb.TypeConv.Tton(&row.VoteAtBlockTime),
		"answer":               row.Answer,
		"histories":            historiesJSON,
	}).Where("voter_address = ?", row.VoterAddress).ToSql()
	if err != nil {
		return fmt.Errorf("error building vote update sql: %v: %w", err, rdb.ErrPrepare)
	}
	result, err := proposalView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error updating vote: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("error updating vote: no rows updated: %w", rdb.ErrWrite)
	}

	return nil
}

func (proposalView *Votes) FindByProposalIdVoter(proposalId string, voterAddress string) (*VoteWithMonikerRow, error) {
	sql, sqlArgs, err := proposalView.rdb.StmtBuilder.Select(
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
	).Where(
		fmt.Sprintf("%s.voter_address = ?", VOTES_TABLE_NAME), voterAddress,
	).ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building vote selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	var row VoteWithMonikerRow
	var historiesJSON *string
	voteAtBlockTimeReader := proposalView.rdb.NtotReader()

	result := proposalView.rdb.QueryRow(sql, sqlArgs...)
	if err != nil {
		return nil, fmt.Errorf("error executing vote select SQL: %v: %w", err, rdb.ErrQuery)
	}
	if err := result.Scan(
		&row.ProposalId,
		&row.VoterAddress,
		&row.MaybeVoterOperatorAddress,
		&row.TransactionHash,
		&row.VoteAtBlockHeight,
		voteAtBlockTimeReader.ScannableArg(),
		&row.Answer,
		&historiesJSON,
		&row.MaybeVoterMoniker,
	); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return nil, rdb.ErrNoRows
		}
		return nil, fmt.Errorf("error scanning proposal row: %v: %w", err, rdb.ErrQuery)
	}

	json.MustUnmarshalFromString(*historiesJSON, &row.Histories)

	voteAtBlockTime, parseErr := voteAtBlockTimeReader.Parse()
	if parseErr != nil {
		return nil, fmt.Errorf("error parsing vote at block time: %v: %w", parseErr, rdb.ErrQuery)
	}
	row.VoteAtBlockTime = *voteAtBlockTime

	return &row, nil
}

func (proposalView *Votes) ListByProposalId(
	proposalId string,
	order VoteListOrder,
	pagination *pagination.Pagination,
) ([]VoteWithMonikerRow, *pagination.PaginationResult, error) {
	stmtBuilder := proposalView.rdb.StmtBuilder.Select(
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
		proposalView.rdb,
	).WithCustomTotalQueryFn(
		func(rdbHandle *rdb.Handle, _ sq.SelectBuilder) (int64, error) {
			totalView := NewVotesTotal(rdbHandle)
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
	voteAtBlockTimeReader := proposalView.rdb.NtotReader()

	rowsResult, err := proposalView.rdb.Query(sql, sqlArgs...)
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
			&row.Answer,
			&historiesJSON,
			&row.MaybeVoterMoniker,
		); scanErr != nil {
			if errors.Is(scanErr, rdb.ErrNoRows) {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning proposal row: %v: %w", scanErr, rdb.ErrQuery)
		}

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
}

type VoteHistory struct {
	TransactionHash   string          `json:"transactionHash"`
	VoteAtBlockHeight int64           `json:"voteAtBlockHeight"`
	VoteAtBlockTime   utctime.UTCTime `json:"voteAtBlockTime"`
	Answer            string          `json:"answer"`
}
