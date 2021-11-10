package view

import (
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	jsoniter "github.com/json-iterator/go"

	"github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/projection/ibc_channel/types"
)

type IBCChannels interface {
	Insert(*IBCChannelRow) error
	UpdateFactualColumns(*IBCChannelRow) error
	Increment(string, string, int64) error
	UpdateSequence(string, string, uint64) error
	UpdateTotalRelayOutSuccessRate(string) error
	UpdateLastActivityTimeAndHeight(string, utctime.UTCTime, int64) error
	UpdateStatus(string, types.Status) error
	UpdateBondedTokens(string, *BondedTokens) error
	FindBondedTokensBy(string) (*BondedTokens, error)
	FindBy(string) (*IBCChannelRow, error)
	List(IBCChannelsListOrder, IBCChannelsListFilter, *pagination.Pagination) ([]IBCChannelRow, *pagination.PaginationResult, error)
	ListChannelsGroupByChainId(IBCChannelsListOrder, IBCChannelsListFilter, *pagination.Pagination) ([]ChainChannels, *pagination.PaginationResult, error)
}

// Channels projection view implemented by relational database
type IBCChannelsView struct {
	rdb *rdb.Handle
}

func NewIBCChannelsView(handle *rdb.Handle) IBCChannels {
	return &IBCChannelsView{
		handle,
	}
}

