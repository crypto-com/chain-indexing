package view

import (
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"

	jsoniter "github.com/json-iterator/go"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/internal/utctime"
)

const MESSAGES_TABLE_NAME = "view_nft_messages"

type Messages struct {
	rdb *rdb.Handle
}

func NewMessages(handle *rdb.Handle) *Messages {
	return &Messages{
		handle,
	}
}

func (nftMessagesView *Messages) Insert(messageRow *MessageRow) error {
	nftMessageDataJSON, err := jsoniter.MarshalToString(messageRow.Data)
	if err != nil {
		return fmt.Errorf("error JSON marshalling NFT message for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	stmtBuilder := nftMessagesView.rdb.StmtBuilder.Insert(
		MESSAGES_TABLE_NAME,
	).Columns(
		"block_height",
		"block_hash",
		"block_time",
		"denom_id",
		"maybe_token_id",
		"maybe_drop",
		"transaction_hash",
		"success",
		"message_index",
		"message_type",
		"data",
	)
	blockTime := nftMessagesView.rdb.Tton(&messageRow.BlockTime)

	stmtBuilder = stmtBuilder.Values(
		messageRow.BlockHeight,
		messageRow.BlockHash,
		blockTime,
		messageRow.DenomId,
		messageRow.MaybeTokenId,
		messageRow.MaybeDrop,
		messageRow.TransactionHash,
		messageRow.Success,
		messageRow.MessageIndex,
		messageRow.MessageType,
		nftMessageDataJSON,
	)
	sql, sqlArgs, err := stmtBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("error building NFT message id insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}
	result, err := nftMessagesView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting NFT messages into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf(
			"error inserting NFT messages into the table: no rows inserted: %w", rdb.ErrWrite,
		)
	}

	return nil
}

func (nftMessagesView *Messages) List(
	filter MessagesListFilter,
	order MessagesListOrder,
	pagination *pagination_interface.Pagination,
) ([]MessageRow, *pagination_interface.PaginationResult, error) {
	stmtBuilder := nftMessagesView.rdb.StmtBuilder.Select(
		"denom_id",
		"maybe_token_id",
		"maybe_drop",
		"block_height",
		"block_hash",
		"block_time",
		"transaction_hash",
		"success",
		"message_index",
		"message_type",
		"data",
	).From(
		MESSAGES_TABLE_NAME,
	)

	if filter.MaybeDenomId != nil {
		stmtBuilder = stmtBuilder.Where("denom_id = ?", *filter.MaybeDenomId)
	}
	if filter.MaybeTokenId != nil {
		stmtBuilder = stmtBuilder.Where("maybe_token_id = ?", *filter.MaybeTokenId)
	}
	if filter.MaybeDrop != nil {
		stmtBuilder = stmtBuilder.Where("maybe_drop = ?", *filter.MaybeDrop)
	}
	if filter.MaybeMsgTypes != nil {
		stmtBuilder = stmtBuilder.Where(sq.Eq{"view_nft_messages.message_type": filter.MaybeMsgTypes})
	}

	if order.Id == view.ORDER_DESC {
		stmtBuilder = stmtBuilder.OrderBy("id DESC")
	} else {
		stmtBuilder = stmtBuilder.OrderBy("id")
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		nftMessagesView.rdb,
	).WithCustomTotalQueryFn(
		func(rdbHandle *rdb.Handle, _ sq.SelectBuilder) (int64, error) {
			var totalIdentities []string
			drop := "-"
			denomId := "-"
			tokenId := "-"
			if filter.MaybeDrop != nil {
				drop = *filter.MaybeDrop
			}
			if filter.MaybeDenomId != nil {
				denomId = *filter.MaybeDenomId
			}
			if filter.MaybeTokenId != nil {
				tokenId = *filter.MaybeTokenId
			}

			if filter.MaybeMsgTypes == nil {
				totalIdentities = []string{fmt.Sprintf("%s:%s:%s:-", drop, denomId, tokenId)}
			} else {
				totalIdentities = make([]string, 0)
				for _, msgType := range filter.MaybeMsgTypes {
					totalIdentities = append(totalIdentities, fmt.Sprintf("%s:%s:%s:%s", drop, denomId, tokenId, msgType))
				}
			}

			totalView := NewMessagesTotal(rdbHandle)
			total, err := totalView.SumBy(totalIdentities)
			if err != nil {
				return int64(0), err
			}
			return total, nil
		},
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf(
			"error building nft messages select SQL: %v, %w", err, rdb.ErrBuildSQLStmt,
		)
	}

	rowsResult, err := nftMessagesView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing nft messages select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	nftMessages := make([]MessageRow, 0)
	for rowsResult.Next() {
		var nftMessage MessageRow
		var accountMessageDataJSON *string
		blockTimeReader := nftMessagesView.rdb.NtotReader()

		if err = rowsResult.Scan(
			&nftMessage.DenomId,
			&nftMessage.MaybeTokenId,
			&nftMessage.MaybeDrop,
			&nftMessage.BlockHeight,
			&nftMessage.BlockHash,
			blockTimeReader.ScannableArg(),
			&nftMessage.TransactionHash,
			&nftMessage.Success,
			&nftMessage.MessageIndex,
			&nftMessage.MessageType,
			&accountMessageDataJSON,
		); err != nil {
			if errors.Is(err, rdb.ErrNoRows) {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning nft message row: %v: %w", err, rdb.ErrQuery)
		}
		blockTime, parseErr := blockTimeReader.Parse()
		if parseErr != nil {
			return nil, nil, fmt.Errorf(
				"error parsing nft message block time: %v: %w", parseErr, rdb.ErrQuery,
			)
		}
		nftMessage.BlockTime = *blockTime

		var data interface{}
		if unmarshalErr := jsoniter.Unmarshal([]byte(*accountMessageDataJSON), &data); unmarshalErr != nil {
			return nil, nil, fmt.Errorf(
				"error unmarshalling nft message data JSON: %v: %w", unmarshalErr, rdb.ErrQuery,
			)
		}
		nftMessage.Data = data

		nftMessages = append(nftMessages, nftMessage)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return nftMessages, paginationResult, nil
}

type MessageRow struct {
	DenomId         string          `json:"denomId"`
	MaybeTokenId    *string         `json:"tokenId"`
	MaybeDrop       *string         `json:"drop"`
	BlockHeight     int64           `json:"blockHeight"`
	BlockHash       string          `json:"blockHash"`
	BlockTime       utctime.UTCTime `json:"blockTime"`
	TransactionHash string          `json:"transactionHash"`
	Success         bool            `json:"success"`
	MessageIndex    int             `json:"messageIndex"`
	MessageType     string          `json:"messageType"`
	Data            interface{}     `json:"data"`
}

type MessagesListFilter struct {
	MaybeDenomId  *string
	MaybeTokenId  *string
	MaybeDrop     *string
	MaybeMsgTypes []string
}

type MessagesListOrder struct {
	Id view.ORDER
}
