package view

import (
	"errors"
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/internal/utctime"

	sq "github.com/Masterminds/squirrel"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	_ "github.com/crypto-com/chain-indexing/test/factory"
)

// Channels projection view implemented by relational database
type Channels struct {
	rdb *rdb.Handle
}

func NewChannels(handle *rdb.Handle) *Channels {
	return &Channels{
		handle,
	}
}

func (channelsView *Channels) Insert(channel *ChannelRow) error {
	sql, sqlArgs, err := channelsView.rdb.StmtBuilder.
		Insert(
			"view_channels",
		).
		Columns(
			"channel_id",
			"port_id",
			"connection_id",
			"counterparty_channel_id",
			"counterparty_port_id",
			"packet_ordering",
			"last_in_packet_sequence",
			"last_out_packet_sequence",
			"total_transfer_in_count",
			"total_transfer_out_count",
			"total_transfer_out_success_count",
			"total_transfer_out_success_rate",
			"last_activity_time",
			"bonded_tokens",
		).
		Values(
			channel.ChannelID,
			channel.PortID,
			channel.ConnectionID,
			channel.CounterpartyChannelID,
			channel.CounterpartyPortID,
			channel.PacketOrdering,
			channel.LastInPacketSequence,
			channel.LastOutPacketSequence,
			channel.TotalTransferInCount,
			channel.TotalTransferOutCount,
			channel.TotalTransferOutSuccessCount,
			channel.TotalTransferOutSuccessRate,
			channel.LastActivityTime,
			channel.BondedTokens,
		).
		ToSql()

	if err != nil {
		return fmt.Errorf("error building channel insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := channelsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting channel into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting channel into the table: no row inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (channelsView *Channels) Update(channel *ChannelRow) error {
	sql, sqlArgs, err := channelsView.rdb.StmtBuilder.
		Update(
			"view_channels",
		).
		SetMap(map[string]interface{}{
			"connection_id":           channel.ConnectionID,
			"counterparty_channel_id": channel.CounterpartyChannelID,
			"counterparty_port_id":    channel.CounterpartyPortID,
		}).
		Where(
			"channel_id = ? AND port_id = ?", channel.ChannelID, channel.PortID,
		).
		ToSql()

	if err != nil {
		return fmt.Errorf("error building channel update sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := channelsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error updating channel into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error updating channel into the table: no row updated: %w", rdb.ErrWrite)
	}

	return nil
}

func (channelsView *Channels) Increment(channelID string, portID string, column string, increaseNumber int64) error {
	if column != "total_transfer_in_count" &&
		column != "total_transfer_out_count" &&
		column != "total_transfer_out_success_count" {
		return fmt.Errorf("error unsupported column in Increment(): %v", column)
	}

	newValue := fmt.Sprintf("%s+%v", column, increaseNumber)

	sql, sqlArgs, err := channelsView.rdb.StmtBuilder.
		Update("view_channels").
		Set(column, sq.Expr(newValue)). // e.g. Set("columnA", sq.Expr("columnA+1"))
		Where(
			"channel_id = ? AND port_id = ?", channelID, portID,
		).
		ToSql()
	if err != nil {
		return fmt.Errorf("error building channel increment sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := channelsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error incresing column: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error incresing column: no row updated: %w", rdb.ErrWrite)
	}

	return nil
}

func (channelsView *Channels) UpdateSequence(channelID string, portID string, column string, sequence uint64) error {
	if column != "last_in_packet_sequence" && column != "last_out_packet_sequence" {
		return fmt.Errorf("error unsupported column in UpdateSequence(): %v", column)
	}

	sql, sqlArgs, err := channelsView.rdb.StmtBuilder.
		Update("view_channels").
		SetMap(map[string]interface{}{
			column: sequence,
		}).
		Where(
			"channel_id = ? AND port_id = ?", channelID, portID,
		).
		ToSql()
	if err != nil {
		return fmt.Errorf("error building channel increment sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := channelsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error incresing column: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error incresing column: no row updated: %w", rdb.ErrWrite)
	}

	return nil
}

func (channelsView *Channels) UpdateTotalTransferOutSuccessRate(channelID string, portID string) error {
	sql, sqlArgs, err := channelsView.rdb.StmtBuilder.
		Select(
			"total_transfer_out_count",
			"total_transfer_out_success_count",
		).
		From("view_channels").
		Where("channel_id = ? AND port_id = ?", channelID, portID).
		ToSql()
	if err != nil {
		return fmt.Errorf("error building selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	var channel ChannelRow
	if err = channelsView.rdb.QueryRow(sql, sqlArgs...).Scan(
		&channel.TotalTransferOutCount,
		&channel.TotalTransferOutSuccessCount,
	); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return rdb.ErrNoRows
		}
		return fmt.Errorf("error scanning channel row: %v: %w", err, rdb.ErrQuery)
	}

	// Rate = totalSuccessCount / totalCount
	totalTransferOutSuccessRate := 0.0
	if channel.TotalTransferOutCount > 0 {
		totalTransferOutSuccessRate = float64(channel.TotalTransferOutSuccessCount) / float64(channel.TotalTransferOutCount)
	}

	sql, sqlArgs, err = channelsView.rdb.StmtBuilder.
		Update("view_channels").
		SetMap(map[string]interface{}{
			"total_transfer_out_success_rate": totalTransferOutSuccessRate,
		}).
		Where("channel_id = ? AND port_id = ?", channelID, portID).
		ToSql()
	if err != nil {
		return fmt.Errorf("error building channel update sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := channelsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error updating total_transfer_out_success_rate: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error updating total_transfer_out_success_rate: no row updated: %w", rdb.ErrWrite)
	}

	return nil
}

func (channelsView *Channels) UpdateLastActivityTime(channelID string, portID string, time utctime.UTCTime) error {
	sql, sqlArgs, err := channelsView.rdb.StmtBuilder.
		Update("view_channels").
		SetMap(map[string]interface{}{
			"last_activity_time": channelsView.rdb.Tton(&time),
		}).
		Where("channel_id = ? AND port_id = ?", channelID, portID).
		ToSql()
	if err != nil {
		return fmt.Errorf("error building channel update sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := channelsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error updating last_activity_time: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error updating last_activity_time: no row updated: %w", rdb.ErrWrite)
	}

	return nil
}

func (channelsView *Channels) List(order ChannelsListOrder, pagination *pagination.Pagination) ([]ChannelRow, *pagination.PaginationResult, error) {
	stmtBuilder := channelsView.rdb.StmtBuilder.Select(
		"channel_id",
		"port_id",
		"connection_id",
		"counterparty_channel_id",
		"counterparty_port_id",
		"packet_ordering",
		"last_in_packet_sequence",
		"last_out_packet_sequence",
		"total_transfer_in_count",
		"total_transfer_out_count",
		"total_transfer_out_success_count",
		"total_transfer_out_success_rate",
	).From(
		"view_channels",
	)

	if order.ChannelId == view.ORDER_DESC {
		stmtBuilder = stmtBuilder.OrderBy("channel_id DESC")
	} else {
		stmtBuilder = stmtBuilder.OrderBy("channel_id")
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		channelsView.rdb,
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building channels select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := channelsView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing channels select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	channels := make([]ChannelRow, 0)
	for rowsResult.Next() {
		var channel ChannelRow
		if err = rowsResult.Scan(
			&channel.ChannelID,
			&channel.PortID,
			&channel.ConnectionID,
			&channel.CounterpartyChannelID,
			&channel.CounterpartyPortID,
			&channel.PacketOrdering,
			&channel.LastInPacketSequence,
			&channel.LastOutPacketSequence,
			&channel.TotalTransferInCount,
			&channel.TotalTransferOutCount,
			&channel.TotalTransferOutSuccessCount,
			&channel.TotalTransferOutSuccessRate,
		); err != nil {
			if errors.Is(err, rdb.ErrNoRows) {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning channel row: %v: %w", err, rdb.ErrQuery)
		}

		channels = append(channels, channel)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return channels, paginationResult, nil
}

type ChannelsListOrder struct {
	ChannelId view.ORDER
}

type ChannelRow struct {
	ChannelID                    string                   `json:"channelId"`
	PortID                       string                   `json:"portId"`
	ConnectionID                 string                   `json:"connectionId"`
	CounterpartyChannelID        string                   `json:"counterpartyChannelId"`
	CounterpartyPortID           string                   `json:"counterpartyPortId"`
	PacketOrdering               string                   `json:"packetOrdering"`
	LastInPacketSequence         int64                    `json:"lastInPacketSequence"`
	LastOutPacketSequence        int64                    `json:"lastOutPacketSequence"`
	TotalTransferInCount         int64                    `json:"totalTransferInCount"`
	TotalTransferOutCount        int64                    `json:"totalTransferOutCount"`
	TotalTransferOutSuccessCount int64                    `json:"totalTransferOutSuccessCount"`
	TotalTransferOutSuccessRate  float64                  `json:"totalTransferOutSuccessRate"`
	BondedTokens                 []map[string]interface{} `json:"bondedTokens"`
	LastActivityTime             int64                    `json:"lastActivityTime"`
}
