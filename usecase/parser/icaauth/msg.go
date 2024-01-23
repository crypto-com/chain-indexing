package icaauth

import (
	"encoding/json"
	"fmt"

	"github.com/crypto-com/chain-indexing/entity/command"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	icaauth_model "github.com/crypto-com/chain-indexing/usecase/model/icaauth"
	"github.com/crypto-com/chain-indexing/usecase/parser/ibc"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
	mapstructure_utils "github.com/crypto-com/chain-indexing/usecase/parser/utils/mapstructure"
)

func ParseChainmainMsgSubmitTx(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var rawMsg icaauth_model.RawMsgSubmitTx
	if err := mapstructure_utils.DefaultMapstructureDecoder.Decode(parserParams.Msg, &rawMsg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgSubmitTx: %v", err))
	}

	var msgSubmitTx icaauth_model.MsgSubmitTx
	msg, err := json.Marshal(rawMsg)
	if err != nil {
		panic(fmt.Errorf("error json marshalling RawMsgSubmitTx: %v", err))
	}

	err = json.Unmarshal(msg, &msgSubmitTx)
	if err != nil {
		panic(fmt.Errorf("error json unmarshlling MsgSubmitTx: %v", err))
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		msgSubmitTxParams := icaauth_model.MsgSubmitTxParams{
			MsgSubmitTx: msgSubmitTx,
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, msgSubmitTxParams.Owner)

		return []command.Command{command_usecase.NewCreateChainmainMsgSubmitTx(
			parserParams.MsgCommonParams,

			msgSubmitTxParams,
		)}, possibleSignerAddresses
	}

	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])

	sendPacketEvents := log.GetEventsByType("send_packet")
	if sendPacketEvents == nil {
		panic("missing `send_packet` event in TxsResult log")
	}
	var packetData string
	var packetTimeoutHeight string
	var packetTimeoutTimestamp string
	var packetSequence string
	var packetSrcPort string
	var packetSrcChannel string
	var packetDstPort string
	var packetDstChannel string
	for _, sendPacketEvent := range sendPacketEvents {
		if sendPacketEvent.HasAttribute("packet_data") {
			packetData = sendPacketEvent.MustGetAttributeByKey("packet_data")
		}
		if sendPacketEvent.HasAttribute("packet_timeout_height") {
			packetTimeoutHeight = sendPacketEvent.MustGetAttributeByKey("packet_timeout_height")
		}
		if sendPacketEvent.HasAttribute("packet_timeout_timestamp") {
			packetTimeoutTimestamp = sendPacketEvent.MustGetAttributeByKey("packet_timeout_timestamp")
		}
		if sendPacketEvent.HasAttribute("packet_sequence") {
			packetSequence = sendPacketEvent.MustGetAttributeByKey("packet_sequence")
		}
		if sendPacketEvent.HasAttribute("packet_src_port") {
			packetSrcPort = sendPacketEvent.MustGetAttributeByKey("packet_src_port")
		}
		if sendPacketEvent.HasAttribute("packet_src_channel") {
			packetSrcChannel = sendPacketEvent.MustGetAttributeByKey("packet_src_channel")
		}
		if sendPacketEvent.HasAttribute("packet_dst_port") {
			packetDstPort = sendPacketEvent.MustGetAttributeByKey("packet_dst_port")
		}
		if sendPacketEvent.HasAttribute("packet_dst_channel") {
			packetDstChannel = sendPacketEvent.MustGetAttributeByKey("packet_dst_channel")
		}
	}

	msgSubmitTxParams := icaauth_model.MsgSubmitTxParams{
		MsgSubmitTx: msgSubmitTx,

		Packet: ibc_model.Packet{
			Sequence:           packetSequence,
			SourcePort:         packetSrcPort,
			SourceChannel:      packetSrcChannel,
			DestinationPort:    packetDstPort,
			DestinationChannel: packetDstChannel,
			Data:               packetData,
			TimeoutHeight:      ibc.MustParseHeight(packetTimeoutHeight),
			TimeoutTimestamp:   packetTimeoutTimestamp,
		},
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	possibleSignerAddresses = append(possibleSignerAddresses, msgSubmitTxParams.Owner)

	return []command.Command{command_usecase.NewCreateChainmainMsgSubmitTx(
		parserParams.MsgCommonParams,

		msgSubmitTxParams,
	)}, possibleSignerAddresses
}

