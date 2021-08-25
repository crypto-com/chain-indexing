package channel

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	entity_projection "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	"github.com/crypto-com/chain-indexing/internal/utctime"
	block_view "github.com/crypto-com/chain-indexing/projection/block/view"
	channel_view "github.com/crypto-com/chain-indexing/projection/channel/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ entity_projection.Projection = &Channel{}

type Channel struct {
	*rdbprojectionbase.Base

	rdbConn rdb.Conn
	logger  applogger.Logger
}

func NewChannel(logger applogger.Logger, rdbConn rdb.Conn) *Channel {
	return &Channel{
		rdbprojectionbase.NewRDbBase(rdbConn.ToHandle(), "Channel"),

		rdbConn,
		logger,
	}
}

func (_ *Channel) GetEventsToListen() []string {
	return []string{
		event_usecase.MSG_IBC_CREATE_CLIENT_CREATED,
		event_usecase.MSG_IBC_CONNECTION_OPEN_ACK_CREATED,
		event_usecase.MSG_IBC_CONNECTION_OPEN_CONFIRM_CREATED,
		event_usecase.MSG_IBC_CHANNEL_OPEN_INIT_CREATED,
		event_usecase.MSG_IBC_CHANNEL_OPEN_TRY_CREATED,
		event_usecase.MSG_IBC_CHANNEL_OPEN_ACK_CREATED,
		event_usecase.MSG_IBC_CHANNEL_OPEN_CONFIRM_CREATED,
		event_usecase.MSG_IBC_TRANSFER_TRANSFER_CREATED,
		event_usecase.MSG_IBC_RECV_PACKET_CREATED,
		event_usecase.MSG_IBC_ACKNOWLEDGEMENT_CREATED,
	}
}

func (projection *Channel) OnInit() error {
	return nil
}

