package bridge_pending_activity

import (
	"fmt"
	"strconv"

	"github.com/crypto-com/chain-indexing/projection/bridge_pending_activity/types"

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

type Config struct {
	ThisChainId           string `mapstructure:"this_chain_id"`
	ThisChainName         string `mapstructure:"this_chain_name"`
	CounterpartyChainName string `mapstructure:"counterparty_chain_name"`
	ChannelId             string `mapstructure:"channel_id"`
	StartingHeight        int64  `mapstructure:"starting_height"`
}

func NewBridgePendingActivity(logger applogger.Logger, rdbConn rdb.Conn) *BridgePendingActivity {
	var config Config

	return &BridgePendingActivity{
		rdbprojectionbase.NewRDbBaseWithOptions(
			rdbConn.ToHandle(), "BridgePendingActivity", rdbprojectionbase.Options{
				MaybeTable:     nil,
				MaybeConfigPtr: &config,
			}),

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

func (projection *BridgePendingActivity) Config() *Config {
	return projection.Base.Config().(*Config)
}

func (projection *BridgePendingActivity) HandleEvents(height int64, events []event_entity.Event) error {
	fmt.Println("---------------------------------")
	projection.logger.Debugf("%d, %+v", height, projection.Config())
	if height < projection.Config().StartingHeight {
		return nil
	}

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
			fmt.Println(msgIBCTransferTransfer.Params.SourceChannel, projection.Config().ChannelId)
			if msgIBCTransferTransfer.Params.SourceChannel != projection.Config().ChannelId {
				continue
			}
			if msgIBCTransferTransfer.TxSuccess() {
				if err := view.Insert(&bridge_pending_activity_view.BridgePendingActivityRow{
					BlockHeight:                   height,
					BlockTime:                     &blockTime,
					MaybeTransactionId:            primptr.String(msgIBCTransferTransfer.TxHash()),
					BridgeType:                    types.BRIDGE_TYPE_IBC,
					LinkId:                        strconv.FormatUint(msgIBCTransferTransfer.Params.PacketSequence, 10),
					Direction:                     types.DIRECTION_OUTGOING,
					FromChainId:                   projection.Config().ThisChainName,
					MaybeFromAddress:              primptr.String(msgIBCTransferTransfer.Params.Sender),
					MaybeFromSmartContractAddress: nil,
					ToChainId:                     projection.Config().CounterpartyChainName,
					ToAddress:                     msgIBCTransferTransfer.Params.Receiver,
					MaybeToSmartContractAddress:   nil,
					MaybeChannelId:                primptr.String(msgIBCTransferTransfer.Params.SourceChannel),
					Amount:                        coin.NewIntFromUint64(msgIBCTransferTransfer.Params.Token.Amount),
					MaybeDenom:                    primptr.String(msgIBCTransferTransfer.Params.Token.Denom),
					MaybeBridgeFeeAmount:          nil,
					MaybeBridgeFeeDenom:           nil,
					Status:                        types.STATUS_PENDING,
					IsProcessed:                   false,
				}); err != nil {
					return fmt.Errorf("error inserting record when MsgIBCTransferTransfer: %w", err)
				}
			} else {
				if err := view.Insert(&bridge_pending_activity_view.BridgePendingActivityRow{
					BlockHeight:                   height,
					BlockTime:                     &blockTime,
					MaybeTransactionId:            primptr.String(msgIBCTransferTransfer.TxHash()),
					BridgeType:                    types.BRIDGE_TYPE_IBC,
					LinkId:                        "",
					Direction:                     types.DIRECTION_OUTGOING,
					FromChainId:                   projection.Config().ThisChainName,
					MaybeFromAddress:              primptr.String(msgIBCTransferTransfer.Params.Sender),
					MaybeFromSmartContractAddress: nil,
					ToChainId:                     projection.Config().CounterpartyChainName,
					ToAddress:                     msgIBCTransferTransfer.Params.Receiver,
					MaybeToSmartContractAddress:   nil,
					MaybeChannelId:                primptr.String(msgIBCTransferTransfer.Params.SourceChannel),
					Amount:                        coin.NewIntFromUint64(msgIBCTransferTransfer.Params.Token.Amount),
					MaybeDenom:                    primptr.String(msgIBCTransferTransfer.Params.Token.Denom),
					MaybeBridgeFeeAmount:          nil,
					MaybeBridgeFeeDenom:           nil,
					Status:                        types.STATUS_FAILED,
					IsProcessed:                   false,
				}); err != nil {
					return fmt.Errorf("error inserting failed record when MsgIBCTransferTransfer: %w", err)
				}
			}

		} else if msgIBCRecvPacket, ok := event.(*event_usecase.MsgIBCRecvPacket); ok {
			if msgIBCTransferTransfer.Params.DestinationChannel != projection.Config().ChannelId {
				continue
			}
			if msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData != nil {
				var status types.Status
				if msgIBCRecvPacket.Params.PacketAck.MaybeError == nil {
					status = types.STATUS_COUNTERPARTY_CONFIRMED
				} else {
					status = types.STATUS_COUNTERPARTY_REJECTED
				}

				if err := view.Insert(&bridge_pending_activity_view.BridgePendingActivityRow{
					BlockHeight:                   height,
					BlockTime:                     &blockTime,
					MaybeTransactionId:            primptr.String(msgIBCRecvPacket.TxHash()),
					BridgeType:                    types.BRIDGE_TYPE_IBC,
					LinkId:                        strconv.FormatUint(msgIBCRecvPacket.Params.PacketSequence, 10),
					Direction:                     types.DIRECTION_INCOMING,
					FromChainId:                   projection.Config().CounterpartyChainName,
					MaybeFromAddress:              primptr.String(msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData.Sender),
					MaybeFromSmartContractAddress: nil,
					ToChainId:                     projection.Config().ThisChainName,
					ToAddress:                     msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData.Receiver,
					MaybeToSmartContractAddress:   nil,
					MaybeChannelId:                primptr.String(msgIBCRecvPacket.Params.Packet.DestinationChannel),
					Amount:                        coin.NewIntFromUint64(msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData.Amount.Uint64()),
					MaybeDenom:                    primptr.String(msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData.Denom),
					MaybeBridgeFeeAmount:          nil,
					MaybeBridgeFeeDenom:           nil,
					Status:                        status,
					IsProcessed:                   false,
				}); err != nil {
					return fmt.Errorf("error inserting record when MsgIBCRecvPacket: %w", err)
				}
			}

		} else if msgIBCAcknowledgement, ok := event.(*event_usecase.MsgIBCAcknowledgement); ok {
			if msgIBCTransferTransfer.Params.SourceChannel != projection.Config().ChannelId {
				continue
			}
			if msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData != nil {
				var status types.Status
				if msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.Success {
					status = types.STATUS_COUNTERPARTY_CONFIRMED
				} else {
					status = types.STATUS_COUNTERPARTY_REJECTED
				}

				if err := view.Insert(&bridge_pending_activity_view.BridgePendingActivityRow{
					BlockHeight:                   height,
					BlockTime:                     &blockTime,
					BridgeType:                    types.BRIDGE_TYPE_IBC,
					MaybeTransactionId:            primptr.String(msgIBCAcknowledgement.TxHash()),
					LinkId:                        strconv.FormatUint(msgIBCAcknowledgement.Params.PacketSequence, 10),
					Direction:                     types.DIRECTION_OUTGOING,
					FromChainId:                   projection.Config().CounterpartyChainName,
					MaybeFromAddress:              primptr.String(msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.Sender),
					MaybeFromSmartContractAddress: nil,
					ToChainId:                     projection.Config().ThisChainName,
					ToAddress:                     msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.Receiver,
					MaybeToSmartContractAddress:   nil,
					MaybeChannelId:                primptr.String(msgIBCAcknowledgement.Params.Packet.SourceChannel),
					Amount:                        coin.NewIntFromUint64(msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.Amount.Uint64()),
					MaybeDenom:                    primptr.String(msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.Denom),
					MaybeBridgeFeeAmount:          nil,
					MaybeBridgeFeeDenom:           nil,
					Status:                        status,
					IsProcessed:                   false,
				}); err != nil {
					return fmt.Errorf("error inserting record when MsgIBCAcknowledgement: %w", err)
				}
			}

		} else if msgIBCTimeout, ok := event.(*event_usecase.MsgIBCTimeout); ok {
			if msgIBCTransferTransfer.Params.SourceChannel != projection.Config().ChannelId {
				continue
			}
			if msgIBCTimeout.Params.MaybeMsgTransfer != nil {
				if err := view.Insert(&bridge_pending_activity_view.BridgePendingActivityRow{
					BlockHeight:                   height,
					BlockTime:                     &blockTime,
					BridgeType:                    types.BRIDGE_TYPE_IBC,
					MaybeTransactionId:            primptr.String(msgIBCTimeout.TxHash()),
					LinkId:                        strconv.FormatUint(msgIBCTimeout.Params.PacketSequence, 10),
					Direction:                     types.DIRECTION_OUTGOING,
					FromChainId:                   projection.Config().CounterpartyChainName,
					MaybeFromAddress:              nil,
					MaybeFromSmartContractAddress: nil,
					ToChainId:                     projection.Config().ThisChainName,
					ToAddress:                     msgIBCTimeout.Params.MaybeMsgTransfer.RefundReceiver,
					MaybeToSmartContractAddress:   nil,
					MaybeChannelId:                primptr.String(msgIBCTimeout.Params.Packet.SourceChannel),
					Amount:                        coin.NewIntFromUint64(msgIBCTimeout.Params.MaybeMsgTransfer.RefundAmount),
					MaybeDenom:                    primptr.String(msgIBCTimeout.Params.MaybeMsgTransfer.RefundDenom),
					MaybeBridgeFeeAmount:          nil,
					MaybeBridgeFeeDenom:           nil,
					Status:                        types.STATUS_COUNTERPARTY_REJECTED,
					IsProcessed:                   false,
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
