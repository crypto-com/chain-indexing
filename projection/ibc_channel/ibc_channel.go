package ibc_channel

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	entity_projection "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper"
	"github.com/crypto-com/chain-indexing/projection/ibc_channel/types"
	ibc_channel_view "github.com/crypto-com/chain-indexing/projection/ibc_channel/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ entity_projection.Projection = &IBCChannel{}

var (
	NewIBCChannels               = ibc_channel_view.NewIBCChannelsView
	NewIBCClients                = ibc_channel_view.NewIBCClientsView
	NewIBCConnections            = ibc_channel_view.NewIBCConnectionsView
	NewIBCDenomHashMapping       = ibc_channel_view.NewIBCDenomHashMappingView
	NewIBCChannelTraces          = ibc_channel_view.NewIBCChannelTraces
	UpdateLastHandledEventHeight = (*IBCChannel).UpdateLastHandledEventHeight
)

type IBCChannel struct {
	*rdbprojectionbase.Base

	rdbConn rdb.Conn
	logger  applogger.Logger

	config          *Config
	migrationHelper migrationhelper.MigrationHelper
}

type Config struct {
	EnableTxMsgTrace bool
}

func NewIBCChannel(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	config *Config,
	migrationHelper migrationhelper.MigrationHelper,
) *IBCChannel {
	projectionID := "IBCChannel"
	if config.EnableTxMsgTrace {
		projectionID = "IBCChannelTxMsgTrace"
	}
	return &IBCChannel{
		rdbprojectionbase.NewRDbBase(
			rdbConn.ToHandle(),
			projectionID,
		),

		rdbConn,
		logger,

		config,
		migrationHelper,
	}
}

func (_ *IBCChannel) GetEventsToListen() []string {
	return []string{
		event_usecase.BLOCK_CREATED,

		event_usecase.MSG_IBC_CREATE_CLIENT_CREATED,
		event_usecase.MSG_IBC_CONNECTION_OPEN_INIT_CREATED,
		event_usecase.MSG_IBC_CONNECTION_OPEN_TRY_CREATED,
		event_usecase.MSG_IBC_CONNECTION_OPEN_ACK_CREATED,
		event_usecase.MSG_IBC_CONNECTION_OPEN_CONFIRM_CREATED,
		event_usecase.MSG_IBC_CHANNEL_OPEN_INIT_CREATED,
		event_usecase.MSG_IBC_CHANNEL_OPEN_TRY_CREATED,
		event_usecase.MSG_IBC_CHANNEL_OPEN_ACK_CREATED,
		event_usecase.MSG_IBC_CHANNEL_OPEN_CONFIRM_CREATED,
		event_usecase.MSG_IBC_CHANNEL_CLOSE_INIT_CREATED,
		event_usecase.MSG_IBC_CHANNEL_CLOSE_CONFIRM_CREATED,
		event_usecase.MSG_IBC_TRANSFER_TRANSFER_CREATED,
		event_usecase.MSG_IBC_RECV_PACKET_CREATED,
		event_usecase.MSG_IBC_ACKNOWLEDGEMENT_CREATED,
		event_usecase.MSG_IBC_TIMEOUT_CREATED,
		event_usecase.MSG_IBC_TIMEOUT_ON_CLOSE_CREATED,
		event_usecase.CHAINMAIN_MSG_REGISTER_ACCOUNT_CREATED,
		event_usecase.MSG_REGISTER_ACCOUNT_CREATED,
	}
}

// Projection IBCChannelTxMsgTrace needs the below const, for debug purpose
const (
	MIGRATION_TABLE_NAME = "ibc_channel_schema_migrations"
	MIGRATION_DIRECOTRY  = "projection/ibc_channel/migrations"
)

func (projection *IBCChannel) OnInit() error {
	if projection.migrationHelper != nil {
		projection.migrationHelper.Migrate()
	}

	return nil
}

