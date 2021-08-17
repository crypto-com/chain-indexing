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
	// The indexing server is polling from source chain
	//
	// When the source chain init a transfer,
	// - on success, the source chain will generate MSG_IBC_ACKNOWLEDGEMENT event
	//							 the destination chain will generate MSG_IBC_RECV_PACKET event
	// - on failure, the source chain will generate MSG_IBC_TIMEOUT or MSG_IBC_TIMEOUT_ON_CLOSE
	//
	// reference: https://hermes.informal.systems/commands/relaying/packets.html#packet-streaming
	//
	return []string{
		event_usecase.MSG_IBC_CHANNEL_OPEN_ACK_CREATED,
		event_usecase.MSG_IBC_CHANNEL_OPEN_CONFIRM_CREATED,
		event_usecase.MSG_IBC_ACKNOWLEDGEMENT_CREATED,
		event_usecase.MSG_IBC_TIMEOUT_CREATED,
		event_usecase.MSG_IBC_TIMEOUT_ON_CLOSE_CREATED,
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

	for _, event := range events {
		if msgIBCChannelOpenAck, ok := event.(*event_usecase.MsgIBCChannelOpenAck); ok {

			// Channel just created, packet count should be 0
			channel := &channel_view.ChannelRow{
				ChannelID:             msgIBCChannelOpenAck.Params.ChannelID,
				PortID:                msgIBCChannelOpenAck.Params.PortID,
				ConnectionID:          msgIBCChannelOpenAck.Params.ConnectionID,
				CounterpartyChannelID: msgIBCChannelOpenAck.Params.CounterpartyChannelID,
				CounterpartyPortID:    msgIBCChannelOpenAck.Params.CounterpartyPortID,
				SuccessPacketCount:    0,
				FailurePacketCount:    0,
			}
			if err := channelsView.Upsert(channel); err != nil {
				return fmt.Errorf("error upserting channel: %w", err)
			}

		} else if msgIBCChannelOpenConfirm, ok := event.(*event_usecase.MsgIBCChannelOpenConfirm); ok {

			// Channel just created, packet count should be 0
			channel := &channel_view.ChannelRow{
				ChannelID:             msgIBCChannelOpenConfirm.Params.ChannelID,
				PortID:                msgIBCChannelOpenConfirm.Params.PortID,
				ConnectionID:          msgIBCChannelOpenConfirm.Params.ConnectionID,
				CounterpartyChannelID: msgIBCChannelOpenConfirm.Params.CounterpartyChannelID,
				CounterpartyPortID:    msgIBCChannelOpenConfirm.Params.CounterpartyPortID,
				SuccessPacketCount:    0,
				FailurePacketCount:    0,
			}
			if err := channelsView.Upsert(channel); err != nil {
				return fmt.Errorf("error upserting channel: %w", err)
			}

		} else if msgIBCAcknowledgement, ok := event.(*event_usecase.MsgIBCAcknowledgement); ok {

			channelId := msgIBCAcknowledgement.Params.Packet.SourceChannel
			portId := msgIBCAcknowledgement.Params.Packet.SourcePort

			if err := channelsView.SuccessPacketCountIncrement(channelId, portId); err != nil {
				return fmt.Errorf("error increasing channel.SuccessPacketCount: %w", err)
			}

		} else if msgIBCTimeout, ok := event.(*event_usecase.MsgIBCTimeout); ok {

			channelId := msgIBCTimeout.Params.Packet.SourceChannel
			portId := msgIBCTimeout.Params.Packet.SourcePort

			if err := channelsView.FailurePacketCountIncrement(channelId, portId); err != nil {
				return fmt.Errorf("error increasing channel.FailurePacketCount: %w", err)
			}

		} else if msgIBCTimeoutOnClose, ok := event.(*event_usecase.MsgIBCTimeoutOnClose); ok {

			channelId := msgIBCTimeoutOnClose.Params.Packet.SourceChannel
			portId := msgIBCTimeoutOnClose.Params.Packet.SourcePort

			if err := channelsView.FailurePacketCountIncrement(channelId, portId); err != nil {
				return fmt.Errorf("error increasing channel.FailurePacketCount: %w", err)
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
