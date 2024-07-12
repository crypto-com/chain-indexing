package ibcmsg

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"time"

	"github.com/crypto-com/chain-indexing/external/json"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"

	"github.com/crypto-com/chain-indexing/entity/command"
	base64_internal "github.com/crypto-com/chain-indexing/internal/base64"
	"github.com/crypto-com/chain-indexing/internal/typeconv"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	"github.com/crypto-com/chain-indexing/usecase/parser/ibc"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
	mapstructure_utils "github.com/crypto-com/chain-indexing/usecase/parser/utils/mapstructure"
)

func ParseMsgRecvPacket(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
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

	var cmds []command.Command
	var possibleSignerAddresses []string

	if !parserParams.MsgCommonParams.TxSuccess {
		return []command.Command{}, []string{}
	}

	if ibc.IsPacketMsgTransfer(rawMsg.Packet, &parserParams.TxsResult.Log[parserParams.MsgIndex]) {
		// Parse IBC transfer
		cmds, possibleSignerAddresses = ParseMsgTransfer(parserParams, rawMsg)

	} else if ibc.IsPacketMsgSubmitTx(rawMsg.Packet) {
		// Parse host chain MsgSubmitTx
		cmds, possibleSignerAddresses = ParseMsgSubmitTx(parserParams, rawMsg)

	} else {
		// unsupported application
		return []command.Command{}, []string{}
	}

	return cmds, possibleSignerAddresses
}

