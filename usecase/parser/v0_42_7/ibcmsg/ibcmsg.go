package ibcmsg

import (
	"fmt"
	"time"

	"github.com/crypto-com/chain-indexing/external/json"
	"github.com/mitchellh/mapstructure"

	"github.com/crypto-com/chain-indexing/entity/command"
	base64_internal "github.com/crypto-com/chain-indexing/internal/base64"
	"github.com/crypto-com/chain-indexing/internal/typeconv"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	"github.com/crypto-com/chain-indexing/usecase/parser/ibcmsg"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
	mapstructure_utils "github.com/crypto-com/chain-indexing/usecase/parser/utils/mapstructure"
)

func ParseMsgRecvPacket(
	parserParams utils.CosmosParserParams,
) []command.Command {
	var rawMsg ibc_model.RawMsgRecvPacket
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			mapstructure_utils.StringToDurationHookFunc(),
			mapstructure_utils.StringToByteSliceHookFunc(),
		),
		Result: &rawMsg,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		panic(fmt.Errorf("error creating RawMsgRecvPacket decoder: %v", decoderErr))
	}
	if err := decoder.Decode(parserParams.Msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgRecvPacket: %v", err))
	}

	if !ibcmsg.IsPacketMsgTransfer(rawMsg.Packet) {
		// unsupported application
		return []command.Command{}
	}

	// Transfer application, MsgTransfer
	var rawFungibleTokenPacketData ibc_model.FungibleTokenPacketData
	rawPacketData := base64_internal.MustDecodeString(rawMsg.Packet.Data)
	json.MustUnmarshal(rawPacketData, &rawFungibleTokenPacketData)

	if !parserParams.MsgCommonParams.TxSuccess {
		msgRecvPacketParams := ibc_model.MsgRecvPacketParams{
			RawMsgRecvPacket: rawMsg,

			Application: "transfer",
			MessageType: "MsgTransfer",
			MaybeFungibleTokenPacketData: &ibc_model.MsgRecvPacketFungibleTokenPacketData{
				FungibleTokenPacketData: rawFungibleTokenPacketData,
				Success:                 false,
			},
		}

		return []command.Command{command_usecase.NewCreateMsgIBCRecvPacket(
			parserParams.MsgCommonParams,

			msgRecvPacketParams,
		)}
	}

	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])

	recvPacketEvent := log.GetEventByType("recv_packet")
	if recvPacketEvent == nil {
		panic("missing `recv_packet` event in TxsResult log")
	}

	fungibleTokenPacketEvent := log.GetEventByType("fungible_token_packet")
	if fungibleTokenPacketEvent == nil {
		// Note: this could happen when the packet is already relayed.
		// https://github.com/cosmos/ibc-go/blob/760d15a3a55397678abe311b7f65203b2e8437d6/modules/core/04-channel/keeper/packet.go#L239
		// https://github.com/cosmos/ibc-go/blob/760d15a3a55397678abe311b7f65203b2e8437d6/modules/core/keeper/msg_server.go#L508

		msgAlreadyRelayedRecvPacketParams := ibc_model.MsgAlreadyRelayedRecvPacketParams{
			RawMsgRecvPacket: rawMsg,

			Application: "transfer",
			MessageType: "MsgTransfer",
			MaybeFungibleTokenPacketData: &ibc_model.MsgAlreadyRelayedRecvPacketFungibleTokenPacketData{
				FungibleTokenPacketData: rawFungibleTokenPacketData,
			},

			PacketSequence: typeconv.MustAtou64(recvPacketEvent.MustGetAttributeByKey("packet_sequence")),
		}

		return []command.Command{command_usecase.NewCreateMsgAlreadyRelayedIBCRecvPacket(
			parserParams.MsgCommonParams,

			msgAlreadyRelayedRecvPacketParams,
		)}
	}

	var maybeDenominationTrace *ibc_model.MsgRecvPacketFungibleTokenDenominationTrace
	denominationTraceEvent := log.GetEventByType("denomination_trace")
	if denominationTraceEvent != nil {
		maybeDenominationTrace = &ibc_model.MsgRecvPacketFungibleTokenDenominationTrace{
			Hash:  denominationTraceEvent.MustGetAttributeByKey("trace_hash"),
			Denom: denominationTraceEvent.MustGetAttributeByKey("denom"),
		}
	}

	writeAckEvent := log.GetEventByType("write_acknowledgement")
	if writeAckEvent == nil {
		panic("missing `write_acknowledgement` event in TxsResult log")
	}
	var packetAck ibc_model.MsgRecvPacketPacketAck
	json.MustUnmarshalFromString(writeAckEvent.MustGetAttributeByKey("packet_ack"), &packetAck)

	msgRecvPacketParams := ibc_model.MsgRecvPacketParams{
		RawMsgRecvPacket: rawMsg,

		Application: "transfer",
		MessageType: "MsgTransfer",
		MaybeFungibleTokenPacketData: &ibc_model.MsgRecvPacketFungibleTokenPacketData{
			FungibleTokenPacketData: rawFungibleTokenPacketData,
			Success:                 fungibleTokenPacketEvent.MustGetAttributeByKey("success") == "true",
			MaybeDenominationTrace:  maybeDenominationTrace,
		},

		PacketSequence:  typeconv.MustAtou64(recvPacketEvent.MustGetAttributeByKey("packet_sequence")),
		ChannelOrdering: recvPacketEvent.MustGetAttributeByKey("packet_channel_ordering"),
		ConnectionID:    recvPacketEvent.MustGetAttributeByKey("packet_connection"),
		PacketAck:       packetAck,
	}

	return []command.Command{command_usecase.NewCreateMsgIBCRecvPacket(
		parserParams.MsgCommonParams,

		msgRecvPacketParams,
	)}
}
