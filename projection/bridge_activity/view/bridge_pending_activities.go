package view

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"

	"github.com/crypto-com/chain-indexing/internal/primptr"

	"github.com/crypto-com/chain-indexing/projection/bridge_activity/types"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/internal/utctime"
	"github.com/crypto-com/chain-indexing/usecase/coin"
)

const TABLE_BRIDGE_PENDING_ACTIVITIES = "view_bridge_pending_activities"

type BridgePendingActivities struct {
	rdb *rdb.Handle
}

func NewBridgePendingActivities(handle *rdb.Handle) *BridgePendingActivities {
	return &BridgePendingActivities{
		handle,
	}
}

func (thisView *BridgePendingActivities) ListAllUnprocessedOutgoing() ([]BridgePendingActivityReadRow, error) {
	return thisView.List(BridgePendingActivitiesFilter{
		MaybeDirections:  []types.Direction{types.DIRECTION_OUTGOING},
		MaybeIsProcessed: primptr.Bool(false),
	}, BridgePendingActivitiesOrder{MaybeId: primptr.String(view.ORDER_ASC)})
}

func (thisView *BridgePendingActivities) ListAllUnprocessedIncoming() ([]BridgePendingActivityReadRow, error) {
	return thisView.List(BridgePendingActivitiesFilter{
		MaybeDirections:  []types.Direction{types.DIRECTION_INCOMING, types.DIRECTION_RESPONSE},
		MaybeIsProcessed: primptr.Bool(false),
	}, BridgePendingActivitiesOrder{MaybeId: primptr.String(view.ORDER_ASC)})
}

