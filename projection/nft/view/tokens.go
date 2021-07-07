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
		"name",
		"uri",
		"data",
		"minter",
		"owner",
		"minted_at",
		"minted_at_block_height",
		"last_edited_at",
		"last_edited_at_block_height",
		"last_transferred_at",
		"last_transferred_at_block_height",
	).Values(
		tokenRow.DenomId,
		tokenRow.TokenId,
		tokenRow.MaybeDrop,
		tokenRow.Name,
		tokenRow.URI,
		tokenRow.Data,
		tokenRow.Minter,
		tokenRow.Owner,
		tokensView.rdb.Tton(&tokenRow.MintedAt),
		tokenRow.MintedAtBlockHeight,
		tokensView.rdb.Tton(&tokenRow.LastEditedAt),
		tokenRow.LastEditedAtBlockHeight,
		tokensView.rdb.Tton(&tokenRow.LastTransferredAt),
		tokenRow.LastTransferredAtBlockHeight,
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

func (tokensView *Tokens) Delete(denomId string, tokenId string) (int64, error) {
	var err error

	sql, sqlArgs, err := tokensView.rdb.StmtBuilder.Delete(
		TOKENS_TABLE_NAME,
	).Where(
		"denom_id = ? AND token_id = ?", denomId, tokenId,
	).ToSql()
	if err != nil {
		return 0, fmt.Errorf("error building NFT token deletion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := tokensView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return 0, fmt.Errorf("error deleting NFT token from the table: %v: %w", err, rdb.ErrWrite)
	}

	return result.RowsAffected(), nil
}

func (tokensView *Tokens) FindById(
	denomId string, tokenId string,
) (*TokenRowWithDenomname, error) {
	selectStmtBuilder := tokensView.rdb.StmtBuilder.Select(
		fmt.Sprintf("%s.denom_id", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.name AS denom_name", DENOMS_TABLE_NAME),
		fmt.Sprintf("%s.schema AS denom_schema", DENOMS_TABLE_NAME),
		fmt.Sprintf("%s.token_id", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.drop", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.name", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.uri", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.data", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.minter", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.owner", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.minted_at", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.minted_at_block_height", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.last_edited_at", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.last_edited_at_block_height", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.last_transferred_at", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.last_transferred_at_block_height", TOKENS_TABLE_NAME),
	).From(
		TOKENS_TABLE_NAME,
	).LeftJoin(
		fmt.Sprintf(
			"%s ON %s.denom_id = %s.denom_id",
			DENOMS_TABLE_NAME, DENOMS_TABLE_NAME, TOKENS_TABLE_NAME,
		),
	).Where(
		fmt.Sprintf(
			"%s.denom_id = ? AND %s.token_id = ?",
			TOKENS_TABLE_NAME, TOKENS_TABLE_NAME,
		),
		denomId, tokenId,
	)

	sql, sqlArgs, err := selectStmtBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building NFT token selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	var row TokenRowWithDenomname
	mintedAtTimeReader := tokensView.rdb.NtotReader()
	lastEditedAtTimeReader := tokensView.rdb.NtotReader()
	lastTransferredAtTimeReader := tokensView.rdb.NtotReader()

	if err = tokensView.rdb.QueryRow(sql, sqlArgs...).Scan(
		&row.DenomId,
		&row.DenomName,
		&row.DenomSchema,
		&row.TokenId,
		&row.MaybeDrop,
		&row.Name,
		&row.URI,
		&row.Data,
		&row.Minter,
		&row.Owner,
		mintedAtTimeReader.ScannableArg(),
		&row.MintedAtBlockHeight,
		lastEditedAtTimeReader.ScannableArg(),
		&row.LastEditedAtBlockHeight,
		lastTransferredAtTimeReader.ScannableArg(),
		&row.LastTransferredAtBlockHeight,
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

	lastTransferredAt, parseLastTransferredAtErr := lastTransferredAtTimeReader.Parse()
	if parseLastTransferredAtErr != nil {
		return nil, fmt.Errorf("error parsing NFT token last transferred time: %v: %w", parseLastTransferredAtErr, rdb.ErrQuery)
	}
	row.LastTransferredAt = *lastTransferredAt

	return &row, nil
}

func (tokensView *Tokens) Update(tokenRow TokenRow) error {
	sql, sqlArgs, err := tokensView.rdb.StmtBuilder.Update(
		TOKENS_TABLE_NAME,
	).SetMap(map[string]interface{}{
		"drop":                             tokenRow.MaybeDrop,
		"name":                             tokenRow.Name,
		"uri":                              tokenRow.URI,
		"data":                             tokenRow.Data,
		"owner":                            tokenRow.Owner,
		"last_edited_at":                   tokensView.rdb.TypeConv.Tton(&tokenRow.LastEditedAt),
		"last_edited_at_block_height":      tokenRow.LastEditedAtBlockHeight,
		"last_transferred_at":              tokensView.rdb.TypeConv.Tton(&tokenRow.LastTransferredAt),
		"last_transferred_at_block_height": tokenRow.LastTransferredAtBlockHeight,
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
) ([]TokenRowWithDenomname, *pagination_interface.PaginationResult, error) {
	stmtBuilder := tokensView.rdb.StmtBuilder.Select(
		fmt.Sprintf("%s.denom_id", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.name AS denom_name", DENOMS_TABLE_NAME),
		fmt.Sprintf("%s.schema AS denom_schema", DENOMS_TABLE_NAME),
		fmt.Sprintf("%s.token_id", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.drop", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.name", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.uri", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.data", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.minter", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.owner", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.minted_at", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.minted_at_block_height", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.last_edited_at", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.last_edited_at_block_height", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.last_transferred_at", TOKENS_TABLE_NAME),
		fmt.Sprintf("%s.last_transferred_at_block_height", TOKENS_TABLE_NAME),
	).From(
		TOKENS_TABLE_NAME,
	).LeftJoin(
		fmt.Sprintf(
			"%s ON %s.denom_id = %s.denom_id",
			DENOMS_TABLE_NAME, DENOMS_TABLE_NAME, TOKENS_TABLE_NAME,
		),
	)

	if filter.MaybeDenomId != nil {
		stmtBuilder = stmtBuilder.Where(fmt.Sprintf("%s.denom_id = ?", DENOMS_TABLE_NAME), *filter.MaybeDenomId)
	}
	if filter.MaybeDrop != nil {
		stmtBuilder = stmtBuilder.Where(fmt.Sprintf("%s.drop = ?", TOKENS_TABLE_NAME), *filter.MaybeDrop)
	}
	if filter.MaybeMinter != nil {
		stmtBuilder = stmtBuilder.Where(fmt.Sprintf("%s.minter = ?", TOKENS_TABLE_NAME), *filter.MaybeMinter)
	}
	if filter.MaybeOwner != nil {
		stmtBuilder = stmtBuilder.Where(fmt.Sprintf("%s.owner = ?", TOKENS_TABLE_NAME), *filter.MaybeOwner)
	}

	if order.MintedAt == view.ORDER_DESC {
		stmtBuilder = stmtBuilder.OrderBy(fmt.Sprintf("%s.minted_at DESC", TOKENS_TABLE_NAME))
	} else if order.LastEditedAt == view.ORDER_ASC {
		stmtBuilder = stmtBuilder.OrderBy(fmt.Sprintf("%s.last_edited_at", TOKENS_TABLE_NAME))
	} else if order.LastEditedAt == view.ORDER_DESC {
		stmtBuilder = stmtBuilder.OrderBy(fmt.Sprintf("%s.last_edited_at DESC", TOKENS_TABLE_NAME))
	} else if order.LastTransferredAt == view.ORDER_ASC {
		stmtBuilder = stmtBuilder.OrderBy(fmt.Sprintf("%s.last_transferred_at", TOKENS_TABLE_NAME))
	} else if order.LastTransferredAt == view.ORDER_DESC {
		stmtBuilder = stmtBuilder.OrderBy(fmt.Sprintf("%s.last_transferred_at DESC", TOKENS_TABLE_NAME))
	} else {
		stmtBuilder = stmtBuilder.OrderBy(fmt.Sprintf("%s.minted_at", TOKENS_TABLE_NAME))
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

	rows := make([]TokenRowWithDenomname, 0)
	for rowsResult.Next() {
		var row TokenRowWithDenomname
		mintedAtTimeReader := tokensView.rdb.NtotReader()
		lastEditedAtTimeReader := tokensView.rdb.NtotReader()
		lastTransferredAtTimeReader := tokensView.rdb.NtotReader()

		if scanErr := rowsResult.Scan(
			&row.DenomId,
			&row.DenomName,
			&row.DenomSchema,
			&row.TokenId,
			&row.MaybeDrop,
			&row.Name,
			&row.URI,
			&row.Data,
			&row.Minter,
			&row.Owner,
			mintedAtTimeReader.ScannableArg(),
			&row.MintedAtBlockHeight,
			lastEditedAtTimeReader.ScannableArg(),
			&row.LastEditedAtBlockHeight,
			lastTransferredAtTimeReader.ScannableArg(),
			&row.LastTransferredAtBlockHeight,
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

		lastTransferredAt, parseLastTransferredAtErr := lastTransferredAtTimeReader.Parse()
		if parseLastTransferredAtErr != nil {
			return nil, nil, fmt.Errorf("error parsing NFT token last transferred time: %v: %w", parseLastTransferredAtErr, rdb.ErrQuery)
		}
		row.LastTransferredAt = *lastTransferredAt

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
	MintedAt          view.ORDER
	LastEditedAt      view.ORDER
	LastTransferredAt view.ORDER
}

type TokenRowWithDenomname struct {
	TokenRow

	DenomName   string `json:"denomName"`
	DenomSchema string `json:"denomSchema"`
}

type TokenRow struct {
	DenomId                      string          `json:"denomId"`
	TokenId                      string          `json:"tokenId"`
	MaybeDrop                    *string         `json:"drop"`
	Name                         string          `json:"tokenName"`
	URI                          string          `json:"tokenURI"`
	Data                         string          `json:"tokenData"`
	Minter                       string          `json:"tokenMinter"`
	Owner                        string          `json:"tokenOwner"`
	MintedAt                     utctime.UTCTime `json:"tokenMintedAt"`
	MintedAtBlockHeight          int64           `json:"tokenMintedAtBlockHeight"`
	LastEditedAt                 utctime.UTCTime `json:"tokenLastEditedAt"`
	LastEditedAtBlockHeight      int64           `json:"tokenLastEditedAtBlockHeight"`
	LastTransferredAt            utctime.UTCTime `json:"tokenLastTransferredAt"`
	LastTransferredAtBlockHeight int64           `json:"tokenLastTransferredAtBlockHeight"`
}
