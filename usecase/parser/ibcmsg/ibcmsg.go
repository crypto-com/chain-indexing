package ibcmsg

import (
	"encoding/base64"
	"fmt"
	"time"

	base64_internal "github.com/crypto-com/chain-indexing/internal/base64"
	"github.com/crypto-com/chain-indexing/internal/json"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"

	"github.com/crypto-com/chain-indexing/entity/command"
	"github.com/crypto-com/chain-indexing/internal/typeconv"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
)

func ParseMsgCreateClient(
	msgCommonParams event.MsgCommonParams,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msg map[string]interface{},
) []command.Command {
	clientStateType := msg["client_state"].(map[string]interface{})["@type"]
	if clientStateType != "/ibc.lightclients.tendermint.v1.ClientState" {
		// TODO: SoloMachine and Localhost LightClient
		return []command.Command{}
	}

	var rawMsg ibc_model.RawMsgCreateTendermintLightClient
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			StringToDurationHookFunc(),
			StringToByteSliceHookFunc(),
		),
		Result: &rawMsg,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		panic(fmt.Errorf("error creating MsgCreateClient decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding MsgCreateClient: %v", err))
	}

	log := utils.NewParsedTxsResultLog(&txsResult.Log[msgIndex])
	// When there is no reward withdrew, `transfer` event would not exist
	event := log.GetEventByType("create_client")
	if event == nil {
		panic("missing `create_client` event in TxsResult log")
	}
	params := ibc_model.MsgCreateClientParams{
		MaybeTendermintLightClient: &ibc_model.TendermintLightClient{
			TendermintClientState:               rawMsg.ClientState,
			TendermintLightClientConsensusState: rawMsg.ConsensusState,
		},
		Signer:     rawMsg.Signer,
		ClientID:   event.MustGetAttributeByKey("client_id"),
		ClientType: event.MustGetAttributeByKey("client_type"),
	}

	return []command.Command{command_usecase.NewCreateMsgIBCCreateClient(
		msgCommonParams,

		params,
	)}
}

func ParseMsgConnectionOpenInit(
	msgCommonParams event.MsgCommonParams,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msg map[string]interface{},
) []command.Command {
	var rawMessage ibc_model.RawMsgConnectionOpenInit
	decoder, decoderErr := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			StringToByteSliceHookFunc(),
		),
		Result: &rawMessage,
	})
	if decoderErr != nil {
		panic(fmt.Errorf("error creating mapstructure decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding message: %v", err))
	}

	log := utils.NewParsedTxsResultLog(&txsResult.Log[msgIndex])
	event := log.GetEventByType("connection_open_init")
	if event == nil {
		panic("missing `connection_open_init` event in TxsResult log")
	}
	params := ibc_model.MsgConnectionOpenInitParams{
		RawMsgConnectionOpenInit: rawMessage,

		ConnectionID: event.MustGetAttributeByKey("connection_id"),
	}

	return []command.Command{command_usecase.NewCreateMsgIBCConnectionOpenInit(
		msgCommonParams,

		params,
	)}
}

func ParseMsgConnectionOpenTry(
	msgCommonParams event.MsgCommonParams,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msg map[string]interface{},
) []command.Command {
	clientStateType := msg["client_state"].(map[string]interface{})["@type"]
	if clientStateType != "/ibc.lightclients.tendermint.v1.ClientState" {
		// TODO: SoloMachine and Localhost LightClient
		return []command.Command{}
	}

	var rawMsg ibc_model.RawMsgConnectionOpenTryTendermintClient
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			StringToDurationHookFunc(),
			StringToByteSliceHookFunc(),
		),
		Result: &rawMsg,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		panic(fmt.Errorf("error creating MsgConnectionOpenTry decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding MsgConnectionOpenTry: %v", err))
	}

	log := utils.NewParsedTxsResultLog(&txsResult.Log[msgIndex])
	event := log.GetEventByType("connection_open_try")
	if event == nil {
		panic("missing `connection_open_try` event in TxsResult log")
	}

	params := ibc_model.MsgConnectionOpenTryParams{
		MsgConnectionOpenTryBaseParams: rawMsg.MsgConnectionOpenTryBaseParams,
		MaybeTendermintClientState:     &rawMsg.TendermintClientState,
		ConnectionID:                   event.MustGetAttributeByKey("connection_id"),
	}

	return []command.Command{command_usecase.NewCreateMsgIBCConnectionOpenTry(
		msgCommonParams,

		params,
	)}
}