func (projection *IBCChannel) HandleEvents(height int64, events []event_entity.Event) error {
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

	ibcChannelsView := NewIBCChannels(rdbTxHandle)
	ibcClientsView := NewIBCClients(rdbTxHandle)
	ibcConnectionsView := NewIBCConnections(rdbTxHandle)
	ibcChannelTracesView := NewIBCChannelTraces(rdbTxHandle)
	ibcDenomHashMappingView := NewIBCDenomHashMapping(rdbTxHandle)

	// Get the block time of current height
	var blockTime utctime.UTCTime
	for _, event := range events {
		if blockCreatedEvent, ok := event.(*event_usecase.BlockCreated); ok {
			blockTime = blockCreatedEvent.Block.Time
		}
	}

	// NOTES: Why four channel open events are all needed?
	//
	// PacketOrdering info can only be found in MsgIBCChannelOpenInit and MsgIBCChannelOpenTry.
	//
	// Full information of ConnectionID, CounterpartyChannelID, and CounterpartyPortID
	// can only be found in MsgIBCChannelOpenAck and MsgIBCChannelOpenConfirm.

	for _, event := range events {
		if msgIBCCreateClient, ok := event.(*event_usecase.MsgIBCCreateClient); ok {
			var client *ibc_channel_view.IBCClientRow
			if msgIBCCreateClient.Params.MaybeTendermintLightClient != nil {
				client = &ibc_channel_view.IBCClientRow{
					ClientID:            msgIBCCreateClient.Params.ClientID,
					CounterpartyChainID: msgIBCCreateClient.Params.MaybeTendermintLightClient.TendermintClientState.ChainID,
				}
			} else if msgIBCCreateClient.Params.MaybeSoloMachineLightClient != nil {
				client = &ibc_channel_view.IBCClientRow{
					ClientID:            msgIBCCreateClient.Params.ClientID,
					CounterpartyChainID: msgIBCCreateClient.Params.MaybeSoloMachineLightClient.SoloMachineLightClientConsensusState.Diversifier,
				}
			} else if msgIBCCreateClient.Params.MaybeLocalhostLightClient != nil {
				client = &ibc_channel_view.IBCClientRow{
					ClientID:            msgIBCCreateClient.Params.ClientID,
					CounterpartyChainID: msgIBCCreateClient.Params.MaybeLocalhostLightClient.LocalhostClientState.ChainID,
				}
			}
			if err := ibcClientsView.Insert(client); err != nil {
				return fmt.Errorf("error inserting client: %w", err)
			}

		} else if msgIBCConnectionOpenInit, ok := event.(*event_usecase.MsgIBCConnectionOpenInit); ok {

			clientID := msgIBCConnectionOpenInit.Params.ClientID
			counterpartyChainID, err := ibcClientsView.FindCounterpartyChainIDBy(clientID)
			if err != nil {
				return fmt.Errorf("error in finding counterparty_chain_id: %w", err)
			}

			connection := &ibc_channel_view.IBCConnectionRow{
				ConnectionID:             msgIBCConnectionOpenInit.Params.ConnectionID,
				ClientID:                 msgIBCConnectionOpenInit.Params.ClientID,
				CounterpartyConnectionID: msgIBCConnectionOpenInit.Params.Counterparty.ConnectionID,
				CounterpartyClientID:     msgIBCConnectionOpenInit.Params.Counterparty.ClientID,
				CounterpartyChainID:      counterpartyChainID,
			}
			if err := ibcConnectionsView.Insert(connection); err != nil {
				return fmt.Errorf("error inserting connection: %w", err)
			}

		} else if msgIBCConnectionOpenTry, ok := event.(*event_usecase.MsgIBCConnectionOpenTry); ok {

			clientID := msgIBCConnectionOpenTry.Params.ClientID
			counterpartyChainID, err := ibcClientsView.FindCounterpartyChainIDBy(clientID)
			if err != nil {
				return fmt.Errorf("error in finding counterparty_chain_id: %w", err)
			}

			connection := &ibc_channel_view.IBCConnectionRow{
				ConnectionID:             msgIBCConnectionOpenTry.Params.ConnectionID,
				ClientID:                 msgIBCConnectionOpenTry.Params.ClientID,
				CounterpartyConnectionID: msgIBCConnectionOpenTry.Params.Counterparty.ConnectionID,
				CounterpartyClientID:     msgIBCConnectionOpenTry.Params.Counterparty.ClientID,
				CounterpartyChainID:      counterpartyChainID,
			}
			if err := ibcConnectionsView.Insert(connection); err != nil {
				return fmt.Errorf("error inserting connection: %w", err)
			}

		} else if msgIBCConnectionOpenConfirm, ok := event.(*event_usecase.MsgIBCConnectionOpenConfirm); ok {

			clientID := msgIBCConnectionOpenConfirm.Params.ClientID
			counterpartyChainID, err := ibcClientsView.FindCounterpartyChainIDBy(clientID)
			if err != nil {
				return fmt.Errorf("error in finding counterparty_chain_id: %w", err)
			}

			connection := &ibc_channel_view.IBCConnectionRow{
				ConnectionID:             msgIBCConnectionOpenConfirm.Params.ConnectionID,
				ClientID:                 msgIBCConnectionOpenConfirm.Params.ClientID,
				CounterpartyConnectionID: msgIBCConnectionOpenConfirm.Params.CounterpartyConnectionID,
				CounterpartyClientID:     msgIBCConnectionOpenConfirm.Params.CounterpartyClientID,
				CounterpartyChainID:      counterpartyChainID,
			}
			if err := ibcConnectionsView.Update(connection); err != nil {
				return fmt.Errorf("error updating connection: %w", err)
			}

		} else if msgIBCConnectionOpenAck, ok := event.(*event_usecase.MsgIBCConnectionOpenAck); ok {

			clientID := msgIBCConnectionOpenAck.Params.ClientID
			counterpartyChainID, err := ibcClientsView.FindCounterpartyChainIDBy(clientID)
			if err != nil {
				return fmt.Errorf("error in finding counterparty_chain_id: %w", err)
			}

			connection := &ibc_channel_view.IBCConnectionRow{
				ConnectionID:             msgIBCConnectionOpenAck.Params.ConnectionID,
				ClientID:                 msgIBCConnectionOpenAck.Params.ClientID,
				CounterpartyConnectionID: msgIBCConnectionOpenAck.Params.CounterpartyConnectionID,
				CounterpartyClientID:     msgIBCConnectionOpenAck.Params.CounterpartyClientID,
				CounterpartyChainID:      counterpartyChainID,
			}
			if err := ibcConnectionsView.Update(connection); err != nil {
				return fmt.Errorf("error updating connection: %w", err)
			}

		} else if msgIBCChannelOpenInit, ok := event.(*event_usecase.MsgIBCChannelOpenInit); ok {

			connectionID := msgIBCChannelOpenInit.Params.ConnectionID
			counterpartyChainID, err := ibcConnectionsView.FindCounterpartyChainIDBy(connectionID)
			if err != nil {
				return fmt.Errorf("error in finding counterparty_chain_id: %w", err)
			}

			channel := &ibc_channel_view.IBCChannelRow{
				ChannelID:                 msgIBCChannelOpenInit.Params.ChannelID,
				PortID:                    msgIBCChannelOpenInit.Params.PortID,
				ConnectionID:              msgIBCChannelOpenInit.Params.ConnectionID,
				CounterpartyChannelID:     msgIBCChannelOpenInit.Params.Channel.Counterparty.ChannelID,
				CounterpartyPortID:        msgIBCChannelOpenInit.Params.Channel.Counterparty.PortID,
				CounterpartyChainID:       counterpartyChainID,
				Status:                    types.STATUS_NOT_ESTABLISHED,
				PacketOrdering:            msgIBCChannelOpenInit.Params.Channel.Ordering,
				LastInPacketSequence:      0,
				LastOutPacketSequence:     0,
				TotalRelayInCount:         0,
				TotalRelayOutCount:        0,
				TotalRelayOutSuccessCount: 0,
				TotalRelayOutSuccessRate:  0,
				CreatedAtBlockTime:        utctime.FromUnixNano(0),
				CreatedAtBlockHeight:      0,
				Verified:                  false,
				Description:               "",
				LastActivityBlockTime:     utctime.FromUnixNano(0),
				LastActivityBlockHeight:   0,
				BondedTokens:              *ibc_channel_view.NewEmptyBondedTokens(),
			}
			if err := ibcChannelsView.Insert(channel); err != nil {
				return fmt.Errorf("error inserting channel: %w", err)
			}

		} else if msgRegisterAccount, ok := event.(*event_usecase.MsgRegisterAccount); ok {

			connectionID := msgRegisterAccount.Params.ConnectionID
			counterpartyChainID, err := ibcConnectionsView.FindCounterpartyChainIDBy(connectionID)
			if err != nil {
				return fmt.Errorf("error in finding counterparty_chain_id: %w", err)
			}

			channel := &ibc_channel_view.IBCChannelRow{
				ChannelID:                 msgRegisterAccount.Params.ChannelID,
				PortID:                    msgRegisterAccount.Params.PortID,
				ConnectionID:              msgRegisterAccount.Params.ConnectionID,
				CounterpartyChannelID:     msgRegisterAccount.Params.CounterpartyChannelID,
				CounterpartyPortID:        msgRegisterAccount.Params.CounterpartyPortID,
				CounterpartyChainID:       counterpartyChainID,
				Status:                    types.STATUS_NOT_ESTABLISHED,
				PacketOrdering:            "ORDER_ORDERED",
				LastInPacketSequence:      0,
				LastOutPacketSequence:     0,
				TotalRelayInCount:         0,
				TotalRelayOutCount:        0,
				TotalRelayOutSuccessCount: 0,
				TotalRelayOutSuccessRate:  0,
				CreatedAtBlockTime:        utctime.FromUnixNano(0),
				CreatedAtBlockHeight:      0,
				Verified:                  false,
				Description:               "",
				LastActivityBlockTime:     utctime.FromUnixNano(0),
				LastActivityBlockHeight:   0,
				BondedTokens:              *ibc_channel_view.NewEmptyBondedTokens(),
			}

			if err := ibcChannelsView.Insert(channel); err != nil {
				return fmt.Errorf("error inserting channel: %w", err)
			}

		} else if msgRegisterAccount, ok := event.(*event_usecase.ChainmainMsgRegisterAccount); ok {

			connectionID := msgRegisterAccount.Params.ConnectionID
			counterpartyChainID, err := ibcConnectionsView.FindCounterpartyChainIDBy(connectionID)
			if err != nil {
				return fmt.Errorf("error in finding counterparty_chain_id: %w", err)
			}

			channel := &ibc_channel_view.IBCChannelRow{
				ChannelID:                 msgRegisterAccount.Params.ChannelID,
				PortID:                    msgRegisterAccount.Params.PortID,
				ConnectionID:              msgRegisterAccount.Params.ConnectionID,
				CounterpartyChannelID:     msgRegisterAccount.Params.CounterpartyChannelID,
				CounterpartyPortID:        msgRegisterAccount.Params.CounterpartyPortID,
				CounterpartyChainID:       counterpartyChainID,
				Status:                    types.STATUS_NOT_ESTABLISHED,
				PacketOrdering:            "ORDER_ORDERED",
				LastInPacketSequence:      0,
				LastOutPacketSequence:     0,
				TotalRelayInCount:         0,
				TotalRelayOutCount:        0,
				TotalRelayOutSuccessCount: 0,
				TotalRelayOutSuccessRate:  0,
				CreatedAtBlockTime:        utctime.FromUnixNano(0),
				CreatedAtBlockHeight:      0,
				Verified:                  false,
				Description:               "",
				LastActivityBlockTime:     utctime.FromUnixNano(0),
				LastActivityBlockHeight:   0,
				BondedTokens:              *ibc_channel_view.NewEmptyBondedTokens(),
			}

			if err := ibcChannelsView.Insert(channel); err != nil {
				return fmt.Errorf("error inserting channel: %w", err)
			}

		} else if msgIBCChannelOpenTry, ok := event.(*event_usecase.MsgIBCChannelOpenTry); ok {

			connectionID := msgIBCChannelOpenTry.Params.ConnectionID
			counterpartyChainID, err := ibcConnectionsView.FindCounterpartyChainIDBy(connectionID)
			if err != nil {
				return fmt.Errorf("error in finding counterparty_chain_id: %w", err)
			}

			channel := &ibc_channel_view.IBCChannelRow{
				ChannelID:                 msgIBCChannelOpenTry.Params.ChannelID,
				PortID:                    msgIBCChannelOpenTry.Params.PortID,
				ConnectionID:              msgIBCChannelOpenTry.Params.ConnectionID,
				CounterpartyChannelID:     msgIBCChannelOpenTry.Params.Channel.Counterparty.ChannelID,
				CounterpartyPortID:        msgIBCChannelOpenTry.Params.Channel.Counterparty.PortID,
				CounterpartyChainID:       counterpartyChainID,
				Status:                    types.STATUS_NOT_ESTABLISHED,
				PacketOrdering:            msgIBCChannelOpenTry.Params.Channel.Ordering,
				LastInPacketSequence:      0,
				LastOutPacketSequence:     0,
				TotalRelayInCount:         0,
				TotalRelayOutCount:        0,
				TotalRelayOutSuccessCount: 0,
				TotalRelayOutSuccessRate:  0,
				CreatedAtBlockTime:        utctime.FromUnixNano(0),
				CreatedAtBlockHeight:      0,
				Verified:                  false,
				Description:               "",
				LastActivityBlockTime:     utctime.FromUnixNano(0),
				LastActivityBlockHeight:   0,
				BondedTokens:              *ibc_channel_view.NewEmptyBondedTokens(),
			}
			if err := ibcChannelsView.Insert(channel); err != nil {
				return fmt.Errorf("error inserting channel: %w", err)
			}

		} else if msgIBCChannelOpenAck, ok := event.(*event_usecase.MsgIBCChannelOpenAck); ok {

			connectionID := msgIBCChannelOpenAck.Params.ConnectionID
			counterpartyChainID, err := ibcConnectionsView.FindCounterpartyChainIDBy(connectionID)
			if err != nil {
				return fmt.Errorf("error in finding counterparty_chain_id: %w", err)
			}

			channel := &ibc_channel_view.IBCChannelRow{
				ChannelID:             msgIBCChannelOpenAck.Params.ChannelID,
				PortID:                msgIBCChannelOpenAck.Params.PortID,
				ConnectionID:          msgIBCChannelOpenAck.Params.ConnectionID,
				CounterpartyChannelID: msgIBCChannelOpenAck.Params.CounterpartyChannelID,
				CounterpartyPortID:    msgIBCChannelOpenAck.Params.CounterpartyPortID,
				CounterpartyChainID:   counterpartyChainID,
				CreatedAtBlockTime:    blockTime,
				CreatedAtBlockHeight:  height,
			}
			if err := ibcChannelsView.UpdateFactualColumns(channel); err != nil {
				return fmt.Errorf("error updating channel: %w", err)
			}

			if err := ibcChannelsView.UpdateStatus(channel.ChannelID, types.STATUS_OPENED); err != nil {
				return fmt.Errorf("error updating channel established: %w", err)
			}

		} else if msgIBCChannelOpenConfirm, ok := event.(*event_usecase.MsgIBCChannelOpenConfirm); ok {

			connectionID := msgIBCChannelOpenConfirm.Params.ConnectionID
			counterpartyChainID, err := ibcConnectionsView.FindCounterpartyChainIDBy(connectionID)
			if err != nil {
				return fmt.Errorf("error in finding counterparty_chain_id: %w", err)
			}

			channel := &ibc_channel_view.IBCChannelRow{
				ChannelID:             msgIBCChannelOpenConfirm.Params.ChannelID,
				PortID:                msgIBCChannelOpenConfirm.Params.PortID,
				ConnectionID:          msgIBCChannelOpenConfirm.Params.ConnectionID,
				CounterpartyChannelID: msgIBCChannelOpenConfirm.Params.CounterpartyChannelID,
				CounterpartyPortID:    msgIBCChannelOpenConfirm.Params.CounterpartyPortID,
				CounterpartyChainID:   counterpartyChainID,
				CreatedAtBlockTime:    blockTime,
				CreatedAtBlockHeight:  height,
			}
			if err := ibcChannelsView.UpdateFactualColumns(channel); err != nil {
				return fmt.Errorf("error updating channel: %w", err)
			}

			if err := ibcChannelsView.UpdateStatus(channel.ChannelID, types.STATUS_OPENED); err != nil {
				return fmt.Errorf("error updating channel established: %w", err)
			}

		} else if msgIBCTransferTransfer, ok := event.(*event_usecase.MsgIBCTransferTransfer); ok {

			// Transfer started by source chain
			channelID := msgIBCTransferTransfer.Params.SourceChannel

			// TotalRelayOutSuccessRate
			if err := ibcChannelsView.Increment(channelID, "total_relay_out_count", 1); err != nil {
				return fmt.Errorf("error increasing total_relay_out_count: %w", err)
			}
			if err := ibcChannelsView.UpdateTotalRelayOutSuccessRate(channelID); err != nil {
				return fmt.Errorf("error updating total_relay_out_success_rate: %w", err)
			}

			lastOutPacketSequence := msgIBCTransferTransfer.Params.PacketSequence
			if err := ibcChannelsView.UpdateSequence(channelID, "last_out_packet_sequence", lastOutPacketSequence); err != nil {
				return fmt.Errorf("error updating last_out_packet_sequence: %w", err)
			}

			if err := ibcChannelsView.UpdateLastActivityTimeAndHeight(channelID, blockTime, height); err != nil {
				return fmt.Errorf("error updating channel last_activity_time: %w", err)
			}

			var amount string
			if msgIBCTransferTransfer.Params.PacketData.Amount == nil {
				return nil
			}
			amount = msgIBCTransferTransfer.Params.PacketData.Amount.String()
			denom := msgIBCTransferTransfer.Params.PacketData.Denom
			destinationChannelID := msgIBCTransferTransfer.Params.DestinationChannel
			destinationPortID := msgIBCTransferTransfer.Params.DestinationPort
			sourceChannelID := msgIBCTransferTransfer.Params.SourceChannel
			sourcePortID := msgIBCTransferTransfer.Params.SourcePort

			if err := projection.updateBondedTokensWhenMsgIBCTransfer(
				ibcChannelsView,
				channelID,
				amount,
				denom,
				destinationChannelID,
				destinationPortID,
				sourceChannelID,
				sourcePortID,
			); err != nil {
				return fmt.Errorf("error updateBondedTokensWhenMsgIBCTransfer: %v", err)
			}

			if projection.config.EnableTxMsgTrace {

				msg, err := msgIBCTransferTransfer.ToJSON()
				if err != nil {
					return fmt.Errorf("error msgIBCTransferTransfer.ToJSON(): %w", err)
				}

				// Here the bonded_tokens has already been updated by the above updateBondedTokensWhenXXXX()
				updatedBondedTokensJSON, err := getBondedTokensInJSON(ibcChannelsView, channelID)
				if err != nil {
					return fmt.Errorf("error getBondedTokensInJSON: %v", err)
				}

				if err := ibcChannelTracesView.Insert(&ibc_channel_view.IBCChannelTraceRow{
					ChannelID:           channelID,
					BlockHeight:         height,
					SourceChannel:       msgIBCTransferTransfer.Params.SourceChannel,
					DestinationChannel:  msgIBCTransferTransfer.Params.DestinationChannel,
					Denom:               msgIBCTransferTransfer.Params.PacketData.Denom,
					Amount:              msgIBCTransferTransfer.Params.PacketData.Amount.String(),
					Success:             "",
					Error:               "",
					MessageType:         msgIBCTransferTransfer.MsgName,
					Message:             msg,
					UpdatedBondedTokens: updatedBondedTokensJSON,
				}); err != nil {
					return fmt.Errorf("error adding tx trace when MsgIBCTransferTransfer: %w", err)
				}

			}

		} else if msgIBCRecvPacket, ok := event.(*event_usecase.MsgIBCRecvPacket); ok {

			// Transfer started by destination chain
			channelID := msgIBCRecvPacket.Params.Packet.DestinationChannel

			if err := ibcChannelsView.Increment(channelID, "total_relay_in_count", 1); err != nil {
				return fmt.Errorf("error increasing total_relay_in_count: %w", err)
			}

			lastInPacketSequence := msgIBCRecvPacket.Params.PacketSequence
			if err := ibcChannelsView.UpdateSequence(channelID, "last_in_packet_sequence", lastInPacketSequence); err != nil {
				return fmt.Errorf("error updating last_in_packet_sequence: %w", err)
			}

			if err := ibcChannelsView.UpdateLastActivityTimeAndHeight(channelID, blockTime, height); err != nil {
				return fmt.Errorf("error updating channel last_activity_time: %w", err)
			}

			if msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData != nil &&
				msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData.Success {
				if msgIBCTransferTransfer.Params.PacketData.Amount == nil {
					return nil
				}

				amount := msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData.Amount.String()
				denom := msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData.Denom
				destinationChannelID := msgIBCRecvPacket.Params.Packet.DestinationChannel
				destinationPortID := msgIBCRecvPacket.Params.Packet.DestinationPort
				sourceChannelID := msgIBCRecvPacket.Params.Packet.SourceChannel
				sourcePortID := msgIBCRecvPacket.Params.Packet.SourcePort

				if err := projection.updateBondedTokensWhenMsgIBCRecvPacket(
					ibcChannelsView,
					ibcDenomHashMappingView,
					channelID,
					amount,
					denom,
					destinationChannelID,
					destinationPortID,
					sourceChannelID,
					sourcePortID,
				); err != nil {
					return fmt.Errorf("error updateChannelBondedTokensWhenMsgIBCRecvPacket: %w", err)
				}

			}

			if projection.config.EnableTxMsgTrace &&
				msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData != nil {
				if msgIBCTransferTransfer.Params.PacketData.Amount == nil {
					return nil
				}

				msg, err := msgIBCRecvPacket.ToJSON()
				if err != nil {
					return fmt.Errorf("error msgIBCRecvPacket.ToJSON(): %w", err)
				}
				success := strconv.FormatBool(msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData.Success)

				maybeError := ""
				if msgIBCRecvPacket.Params.PacketAck.MaybeError != nil {
					maybeError = *msgIBCRecvPacket.Params.PacketAck.MaybeError
				}

				// Here the bonded_tokens has already been updated by the above updateBondedTokensWhenXXXX()
				updatedBondedTokensJSON, err := getBondedTokensInJSON(ibcChannelsView, channelID)
				if err != nil {
					return fmt.Errorf("error getBondedTokensInJSON: %v", err)
				}

				if err := ibcChannelTracesView.Insert(&ibc_channel_view.IBCChannelTraceRow{
					ChannelID:           channelID,
					BlockHeight:         height,
					SourceChannel:       msgIBCRecvPacket.Params.Packet.SourceChannel,
					DestinationChannel:  msgIBCRecvPacket.Params.Packet.DestinationChannel,
					Denom:               msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData.Denom,
					Amount:              msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData.Amount.String(),
					Success:             success,
					Error:               maybeError,
					MessageType:         msgIBCRecvPacket.MsgName,
					Message:             msg,
					UpdatedBondedTokens: updatedBondedTokensJSON,
				}); err != nil {
					return fmt.Errorf("error adding tx trace when MsgIBCRecvPacket: %w", err)
				}

			}

		} else if msgIBCAcknowledgement, ok := event.(*event_usecase.MsgIBCAcknowledgement); ok {

			// Transfer started by source chain
			channelID := msgIBCAcknowledgement.Params.Packet.SourceChannel

			if err := ibcChannelsView.UpdateLastActivityTimeAndHeight(channelID, blockTime, height); err != nil {
				return fmt.Errorf("error updating channel last_activity_time: %w", err)
			}

			// TotalRelayOutSuccessRate
			if err := ibcChannelsView.Increment(channelID, "total_relay_out_success_count", 1); err != nil {
				return fmt.Errorf("error increasing total_relay_out_success_count: %w", err)
			}
			if err := ibcChannelsView.UpdateTotalRelayOutSuccessRate(channelID); err != nil {
				return fmt.Errorf("error updating total_relay_out_success_rate: %w", err)
			}

			if msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData != nil {

				if !msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.Success {

					amount := msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.Amount.String()
					denom := msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.Denom
					destinationChannelID := msgIBCAcknowledgement.Params.Packet.DestinationChannel
					destinationPortID := msgIBCAcknowledgement.Params.Packet.DestinationPort
					sourceChannelID := msgIBCAcknowledgement.Params.Packet.SourceChannel
					sourcePortID := msgIBCAcknowledgement.Params.Packet.SourcePort

					if err := revertUpdateBondedTokensWhenMsgIBCTransfer(
						ibcChannelsView,
						channelID,
						amount,
						denom,
						destinationChannelID,
						destinationPortID,
						sourceChannelID,
						sourcePortID,
					); err != nil {
						return fmt.Errorf("error revertUpdateBondedTokensWhenMsgIBCTransfer: %v", err)
					}

				}

			}

			if projection.config.EnableTxMsgTrace && msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData != nil {

				msg, err := msgIBCAcknowledgement.ToJSON()
				if err != nil {
					return fmt.Errorf("error msgIBCAcknowledgement.ToJSON(): %w", err)
				}
				success := strconv.FormatBool(msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.Success)

				maybeError := ""
				if msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.MaybeError != nil {
					maybeError = *msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.MaybeError
				}

				// Here the bonded_tokens has already been updated by the above updateBondedTokensWhenXXXX()
				updatedBondedTokensJSON, err := getBondedTokensInJSON(ibcChannelsView, channelID)
				if err != nil {
					return fmt.Errorf("error getBondedTokensInJSON: %v", err)
				}

				if err := ibcChannelTracesView.Insert(&ibc_channel_view.IBCChannelTraceRow{
					ChannelID:           channelID,
					BlockHeight:         height,
					SourceChannel:       msgIBCAcknowledgement.Params.Packet.SourceChannel,
					DestinationChannel:  msgIBCAcknowledgement.Params.Packet.DestinationChannel,
					Denom:               msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.Denom,
					Amount:              msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.Amount.String(),
					Success:             success,
					Error:               maybeError,
					MessageType:         msgIBCAcknowledgement.MsgName,
					Message:             msg,
					UpdatedBondedTokens: updatedBondedTokensJSON,
				}); err != nil {
					return fmt.Errorf("error adding tx trace when MsgIBCAcknowledgement: %w", err)
				}

			}

		} else if msgIBCTimeout, ok := event.(*event_usecase.MsgIBCTimeout); ok {

			// Transfer started by source chain
			channelID := msgIBCTimeout.Params.Packet.SourceChannel

			if err := ibcChannelsView.UpdateLastActivityTimeAndHeight(channelID, blockTime, height); err != nil {
				return fmt.Errorf("error updating channel last_activity_time: %w", err)
			}

			// TotalRelayOutSuccessRate
			if err := ibcChannelsView.Increment(channelID, "total_relay_out_success_count", 1); err != nil {
				return fmt.Errorf("error increasing total_relay_out_success_count: %w", err)
			}
			if err := ibcChannelsView.UpdateTotalRelayOutSuccessRate(channelID); err != nil {
				return fmt.Errorf("error updating total_relay_out_success_rate: %w", err)
			}

			if msgIBCTimeout.Params.MaybeMsgTransfer != nil {

				amount := msgIBCTimeout.Params.MaybeMsgTransfer.RefundAmount.String()
				denom := msgIBCTimeout.Params.MaybeMsgTransfer.RefundDenom
				destinationChannelID := msgIBCTimeout.Params.Packet.DestinationChannel
				destinationPortID := msgIBCTimeout.Params.Packet.DestinationPort
				sourceChannelID := msgIBCTimeout.Params.Packet.SourceChannel
				sourcePortID := msgIBCTimeout.Params.Packet.SourcePort

				if err := revertUpdateBondedTokensWhenMsgIBCTransfer(
					ibcChannelsView,
					channelID,
					amount,
					denom,
					destinationChannelID,
					destinationPortID,
					sourceChannelID,
					sourcePortID,
				); err != nil {
					return fmt.Errorf("error revertUpdateBondedTokensWhenMsgIBCTransfer: %v", err)
				}

			}

			if projection.config.EnableTxMsgTrace && msgIBCTimeout.Params.MaybeMsgTransfer != nil {

				amount := msgIBCTimeout.Params.MaybeMsgTransfer.RefundAmount.String()
				msg, err := msgIBCTimeout.ToJSON()
				if err != nil {
					return fmt.Errorf("error msgIBCTimeout.ToJSON(): %w", err)
				}

				// Here the bonded_tokens has already been updated by the above updateBondedTokensWhenXXXX()
				updatedBondedTokensJSON, err := getBondedTokensInJSON(ibcChannelsView, channelID)
				if err != nil {
					return fmt.Errorf("error getBondedTokensInJSON: %v", err)
				}

				if err := ibcChannelTracesView.Insert(&ibc_channel_view.IBCChannelTraceRow{
					ChannelID:           channelID,
					BlockHeight:         height,
					SourceChannel:       msgIBCTimeout.Params.Packet.SourceChannel,
					DestinationChannel:  msgIBCTimeout.Params.Packet.DestinationChannel,
					Denom:               msgIBCTimeout.Params.MaybeMsgTransfer.RefundDenom,
					Amount:              amount,
					Success:             "",
					Error:               "",
					MessageType:         msgIBCTimeout.MsgName,
					Message:             msg,
					UpdatedBondedTokens: updatedBondedTokensJSON,
				}); err != nil {
					return fmt.Errorf("error adding tx trace when MsgIBCTimeout: %w", err)
				}

			}

		} else if msgIBCTimeoutOnClose, ok := event.(*event_usecase.MsgIBCTimeoutOnClose); ok {

			// Transfer started by source chain
			channelID := msgIBCTimeoutOnClose.Params.Packet.SourceChannel

			if err := ibcChannelsView.UpdateLastActivityTimeAndHeight(channelID, blockTime, height); err != nil {
				return fmt.Errorf("error updating channel last_activity_time: %w", err)
			}

			// TotalRelayOutSuccessRate
			if err := ibcChannelsView.Increment(channelID, "total_relay_out_success_count", 1); err != nil {
				return fmt.Errorf("error increasing total_relay_out_success_count: %w", err)
			}
			if err := ibcChannelsView.UpdateTotalRelayOutSuccessRate(channelID); err != nil {
				return fmt.Errorf("error updating total_relay_out_success_rate: %w", err)
			}

			if msgIBCTimeoutOnClose.Params.MaybeMsgTransfer != nil {

				amount := msgIBCTimeoutOnClose.Params.MaybeMsgTransfer.RefundAmount
				denom := msgIBCTimeoutOnClose.Params.MaybeMsgTransfer.RefundDenom
				destinationChannelID := msgIBCTimeoutOnClose.Params.Packet.DestinationChannel
				destinationPortID := msgIBCTimeoutOnClose.Params.Packet.DestinationPort
				sourceChannelID := msgIBCTimeoutOnClose.Params.Packet.SourceChannel
				sourcePortID := msgIBCTimeoutOnClose.Params.Packet.SourcePort

				if err := revertUpdateBondedTokensWhenMsgIBCTransfer(
					ibcChannelsView,
					channelID,
					amount.String(),
					denom,
					destinationChannelID,
					destinationPortID,
					sourceChannelID,
					sourcePortID,
				); err != nil {
					return fmt.Errorf("error revertUpdateBondedTokensWhenMsgIBCTransfer: %v", err)
				}

			}

			if projection.config.EnableTxMsgTrace && msgIBCTimeoutOnClose.Params.MaybeMsgTransfer != nil {

				amount := msgIBCTimeoutOnClose.Params.MaybeMsgTransfer.RefundAmount.String()
				msg, err := msgIBCTimeoutOnClose.ToJSON()
				if err != nil {
					return fmt.Errorf("error msgIBCTimeoutOnClose.ToJSON(): %w", err)
				}

				// Here the bonded_tokens has already been updated by the above updateBondedTokensWhenXXXX()
				updatedBondedTokensJSON, err := getBondedTokensInJSON(ibcChannelsView, channelID)
				if err != nil {
					return fmt.Errorf("error getBondedTokensInJSON: %v", err)
				}

				if err := ibcChannelTracesView.Insert(&ibc_channel_view.IBCChannelTraceRow{
					ChannelID:           channelID,
					BlockHeight:         height,
					SourceChannel:       msgIBCTimeoutOnClose.Params.Packet.SourceChannel,
					DestinationChannel:  msgIBCTimeoutOnClose.Params.Packet.DestinationChannel,
					Denom:               msgIBCTimeoutOnClose.Params.MaybeMsgTransfer.RefundDenom,
					Amount:              amount,
					Success:             "",
					Error:               "",
					MessageType:         msgIBCTimeoutOnClose.MsgName,
					Message:             msg,
					UpdatedBondedTokens: updatedBondedTokensJSON,
				}); err != nil {
					return fmt.Errorf("error adding tx trace when MsgIBCTimeoutOnClose: %w", err)
				}

			}

		} else if msgIBCChannelCloseInit, ok := event.(*event_usecase.MsgIBCChannelCloseInit); ok {

			channelID := msgIBCChannelCloseInit.Params.ChannelID

			if err := ibcChannelsView.UpdateStatus(channelID, types.STATUS_CLOSED); err != nil {
				return fmt.Errorf("error updating channel closed: %w", err)
			}

			if err := ibcChannelsView.UpdateLastActivityTimeAndHeight(channelID, blockTime, height); err != nil {
				return fmt.Errorf("error updating channel last_activity_time: %w", err)
			}

		} else if msgIBCChannelCloseConfirm, ok := event.(*event_usecase.MsgIBCChannelCloseConfirm); ok {

			channelID := msgIBCChannelCloseConfirm.Params.ChannelID

			if err := ibcChannelsView.UpdateStatus(channelID, types.STATUS_CLOSED); err != nil {
				return fmt.Errorf("error updating channel closed: %w", err)
			}

			if err := ibcChannelsView.UpdateLastActivityTimeAndHeight(channelID, blockTime, height); err != nil {
				return fmt.Errorf("error updating channel last_activity_time: %w", err)
			}

		}
	}

	if err := UpdateLastHandledEventHeight(projection, rdbTxHandle, height); err != nil {
		return fmt.Errorf("error updating last handled event height: %v", err)
	}

	if err := rdbTx.Commit(); err != nil {
		return fmt.Errorf("error committing changes: %v", err)
	}
	committed = true

	return nil
}

