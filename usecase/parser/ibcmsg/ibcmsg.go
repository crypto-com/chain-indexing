package ibcmsg

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
	log2 "log"
	"strings"
	"time"

	base64_internal "github.com/crypto-com/chain-indexing/internal/base64"
	"github.com/crypto-com/chain-indexing/internal/json"
	"github.com/crypto-com/chain-indexing/internal/primptr"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"

	"github.com/crypto-com/chain-indexing/entity/command"
	"github.com/crypto-com/chain-indexing/internal/typeconv"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

const tendermintClientStateTypeV1 = "/ibc.lightclients.tendermint.v1.ClientState"
const soloMachineClientStateTypeV2 = "/ibc.lightclients.solomachine.v2.ClientState"
const tendermintHeaderTypeV1 = "/ibc.lightclients.tendermint.v1.Header"
const soloMachineHeaderTypeV2 = "/ibc.lightclients.solomachine.v2.Header"

func ParseMsgCreateClient(
	msgCommonParams event.MsgCommonParams,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msg map[string]interface{},
) []command.Command {
	clientStateType := msg["client_state"].(map[string]interface{})["@type"]

	switch clientStateType {
	case tendermintClientStateTypeV1:
		return parseRawMsgCreateTendermintLightClient(
			msgCommonParams,
			txsResult,
			msgIndex,
			msg,
		)
	case soloMachineClientStateTypeV2:
		return parseRawMsgCreateSoloMachineLightClient(
			msgCommonParams,
			txsResult,
			msgIndex,
			msg,
		)
	default:
		return []command.Command{}

	}
}

