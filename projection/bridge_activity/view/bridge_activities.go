package view

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"

	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/projection/bridge_activity/types"
	"github.com/crypto-com/chain-indexing/usecase/coin"
)

type BridgeActivities interface {
	FindByLinkId(linkId string) (BridgeActivityReadRow, error)
	FindBy(filter BridgeActivitiesFindByFilter) (BridgeActivityReadRow, error)
	ListByChainAddress(
		chain string,
		address string,
		order BridgeActivitiesListOrder,
		pagination *pagination_interface.Pagination,
	) ([]BridgeActivityReadRow, *pagination_interface.PaginationResult, error)
	List(
		addressFilter BridgeActivitiesListAddressFilter,
		filter BridgeActivitiesListFilter,
		order BridgeActivitiesListOrder,
		pagination *pagination_interface.Pagination,
	) ([]BridgeActivityReadRow, *pagination_interface.PaginationResult, error)
	Insert(activity *BridgeActivityInsertRow) error
	Update(activity *BridgeActivityReadRow) error
}

const TABLE_BRIDGE_ACTIVITIES = "view_bridge_activities"

var _ BridgeActivities = &BridgeActivitiesView{}

type BridgeActivitiesView struct {
	rdb *rdb.Handle
}

func NewBridgeActivitiesView(handle *rdb.Handle) BridgeActivities {
	return &BridgeActivitiesView{
		handle,
	}
}

func (view *BridgeActivitiesView) FindByLinkId(linkId string) (BridgeActivityReadRow, error) {
	return view.FindBy(BridgeActivitiesFindByFilter{
		MaybeLinkId:        &linkId,
		MaybeTransactionId: nil,
	})
}

func (view *BridgeActivitiesView) FindBy(filter BridgeActivitiesFindByFilter) (BridgeActivityReadRow, error) {
	stmtBuilder := view.rdb.StmtBuilder.Select(
		"id",
		"uuid",
		"bridge_type",
		"link_id",

		"source_block_height",
		"source_block_time",
		"source_transaction_id",
		"source_chain",
		"source_address",
		"maybe_source_smart_contract_address",

		"maybe_destination_block_height",
		"maybe_destination_block_time",
		"maybe_destination_transaction_id",
		"destination_chain",
		"destination_address",
		"maybe_destination_smart_contract_address",

		"maybe_channel_id",
		"amount",
		"denom",
		"maybe_bridge_fee_amount",
		"maybe_bridge_fee_denom",
		"status",
		"created_at",
		"updated_at",
	).From(
		TABLE_BRIDGE_ACTIVITIES,
	)

	if filter.MaybeLinkId != nil {
		stmtBuilder = stmtBuilder.Where("link_id = ?", *filter.MaybeLinkId)
	}

	if filter.MaybeTransactionId != nil {
		stmtBuilder = stmtBuilder.Where("source_transaction_id = ?", *filter.MaybeTransactionId)
	}

	sql, sqlArgs, sqlErr := stmtBuilder.ToSql()

	if sqlErr != nil {
		return BridgeActivityReadRow{}, fmt.Errorf(
			"error building bridge activity select SQL: %v, %w",
			sqlErr, rdb.ErrBuildSQLStmt,
		)
	}

	var row BridgeActivityReadRow
	var bridgeType string
	var amount string
	var maybeBridgeFeeAmount *string
	sourceBlockTimeReader := view.rdb.NtotReader()
	maybeDestinationBlockTimeReader := view.rdb.NtotReader()
	createdAtTimeReader := view.rdb.NtotReader()
	updatedAtTimeReader := view.rdb.NtotReader()

	if err := view.rdb.QueryRow(sql, sqlArgs...).Scan(
		&row.Id,
		&row.UUID,
		&bridgeType,
		&row.LinkId,

		&row.SourceBlockHeight,
		sourceBlockTimeReader.ScannableArg(),
		&row.SourceTransactionId,
		&row.SourceChain,
		&row.SourceAddress,
		&row.MaybeSourceSmartContractAddress,

		&row.MaybeDestinationBlockHeight,
		maybeDestinationBlockTimeReader.ScannableArg(),
		&row.MaybeDestinationTransactionId,
		&row.DestinationChain,
		&row.DestinationAddress,
		&row.MaybeDestinationSmartContractAddress,

		&row.MaybeChannelId,
		&amount,
		&row.MaybeDenom,
		&maybeBridgeFeeAmount,
		&row.MaybeBridgeFeeDenom,
		&row.Status,
		createdAtTimeReader.ScannableArg(),
		updatedAtTimeReader.ScannableArg(),
	); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return BridgeActivityReadRow{}, rdb.ErrNoRows
		}
		return BridgeActivityReadRow{}, fmt.Errorf("error scanning bridge activity row: %v: %w", err, rdb.ErrQuery)
	}

	if bridgeType == types.BRIDGE_TYPE_GRAVITY {
		row.BridgeType = types.BRIDGE_TYPE_GRAVITY
	} else if bridgeType == types.BRIDGE_TYPE_IBC {
		row.BridgeType = types.BRIDGE_TYPE_IBC
	} else {
		panic(fmt.Sprintf("unrecognized bridge type: %s", bridgeType))
	}

	sourceBlockTime, sourceBlockTimeErr := sourceBlockTimeReader.Parse()
	if sourceBlockTimeErr != nil {
		return BridgeActivityReadRow{}, fmt.Errorf("error parsing source block time: %v", sourceBlockTimeErr)
	}
	row.SourceBlockTime = sourceBlockTime

	maybeDestinationBlockTime, maybeDestinationBlockTimeErr := maybeDestinationBlockTimeReader.Parse()
	if maybeDestinationBlockTimeErr != nil {
		return BridgeActivityReadRow{}, fmt.Errorf("error parsing destinoation block time: %v", maybeDestinationBlockTimeErr)
	}
	row.MaybeDestinationBlockTime = maybeDestinationBlockTime

	rowAmount, rowAmountOk := coin.NewIntFromString(amount)
	if !rowAmountOk {
		return BridgeActivityReadRow{}, fmt.Errorf("error parsing amount %s", amount)
	}
	row.Amount = rowAmount

	if maybeBridgeFeeAmount != nil {
		rowBridgeFeeAmount, rowBridgeFeeAmountOk := coin.NewIntFromString(*maybeBridgeFeeAmount)
		if !rowBridgeFeeAmountOk {
			return BridgeActivityReadRow{}, fmt.Errorf("error parsing bridge fee amount %s", *maybeBridgeFeeAmount)
		}
		row.MaybeBridgeFeeAmount = &rowBridgeFeeAmount
	}

	createdAt, createdAtErr := createdAtTimeReader.Parse()
	if createdAtErr != nil {
		return BridgeActivityReadRow{}, fmt.Errorf("error parsing created at: %v", createdAtErr)
	}
	row.CreatedAt = createdAt

	updatedAt, updatedAtErr := updatedAtTimeReader.Parse()
	if updatedAtErr != nil {
		return BridgeActivityReadRow{}, fmt.Errorf("error parsing updated at: %v", updatedAtErr)
	}
	row.UpdatedAt = updatedAt

	return row, nil
}