func ParseMsgConnectionOpenAck(
	msgCommonParams event.MsgCommonParams,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msg map[string]interface{},
) []command.Command {
	clientStateType := msg["client_state"].(map[string]interface{})["@type"]
	if clientStateType != "/ibc.lightclients.tendermint.v1.ClientState" {
		// TODO: SoloMachine and Localhost LightClient
		return []command.Command{}
	}

	var rawMsg ibc_model.RawMsgConnectionOpenAckTendermintClient
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			StringToDurationHookFunc(),
			StringToByteSliceHookFunc(),
		),
		Result: &rawMsg,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		panic(fmt.Errorf("error creating MsgConnectionOpenAck decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding MsgConnectionOpenAck: %v", err))
	}

	log := utils.NewParsedTxsResultLog(&txsResult.Log[msgIndex])
	event := log.GetEventByType("connection_open_ack")
	if event == nil {
		panic("missing `connection_open_ack` event in TxsResult log")
	}

	params := ibc_model.MsgConnectionOpenAckParams{
		MsgConnectionOpenAckBaseParams: rawMsg.MsgConnectionOpenAckBaseParams,
		MaybeTendermintClientState:     &rawMsg.TendermintClientState,

		ClientID:             event.MustGetAttributeByKey("client_id"),
		CounterpartyClientID: event.MustGetAttributeByKey("counterparty_client_id"),
	}

	return []command.Command{command_usecase.NewCreateMsgIBCConnectionOpenAck(
		msgCommonParams,

		params,
	)}
}

func ParseMsgConnectionOpenConfirm(
	msgCommonParams event.MsgCommonParams,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msg map[string]interface{},
) []command.Command {
	var rawMsg ibc_model.RawMsgConnectionOpenConfirm
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			StringToDurationHookFunc(),
			StringToByteSliceHookFunc(),
		),
		Result: &rawMsg,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		panic(fmt.Errorf("error creating MsgConnectionOpenConfirm decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding MsgConnectionOpenConfirm: %v", err))
	}

	log := utils.NewParsedTxsResultLog(&txsResult.Log[msgIndex])
	event := log.GetEventByType("connection_open_confirm")
	if event == nil {
		panic("missing `connection_open_confirm` event in TxsResult log")
	}

	params := ibc_model.MsgConnectionOpenConfirmParams{
		RawMsgConnectionOpenConfirm: rawMsg,

		ClientID:                 event.MustGetAttributeByKey("client_id"),
		CounterpartyClientID:     event.MustGetAttributeByKey("counterparty_client_id"),
		CounterpartyConnectionID: event.MustGetAttributeByKey("counterparty_connection_id"),
	}

	return []command.Command{command_usecase.NewCreateMsgIBCConnectionOpenConfirm(
		msgCommonParams,

		params,
	)}
}

func ParseMsgChannelOpenInit(
	msgCommonParams event.MsgCommonParams,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msg map[string]interface{},
) []command.Command {
	var rawMsg ibc_model.RawMsgChannelOpenInit
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			StringToDurationHookFunc(),
			StringToByteSliceHookFunc(),
		),
		Result: &rawMsg,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		panic(fmt.Errorf("error creating RawMsgChannelOpenInit decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgChannelOpenInit: %v", err))
	}

	log := utils.NewParsedTxsResultLog(&txsResult.Log[msgIndex])
	event := log.GetEventByType("channel_open_init")
	if event == nil {
		panic("missing `channel_open_init` event in TxsResult log")
	}

	params := ibc_model.MsgChannelOpenInitParams{
		RawMsgChannelOpenInit: rawMsg,

		ChannelID: event.MustGetAttributeByKey("channel_id"),
	}

	return []command.Command{command_usecase.NewCreateMsgIBCChannelOpenInit(
		msgCommonParams,

		params,
	)}
}

func ParseMsgChannelOpenTry(
	msgCommonParams event.MsgCommonParams,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msg map[string]interface{},
) []command.Command {
	var rawMsg ibc_model.RawMsgChannelOpenTry
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			StringToDurationHookFunc(),
			StringToByteSliceHookFunc(),
		),
		Result: &rawMsg,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		panic(fmt.Errorf("error creating RawMsgChannelOpenTry decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgChannelOpenTry: %v", err))
	}

	log := utils.NewParsedTxsResultLog(&txsResult.Log[msgIndex])
	event := log.GetEventByType("channel_open_try")
	if event == nil {
		panic("missing `channel_open_try` event in TxsResult log")
	}

	params := ibc_model.MsgChannelOpenTryParams{
		RawMsgChannelOpenTry: rawMsg,

		ChannelID:    event.MustGetAttributeByKey("channel_id"),
		ConnectionID: event.MustGetAttributeByKey("connection_id"),
	}

	return []command.Command{command_usecase.NewCreateMsgIBCChannelOpenTry(
		msgCommonParams,

		params,
	)}
}