func (projection *IBCChannel) updateBondedTokensWhenMsgIBCRecvPacket(
	ibcChannelsView ibc_channel_view.IBCChannels,
	ibcDenomHashMappingView ibc_channel_view.IBCDenomHashMapping,
	channelID string,
	amount string,
	denom string,
	destinationChannelID string,
	destinationPortID string,
	sourceChannelID string,
	sourcePortID string,
) error {

	bondedTokens, err := ibcChannelsView.FindBondedTokensBy(channelID)
	if err != nil {
		return fmt.Errorf("error finding channel bonded_tokens: %v", err)
	}

	amountInCoinInt, amountInCoinIntOk := coin.NewIntFromString(amount)
	if !amountInCoinIntOk {
		return fmt.Errorf("error creating coin from string: %s", amount)
	}

	if receiverChainIsTokenSource(denom, sourceChannelID, sourcePortID) {
		// This chain is token source, it is unbonded now.
		// Subtract it from the bondedTokens.OnCounterpartyChain
		token := ibc_channel_view.NewBondedToken(denom, amountInCoinInt)
		if subtractErr := subtractTokenOnCounterpartyChain(bondedTokens, token); subtractErr != nil {
			projection.logger.Warnf("error subtractTokenOnCounterpartyChain: %v", subtractErr)
		}
	} else {
		// Counterparty chain is token source, it is bonded to this chain now.
		// Add it to bondedTokens.OnThisChain
		newDenom := fmt.Sprintf("%s/%s/%s", destinationPortID, destinationChannelID, denom)
		token := ibc_channel_view.NewBondedToken(newDenom, amountInCoinInt)
		addTokenOnThisChain(bondedTokens, token)

		// For keeping record of denom-hash, we only interested in tokens sending to our chain
		if insertErr := insertDenomHashMappingIfNotExist(ibcDenomHashMappingView, newDenom); insertErr != nil {
			return fmt.Errorf("error insertDenomHashMappingIfNotExist: %v", insertErr)
		}
	}

	if err := ibcChannelsView.UpdateBondedTokens(channelID, bondedTokens); err != nil {
		return fmt.Errorf("error ibcChannelsView.UpdateBondedTokens: %v", err)
	}

	return nil
}

