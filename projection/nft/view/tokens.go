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
const TOKENS_TABLE_NAME = "view_nft_tokens"

type Tokens struct {
	rdb *rdb.Handle
}

func NewTokens(handle *rdb.Handle) *Tokens {
	return &Tokens{
		handle,
	}
}

func (tokensView *Tokens) Insert(tokenRow *TokenRow) error {
	var err error

	sql, sqlArgs, err := tokensView.rdb.StmtBuilder.Insert(
		TOKENS_TABLE_NAME,
	).Columns(
		"denom_id",
		"token_id",
		"drop",
		"burned",
		"name",
		"uri",
		"data",
		"minter",
		"owner",
		"minted_at",
		"last_edited_at",
	).Values(
		tokenRow.DenomId,
		tokenRow.TokenId,
		tokenRow.Drop,
		tokenRow.Burned,
		tokenRow.Name,
		tokenRow.URI,
		tokenRow.Data,
		tokenRow.Minter,
		tokenRow.Owner,
		tokensView.rdb.Tton(&tokenRow.MintedAt),
		tokensView.rdb.Tton(&tokenRow.LastEditedAt),
	).ToSql()
	if err != nil {
		return fmt.Errorf("error building NFT token insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := tokensView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting NFT token into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting NFT token into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (tokensView *Tokens) FindById(denomId string, tokenId string) (*TokenRow, error) {
	selectStmtBuilder := tokensView.rdb.StmtBuilder.Select(
		"denom_id",
		"token_id",
		"drop",
		"burned",
		"name",
		"uri",
		"data",
		"minter",
		"owner",
		"minted_at",
		"last_edited_at",
	).From(
		TOKENS_TABLE_NAME,
	).Where(
		"denom_id = ? AND token_id = ?", denomId, tokenId,
	)

	sql, sqlArgs, err := selectStmtBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building NFT token selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	var row TokenRow
	mintedAtTimeReader := tokensView.rdb.NtotReader()
	lastEditedAtTimeReader := tokensView.rdb.NtotReader()

	if err = tokensView.rdb.QueryRow(sql, sqlArgs...).Scan(
		&row.DenomId,
		&row.TokenId,
		&row.Drop,
		&row.Burned,
		&row.Name,
		&row.URI,
		&row.Data,
		&row.Minter,
		&row.Owner,
		mintedAtTimeReader.ScannableArg(),
		lastEditedAtTimeReader.ScannableArg(),
	); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return nil, rdb.ErrNoRows
		}
		return nil, fmt.Errorf("error scanning NFT denom row: %v: %w", err, rdb.ErrQuery)
	}

	mintedAt, parseCreatedAtErr := mintedAtTimeReader.Parse()
	if parseCreatedAtErr != nil {
		return nil, fmt.Errorf("error parsing NFT token minted time: %v: %w", parseCreatedAtErr, rdb.ErrQuery)
	}
	row.MintedAt = *mintedAt

	lastEditedAt, parseLastEditedAtErr := lastEditedAtTimeReader.Parse()
	if parseLastEditedAtErr != nil {
		return nil, fmt.Errorf("error parsing NFT token last edited time: %v: %w", parseLastEditedAtErr, rdb.ErrQuery)
	}
	row.LastEditedAt = *lastEditedAt

	return &row, nil
}

func (tokensView *Tokens) Update(tokenRow TokenRow) error {
	sql, sqlArgs, err := tokensView.rdb.StmtBuilder.Update(
		TOKENS_TABLE_NAME,
	).SetMap(map[string]interface{}{
		"drop":           tokenRow.Drop,
		"burned":         tokenRow.Burned,
		"name":           tokenRow.Name,
		"uri":            tokenRow.URI,
		"data":           tokenRow.Data,
		"owner":          tokenRow.Owner,
		"last_edited_at": tokensView.rdb.TypeConv.Tton(&tokenRow.LastEditedAt),
	}).Where(
		"denom_id = ? AND token_id = ?", tokenRow.DenomId, tokenRow.TokenId,
	).ToSql()
	if err != nil {
		return fmt.Errorf("error building NFT token update sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := tokensView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error updating NFT token into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error updating NFT token: no rows updated: %w", rdb.ErrWrite)
	}

	return nil
}

