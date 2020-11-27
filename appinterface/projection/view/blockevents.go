package view

import (
	"fmt"
	"time"

	pagination_interface "github.com/crypto-com/chainindex/appinterface/pagination"
	"github.com/crypto-com/chainindex/appinterface/rdb"
	"github.com/crypto-com/chainindex/internal/utctime"
	jsoniter "github.com/json-iterator/go"
)

// BlockEvents projection view implemented by relational database
type BlockEvents struct {
	rdb *rdb.Handle
}

func NewBlockEvents(handle *rdb.Handle) *BlockEvents {
	return &BlockEvents{
		handle,
	}
}

func (view *BlockEvents) Insert(blockEvent *BlockEventRow) error {
	var err error

	var sql string
	sql, _, err = view.rdb.StmtBuilder.Insert(
		"view_block_events",
	).Columns(
		"block_height",
		"block_hash",
		"block_time",
		"data",
	).Values("?", "?", "?", "?").ToSql()
	if err != nil {
		return fmt.Errorf("error building events insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	var blockEventDataJSON string
	if blockEventDataJSON, err = jsoniter.MarshalToString(blockEvent.Data); err != nil {
		return fmt.Errorf("error JSON marshalling block transation messages for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := view.rdb.Exec(sql,
		blockEvent.BlockHeight,
		blockEvent.BlockHash,
		view.rdb.Tton(&blockEvent.BlockTime),
		blockEventDataJSON,
	)
	if err != nil {
		return fmt.Errorf("error inserting block transaction into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting block transaction into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (view *BlockEvents) InsertAll(blockEvents []BlockEventRow) error {
	if len(blockEvents) == 0 {
		return nil
	}

	stmtBuilder := view.rdb.StmtBuilder.Insert(
		"view_block_events",
	).Columns(
		"block_height",
		"block_hash",
		"block_time",
		"data",
	)
	for _, blockEvent := range blockEvents {
		blockEventDataJSON, marshalErr := jsoniter.MarshalToString(blockEvent.Data)
		if marshalErr != nil {
			return fmt.Errorf(
				"error JSON marshalling block transation messages for insertion: %v: %w",
				marshalErr, rdb.ErrBuildSQLStmt,
			)
		}

		stmtBuilder = stmtBuilder.Values(
			blockEvent.BlockHeight,
			blockEvent.BlockHash,
			view.rdb.Tton(&blockEvent.BlockTime),
			blockEventDataJSON,
		)
	}

	sql, sqlArgs, err := stmtBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("error building events batch insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := view.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting block transaction into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != int64(len(blockEvents)) {
		return fmt.Errorf(
			"error batch inserting block transaction into the table: mismatched number of rows inserted: %w",
			rdb.ErrWrite,
		)
	}

	return nil
}

func (view *BlockEvents) FindById(id int64) (*BlockEventRow, error) {
	var err error

	selectStmtBuilder := view.rdb.StmtBuilder.Select(
		"id",
		"block_height",
		"block_hash",
		"block_time",
		"data",
	).From(
		"view_block_events",
	).Where(
		"id = ?", id,
	)

	sql, sqlArgs, err := selectStmtBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building block events selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	var blockEvent BlockEventRow
	var blockEventDataJSON *string
	blockTimeReader := view.rdb.NtotReader()

	if err = view.rdb.QueryRow(sql, sqlArgs...).Scan(
		&blockEvent.MaybeId,
		&blockEvent.BlockHeight,
		&blockEvent.BlockHash,
		blockTimeReader.ScannableArg(),
		&blockEventDataJSON,
	); err != nil {
		if err == rdb.ErrNoRows {
			return nil, rdb.ErrNoRows
		}
		return nil, fmt.Errorf("error scanning transaction row: %v: %w", err, rdb.ErrQuery)
	}
	blockTime, parseErr := blockTimeReader.Parse()
	if parseErr != nil {
		return nil, fmt.Errorf("error parsing transaction block time: %v: %w", parseErr, rdb.ErrQuery)
	}
	blockEvent.BlockTime = *blockTime

	var blockEventData BlockEventRowData
	if unmarshalErr := jsoniter.Unmarshal([]byte(*blockEventDataJSON), &blockEventData); unmarshalErr != nil {
		return nil, fmt.Errorf("error unmarshalling transaction messages JSON: %v: %w", unmarshalErr, rdb.ErrQuery)
	}
	blockEvent.Data = blockEventData

	return &blockEvent, nil
}

func (view *BlockEvents) List(
	filter BlockEventsListFilter,
	pagination *pagination_interface.Pagination,
) ([]BlockEventRow, *pagination_interface.PaginationResult, error) {
	stmtBuilder := view.rdb.StmtBuilder.Select(
		"id",
		"block_height",
		"block_hash",
		"block_time",
		"data",
	).From(
		"view_block_events",
	).OrderBy(
		"id",
	)

	if filter.MaybeBlockHeight != nil {
		stmtBuilder = stmtBuilder.Where("block_height = ?", *filter.MaybeBlockHeight)
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		view.rdb.Runner,
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building transactions select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := view.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing transactions select SQL: %v: %w", err, rdb.ErrQuery)
	}

	blockEvents := make([]BlockEventRow, 0)
	for rowsResult.Next() {
		fmt.Println(time.Now())
		var blockEvent BlockEventRow
		var blockEventDataJSON *string
		blockTimeReader := view.rdb.NtotReader()

		if err = rowsResult.Scan(
			&blockEvent.MaybeId,
			&blockEvent.BlockHeight,
			&blockEvent.BlockHash,
			blockTimeReader.ScannableArg(),
			&blockEventDataJSON,
		); err != nil {
			if err == rdb.ErrNoRows {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning transaction row: %v: %w", err, rdb.ErrQuery)
		}
		blockTime, parseErr := blockTimeReader.Parse()
		if parseErr != nil {
			return nil, nil, fmt.Errorf("error parsing transaction block time: %v: %w", parseErr, rdb.ErrQuery)
		}
		blockEvent.BlockTime = *blockTime

		var blockEventData BlockEventRowData
		if unmarshalErr := jsoniter.Unmarshal([]byte(*blockEventDataJSON), &blockEventData); unmarshalErr != nil {
			return nil, nil, fmt.Errorf("error unmarshalling transaction messages JSON: %v: %w", unmarshalErr, rdb.ErrQuery)
		}
		blockEvent.Data = blockEventData

		blockEvents = append(blockEvents, blockEvent)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return blockEvents, paginationResult, nil
}

type BlockEventsListFilter struct {
	MaybeBlockHeight *int64
}

type BlockEventRow struct {
	MaybeId     *int64            `json:"id"`
	BlockHeight int64             `json:"blockHeight"`
	BlockHash   string            `json:"blockHash"`
	BlockTime   utctime.UTCTime   `json:"blockTime"`
	Data        BlockEventRowData `json:"data"`
}

type BlockEventRowData struct {
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
}