func (projection *IBCChannel) updateBondedTokensWhenMsgIBCTransfer(
	ibcChannelsView ibc_channel_view.IBCChannels,
	channelID string,
	amount string,
	denom string,
	destinationChannelID string,
	destinationPortID string,
	sourceChannelID string,
	sourcePortID string,
) error {

	bondedTokens, err := ibcChannelsView.FindBondedTokensBy(channelID)
	if err != nil {
		return fmt.Errorf("error ibcChannelsView.FindBondedTokensBy: %v", err)
	}

	amountInCoinInt, amountInCoinIntOk := coin.NewIntFromString(amount)
	if !amountInCoinIntOk {
		return fmt.Errorf("error creating coin from sting: %s", amount)
	}

	if receiverChainIsTokenSource(denom, sourceChannelID, sourcePortID) {
		// Counterparty chain is token source, it is unbonded now.
		// Subtract it from the bondedTokens.OnThisChain
		token := ibc_channel_view.NewBondedToken(denom, amountInCoinInt)
		if subtractErr := subtractTokenOnThisChain(bondedTokens, token); subtractErr != nil {
			projection.logger.Warnf("error subtractTokenOnThisChain: %v", subtractErr)
		}
	} else {
		// This chain is token source, it is bonded to counterparty chain now.
		// Add it to bondedTokens.OnCounterpartyChain
		newDenom := fmt.Sprintf("%s/%s/%s", destinationPortID, destinationChannelID, denom)
		token := ibc_channel_view.NewBondedToken(newDenom, amountInCoinInt)
		addTokenOnCounterpartyChain(bondedTokens, token)
	}

	if err := ibcChannelsView.UpdateBondedTokens(channelID, bondedTokens); err != nil {
		return fmt.Errorf("error ibcChannelsView.UpdateBondedTokens: %v", err)
	}

	return nil
}

