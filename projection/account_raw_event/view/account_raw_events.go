package view

import (
	"errors"
	"fmt"
	"strconv"

	sq "github.com/Masterminds/squirrel"

	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/json"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/usecase/model"
	jsoniter "github.com/json-iterator/go"
)

// AccountRawEvents projection view implemented by relational database
type AccountRawEvents struct {
	rdb *rdb.Handle
}

func NewAccountRawEvents(handle *rdb.Handle) *AccountRawEvents {
	return &AccountRawEvents{
		handle,
	}
}

func (eventsView *AccountRawEvents) Insert(accountRawEvent *AccountRawEventRow, accounts []string) error {
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
	for _, account := range accounts {
		stmtBuilder = stmtBuilder.Values(
			account,
			accountRawEvent.BlockHeight,
			accountRawEvent.BlockHash,
			blockTime,
			accountRawEventDataJSON,
		)
	}
	sql, sqlArgs, err := stmtBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("error building account raw event id insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}
	result, err := eventsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting ccount raw event into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != int64(len(accounts)) {
		return fmt.Errorf(
			"error inserting ccount raw event into the table: mismatched rows inserted: %w", rdb.ErrWrite,
		)
	}

	return nil
}

func (eventsView *AccountRawEvents) InsertAll(accountRawEvents []AccountRawEventRow) error {
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

func (eventsView *AccountRawEvents) FindById(id int64) (*AccountRawEventRow, error) {
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

func (eventsView *AccountRawEvents) List(
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
	)

	if order.Height == view.ORDER_DESC {
		stmtBuilder = stmtBuilder.OrderBy("account DESC, id DESC")
	} else {
		stmtBuilder = stmtBuilder.OrderBy("account, id")
	}

	if filter.MaybeAccount != nil {
		stmtBuilder = stmtBuilder.Where("account = ?", *filter.MaybeAccount)
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		eventsView.rdb,
	).WithCustomTotalQueryFn(
		func(rdbHandle *rdb.Handle, _ sq.SelectBuilder) (int64, error) {
			identity := "-"
			if filter.MaybeAccount != nil {
				identity = strconv.FormatInt(*filter.MaybeAccount, 10)
			}
			totalView := NewAccountRawEventsTotal(rdbHandle)
			total, err := totalView.FindBy(identity)
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
	MaybeAccount *int64
}

type AccountRawEventsListOrder struct {
	Height view.ORDER
}

type AccountRawEventRecord struct {
	Row       AccountRawEventRow
	Accounts []string
	Events 	 []model.BlockEvent
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
