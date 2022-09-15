package view

import (
	"errors"
	"fmt"
	"strconv"

	sq "github.com/Masterminds/squirrel"

	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/usecase/model"
	jsoniter "github.com/json-iterator/go"
)

// BlockRawEvents projection view implemented by relational database
type BlockRawEvents struct {
	rdb *rdb.Handle
}

func NewBlockRawEvents(handle *rdb.Handle) *BlockRawEvents {
	return &BlockRawEvents{
		handle,
	}
}

func (eventsView *BlockRawEvents) Insert(blockRawEvent *BlockRawEventRow) error {
	var err error

	var sql string
	sql, _, err = eventsView.rdb.StmtBuilder.Insert(
		"view_block_raw_events",
	).Columns(
		"block_height",
		"block_hash",
		"block_time",
		"from_result",
		"data",
	).Values("?", "?", "?", "?", "?").ToSql()
	if err != nil {
		return fmt.Errorf("error building events insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	var blockRawEventDataJSON string
	if blockRawEventDataJSON, err = jsoniter.MarshalToString(blockRawEvent.Data); err != nil {
		return fmt.Errorf("error JSON marshalling block raw event data for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := eventsView.rdb.Exec(sql,
		blockRawEvent.BlockHeight,
		blockRawEvent.BlockHash,
		eventsView.rdb.Tton(&blockRawEvent.BlockTime),
		blockRawEvent.FromResult,
		blockRawEventDataJSON,
	)
	if err != nil {
		return fmt.Errorf("error inserting block raw event data into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting block raw event data into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (eventsView *BlockRawEvents) InsertAll(blockRawEvents []BlockRawEventRow) error {
	if len(blockRawEvents) == 0 {
		return nil
	}

	pendingRowCount := 0
	var stmtBuilder sq.InsertBuilder

	blockRawEventCount := len(blockRawEvents)
	for i, blockRawEvent := range blockRawEvents {

		if pendingRowCount == 0 {
			stmtBuilder = eventsView.rdb.StmtBuilder.Insert(
				"view_block_raw_events",
			).Columns(
				"block_height",
				"block_hash",
				"block_time",
				"from_result",
				"data",
			)
		}

		blockRawEventDataJSON, marshalErr := jsoniter.MarshalToString(blockRawEvent.Data)
		if marshalErr != nil {
			return fmt.Errorf(
				"error JSON marshalling block raw event data for insertion: %v: %w",
				marshalErr, rdb.ErrBuildSQLStmt,
			)
		}

		stmtBuilder = stmtBuilder.Values(
			blockRawEvent.BlockHeight,
			blockRawEvent.BlockHash,
			eventsView.rdb.Tton(&blockRawEvent.BlockTime),
			blockRawEvent.FromResult,
			blockRawEventDataJSON,
		)
		pendingRowCount += 1

		// Postgres has a limit of 65536 parameters.
		if pendingRowCount == 500 || i+1 == blockRawEventCount {
			sql, sqlArgs, err := stmtBuilder.ToSql()
			if err != nil {
				return fmt.Errorf("error building raw block events batch insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
			}

			result, err := eventsView.rdb.Exec(sql, sqlArgs...)
			if err != nil {
				return fmt.Errorf("error batch inserting raw block events into the table: %v: %w", err, rdb.ErrWrite)
			}
			if result.RowsAffected() != int64(pendingRowCount) {
				return fmt.Errorf(
					"error batch inserting raw block events into the table: mismatched number of rows inserted: %w",
					rdb.ErrWrite,
				)
			}
			pendingRowCount = 0
		}

	}

	return nil
}

func (eventsView *BlockRawEvents) FindById(id int64) (*BlockRawEventRow, error) {
	var err error

	selectStmtBuilder := eventsView.rdb.StmtBuilder.Select(
		"id",
		"block_height",
		"block_hash",
		"block_time",
		"from_result",
		"data",
	).From(
		"view_block_raw_events",
	).Where(
		"id = ?", id,
	)

	sql, sqlArgs, err := selectStmtBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building raw block events selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	var blockRawEvent BlockRawEventRow
	var blockRawEventDataJSON *string
	blockTimeReader := eventsView.rdb.NtotReader()

	if err = eventsView.rdb.QueryRow(sql, sqlArgs...).Scan(
		&blockRawEvent.MaybeId,
		&blockRawEvent.BlockHeight,
		&blockRawEvent.BlockHash,
		blockTimeReader.ScannableArg(),
		&blockRawEvent.FromResult,
		&blockRawEventDataJSON,
	); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return nil, rdb.ErrNoRows
		}
		return nil, fmt.Errorf("error scanning block raw event row: %v: %w", err, rdb.ErrQuery)
	}
	blockTime, parseErr := blockTimeReader.Parse()
	if parseErr != nil {
		return nil, fmt.Errorf("error parsing block raw event time: %v: %w", parseErr, rdb.ErrQuery)
	}
	blockRawEvent.BlockTime = *blockTime

	var blockRawEventData BlockRawEventRowData
	if unmarshalErr := jsoniter.Unmarshal([]byte(*blockRawEventDataJSON), &blockRawEventData); unmarshalErr != nil {
		return nil, fmt.Errorf("error unmarshalling block raw event data JSON: %v: %w", unmarshalErr, rdb.ErrQuery)
	}
	blockRawEvent.Data = blockRawEventData

	return &blockRawEvent, nil
}

func (eventsView *BlockRawEvents) List(
	filter BlockRawEventsListFilter,
	order BlockRawEventsListOrder,
	pagination *pagination_interface.Pagination,
) ([]BlockRawEventRow, *pagination_interface.PaginationResult, error) {
	stmtBuilder := eventsView.rdb.StmtBuilder.Select(
		"id",
		"block_height",
		"block_hash",
		"block_time",
		"from_result",
		"data",
	).From(
		"view_block_raw_events",
	)

	if order.Height == view.ORDER_DESC {
		stmtBuilder = stmtBuilder.OrderBy("block_height DESC, id DESC")
	} else {
		stmtBuilder = stmtBuilder.OrderBy("block_height, id")
	}

	if filter.MaybeBlockHeight != nil {
		stmtBuilder = stmtBuilder.Where("block_height = ?", *filter.MaybeBlockHeight)
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		eventsView.rdb,
	).WithCustomTotalQueryFn(
		func(rdbHandle *rdb.Handle, _ sq.SelectBuilder) (int64, error) {
			identity := "-"
			if filter.MaybeBlockHeight != nil {
				identity = strconv.FormatInt(*filter.MaybeBlockHeight, 10)
			}
			totalView := NewBlockRawEventsTotal(rdbHandle)
			total, err := totalView.FindBy(identity)
			if err != nil {
				return int64(0), err
			}
			return total, nil
		},
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building raw block events select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := eventsView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing raw block events select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	blockRawEvents := make([]BlockRawEventRow, 0)
	for rowsResult.Next() {
		var blockRawEvent BlockRawEventRow
		var blockRawEventDataJSON *string
		blockTimeReader := eventsView.rdb.NtotReader()

		if err = rowsResult.Scan(
			&blockRawEvent.MaybeId,
			&blockRawEvent.BlockHeight,
			&blockRawEvent.BlockHash,
			blockTimeReader.ScannableArg(),
			&blockRawEvent.FromResult,
			&blockRawEventDataJSON,
		); err != nil {
			if errors.Is(err, rdb.ErrNoRows) {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning block event row: %v: %w", err, rdb.ErrQuery)
		}
		blockTime, parseErr := blockTimeReader.Parse()
		if parseErr != nil {
			return nil, nil, fmt.Errorf("error parsing block event time: %v: %w", parseErr, rdb.ErrQuery)
		}
		blockRawEvent.BlockTime = *blockTime

		var blockRawEventData BlockRawEventRowData
		if unmarshalErr := jsoniter.Unmarshal([]byte(*blockRawEventDataJSON), &blockRawEventData); unmarshalErr != nil {
			return nil, nil, fmt.Errorf("error unmarshalling block raw event data JSON: %v: %w", unmarshalErr, rdb.ErrQuery)
		}
		blockRawEvent.Data = blockRawEventData

		blockRawEvents = append(blockRawEvents, blockRawEvent)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return blockRawEvents, paginationResult, nil
}

type BlockRawEventsListFilter struct {
	MaybeBlockHeight *int64
}

type BlockRawEventsListOrder struct {
	Height view.ORDER
}

type BlockRawEventRow struct {
	MaybeId     *int64               `json:"id"`
	BlockHeight int64                `json:"blockHeight"`
	BlockHash   string               `json:"blockHash"`
	BlockTime   utctime.UTCTime      `json:"blockTime"`
	FromResult  string               `json:"fromResult"`
	Data        BlockRawEventRowData `json:"data"`
}

type BlockRawEventRowData struct {
	Type    string                  `json:"type"`
	Content model.BlockResultsEvent `json:"content"`
}