func (tokensView *Tokens) List(
	filter TokenListFilter,
	order TokenListOrder,
	pagination *pagination_interface.Pagination,
) ([]TokenRow, *pagination_interface.PaginationResult, error) {
	stmtBuilder := tokensView.rdb.StmtBuilder.Select(
		"denom_id",
		"token_id",
		"drop",
		"burned",
		"name",
		"uri",
		"data",
		"minter",
		"owner",
		"minted_at",
		"last_edited_at",
	).From(
		TOKENS_TABLE_NAME,
	)

	if filter.MaybeDenomId != nil {
		stmtBuilder = stmtBuilder.Where("denom_id = ?", *filter.MaybeDenomId)
	}
	if filter.MaybeDrop != nil {
		stmtBuilder = stmtBuilder.Where("drop = ?", *filter.MaybeDrop)
	}
	if filter.MaybeMinter != nil {
		stmtBuilder = stmtBuilder.Where("minter = ?", *filter.MaybeMinter)
	}
	if filter.MaybeOwner != nil {
		stmtBuilder = stmtBuilder.Where("owner = ?", *filter.MaybeOwner)
	}

	if order.MintedAt == view.ORDER_DESC {
		stmtBuilder = stmtBuilder.OrderBy("minted_at DESC")
	} else {
		stmtBuilder = stmtBuilder.OrderBy("minted_at")
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		tokensView.rdb,
	).WithCustomTotalQueryFn(
		func(rdbHandle *rdb.Handle, _ sq.SelectBuilder) (int64, error) {
			totalView := NewTokensTotal(rdbHandle)

			denomIdIdentifier := "-"
			dropIdentifier := "-"
			minterIdentifier := "-"
			ownerIdentifier := "-"
			if filter.MaybeDenomId != nil {
				denomIdIdentifier = *filter.MaybeDenomId
			}
			if filter.MaybeDrop != nil {
				dropIdentifier = *filter.MaybeDrop
			}
			if filter.MaybeMinter != nil {
				minterIdentifier = *filter.MaybeMinter
			}
			if filter.MaybeOwner != nil {
				ownerIdentifier = *filter.MaybeOwner
			}
			identifier := fmt.Sprintf(
				"%s:%s:%s:%s",
				denomIdIdentifier, dropIdentifier, minterIdentifier, ownerIdentifier,
			)

			total, err := totalView.FindBy(identifier)
			if err != nil {
				return int64(0), err
			}
			return total, nil
		},
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building NFT tokens list SQL: %v: %w", err, rdb.ErrPrepare)
	}

	rowsResult, err := tokensView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing NFT tokens list query: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	rows := make([]TokenRow, 0)
	for rowsResult.Next() {
		var row TokenRow
		mintedAtTimeReader := tokensView.rdb.NtotReader()
		lastEditedAtTimeReader := tokensView.rdb.NtotReader()

		if scanErr := rowsResult.Scan(
			&row.DenomId,
			&row.TokenId,
			&row.Drop,
			&row.Burned,
			&row.Name,
			&row.URI,
			&row.Data,
			&row.Minter,
			&row.Owner,
			mintedAtTimeReader.ScannableArg(),
			lastEditedAtTimeReader.ScannableArg(),
		); scanErr != nil {
			if errors.Is(scanErr, rdb.ErrNoRows) {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning NFT token row: %v: %w", scanErr, rdb.ErrQuery)
		}

		mintedAt, parseCreatedAtErr := mintedAtTimeReader.Parse()
		if parseCreatedAtErr != nil {
			return nil, nil, fmt.Errorf("error parsing NFT token minted time: %v: %w", parseCreatedAtErr, rdb.ErrQuery)
		}
		row.MintedAt = *mintedAt

		lastEditedAt, parseLastEditedAtErr := lastEditedAtTimeReader.Parse()
		if parseLastEditedAtErr != nil {
			return nil, nil, fmt.Errorf("error parsing NFT token last edited time: %v: %w", parseLastEditedAtErr, rdb.ErrQuery)
		}
		row.LastEditedAt = *lastEditedAt

		rows = append(rows, row)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return rows, paginationResult, nil
}

func (tokensView *Tokens) ListDrops(
	pagination *pagination_interface.Pagination,
) ([]string, *pagination_interface.PaginationResult, error) {
	stmtBuilder := tokensView.rdb.StmtBuilder.Select(
		"DISTINCT drop",
	).From(
		TOKENS_TABLE_NAME,
	).Where(
		"drop IS NOT NULL AND drop <> ''",
	).OrderBy("drop")

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		tokensView.rdb,
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building NFT drops list SQL: %v: %w", err, rdb.ErrPrepare)
	}

	rowsResult, err := tokensView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing NFT drops list query: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	rows := make([]string, 0)
	for rowsResult.Next() {
		var drop string
		if scanErr := rowsResult.Scan(
			&drop,
		); scanErr != nil {
			if errors.Is(scanErr, rdb.ErrNoRows) {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning NFT drop row: %v: %w", scanErr, rdb.ErrQuery)
		}

		rows = append(rows, drop)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return rows, paginationResult, nil
}

type TokenListFilter struct {
	MaybeDenomId *string
	MaybeDrop    *string
	MaybeMinter  *string
	MaybeOwner   *string
}

type TokenListOrder struct {
	MintedAt view.ORDER
}

type TokenRow struct {
	DenomId      string          `json:"denomId"`
	TokenId      string          `json:"tokenId"`
	Drop         string          `json:"drop"`
	Burned       bool            `json:"burned"`
	Name         string          `json:"name"`
	URI          string          `json:"uri"`
	Data         string          `json:"data"`
	Minter       string          `json:"minter"`
	Owner        string          `json:"owner"`
	MintedAt     utctime.UTCTime `json:"mintedAt"`
	LastEditedAt utctime.UTCTime `json:"lastEditedAt"`
}