func (thisView *BridgePendingActivities) List(
	filter BridgePendingActivitiesFilter,
	order BridgePendingActivitiesOrder,
) ([]BridgePendingActivityReadRow, error) {
	stmtBuilder := thisView.rdb.StmtBuilder.Select(
		"id",
		"block_height",
		"block_time",
		"maybe_transaction_id",
		"bridge_type",
		"link_id",
		"direction",
		"from_chain_id",
		"maybe_from_address",
		"maybe_from_smart_contract_address",
		"to_chain_id",
		"to_address",
		"maybe_to_smart_contract_address",
		"maybe_channel_id",
		"amount",
		"maybe_denom",
		"maybe_bridge_fee_amount",
		"maybe_bridge_fee_denom",
		"status",
		"is_processed",
		"created_at",
		"updated_at",
	).From(
		TABLE_BRIDGE_PENDING_ACTIVITIES,
	)

	if filter.MaybeDirections != nil {
		var orCond sq.Or
		for _, direction := range filter.MaybeDirections {
			orCond = append(orCond, sq.Eq{
				"direction": direction,
			})
		}
		stmtBuilder = stmtBuilder.Where(orCond)
	}

	if filter.MaybeIsProcessed != nil {
		stmtBuilder = stmtBuilder.Where("is_processed = ?", filter.MaybeIsProcessed)
	}

	if order.MaybeId != nil {
		stmtBuilder = stmtBuilder.OrderBy(fmt.Sprintf("id %s", *order.MaybeId))
	}

	sql, sqlArgs, sqlErr := stmtBuilder.ToSql()
	if sqlErr != nil {
		return nil, fmt.Errorf("error building bridge pending activities list unprocessed outgoing SQL: %v, %w", sqlErr, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := thisView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, fmt.Errorf("error executing bridge pending activities list unprocessed outoing SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	rows := make([]BridgePendingActivityReadRow, 0)
	for rowsResult.Next() {
		var row BridgePendingActivityReadRow

		blockTimeReader := thisView.rdb.NtotReader()
		var amount string
		var maybeBridgeFeeAmount *string
		createdAtTimeReader := thisView.rdb.NtotReader()
		updatedAtTimeReader := thisView.rdb.NtotReader()
		if err = rowsResult.Scan(
			&row.Id,
			&row.BlockHeight,
			blockTimeReader.ScannableArg(),
			&row.MaybeTransactionId,
			&row.BridgeType,
			&row.LinkId,
			&row.Direction,
			&row.FromChainId,
			&row.MaybeFromAddress,
			&row.MaybeFromSmartContractAddress,
			&row.ToChainId,
			&row.ToAddress,
			&row.MaybeToSmartContractAddress,
			&row.MaybeChannelId,
			&amount,
			&row.MaybeDenom,
			&maybeBridgeFeeAmount,
			&row.MaybeBridgeFeeDenom,
			&row.Status,
			&row.IsProcessed,
			createdAtTimeReader.ScannableArg(),
			updatedAtTimeReader.ScannableArg(),
		); err != nil {
			return nil, fmt.Errorf("error scanning unprocessed outgoing bridge pending activities row: %v: %w", err, rdb.ErrQuery)
		}

		blockTime, blockTimeErr := blockTimeReader.Parse()
		if blockTimeErr != nil {
			return nil, fmt.Errorf("error parsing block time: %v", blockTimeErr)
		}
		row.BlockTime = blockTime

		rowAmount, rowAmountOk := coin.NewIntFromString(amount)
		if !rowAmountOk {
			return nil, fmt.Errorf("error parsing amount %s", amount)
		}
		row.Amount = rowAmount

		if maybeBridgeFeeAmount != nil {
			rowBridgeFeeAmount, rowBridgeFeeAmountOk := coin.NewIntFromString(*maybeBridgeFeeAmount)
			if !rowBridgeFeeAmountOk {
				return nil, fmt.Errorf("error parsing bridge fee amount %s", *maybeBridgeFeeAmount)
			}
			row.MaybeBridgeFeeAmount = &rowBridgeFeeAmount
		}

		createdAt, createdAtErr := createdAtTimeReader.Parse()
		if createdAtErr != nil {
			return nil, fmt.Errorf("error parsing created at: %v", createdAtErr)
		}
		row.CreatedAt = createdAt

		updatedAt, updatedAtErr := updatedAtTimeReader.Parse()
		if updatedAtErr != nil {
			return nil, fmt.Errorf("error parsing updated at: %v", updatedAtErr)
		}
		row.UpdatedAt = updatedAt

		rows = append(rows, row)
	}

	return rows, nil
}

func (thisView *BridgePendingActivities) UpdateToProcessed(id int64) error {
	sql, sqlArgs, sqlErr := thisView.rdb.StmtBuilder.Update(
		TABLE_BRIDGE_PENDING_ACTIVITIES,
	).Set(
		"is_processed", true,
	).Where(
		"id = ?", id,
	).ToSql()

	if sqlErr != nil {
		return fmt.Errorf("error building activiy row update to processed sql: %v: %w", sqlErr, rdb.ErrBuildSQLStmt)
	}

	result, err := thisView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error updating activity row to processed into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error updating activity row to processed into the table: no row updated: %w", rdb.ErrWrite)
	}

	return nil
}

func (thisView *BridgePendingActivities) Insert(activity *BridgePendingActivityInsertRow) error {
	timeNow := utctime.Now()
	var maybeBridgeFeeAmount *string
	if activity.MaybeBridgeFeeAmount != nil {
		maybeBridgeFeeAmount = primptr.String(activity.MaybeBridgeFeeAmount.String())
	}
	sql, sqlArgs, err := thisView.rdb.StmtBuilder.
		Insert(TABLE_BRIDGE_PENDING_ACTIVITIES).
		Columns(
			"block_height",
			"block_time",
			"maybe_transaction_id",
			"bridge_type",
			"link_id",
			"direction",
			"from_chain_id",
			"maybe_from_address",
			"maybe_from_smart_contract_address",
			"to_chain_id",
			"to_address",
			"maybe_to_smart_contract_address",
			"maybe_channel_id",
			"amount",
			"maybe_denom",
			"maybe_bridge_fee_amount",
			"maybe_bridge_fee_denom",
			"status",
			"is_processed",
			"created_at",
			"updated_at",
		).
		Values(
			activity.BlockHeight,
			thisView.rdb.Tton(activity.BlockTime),
			activity.MaybeTransactionId,
			activity.BridgeType,
			activity.LinkId,
			activity.Direction,
			activity.FromChainId,
			activity.MaybeFromAddress,
			activity.MaybeFromSmartContractAddress,
			activity.ToChainId,
			activity.ToAddress,
			activity.MaybeToSmartContractAddress,
			activity.MaybeChannelId,
			activity.Amount.String(),
			activity.MaybeDenom,
			maybeBridgeFeeAmount,
			activity.MaybeBridgeFeeDenom,
			activity.Status,
			activity.IsProcessed,
			thisView.rdb.Tton(&timeNow),
			thisView.rdb.Tton(&timeNow),
		).
		ToSql()

	if err != nil {
		return fmt.Errorf("error building activiy row insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := thisView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting activity row into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting activity row into the table: no row inserted: %w", rdb.ErrWrite)
	}

	return nil
}

type BridgePendingActivitiesFilter struct {
	MaybeDirections  []types.Direction
	MaybeIsProcessed *bool
}

type BridgePendingActivitiesOrder struct {
	MaybeId *view.ORDER
}

type BridgePendingActivityReadRow struct {
	BridgePendingActivityInsertRow

	Id        int64            `json:"id"`
	CreatedAt *utctime.UTCTime `json:"createdAt"`
	UpdatedAt *utctime.UTCTime `json:"updatedAt"`
}

type BridgePendingActivityInsertRow struct {
	BlockHeight                   int64            `json:"blockHeight"`
	BlockTime                     *utctime.UTCTime `json:"blockTime"`
	MaybeTransactionId            *string          `json:"maybeTransactionId"`
	BridgeType                    types.BridgeType `json:"bridgeType"`
	LinkId                        string           `json:"linkId"`
	Direction                     types.Direction  `json:"direction"`
	FromChainId                   string           `json:"fromChainId"`
	MaybeFromAddress              *string          `json:"maybeFromAddress"`
	MaybeFromSmartContractAddress *string          `json:"maybeFromSmartContractAddress"`
	ToChainId                     string           `json:"toChainId"`
	ToAddress                     string           `json:"toAddress"`
	MaybeToSmartContractAddress   *string          `json:"maybeToSmartContractAddress"`
	MaybeChannelId                *string          `json:"maybeChannelId"`
	Amount                        coin.Int         `json:"amount"`
	MaybeDenom                    *string          `json:"maybeDenom"`
	MaybeBridgeFeeAmount          *coin.Int        `json:"maybeBridgeFeeAmount"`
	MaybeBridgeFeeDenom           *string          `json:"maybeBridgeFeeDenom"`
	Status                        types.Status     `json:"status"`
	IsProcessed                   bool             `json:"isProcessed"`
}
