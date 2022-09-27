package view

import (
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/json"
	"github.com/crypto-com/chain-indexing/external/utctime"
	jsoniter "github.com/json-iterator/go"
)

type AccountRawEvents interface {
	Insert(*AccountRawEventRow) error
	InsertAll([]AccountRawEventRow) error
	List(AccountRawEventsListFilter, AccountRawEventsListOrder, *pagination_interface.Pagination) ([]AccountRawEventRow, *pagination_interface.PaginationResult, error)
}

// AccountRawEventsView projection view implemented by relational database
type AccountRawEventsView struct {
	rdb *rdb.Handle
}

func NewAccountRawEventsView(handle *rdb.Handle) AccountRawEvents {
	return &AccountRawEventsView{
		handle,
	}
}

func (eventsView *AccountRawEventsView) Insert(accountRawEvent *AccountRawEventRow) error {
	accountRawEventDataJSON, err := json.MarshalToString(accountRawEvent.Data)
	if err != nil {
		return fmt.Errorf("error JSON marshalling account raw event for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}
	stmtBuilder := eventsView.rdb.StmtBuilder.Insert(
		"view_account_raw_events",
	).Columns(
		"account",
		"block_height",
		"block_hash",
		"block_time",
		"data",
	)

	blockTime := eventsView.rdb.Tton(&accountRawEvent.BlockTime)
	stmtBuilder = stmtBuilder.Values(
		accountRawEvent.Account,
		accountRawEvent.BlockHeight,
		accountRawEvent.BlockHash,
		blockTime,
		accountRawEventDataJSON,
	)
	sql, sqlArgs, err := stmtBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("error building account raw event id insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}
	result, err := eventsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting ccount raw event into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != int64(1) {
		return fmt.Errorf(
			"error inserting ccount raw event into the table: mismatched rows inserted: %w", rdb.ErrWrite,
		)
	}

	return nil
}

func (eventsView *AccountRawEventsView) InsertAll(accountRawEvents []AccountRawEventRow) error {
	if len(accountRawEvents) == 0 {
		return nil
	}

	pendingRowCount := 0
	var stmtBuilder sq.InsertBuilder

	blockEventCount := len(accountRawEvents)
	for i, accountRawEvent := range accountRawEvents {
		if pendingRowCount == 0 {
			stmtBuilder = eventsView.rdb.StmtBuilder.Insert(
				"view_account_raw_events",
			).Columns(
				"account",
				"block_height",
				"block_hash",
				"block_time",
				"data",
			)
		}

		accountRawEventDataJSON, marshalErr := jsoniter.MarshalToString(accountRawEvent.Data)
		if marshalErr != nil {
			return fmt.Errorf(
				"error JSON marshalling account raw event data for insertion: %v: %w",
				marshalErr, rdb.ErrBuildSQLStmt,
			)
		}

		stmtBuilder = stmtBuilder.Values(
			accountRawEvent.Account,
			accountRawEvent.BlockHeight,
			accountRawEvent.BlockHash,
			eventsView.rdb.Tton(&accountRawEvent.BlockTime),
			accountRawEventDataJSON,
		)
		pendingRowCount += 1

		// Postgres has a limit of 65536 parameters.
		if pendingRowCount == 500 || i+1 == blockEventCount {
			sql, sqlArgs, err := stmtBuilder.ToSql()
			if err != nil {
				return fmt.Errorf("error building block events batch insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
			}

			result, err := eventsView.rdb.Exec(sql, sqlArgs...)
			if err != nil {
				return fmt.Errorf("error batch inserting block events into the table: %v: %w", err, rdb.ErrWrite)
			}
			if result.RowsAffected() != int64(pendingRowCount) {
				return fmt.Errorf(
					"error batch inserting block events into the table: mismatched number of rows inserted: %w",
					rdb.ErrWrite,
				)
			}
			pendingRowCount = 0
		}

	}

	return nil
}

func (eventsView *AccountRawEventsView) FindById(id int64) (*AccountRawEventRow, error) {
	var err error

	selectStmtBuilder := eventsView.rdb.StmtBuilder.Select(
		"id",
		"account",
		"block_height",
		"block_hash",
		"block_time",
		"data",
	).From(
		"view_account_raw_events",
	).Where(
		"id = ?", id,
	)

	sql, sqlArgs, err := selectStmtBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building block events selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	var accountRawEvent AccountRawEventRow
	var accountRawEventDataJSON *string
	blockTimeReader := eventsView.rdb.NtotReader()

	if err = eventsView.rdb.QueryRow(sql, sqlArgs...).Scan(
		&accountRawEvent.MaybeId,
		&accountRawEvent.Account,
		&accountRawEvent.BlockHeight,
		&accountRawEvent.BlockHash,
		blockTimeReader.ScannableArg(),
		&accountRawEventDataJSON,
	); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return nil, rdb.ErrNoRows
		}
		return nil, fmt.Errorf("error scanning account raw event row: %v: %w", err, rdb.ErrQuery)
	}
	blockTime, parseErr := blockTimeReader.Parse()
	if parseErr != nil {
		return nil, fmt.Errorf("error parsing account raw event time: %v: %w", parseErr, rdb.ErrQuery)
	}
	accountRawEvent.BlockTime = *blockTime

	var accountRawEventData AccountRawEventRowData
	if unmarshalErr := jsoniter.Unmarshal([]byte(*accountRawEventDataJSON), &accountRawEventData); unmarshalErr != nil {
		return nil, fmt.Errorf("error unmarshalling account raw event data JSON: %v: %w", unmarshalErr, rdb.ErrQuery)
	}
	accountRawEvent.Data = accountRawEventData

	return &accountRawEvent, nil
}

func (eventsView *AccountRawEventsView) List(
	filter AccountRawEventsListFilter,
	order AccountRawEventsListOrder,
	pagination *pagination_interface.Pagination,
) ([]AccountRawEventRow, *pagination_interface.PaginationResult, error) {
	stmtBuilder := eventsView.rdb.StmtBuilder.Select(
		"id",
		"account",
		"block_height",
		"block_hash",
		"block_time",
		"data",
	).From(
		"view_account_raw_events",
	).Where(
		"view_account_raw_events.account = ?", filter.Account,
	)

	var totalIdentities []string
	if filter.MaybeMsgTypes == nil {
		totalIdentities = []string{fmt.Sprintf("%s:-", filter.Account)}
	}

	if order.Id == view.ORDER_DESC {
		stmtBuilder = stmtBuilder.OrderBy("id DESC")
	} else {
		stmtBuilder = stmtBuilder.OrderBy("id")
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		eventsView.rdb,
	).WithCustomTotalQueryFn(
		func(rdbHandle *rdb.Handle, _ sq.SelectBuilder) (int64, error) {
			totalView := NewAccountRawEventsTotalView(rdbHandle)
			total, err := totalView.SumBy(totalIdentities)
			if err != nil {
				return int64(0), err
			}
			return total, nil
		},
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building block events select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := eventsView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing block events select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	accountRawEvents := make([]AccountRawEventRow, 0)
	for rowsResult.Next() {
		var accountRawEvent AccountRawEventRow
		var accountRawEventDataJSON *string
		blockTimeReader := eventsView.rdb.NtotReader()

		if err = rowsResult.Scan(
			&accountRawEvent.MaybeId,
			&accountRawEvent.Account,
			&accountRawEvent.BlockHeight,
			&accountRawEvent.BlockHash,
			blockTimeReader.ScannableArg(),
			&accountRawEventDataJSON,
		); err != nil {
			if errors.Is(err, rdb.ErrNoRows) {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning account raw event row: %v: %w", err, rdb.ErrQuery)
		}
		blockTime, parseErr := blockTimeReader.Parse()
		if parseErr != nil {
			return nil, nil, fmt.Errorf("error parsing account raw event time: %v: %w", parseErr, rdb.ErrQuery)
		}
		accountRawEvent.BlockTime = *blockTime

		var accountRawEventData AccountRawEventRowData
		if unmarshalErr := jsoniter.Unmarshal([]byte(*accountRawEventDataJSON), &accountRawEventData); unmarshalErr != nil {
			return nil, nil, fmt.Errorf("error unmarshalling account raw event data JSON: %v: %w", unmarshalErr, rdb.ErrQuery)
		}
		accountRawEvent.Data = accountRawEventData

		accountRawEvents = append(accountRawEvents, accountRawEvent)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return accountRawEvents, paginationResult, nil
}

type AccountRawEventsListFilter struct {
	// Required account filter
	Account string

	// Optional filtering
	MaybeMsgTypes []string
}

type AccountRawEventsListOrder struct {
	Id view.ORDER
}

type AccountRawEventRow struct {
	MaybeId     *int64                 `json:"id"`
	Account     string                 `json:"account"`
	BlockHeight int64                  `json:"blockHeight"`
	BlockHash   string                 `json:"blockHash"`
	BlockTime   utctime.UTCTime        `json:"blockTime"`
	Data        AccountRawEventRowData `json:"data"`
}

type AccountRawEventRowData struct {
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
}
