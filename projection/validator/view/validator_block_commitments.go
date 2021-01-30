package view

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/crypto-com/chain-indexing/internal/utctime"

	sq "github.com/Masterminds/squirrel"

	"github.com/crypto-com/chain-indexing/appinterface/pagination"
	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type ValidatorBlockCommitments struct {
	rdb *rdb.Handle
}

func NewValidatorBlockCommitments(handle *rdb.Handle) *ValidatorBlockCommitments {
	return &ValidatorBlockCommitments{
		handle,
	}
}

func (commitmentsView *ValidatorBlockCommitments) InsertAll(commitments []ValidatorBlockCommitmentRow) error {
	pendingRowCount := 0
	var stmtBuilder sq.InsertBuilder

	totalCommitmentCount := len(commitments)
	for i, commitment := range commitments {
		if pendingRowCount == 0 {
			stmtBuilder = commitmentsView.rdb.StmtBuilder.Insert(
				"view_validator_block_commitments",
			).Columns(
				"consensus_node_address",
				"block_height",
				"signature",
				"timestamp",
			)
		}

		stmtBuilder = stmtBuilder.Values(
			commitment.ConsensusNodeAddress,
			commitment.BlockHeight,
			commitment.Signature,
			commitmentsView.rdb.Tton(&commitment.Timestamp),
		)
		pendingRowCount += 1

		if pendingRowCount == 500 || i+1 == totalCommitmentCount {
			sql, sqlArgs, err := stmtBuilder.ToSql()
			if err != nil {
				return fmt.Errorf("error building validator block commitment insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
			}

			result, err := commitmentsView.rdb.Exec(sql, sqlArgs...)
			if err != nil {
				return fmt.Errorf("error inserting validator block commitment into the table: %v: %w", err, rdb.ErrWrite)
			}
			if result.RowsAffected() != int64(pendingRowCount) {
				return fmt.Errorf("error inserting validator block commitment into the table: no rows inserted: %w", rdb.ErrWrite)
			}

			pendingRowCount = 0
		}
	}

	return nil
}

func (commitmentsView *ValidatorBlockCommitments) Insert(commitment ValidatorBlockCommitmentRow) error {
	sql, sqlArgs, err := commitmentsView.rdb.StmtBuilder.Insert(
		"view_validator_block_commitments",
	).Columns(
		"consensus_node_address",
		"block_height",
		"signature",
		"timestamp",
	).Values(
		commitment.ConsensusNodeAddress,
		commitment.BlockHeight,
		commitment.Signature,
		commitmentsView.rdb.Tton(&commitment.Timestamp),
	).ToSql()

	if err != nil {
		return fmt.Errorf("error building validator block commitment insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := commitmentsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting validator block commitment into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting validator block commitment into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (commitmentsView *ValidatorBlockCommitments) List(
	filter ValidatorBlockCommitmentsListFilter,
	pagination *pagination_interface.Pagination,
) ([]ListValidatorBlockCommitmentRow, *pagination.PaginationResult, error) {
	if pagination.Type() != pagination_interface.PAGINATION_OFFSET {
		panic(fmt.Sprintf("unsupported pagination type: %s", pagination.Type()))
	}

	stmtBuilder := commitmentsView.rdb.StmtBuilder.Select(
		"view_validator_block_commitments.consensus_node_address",
		"view_validator_block_commitments.block_height",
		"view_validator_block_commitments.signature",
		"view_validator_block_commitments.timestamp",
		"view_validators.moniker",
	).InnerJoin(
		"view_validators ON view_validator_block_commitments.consensus_node_address=view_validators.consensus_node_address",
	).From(
		"view_validator_block_commitments",
	).OrderBy(
		"view_validator_block_commitments.id",
	)

	if filter.MaybeBlockHeight != nil {
		stmtBuilder = stmtBuilder.Where("block_height = ?", *filter.MaybeBlockHeight)
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		commitmentsView.rdb,
	).WithCustomTotalQueryFn(
		func(rdbHandle *rdb.Handle, _ sq.SelectBuilder) (int64, error) {
			identity := "-"
			if filter.MaybeBlockHeight != nil {
				identity = strconv.FormatInt(*filter.MaybeBlockHeight, 10)
			}
			totalView := NewValidatorBlockCommitmentsTotal(rdbHandle)
			total, err := totalView.FindBy(identity)
			if err != nil {
				return int64(0), err
			}
			return total, nil
		},
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building validator block commitments select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := commitmentsView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing validator block commitments select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	validatorBlockCommitments := make([]ListValidatorBlockCommitmentRow, 0)
	for rowsResult.Next() {
		var validatorBlockCommitment ListValidatorBlockCommitmentRow

		timestampParser := commitmentsView.rdb.NtotReader()
		if scanErr := rowsResult.Scan(
			&validatorBlockCommitment.ConsensusNodeAddress,
			&validatorBlockCommitment.BlockHeight,
			&validatorBlockCommitment.Signature,
			timestampParser.ScannableArg(),
			&validatorBlockCommitment.Moniker,
		); scanErr != nil {
			if errors.Is(scanErr, rdb.ErrNoRows) {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning validator block commitment row: %v: %w", scanErr, rdb.ErrQuery)
		}

		timestamp, parseErr := timestampParser.Parse()
		if parseErr != nil {
			return nil, nil, fmt.Errorf("error parsing validator block comitment timestamp: %v", parseErr)
		}
		validatorBlockCommitment.Timestamp = *timestamp

		validatorBlockCommitments = append(validatorBlockCommitments, validatorBlockCommitment)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return validatorBlockCommitments, paginationResult, nil
}

type ValidatorBlockCommitmentsListFilter struct {
	MaybeBlockHeight *int64
}

type ValidatorBlockCommitmentRow struct {
	MaybeId              *int64          `json:"-"`
	ConsensusNodeAddress string          `json:"consensusNodeAddress"`
	BlockHeight          int64           `json:"blockHeight"`
	Signature            string          `json:"signature"`
	Timestamp            utctime.UTCTime `json:"timestamp"`
}

type ListValidatorBlockCommitmentRow struct {
	ValidatorBlockCommitmentRow

	Moniker string `json:"moniker"`
}
