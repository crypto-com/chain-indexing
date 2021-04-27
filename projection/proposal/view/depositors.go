package view

import (
	"errors"
	"fmt"

	"github.com/crypto-com/chain-indexing/usecase/coin"

	sq "github.com/Masterminds/squirrel"

	"github.com/crypto-com/chain-indexing/appinterface/pagination"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/internal/json"

	"github.com/crypto-com/chain-indexing/internal/utctime"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

const DEPOSITORS_TABLE_NAME = "view_proposal_depositors"

type Depositors struct {
	rdb *rdb.Handle
}

func NewDepositors(handle *rdb.Handle) *Depositors {
	return &Depositors{
		handle,
	}
}

func (proposalView *Depositors) Insert(row *DepositorRow) error {
	var err error

	sql, sqlArgs, err := proposalView.rdb.StmtBuilder.Insert(
		DEPOSITORS_TABLE_NAME,
	).Columns(
		"proposal_id",
		"depositor_address",
		"maybe_depositor_operator_address",
		"transaction_hash",
		"deposit_at_block_height",
		"deposit_at_block_time",
		"amount",
	).Values(
		row.ProposalId,
		row.DepositorAddress,
		row.MaybeDepositorOperatorAddress,
		row.TransactionHash,
		row.DepositAtBlockHeight,
		proposalView.rdb.TypeConv.Tton(&row.DepositAtBlockTime),
		json.MustMarshalToString(row.Amount),
	).ToSql()
	if err != nil {
		return fmt.Errorf("error building depositor insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := proposalView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting depositor into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting depositor into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (proposalView *Depositors) ListByProposalId(
	proposalId string,
	order DepositorListOrder,
	pagination *pagination.Pagination,
) ([]DepositorWithMonikerRow, *pagination.PaginationResult, error) {
	stmtBuilder := proposalView.rdb.StmtBuilder.Select(
		fmt.Sprintf("%s.proposal_id", DEPOSITORS_TABLE_NAME),
		fmt.Sprintf("%s.depositor_address", DEPOSITORS_TABLE_NAME),
		fmt.Sprintf("%s.maybe_depositor_operator_address", DEPOSITORS_TABLE_NAME),
		fmt.Sprintf("%s.transaction_hash", DEPOSITORS_TABLE_NAME),
		fmt.Sprintf("%s.deposit_at_block_height", DEPOSITORS_TABLE_NAME),
		fmt.Sprintf("%s.deposit_at_block_time", DEPOSITORS_TABLE_NAME),
		fmt.Sprintf("%s.amount", DEPOSITORS_TABLE_NAME),
		fmt.Sprintf("%s.moniker", VALIDATORS_TABLE_NAME),
	).From(
		DEPOSITORS_TABLE_NAME,
	).LeftJoin(
		fmt.Sprintf(
			"%s ON %s.maybe_depositor_operator_address=%s.operator_address",
			VALIDATORS_TABLE_NAME, DEPOSITORS_TABLE_NAME, VALIDATORS_TABLE_NAME,
		),
	).Where(
		fmt.Sprintf("%s.proposal_id = ?", DEPOSITORS_TABLE_NAME), proposalId,
	)

	if order.DepositAtBlockHeight == view.ORDER_DESC {
		stmtBuilder = stmtBuilder.OrderBy(fmt.Sprintf("%s.id DESC", DEPOSITORS_TABLE_NAME))
	} else {
		stmtBuilder = stmtBuilder.OrderBy(fmt.Sprintf("%s.id", DEPOSITORS_TABLE_NAME))
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		proposalView.rdb,
	).WithCustomTotalQueryFn(
		func(rdbHandle *rdb.Handle, _ sq.SelectBuilder) (int64, error) {
			totalView := NewDepositorsTotal(rdbHandle)
			identity := fmt.Sprintf("%s:-", proposalId)
			total, err := totalView.FindBy(identity)
			if err != nil {
				return int64(0), err
			}
			return total, nil
		},
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building depositor list SQL: %v: %w", err, rdb.ErrPrepare)
	}

	rowsResult, err := proposalView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing depositor list query: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	rows := make([]DepositorWithMonikerRow, 0)
	for rowsResult.Next() {
		var row DepositorWithMonikerRow
		var amountJSON *string
		depositAtBlockTimeReader := proposalView.rdb.NtotReader()

		if scanErr := rowsResult.Scan(
			&row.ProposalId,
			&row.DepositorAddress,
			&row.MaybeDepositorOperatorAddress,
			&row.TransactionHash,
			&row.DepositAtBlockHeight,
			depositAtBlockTimeReader.ScannableArg(),
			&amountJSON,
			&row.MaybeDepositorMoniker,
		); scanErr != nil {
			if errors.Is(scanErr, rdb.ErrNoRows) {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning proposal row: %v: %w", scanErr, rdb.ErrQuery)
		}

		json.MustUnmarshalFromString(*amountJSON, &row.Amount)

		depositAtBlockTime, parseErr := depositAtBlockTimeReader.Parse()
		if parseErr != nil {
			return nil, nil, fmt.Errorf("error parsing deposit at block time: %v: %w", parseErr, rdb.ErrQuery)
		}
		row.DepositAtBlockTime = *depositAtBlockTime

		rows = append(rows, row)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return rows, paginationResult, nil
}

type DepositorListOrder struct {
	DepositAtBlockHeight view.ORDER
}

type DepositorWithMonikerRow struct {
	DepositorRow

	MaybeDepositorMoniker *string `json:"maybeDepositorMoniker"`
}

type DepositorRow struct {
	ProposalId                    string          `json:"proposalId"`
	DepositorAddress              string          `json:"depositorAddress"`
	MaybeDepositorOperatorAddress *string         `json:"maybeDepositorValidatorAddress"`
	TransactionHash               string          `json:"transactionHash"`
	DepositAtBlockHeight          int64           `json:"depositAtBlockHeight"`
	DepositAtBlockTime            utctime.UTCTime `json:"depositAtBlockTime"`
	Amount                        coin.Coins      `json:"amount"`
}