func ParseMsgTransfer(
	parserParams utils.CosmosParserParams,
	rawMsg ibc_model.RawMsgRecvPacket,
) ([]command.Command, []string) {
	// Transfer application, MsgTransfer
	var rawFungibleTokenPacketData ibc_model.FungibleTokenPacketData
	rawPacketData, err := base64_internal.DecodeString(rawMsg.Packet.Data)
	if err != nil {
		rawFungibleTokenPacketData = ibc_model.FungibleTokenPacketData{}
	} else {
		if err = json.Unmarshal(rawPacketData, &rawFungibleTokenPacketData); err != nil {
			rawFungibleTokenPacketData = ibc_model.FungibleTokenPacketData{}
		}
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		msgRecvPacketParams := ibc_model.MsgRecvPacketParams{
			RawMsgRecvPacket: rawMsg,

			Application: "transfer",
			MessageType: "/ibc.applications.transfer.v1.MsgTransfer",
			MaybeFungibleTokenPacketData: &ibc_model.MsgRecvPacketFungibleTokenPacketData{
				FungibleTokenPacketData: rawFungibleTokenPacketData,
				Success:                 false,
			},
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, msgRecvPacketParams.Signer)

		return []command.Command{command_usecase.NewCreateMsgIBCRecvPacket(
			parserParams.MsgCommonParams,

			msgRecvPacketParams,
		)}, possibleSignerAddresses
	}

	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])

	if !log.HasEvent("recv_packet") {
		msgRecvPacketParams := ibc_model.MsgRecvPacketParams{
			RawMsgRecvPacket: rawMsg,

			Application: "transfer",
			MessageType: "/ibc.applications.transfer.v1.MsgTransfer",
			MaybeFungibleTokenPacketData: &ibc_model.MsgRecvPacketFungibleTokenPacketData{
				FungibleTokenPacketData: rawFungibleTokenPacketData,
				Success:                 false,
			},
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, msgRecvPacketParams.Signer)

		return []command.Command{command_usecase.NewCreateMsgIBCRecvPacket(
			parserParams.MsgCommonParams,

			msgRecvPacketParams,
		)}, possibleSignerAddresses
	}

	recvPacketEvents := log.GetEventsByType("recv_packet")
	var packetSequence uint64
	var channelOrdering string
	var connectionID string
	for _, recvPacketEvent := range recvPacketEvents {
		if recvPacketEvent.HasAttribute("packet_sequence") {
			packetSequence = typeconv.MustAtou64(recvPacketEvent.MustGetAttributeByKey("packet_sequence"))
		}
		if recvPacketEvent.HasAttribute("packet_channel_ordering") {
			channelOrdering = recvPacketEvent.MustGetAttributeByKey("packet_channel_ordering")
		}
		if recvPacketEvent.HasAttribute("packet_connection") {
			connectionID = recvPacketEvent.MustGetAttributeByKey("packet_connection")
		}
	}

	fungibleTokenPacketEvents := log.GetEventsByType("fungible_token_packet")
	if fungibleTokenPacketEvents == nil {
		// Note: this could happen when the packet is already relayed.
		// https://github.com/cosmos/ibc-go/blob/760d15a3a55397678abe311b7f65203b2e8437d6/modules/core/04-channel/keeper/packet.go#L239
		// https://github.com/cosmos/ibc-go/blob/760d15a3a55397678abe311b7f65203b2e8437d6/modules/core/keeper/msg_server.go#L508

		msgAlreadyRelayedRecvPacketParams := ibc_model.MsgAlreadyRelayedRecvPacketParams{
			RawMsgRecvPacket: rawMsg,

			Application: "transfer",
			MessageType: "/ibc.applications.transfer.v1.MsgTransfer",
			MaybeFungibleTokenPacketData: &ibc_model.MsgAlreadyRelayedRecvPacketFungibleTokenPacketData{
				FungibleTokenPacketData: rawFungibleTokenPacketData,
			},

			PacketSequence: packetSequence,
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, msgAlreadyRelayedRecvPacketParams.Signer)

		return []command.Command{command_usecase.NewCreateMsgAlreadyRelayedIBCRecvPacket(
			parserParams.MsgCommonParams,

			msgAlreadyRelayedRecvPacketParams,
		)}, possibleSignerAddresses
	}
	var success bool
	for _, fungibleTokenPacketEvent := range fungibleTokenPacketEvents {
		if fungibleTokenPacketEvent.HasAttribute("success") {
			success = fungibleTokenPacketEvent.MustGetAttributeByKey("success") == "true"
		}
	}

	var maybeDenominationTrace *ibc_model.MsgRecvPacketFungibleTokenDenominationTrace
	denominationTraceEvents := log.GetEventsByType("denomination_trace")
	var denominationTraceEventTraceHash string
	var denominationTraceEventDenom string
	for _, denominationTraceEvent := range denominationTraceEvents {
		if denominationTraceEvent.HasAttribute("trace_hash") {
			denominationTraceEventTraceHash = denominationTraceEvent.MustGetAttributeByKey("trace_hash")
		}
		if denominationTraceEvent.HasAttribute("denom") {
			denominationTraceEventDenom = denominationTraceEvent.MustGetAttributeByKey("denom")
		}
	}
	if denominationTraceEvents != nil {
		maybeDenominationTrace = &ibc_model.MsgRecvPacketFungibleTokenDenominationTrace{
			Hash:  denominationTraceEventTraceHash,
			Denom: denominationTraceEventDenom,
		}
	}

	var packetAck ibc_model.MsgRecvPacketPacketAck
	writeAckEvents := log.GetEventsByType("write_acknowledgement")
	if writeAckEvents == nil {
		parserParams.Logger.Warnf("missing `write_acknowledgement` event in TxsResult log on TxHash: %s", parserParams.MsgCommonParams.TxHash)
	} else {
		var writeAckEventPacketAck string
		for _, writeAckEvent := range writeAckEvents {
			if writeAckEvent.HasAttribute("packet_ack") {
				writeAckEventPacketAck = writeAckEvent.MustGetAttributeByKey("packet_ack")
			}
		}

		if writeAckEventPacketAck != "" {
			json.MustUnmarshalFromString(writeAckEventPacketAck, &packetAck)
		}
	}

	msgRecvPacketParams := ibc_model.MsgRecvPacketParams{
		RawMsgRecvPacket: rawMsg,

		Application: "transfer",
		MessageType: "/ibc.applications.transfer.v1.MsgTransfer",
		MaybeFungibleTokenPacketData: &ibc_model.MsgRecvPacketFungibleTokenPacketData{
			FungibleTokenPacketData: rawFungibleTokenPacketData,
			Success:                 success,
			MaybeDenominationTrace:  maybeDenominationTrace,
		},

		PacketSequence:  packetSequence,
		ChannelOrdering: channelOrdering,
		ConnectionID:    connectionID,
		PacketAck:       packetAck,
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	possibleSignerAddresses = append(possibleSignerAddresses, msgRecvPacketParams.Signer)

	return []command.Command{command_usecase.NewCreateMsgIBCRecvPacket(
		parserParams.MsgCommonParams,

		msgRecvPacketParams,
	)}, possibleSignerAddresses
}

