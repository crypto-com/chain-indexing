package ibcmsg

import (
	"fmt"
	"time"

	"github.com/crypto-com/chain-indexing/usecase/parser/utils"

	base64_internal "github.com/crypto-com/chain-indexing/internal/base64"
	"github.com/crypto-com/chain-indexing/internal/json"
	"github.com/mitchellh/mapstructure"

	"github.com/crypto-com/chain-indexing/entity/command"
	"github.com/crypto-com/chain-indexing/internal/typeconv"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	"github.com/crypto-com/chain-indexing/usecase/parser/ibcmsg"
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
			utils.StringToDurationHookFunc(),
			utils.StringToByteSliceHookFunc(),
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
			},
		}

		return []command.Command{command_usecase.NewCreateMsgIBCRecvPacket(
			parserParams.MsgCommonParams,

			msgRecvPacketParams,
		)}
	}

	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])

	fungibleTokenPacketEvent := log.GetEventByType("fungible_token_packet")
	if fungibleTokenPacketEvent == nil {
		panic("missing `fungible_token_packet` event in TxsResult log")
	}

	var maybeDenominationTrace *ibc_model.MsgRecvPacketFungibleTokenDenominationTrace
	denominationTraceEvent := log.GetEventByType("denomination_trace")
	if denominationTraceEvent != nil {
		maybeDenominationTrace = &ibc_model.MsgRecvPacketFungibleTokenDenominationTrace{
			Hash:  denominationTraceEvent.MustGetAttributeByKey("trace_hash"),
			Denom: denominationTraceEvent.MustGetAttributeByKey("denom"),
		}
	}

	recvPacketEvent := log.GetEventByType("recv_packet")
	if recvPacketEvent == nil {
		panic("missing `recv_packet` event in TxsResult log")
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