func ParseChainmainMsgRegisterAccount(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var rawMsg icaauth_model.RawChainmainMsgRegisterAccount
	if err := mapstructure_utils.DefaultMapstructureDecoder.Decode(parserParams.Msg, &rawMsg); err != nil {
		panic(fmt.Errorf("error decoding RawChainmainMsgRegisterAccount: %v", err))
	}

	var msgRegisterAccount icaauth_model.MsgRegisterAccount
	msg, err := json.Marshal(rawMsg)
	if err != nil {
		panic(fmt.Errorf("error json marshalling RawChainmainMsgRegisterAccount: %v", err))
	}

	err = json.Unmarshal(msg, &msgRegisterAccount)
	if err != nil {
		panic(fmt.Errorf("error json unmarshalling MsgRegisterAccount: %v", err))
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, rawMsg.Owner)

		return []command.Command{command_usecase.NewCreateChainmainMsgRegisterAccount(
			parserParams.MsgCommonParams,

			icaauth_model.MsgRegisterAccountParams{
				MsgRegisterAccount: msgRegisterAccount,
			},
		)}, possibleSignerAddresses
	}

	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])
	event := log.GetEventByType("channel_open_init")
	if event == nil {
		panic("missing `channel_open_init` event in TxsResult log")
	}

	msgRegisterAccountParams := icaauth_model.MsgRegisterAccountParams{
		MsgRegisterAccount: msgRegisterAccount,

		PortID:                event.MustGetAttributeByKey("port_id"),
		ChannelID:             event.MustGetAttributeByKey("channel_id"),
		CounterpartyPortID:    event.MustGetAttributeByKey("counterparty_port_id"),
		CounterpartyChannelID: event.MustGetAttributeByKey("counterparty_channel_id"),
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	possibleSignerAddresses = append(possibleSignerAddresses, msgRegisterAccountParams.Owner)

	return []command.Command{command_usecase.NewCreateChainmainMsgRegisterAccount(
		parserParams.MsgCommonParams,

		msgRegisterAccountParams,
	)}, possibleSignerAddresses
}

func ParseMsgSubmitTx(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var rawMsg icaauth_model.RawMsgSubmitTx
	if err := mapstructure_utils.DefaultMapstructureDecoder.Decode(parserParams.Msg, &rawMsg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgSubmitTx: %v", err))
	}

	var msgSubmitTx icaauth_model.MsgSubmitTx
	msg, err := json.Marshal(rawMsg)
	if err != nil {
		panic(fmt.Errorf("error json marshalling RawMsgSubmitTx: %v", err))
	}

	err = json.Unmarshal(msg, &msgSubmitTx)
	if err != nil {
		panic(fmt.Errorf("error json unmarshlling MsgSubmitTx: %v", err))
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		msgSubmitTxParams := icaauth_model.MsgSubmitTxParams{
			MsgSubmitTx: msgSubmitTx,
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, msgSubmitTxParams.Owner)

		return []command.Command{command_usecase.NewCreateMsgSubmitTx(
			parserParams.MsgCommonParams,

			msgSubmitTxParams,
		)}, possibleSignerAddresses
	}

	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])

	sendPacketEvents := log.GetEventsByType("send_packet")
	if sendPacketEvents == nil {
		panic("missing `send_packet` event in TxsResult log")
	}
	var packetData string
	var packetTimeoutHeight string
	var packetTimeoutTimestamp string
	var packetSequence string
	var packetSrcPort string
	var packetSrcChannel string
	var packetDstPort string
	var packetDstChannel string
	for _, sendPacketEvent := range sendPacketEvents {
		if sendPacketEvent.HasAttribute("packet_data") {
			packetData = sendPacketEvent.MustGetAttributeByKey("packet_data")
		}
		if sendPacketEvent.HasAttribute("packet_timeout_height") {
			packetTimeoutHeight = sendPacketEvent.MustGetAttributeByKey("packet_timeout_height")
		}
		if sendPacketEvent.HasAttribute("packet_timeout_timestamp") {
			packetTimeoutTimestamp = sendPacketEvent.MustGetAttributeByKey("packet_timeout_timestamp")
		}
		if sendPacketEvent.HasAttribute("packet_sequence") {
			packetSequence = sendPacketEvent.MustGetAttributeByKey("packet_sequence")
		}
		if sendPacketEvent.HasAttribute("packet_src_port") {
			packetSrcPort = sendPacketEvent.MustGetAttributeByKey("packet_src_port")
		}
		if sendPacketEvent.HasAttribute("packet_src_channel") {
			packetSrcChannel = sendPacketEvent.MustGetAttributeByKey("packet_src_channel")
		}
		if sendPacketEvent.HasAttribute("packet_dst_port") {
			packetDstPort = sendPacketEvent.MustGetAttributeByKey("packet_dst_port")
		}
		if sendPacketEvent.HasAttribute("packet_dst_channel") {
			packetDstChannel = sendPacketEvent.MustGetAttributeByKey("packet_dst_channel")
		}
	}

	msgSubmitTxParams := icaauth_model.MsgSubmitTxParams{
		MsgSubmitTx: msgSubmitTx,

		Packet: ibc_model.Packet{
			Sequence:           packetSequence,
			SourcePort:         packetSrcPort,
			SourceChannel:      packetSrcChannel,
			DestinationPort:    packetDstPort,
			DestinationChannel: packetDstChannel,
			Data:               packetData,
			TimeoutHeight:      ibc.MustParseHeight(packetTimeoutHeight),
			TimeoutTimestamp:   packetTimeoutTimestamp,
		},
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	possibleSignerAddresses = append(possibleSignerAddresses, msgSubmitTxParams.Owner)

	return []command.Command{command_usecase.NewCreateMsgSubmitTx(
		parserParams.MsgCommonParams,

		msgSubmitTxParams,
	)}, possibleSignerAddresses
}

