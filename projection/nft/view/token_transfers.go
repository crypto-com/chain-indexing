package view

import (
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/internal/utctime"
)

// nolint:gosec
const TOKEN_TRANSFERS_TABLE_NAME = "view_nft_token_transfers"

type TokenTransfers struct {
	rdb *rdb.Handle
}

func NewTokenTransfers(handle *rdb.Handle) *TokenTransfers {
	return &TokenTransfers{
		handle,
	}
}

func (tokenTransfersView *TokenTransfers) Insert(
	tokenTransferRow TokenTransferRow,
) error {
	var err error

	sql, sqlArgs, err := tokenTransfersView.rdb.StmtBuilder.Insert(
		TOKEN_TRANSFERS_TABLE_NAME,
	).Columns(
		"denom_id",
		"token_id",
		"drop",
		"block_height",
		"transaction_hash",
		"sender",
		"recipient",
		"transferred_at",
	).Values(
		tokenTransferRow.DenomId,
		tokenTransferRow.TokenId,
		tokenTransferRow.Drop,
		tokenTransferRow.BlockHeight,
		tokenTransferRow.TransactionHash,
		tokenTransferRow.Sender,
		tokenTransferRow.Recipient,
		tokenTransfersView.rdb.Tton(&tokenTransferRow.TransferredAt),
	).ToSql()
	if err != nil {
		return fmt.Errorf("error building NFT token transfer insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := tokenTransfersView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting NFT token transfer into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting NFT token transfer into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (tokenTransfersView *TokenTransfers) List(
	filter TokenTransferListFilter,
	order TokenTransferListOrder,
	pagination *pagination_interface.Pagination,
) ([]TokenTransferRow, *pagination_interface.PaginationResult, error) {
	stmtBuilder := tokenTransfersView.rdb.StmtBuilder.Select(
		"denom_id",
		"token_id",
		"drop",
		"block_height",
		"transaction_hash",
		"sender",
		"recipient",
		"transferred_at",
	).From(
		TOKEN_TRANSFERS_TABLE_NAME,
	)

	if filter.MaybeDenomId != nil {
		stmtBuilder = stmtBuilder.Where("denom_id = ?", *filter.MaybeDenomId)
	}
	if filter.MaybeTokenId != nil {
		stmtBuilder = stmtBuilder.Where("token_id = ?", *filter.MaybeTokenId)
	}
	if filter.MaybeDrop != nil {
		stmtBuilder = stmtBuilder.Where("drop = ?", *filter.MaybeDrop)
	}
	if filter.MaybeBlockHeight != nil {
		stmtBuilder = stmtBuilder.Where("block_height = ?", *filter.MaybeBlockHeight)
	}
	if filter.MaybeSender != nil {
		stmtBuilder = stmtBuilder.Where("sender = ?", *filter.MaybeSender)
	}
	if filter.MaybeRecipient != nil {
		stmtBuilder = stmtBuilder.Where("recipient = ?", *filter.MaybeRecipient)
	}
	if filter.MaybeAccount != nil {
		stmtBuilder = stmtBuilder.Where(
			sq.Or{
				sq.Eq{
					"sender": *filter.MaybeAccount,
				},
				sq.Eq{
					"recipient": *filter.MaybeAccount,
				},
			},
		)
	}

	if order.Id == view.ORDER_DESC {
		stmtBuilder = stmtBuilder.OrderBy("id DESC")
	} else {
		stmtBuilder = stmtBuilder.OrderBy("id")
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		tokenTransfersView.rdb,
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building NFT token transfers list SQL: %v: %w", err, rdb.ErrPrepare)
	}

	rowsResult, err := tokenTransfersView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing NFT token transfers list query: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	rows := make([]TokenTransferRow, 0)
	for rowsResult.Next() {
		var row TokenTransferRow
		transferredAtTimeReader := tokenTransfersView.rdb.NtotReader()

		if scanErr := rowsResult.Scan(
			&row.DenomId,
			&row.TokenId,
			&row.Drop,
			&row.BlockHeight,
			&row.TransactionHash,
			&row.Sender,
			&row.Recipient,
			transferredAtTimeReader.ScannableArg(),
		); scanErr != nil {
			if errors.Is(scanErr, rdb.ErrNoRows) {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning NFT token transfer row: %v: %w", scanErr, rdb.ErrQuery)
		}

		transferredAt, parseTransferredAtErr := transferredAtTimeReader.Parse()
		if parseTransferredAtErr != nil {
			return nil, nil, fmt.Errorf("error parsing NFT token transfer time: %v: %w", parseTransferredAtErr, rdb.ErrQuery)
		}
		row.TransferredAt = *transferredAt

		rows = append(rows, row)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return rows, paginationResult, nil
}

type TokenTransferListFilter struct {
	MaybeDenomId     *string
	MaybeTokenId     *string
	MaybeDrop        *string
	MaybeBlockHeight *string
	MaybeSender      *string
	MaybeRecipient   *string
	MaybeAccount     *string // Sender OR Recipient
}

type TokenTransferListOrder struct {
	Id view.ORDER
}

type TokenTransferRow struct {
	DenomId         string          `json:"denomId"`
	TokenId         string          `json:"tokenId"`
	Drop            string          `json:"drop"`
	BlockHeight     int64           `json:"blockHeight"`
	TransactionHash string          `json:"transactionHash"`
	Sender          string          `json:"sender"`
	Recipient       string          `json:"recipient"`
	TransferredAt   utctime.UTCTime `json:"transferredAt"`
}
