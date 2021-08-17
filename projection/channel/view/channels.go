package view

import (
	"errors"
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"

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

func (channelsView *Channels) Upsert(channel *ChannelRow) error {
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
			"success_packet_count",
			"failure_packet_count",
		).
		Values(
			channel.ChannelID,
			channel.PortID,
			channel.ConnectionID,
			channel.CounterpartyChannelID,
			channel.CounterpartyPortID,
			channel.SuccessPacketCount,
			channel.FailurePacketCount,
		).
		Suffix(`ON CONFLICT (channel_id, port_id) DO UPDATE SET
			connection_id = EXCLUDED.connection_id,
			counterparty_channel_id = EXCLUDED.counterparty_channel_id,
			counterparty_port_id = EXCLUDED.counterparty_port_id,
			success_packet_count = EXCLUDED.success_packet_count,
			failure_packet_count = EXCLUDED.failure_packet_count
		`).
		ToSql()

	if err != nil {
		return fmt.Errorf("error building channel insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := channelsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting channel into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting channel into the table: no rows inserted: %w", rdb.ErrWrite)
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
		"success_packet_count",
		"failure_packet_count",
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
			&channel.SuccessPacketCount,
			&channel.FailurePacketCount,
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

func (channelsView *Channels) SuccessPacketCountIncrement(channelId string, portId string) error {
	sql, sqlArgs, err := channelsView.rdb.StmtBuilder.
		Insert("view_channels AS ORIGINAL").
		Columns("channel_id", "port_id", "success_packet_count", "failure_packet_count").
		Values(channelId, portId, 1, 0).
		Suffix("ON CONFLICT (channel_id, port_id) DO UPDATE SET success_packet_count = ORIGINAL.success_packet_count + EXCLUDED.success_packet_count").
		ToSql()
	if err != nil {
		return fmt.Errorf("error building channel insert sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := channelsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting success_packet_count: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting success_packet_count: no row inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (channelsView *Channels) FailurePacketCountIncrement(channelId string, portId string) error {
	sql, sqlArgs, err := channelsView.rdb.StmtBuilder.
		Insert("view_channels AS ORIGINAL").
		Columns("channel_id", "port_id", "success_packet_count", "failure_packet_count").
		Values(channelId, portId, 0, 1).
		Suffix("ON CONFLICT (channel_id, port_id) DO UPDATE SET failure_packet_count = ORIGINAL.failure_packet_count + EXCLUDED.failure_packet_count").
		ToSql()
	if err != nil {
		return fmt.Errorf("error building channel insert sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := channelsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting failure_packet_count: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting failure_packet_count: no row inserted: %w", rdb.ErrWrite)
	}

	return nil
}

type ChannelsListOrder struct {
	ChannelId view.ORDER
}

type ChannelRow struct {
	ChannelID             string `json:"channelId"`
	PortID                string `json:"portId"`
	ConnectionID          string `json:"connectionId"`
	CounterpartyChannelID string `json:"counterpartyChannelId"`
	CounterpartyPortID    string `json:"counterpartyPortId"`
	SuccessPacketCount    int    `json:"successPacketCount"`
	FailurePacketCount    int    `json:"failurePacketCount"`
}