func (ibcChannelsView *IBCChannelsView) Insert(channel *IBCChannelRow) error {
	bondedTokensJSON, err := jsoniter.MarshalToString(channel.BondedTokens)
	if err != nil {
		return fmt.Errorf("error JSON marshalling channel bonded_tokens for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	sql, sqlArgs, err := ibcChannelsView.rdb.StmtBuilder.
		Insert(
			"view_ibc_channels",
		).
		Columns(
			"channel_id",
			"port_id",
			"connection_id",
			"counterparty_channel_id",
			"counterparty_port_id",
			"counterparty_chain_id",
			"status",
			"packet_ordering",
			"last_in_packet_sequence",
			"last_out_packet_sequence",
			"total_relay_in_count",
			"total_relay_out_count",
			"total_relay_out_success_count",
			"total_relay_out_success_rate",
			"created_at_block_time",
			"created_at_block_height",
			"verified",
			"description",
			"last_activity_block_time",
			"last_activity_block_height",
			"bonded_tokens",
		).
		Values(
			channel.ChannelID,
			channel.PortID,
			channel.ConnectionID,
			channel.CounterpartyChannelID,
			channel.CounterpartyPortID,
			channel.CounterpartyChainID,
			channel.Status,
			channel.PacketOrdering,
			channel.LastInPacketSequence,
			channel.LastOutPacketSequence,
			channel.TotalRelayInCount,
			channel.TotalRelayOutCount,
			channel.TotalRelayOutSuccessCount,
			channel.TotalRelayOutSuccessRate,
			ibcChannelsView.rdb.Tton(&channel.CreatedAtBlockTime),
			channel.CreatedAtBlockHeight,
			channel.Verified,
			channel.Description,
			ibcChannelsView.rdb.Tton(&channel.LastActivityBlockTime),
			channel.LastActivityBlockHeight,
			bondedTokensJSON,
		).
		ToSql()

	if err != nil {
		return fmt.Errorf("error building channel insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := ibcChannelsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting channel into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting channel into the table: no row inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (ibcChannelsView *IBCChannelsView) UpdateFactualColumns(channel *IBCChannelRow) error {
	sql, sqlArgs, err := ibcChannelsView.rdb.StmtBuilder.
		Update(
			"view_ibc_channels",
		).
		SetMap(map[string]interface{}{
			"connection_id":           channel.ConnectionID,
			"counterparty_channel_id": channel.CounterpartyChannelID,
			"counterparty_port_id":    channel.CounterpartyPortID,
			"counterparty_chain_id":   channel.CounterpartyChainID,
			"created_at_block_time":   ibcChannelsView.rdb.Tton(&channel.CreatedAtBlockTime),
			"created_at_block_height": channel.CreatedAtBlockHeight,
		}).
		Where(
			"channel_id = ?", channel.ChannelID,
		).
		ToSql()

	if err != nil {
		return fmt.Errorf("error building channel update sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := ibcChannelsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error updating channel into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error updating channel into the table: no row updated: %w", rdb.ErrWrite)
	}

	return nil
}

func (ibcChannelsView *IBCChannelsView) Increment(channelID string, column string, increaseNumber int64) error {
	if column != "total_relay_in_count" &&
		column != "total_relay_out_count" &&
		column != "total_relay_out_success_count" {
		return fmt.Errorf("error unsupported column in Increment(): %v", column)
	}

	newValue := fmt.Sprintf("%s+%v", column, increaseNumber)

	sql, sqlArgs, err := ibcChannelsView.rdb.StmtBuilder.
		Update("view_ibc_channels").
		Set(column, sq.Expr(newValue)). // e.g. Set("columnA", sq.Expr("columnA+1"))
		Where("channel_id = ?", channelID).
		ToSql()
	if err != nil {
		return fmt.Errorf("error building channel increment sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := ibcChannelsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error incresing column: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error incresing column: no row updated: %w", rdb.ErrWrite)
	}

	return nil
}

func (ibcChannelsView *IBCChannelsView) UpdateSequence(channelID string, column string, sequence uint64) error {
	if column != "last_in_packet_sequence" && column != "last_out_packet_sequence" {
		return fmt.Errorf("error unsupported column in UpdateSequence(): %v", column)
	}

	sql, sqlArgs, err := ibcChannelsView.rdb.StmtBuilder.
		Update("view_ibc_channels").
		SetMap(map[string]interface{}{
			column: sequence,
		}).
		Where("channel_id = ?", channelID).
		ToSql()
	if err != nil {
		return fmt.Errorf("error building channel increment sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := ibcChannelsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error incresing column: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error incresing column: no row updated: %w", rdb.ErrWrite)
	}

	return nil
}

func (ibcChannelsView *IBCChannelsView) UpdateTotalRelayOutSuccessRate(channelID string) error {
	sql, sqlArgs, err := ibcChannelsView.rdb.StmtBuilder.
		Select(
			"total_relay_out_count",
			"total_relay_out_success_count",
		).
		From("view_ibc_channels").
		Where("channel_id = ?", channelID).
		ToSql()
	if err != nil {
		return fmt.Errorf("error building selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	var channel IBCChannelRow
	if err = ibcChannelsView.rdb.QueryRow(sql, sqlArgs...).Scan(
		&channel.TotalRelayOutCount,
		&channel.TotalRelayOutSuccessCount,
	); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return rdb.ErrNoRows
		}
		return fmt.Errorf("error scanning channel row: %v: %w", err, rdb.ErrQuery)
	}

	// Rate = totalSuccessCount / totalCount
	totalRelayOutSuccessRate := 0.0
	if channel.TotalRelayOutCount > 0 {
		totalRelayOutSuccessRate = float64(channel.TotalRelayOutSuccessCount) / float64(channel.TotalRelayOutCount)
	}

	sql, sqlArgs, err = ibcChannelsView.rdb.StmtBuilder.
		Update("view_ibc_channels").
		SetMap(map[string]interface{}{
			"total_relay_out_success_rate": totalRelayOutSuccessRate,
		}).
		Where("channel_id = ?", channelID).
		ToSql()
	if err != nil {
		return fmt.Errorf("error building channel update sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := ibcChannelsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error updating total_relay_out_success_rate: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error updating total_relay_out_success_rate: no row updated: %w", rdb.ErrWrite)
	}

	return nil
}

func (ibcChannelsView *IBCChannelsView) UpdateLastActivityTimeAndHeight(channelID string, time utctime.UTCTime, height int64) error {
	sql, sqlArgs, err := ibcChannelsView.rdb.StmtBuilder.
		Update("view_ibc_channels").
		SetMap(map[string]interface{}{
			"last_activity_block_time":   ibcChannelsView.rdb.Tton(&time),
			"last_activity_block_height": height,
		}).
		Where("channel_id = ?", channelID).
		ToSql()
	if err != nil {
		return fmt.Errorf("error building channel update sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := ibcChannelsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error updating last_activity time and height: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error updating last_activity time and height: no row updated: %w", rdb.ErrWrite)
	}

	return nil
}

func (ibcChannelsView *IBCChannelsView) UpdateStatus(channelID string, status types.Status) error {
	sql, sqlArgs, err := ibcChannelsView.rdb.StmtBuilder.
		Update("view_ibc_channels").
		SetMap(map[string]interface{}{
			"status": status,
		}).
		Where("channel_id = ?", channelID).
		ToSql()
	if err != nil {
		return fmt.Errorf("error building channel.Status update sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := ibcChannelsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error updating channel.Status: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error updating channel.Status: no row updated: %w", rdb.ErrWrite)
	}

	return nil
}

func (ibcChannelsView *IBCChannelsView) UpdateBondedTokens(channelID string, bondedTokens *BondedTokens) error {
	bondedTokensJSON, err := jsoniter.MarshalToString(*bondedTokens)
	if err != nil {
		return fmt.Errorf("error JSON marshalling channel bonded_tokens for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	sql, sqlArgs, err := ibcChannelsView.rdb.StmtBuilder.
		Update("view_ibc_channels").
		SetMap(map[string]interface{}{
			"bonded_tokens": bondedTokensJSON,
		}).
		Where("channel_id = ?", channelID).
		ToSql()
	if err != nil {
		return fmt.Errorf("error building channel.bonded_tokens update sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := ibcChannelsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error updating channel.bonded_tokens: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error updating channel.bonded_tokens: no row updated: %w", rdb.ErrWrite)
	}

	return nil
}

func (ibcChannelsView *IBCChannelsView) FindBondedTokensBy(channelID string) (*BondedTokens, error) {
	sql, sqlArgs, err := ibcChannelsView.rdb.StmtBuilder.
		Select("bonded_tokens").
		From("view_ibc_channels").
		Where("channel_id = ?", channelID).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building select bonded_tokens sql: %v: %w", err, rdb.ErrPrepare)
	}

	var bondedTokensJSON string
	if err = ibcChannelsView.rdb.QueryRow(sql, sqlArgs...).Scan(
		&bondedTokensJSON,
	); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return nil, rdb.ErrNoRows
		}
		return nil, fmt.Errorf("error scanning channel bonded_tokens: %v: %w", err, rdb.ErrQuery)
	}

	var bondedTokens BondedTokens
	if err = jsoniter.UnmarshalFromString(bondedTokensJSON, &bondedTokens); err != nil {
		return nil, fmt.Errorf("error unmarshalling channel bonded_tokens JSON: %v: %w", err, rdb.ErrQuery)
	}

	return &bondedTokens, nil
}

func (ibcChannelsView *IBCChannelsView) FindBy(channelID string) (*IBCChannelRow, error) {
	sql, sqlArgs, err := ibcChannelsView.rdb.StmtBuilder.
		Select(
			"channel_id",
			"port_id",
			"connection_id",
			"counterparty_channel_id",
			"counterparty_port_id",
			"counterparty_chain_id",
			"status",
			"packet_ordering",
			"last_in_packet_sequence",
			"last_out_packet_sequence",
			"total_relay_in_count",
			"total_relay_out_count",
			"total_relay_out_success_count",
			"total_relay_out_success_rate",
			"created_at_block_time",
			"created_at_block_height",
			"verified",
			"description",
			"last_activity_block_time",
			"last_activity_block_height",
			"bonded_tokens",
		).
		From("view_ibc_channels").
		Where("channel_id = ?", channelID).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building select channel sql: %v: %w", err, rdb.ErrPrepare)
	}

	var channel IBCChannelRow
	var bondedTokensJSON string
	lastActivityTimeReader := ibcChannelsView.rdb.NtotReader()
	createdAtTimeReader := ibcChannelsView.rdb.NtotReader()
	if err = ibcChannelsView.rdb.QueryRow(sql, sqlArgs...).Scan(
		&channel.ChannelID,
		&channel.PortID,
		&channel.ConnectionID,
		&channel.CounterpartyChannelID,
		&channel.CounterpartyPortID,
		&channel.CounterpartyChainID,
		&channel.Status,
		&channel.PacketOrdering,
		&channel.LastInPacketSequence,
		&channel.LastOutPacketSequence,
		&channel.TotalRelayInCount,
		&channel.TotalRelayOutCount,
		&channel.TotalRelayOutSuccessCount,
		&channel.TotalRelayOutSuccessRate,
		createdAtTimeReader.ScannableArg(),
		&channel.CreatedAtBlockHeight,
		&channel.Verified,
		&channel.Description,
		lastActivityTimeReader.ScannableArg(),
		&channel.LastActivityBlockHeight,
		&bondedTokensJSON,
	); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return nil, rdb.ErrNoRows
		}
		return nil, fmt.Errorf("error scanning channel: %v: %w", err, rdb.ErrQuery)
	}

	createdAtBlockTime, parseErr := createdAtTimeReader.Parse()
	if parseErr != nil {
		return nil, fmt.Errorf("error parsing createdAtBlockTime: %v: %w", parseErr, rdb.ErrQuery)
	}
	channel.CreatedAtBlockTime = *createdAtBlockTime

	lastActivityBlockTime, parseErr := lastActivityTimeReader.Parse()
	if parseErr != nil {
		return nil, fmt.Errorf("error parsing lastActivityBlockTime: %v: %w", parseErr, rdb.ErrQuery)
	}
	channel.LastActivityBlockTime = *lastActivityBlockTime

	if err = jsoniter.UnmarshalFromString(bondedTokensJSON, &channel.BondedTokens); err != nil {
		return nil, fmt.Errorf("error unmarshalling channel bonded_tokens JSON: %v: %w", err, rdb.ErrQuery)
	}

	return &channel, nil
}