func revertUpdateBondedTokensWhenMsgIBCTransfer(
	ibcChannelsView ibc_channel_view.IBCChannels,
	channelID string,
	amount string,
	denom string,
	destinationChannelID string,
	destinationPortID string,
	sourceChannelID string,
	sourcePortID string,
) error {

	bondedTokens, err := ibcChannelsView.FindBondedTokensBy(channelID)
	if err != nil {
		return fmt.Errorf("error ibcChannelsView.FindBondedTokensBy: %v", err)
	}

	amountInCoinInt, amountInCoinIntOk := coin.NewIntFromString(amount)
	if !amountInCoinIntOk {
		return fmt.Errorf("error creating coin from amount: %s", amount)
	}

	// Revert the operation in updateBondedTokensWhenMsgIBCTransfer()
	if receiverChainIsTokenSource(denom, sourceChannelID, sourcePortID) {
		token := ibc_channel_view.NewBondedToken(denom, amountInCoinInt)
		addTokenOnThisChain(bondedTokens, token)
	} else {
		newDenom := fmt.Sprintf("%s/%s/%s", destinationPortID, destinationChannelID, denom)
		token := ibc_channel_view.NewBondedToken(newDenom, amountInCoinInt)
		if subtractErr := subtractTokenOnCounterpartyChain(bondedTokens, token); subtractErr != nil {
			return fmt.Errorf("error subtractTokenOnCounterpartyChain: %v", subtractErr)
		}
	}

	if err := ibcChannelsView.UpdateBondedTokens(channelID, bondedTokens); err != nil {
		return fmt.Errorf("error ibcChannelsView.UpdateBondedTokens: %v", err)
	}

	return nil
}