func parseRawMsgCreateTendermintLightClient(
	msgCommonParams event.MsgCommonParams,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msg map[string]interface{},
) []command.Command {
	var rawMsg ibc_model.RawMsgCreateTendermintLightClient
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			utils.StringToDurationHookFunc(),
			utils.StringToByteSliceHookFunc(),
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

	if !msgCommonParams.TxSuccess {
		params := ibc_model.MsgCreateClientParams{
			MaybeTendermintLightClient: &ibc_model.TendermintLightClient{
				TendermintClientState:               rawMsg.ClientState,
				TendermintLightClientConsensusState: rawMsg.ConsensusState,
			},
			Signer: rawMsg.Signer,
		}

		return []command.Command{command_usecase.NewCreateMsgIBCCreateClient(
			msgCommonParams,

			params,
		)}
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

func parseRawMsgCreateSoloMachineLightClient(
	msgCommonParams event.MsgCommonParams,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msg map[string]interface{},
) []command.Command {
	var rawMsg ibc_model.RawMsgCreateSoloMachineLightClient
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			utils.StringToDurationHookFunc(),
			utils.StringToByteSliceHookFunc(),
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

	if !msgCommonParams.TxSuccess {
		params := ibc_model.MsgCreateClientParams{
			MaybeSoloMachineLightClient: &ibc_model.SoloMachineLightClient{
				SoloMachineClientState:               rawMsg.ClientState,
				SoloMachineLightClientConsensusState: rawMsg.ConsensusState,
			},
			Signer: rawMsg.Signer,
		}

		return []command.Command{command_usecase.NewCreateMsgIBCCreateClient(
			msgCommonParams,

			params,
		)}
	}

	log := utils.NewParsedTxsResultLog(&txsResult.Log[msgIndex])
	// When there is no reward withdrew, `transfer` event would not exist
	event := log.GetEventByType("create_client")
	if event == nil {
		panic("missing `create_client` event in TxsResult log")
	}
	params := ibc_model.MsgCreateClientParams{
		MaybeSoloMachineLightClient: &ibc_model.SoloMachineLightClient{
			SoloMachineClientState:               rawMsg.ClientState,
			SoloMachineLightClientConsensusState: rawMsg.ConsensusState,
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
			utils.StringToByteSliceHookFunc(),
		),
		Result: &rawMessage,
	})
	if decoderErr != nil {
		panic(fmt.Errorf("error creating mapstructure decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding message: %v", err))
	}

	if !msgCommonParams.TxSuccess {
		return []command.Command{command_usecase.NewCreateMsgIBCConnectionOpenInit(
			msgCommonParams,

			ibc_model.MsgConnectionOpenInitParams{
				RawMsgConnectionOpenInit: rawMessage,
			},
		)}
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
			utils.StringToDurationHookFunc(),
			utils.StringToByteSliceHookFunc(),
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

	if !msgCommonParams.TxSuccess {
		params := ibc_model.MsgConnectionOpenTryParams{
			MsgConnectionOpenTryBaseParams: rawMsg.MsgConnectionOpenTryBaseParams,
			MaybeTendermintClientState:     &rawMsg.TendermintClientState,
		}

		return []command.Command{command_usecase.NewCreateMsgIBCConnectionOpenTry(
			msgCommonParams,

			params,
		)}
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
			utils.StringToDurationHookFunc(),
			utils.StringToByteSliceHookFunc(),
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

	if !msgCommonParams.TxSuccess {
		params := ibc_model.MsgConnectionOpenAckParams{
			MsgConnectionOpenAckBaseParams: rawMsg.MsgConnectionOpenAckBaseParams,
			MaybeTendermintClientState:     &rawMsg.TendermintClientState,
		}

		return []command.Command{command_usecase.NewCreateMsgIBCConnectionOpenAck(
			msgCommonParams,

			params,
		)}
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
			utils.StringToDurationHookFunc(),
			utils.StringToByteSliceHookFunc(),
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

	if !msgCommonParams.TxSuccess {
		return []command.Command{command_usecase.NewCreateMsgIBCConnectionOpenConfirm(
			msgCommonParams,

			ibc_model.MsgConnectionOpenConfirmParams{
				RawMsgConnectionOpenConfirm: rawMsg,
			},
		)}
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
			utils.StringToDurationHookFunc(),
			utils.StringToByteSliceHookFunc(),
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

	if !msgCommonParams.TxSuccess {
		return []command.Command{command_usecase.NewCreateMsgIBCChannelOpenInit(
			msgCommonParams,

			ibc_model.MsgChannelOpenInitParams{
				RawMsgChannelOpenInit: rawMsg,
			},
		)}
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
			utils.StringToDurationHookFunc(),
			utils.StringToByteSliceHookFunc(),
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

	if !msgCommonParams.TxSuccess {
		return []command.Command{command_usecase.NewCreateMsgIBCChannelOpenTry(
			msgCommonParams,

			ibc_model.MsgChannelOpenTryParams{
				RawMsgChannelOpenTry: rawMsg,
			},
		)}
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
			utils.StringToDurationHookFunc(),
			utils.StringToByteSliceHookFunc(),
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

	if !msgCommonParams.TxSuccess {
		return []command.Command{command_usecase.NewCreateMsgIBCChannelOpenAck(
			msgCommonParams,

			ibc_model.MsgChannelOpenAckParams{
				RawMsgChannelOpenAck: rawMsg,
			},
		)}
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
			utils.StringToDurationHookFunc(),
			utils.StringToByteSliceHookFunc(),
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

	if !msgCommonParams.TxSuccess {
		return []command.Command{command_usecase.NewCreateMsgIBCChannelOpenConfirm(
			msgCommonParams,

			ibc_model.MsgChannelOpenConfirmParams{
				RawMsgChannelOpenConfirm: rawMsg,
			},
		)}
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

	switch headerType {
	case tendermintHeaderTypeV1:
		return parseMsgUpdateTendermintLightClient(
			msgCommonParams,
			txsResult,
			msgIndex,
			msg,
		)
	case soloMachineHeaderTypeV2:
		return parseMsgUpdateSolomachineLightClient(
			msgCommonParams,
			txsResult,
			msgIndex,
			msg,
		)
	default:
		return []command.Command{}

	}
}

func parseMsgUpdateTendermintLightClient(
	msgCommonParams event.MsgCommonParams,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msg map[string]interface{},
) []command.Command {
	var rawMsg ibc_model.RawMsgUpdateTendermintLightClient
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			utils.StringToByteSliceHookFunc(),
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

	if !msgCommonParams.TxSuccess {
		params := ibc_model.MsgUpdateClientParams{
			MaybeTendermintLightClientUpdate: &ibc_model.TendermintLightClientUpdate{Header: rawMsg.Header},

			ClientID:        rawMsg.ClientID,
			ClientType:      "",
			ConsensusHeight: ibc_model.Height{},
			Signer:          rawMsg.Signer,
		}

		return []command.Command{command_usecase.NewCreateMsgIBCUpdateClient(
			msgCommonParams,

			params,
		)}
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
		ConsensusHeight: mustParseHeight(event.MustGetAttributeByKey("consensus_height")),
		Signer:          rawMsg.Signer,
	}

	return []command.Command{command_usecase.NewCreateMsgIBCUpdateClient(
		msgCommonParams,

		params,
	)}
}

func parseMsgUpdateSolomachineLightClient(
	msgCommonParams event.MsgCommonParams,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msg map[string]interface{},
) []command.Command {
	var rawMsg ibc_model.RawMsgUpdateSoloMachineLightClient
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			utils.StringToByteSliceHookFunc(),
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

	if !msgCommonParams.TxSuccess {
		params := ibc_model.MsgUpdateClientParams{
			MaybeSoloMachineLightClientUpdate: &ibc_model.SoloMachineLightClientUpdate{Header: rawMsg.Header},

			ClientID:        rawMsg.ClientID,
			ClientType:      "",
			ConsensusHeight: ibc_model.Height{},
			Signer:          rawMsg.Signer,
		}

		return []command.Command{command_usecase.NewCreateMsgIBCUpdateClient(
			msgCommonParams,

			params,
		)}
	}

	log := utils.NewParsedTxsResultLog(&txsResult.Log[msgIndex])
	event := log.GetEventByType("update_client")
	if event == nil {
		panic("missing `update_client` event in TxsResult log")
	}

	params := ibc_model.MsgUpdateClientParams{
		MaybeSoloMachineLightClientUpdate: &ibc_model.SoloMachineLightClientUpdate{Header: rawMsg.Header},

		ClientID:        rawMsg.ClientID,
		ClientType:      event.MustGetAttributeByKey("client_type"),
		ConsensusHeight: mustParseHeight(event.MustGetAttributeByKey("consensus_height")),
		Signer:          rawMsg.Signer,
	}

	return []command.Command{command_usecase.NewCreateMsgIBCUpdateClient(
		msgCommonParams,

		params,
	)}
}

func mustParseHeight(height string) ibc_model.Height {
	heightTokens := strings.Split(height, "-")
	if len(heightTokens) != 2 {
		panic("invalid height")
	}

	revisionNumber := typeconv.MustAtou64(heightTokens[0])
	revisionHeight := typeconv.MustAtou64(heightTokens[1])
	return ibc_model.Height{
		RevisionNumber: revisionNumber,
		RevisionHeight: revisionHeight,
	}
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
			utils.StringToDurationHookFunc(),
			utils.StringToByteSliceHookFunc(),
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

	if !msgCommonParams.TxSuccess {
		return []command.Command{command_usecase.NewCreateMsgIBCTransferTransfer(
			msgCommonParams,
			ibc_model.MsgTransferParams{
				RawMsgTransfer: rawMsg,
			},
		)}
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
	cosmosParserParams utils.CosmosParserParams,
) []command.Command {
	log2.Println("v0 ParseMsgRecvPacket with block height:", cosmosParserParams.MsgCommonParams.BlockHeight)
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
	if err := decoder.Decode(cosmosParserParams.Msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgRecvPacket: %v", err))
	}

	if !IsPacketMsgTransfer(rawMsg.Packet) {
		// unsupported application
		return []command.Command{}
	}

	// Transfer application, MsgTransfer
	var rawFungibleTokenPacketData ibc_model.FungibleTokenPacketData
	rawPacketData := base64_internal.MustDecodeString(rawMsg.Packet.Data)
	json.MustUnmarshal(rawPacketData, &rawFungibleTokenPacketData)

	if !cosmosParserParams.MsgCommonParams.TxSuccess {
		params := ibc_model.MsgRecvPacketParams{
			RawMsgRecvPacket: rawMsg,

			Application: "transfer",
			MessageType: "MsgTransfer",
			MaybeFungibleTokenPacketData: &ibc_model.MsgRecvPacketFungibleTokenPacketData{
				FungibleTokenPacketData: rawFungibleTokenPacketData,
			},
		}

		return []command.Command{command_usecase.NewCreateMsgIBCRecvPacket(
			cosmosParserParams.MsgCommonParams,

			params,
		)}
	}

	log := utils.NewParsedTxsResultLog(&cosmosParserParams.TxsResult.Log[cosmosParserParams.MsgIndex])

	recvPacketEvent := log.GetEventByType("recv_packet")
	if recvPacketEvent == nil {
		panic("missing `recv_packet` event in TxsResult log")
	}

	fungibleTokenPacketEvent := log.GetEventByType("fungible_token_packet")
	if fungibleTokenPacketEvent == nil {
		// Note: this could happen when the packet is already relayed.
		// https://github.com/cosmos/ibc-go/blob/760d15a3a55397678abe311b7f65203b2e8437d6/modules/core/04-channel/keeper/packet.go#L239
		// https://github.com/cosmos/ibc-go/blob/760d15a3a55397678abe311b7f65203b2e8437d6/modules/core/keeper/msg_server.go#L508

		packetAck := ibc_model.MsgRecvPacketPacketAck{
			MaybeError: primptr.String("missing `fungible_token_packet` event in TxsResult log, this could happen when the packet is already relayed"),
		}

		params := ibc_model.MsgRecvPacketParams{
			RawMsgRecvPacket: rawMsg,

			Application: "transfer",
			MessageType: "MsgTransfer",
			MaybeFungibleTokenPacketData: &ibc_model.MsgRecvPacketFungibleTokenPacketData{
				FungibleTokenPacketData: rawFungibleTokenPacketData,
				Success:                 false,
			},

			PacketSequence: typeconv.MustAtou64(recvPacketEvent.MustGetAttributeByKey("packet_sequence")),
			// TODO: Remove this when IBCChannel projection no longer relies on packetAck.MaybeError to check if MsgRecvPacket is success or not
			PacketAck: packetAck,
		}

		return []command.Command{command_usecase.NewCreateMsgIBCRecvPacket(
			msgCommonParams,

			params,
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

	params := ibc_model.MsgRecvPacketParams{
		RawMsgRecvPacket: rawMsg,

		Application: "transfer",
		MessageType: "MsgTransfer",
		MaybeFungibleTokenPacketData: &ibc_model.MsgRecvPacketFungibleTokenPacketData{
			FungibleTokenPacketData: rawFungibleTokenPacketData,
			// success value inverted bug: https://github.com/cosmos/cosmos-sdk/pull/9640
			// TODO: after fixing this, please also update `projection/ibc_channel/ibc_channel.go -> HandleEvents()`
			//			 when event type is `MsgIBCRecvPacket`.
			Success:                fungibleTokenPacketEvent.MustGetAttributeByKey("success") == "false",
			MaybeDenominationTrace: maybeDenominationTrace,
		},

		PacketSequence:  typeconv.MustAtou64(recvPacketEvent.MustGetAttributeByKey("packet_sequence")),
		ChannelOrdering: recvPacketEvent.MustGetAttributeByKey("packet_channel_ordering"),
		ConnectionID:    recvPacketEvent.MustGetAttributeByKey("packet_connection"),
		PacketAck:       packetAck,
	}

	return []command.Command{command_usecase.NewCreateMsgIBCRecvPacket(
		cosmosParserParams.MsgCommonParams,

		params,
	)}
}

func ParseMsgAcknowledgement(
	msgCommonParams event.MsgCommonParams,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msg map[string]interface{},
) []command.Command {
	var rawMsg ibc_model.RawMsgAcknowledgement
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
		panic(fmt.Errorf("error creating RawMsgAcknowledgement decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgAcknowledgement: %v", err))
	}

	if !IsPacketMsgTransfer(rawMsg.Packet) {
		// unsupported application
		return []command.Command{}
	}

	// Transfer application, MsgTransfer
	var rawFungibleTokenPacketData ibc_model.FungibleTokenPacketData
	rawPacketData := base64_internal.MustDecodeString(rawMsg.Packet.Data)
	json.MustUnmarshal(rawPacketData, &rawFungibleTokenPacketData)

	if !msgCommonParams.TxSuccess {
		params := ibc_model.MsgAcknowledgementParams{
			RawMsgAcknowledgement: rawMsg,

			Application: "transfer",
			MessageType: "MsgTransfer",
			MaybeFungibleTokenPacketData: &ibc_model.MsgAcknowledgementFungibleTokenPacketData{
				FungibleTokenPacketData: rawFungibleTokenPacketData,
			},
		}

		return []command.Command{command_usecase.NewCreateMsgIBCAcknowledgement(
			msgCommonParams,

			params,
		)}
	}

	log := utils.NewParsedTxsResultLog(&txsResult.Log[msgIndex])

	fungibleTokenPacketEvent := log.GetEventByType("fungible_token_packet")
	if fungibleTokenPacketEvent == nil {
		panic("missing `fungible_token_packet` event in TxsResult log")
	}

	acknowledgePacketEvent := log.GetEventByType("acknowledge_packet")
	if acknowledgePacketEvent == nil {
		panic("missing `acknowledge_packet` event in TxsResult log")
	}

	params := ibc_model.MsgAcknowledgementParams{
		RawMsgAcknowledgement: rawMsg,

		Application: "transfer",
		MessageType: "MsgTransfer",
		MaybeFungibleTokenPacketData: &ibc_model.MsgAcknowledgementFungibleTokenPacketData{
			FungibleTokenPacketData: rawFungibleTokenPacketData,
			// https://github.com/cosmos/ibc-go/blob/e98838612a4fa5d240e392aad3409db5ec428f50/modules/apps/transfer/module.go#L327
			Success: fungibleTokenPacketEvent.HasAttribute("success") &&
				bytes.Equal(
					[]byte(fungibleTokenPacketEvent.MustGetAttributeByKey("success")),
					[]byte{byte(1)},
				),
			Acknowledgement: fungibleTokenPacketEvent.MustGetAttributeByKey("acknowledgement"),
			// https://github.com/cosmos/ibc-go/blob/e98838612a4fa5d240e392aad3409db5ec428f50/modules/apps/transfer/module.go#L364
			MaybeError: fungibleTokenPacketEvent.GetAttributeByKey("error"),
		},

		PacketSequence:  typeconv.MustAtou64(acknowledgePacketEvent.MustGetAttributeByKey("packet_sequence")),
		ChannelOrdering: acknowledgePacketEvent.MustGetAttributeByKey("packet_channel_ordering"),
		ConnectionID:    acknowledgePacketEvent.MustGetAttributeByKey("packet_connection"),
	}

	return []command.Command{command_usecase.NewCreateMsgIBCAcknowledgement(
		msgCommonParams,

		params,
	)}
}

func IsPacketMsgTransfer(
	packet ibc_model.Packet,
) bool {
	if packet.DestinationPort != "transfer" {
		return false
	}

	var fungiblePacketData ibc_model.FungibleTokenPacketData
	rawPacketData, decodeDataErr := base64.StdEncoding.DecodeString(packet.Data)
	if decodeDataErr != nil {
		return false
	}
	if unmarshalErr := jsoniter.Unmarshal(rawPacketData, &fungiblePacketData); unmarshalErr != nil {
		return false
	}

	return true
}

func ParseMsgTimeout(
	msgCommonParams event.MsgCommonParams,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msg map[string]interface{},
) []command.Command {
	var rawMsg ibc_model.RawMsgTimeout
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
		panic(fmt.Errorf("error creating RawMsgTimeout decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgTimeout: %v", err))
	}

	if !IsPacketMsgTransfer(rawMsg.Packet) {
		// unsupported application
		return []command.Command{}
	}

	if !msgCommonParams.TxSuccess {
		params := ibc_model.MsgTimeoutParams{
			RawMsgTimeout: rawMsg,

			Application: "transfer",
			MessageType: "MsgTransfer",
		}

		return []command.Command{command_usecase.NewCreateMsgIBCTimeout(
			msgCommonParams,

			params,
		)}
	}

	log := utils.NewParsedTxsResultLog(&txsResult.Log[msgIndex])

	// Transfer application, MsgTransfer
	var rawFungibleTokenPacketData ibc_model.FungibleTokenPacketData
	rawPacketData := base64_internal.MustDecodeString(rawMsg.Packet.Data)
	json.MustUnmarshal(rawPacketData, &rawFungibleTokenPacketData)

	timeoutPacketEvent := log.GetEventByType("timeout_packet")
	if timeoutPacketEvent == nil {
		panic("missing `timeout_packet` event in TxsResult log")
	}

	params := ibc_model.MsgTimeoutParams{
		RawMsgTimeout: rawMsg,

		Application: "transfer",
		MessageType: "MsgTransfer",
		MaybeMsgTransfer: &ibc_model.MsgTimeoutMsgTransfer{
			RefundReceiver: rawFungibleTokenPacketData.Sender,
			RefundDenom:    rawFungibleTokenPacketData.Denom,
			RefundAmount:   rawFungibleTokenPacketData.Amount.Uint64(),
		},

		PacketTimeoutHeight: mustParseHeight(
			timeoutPacketEvent.MustGetAttributeByKey("packet_timeout_height"),
		),
		PacketTimeoutTimestamp: typeconv.MustAtou64(
			timeoutPacketEvent.MustGetAttributeByKey("packet_timeout_timestamp"),
		),

		PacketSequence:  typeconv.MustAtou64(timeoutPacketEvent.MustGetAttributeByKey("packet_sequence")),
		ChannelOrdering: timeoutPacketEvent.MustGetAttributeByKey("packet_channel_ordering"),
	}

	return []command.Command{command_usecase.NewCreateMsgIBCTimeout(
		msgCommonParams,

		params,
	)}
}

func ParseMsgTimeoutOnClose(
	msgCommonParams event.MsgCommonParams,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msg map[string]interface{},
) []command.Command {
	var rawMsg ibc_model.RawMsgTimeoutOnClose
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
		panic(fmt.Errorf("error creating RawMsgTimeoutOnClose decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgTimeoutOnClose: %v", err))
	}

	if !IsPacketMsgTransfer(rawMsg.Packet) {
		// unsupported application
		return []command.Command{}
	}

	if !msgCommonParams.TxSuccess {
		params := ibc_model.MsgTimeoutOnCloseParams{
			RawMsgTimeoutOnClose: rawMsg,

			Application: "transfer",
			MessageType: "MsgTransfer",
		}

		return []command.Command{command_usecase.NewCreateMsgIBCTimeoutOnClose(
			msgCommonParams,

			params,
		)}
	}

	log := utils.NewParsedTxsResultLog(&txsResult.Log[msgIndex])

	// Transfer application, MsgTransfer
	var rawFungibleTokenPacketData ibc_model.FungibleTokenPacketData
	rawPacketData := base64_internal.MustDecodeString(rawMsg.Packet.Data)
	json.MustUnmarshal(rawPacketData, &rawFungibleTokenPacketData)

	timeoutPacketEvent := log.GetEventByType("timeout_packet")
	if timeoutPacketEvent == nil {
		panic("missing `timeout_packet` event in TxsResult log")
	}

	params := ibc_model.MsgTimeoutOnCloseParams{
		RawMsgTimeoutOnClose: rawMsg,

		Application: "transfer",
		MessageType: "MsgTransfer",
		MaybeMsgTransfer: &ibc_model.MsgTimeoutMsgTransfer{
			RefundReceiver: rawFungibleTokenPacketData.Sender,
			RefundDenom:    rawFungibleTokenPacketData.Denom,
			RefundAmount:   rawFungibleTokenPacketData.Amount.Uint64(),
		},

		PacketTimeoutHeight: mustParseHeight(
			timeoutPacketEvent.MustGetAttributeByKey("packet_timeout_height"),
		),
		PacketTimeoutTimestamp: typeconv.MustAtou64(
			timeoutPacketEvent.MustGetAttributeByKey("packet_timeout_timestamp"),
		),

		PacketSequence:  typeconv.MustAtou64(timeoutPacketEvent.MustGetAttributeByKey("packet_sequence")),
		ChannelOrdering: timeoutPacketEvent.MustGetAttributeByKey("packet_channel_ordering"),
	}

	return []command.Command{command_usecase.NewCreateMsgIBCTimeoutOnClose(
		msgCommonParams,

		params,
	)}
}
