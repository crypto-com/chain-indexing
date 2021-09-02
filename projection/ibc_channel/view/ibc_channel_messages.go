package view

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type IBCChannelMessages struct {
	rdb *rdb.Handle
}

func NewIBCChannelMessages(handle *rdb.Handle) *IBCChannelMessages {
	return &IBCChannelMessages{
		handle,
	}
}

func (ibcClientsView *IBCChannelMessages) Insert(ibcChannelMessage *IBCChannelMessageRow) error {
	sql, sqlArgs, err := ibcClientsView.rdb.StmtBuilder.
		Insert("view_ibc_channel_messages").
		Columns(
			"channel_id",
			"block_height",
			"source_channel",
			"destination_channel",
			"denom",
			"amount",
			"success",
			"error",
			"message_type",
			"message",
			"updated_bonded_tokens",
		).
		Values(
			ibcChannelMessage.ChannelID,
			ibcChannelMessage.BlockHeight,
			ibcChannelMessage.SourceChannel,
			ibcChannelMessage.DestinationChannel,
			ibcChannelMessage.Denom,
			ibcChannelMessage.Amount,
			ibcChannelMessage.Success,
			ibcChannelMessage.Error,
			ibcChannelMessage.MessageType,
			ibcChannelMessage.Message,
			ibcChannelMessage.UpdatedBondedTokens,
		).
		ToSql()

	if err != nil {
		return fmt.Errorf("error building view_ibc_channel_messages insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := ibcClientsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting view_ibc_channel_messages into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting view_ibc_channel_messages into the table: no row inserted: %w", rdb.ErrWrite)
	}

	return nil
}

type IBCChannelMessageRow struct {
	ChannelID           string `json:"channelId"`
	BlockHeight         int64  `json:"blockHeight"`
	SourceChannel       string `json:"sourceChannel"`
	DestinationChannel  string `json:"destinationChannel"`
	Denom               string `json:"denom"`
	Amount              string `json:"amount"`
	Success             string `json:"success"`
	Error               string `json:"error"`
	MessageType         string `json:"messageType"`
	Message             string `json:"message"`
	UpdatedBondedTokens string `json:"updatedBondedTokens"`
}
