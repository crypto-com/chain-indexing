package bridge_pending_activity

import (
	"fmt"
	"strconv"

	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	appprojection "github.com/crypto-com/chain-indexing/projection"
	"github.com/crypto-com/chain-indexing/projection/bridge_activity/types"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	entity_projection "github.com/crypto-com/chain-indexing/entity/projection"
	"github.com/crypto-com/chain-indexing/external/utctime"
	github_mh "github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper/github"
	bridge_pending_activity_view "github.com/crypto-com/chain-indexing/projection/bridge_activity/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ entity_projection.Projection = &BridgePendingActivity{}

type BridgePendingActivity struct {
	*rdbprojectionbase.Base

	rdbConn rdb.Conn
	logger  applogger.Logger

	config *appprojection.Config
}

type Config struct {
	ThisChainId           string `mapstructure:"this_chain_id"`
	ThisChainName         string `mapstructure:"this_chain_name"`
	CounterpartyChainName string `mapstructure:"counterparty_chain_name"`
	ChannelId             string `mapstructure:"channel_id"`
	StartingHeight        int64  `mapstructure:"starting_height"`
}

const (
	MIGRATION_DIRECOTRY = "projection/bridge_activity/bridge_pending_activity/migrations"
)

func NewBridgePendingActivity(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	config *appprojection.Config,
) *BridgePendingActivity {

	// Here we customize the migrationSourceURL, as it is not following the same path convention
	// in other projections. (extra `bridge_activity` folder in the path)
	migrationSourceURL := github_mh.GenerateSourceURL(
		github_mh.MIGRATION_GITHUB_URL_FORMAT,
		config.GithubAPIUser,
		config.GithubAPIToken,
		MIGRATION_DIRECOTRY,
		config.MigrationRepoRef,
	)

	return &BridgePendingActivity{
		rdbprojectionbase.NewRDbBaseWithOptions(
			rdbConn.ToHandle(),
			"BridgePendingActivity",

			rdbprojectionbase.Options{
				MaybeTable:     nil,
				MaybeConfigPtr: &Config{},
				MaybeMigrationHelper: github_mh.NewGithubMigrationHelper(
					github_mh.Config{
						Config:     *config,
						ConnString: rdbConn.(*pg.PgxConn).ConnString(),
					},
					&migrationSourceURL,
					nil,
				),
			},
		),

		rdbConn,
		logger,
		config,
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

		event_usecase.GRAVITY_ETHEREUM_SEND_TO_COSMOS_HANDLED,
	}
}

func (projection *BridgePendingActivity) OnInit() error {
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
		if err := projection.UpdateLastHandledEventHeight(rdbTxHandle, height); err != nil {
			return fmt.Errorf("error updating last handled event height: %v", err)
		}

		if err := rdbTx.Commit(); err != nil {
			return fmt.Errorf("error committing changes: %v", err)
		}
		committed = true

		return nil
	}

	if height < projection.Config().StartingHeight {
		return commit()
	}

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
			if msgIBCTransferTransfer.Params.SourceChannel != projection.Config().ChannelId {
				continue
			}
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
					ToChainId:                     projection.Config().CounterpartyChainName,
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
					ToChainId:                     projection.Config().CounterpartyChainName,
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
			if msgIBCRecvPacket.Params.Packet.DestinationChannel != projection.Config().ChannelId {
				continue
			}
			if msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData != nil {
				var status types.Status
				var isProcessed bool
				if msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData.Success {
					status = types.STATUS_COUNTERPARTY_CONFIRMED
					isProcessed = false
				} else {
					if msgIBCRecvPacket.Params.PacketAck.MaybeError != nil {
						status = types.STATUS_COUNTERPARTY_REJECTED
						isProcessed = false
					} else {
						status = types.STATUS_NO_OPERATION
						isProcessed = true
					}
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
					LinkId:                        ibcLinkId(projection.Config().CounterpartyChainName, msgIBCRecvPacket.Params.PacketSequence),
					Direction:                     types.DIRECTION_INCOMING,
					FromChainId:                   projection.Config().CounterpartyChainName,
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
			if msgIBCAcknowledgement.Params.Packet.SourceChannel != projection.Config().ChannelId {
				continue
			}
			if msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData != nil {
				var status types.Status
				var isProcessed bool
				if msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.Success {
					status = types.STATUS_COUNTERPARTY_CONFIRMED
					isProcessed = false
				} else {
					if msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.MaybeError != nil {
						status = types.STATUS_COUNTERPARTY_REJECTED
						isProcessed = false
					} else {
						status = types.STATUS_NO_OPERATION
						isProcessed = true
					}
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
					ToChainId:                     projection.Config().CounterpartyChainName,
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
			if msgIBCTimeout.Params.Packet.SourceChannel != projection.Config().ChannelId {
				continue
			}

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
					ToChainId:                     projection.Config().CounterpartyChainName,
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
			if cronosSendToIBCCreatedEvent.Params.SourceChannel != projection.Config().ChannelId {
				continue
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
				ToChainId:                     projection.Config().CounterpartyChainName,
				ToAddress:                     cronosSendToIBCCreatedEvent.Params.Receiver,
				MaybeToSmartContractAddress:   nil,
				MaybeChannelId:                primptr.String(cronosSendToIBCCreatedEvent.Params.SourceChannel),
				Amount:                        coin.NewIntFromUint64(cronosSendToIBCCreatedEvent.Params.Token.Amount.Uint64()),
				MaybeDenom:                    primptr.String(cronosSendToIBCCreatedEvent.Params.Token.Denom),
				MaybeBridgeFeeAmount:          nil,
				MaybeBridgeFeeDenom:           nil,
				Status:                        types.STATUS_PENDING,
				IsProcessed:                   false,
			}); err != nil {
				return fmt.Errorf("error inserting record when CronosSendToIBCCreated: %w", err)
			}

		} else if ethereumSendToCosmosHandledEvent, ok := event.(*event_usecase.GravityEthereumSendToCosmosHandled); ok {
			if len(ethereumSendToCosmosHandledEvent.Params.Amount) == 0 {
				continue
			}

			sourceChainName := types.CHAIN_ETHEREUM
			if err := view.Insert(&bridge_pending_activity_view.BridgePendingActivityInsertRow{
				BlockHeight:        height,
				BlockTime:          &blockTime,
				BridgeType:         types.BRIDGE_TYPE_GRAVITY,
				MaybeTransactionId: nil,
				LinkId: fmt.Sprintf(
					"source:%s-receiver:%s-amount:%s-contract:%s",
					sourceChainName,
					ethereumSendToCosmosHandledEvent.Params.Receiver,
					ethereumSendToCosmosHandledEvent.Params.Amount,
					ethereumSendToCosmosHandledEvent.Params.EthereumTokenContract,
				),
				Direction:                     types.DIRECTION_INCOMING,
				FromChainId:                   sourceChainName, // TODO: Parse Params.BridgeChainId?
				MaybeFromAddress:              primptr.String(ethereumSendToCosmosHandledEvent.Params.Sender),
				MaybeFromSmartContractAddress: primptr.String(ethereumSendToCosmosHandledEvent.Params.EthereumTokenContract),
				ToChainId:                     projection.Config().ThisChainName,
				ToAddress:                     ethereumSendToCosmosHandledEvent.Params.Receiver,
				MaybeToSmartContractAddress:   nil,
				MaybeChannelId:                nil,
				Amount:                        ethereumSendToCosmosHandledEvent.Params.Amount[0].Amount,
				MaybeDenom:                    primptr.String(ethereumSendToCosmosHandledEvent.Params.Amount[0].Denom),
				MaybeBridgeFeeAmount:          nil,
				MaybeBridgeFeeDenom:           nil,
				Status:                        types.STATUS_COUNTERPARTY_CONFIRMED,
				IsProcessed:                   false,
			}); err != nil {
				return fmt.Errorf("error inserting record when GravityEthereumSendToCosmosHandled: %w", err)
			}
		}
	}

	return commit()
}

func ibcLinkId(sourceChain string, sequence uint64) string {
	return fmt.Sprintf("source:%s-sequence:%s", sourceChain, strconv.FormatUint(sequence, 10))
}
