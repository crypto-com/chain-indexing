package bridge_pending_activity

import (
	"fmt"
	"strconv"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	entity_projection "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper"
	"github.com/crypto-com/chain-indexing/projection/bridge_activity/types"
	bridge_pending_activity_view "github.com/crypto-com/chain-indexing/projection/bridge_activity/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var (
	NewBridgePendingActivities   = bridge_pending_activity_view.NewBridgePendingActivitiesView
	UpdateLastHandledEventHeight = (*BridgePendingActivity).UpdateLastHandledEventHeight
)

var _ entity_projection.Projection = &BridgePendingActivity{}

type ChannelId = string

type BridgePendingActivity struct {
	*rdbprojectionbase.Base

	rdbConn rdb.Conn
	logger  applogger.Logger

	migrationHelper migrationhelper.MigrationHelper

	listenedChannelIdToConfig map[ChannelId]*CounterPartyChainConfig
}

type Config struct {
	ThisChainName      string                    `mapstructure:"this_chain_name"`
	CounterPartyChains []CounterPartyChainConfig `mapstructure:"counteparty"`
}

type CounterPartyChainConfig struct {
	ChainName      string `mapstructure:"chain_name"`
	ChannelId      string `mapstructure:"channel_id"`
	StartingHeight int64  `mapstructure:"starting_height"`
}

const (
	MIGRATION_DIRECOTRY = "projection/bridge_activity/bridge_pending_activity/migrations"
)

func New(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	migrationHelper migrationhelper.MigrationHelper,
) *BridgePendingActivity {
	return &BridgePendingActivity{
		rdbprojectionbase.NewRDbBaseWithOptions(
			rdbConn.ToHandle(),
			"BridgePendingActivity",
			rdbprojectionbase.Options{
				MaybeTable:         nil,
				MaybeReadConfigPtr: &Config{},
			},
		),

		rdbConn,
		logger,
		migrationHelper,
		make(map[ChannelId]*CounterPartyChainConfig),
	}
}

func NewWithConfig(
	config *Config,
	logger applogger.Logger,
	rdbConn rdb.Conn,
	migrationHelper migrationhelper.MigrationHelper,
) *BridgePendingActivity {
	return &BridgePendingActivity{
		rdbprojectionbase.NewRDbBaseWithOptions(
			rdbConn.ToHandle(),
			"BridgePendingActivity",
			rdbprojectionbase.Options{
				MaybeTable:     nil,
				MaybeConfigPtr: config,
			},
		),

		rdbConn,
		logger,
		migrationHelper,
		make(map[ChannelId]*CounterPartyChainConfig),
	}
}

func (_ *BridgePendingActivity) GetEventsToListen() []string {
	return []string{
		event_usecase.BLOCK_CREATED,

		event_usecase.MSG_IBC_TRANSFER_TRANSFER_CREATED,
		event_usecase.MSG_IBC_TRANSFER_TRANSFER_FAILED,
		event_usecase.MSG_IBC_RECV_PACKET_CREATED,
		event_usecase.MSG_IBC_ACKNOWLEDGEMENT_CREATED,
		event_usecase.MSG_IBC_TIMEOUT_CREATED,

		event_usecase.CRONOS_SEND_TO_IBC_CREATED,
	}
}

func (projection *BridgePendingActivity) OnInit() error {
	if projection.migrationHelper != nil {
		projection.migrationHelper.Migrate()
	}

	counterPartyChainConfigs := projection.Config().CounterPartyChains
	for i, chainConfig := range projection.Config().CounterPartyChains {
		projection.listenedChannelIdToConfig[chainConfig.ChannelId] = &counterPartyChainConfigs[i]
	}
	return nil
}

func (projection *BridgePendingActivity) Config() *Config {
	return projection.Base.Config().(*Config)
}