func ParseMsgRegisterAccount(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var rawMsg icaauth_model.RawMsgRegisterAccount
	if err := mapstructure_utils.DefaultMapstructureDecoder.Decode(parserParams.Msg, &rawMsg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgRegisterAccount: %v", err))
	}

	var normalizedRawMsg icaauth_model.RawChainmainMsgRegisterAccount
	normalizedMsgRegisterAccount, _ := json.MarshalIndent(icaauth_model.RawChainmainMsgRegisterAccount(rawMsg), "", "  ")
	err := json.Unmarshal(normalizedMsgRegisterAccount, &normalizedRawMsg)
	if err != nil {
		panic(fmt.Errorf("error json unmarshalling MsgRegisterAccount: %v", err))
	}

	fmt.Println(parserParams.Msg)
	fmt.Println(normalizedRawMsg)

	var msgRegisterAccount icaauth_model.MsgRegisterAccount
	msg, err := json.Marshal(normalizedRawMsg)
	if err != nil {
		panic(fmt.Errorf("error json marshalling RawMsgRegisterAccount: %v", err))
	}

	err = json.Unmarshal(msg, &msgRegisterAccount)
	if err != nil {
		panic(fmt.Errorf("error json unmarshalling MsgRegisterAccount: %v", err))
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, rawMsg.Owner)

		return []command.Command{command_usecase.NewCreateMsgRegisterAccount(
			parserParams.MsgCommonParams,

			icaauth_model.MsgRegisterAccountParams{
				MsgRegisterAccount: msgRegisterAccount,
			},
		)}, possibleSignerAddresses
	}

	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])

	if parserParams.IsEthereumTxInnerMsg {
		var commands []command.Command
		var possibleSignerAddresses []string

		events := log.GetEventsByType("channel_open_init")
		msgRegisterAccountParams := make([]icaauth_model.MsgRegisterAccountParams, 0)

		for _, event := range events {
			param := icaauth_model.MsgRegisterAccountParams{
				MsgRegisterAccount: icaauth_model.MsgRegisterAccount{
					Owner:        msgRegisterAccount.Owner,
					ConnectionID: event.MustGetAttributeByKey("connection_id"),
					Version:      event.MustGetAttributeByKey("version"),
				},

				PortID:                event.MustGetAttributeByKey("port_id"),
				ChannelID:             event.MustGetAttributeByKey("channel_id"),
				CounterpartyPortID:    event.MustGetAttributeByKey("counterparty_port_id"),
				CounterpartyChannelID: event.MustGetAttributeByKey("counterparty_channel_id"),
			}

			isExisted := false
			for _, msgRegisterAccountParam := range msgRegisterAccountParams {
				if msgRegisterAccountParam.ChannelID == param.ChannelID && msgRegisterAccountParam.PortID == param.PortID && msgRegisterAccountParam.CounterpartyChannelID == param.CounterpartyChannelID {
					isExisted = true
					break
				}
			}

			if !isExisted {
				// append non-duplicated msg register account params extracted from channel open init event
				msgRegisterAccountParams = append(msgRegisterAccountParams, param)

				// append commands
				commands = append(commands, command_usecase.NewCreateMsgRegisterAccount(
					parserParams.MsgCommonParams,
					param,
				))

				// append possible signer addresses
				possibleSignerAddresses = append(possibleSignerAddresses, rawMsg.Owner)
			}
		}
		return commands, possibleSignerAddresses
	}

	event := log.GetEventByType("channel_open_init")
	if event == nil {
		panic("missing `channel_open_init` event in TxsResult log")
	}

	msgRegisterAccountParams := icaauth_model.MsgRegisterAccountParams{
		MsgRegisterAccount: msgRegisterAccount,

		PortID:                event.MustGetAttributeByKey("port_id"),
		ChannelID:             event.MustGetAttributeByKey("channel_id"),
		CounterpartyPortID:    event.MustGetAttributeByKey("counterparty_port_id"),
		CounterpartyChannelID: event.MustGetAttributeByKey("counterparty_channel_id"),
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	possibleSignerAddresses = append(possibleSignerAddresses, msgRegisterAccountParams.Owner)

	return []command.Command{command_usecase.NewCreateMsgRegisterAccount(
		parserParams.MsgCommonParams,

		msgRegisterAccountParams,
	)}, possibleSignerAddresses
}