func ParseMsgChannelOpenAck(
	msgCommonParams event.MsgCommonParams,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msg map[string]interface{},
) []command.Command {
	var rawMsg ibc_model.RawMsgChannelOpenAck
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			StringToDurationHookFunc(),
			StringToByteSliceHookFunc(),
		),
		Result: &rawMsg,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		panic(fmt.Errorf("error creating RawMsgChannelOpenAck decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgChannelOpenAck: %v", err))
	}

	log := utils.NewParsedTxsResultLog(&txsResult.Log[msgIndex])
	event := log.GetEventByType("channel_open_ack")
	if event == nil {
		panic("missing `channel_open_ack` event in TxsResult log")
	}

	params := ibc_model.MsgChannelOpenAckParams{
		RawMsgChannelOpenAck: rawMsg,

		CounterpartyPortID: event.MustGetAttributeByKey("counterparty_port_id"),
		ConnectionID:       event.MustGetAttributeByKey("connection_id"),
	}

	return []command.Command{command_usecase.NewCreateMsgIBCChannelOpenAck(
		msgCommonParams,

		params,
	)}
}

func ParseMsgChannelOpenConfirm(
	msgCommonParams event.MsgCommonParams,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msg map[string]interface{},
) []command.Command {
	var rawMsg ibc_model.RawMsgChannelOpenConfirm
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			StringToDurationHookFunc(),
			StringToByteSliceHookFunc(),
		),
		Result: &rawMsg,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		panic(fmt.Errorf("error creating RawMsgChannelOpenConfirm decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgChannelOpenConfirm: %v", err))
	}

	log := utils.NewParsedTxsResultLog(&txsResult.Log[msgIndex])
	event := log.GetEventByType("channel_open_confirm")
	if event == nil {
		panic("missing `channel_open_confirm` event in TxsResult log")
	}

	params := ibc_model.MsgChannelOpenConfirmParams{
		RawMsgChannelOpenConfirm: rawMsg,

		CounterpartyChannelID: event.MustGetAttributeByKey("counterparty_channel_id"),
		CounterpartyPortID:    event.MustGetAttributeByKey("counterparty_port_id"),
		ConnectionID:          event.MustGetAttributeByKey("connection_id"),
	}

	return []command.Command{command_usecase.NewCreateMsgIBCChannelOpenConfirm(
		msgCommonParams,

		params,
	)}
}

func ParseMsgUpdateClient(
	msgCommonParams event.MsgCommonParams,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msg map[string]interface{},
) []command.Command {
	headerType := msg["header"].(map[string]interface{})["@type"]
	if headerType != "/ibc.lightclients.tendermint.v1.Header" {
		// TODO: SoloMachine and Localhost LightClient
		return []command.Command{}
	}

	var rawMsg ibc_model.RawMsgUpdateTendermintLightClient
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			StringToByteSliceHookFunc(),
		),
		Result: &rawMsg,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		panic(fmt.Errorf("error creating MsgUpdateClient decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding MsgUpdateClient: %v", err))
	}

	log := utils.NewParsedTxsResultLog(&txsResult.Log[msgIndex])
	event := log.GetEventByType("update_client")
	if event == nil {
		panic("missing `update_client` event in TxsResult log")
	}

	params := ibc_model.MsgUpdateClientParams{
		MaybeTendermintLightClientUpdate: &ibc_model.TendermintLightClientUpdate{Header: rawMsg.Header},

		ClientID:        rawMsg.ClientID,
		ClientType:      event.MustGetAttributeByKey("client_type"),
		ConsensusHeight: event.MustGetAttributeByKey("consensus_height"),
		Signer:          rawMsg.Signer,
	}

	return []command.Command{command_usecase.NewCreateMsgIBCUpdateClient(
		msgCommonParams,

		params,
	)}
}