func (projection *Channel) HandleEvents(height int64, events []event_entity.Event) error {
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

	ibcChannelsView := channel_view.NewIBCChannels(rdbTxHandle)
	ibcClientsView := channel_view.NewIBCClients(rdbTxHandle)
	ibcConnectionsView := channel_view.NewIBCConnections(rdbTxHandle)
	blocksView := block_view.NewBlocks(rdbTxHandle)

	// NOTES: Why four channel open events are all needed?
	//
	// PacketOrdering info can only be found in MsgIBCChannelOpenInit and MsgIBCChannelOpenTry.
	//
	// Full information of ConnectionID, CounterpartyChannelID, and CounterpartyPortID
	// can only be found in MsgIBCChannelOpenAck and MsgIBCChannelOpenConfirm.

	for _, event := range events {
		if msgIBCCreateClient, ok := event.(*event_usecase.MsgIBCCreateClient); ok {

			client := &channel_view.IBCClientRow{
				ClientID:            msgIBCCreateClient.Params.ClientID,
				CounterpartyChainID: msgIBCCreateClient.Params.MaybeTendermintLightClient.TendermintClientState.ChainID,
			}
			if err := ibcClientsView.Insert(client); err != nil {
				return fmt.Errorf("error inserting client: %w", err)
			}

		} else if msgIBCConnectionOpenConfirm, ok := event.(*event_usecase.MsgIBCConnectionOpenConfirm); ok {

			clientID := msgIBCConnectionOpenConfirm.Params.ClientID
			counterpartyChainID, err := ibcClientsView.FindCounterpartyChainIDBy(clientID)
			if err != nil {
				return fmt.Errorf("error in finding counterparty_chain_id: %w", err)
			}

			connection := &channel_view.IBCConnectionRow{
				ConnectionID:             msgIBCConnectionOpenConfirm.Params.ConnectionID,
				ClientID:                 msgIBCConnectionOpenConfirm.Params.ClientID,
				CounterpartyConnectionID: msgIBCConnectionOpenConfirm.Params.CounterpartyConnectionID,
				CounterpartyClientID:     msgIBCConnectionOpenConfirm.Params.CounterpartyClientID,
				CounterpartyChainID:      counterpartyChainID,
			}
			if err := ibcConnectionsView.Insert(connection); err != nil {
				return fmt.Errorf("error inserting connection: %w", err)
			}

		} else if msgIBCConnectionOpenAck, ok := event.(*event_usecase.MsgIBCConnectionOpenAck); ok {

			clientID := msgIBCConnectionOpenAck.Params.ClientID
			counterpartyChainID, err := ibcClientsView.FindCounterpartyChainIDBy(clientID)
			if err != nil {
				return fmt.Errorf("error in finding counterparty_chain_id: %w", err)
			}

			connection := &channel_view.IBCConnectionRow{
				ConnectionID:             msgIBCConnectionOpenAck.Params.ConnectionID,
				ClientID:                 msgIBCConnectionOpenAck.Params.ClientID,
				CounterpartyConnectionID: msgIBCConnectionOpenAck.Params.CounterpartyConnectionID,
				CounterpartyClientID:     msgIBCConnectionOpenAck.Params.CounterpartyClientID,
				CounterpartyChainID:      counterpartyChainID,
			}
			if err := ibcConnectionsView.Insert(connection); err != nil {
				return fmt.Errorf("error inserting connection: %w", err)
			}

		} else if msgIBCChannelOpenInit, ok := event.(*event_usecase.MsgIBCChannelOpenInit); ok {

			channel := &channel_view.ChannelRow{
				ChannelID:                    msgIBCChannelOpenInit.Params.ChannelID,
				PortID:                       msgIBCChannelOpenInit.Params.PortID,
				ConnectionID:                 "",
				CounterpartyChannelID:        msgIBCChannelOpenInit.Params.Channel.Counterparty.ChannelID,
				CounterpartyPortID:           msgIBCChannelOpenInit.Params.Channel.Counterparty.PortID,
				CounterpartyChainID:          "",
				Status:                       false,
				PacketOrdering:               msgIBCChannelOpenInit.Params.Channel.Ordering,
				LastInPacketSequence:         0,
				LastOutPacketSequence:        0,
				TotalTransferInCount:         0,
				TotalTransferOutCount:        0,
				TotalTransferOutSuccessCount: 0,
				TotalTransferOutSuccessRate:  0,
				CreatedAtBlockTime:           utctime.FromUnixNano(0),
				CreatedAtBlockHeight:         0,
				Verified:                     false,
				Description:                  "",
				LastActivityBlockTime:        utctime.FromUnixNano(0),
				LastActivityBlockHeight:      0,
				BondedTokens:                 *channel_view.NewEmptyBondedTokens(),
			}
			if err := ibcChannelsView.Insert(channel); err != nil {
				return fmt.Errorf("error inserting channel: %w", err)
			}

		} else if msgIBCChannelOpenTry, ok := event.(*event_usecase.MsgIBCChannelOpenTry); ok {

			channel := &channel_view.ChannelRow{
				ChannelID:                    msgIBCChannelOpenTry.Params.ChannelID,
				PortID:                       msgIBCChannelOpenTry.Params.PortID,
				ConnectionID:                 msgIBCChannelOpenTry.Params.ConnectionID,
				CounterpartyChannelID:        msgIBCChannelOpenTry.Params.Channel.Counterparty.ChannelID,
				CounterpartyPortID:           msgIBCChannelOpenTry.Params.Channel.Counterparty.PortID,
				CounterpartyChainID:          "",
				Status:                       false,
				PacketOrdering:               msgIBCChannelOpenTry.Params.Channel.Ordering,
				LastInPacketSequence:         0,
				LastOutPacketSequence:        0,
				TotalTransferInCount:         0,
				TotalTransferOutCount:        0,
				TotalTransferOutSuccessCount: 0,
				TotalTransferOutSuccessRate:  0,
				CreatedAtBlockTime:           utctime.FromUnixNano(0),
				CreatedAtBlockHeight:         0,
				Verified:                     false,
				Description:                  "",
				LastActivityBlockTime:        utctime.FromUnixNano(0),
				LastActivityBlockHeight:      0,
				BondedTokens:                 *channel_view.NewEmptyBondedTokens(),
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

			identity := block_view.BlockIdentity{MaybeHeight: &height}
			block, err := blocksView.FindBy(&identity)
			if err != nil {
				return fmt.Errorf("error finding the block: %w", err)
			}

			channel := &channel_view.ChannelRow{
				ChannelID:             msgIBCChannelOpenAck.Params.ChannelID,
				PortID:                msgIBCChannelOpenAck.Params.PortID,
				ConnectionID:          msgIBCChannelOpenAck.Params.ConnectionID,
				CounterpartyChannelID: msgIBCChannelOpenAck.Params.CounterpartyChannelID,
				CounterpartyPortID:    msgIBCChannelOpenAck.Params.CounterpartyPortID,
				CounterpartyChainID:   counterpartyChainID,
				CreatedAtBlockTime:    block.Time,
				CreatedAtBlockHeight:  height,
			}
			if err := ibcChannelsView.UpdateFactualColumns(channel); err != nil {
				return fmt.Errorf("error updating channel: %w", err)
			}

			if err := ibcChannelsView.UpdateStatus(channel.ChannelID, true); err != nil {
				return fmt.Errorf("error updating channel status: %w", err)
			}

		} else if msgIBCChannelOpenConfirm, ok := event.(*event_usecase.MsgIBCChannelOpenConfirm); ok {

			connectionID := msgIBCChannelOpenConfirm.Params.ConnectionID
			counterpartyChainID, err := ibcConnectionsView.FindCounterpartyChainIDBy(connectionID)
			if err != nil {
				return fmt.Errorf("error in finding counterparty_chain_id: %w", err)
			}

			identity := block_view.BlockIdentity{MaybeHeight: &height}
			block, err := blocksView.FindBy(&identity)
			if err != nil {
				return fmt.Errorf("error finding the block: %w", err)
			}

			channel := &channel_view.ChannelRow{
				ChannelID:             msgIBCChannelOpenConfirm.Params.ChannelID,
				PortID:                msgIBCChannelOpenConfirm.Params.PortID,
				ConnectionID:          msgIBCChannelOpenConfirm.Params.ConnectionID,
				CounterpartyChannelID: msgIBCChannelOpenConfirm.Params.CounterpartyChannelID,
				CounterpartyPortID:    msgIBCChannelOpenConfirm.Params.CounterpartyPortID,
				CounterpartyChainID:   counterpartyChainID,
				CreatedAtBlockTime:    block.Time,
				CreatedAtBlockHeight:  height,
			}
			if err := ibcChannelsView.UpdateFactualColumns(channel); err != nil {
				return fmt.Errorf("error updating channel: %w", err)
			}

			if err := ibcChannelsView.UpdateStatus(channel.ChannelID, true); err != nil {
				return fmt.Errorf("error updating channel status: %w", err)
			}

		} else if msgIBCTransferTransfer, ok := event.(*event_usecase.MsgIBCTransferTransfer); ok {

			// Transfer started by source chain
			channelID := msgIBCTransferTransfer.Params.SourceChannel

			// TotalTransferOutSuccessRate
			if err := ibcChannelsView.Increment(channelID, "total_transfer_out_count", 1); err != nil {
				return fmt.Errorf("error increasing total_transfer_out_count: %w", err)
			}
			if err := ibcChannelsView.UpdateTotalTransferOutSuccessRate(channelID); err != nil {
				return fmt.Errorf("error updating total_transfer_out_success_rate: %w", err)
			}

			lastOutPacketSequence := msgIBCTransferTransfer.Params.PacketSequence
			if err := ibcChannelsView.UpdateSequence(channelID, "last_out_packet_sequence", lastOutPacketSequence); err != nil {
				return fmt.Errorf("error updating last_out_packet_sequence: %w", err)
			}

			if err := projection.updateLastActivityTimeAndHeight(blocksView, ibcChannelsView, channelID, height); err != nil {
				return fmt.Errorf("error updating channel last_activity_time: %w", err)
			}

		} else if msgIBCRecvPacket, ok := event.(*event_usecase.MsgIBCRecvPacket); ok {

			// Transfer started by destination chain
			channelID := msgIBCRecvPacket.Params.Packet.DestinationChannel

			if err := ibcChannelsView.Increment(channelID, "total_transfer_in_count", 1); err != nil {
				return fmt.Errorf("error increasing total_transfer_in_count: %w", err)
			}

			lastInPacketSequence := msgIBCRecvPacket.Params.PacketSequence
			if err := ibcChannelsView.UpdateSequence(channelID, "last_in_packet_sequence", lastInPacketSequence); err != nil {
				return fmt.Errorf("error updating last_in_packet_sequence: %w", err)
			}

			if err := projection.updateLastActivityTimeAndHeight(blocksView, ibcChannelsView, channelID, height); err != nil {
				return fmt.Errorf("error updating channel last_activity_time: %w", err)
			}

			amount := msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData.Amount
			denom := msgIBCRecvPacket.Params.MaybeFungibleTokenPacketData.Denom
			destinationChannelID := msgIBCRecvPacket.Params.Packet.DestinationChannel
			destinationPortID := msgIBCRecvPacket.Params.Packet.DestinationPort

			if err := projection.updateBondedTokensWhenMsgIBCRecvPacket(
				ibcChannelsView,
				channelID,
				amount,
				denom,
				destinationChannelID,
				destinationPortID,
			); err != nil {
				return fmt.Errorf("error updateBondedTokensWhenMsgIBCRecvPacket: %w", err)
			}

		} else if msgIBCAcknowledgement, ok := event.(*event_usecase.MsgIBCAcknowledgement); ok {

			// Transfer started by source chain
			channelID := msgIBCAcknowledgement.Params.Packet.SourceChannel

			// TotalTransferOutSuccessRate
			if err := ibcChannelsView.Increment(channelID, "total_transfer_out_success_count", 1); err != nil {
				return fmt.Errorf("error increasing total_transfer_out_success_count: %w", err)
			}
			if err := ibcChannelsView.UpdateTotalTransferOutSuccessRate(channelID); err != nil {
				return fmt.Errorf("error updating total_transfer_out_success_rate: %w", err)
			}

			lastOutPacketSequence := msgIBCAcknowledgement.Params.PacketSequence
			if err := ibcChannelsView.UpdateSequence(channelID, "last_out_packet_sequence", lastOutPacketSequence); err != nil {
				return fmt.Errorf("error updating last_out_packet_sequence: %w", err)
			}

			if err := projection.updateLastActivityTimeAndHeight(blocksView, ibcChannelsView, channelID, height); err != nil {
				return fmt.Errorf("error updating channel last_activity_time: %w", err)
			}

			amount := msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.Amount
			denom := msgIBCAcknowledgement.Params.MaybeFungibleTokenPacketData.Denom
			destinationChannelID := msgIBCAcknowledgement.Params.Packet.DestinationChannel
			destinationPortID := msgIBCAcknowledgement.Params.Packet.DestinationPort

			if err := projection.updateBondedTokensWhenMsgIBCAcknowledgement(
				ibcChannelsView,
				channelID,
				amount,
				denom,
				destinationChannelID,
				destinationPortID,
			); err != nil {
				return fmt.Errorf("error updateBondedTokensWhenMsgIBCAcknowledgement: %w", err)
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

func (projection *Channel) updateLastActivityTimeAndHeight(
	blocksView *block_view.Blocks,
	ibcChannelsView *channel_view.IBCChannels,
	channelID string,
	height int64,
) error {

	identity := block_view.BlockIdentity{MaybeHeight: &height}
	block, err := blocksView.FindBy(&identity)
	if err != nil {
		return fmt.Errorf("error finding the block: %w", err)
	}
	if err := ibcChannelsView.UpdateLastActivityTimeAndHeight(channelID, block.Time, height); err != nil {
		return fmt.Errorf("error updating channel last_activity_time: %w", err)
	}

	return nil
}

func (projection *Channel) updateBondedTokensWhenMsgIBCRecvPacket(
	ibcChannelsView *channel_view.IBCChannels,
	channelID string,
	amount uint64,
	denom string,
	destinationChannelID string,
	destinationPortID string,
) error {

	bondedTokens, err := ibcChannelsView.FindBondedTokensBy(channelID)
	if err != nil {
		return fmt.Errorf("error finding channel bonded_tokens: %w", err)
	}

	// Check if the token is originally from this chain
	tokenOriginallyFromThisChain := false
	tokenIndex := -1
	for i, token := range bondedTokens.OnCounterpartyChain {
		if token.Denom == denom {
			tokenOriginallyFromThisChain = true
			tokenIndex = i
			break
		}
	}
	if tokenOriginallyFromThisChain {
		// This token is originally from this chain, now it is sent back from counterparty chain.
		// Subtract it from the bondedTokens.OnCouterpartyChain[i]
		amountInCoinInt := coin.NewIntFromUint64(amount)
		bondedTokens.OnCounterpartyChain[tokenIndex].Amount = bondedTokens.OnCounterpartyChain[tokenIndex].Amount.Sub(amountInCoinInt)
	} else {
		// This token is new to this chain.
		// Append it to bondedTokens.OnThisChain
		newDenom := fmt.Sprintf("%s/%s/%s", destinationPortID, destinationChannelID, denom)
		amountInCoinInt := coin.NewIntFromUint64(amount)
		newToken := channel_view.NewBondedToken(newDenom, amountInCoinInt)
		bondedTokens.OnThisChain = append(bondedTokens.OnThisChain, *newToken)
	}

	if err := ibcChannelsView.UpdateBondedTokens(channelID, bondedTokens); err != nil {
		return fmt.Errorf("error ibcChannelsView.UpdateBondedTokens: %w", err)
	}

	return nil
}

func (projection *Channel) updateBondedTokensWhenMsgIBCAcknowledgement(
	ibcChannelsView *channel_view.IBCChannels,
	channelID string,
	amount uint64,
	denom string,
	destinationChannelID string,
	destinationPortID string,
) error {

	bondedTokens, err := ibcChannelsView.FindBondedTokensBy(channelID)
	if err != nil {
		return fmt.Errorf("error ibcChannelsView.FindBondedTokensBy: %w", err)
	}

	// Check if the token is originally from counterparty chain
	tokenOriginallyFromCounterpartyChain := false
	tokenIndex := -1
	for i, token := range bondedTokens.OnThisChain {
		if token.Denom == denom {
			tokenOriginallyFromCounterpartyChain = true
			tokenIndex = i
			break
		}
	}
	if tokenOriginallyFromCounterpartyChain {
		// This token is from counterparty chain, now it is sent back to counterparty chain.
		// Subtract it from the bondedTokens.OnThisChain[i]
		amountInCoinInt := coin.NewIntFromUint64(amount)
		bondedTokens.OnThisChain[tokenIndex].Amount = bondedTokens.OnThisChain[tokenIndex].Amount.Sub(amountInCoinInt)
	} else {
		// This token is new to the counterparty chain.
		// Append it to bondedTokens.OnCounterpartyChain
		newDenom := fmt.Sprintf("%s/%s/%s", destinationPortID, destinationChannelID, denom)
		amountInCoinInt := coin.NewIntFromUint64(amount)
		newToken := channel_view.NewBondedToken(newDenom, amountInCoinInt)
		bondedTokens.OnCounterpartyChain = append(bondedTokens.OnCounterpartyChain, *newToken)
	}

	if err := ibcChannelsView.UpdateBondedTokens(channelID, bondedTokens); err != nil {
		return fmt.Errorf("error ibcChannelsView.UpdateBondedTokens: %w", err)
	}

	return nil
}
