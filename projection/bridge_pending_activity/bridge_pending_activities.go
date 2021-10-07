package bridge_pending_activity

import (
	"fmt"
	"strconv"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	entity_projection "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	"github.com/crypto-com/chain-indexing/internal/primptr"
	"github.com/crypto-com/chain-indexing/internal/utctime"
	bridge_pending_activity_view "github.com/crypto-com/chain-indexing/projection/bridge_pending_activity/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ entity_projection.Projection = &BridgePendingActivity{}

type BridgePendingActivity struct {
	*rdbprojectionbase.Base

	rdbConn rdb.Conn
	logger  applogger.Logger
}

type ProjectionConfig struct {
	BridgePendingActivity Config `toml:"bridge_pending_activity"`
}

type Config struct {
	ThisChainId         string `toml:"this_chain_id"`
	ChannelId           string `toml:"channel_id"`
	CounterpartyChainId string `toml:"counterparty_chain_id"`
}

func NewBridgePendingActivity(logger applogger.Logger, rdbConn rdb.Conn) *BridgePendingActivity {
	var config ProjectionConfig
	projectionBase := rdbprojectionbase.NewRDbBaseWithOptions(
		rdbConn.ToHandle(), "BridgePendingActivity", rdbprojectionbase.Options{
			MaybeTable:     nil,
			MaybeConfigPtr: &config,
		})
	return &BridgePendingActivity{
		projectionBase,

		rdbConn,
		logger,
	}
}

func (_ *BridgePendingActivity) GetEventsToListen() []string {
	return []string{
		event_usecase.BLOCK_CREATED,

		event_usecase.MSG_IBC_TRANSFER_TRANSFER_CREATED,
		event_usecase.MSG_IBC_TRANSFER_TRANSFER_FAILED,
		event_usecase.MSG_IBC_RECV_PACKET_CREATED,
		event_usecase.MSG_IBC_ACKNOWLEDGEMENT_CREATED,
		event_usecase.MSG_IBC_TIMEOUT,
	}
}

func (projection *BridgePendingActivity) OnInit() error {
	return nil
}

func (projection *BridgePendingActivity) Config() Config {
	return projection.Base.Config().(ProjectionConfig).BridgePendingActivity
}

func (projection *BridgePendingActivity) HandleEvents(height int64, events []event_entity.Event) error {
	fmt.Println(projection.Config())
	rdbTx, err := projection.rdbConn.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction: %v", err)
	}

	committed := false
	defer func() {
		if !committed {
			_ = rdbTx.Rollback()
		}
	}()

	rdbTxHandle := rdbTx.ToHandle()

	view := bridge_pending_activity_view.NewBridgePendingActivities(rdbTxHandle)

	// Get the block time of current height
	var blockTime utctime.UTCTime
	for _, event := range events {
		if blockCreatedEvent, ok := event.(*event_usecase.BlockCreated); ok {
			blockTime = blockCreatedEvent.Block.Time
		}
	}

	for _, event := range events {
		if msgIBCTransferTransfer, ok := event.(*event_usecase.MsgIBCTransferTransfer); ok {
			if msgIBCTransferTransfer.TxSuccess() {
				if err := view.Insert(&bridge_pending_activity_view.BridgePendingActivityRow{
					BlockHeight:        height,
					BlockTime:          &blockTime,
					MaybeTransactionId: primptr.String(msgIBCTransferTransfer.TxHash()),
					LinkId:             strconv.FormatUint(msgIBCTransferTransfer.Params.PacketSequence, 10),
					FromChainId:        projection.Config().ThisChainId,
					MaybeFromAddress:   primptr.String(msgIBCTransferTransfer.Params.Sender),
					ToChainId:          projection.Config().CounterpartyChainId,
					ToAddress:          msgIBCTransferTransfer.Params.Receiver,
					MaybeChannelId:     primptr.String(msgIBCTransferTransfer.Params.SourceChannel),
					TokenId:            msgIBCTransferTransfer.Params.Token.Denom,
					Amount:             coin.NewIntFromUint64(msgIBCTransferTransfer.Params.Token.Amount),
					Status:             bridge_pending_activity_view.STATUS_PENDING,
					IsProcessed:        false,
				}); err != nil {
					return fmt.Errorf("error inserting record when MsgIBCTransferTransfer: %w", err)
				}
			} else {
				if err := view.Insert(&bridge_pending_activity_view.BridgePendingActivityRow{
					BlockHeight:        height,
					BlockTime:          &blockTime,
					MaybeTransactionId: primptr.String(msgIBCTransferTransfer.TxHash()),
					LinkId:             strconv.FormatUint(msgIBCTransferTransfer.Params.PacketSequence, 10),
					FromChainId:        projection.Config().ThisChainId,
					MaybeFromAddress:   primptr.String(msgIBCTransferTransfer.Params.Sender),
					ToChainId:          projection.Config().CounterpartyChainId,
					ToAddress:          msgIBCTransferTransfer.Params.Receiver,
					MaybeChannelId:     primptr.String(msgIBCTransferTransfer.Params.SourceChannel),
					TokenId:            msgIBCTransferTransfer.Params.Token.Denom,
					Amount:             coin.NewIntFromUint64(msgIBCTransferTransfer.Params.Token.Amount),
					Status:             bridge_pending_activity_view.STATUS_FAILED,
					IsProcessed:        false,
				}); err != nil {
					return fmt.Errorf("error inserting failed record when MsgIBCTransferTransfer: %w", err)
				}
			}

		} else if msgIBCRecvPacket, ok := event.(*event_usecase.MsgIBCRecvPacket); ok {
			if msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData != nil {
				var status bridge_pending_activity_view.Status
				if msgIBCRecvPacket.Params.PacketAck.MaybeError == nil {
					status = bridge_pending_activity_view.STATUS_COUNTERPARTY_CONFIRMED
				} else {
					status = bridge_pending_activity_view.STATUS_COUNTERPARTY_REJECTED
				}

				if err := view.Insert(&bridge_pending_activity_view.BridgePendingActivityRow{
					BlockHeight:        height,
					BlockTime:          &blockTime,
					MaybeTransactionId: primptr.String(msgIBCRecvPacket.TxHash()),
					LinkId:             strconv.FormatUint(msgIBCRecvPacket.Params.PacketSequence, 10),
					FromChainId:        projection.Config().CounterpartyChainId,
					MaybeFromAddress:   primptr.String(msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData.Sender),
					ToChainId:          projection.Config().ThisChainId,
					ToAddress:          msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData.Receiver,
					MaybeChannelId:     primptr.String(msgIBCRecvPacket.Params.Packet.DestinationChannel),
					TokenId:            msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData.Denom,
					Amount:             coin.NewIntFromUint64(msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData.Amount.Uint64()),
					Status:             status,
					IsProcessed:        false,
				}); err != nil {
					return fmt.Errorf("error inserting record when MsgIBCRecvPacket: %w", err)
				}
			}

		} else if msgIBCAcknowledgement, ok := event.(*event_usecase.MsgIBCAcknowledgement); ok {
			if msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData != nil {
				var status bridge_pending_activity_view.Status
				if msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.Success {
					status = bridge_pending_activity_view.STATUS_COUNTERPARTY_CONFIRMED
				} else {
					status = bridge_pending_activity_view.STATUS_COUNTERPARTY_REJECTED
				}

				if err := view.Insert(&bridge_pending_activity_view.BridgePendingActivityRow{
					BlockHeight:        height,
					BlockTime:          &blockTime,
					MaybeTransactionId: primptr.String(msgIBCAcknowledgement.TxHash()),
					LinkId:             strconv.FormatUint(msgIBCAcknowledgement.Params.PacketSequence, 10),
					FromChainId:        projection.Config().CounterpartyChainId,
					MaybeFromAddress:   primptr.String(msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.Sender),
					ToChainId:          projection.Config().ThisChainId,
					ToAddress:          msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.Receiver,
					MaybeChannelId:     primptr.String(msgIBCAcknowledgement.Params.Packet.DestinationChannel),
					TokenId:            msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.Denom,
					Amount:             coin.NewIntFromUint64(msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.Amount.Uint64()),
					Status:             status,
					IsProcessed:        false,
				}); err != nil {
					return fmt.Errorf("error inserting record when MsgIBCAcknowledgement: %w", err)
				}
			}

		} else if msgIBCTimeout, ok := event.(*event_usecase.MsgIBCTimeout); ok {
			if msgIBCTimeout.Params.MaybeMsgTransfer != nil {
				if err := view.Insert(&bridge_pending_activity_view.BridgePendingActivityRow{
					BlockHeight:        height,
					BlockTime:          &blockTime,
					MaybeTransactionId: primptr.String(msgIBCTimeout.TxHash()),
					LinkId:             strconv.FormatUint(msgIBCTimeout.Params.PacketSequence, 10),
					FromChainId:        projection.Config().CounterpartyChainId,
					MaybeFromAddress:   nil,
					ToChainId:          projection.Config().ThisChainId,
					ToAddress:          msgIBCTimeout.Params.MaybeMsgTransfer.RefundReceiver,
					MaybeChannelId:     primptr.String(msgIBCTimeout.Params.Packet.SourceChannel),
					TokenId:            msgIBCTimeout.Params.MaybeMsgTransfer.RefundDenom,
					Amount:             coin.NewIntFromUint64(msgIBCTimeout.Params.MaybeMsgTransfer.RefundAmount),
					Status:             bridge_pending_activity_view.STATUS_COUNTERPARTY_REJECTED,
					IsProcessed:        false,
				}); err != nil {
					return fmt.Errorf("error inserting record when MsgIBCTimeout: %w", err)
				}
			}

		}
	}

	if err := projection.UpdateLastHandledEventHeight(rdbTxHandle, height); err != nil {
		return fmt.Errorf("error updating last handled event height: %v", err)
	}

	if err := rdbTx.Commit(); err != nil {
		return fmt.Errorf("error committing changes: %v", err)
	}
	committed = true

	return nil
}
