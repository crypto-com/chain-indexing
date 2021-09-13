package view

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type IBCChannelTraces struct {
	rdb *rdb.Handle
}

func NewIBCChannelTraces(handle *rdb.Handle) *IBCChannelTraces {
	return &IBCChannelTraces{
		handle,
	}
}

func (ibcChannelTracesView *IBCChannelTraces) Insert(ibcChannelMessage *IBCChannelTraceRow) error {
	sql, sqlArgs, err := ibcChannelTracesView.rdb.StmtBuilder.
		Insert("view_ibc_channel_traces").
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
		return fmt.Errorf("error building view_ibc_channel_traces insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := ibcChannelTracesView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting view_ibc_channel_traces into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting view_ibc_channel_traces into the table: no row inserted: %w", rdb.ErrWrite)
	}

	return nil
}

type IBCChannelTraceRow struct {
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