func receiverChainIsTokenSource(denom, packetSourceChannelID, packetSourcePortID string) bool {
	denomPrefix := fmt.Sprintf("%s/%s", packetSourcePortID, packetSourceChannelID)
	return strings.HasPrefix(denom, denomPrefix)
}

func subtractTokenOnThisChain(
	bondedTokens *ibc_channel_view.BondedTokens,
	newToken *ibc_channel_view.BondedToken,
) error {
	for i, token := range bondedTokens.OnThisChain {
		if token.Denom == newToken.Denom {
			bondedTokens.OnThisChain[i].Amount = bondedTokens.OnThisChain[i].Amount.Sub(newToken.Amount)
			if bondedTokens.OnThisChain[i].Amount.IsZero() {
				bondedTokens.OnThisChain = append(bondedTokens.OnThisChain[:i], bondedTokens.OnThisChain[i+1:]...)
			}
			return nil
		}
	}
	return fmt.Errorf("denom %s is not found on bondedTokens.OnThisChain", newToken.Denom)
}

func subtractTokenOnCounterpartyChain(
	bondedTokens *ibc_channel_view.BondedTokens,
	newToken *ibc_channel_view.BondedToken,
) error {
	for i, token := range bondedTokens.OnCounterpartyChain {
		if token.Denom == newToken.Denom {
			bondedTokens.OnCounterpartyChain[i].Amount = bondedTokens.OnCounterpartyChain[i].Amount.Sub(newToken.Amount)
			if bondedTokens.OnCounterpartyChain[i].Amount.IsZero() {
				bondedTokens.OnCounterpartyChain = append(bondedTokens.OnCounterpartyChain[:i], bondedTokens.OnCounterpartyChain[i+1:]...)
			}
			return nil
		}
	}

	return fmt.Errorf("denom %s is not found on bondedTokens.OnCounterpartyChain", newToken.Denom)
}