func (ibcChannelsView *IBCChannelsView) List(
	order IBCChannelsListOrder,
	filter IBCChannelsListFilter,
	pagination *pagination.Pagination,
) (
	[]IBCChannelRow,
	*pagination.PaginationResult,
	error,
) {

	stmtBuilder := ibcChannelsView.rdb.StmtBuilder.Select(
		"channel_id",
		"port_id",
		"connection_id",
		"counterparty_channel_id",
		"counterparty_port_id",
		"counterparty_chain_id",
		"status",
		"packet_ordering",
		"last_in_packet_sequence",
		"last_out_packet_sequence",
		"total_relay_in_count",
		"total_relay_out_count",
		"total_relay_out_success_count",
		"total_relay_out_success_rate",
		"created_at_block_time",
		"created_at_block_height",
		"verified",
		"description",
		"last_activity_block_time",
		"last_activity_block_height",
		"bonded_tokens",
	).From(
		"view_ibc_channels",
	)

	if filter.MaybeStatus != nil {
		stmtBuilder = stmtBuilder.Where("status = ?", *filter.MaybeStatus)
	}

	// MaybeLastActivityBlockTime has a higher priority than MaybeCreatedAtBlockTime
	if order.MaybeLastActivityBlockTime != nil {
		if *order.MaybeLastActivityBlockTime == view.ORDER_DESC {
			stmtBuilder = stmtBuilder.OrderBy("last_activity_block_time DESC")
		} else {
			stmtBuilder = stmtBuilder.OrderBy("last_activity_block_time")
		}
	} else if order.MaybeCreatedAtBlockTime != nil {
		if *order.MaybeCreatedAtBlockTime == view.ORDER_DESC {
			stmtBuilder = stmtBuilder.OrderBy("created_at_block_time DESC")
		} else {
			stmtBuilder = stmtBuilder.OrderBy("created_at_block_time")
		}
	} else {
		// By default, sort by last_activity_block_time in descending order
		stmtBuilder = stmtBuilder.OrderBy("last_activity_block_time DESC")
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		ibcChannelsView.rdb,
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building channels select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := ibcChannelsView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing channels select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	channels := make([]IBCChannelRow, 0)
	for rowsResult.Next() {
		var channel IBCChannelRow
		var bondedTokensJSON string
		lastActivityTimeReader := ibcChannelsView.rdb.NtotReader()
		createdAtTimeReader := ibcChannelsView.rdb.NtotReader()
		if err = rowsResult.Scan(
			&channel.ChannelID,
			&channel.PortID,
			&channel.ConnectionID,
			&channel.CounterpartyChannelID,
			&channel.CounterpartyPortID,
			&channel.CounterpartyChainID,
			&channel.Status,
			&channel.PacketOrdering,
			&channel.LastInPacketSequence,
			&channel.LastOutPacketSequence,
			&channel.TotalRelayInCount,
			&channel.TotalRelayOutCount,
			&channel.TotalRelayOutSuccessCount,
			&channel.TotalRelayOutSuccessRate,
			createdAtTimeReader.ScannableArg(),
			&channel.CreatedAtBlockHeight,
			&channel.Verified,
			&channel.Description,
			lastActivityTimeReader.ScannableArg(),
			&channel.LastActivityBlockHeight,
			&bondedTokensJSON,
		); err != nil {
			if errors.Is(err, rdb.ErrNoRows) {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning channel row: %v: %w", err, rdb.ErrQuery)
		}

		createdAtBlockTime, parseErr := createdAtTimeReader.Parse()
		if parseErr != nil {
			return nil, nil, fmt.Errorf("error parsing createdAtBlockTime: %v: %w", parseErr, rdb.ErrQuery)
		}
		channel.CreatedAtBlockTime = *createdAtBlockTime

		lastActivityBlockTime, parseErr := lastActivityTimeReader.Parse()
		if parseErr != nil {
			return nil, nil, fmt.Errorf("error parsing lastActivityBlockTime: %v: %w", parseErr, rdb.ErrQuery)
		}
		channel.LastActivityBlockTime = *lastActivityBlockTime

		if err = jsoniter.UnmarshalFromString(bondedTokensJSON, &channel.BondedTokens); err != nil {
			return nil, nil, fmt.Errorf("error unmarshalling channel bonded_tokens JSON: %v: %w", err, rdb.ErrQuery)
		}

		channels = append(channels, channel)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return channels, paginationResult, nil
}

func (ibcChannelsView *IBCChannelsView) ListChannelsGroupByChainId(
	order IBCChannelsListOrder,
	filter IBCChannelsListFilter,
	pagination *pagination.Pagination,
) (
	[]ChainChannels,
	*pagination.PaginationResult,
	error,
) {

	stmtBuilder := ibcChannelsView.rdb.StmtBuilder.Select(
		"channel_id",
		"port_id",
		"connection_id",
		"counterparty_channel_id",
		"counterparty_port_id",
		"counterparty_chain_id",
		"status",
		"packet_ordering",
		"last_in_packet_sequence",
		"last_out_packet_sequence",
		"total_relay_in_count",
		"total_relay_out_count",
		"total_relay_out_success_count",
		"total_relay_out_success_rate",
		"created_at_block_time",
		"created_at_block_height",
		"verified",
		"description",
		"last_activity_block_time",
		"last_activity_block_height",
		"bonded_tokens",
	).From(
		"view_ibc_channels",
	)

	if filter.MaybeStatus != nil {
		stmtBuilder = stmtBuilder.Where("status = ?", *filter.MaybeStatus)
	}

	// MaybeLastActivityBlockTime has a higher priority than MaybeCreatedAtBlockTime
	if order.MaybeLastActivityBlockTime != nil {
		if *order.MaybeLastActivityBlockTime == view.ORDER_DESC {
			stmtBuilder = stmtBuilder.OrderBy("last_activity_block_time DESC")
		} else {
			stmtBuilder = stmtBuilder.OrderBy("last_activity_block_time")
		}
	} else if order.MaybeCreatedAtBlockTime != nil {
		if *order.MaybeCreatedAtBlockTime == view.ORDER_DESC {
			stmtBuilder = stmtBuilder.OrderBy("created_at_block_time DESC")
		} else {
			stmtBuilder = stmtBuilder.OrderBy("created_at_block_time")
		}
	} else {
		// By default, sort by last_activity_block_time in descending order
		stmtBuilder = stmtBuilder.OrderBy("last_activity_block_time DESC")
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		ibcChannelsView.rdb,
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building channels select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := ibcChannelsView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing channels select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	chainChannelsMap := make(map[IBCChainID][]IBCChannelRow)

	for rowsResult.Next() {
		var channel IBCChannelRow
		var bondedTokensJSON string
		lastActivityTimeReader := ibcChannelsView.rdb.NtotReader()
		createdAtTimeReader := ibcChannelsView.rdb.NtotReader()
		if err = rowsResult.Scan(
			&channel.ChannelID,
			&channel.PortID,
			&channel.ConnectionID,
			&channel.CounterpartyChannelID,
			&channel.CounterpartyPortID,
			&channel.CounterpartyChainID,
			&channel.Status,
			&channel.PacketOrdering,
			&channel.LastInPacketSequence,
			&channel.LastOutPacketSequence,
			&channel.TotalRelayInCount,
			&channel.TotalRelayOutCount,
			&channel.TotalRelayOutSuccessCount,
			&channel.TotalRelayOutSuccessRate,
			createdAtTimeReader.ScannableArg(),
			&channel.CreatedAtBlockHeight,
			&channel.Verified,
			&channel.Description,
			lastActivityTimeReader.ScannableArg(),
			&channel.LastActivityBlockHeight,
			&bondedTokensJSON,
		); err != nil {
			if errors.Is(err, rdb.ErrNoRows) {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning channel row: %v: %w", err, rdb.ErrQuery)
		}

		createdAtBlockTime, parseErr := createdAtTimeReader.Parse()
		if parseErr != nil {
			return nil, nil, fmt.Errorf("error parsing createdAtBlockTime: %v: %w", parseErr, rdb.ErrQuery)
		}
		channel.CreatedAtBlockTime = *createdAtBlockTime

		lastActivityBlockTime, parseErr := lastActivityTimeReader.Parse()
		if parseErr != nil {
			return nil, nil, fmt.Errorf("error parsing lastActivityBlockTime: %v: %w", parseErr, rdb.ErrQuery)
		}
		channel.LastActivityBlockTime = *lastActivityBlockTime

		if err = jsoniter.UnmarshalFromString(bondedTokensJSON, &channel.BondedTokens); err != nil {
			return nil, nil, fmt.Errorf("error unmarshalling channel bonded_tokens JSON: %v: %w", err, rdb.ErrQuery)
		}

		// Append channel to chainChannelsMap[chainID]
		chainID := IBCChainID(channel.CounterpartyChainID)
		if channels, ok := chainChannelsMap[chainID]; ok {
			channels = append(channels, channel)
			chainChannelsMap[chainID] = channels
		} else {
			newChannels := make([]IBCChannelRow, 0)
			newChannels = append(newChannels, channel)
			chainChannelsMap[chainID] = newChannels
		}
	}

	chainChannelsList := convertChainChannelsMapToList(chainChannelsMap)

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return chainChannelsList, paginationResult, nil
}

func convertChainChannelsMapToList(chainChannelsMap map[IBCChainID][]IBCChannelRow) []ChainChannels {
	var chainChannelsList []ChainChannels
	for chainID, channels := range chainChannelsMap {
		chainChannels := ChainChannels{
			ChainID:  string(chainID),
			Channels: channels,
		}
		chainChannelsList = append(chainChannelsList, chainChannels)
	}
	return chainChannelsList
}

type IBCChainID string

type IBCChannelsListFilter struct {
	MaybeStatus *types.Status
}

type IBCChannelsListOrder struct {
	MaybeCreatedAtBlockTime    *view.ORDER
	MaybeLastActivityBlockTime *view.ORDER
}

type ChainChannels struct {
	ChainID  string          `json:"chainId"`
	Channels []IBCChannelRow `json:"channels"`
}

type IBCChannelRow struct {
	ChannelID                 string          `json:"channelId"`
	PortID                    string          `json:"portId"`
	ConnectionID              string          `json:"connectionId"`
	CounterpartyChannelID     string          `json:"counterpartyChannelId"`
	CounterpartyPortID        string          `json:"counterpartyPortId"`
	CounterpartyChainID       string          `json:"counterpartyChainId"`
	Status                    string          `json:"status"`
	PacketOrdering            string          `json:"packetOrdering"`
	LastInPacketSequence      int64           `json:"lastInPacketSequence"`
	LastOutPacketSequence     int64           `json:"lastOutPacketSequence"`
	TotalRelayInCount         int64           `json:"totalRelayInCount"`
	TotalRelayOutCount        int64           `json:"totalRelayOutCount"`
	TotalRelayOutSuccessCount int64           `json:"totalRelayOutSuccessCount"`
	TotalRelayOutSuccessRate  float64         `json:"totalRelayOutSuccessRate"`
	CreatedAtBlockTime        utctime.UTCTime `json:"createdAtBlockTime"`
	CreatedAtBlockHeight      int64           `json:"createdAtBlockHeight"`
	Verified                  bool            `json:"verified"`
	Description               string          `json:"description"`
	LastActivityBlockTime     utctime.UTCTime `json:"lastActivityBlockTime"`
	LastActivityBlockHeight   int64           `json:"lastActivityBlockHeight"`
	BondedTokens              BondedTokens    `json:"bondedTokens"`
}