func ParseMsgSubmitTx(
	parserParams utils.CosmosParserParams,
	rawMsg ibc_model.RawMsgRecvPacket,
) ([]command.Command, []string) {
	// MsgSubmitTx
	var rawInterchainAccountPacketData ibc_model.InterchainAccountPacketData
	rawPacketData, err := base64.StdEncoding.DecodeString(rawMsg.Packet.Data)
	if err != nil {
		rawInterchainAccountPacketData = ibc_model.InterchainAccountPacketData{}
	}
	if unmarshalErr := jsoniter.Unmarshal(rawPacketData, &rawInterchainAccountPacketData); unmarshalErr != nil {
		rawInterchainAccountPacketData = ibc_model.InterchainAccountPacketData{}
	}

	cosmosTx, innerMsgs, deserializeErr := parserParams.TxDecoder.DeserializeCosmosTx(rawInterchainAccountPacketData.Data)
	if deserializeErr != nil {
		panic(fmt.Sprintf("error deserializing cosmos tx: %v", err))
	}

	var cmds []command.Command
	var possibleSignerAddresses []string

	if !parserParams.MsgCommonParams.TxSuccess {
		msgRecvPacketParams := ibc_model.MsgRecvPacketParams{
			RawMsgRecvPacket: rawMsg,

			Application: "icahost",
			MessageType: "/chainmain.icaauth.v1.MsgSubmitTx",
			MaybeFungibleTokenPacketData: &ibc_model.MsgRecvPacketFungibleTokenPacketData{
				Success: false,
			},
		}

		// Getting possible signer address from Msg
		possibleSignerAddresses = append(possibleSignerAddresses, msgRecvPacketParams.Signer)
		cmds = append(cmds, command_usecase.NewCreateMsgIBCRecvPacket(
			parserParams.MsgCommonParams,

			msgRecvPacketParams,
		))
		return cmds, possibleSignerAddresses
	}

	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])

	if !log.HasEvent("recv_packet") {
		msgRecvPacketParams := ibc_model.MsgRecvPacketParams{
			RawMsgRecvPacket: rawMsg,

			Application: "icahost",
			MessageType: "/chainmain.icaauth.v1.MsgSubmitTx",
			MaybeFungibleTokenPacketData: &ibc_model.MsgRecvPacketFungibleTokenPacketData{
				Success: false,
			},
		}

		// Getting possible signer address from Msg
		possibleSignerAddresses = append(possibleSignerAddresses, msgRecvPacketParams.Signer)
		cmds = append(cmds, command_usecase.NewCreateMsgIBCRecvPacket(
			parserParams.MsgCommonParams,

			msgRecvPacketParams,
		))
	}

	recvPacketEvents := log.GetEventsByType("recv_packet")
	if recvPacketEvents == nil {
		parserParams.Logger.Warnf("missing `recv_packet` event in TxsResult log on TxHash: %s", parserParams.MsgCommonParams.TxHash)
	}

	var packetSequence uint64
	var channelOrdering string
	var connectionID string
	for _, recvPacketEvent := range recvPacketEvents {
		if recvPacketEvent.HasAttribute("packet_sequence") {
			packetSequence = typeconv.MustAtou64(recvPacketEvent.MustGetAttributeByKey("packet_sequence"))
		}
		if recvPacketEvent.HasAttribute("packet_channel_ordering") {
			channelOrdering = recvPacketEvent.MustGetAttributeByKey("packet_channel_ordering")
		}
		if recvPacketEvent.HasAttribute("packet_connection") {
			connectionID = recvPacketEvent.MustGetAttributeByKey("packet_connection")
		}
	}

	ics27PacketEvents := log.GetEventsByType("ics27_packet")
	var success bool
	for _, ics27PacketEvent := range ics27PacketEvents {
		if ics27PacketEvent.HasAttribute("success") {
			success = ics27PacketEvent.MustGetAttributeByKey("success") == "true"
		}
	}

	var packetAck ibc_model.MsgRecvPacketPacketAck
	writeAckEvents := log.GetEventsByType("write_acknowledgement")
	if writeAckEvents != nil {
		var writeAckEventPacketAck string
		for _, writeAckEvent := range writeAckEvents {
			if writeAckEvent.HasAttribute("packet_ack") {
				writeAckEventPacketAck = writeAckEvent.MustGetAttributeByKey("packet_ack")
			}
		}

		json.MustUnmarshalFromString(writeAckEventPacketAck, &packetAck)
	}

	msgRecvPacketParams := ibc_model.MsgRecvPacketParams{
		RawMsgRecvPacket: rawMsg,

		Application: "icahost",
		MessageType: "/chainmain.icaauth.v1.MsgSubmitTx",
		MaybeFungibleTokenPacketData: &ibc_model.MsgRecvPacketFungibleTokenPacketData{
			Success: success,
		},

		PacketSequence:  packetSequence,
		ChannelOrdering: channelOrdering,
		ConnectionID:    connectionID,
		PacketAck:       packetAck,
	}

	// Getting possible signer address from Msg
	possibleSignerAddresses = append(possibleSignerAddresses, msgRecvPacketParams.Signer)
	cmds = append(cmds, command_usecase.NewCreateMsgIBCRecvPacket(
		parserParams.MsgCommonParams,

		msgRecvPacketParams,
	))

	if success {
		for msgIndex, message := range cosmosTx.Messages {
			switch message.Type {
			case
				// cosmos distribution
				"/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward",
				"/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission",

				// cosmos staking
				"/cosmos.staking.v1beta1.MsgDelegate",
				"/cosmos.staking.v1beta1.MsgUndelegate",
				"/cosmos.staking.v1beta1.MsgBeginRedelegate":
				// Note: will encounter the same events mapping issue as MsgExec when having multiple inner msg
				// https://github.com/crypto-com/chain-indexing/issues/673
				continue
			default:
				parser := parserParams.ParserManager.GetParser(utils.CosmosParserKey(message.Type), utils.ParserBlockHeight(parserParams.MsgCommonParams.BlockHeight))
				innerMsgs[msgIndex]["msg_index"] = strconv.Itoa(msgIndex)
				msgCommands, signers := parser(utils.CosmosParserParams{
					AddressPrefix:      parserParams.AddressPrefix,
					StakingDenom:       parserParams.StakingDenom,
					TxsResult:          parserParams.TxsResult,
					MsgCommonParams:    parserParams.MsgCommonParams,
					Msg:                innerMsgs[msgIndex],
					MsgIndex:           parserParams.MsgIndex,
					ParserManager:      parserParams.ParserManager,
					IsProposalInnerMsg: false,
				})

				possibleSignerAddresses = append(possibleSignerAddresses, signers...)
				cmds = append(cmds, msgCommands...)
			}
		}
	}

	return cmds, possibleSignerAddresses
}
