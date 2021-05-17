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
		"block_height",
		"transaction_hash",
		"sender",
		"recipient",
		"transferred_at",
	).Values(
		tokenTransferRow.DenomId,
		tokenTransferRow.TokenId,
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
) ([]TokenTransferRowWithDenomAndTokenDetails, *pagination_interface.PaginationResult, error) {
	stmtBuilder := tokenTransfersView.rdb.StmtBuilder.Select(
		fmt.Sprintf("%s.denom_id", TOKEN_TRANSFERS_TABLE_NAME),
		fmt.Sprintf("%s.name AS denom_name", DENOMS_TABLE_NAME),
		fmt.Sprintf("%s.schema AS denom_schema", DENOMS_TABLE_NAME),
		fmt.Sprintf("%s.token_id", TOKEN_TRANSFERS_TABLE_NAME),
		fmt.Sprintf("%s.drop AS token_drop", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.name AS token_name", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.uri AS token_uri", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.data AS token_data", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.minter AS token_minter", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.block_height", TOKEN_TRANSFERS_TABLE_NAME),
		fmt.Sprintf("%s.transaction_hash", TOKEN_TRANSFERS_TABLE_NAME),
		fmt.Sprintf("%s.sender", TOKEN_TRANSFERS_TABLE_NAME),
		fmt.Sprintf("%s.recipient", TOKEN_TRANSFERS_TABLE_NAME),
		fmt.Sprintf("%s.transferred_at", TOKEN_TRANSFERS_TABLE_NAME),
	).From(
		TOKEN_TRANSFERS_TABLE_NAME,
	).LeftJoin(
		fmt.Sprintf(
			"%s ON %s.denom_id = %s.denom_id",
			DENOMS_TABLE_NAME, DENOMS_TABLE_NAME, TOKEN_TRANSFERS_TABLE_NAME,
		),
	).LeftJoin(
		fmt.Sprintf(
			"%s ON %s.token_id = %s.token_id",
			TOKENS_TABLE_NAME, TOKENS_TABLE_NAME, TOKEN_TRANSFERS_TABLE_NAME,
		),
	)

	if filter.MaybeDenomId != nil {
		stmtBuilder = stmtBuilder.Where(
			fmt.Sprintf("%s.denom_id = ?", TOKEN_TRANSFERS_TABLE_NAME),
			*filter.MaybeDenomId,
		)
	}
	if filter.MaybeTokenId != nil {
		stmtBuilder = stmtBuilder.Where(
			fmt.Sprintf("%s.token_id = ?", TOKEN_TRANSFERS_TABLE_NAME),
			*filter.MaybeTokenId,
		)
	}
	if filter.MaybeDrop != nil {
		stmtBuilder = stmtBuilder.Where(
			fmt.Sprintf("%s.drop = ?", TOKENS_TABLE_NAME),
			*filter.MaybeDrop,
		)
	}
	if filter.MaybeBlockHeight != nil {
		stmtBuilder = stmtBuilder.Where(
			fmt.Sprintf("%s.block_height = ?", TOKEN_TRANSFERS_TABLE_NAME),
			*filter.MaybeBlockHeight,
		)
	}
	if filter.MaybeSender != nil {
		stmtBuilder = stmtBuilder.Where(
			fmt.Sprintf("%s.sender = ?", TOKEN_TRANSFERS_TABLE_NAME),
			*filter.MaybeSender,
		)
	}
	if filter.MaybeRecipient != nil {
		stmtBuilder = stmtBuilder.Where(
			fmt.Sprintf("%s.recipient = ?", TOKEN_TRANSFERS_TABLE_NAME),
			*filter.MaybeRecipient,
		)
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
		stmtBuilder = stmtBuilder.OrderBy(
			fmt.Sprintf("%s.id DESC", TOKEN_TRANSFERS_TABLE_NAME),
		)
	} else {
		stmtBuilder = stmtBuilder.OrderBy(
			fmt.Sprintf("%s.id", TOKEN_TRANSFERS_TABLE_NAME),
		)
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

	rows := make([]TokenTransferRowWithDenomAndTokenDetails, 0)
	for rowsResult.Next() {
		var row TokenTransferRowWithDenomAndTokenDetails
		transferredAtTimeReader := tokenTransfersView.rdb.NtotReader()

		if scanErr := rowsResult.Scan(
			&row.DenomId,
			&row.DenomName,
			&row.DenomSchema,
			&row.TokenId,
			&row.Drop,
			&row.TokenName,
			&row.TokenURI,
			&row.TokenData,
			&row.TokenMinter,
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

type TokenTransferRowWithDenomAndTokenDetails struct {
	TokenTransferRow

	DenomName   string `json:"denomName"`
	DenomSchema string `json:"denomSchema"`
	TokenName   string `json:"tokenName"`
	Drop        string `json:"drop"`
	TokenURI    string `json:"tokenURI"`
	TokenData   string `json:"tokenData"`
	TokenMinter string `json:"tokenMinter"`
}

type TokenTransferRow struct {
	DenomId         string          `json:"denomId"`
	TokenId         string          `json:"tokenId"`
	BlockHeight     int64           `json:"blockHeight"`
	TransactionHash string          `json:"transactionHash"`
	Sender          string          `json:"sender"`
	Recipient       string          `json:"recipient"`
	TransferredAt   utctime.UTCTime `json:"transferredAt"`
}
