package channel

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	entity_projection "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	channel_view "github.com/crypto-com/chain-indexing/projection/channel/view"
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

	channelsView := channel_view.NewChannels(rdbTxHandle)

	// NOTES: Why four channel open events are all needed?
	//
	// PacketOrdering info can only be found in MsgIBCChannelOpenInit and MsgIBCChannelOpenTry.
	//
	// Full information of ConnectionID, CounterpartyChannelID, and CounterpartyPortID
	// can only be found in MsgIBCChannelOpenAck and MsgIBCChannelOpenConfirm.

	for _, event := range events {
		if msgIBCChannelOpenInit, ok := event.(*event_usecase.MsgIBCChannelOpenInit); ok {

			channel := &channel_view.ChannelRow{
				ChannelID:                    msgIBCChannelOpenInit.Params.ChannelID,
				PortID:                       msgIBCChannelOpenInit.Params.PortID,
				ConnectionID:                 "",
				CounterpartyChannelID:        msgIBCChannelOpenInit.Params.Channel.Counterparty.ChannelID,
				CounterpartyPortID:           msgIBCChannelOpenInit.Params.Channel.Counterparty.PortID,
				PacketOrdering:               msgIBCChannelOpenInit.Params.Channel.Ordering,
				LastInPacketSequence:         0,
				LastOutPacketSequence:        0,
				TotalTransferInCount:         0,
				TotalTransferOutCount:        0,
				TotalTransferOutSuccessCount: 0,
				TotalTransferOutSuccessRate:  0,
			}
			if err := channelsView.Insert(channel); err != nil {
				return fmt.Errorf("error inserting channel: %w", err)
			}

		} else if msgIBCChannelOpenTry, ok := event.(*event_usecase.MsgIBCChannelOpenTry); ok {

			channel := &channel_view.ChannelRow{
				ChannelID:                    msgIBCChannelOpenTry.Params.ChannelID,
				PortID:                       msgIBCChannelOpenTry.Params.PortID,
				ConnectionID:                 msgIBCChannelOpenTry.Params.ConnectionID,
				CounterpartyChannelID:        msgIBCChannelOpenTry.Params.Channel.Counterparty.ChannelID,
				CounterpartyPortID:           msgIBCChannelOpenTry.Params.Channel.Counterparty.PortID,
				PacketOrdering:               msgIBCChannelOpenTry.Params.Channel.Ordering,
				LastInPacketSequence:         0,
				LastOutPacketSequence:        0,
				TotalTransferInCount:         0,
				TotalTransferOutCount:        0,
				TotalTransferOutSuccessCount: 0,
				TotalTransferOutSuccessRate:  0,
			}
			if err := channelsView.Insert(channel); err != nil {
				return fmt.Errorf("error inserting channel: %w", err)
			}

		} else if msgIBCChannelOpenAck, ok := event.(*event_usecase.MsgIBCChannelOpenAck); ok {

			channel := &channel_view.ChannelRow{
				ChannelID:                    msgIBCChannelOpenAck.Params.ChannelID,
				PortID:                       msgIBCChannelOpenAck.Params.PortID,
				ConnectionID:                 msgIBCChannelOpenAck.Params.ConnectionID,
				CounterpartyChannelID:        msgIBCChannelOpenAck.Params.CounterpartyChannelID,
				CounterpartyPortID:           msgIBCChannelOpenAck.Params.CounterpartyPortID,
				LastInPacketSequence:         0,
				LastOutPacketSequence:        0,
				TotalTransferInCount:         0,
				TotalTransferOutCount:        0,
				TotalTransferOutSuccessCount: 0,
				TotalTransferOutSuccessRate:  0,
			}
			if err := channelsView.Update(channel); err != nil {
				return fmt.Errorf("error updating channel: %w", err)
			}

		} else if msgIBCChannelOpenConfirm, ok := event.(*event_usecase.MsgIBCChannelOpenConfirm); ok {

			channel := &channel_view.ChannelRow{
				ChannelID:                    msgIBCChannelOpenConfirm.Params.ChannelID,
				PortID:                       msgIBCChannelOpenConfirm.Params.PortID,
				ConnectionID:                 msgIBCChannelOpenConfirm.Params.ConnectionID,
				CounterpartyChannelID:        msgIBCChannelOpenConfirm.Params.CounterpartyChannelID,
				CounterpartyPortID:           msgIBCChannelOpenConfirm.Params.CounterpartyPortID,
				LastInPacketSequence:         0,
				LastOutPacketSequence:        0,
				TotalTransferInCount:         0,
				TotalTransferOutCount:        0,
				TotalTransferOutSuccessCount: 0,
				TotalTransferOutSuccessRate:  0,
			}
			if err := channelsView.Update(channel); err != nil {
				return fmt.Errorf("error updating channel: %w", err)
			}

		} else if msgIBCTransferTransfer, ok := event.(*event_usecase.MsgIBCTransferTransfer); ok {

			// Transfer started by source chain
			channelId := msgIBCTransferTransfer.Params.SourceChannel
			portId := msgIBCTransferTransfer.Params.SourcePort

			// TotalTransferOutSuccessRate
			if err := channelsView.Increment(channelId, portId, "total_transfer_out_count", 1); err != nil {
				return fmt.Errorf("error increasing total_transfer_out_count: %w", err)
			}
			if err := channelsView.UpdateTotalTransferOutSuccessRate(channelId, portId); err != nil {
				return fmt.Errorf("error updating total_transfer_out_success_rate: %w", err)
			}

			lastOutPacketSequence := msgIBCTransferTransfer.Params.PacketSequence
			if err := channelsView.UpdateSequence(channelId, portId, "last_out_packet_sequence", lastOutPacketSequence); err != nil {
				return fmt.Errorf("error updating last_out_packet_sequence: %w", err)
			}

		} else if msgIBCRecvPacket, ok := event.(*event_usecase.MsgIBCRecvPacket); ok {

			// Transfer started by destination chain
			channelId := msgIBCRecvPacket.Params.Packet.DestinationChannel
			portId := msgIBCRecvPacket.Params.Packet.DestinationPort

			if err := channelsView.Increment(channelId, portId, "total_transfer_in_count", 1); err != nil {
				return fmt.Errorf("error increasing total_transfer_in_count: %w", err)
			}

			lastInPacketSequence := msgIBCRecvPacket.Params.PacketSequence
			if err := channelsView.UpdateSequence(channelId, portId, "last_in_packet_sequence", lastInPacketSequence); err != nil {
				return fmt.Errorf("error updating last_in_packet_sequence: %w", err)
			}

		} else if msgIBCAcknowledgement, ok := event.(*event_usecase.MsgIBCAcknowledgement); ok {

			// Transfer started by source chain
			channelId := msgIBCAcknowledgement.Params.Packet.SourceChannel
			portId := msgIBCAcknowledgement.Params.Packet.SourcePort

			// TotalTransferOutSuccessRate
			if err := channelsView.Increment(channelId, portId, "total_transfer_out_success_count", 1); err != nil {
				return fmt.Errorf("error increasing total_transfer_out_success_count: %w", err)
			}
			if err := channelsView.UpdateTotalTransferOutSuccessRate(channelId, portId); err != nil {
				return fmt.Errorf("error updating total_transfer_out_success_rate: %w", err)
			}

			lastOutPacketSequence := msgIBCAcknowledgement.Params.PacketSequence
			if err := channelsView.UpdateSequence(channelId, portId, "last_out_packet_sequence", lastOutPacketSequence); err != nil {
				return fmt.Errorf("error updating last_out_packet_sequence: %w", err)
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