func addTokenOnCounterpartyChain(
	bondedTokens *ibc_channel_view.BondedTokens,
	newToken *ibc_channel_view.BondedToken,
) {
	for i, token := range bondedTokens.OnCounterpartyChain {
		if token.Denom == newToken.Denom {
			// This token already has a record on bondedTokens.OnCounterpartyChain
			bondedTokens.OnCounterpartyChain[i].Amount = bondedTokens.OnCounterpartyChain[i].Amount.Add(newToken.Amount)
			return
		}
	}
	// This token has NO record on bondedTokens.OnCounterpartyChain yet
	bondedTokens.OnCounterpartyChain = append(bondedTokens.OnCounterpartyChain, *newToken)
}

func addTokenOnThisChain(
	bondedTokens *ibc_channel_view.BondedTokens,
	newToken *ibc_channel_view.BondedToken,
) {
	for i, token := range bondedTokens.OnThisChain {
		if token.Denom == newToken.Denom {
			// This token already has a record on bondedTokens.OnThisChain
			bondedTokens.OnThisChain[i].Amount = bondedTokens.OnThisChain[i].Amount.Add(newToken.Amount)
			return
		}
	}
	// This token has NO record on bondedTokens.OnThisChain yet
	bondedTokens.OnThisChain = append(bondedTokens.OnThisChain, *newToken)
}

func insertDenomHashMappingIfNotExist(
	ibcDenomHashMappingView ibc_channel_view.IBCDenomHashMapping,
	denom string,
) error {

	exist, err := ibcDenomHashMappingView.IfDenomExist(denom)
	if err != nil {
		return fmt.Errorf("error invoking IfDenomExist(): %v", err)
	}

	// The record does not exist, insert one
	if !exist {
		sum := sha256.Sum256([]byte(denom))
		hashBase64 := hex.EncodeToString((sum[:]))

		return ibcDenomHashMappingView.Insert(&ibc_channel_view.IBCDenomHashMappingRow{
			Denom: denom,
			Hash:  strings.ToUpper(hashBase64),
		})
	}

	// The record exists, do nothing
	return nil
}

func getBondedTokensInJSON(
	ibcChannelsView ibc_channel_view.IBCChannels,
	channelID string,
) (string, error) {

	updatedBondedTokens, err := ibcChannelsView.FindBondedTokensBy(channelID)
	if err != nil {
		return "", fmt.Errorf("error finding channel bonded_tokens: %w", err)
	}
	return jsoniter.MarshalToString(updatedBondedTokens)

}
