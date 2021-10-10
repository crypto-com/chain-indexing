package view

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/projection/bridge_pending_activity/types"

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

func (view *BridgePendingActivities) Insert(activity *BridgePendingActivityRow) error {
	sql, sqlArgs, err := view.rdb.StmtBuilder.
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
		).
		Values(
			activity.BlockHeight,
			view.rdb.Tton(activity.BlockTime),
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
			activity.MaybeBridgeFeeAmount,
			activity.MaybeBridgeFeeDenom,
			activity.Status,
			activity.IsProcessed,
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

type BridgePendingActivityRow struct {
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