type BridgeActivitiesFindByFilter struct {
	MaybeLinkId        *string
	MaybeTransactionId *string
}

func (view *BridgeActivitiesView) ListByChainAddress(
	chain string,
	address string,
	order BridgeActivitiesListOrder,
	pagination *pagination_interface.Pagination,
) ([]BridgeActivityReadRow, *pagination_interface.PaginationResult, error) {
	addressFilter := []BridgeActivitiesListAddressFilterCond{
		{
			Chain:   chain,
			Address: address,
		},
	}
	filter := BridgeActivitiesListFilter{}
	return view.List(addressFilter, filter, order, pagination)
}

func (view *BridgeActivitiesView) List(
	addressFilter BridgeActivitiesListAddressFilter,
	filter BridgeActivitiesListFilter,
	order BridgeActivitiesListOrder,
	pagination *pagination_interface.Pagination,
) ([]BridgeActivityReadRow, *pagination_interface.PaginationResult, error) {
	if pagination.Type() != pagination_interface.PAGINATION_OFFSET {
		panic(fmt.Sprintf("unsupported pagination type: %s", pagination.Type()))
	}

	stmtBuilder := view.rdb.StmtBuilder.Select(
		"id",
		"uuid",
		"bridge_type",
		"link_id",

		"source_block_height",
		"source_block_time",
		"source_transaction_id",
		"source_chain",
		"source_address",
		"maybe_source_smart_contract_address",

		"maybe_destination_block_height",
		"maybe_destination_block_time",
		"maybe_destination_transaction_id",
		"destination_chain",
		"destination_address",
		"maybe_destination_smart_contract_address",

		"maybe_channel_id",
		"amount",
		"denom",
		"maybe_bridge_fee_amount",
		"maybe_bridge_fee_denom",
		"status",
		"created_at",
		"updated_at",
	).From(
		TABLE_BRIDGE_ACTIVITIES,
	)

	preds := make([]string, 0)
	args := make([]interface{}, 0)
	for _, filter := range addressFilter {
		preds = append(preds, "( source_chain = ? AND source_address = ?) OR (destination_chain = ? AND destination_address = ?)")
		args = append(args, filter.Chain, filter.Address, filter.Chain, filter.Address)
	}
	pred := strings.Join(preds, " OR ")
	stmtBuilder = stmtBuilder.Where(pred, args...)

	if filter.MaybeStatus != nil {
		stmtBuilder = stmtBuilder.Where(
			"status = ?", filter.MaybeStatus,
		)
	}
	if filter.MaybeIdGt != nil {
		stmtBuilder = stmtBuilder.Where(
			"id > ?", filter.MaybeIdGt,
		)
	}
	if filter.MaybeCreatedAtLt != nil {
		stmtBuilder = stmtBuilder.Where(
			"created_at < ?", view.rdb.Tton(filter.MaybeCreatedAtLt),
		)
	}
	if filter.MaybeCreatedAtGt != nil {
		stmtBuilder = stmtBuilder.Where(
			"created_at > ?", view.rdb.Tton(filter.MaybeCreatedAtGt),
		)
	}
	if filter.MaybeUpdatedAtGt != nil {
		stmtBuilder = stmtBuilder.Where(
			"updated_at > ?", view.rdb.Tton(filter.MaybeUpdatedAtGt),
		)
	}
	if filter.MaybeSourceBlockTimeLt != nil {
		stmtBuilder = stmtBuilder.Where(
			"source_block_time < ?", view.rdb.Tton(filter.MaybeSourceBlockTimeLt),
		)
	}
	if filter.MaybeSourceBlockTimeGt != nil {
		stmtBuilder = stmtBuilder.Where(
			"source_block_time > ?", view.rdb.Tton(filter.MaybeSourceBlockTimeGt),
		)
	}
	if filter.MaybeDestinationBlockTimeGt != nil {
		stmtBuilder = stmtBuilder.Where(
			"maybe_destination_block_time > ?", view.rdb.Tton(filter.MaybeDestinationBlockTimeGt),
		)
	}

	if order.MaybeSourceBlockTime != nil {
		stmtBuilder = stmtBuilder.OrderBy(fmt.Sprintf("source_block_time %s", *order.MaybeSourceBlockTime))
	} else {
		stmtBuilder = stmtBuilder.OrderBy("source_block_time ASC")
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		view.rdb,
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, sqlErr := rDbPagination.ToStmtBuilder().ToSql()
	if sqlErr != nil {
		return nil, nil, fmt.Errorf(
			"error building bridge activity select SQL: %v, %w",
			sqlErr, rdb.ErrBuildSQLStmt,
		)
	}

	rowsResult, queryErr := view.rdb.Query(sql, sqlArgs...)
	if queryErr != nil {
		return nil, nil, fmt.Errorf("error executing bridge activities list SQL: %v: %w", queryErr, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	rows := make([]BridgeActivityReadRow, 0)
	for rowsResult.Next() {
		var row BridgeActivityReadRow
		var bridgeType string
		var amount string
		var maybeBridgeFeeAmount *string
		sourceBlockTimeReader := view.rdb.NtotReader()
		maybeDestinationBlockTimeReader := view.rdb.NtotReader()
		createdAtTimeReader := view.rdb.NtotReader()
		updatedAtTimeReader := view.rdb.NtotReader()

		if scanErr := rowsResult.Scan(
			&row.Id,
			&row.UUID,
			&bridgeType,
			&row.LinkId,

			&row.SourceBlockHeight,
			sourceBlockTimeReader.ScannableArg(),
			&row.SourceTransactionId,
			&row.SourceChain,
			&row.SourceAddress,
			&row.MaybeSourceSmartContractAddress,

			&row.MaybeDestinationBlockHeight,
			maybeDestinationBlockTimeReader.ScannableArg(),
			&row.MaybeDestinationTransactionId,
			&row.DestinationChain,
			&row.DestinationAddress,
			&row.MaybeDestinationSmartContractAddress,

			&row.MaybeChannelId,
			&amount,
			&row.MaybeDenom,
			&maybeBridgeFeeAmount,
			&row.MaybeBridgeFeeDenom,
			&row.Status,
			createdAtTimeReader.ScannableArg(),
			updatedAtTimeReader.ScannableArg(),
		); scanErr != nil {
			return nil, nil, fmt.Errorf("error scanning bridge activity row: %v: %w", scanErr, rdb.ErrQuery)
		}

		if bridgeType == types.BRIDGE_TYPE_GRAVITY {
			row.BridgeType = types.BRIDGE_TYPE_GRAVITY
		} else if bridgeType == types.BRIDGE_TYPE_IBC {
			row.BridgeType = types.BRIDGE_TYPE_IBC
		} else {
			return nil, nil, fmt.Errorf("unrecognized bridge type: %s", bridgeType)
		}

		sourceBlockTime, sourceBlockTimeErr := sourceBlockTimeReader.Parse()
		if sourceBlockTimeErr != nil {
			return nil, nil, fmt.Errorf("error parsing source block time: %v", sourceBlockTimeErr)
		}
		row.SourceBlockTime = sourceBlockTime

		maybeDestinationBlockTime, maybeDestinationBlockTimeErr := maybeDestinationBlockTimeReader.Parse()
		if maybeDestinationBlockTimeErr != nil {
			return nil, nil, fmt.Errorf("error parsing destinoation block time: %v", maybeDestinationBlockTimeErr)
		}
		row.MaybeDestinationBlockTime = maybeDestinationBlockTime

		rowAmount, rowAmountOk := coin.NewIntFromString(amount)
		if !rowAmountOk {
			return nil, nil, fmt.Errorf("error parsing amount %s", amount)
		}
		row.Amount = rowAmount

		if maybeBridgeFeeAmount != nil {
			rowBridgeFeeAmount, rowBridgeFeeAmountOk := coin.NewIntFromString(*maybeBridgeFeeAmount)
			if !rowBridgeFeeAmountOk {
				return nil, nil, fmt.Errorf("error parsing bridge fee amount %s", *maybeBridgeFeeAmount)
			}
			row.MaybeBridgeFeeAmount = &rowBridgeFeeAmount
		}

		createdAt, createdAtErr := createdAtTimeReader.Parse()
		if createdAtErr != nil {
			return nil, nil, fmt.Errorf("error parsing created at: %v", createdAtErr)
		}
		row.CreatedAt = createdAt

		updatedAt, updatedAtErr := updatedAtTimeReader.Parse()
		if updatedAtErr != nil {
			return nil, nil, fmt.Errorf("error parsing updated at: %v", updatedAtErr)
		}
		row.UpdatedAt = updatedAt

		rows = append(rows, row)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return rows, paginationResult, nil
}

type BridgeActivitiesListFilter struct {
	MaybeStatus                 *types.Status
	MaybeIdGt                   *string
	MaybeCreatedAtLt            *utctime.UTCTime
	MaybeCreatedAtGt            *utctime.UTCTime
	MaybeUpdatedAtGt            *utctime.UTCTime
	MaybeSourceBlockTimeLt      *utctime.UTCTime
	MaybeSourceBlockTimeGt      *utctime.UTCTime
	MaybeDestinationBlockTimeGt *utctime.UTCTime
}

type BridgeActivitiesListAddressFilter = []BridgeActivitiesListAddressFilterCond

type BridgeActivitiesListAddressFilterCond struct {
	Chain   string
	Address string
}

type BridgeActivitiesListOrder struct {
	MaybeSourceBlockTime *view.ORDER
}

func (view *BridgeActivitiesView) Insert(activity *BridgeActivityInsertRow) error {
	timeNow := utctime.Now()
	var maybeBridgeFeeAmount *string
	if activity.MaybeBridgeFeeAmount != nil {
		maybeBridgeFeeAmount = primptr.String(activity.MaybeBridgeFeeAmount.String())
	}
	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Insert(TABLE_BRIDGE_ACTIVITIES).
		Columns(
			"uuid",
			"bridge_type",
			"link_id",

			"source_block_height",
			"source_block_time",
			"source_transaction_id",
			"source_chain",
			"source_address",
			"maybe_source_smart_contract_address",

			"maybe_destination_block_height",
			"maybe_destination_block_time",
			"maybe_destination_transaction_id",
			"destination_chain",
			"destination_address",
			"maybe_destination_smart_contract_address",

			"maybe_channel_id",
			"amount",
			"denom",
			"maybe_bridge_fee_amount",
			"maybe_bridge_fee_denom",
			"status",
			"created_at",
			"updated_at",
		).
		Values(
			uuid.New().String(),
			activity.BridgeType,
			activity.LinkId,

			activity.SourceBlockHeight,
			view.rdb.Tton(activity.SourceBlockTime),
			activity.SourceTransactionId,
			activity.SourceChain,
			activity.SourceAddress,
			activity.MaybeSourceSmartContractAddress,

			activity.MaybeDestinationBlockHeight,
			view.rdb.Tton(activity.MaybeDestinationBlockTime),
			activity.MaybeDestinationTransactionId,
			activity.DestinationChain,
			activity.DestinationAddress,
			activity.MaybeDestinationSmartContractAddress,

			activity.MaybeChannelId,
			activity.Amount.String(),
			activity.MaybeDenom,
			maybeBridgeFeeAmount,
			activity.MaybeBridgeFeeDenom,
			activity.Status,
			view.rdb.Tton(&timeNow),
			view.rdb.Tton(&timeNow),
		).
		ToSql()

	if err != nil {
		return fmt.Errorf("error building activiy row insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := view.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting activity row into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting activity row into the table: no row inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (view *BridgeActivitiesView) Update(activity *BridgeActivityReadRow) error {
	timeNow := utctime.Now()
	var maybeBridgeFeeAmount *string
	if activity.MaybeBridgeFeeAmount != nil {
		maybeBridgeFeeAmount = primptr.String(activity.MaybeBridgeFeeAmount.String())
	}
	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Update(TABLE_BRIDGE_ACTIVITIES).
		SetMap(map[string]interface{}{
			//"uuid":        activity.UUID,
			"bridge_type": activity.BridgeType,
			"link_id":     activity.LinkId,

			"source_block_height":                 activity.SourceBlockHeight,
			"source_block_time":                   view.rdb.Tton(activity.SourceBlockTime),
			"source_transaction_id":               activity.SourceTransactionId,
			"source_chain":                        activity.SourceChain,
			"source_address":                      activity.SourceAddress,
			"maybe_source_smart_contract_address": activity.MaybeSourceSmartContractAddress,

			"maybe_destination_block_height":           activity.MaybeDestinationBlockHeight,
			"maybe_destination_block_time":             view.rdb.Tton(activity.MaybeDestinationBlockTime),
			"maybe_destination_transaction_id":         activity.MaybeDestinationTransactionId,
			"destination_chain":                        activity.DestinationChain,
			"destination_address":                      activity.DestinationAddress,
			"maybe_destination_smart_contract_address": activity.MaybeDestinationSmartContractAddress,

			"maybe_channel_id":        activity.MaybeChannelId,
			"amount":                  activity.Amount.String(),
			"denom":                   activity.MaybeDenom,
			"maybe_bridge_fee_amount": maybeBridgeFeeAmount,
			"maybe_bridge_fee_denom":  activity.MaybeBridgeFeeDenom,
			"status":                  activity.Status,
			//"created_at": view.rdb.Tton(activity.CreatedAt),
			"updated_at": view.rdb.Tton(&timeNow),
		}).
		Where(
			"id = ?", activity.Id,
		).ToSql()
	if err != nil {
		return fmt.Errorf("error building activiy row update sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := view.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error updating activity row into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error updating activity row into the table: no row updated: %w", rdb.ErrWrite)
	}

	return nil
}

type BridgeActivityReadRow struct {
	BridgeActivityInsertRow

	Id        int64            `json:"id"`
	UUID      string           `json:"uuid"`
	CreatedAt *utctime.UTCTime `json:"createdAt"`
	UpdatedAt *utctime.UTCTime `json:"updatedAt"`
}

type BridgeActivityInsertRow struct {
	BridgeType                           types.BridgeType `json:"bridgeType"`
	SourceBlockHeight                    int64            `json:"sourceBlockHeight"`
	SourceBlockTime                      *utctime.UTCTime `json:"sourceBlockTime"`
	SourceTransactionId                  string           `json:"sourceTransactionId"`
	SourceChain                          string           `json:"sourceChain"`
	SourceAddress                        string           `json:"sourceAddress"`
	MaybeSourceSmartContractAddress      *string          `json:"sourceSmartContractAddress"`
	MaybeDestinationBlockHeight          *int64           `json:"destinationBlockHeight"`
	MaybeDestinationBlockTime            *utctime.UTCTime `json:"destinationBlockTime"`
	MaybeDestinationTransactionId        *string          `json:"destinationTransactionId"`
	DestinationChain                     string           `json:"destinationChain"`
	DestinationAddress                   string           `json:"destinationAddress"`
	MaybeDestinationSmartContractAddress *string          `json:"destinationSmartContractAddress"`
	MaybeChannelId                       *string          `json:"channelId"`
	LinkId                               string           `json:"-"`
	Amount                               coin.Int         `json:"amount"`
	MaybeDenom                           *string          `json:"denom"`
	MaybeBridgeFeeAmount                 *coin.Int        `json:"bridgeFeeAmount"`
	MaybeBridgeFeeDenom                  *string          `json:"bridgeFeeDenom"`
	Status                               types.Status     `json:"status"`
}