func ParseMsgTransfer(
	msgCommonParams event.MsgCommonParams,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msg map[string]interface{},
) []command.Command {
	var rawMsg ibc_model.RawMsgTransfer
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			StringToDurationHookFunc(),
			StringToByteSliceHookFunc(),
		),
		Result: &rawMsg,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		panic(fmt.Errorf("error creating RawMsgTransfer decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgTransfer: %v", err))
	}

	log := utils.NewParsedTxsResultLog(&txsResult.Log[msgIndex])
	event := log.GetEventByType("send_packet")
	if event == nil {
		panic("missing `send_packet` event in TxsResult log")
	}

	params := ibc_model.MsgTransferParams{
		RawMsgTransfer: rawMsg,

		PacketSequence:     typeconv.MustAtou64(event.MustGetAttributeByKey("packet_sequence")),
		DestinationPort:    event.MustGetAttributeByKey("packet_dst_port"),
		DestinationChannel: event.MustGetAttributeByKey("packet_dst_channel"),
		ChannelOrdering:    event.MustGetAttributeByKey("packet_channel_ordering"),
		ConnectionID:       event.MustGetAttributeByKey("packet_connection"),
	}

	return []command.Command{command_usecase.NewCreateMsgIBCTransferTransfer(
		msgCommonParams,

		params,
	)}
}

func ParseMsgRecvPacket(
	msgCommonParams event.MsgCommonParams,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msg map[string]interface{},
) []command.Command {
	log := utils.NewParsedTxsResultLog(&txsResult.Log[msgIndex])
	fmt.Printf("%+v", txsResult.Log[msgIndex])

	var rawMsg ibc_model.RawMsgRecvPacket
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			StringToDurationHookFunc(),
			StringToByteSliceHookFunc(),
		),
		Result: &rawMsg,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		panic(fmt.Errorf("error creating RawMsgRecvPacket decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgRecvPacket: %v", err))
	}

	if !isPacketMsgTransfer(log, rawMsg) {
		// unsupported application
		return []command.Command{}
	}

	// Transfer application, MsgTransfer
	var rawFungibleTokenPacketData ibc_model.RawMsgRecvPacketFungibleTokenPacketData
	rawPacketData := base64_internal.MustDecodeString(rawMsg.Packet.Data)
	json.MustUnmarshal(rawPacketData, &rawFungibleTokenPacketData)

	fungibleTokenPacketEvent := log.GetEventByType("fungible_token_packet")
	if fungibleTokenPacketEvent == nil {
		panic("missing `fungible_token_packet` event in TxsResult log")
	}

	denominationTraceEvent := log.GetEventByType("denomination_trace")
	if denominationTraceEvent == nil {
		panic("missing `denomination_trace` event in TxsResult log")
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

	params := ibc_model.MsgRecvPacketParams{
		RawMsgRecvPacket: rawMsg,

		Application: "transfer",
		MessageType: "MsgTransfer",
		MaybeFungibleTokenPacketData: &ibc_model.MsgRecvPacketFungibleTokenPacketData{
			RawMsgRecvPacketFungibleTokenPacketData: rawFungibleTokenPacketData,
			// success value inverted bug: https://github.com/cosmos/cosmos-sdk/pull/9640
			Success:                fungibleTokenPacketEvent.MustGetAttributeByKey("success") == "false",
			DenominationTraceHash:  denominationTraceEvent.MustGetAttributeByKey("trace_hash"),
			DenominationTraceDenom: denominationTraceEvent.MustGetAttributeByKey("denom"),
		},

		PacketSequence:  typeconv.MustAtou64(recvPacketEvent.MustGetAttributeByKey("packet_sequence")),
		ChannelOrdering: recvPacketEvent.MustGetAttributeByKey("packet_channel_ordering"),
		ConnectionID:    recvPacketEvent.MustGetAttributeByKey("packet_connection"),
		PacketAck:       packetAck,
	}

	return []command.Command{command_usecase.NewCreateMsgIBCCoreRecvPacket(
		msgCommonParams,

		params,
	)}
}

func isPacketMsgTransfer(
	log *utils.ParsedTxsResultLog,
	rawMsgRecvPacket ibc_model.RawMsgRecvPacket,
) bool {
	if !log.HasEvent("fungible_token_packet") {
		return false
	}

	if rawMsgRecvPacket.Packet.DestinationPort != "transfer" {
		return false
	}

	var fungiblePacketData ibc_model.RawMsgRecvPacketFungibleTokenPacketData
	rawPacketData, decodeDataErr := base64.StdEncoding.DecodeString(rawMsgRecvPacket.Packet.Data)
	if decodeDataErr != nil {
		return false
	}
	if err := jsoniter.Unmarshal(rawPacketData, &fungiblePacketData); err != nil {
		fmt.Println("packet gg")
		return false
	}

	return true
}
