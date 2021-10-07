package view

import (
	"fmt"

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
			"link_id",
			"from_chain_id",
			"maybe_from_address",
			"to_chain_id",
			"to_address",
			"maybe_channel_id",
			"token_id",
			"amount",
			"status",
			"is_processed",
		).
		Values(
			activity.BlockHeight,
			view.rdb.Tton(activity.BlockTime),
			activity.MaybeTransactionId,
			activity.LinkId,
			activity.FromChainId,
			activity.MaybeFromAddress,
			activity.ToChainId,
			activity.ToAddress,
			activity.MaybeChannelId,
			activity.TokenId,
			activity.Amount.String(),
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
	BlockHeight        int64            `json:"blockHeight"`
	BlockTime          *utctime.UTCTime `json:"blockTime"`
	MaybeTransactionId *string          `json:"maybeTransactionId"`
	LinkId             string           `json:"linkId"`
	FromChainId        string           `json:"fromChainId"`
	MaybeFromAddress   *string          `json:"maybeFromAddress"`
	ToChainId          string           `json:"toChainId"`
	ToAddress          string           `json:"toAddress"`
	MaybeChannelId     *string          `json:"maybeChannelId"`
	TokenId            string           `json:"tokenId"`
	Amount             coin.Int         `json:"amount"`
	Status             Status           `json:"status"`
	IsProcessed        bool             `json:"isProcessed"`
}
