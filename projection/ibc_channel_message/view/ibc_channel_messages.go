package view

import (
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/internal/utctime"
	jsoniter "github.com/json-iterator/go"
)

type IBCChannelMessages struct {
	rdb *rdb.Handle
}

func NewIBCChannelMessages(handle *rdb.Handle) *IBCChannelMessages {
	return &IBCChannelMessages{
		handle,
	}
}

func (ibcChannelMessagesView *IBCChannelMessages) Insert(ibcChannelMessage *IBCChannelMessageRow) error {
	messageJSON, err := jsoniter.MarshalToString(ibcChannelMessage.Message)
	if err != nil {
		return fmt.Errorf("error JSON marshalling ibcChannelMessage.Message for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	sql, sqlArgs, err := ibcChannelMessagesView.rdb.StmtBuilder.
		Insert("view_ibc_channel_messages").
		Columns(
			"channel_id",
			"block_height",
			"block_time",
			"transaction_hash",
			"relayer",
			"success",
			"error",
			"sender",
			"receiver",
			"denom",
			"amount",
			"message_type",
			"message",
		).
		Values(
			ibcChannelMessage.ChannelID,
			ibcChannelMessage.BlockHeight,
			ibcChannelMessagesView.rdb.Tton(&ibcChannelMessage.BlockTime),
			ibcChannelMessage.TransactionHash,
			ibcChannelMessage.MaybeRelayer,
			ibcChannelMessage.MaybeSuccess,
			ibcChannelMessage.MaybeError,
			ibcChannelMessage.MaybeSender,
			ibcChannelMessage.MaybeReceiver,
			ibcChannelMessage.MaybeDenom,
			ibcChannelMessage.MaybeAmount,
			ibcChannelMessage.MessageType,
			messageJSON,
		).
		ToSql()

	if err != nil {
		return fmt.Errorf("error building view_ibc_channel_messages insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := ibcChannelMessagesView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting view_ibc_channel_messages into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting view_ibc_channel_messages into the table: no row inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (ibcChannelMessagesView *IBCChannelMessages) ListByChannelID(
	channelID string,
	order IBCChannelMessagesListOrder,
	filter IBCChannelMessagesListFilter,
	pagination *pagination.Pagination,
) (
	[]IBCChannelMessageRow,
	*pagination.PaginationResult,
	error,
) {
	stmtBuilder := ibcChannelMessagesView.rdb.StmtBuilder.Select(
		"channel_id",
		"block_height",
		"block_time",
		"transaction_hash",
		"relayer",
		"success",
		"error",
		"sender",
		"receiver",
		"denom",
		"amount",
		"message_type",
		"message",
	).From(
		"view_ibc_channel_messages",
	).Where(
		"channel_id = ?", channelID,
	)

	if order.BlockTime == view.ORDER_DESC {
		stmtBuilder = stmtBuilder.OrderBy("block_time DESC")
	} else {
		stmtBuilder = stmtBuilder.OrderBy("block_time")
	}

	totalIdentities := make([]string, 0)
	if filter.MaybeMsgTypes == nil {
		totalIdentities = []string{fmt.Sprintf("%s:-", channelID)}
	} else {
		stmtBuilder = stmtBuilder.Where(sq.Eq{"message_type": filter.MaybeMsgTypes})
		for _, msgType := range filter.MaybeMsgTypes {
			totalIdentities = append(totalIdentities, fmt.Sprintf("%s:%s", channelID, msgType))
		}
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		ibcChannelMessagesView.rdb,
	).WithCustomTotalQueryFn(
		func(rdbHandle *rdb.Handle, _ sq.SelectBuilder) (int64, error) {
			totalView := NewIBCChannelMessagesTotal(rdbHandle)
			total, err := totalView.SumBy(totalIdentities)
			if err != nil {
				return int64(0), err
			}
			return total, nil
		},
	).BuildStmt(stmtBuilder)

	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building IBCChannelMessage select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := ibcChannelMessagesView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing IBCChannelMessage select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	messages := make([]IBCChannelMessageRow, 0)
	for rowsResult.Next() {
		var message IBCChannelMessageRow
		var messageJSON string
		BlockTimeReader := ibcChannelMessagesView.rdb.NtotReader()
		if err = rowsResult.Scan(
			&message.ChannelID,
			&message.BlockHeight,
			BlockTimeReader.ScannableArg(),
			&message.TransactionHash,
			&message.MaybeRelayer,
			&message.MaybeSuccess,
			&message.MaybeError,
			&message.MaybeSender,
			&message.MaybeReceiver,
			&message.MaybeDenom,
			&message.MaybeAmount,
			&message.MessageType,
			&messageJSON,
		); err != nil {
			if errors.Is(err, rdb.ErrNoRows) {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning IBCChannelMessage row: %v: %w", err, rdb.ErrQuery)
		}

		blockTime, parseErr := BlockTimeReader.Parse()
		if parseErr != nil {
			return nil, nil, fmt.Errorf("error parsing IBCChannelMessage BlockTime: %v: %w", parseErr, rdb.ErrQuery)
		}
		message.BlockTime = *blockTime

		if err = jsoniter.UnmarshalFromString(messageJSON, &message.Message); err != nil {
			return nil, nil, fmt.Errorf("error unmarshalling IBCChannelMessage message JSON: %v: %w", err, rdb.ErrQuery)
		}

		messages = append(messages, message)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return messages, paginationResult, nil
}

type IBCChannelMessagesListOrder struct {
	BlockTime view.ORDER
}

type IBCChannelMessagesListFilter struct {
	MaybeMsgTypes []string
}

type IBCChannelMessageRow struct {
	ChannelID       string          `json:"channelId"`
	BlockHeight     int64           `json:"blockHeight"`
	BlockTime       utctime.UTCTime `json:"blockTime"`
	TransactionHash string          `json:"transactionHash"`
	MaybeRelayer    *string         `json:"relayer"`
	MaybeSuccess    *bool           `json:"success"`
	MaybeError      *string         `json:"error"`
	MaybeSender     *string         `json:"sender"`
	MaybeReceiver   *string         `json:"receiver"`
	MaybeDenom      *string         `json:"denom"`
	MaybeAmount     *string         `json:"amount"`
	MessageType     string          `json:"messageType"`
	Message         interface{}     `json:"message"`
}
