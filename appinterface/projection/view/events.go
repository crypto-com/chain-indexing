package view

import (
	"fmt"

	pagination_interface "github.com/crypto-com/chainindex/appinterface/pagination"
	"github.com/crypto-com/chainindex/appinterface/rdb"
	"github.com/crypto-com/chainindex/internal/utctime"
	jsoniter "github.com/json-iterator/go"
)

// Events projection view implemented by relational database
type Events struct {
	rdb *rdb.Handle
}

func NewEvents(handle *rdb.Handle) *Events {
	return &Events{
		handle,
	}
}

type EventsListFilter struct {
	MaybeBlockHeight *int64
}

func (view *Events) Insert(transaction *Transaction) error {
	var err error

	var sql string
	sql, _, err = view.rdb.StmtBuilder.Insert(
		"view_events",
	).Columns(
		"block_height",
		"block_hash",
		"block_time",
		"events",
	).Values("?", "?", "?", "?").ToSql()
	if err != nil {
		return fmt.Errorf("error building events insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	var eventsJSON string
	if eventsJSON, err = jsoniter.MarshalToString(transaction.Messages); err != nil {
		return fmt.Errorf("error JSON marshalling block transation messages for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := view.rdb.Exec(sql,
		transaction.BlockHeight,
		transaction.BlockHash,
		view.rdb.Tton(&transaction.BlockTime),
		eventsJSON,
	)
	if err != nil {
		return fmt.Errorf("error inserting block transaction into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting block transaction into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (view *Events) List(
	filter EventsListFilter,
	pagination *pagination_interface.Pagination,
) ([]Event, *pagination_interface.PaginationResult, error) {
	stmtBuilder := view.rdb.StmtBuilder.Select(
		"block_height",
		"block_hash",
		"block_time",
		"events",
	).From(
		"view_events",
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

	events := make([]Event, 0)
	for rowsResult.Next() {
		var event Event
		var eventsJSON *string
		blockTimeReader := view.rdb.NtotReader()
		var fee string

		if err = rowsResult.Scan(
			&event.BlockHeight,
			&event.BlockHash,
			blockTimeReader.ScannableArg(),
			&eventsJSON,
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
		event.BlockTime = *blockTime

		var eventContents []EventContent
		if unmarshalErr := jsoniter.Unmarshal([]byte(*eventsJSON), &eventContents); unmarshalErr != nil {
			return nil, nil, fmt.Errorf("error unmarshalling transaction messages JSON: %v: %w", unmarshalErr, rdb.ErrQuery)
		}
		event.Events = eventContents

		events = append(events, event)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return events, paginationResult, nil
}

type Event struct {
	BlockHeight int64           `json:"blockHeight"`
	BlockHash   string          `json:"blockHash"`
	BlockTime   utctime.UTCTime `json:"blockTime"`
	Events      []EventContent  `json:"events"`
}

type EventContent struct {
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
}
