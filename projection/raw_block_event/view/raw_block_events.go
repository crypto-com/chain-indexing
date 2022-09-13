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

// RawBlockEvents projection view implemented by relational database
type RawBlockEvents struct {
	rdb *rdb.Handle
}

func NewRawBlockEvents(handle *rdb.Handle) *RawBlockEvents {
	return &RawBlockEvents{
		handle,
	}
}

func (eventsView *RawBlockEvents) Insert(rawBlockEvent *RawBlockEventRow) error {
	var err error

	var sql string
	sql, _, err = eventsView.rdb.StmtBuilder.Insert(
		"view_raw_block_events",
	).Columns(
		"block_height",
		"block_hash",
		"block_time",
		"from_result",
		"raw_data",
	).Values("?", "?", "?", "?", "?").ToSql()
	if err != nil {
		return fmt.Errorf("error building events insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	var rawBlockEventDataJSON string
	if rawBlockEventDataJSON, err = jsoniter.MarshalToString(rawBlockEvent.RawData); err != nil {
		return fmt.Errorf("error JSON marshalling raw block event data for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := eventsView.rdb.Exec(sql,
		rawBlockEvent.BlockHeight,
		rawBlockEvent.BlockHash,
		eventsView.rdb.Tton(&rawBlockEvent.BlockTime),
		rawBlockEvent.FromResult,
		rawBlockEventDataJSON,
	)
	if err != nil {
		return fmt.Errorf("error inserting raw block event data into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting raw block event data into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (eventsView *RawBlockEvents) InsertAll(rawBlockEvents []RawBlockEventRow) error {
	if len(rawBlockEvents) == 0 {
		return nil
	}

	pendingRowCount := 0
	var stmtBuilder sq.InsertBuilder

	rawBlockEventCount := len(rawBlockEvents)
	for i, rawBlockEvent := range rawBlockEvents {

		if pendingRowCount == 0 {
			stmtBuilder = eventsView.rdb.StmtBuilder.Insert(
				"view_raw_block_events",
			).Columns(
				"block_height",
				"block_hash",
				"block_time",
				"from_result",
				"raw_data",
			)
		}

		rawBlockEventDataJSON, marshalErr := jsoniter.MarshalToString(rawBlockEvent.RawData)
		if marshalErr != nil {
			return fmt.Errorf(
				"error JSON marshalling raw block event data for insertion: %v: %w",
				marshalErr, rdb.ErrBuildSQLStmt,
			)
		}

		stmtBuilder = stmtBuilder.Values(
			rawBlockEvent.BlockHeight,
			rawBlockEvent.BlockHash,
			eventsView.rdb.Tton(&rawBlockEvent.BlockTime),
			rawBlockEvent.FromResult,
			rawBlockEventDataJSON,
		)
		pendingRowCount += 1

		// Postgres has a limit of 65536 parameters.
		if pendingRowCount == 500 || i+1 == rawBlockEventCount {
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

func (eventsView *RawBlockEvents) FindById(id int64) (*RawBlockEventRow, error) {
	var err error

	selectStmtBuilder := eventsView.rdb.StmtBuilder.Select(
		"id",
		"block_height",
		"block_hash",
		"block_time",
		"from_result",
		"raw_data",
	).From(
		"view_raw_block_events",
	).Where(
		"id = ?", id,
	)

	sql, sqlArgs, err := selectStmtBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building raw block events selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	var rawBlockEvent RawBlockEventRow
	var rawBlockEventDataJSON *string
	blockTimeReader := eventsView.rdb.NtotReader()

	if err = eventsView.rdb.QueryRow(sql, sqlArgs...).Scan(
		&rawBlockEvent.MaybeId,
		&rawBlockEvent.BlockHeight,
		&rawBlockEvent.BlockHash,
		blockTimeReader.ScannableArg(),
		&rawBlockEvent.FromResult,
		&rawBlockEventDataJSON,
	); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return nil, rdb.ErrNoRows
		}
		return nil, fmt.Errorf("error scanning raw block event row: %v: %w", err, rdb.ErrQuery)
	}
	blockTime, parseErr := blockTimeReader.Parse()
	if parseErr != nil {
		return nil, fmt.Errorf("error parsing raw block event time: %v: %w", parseErr, rdb.ErrQuery)
	}
	rawBlockEvent.BlockTime = *blockTime

	var rawBlockEventData RawBlockEventRowData
	if unmarshalErr := jsoniter.Unmarshal([]byte(*rawBlockEventDataJSON), &rawBlockEventData); unmarshalErr != nil {
		return nil, fmt.Errorf("error unmarshalling raw block event data JSON: %v: %w", unmarshalErr, rdb.ErrQuery)
	}
	rawBlockEvent.RawData = rawBlockEventData

	return &rawBlockEvent, nil
}

func (eventsView *RawBlockEvents) List(
	filter RawBlockEventsListFilter,
	order RawBlockEventsListOrder,
	pagination *pagination_interface.Pagination,
) ([]RawBlockEventRow, *pagination_interface.PaginationResult, error) {
	stmtBuilder := eventsView.rdb.StmtBuilder.Select(
		"id",
		"block_height",
		"block_hash",
		"block_time",
		"from_result",
		"raw_data",
	).From(
		"view_raw_block_events",
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
			totalView := NewRawBlockEventsTotal(rdbHandle)
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

	rawBlockEvents := make([]RawBlockEventRow, 0)
	for rowsResult.Next() {
		var rawBlockEvent RawBlockEventRow
		var rawBlockEventDataJSON *string
		blockTimeReader := eventsView.rdb.NtotReader()

		if err = rowsResult.Scan(
			&rawBlockEvent.MaybeId,
			&rawBlockEvent.BlockHeight,
			&rawBlockEvent.BlockHash,
			blockTimeReader.ScannableArg(),
			&rawBlockEvent.FromResult,
			&rawBlockEventDataJSON,
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
		rawBlockEvent.BlockTime = *blockTime

		var rawBlockEventData RawBlockEventRowData
		if unmarshalErr := jsoniter.Unmarshal([]byte(*rawBlockEventDataJSON), &rawBlockEventData); unmarshalErr != nil {
			return nil, nil, fmt.Errorf("error unmarshalling raw block event data JSON: %v: %w", unmarshalErr, rdb.ErrQuery)
		}
		rawBlockEvent.RawData = rawBlockEventData

		rawBlockEvents = append(rawBlockEvents, rawBlockEvent)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return rawBlockEvents, paginationResult, nil
}

type RawBlockEventsListFilter struct {
	MaybeBlockHeight *int64
}

type RawBlockEventsListOrder struct {
	Height view.ORDER
}

type RawBlockEventRow struct {
	MaybeId     *int64               `json:"id"`
	BlockHeight int64                `json:"blockHeight"`
	BlockHash   string               `json:"blockHash"`
	BlockTime   utctime.UTCTime      `json:"blockTime"`
	FromResult  string               `json:"fromResult"`
	RawData     RawBlockEventRowData `json:"rawData"`
}

type RawBlockEventRowData struct {
	Type    string                  `json:"type"`
	Content model.BlockResultsEvent `json:"content"`
}