func (projection *BridgePendingActivity) HandleEvents(height int64, events []event_entity.Event) error {
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

	commit := func() error {
		if err := UpdateLastHandledEventHeight(projection, rdbTxHandle, height); err != nil {
			return fmt.Errorf("error updating last handled event height: %v", err)
		}

		if err := rdbTx.Commit(); err != nil {
			return fmt.Errorf("error committing changes: %v", err)
		}
		committed = true

		return nil
	}

	view := NewBridgePendingActivities(rdbTxHandle)

	// Get the block time of current height
	var blockTime utctime.UTCTime
	for _, event := range events {
		if blockCreatedEvent, ok := event.(*event_usecase.BlockCreated); ok {
			blockTime = blockCreatedEvent.Block.Time
		}
	}

	for _, event := range events {
		if msgIBCTransferTransfer, ok := event.(*event_usecase.MsgIBCTransferTransfer); ok {
			channelId := msgIBCTransferTransfer.Params.SourceChannel
			if !projection.isListenedChannel(channelId) {
				continue
			}
			counterpartyConfig := projection.mustGetCounterpartyChainConfigByListenedChannel(channelId)
			if msgIBCTransferTransfer.TxSuccess() {
				tokenAmount, tokenAmountOk := coin.NewIntFromString(msgIBCTransferTransfer.Params.Token.Amount.String())
				if !tokenAmountOk {
					return fmt.Errorf(
						"error creating coin from token amount: %s",
						msgIBCTransferTransfer.Params.Token.Amount.String(),
					)
				}
				if err := view.Insert(&bridge_pending_activity_view.BridgePendingActivityInsertRow{
					BlockHeight:                   height,
					BlockTime:                     &blockTime,
					MaybeTransactionId:            primptr.String(msgIBCTransferTransfer.TxHash()),
					BridgeType:                    types.BRIDGE_TYPE_IBC,
					LinkId:                        ibcLinkId(projection.Config().ThisChainName, msgIBCTransferTransfer.Params.PacketSequence),
					Direction:                     types.DIRECTION_OUTGOING,
					FromChainId:                   projection.Config().ThisChainName,
					MaybeFromAddress:              primptr.String(msgIBCTransferTransfer.Params.Sender),
					MaybeFromSmartContractAddress: nil,
					ToChainId:                     counterpartyConfig.ChainName,
					ToAddress:                     msgIBCTransferTransfer.Params.Receiver,
					MaybeToSmartContractAddress:   nil,
					MaybeChannelId:                primptr.String(msgIBCTransferTransfer.Params.SourceChannel),
					Amount:                        tokenAmount,
					MaybeDenom:                    primptr.String(msgIBCTransferTransfer.Params.Token.Denom),
					MaybeBridgeFeeAmount:          nil,
					MaybeBridgeFeeDenom:           nil,
					Status:                        types.STATUS_PENDING,
					IsProcessed:                   false,
				}); err != nil {
					return fmt.Errorf("error inserting record when MsgIBCTransferTransfer: %w", err)
				}
			} else {
				tokenAmount, tokenAmountOk := coin.NewIntFromString(msgIBCTransferTransfer.Params.Token.Amount.String())
				if !tokenAmountOk {
					return fmt.Errorf(
						"error creating coing from token amount: %s",
						msgIBCTransferTransfer.Params.Token.Amount.String(),
					)
				}
				if err := view.Insert(&bridge_pending_activity_view.BridgePendingActivityInsertRow{
					BlockHeight:        height,
					BlockTime:          &blockTime,
					MaybeTransactionId: primptr.String(msgIBCTransferTransfer.TxHash()),
					BridgeType:         types.BRIDGE_TYPE_IBC,
					LinkId: fmt.Sprintf(
						"source:%s-transactionId:%s-failedOnChain",
						projection.Config().ThisChainName, msgIBCTransferTransfer.TxHash(),
					),
					Direction:                     types.DIRECTION_OUTGOING,
					FromChainId:                   projection.Config().ThisChainName,
					MaybeFromAddress:              primptr.String(msgIBCTransferTransfer.Params.Sender),
					MaybeFromSmartContractAddress: nil,
					ToChainId:                     counterpartyConfig.ChainName,
					ToAddress:                     msgIBCTransferTransfer.Params.Receiver,
					MaybeToSmartContractAddress:   nil,
					MaybeChannelId:                primptr.String(msgIBCTransferTransfer.Params.SourceChannel),
					Amount:                        tokenAmount,
					MaybeDenom:                    primptr.String(msgIBCTransferTransfer.Params.Token.Denom),
					MaybeBridgeFeeAmount:          nil,
					MaybeBridgeFeeDenom:           nil,
					Status:                        types.STATUS_FAILED_ON_CHAIN,
					IsProcessed:                   false,
				}); err != nil {
					return fmt.Errorf("error inserting failed record when MsgIBCTransferTransfer: %w", err)
				}
			}

		} else if msgIBCRecvPacket, ok := event.(*event_usecase.MsgIBCRecvPacket); ok {
			channelId := msgIBCRecvPacket.Params.Packet.DestinationChannel
			if !projection.isListenedChannel(channelId) {
				continue
			}
			counterpartyConfig := projection.mustGetCounterpartyChainConfigByListenedChannel(channelId)
			if msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData != nil {
				var status types.Status
				var isProcessed bool
				if msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData.Success {
					status = types.STATUS_COUNTERPARTY_CONFIRMED
					isProcessed = false
				} else {
					status = types.STATUS_COUNTERPARTY_REJECTED
					isProcessed = false
				}

				amount, amountOk := coin.NewIntFromString(msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData.Amount.String())
				if !amountOk {
					return fmt.Errorf(
						"error creating coin from token amount: %s",
						msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData.Amount.String(),
					)
				}

				if err := view.Insert(&bridge_pending_activity_view.BridgePendingActivityInsertRow{
					BlockHeight:                   height,
					BlockTime:                     &blockTime,
					MaybeTransactionId:            primptr.String(msgIBCRecvPacket.TxHash()),
					BridgeType:                    types.BRIDGE_TYPE_IBC,
					LinkId:                        ibcLinkId(counterpartyConfig.ChainName, msgIBCRecvPacket.Params.PacketSequence),
					Direction:                     types.DIRECTION_INCOMING,
					FromChainId:                   counterpartyConfig.ChainName,
					MaybeFromAddress:              primptr.String(msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData.Sender),
					MaybeFromSmartContractAddress: nil,
					ToChainId:                     projection.Config().ThisChainName,
					ToAddress:                     msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData.Receiver,
					MaybeToSmartContractAddress:   nil,
					MaybeChannelId:                primptr.String(msgIBCRecvPacket.Params.Packet.DestinationChannel),
					Amount:                        amount,
					MaybeDenom:                    primptr.String(msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData.Denom),
					MaybeBridgeFeeAmount:          nil,
					MaybeBridgeFeeDenom:           nil,
					Status:                        status,
					IsProcessed:                   isProcessed,
				}); err != nil {
					return fmt.Errorf("error inserting record when MsgIBCRecvPacket: %w", err)
				}
			}

		} else if msgIBCAcknowledgement, ok := event.(*event_usecase.MsgIBCAcknowledgement); ok {
			channelId := msgIBCAcknowledgement.Params.Packet.SourceChannel
			if !projection.isListenedChannel(channelId) {
				continue
			}
			counterpartyConfig := projection.mustGetCounterpartyChainConfigByListenedChannel(channelId)
			if msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData != nil {
				var status types.Status
				var isProcessed bool
				if msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.Success {
					status = types.STATUS_COUNTERPARTY_CONFIRMED
					isProcessed = false
				} else {
					status = types.STATUS_COUNTERPARTY_REJECTED
					isProcessed = false
				}

				amount, amountOk := coin.NewIntFromString(msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.Amount.String())
				if !amountOk {
					return fmt.Errorf(
						"error creating coin from token amount: %s",
						msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.Amount.String(),
					)
				}

				if err := view.Insert(&bridge_pending_activity_view.BridgePendingActivityInsertRow{
					BlockHeight:                   height,
					BlockTime:                     &blockTime,
					BridgeType:                    types.BRIDGE_TYPE_IBC,
					MaybeTransactionId:            primptr.String(msgIBCAcknowledgement.TxHash()),
					LinkId:                        ibcLinkId(projection.Config().ThisChainName, msgIBCAcknowledgement.Params.PacketSequence),
					Direction:                     types.DIRECTION_RESPONSE,
					FromChainId:                   projection.Config().ThisChainName,
					MaybeFromAddress:              primptr.String(msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.Sender),
					MaybeFromSmartContractAddress: nil,
					ToChainId:                     counterpartyConfig.ChainName,
					ToAddress:                     msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.Receiver,
					MaybeToSmartContractAddress:   nil,
					MaybeChannelId:                primptr.String(msgIBCAcknowledgement.Params.Packet.SourceChannel),
					Amount:                        amount,
					MaybeDenom:                    primptr.String(msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.Denom),
					MaybeBridgeFeeAmount:          nil,
					MaybeBridgeFeeDenom:           nil,
					Status:                        status,
					IsProcessed:                   isProcessed,
				}); err != nil {
					return fmt.Errorf("error inserting record when MsgIBCAcknowledgement: %w", err)
				}
			}

		} else if msgIBCTimeout, ok := event.(*event_usecase.MsgIBCTimeout); ok {
			channelId := msgIBCTimeout.Params.Packet.SourceChannel
			if !projection.isListenedChannel(channelId) {
				continue
			}
			counterpartyConfig := projection.mustGetCounterpartyChainConfigByListenedChannel(channelId)

			amount, amountOk := coin.NewIntFromString(msgIBCTimeout.Params.MaybeMsgTransfer.RefundAmount.String())
			if !amountOk {
				return fmt.Errorf(
					"error creating coin from token amount: %s",
					msgIBCTimeout.Params.MaybeMsgTransfer.RefundAmount.String(),
				)
			}

			if msgIBCTimeout.Params.MaybeMsgTransfer != nil {
				if err := view.Insert(&bridge_pending_activity_view.BridgePendingActivityInsertRow{
					BlockHeight:                   height,
					BlockTime:                     &blockTime,
					BridgeType:                    types.BRIDGE_TYPE_IBC,
					MaybeTransactionId:            primptr.String(msgIBCTimeout.TxHash()),
					LinkId:                        ibcLinkId(projection.Config().ThisChainName, msgIBCTimeout.Params.PacketSequence),
					Direction:                     types.DIRECTION_RESPONSE,
					FromChainId:                   projection.Config().ThisChainName,
					MaybeFromAddress:              nil,
					MaybeFromSmartContractAddress: nil,
					ToChainId:                     counterpartyConfig.ChainName,
					ToAddress:                     msgIBCTimeout.Params.MaybeMsgTransfer.RefundReceiver,
					MaybeToSmartContractAddress:   nil,
					MaybeChannelId:                primptr.String(msgIBCTimeout.Params.Packet.SourceChannel),
					Amount:                        amount,
					MaybeDenom:                    primptr.String(msgIBCTimeout.Params.MaybeMsgTransfer.RefundDenom),
					MaybeBridgeFeeAmount:          nil,
					MaybeBridgeFeeDenom:           nil,
					Status:                        types.STATUS_COUNTERPARTY_REJECTED,
					IsProcessed:                   false,
				}); err != nil {
					return fmt.Errorf("error inserting record when MsgIBCTimeout: %w", err)
				}
			}

		} else if cronosSendToIBCCreatedEvent, ok := event.(*event_usecase.CronosSendToIBCCreated); ok {
			channelId := cronosSendToIBCCreatedEvent.Params.SourceChannel
			if !projection.isListenedChannel(channelId) {
				continue
			}
			counterpartyConfig := projection.mustGetCounterpartyChainConfigByListenedChannel(channelId)

			amount, amountOk := coin.NewIntFromString(cronosSendToIBCCreatedEvent.Params.Token.Amount.String())
			if !amountOk {
				return fmt.Errorf(
					"error creating coin from token amount: %s",
					msgIBCTimeout.Params.MaybeMsgTransfer.RefundAmount.String(),
				)
			}

			if err := view.Insert(&bridge_pending_activity_view.BridgePendingActivityInsertRow{
				BlockHeight:                   height,
				BlockTime:                     &blockTime,
				MaybeTransactionId:            primptr.String(cronosSendToIBCCreatedEvent.Params.EthereumTxHash),
				BridgeType:                    types.BRIDGE_TYPE_IBC,
				LinkId:                        ibcLinkId(projection.Config().ThisChainName, cronosSendToIBCCreatedEvent.Params.PacketSequence),
				Direction:                     types.DIRECTION_OUTGOING,
				FromChainId:                   projection.Config().ThisChainName,
				MaybeFromAddress:              primptr.String(cronosSendToIBCCreatedEvent.Params.Sender),
				MaybeFromSmartContractAddress: nil,
				ToChainId:                     counterpartyConfig.ChainName,
				ToAddress:                     cronosSendToIBCCreatedEvent.Params.Receiver,
				MaybeToSmartContractAddress:   nil,
				MaybeChannelId:                primptr.String(cronosSendToIBCCreatedEvent.Params.SourceChannel),
				Amount:                        amount,
				MaybeDenom:                    primptr.String(cronosSendToIBCCreatedEvent.Params.Token.Denom),
				MaybeBridgeFeeAmount:          nil,
				MaybeBridgeFeeDenom:           nil,
				Status:                        types.STATUS_PENDING,
				IsProcessed:                   false,
			}); err != nil {
				return fmt.Errorf("error inserting record when CronosSendToIBCCreated: %w", err)
			}
		}
	}

	return commit()
}

func (projection *BridgePendingActivity) isListenedChannel(channelId string) bool {
	_, exists := projection.listenedChannelIdToConfig[channelId]
	return exists
}

func (projection *BridgePendingActivity) mustGetCounterpartyChainConfigByListenedChannel(channelId string) *CounterPartyChainConfig {
	if !projection.isListenedChannel(channelId) {
		panic(fmt.Sprintf("channel id %s not found", channelId))
	}
	return projection.listenedChannelIdToConfig[channelId]
}

func ibcLinkId(sourceChain string, sequence uint64) string {
	return fmt.Sprintf("source:%s-sequence:%s", sourceChain, strconv.FormatUint(sequence, 10))
}
