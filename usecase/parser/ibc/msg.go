package ibc

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/crypto-com/chain-indexing/external/json"
	jsoniter "github.com/json-iterator/go"

	"github.com/crypto-com/chain-indexing/entity/command"
	base64_internal "github.com/crypto-com/chain-indexing/internal/base64"
	"github.com/crypto-com/chain-indexing/internal/typeconv"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
	mapstructure_utils "github.com/crypto-com/chain-indexing/usecase/parser/utils/mapstructure"
)

const tendermintClientStateTypeV1 = "/ibc.lightclients.tendermint.v1.ClientState"
const soloMachineClientStateTypeV2 = "/ibc.lightclients.solomachine.v2.ClientState"
const tendermintHeaderTypeV1 = "/ibc.lightclients.tendermint.v1.Header"
const soloMachineHeaderTypeV2 = "/ibc.lightclients.solomachine.v2.Header"
const soloMachineHeaderTypeV3 = "/ibc.lightclients.solomachine.v3.Header"

func ParseMsgCreateClient(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	clientStateType := parserParams.Msg["client_state"].(map[string]interface{})["@type"]

	switch clientStateType {
	case tendermintClientStateTypeV1:
		return parseRawMsgCreateTendermintLightClient(
			parserParams.MsgCommonParams,
			parserParams.TxsResult,
			parserParams.MsgIndex,
			parserParams.Msg,
		)
	case soloMachineClientStateTypeV2:
		return parseRawMsgCreateSoloMachineLightClient(
			parserParams.MsgCommonParams,
			parserParams.TxsResult,
			parserParams.MsgIndex,
			parserParams.Msg,
		)
	default:
		return []command.Command{}, []string{}

	}
}

func parseRawMsgCreateTendermintLightClient(
	msgCommonParams event.MsgCommonParams,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msg map[string]interface{},
) ([]command.Command, []string) {
	var rawMsg ibc_model.RawMsgCreateTendermintLightClient
	if err := mapstructure_utils.DefaultMapstructureDecoder.Decode(msg, &rawMsg); err != nil {
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

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, params.Signer)

		return []command.Command{command_usecase.NewCreateMsgIBCCreateClient(
			msgCommonParams,

			params,
		)}, possibleSignerAddresses
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

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	possibleSignerAddresses = append(possibleSignerAddresses, params.Signer)

	return []command.Command{command_usecase.NewCreateMsgIBCCreateClient(
		msgCommonParams,

		params,
	)}, possibleSignerAddresses
}

func parseRawMsgCreateSoloMachineLightClient(
	msgCommonParams event.MsgCommonParams,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msg map[string]interface{},
) ([]command.Command, []string) {
	var rawMsg ibc_model.RawMsgCreateSoloMachineLightClient
	if err := mapstructure_utils.DefaultMapstructureDecoder.Decode(msg, &rawMsg); err != nil {
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

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, params.Signer)

		return []command.Command{command_usecase.NewCreateMsgIBCCreateClient(
			msgCommonParams,

			params,
		)}, possibleSignerAddresses
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

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	possibleSignerAddresses = append(possibleSignerAddresses, params.Signer)

	return []command.Command{command_usecase.NewCreateMsgIBCCreateClient(
		msgCommonParams,

		params,
	)}, possibleSignerAddresses
}

func ParseMsgConnectionOpenInit(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var rawMessage ibc_model.RawMsgConnectionOpenInit
	if err := mapstructure_utils.DefaultMapstructureDecoder.Decode(parserParams.Msg, &rawMessage); err != nil {
		panic(fmt.Errorf("error decoding message: %v", err))
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, rawMessage.Signer)

		return []command.Command{command_usecase.NewCreateMsgIBCConnectionOpenInit(
			parserParams.MsgCommonParams,

			ibc_model.MsgConnectionOpenInitParams{
				RawMsgConnectionOpenInit: rawMessage,
			},
		)}, possibleSignerAddresses
	}

	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])
	event := log.GetEventByType("connection_open_init")
	if event == nil {
		panic("missing `connection_open_init` event in TxsResult log")
	}

	if parserParams.IsEthereumTxInnerMsg {
		msgConnectionOpenInitParams := ibc_model.MsgConnectionOpenInitParams{
			RawMsgConnectionOpenInit: ibc_model.RawMsgConnectionOpenInit{
				ClientID: event.MustGetAttributeByKey("client_id"),
				Counterparty: ibc_model.ConnectionCounterparty{
					ClientID:     event.MustGetAttributeByKey("counterparty_client_id"),
					ConnectionID: event.MustGetAttributeByKey("counterparty_connection_id"),
				},
			},

			ConnectionID: event.MustGetAttributeByKey("connection_id"),
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, msgConnectionOpenInitParams.Signer)

		return []command.Command{command_usecase.NewCreateMsgIBCConnectionOpenInit(
			parserParams.MsgCommonParams,

			msgConnectionOpenInitParams,
		)}, possibleSignerAddresses
	}

	msgConnectionOpenInitParams := ibc_model.MsgConnectionOpenInitParams{
		RawMsgConnectionOpenInit: rawMessage,

		ConnectionID: event.MustGetAttributeByKey("connection_id"),
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	possibleSignerAddresses = append(possibleSignerAddresses, msgConnectionOpenInitParams.Signer)

	return []command.Command{command_usecase.NewCreateMsgIBCConnectionOpenInit(
		parserParams.MsgCommonParams,

		msgConnectionOpenInitParams,
	)}, possibleSignerAddresses
}

func ParseMsgConnectionOpenTry(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	clientState, ok := parserParams.Msg["client_state"]
	if ok {
		clientStateType := clientState.(map[string]interface{})["@type"]
		if clientStateType != "/ibc.lightclients.tendermint.v1.ClientState" {
			// TODO: SoloMachine and Localhost LightClient
			return []command.Command{}, []string{}
		}
	}

	var rawMsg ibc_model.RawMsgConnectionOpenTryTendermintClient
	if err := mapstructure_utils.DefaultMapstructureDecoder.Decode(parserParams.Msg, &rawMsg); err != nil {
		panic(fmt.Errorf("error decoding MsgConnectionOpenTry: %v", err))
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		msgConnectionOpenTryParams := ibc_model.MsgConnectionOpenTryParams{
			MsgConnectionOpenTryBaseParams: rawMsg.MsgConnectionOpenTryBaseParams,
			MaybeTendermintClientState:     &rawMsg.TendermintClientState,
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, msgConnectionOpenTryParams.Signer)

		return []command.Command{command_usecase.NewCreateMsgIBCConnectionOpenTry(
			parserParams.MsgCommonParams,

			msgConnectionOpenTryParams,
		)}, possibleSignerAddresses
	}

	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])
	event := log.GetEventByType("connection_open_try")
	if event == nil {
		panic("missing `connection_open_try` event in TxsResult log")
	}

	if parserParams.IsEthereumTxInnerMsg {
		msgConnectionOpenTryParams := ibc_model.MsgConnectionOpenTryParams{
			MsgConnectionOpenTryBaseParams: ibc_model.MsgConnectionOpenTryBaseParams{
				ClientID: event.MustGetAttributeByKey("client_id"),
				Counterparty: ibc_model.ConnectionCounterparty{
					ClientID:     event.MustGetAttributeByKey("counterparty_client_id"),
					ConnectionID: event.MustGetAttributeByKey("counterparty_connection_id"),
				},
			},

			ConnectionID: event.MustGetAttributeByKey("connection_id"),
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, msgConnectionOpenTryParams.Signer)

		return []command.Command{command_usecase.NewCreateMsgIBCConnectionOpenTry(
			parserParams.MsgCommonParams,

			msgConnectionOpenTryParams,
		)}, possibleSignerAddresses
	}

	msgConnectionOpenTryParams := ibc_model.MsgConnectionOpenTryParams{
		MsgConnectionOpenTryBaseParams: rawMsg.MsgConnectionOpenTryBaseParams,
		MaybeTendermintClientState:     &rawMsg.TendermintClientState,
		ConnectionID:                   event.MustGetAttributeByKey("connection_id"),
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	possibleSignerAddresses = append(possibleSignerAddresses, msgConnectionOpenTryParams.Signer)

	return []command.Command{command_usecase.NewCreateMsgIBCConnectionOpenTry(
		parserParams.MsgCommonParams,

		msgConnectionOpenTryParams,
	)}, possibleSignerAddresses
}

func ParseMsgConnectionOpenAck(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	clientState, ok := parserParams.Msg["client_state"]
	if ok {
		clientStateType := clientState.(map[string]interface{})["@type"]
		if clientStateType != "/ibc.lightclients.tendermint.v1.ClientState" {
			// TODO: SoloMachine and Localhost LightClient
			return []command.Command{}, []string{}
		}
	}

	var rawMsg ibc_model.RawMsgConnectionOpenAckTendermintClient
	if err := mapstructure_utils.DefaultMapstructureDecoder.Decode(parserParams.Msg, &rawMsg); err != nil {
		panic(fmt.Errorf("error decoding MsgConnectionOpenAck: %v", err))
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		msgConnectionOpenAckParams := ibc_model.MsgConnectionOpenAckParams{
			MsgConnectionOpenAckBaseParams: rawMsg.MsgConnectionOpenAckBaseParams,
			MaybeTendermintClientState:     &rawMsg.TendermintClientState,
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, msgConnectionOpenAckParams.Signer)

		return []command.Command{command_usecase.NewCreateMsgIBCConnectionOpenAck(
			parserParams.MsgCommonParams,

			msgConnectionOpenAckParams,
		)}, possibleSignerAddresses
	}

	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])
	event := log.GetEventByType("connection_open_ack")
	if event == nil {
		panic("missing `connection_open_ack` event in TxsResult log")
	}

	if parserParams.IsEthereumTxInnerMsg {
		msgConnectionOpenAckParams := ibc_model.MsgConnectionOpenAckParams{
			MsgConnectionOpenAckBaseParams: ibc_model.MsgConnectionOpenAckBaseParams{
				ConnectionID:             event.MustGetAttributeByKey("connection_id"),
				CounterpartyConnectionID: event.MustGetAttributeByKey("counterparty_connection_id"),
			},

			ClientID:             event.MustGetAttributeByKey("client_id"),
			CounterpartyClientID: event.MustGetAttributeByKey("counterparty_client_id"),
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, msgConnectionOpenAckParams.Signer)

		return []command.Command{command_usecase.NewCreateMsgIBCConnectionOpenAck(
			parserParams.MsgCommonParams,

			msgConnectionOpenAckParams,
		)}, possibleSignerAddresses
	}

	msgConnectionOpenAckParams := ibc_model.MsgConnectionOpenAckParams{
		MsgConnectionOpenAckBaseParams: rawMsg.MsgConnectionOpenAckBaseParams,
		MaybeTendermintClientState:     &rawMsg.TendermintClientState,

		ClientID:             event.MustGetAttributeByKey("client_id"),
		CounterpartyClientID: event.MustGetAttributeByKey("counterparty_client_id"),
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	possibleSignerAddresses = append(possibleSignerAddresses, msgConnectionOpenAckParams.Signer)

	return []command.Command{command_usecase.NewCreateMsgIBCConnectionOpenAck(
		parserParams.MsgCommonParams,

		msgConnectionOpenAckParams,
	)}, possibleSignerAddresses
}

func ParseMsgConnectionOpenConfirm(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var rawMsg ibc_model.RawMsgConnectionOpenConfirm
	if err := mapstructure_utils.DefaultMapstructureDecoder.Decode(parserParams.Msg, &rawMsg); err != nil {
		panic(fmt.Errorf("error decoding MsgConnectionOpenConfirm: %v", err))
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, rawMsg.Signer)

		return []command.Command{command_usecase.NewCreateMsgIBCConnectionOpenConfirm(
			parserParams.MsgCommonParams,

			ibc_model.MsgConnectionOpenConfirmParams{
				RawMsgConnectionOpenConfirm: rawMsg,
			},
		)}, possibleSignerAddresses
	}

	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])
	event := log.GetEventByType("connection_open_confirm")
	if event == nil {
		panic("missing `connection_open_confirm` event in TxsResult log")
	}

	if parserParams.IsEthereumTxInnerMsg {
		msgConnectionOpenConfirmParams := ibc_model.MsgConnectionOpenConfirmParams{
			RawMsgConnectionOpenConfirm: ibc_model.RawMsgConnectionOpenConfirm{
				ConnectionID: event.MustGetAttributeByKey("connection_id"),
			},

			ClientID:                 event.MustGetAttributeByKey("client_id"),
			CounterpartyClientID:     event.MustGetAttributeByKey("counterparty_client_id"),
			CounterpartyConnectionID: event.MustGetAttributeByKey("counterparty_connection_id"),
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, msgConnectionOpenConfirmParams.Signer)

		return []command.Command{command_usecase.NewCreateMsgIBCConnectionOpenConfirm(
			parserParams.MsgCommonParams,

			msgConnectionOpenConfirmParams,
		)}, possibleSignerAddresses
	}

	msgConnectionOpenConfirmParams := ibc_model.MsgConnectionOpenConfirmParams{
		RawMsgConnectionOpenConfirm: rawMsg,

		ClientID:                 event.MustGetAttributeByKey("client_id"),
		CounterpartyClientID:     event.MustGetAttributeByKey("counterparty_client_id"),
		CounterpartyConnectionID: event.MustGetAttributeByKey("counterparty_connection_id"),
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	possibleSignerAddresses = append(possibleSignerAddresses, msgConnectionOpenConfirmParams.Signer)

	return []command.Command{command_usecase.NewCreateMsgIBCConnectionOpenConfirm(
		parserParams.MsgCommonParams,

		msgConnectionOpenConfirmParams,
	)}, possibleSignerAddresses
}

func ParseMsgChannelOpenInit(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var rawMsg ibc_model.RawMsgChannelOpenInit
	if err := mapstructure_utils.DefaultMapstructureDecoder.Decode(parserParams.Msg, &rawMsg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgChannelOpenInit: %v", err))
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, rawMsg.Signer)

		return []command.Command{command_usecase.NewCreateMsgIBCChannelOpenInit(
			parserParams.MsgCommonParams,

			ibc_model.MsgChannelOpenInitParams{
				RawMsgChannelOpenInit: rawMsg,
			},
		)}, possibleSignerAddresses
	}

	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])
	if parserParams.IsEthereumTxInnerMsg {
		events := log.GetEventsByType("channel_open_init")
		for _, event := range events {
			msgChannelOpenInitParams := ibc_model.MsgChannelOpenInitParams{
				RawMsgChannelOpenInit: ibc_model.RawMsgChannelOpenInit{
					PortID: event.MustGetAttributeByKey("port_id"),
					Channel: ibc_model.Channel{
						Counterparty: ibc_model.ChannelCounterparty{
							PortID:    event.MustGetAttributeByKey("counterparty_port_id"),
							ChannelID: event.MustGetAttributeByKey("counterparty_channel_id"),
						},
					},
				},

				ChannelID:    event.MustGetAttributeByKey("channel_id"),
				ConnectionID: event.MustGetAttributeByKey("connection_id"),
			}
			// Getting possible signer address from Msg
			var possibleSignerAddresses []string
			possibleSignerAddresses = append(possibleSignerAddresses, msgChannelOpenInitParams.Signer)

			return []command.Command{command_usecase.NewCreateMsgIBCChannelOpenInit(
				parserParams.MsgCommonParams,

				msgChannelOpenInitParams,
			)}, possibleSignerAddresses
		}
	}

	event := log.GetEventByType("channel_open_init")
	if event == nil {
		panic("missing `channel_open_init` event in TxsResult log")
	}

	msgChannelOpenInitParams := ibc_model.MsgChannelOpenInitParams{
		RawMsgChannelOpenInit: rawMsg,

		ChannelID:    event.MustGetAttributeByKey("channel_id"),
		ConnectionID: event.MustGetAttributeByKey("connection_id"),
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	possibleSignerAddresses = append(possibleSignerAddresses, msgChannelOpenInitParams.Signer)

	return []command.Command{command_usecase.NewCreateMsgIBCChannelOpenInit(
		parserParams.MsgCommonParams,

		msgChannelOpenInitParams,
	)}, possibleSignerAddresses
}

func ParseMsgChannelOpenTry(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var rawMsg ibc_model.RawMsgChannelOpenTry
	if err := mapstructure_utils.DefaultMapstructureDecoder.Decode(parserParams.Msg, &rawMsg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgChannelOpenTry: %v", err))
	}

	if !parserParams.MsgCommonParams.TxSuccess {

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, rawMsg.Signer)

		return []command.Command{command_usecase.NewCreateMsgIBCChannelOpenTry(
			parserParams.MsgCommonParams,

			ibc_model.MsgChannelOpenTryParams{
				RawMsgChannelOpenTry: rawMsg,
			},
		)}, possibleSignerAddresses
	}

	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])
	event := log.GetEventByType("channel_open_try")
	if event == nil {
		panic("missing `channel_open_try` event in TxsResult log")
	}

	if parserParams.IsEthereumTxInnerMsg {
		msgChannelOpenTryParams := ibc_model.MsgChannelOpenTryParams{
			RawMsgChannelOpenTry: ibc_model.RawMsgChannelOpenTry{
				PortID: event.MustGetAttributeByKey("port_id"),
				Channel: ibc_model.Channel{
					Counterparty: ibc_model.ChannelCounterparty{
						PortID:    event.MustGetAttributeByKey("counterparty_port_id"),
						ChannelID: event.MustGetAttributeByKey("counterparty_channel_id"),
					},
				},
			},

			ChannelID:    event.MustGetAttributeByKey("channel_id"),
			ConnectionID: event.MustGetAttributeByKey("connection_id"),
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, msgChannelOpenTryParams.Signer)

		return []command.Command{command_usecase.NewCreateMsgIBCChannelOpenTry(
			parserParams.MsgCommonParams,

			msgChannelOpenTryParams,
		)}, possibleSignerAddresses
	}

	msgChannelOpenTryParams := ibc_model.MsgChannelOpenTryParams{
		RawMsgChannelOpenTry: rawMsg,

		ChannelID:    event.MustGetAttributeByKey("channel_id"),
		ConnectionID: event.MustGetAttributeByKey("connection_id"),
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	possibleSignerAddresses = append(possibleSignerAddresses, msgChannelOpenTryParams.Signer)

	return []command.Command{command_usecase.NewCreateMsgIBCChannelOpenTry(
		parserParams.MsgCommonParams,

		msgChannelOpenTryParams,
	)}, possibleSignerAddresses
}

func ParseMsgChannelOpenAck(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var rawMsg ibc_model.RawMsgChannelOpenAck
	if err := mapstructure_utils.DefaultMapstructureDecoder.Decode(parserParams.Msg, &rawMsg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgChannelOpenAck: %v", err))
	}

	if !parserParams.MsgCommonParams.TxSuccess {

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, rawMsg.Signer)

		return []command.Command{command_usecase.NewCreateMsgIBCChannelOpenAck(
			parserParams.MsgCommonParams,

			ibc_model.MsgChannelOpenAckParams{
				RawMsgChannelOpenAck: rawMsg,
			},
		)}, possibleSignerAddresses
	}

	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])
	event := log.GetEventByType("channel_open_ack")
	if event == nil {
		panic("missing `channel_open_ack` event in TxsResult log")
	}

	if parserParams.IsEthereumTxInnerMsg {
		msgChannelOpenAckParams := ibc_model.MsgChannelOpenAckParams{
			RawMsgChannelOpenAck: ibc_model.RawMsgChannelOpenAck{
				PortID:                event.MustGetAttributeByKey("port_id"),
				ChannelID:             event.MustGetAttributeByKey("channel_id"),
				CounterpartyChannelID: event.MustGetAttributeByKey("counterparty_channel_id"),
			},

			CounterpartyPortID: event.MustGetAttributeByKey("counterparty_port_id"),
			ConnectionID:       event.MustGetAttributeByKey("connection_id"),
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, msgChannelOpenAckParams.Signer)

		return []command.Command{command_usecase.NewCreateMsgIBCChannelOpenAck(
			parserParams.MsgCommonParams,

			msgChannelOpenAckParams,
		)}, possibleSignerAddresses
	}

	msgChannelOpenAckParams := ibc_model.MsgChannelOpenAckParams{
		RawMsgChannelOpenAck: rawMsg,

		CounterpartyPortID: event.MustGetAttributeByKey("counterparty_port_id"),
		ConnectionID:       event.MustGetAttributeByKey("connection_id"),
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	possibleSignerAddresses = append(possibleSignerAddresses, msgChannelOpenAckParams.Signer)

	return []command.Command{command_usecase.NewCreateMsgIBCChannelOpenAck(
		parserParams.MsgCommonParams,

		msgChannelOpenAckParams,
	)}, possibleSignerAddresses
}

func ParseMsgChannelOpenConfirm(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var rawMsg ibc_model.RawMsgChannelOpenConfirm
	if err := mapstructure_utils.DefaultMapstructureDecoder.Decode(parserParams.Msg, &rawMsg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgChannelOpenConfirm: %v", err))
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, rawMsg.Signer)

		return []command.Command{command_usecase.NewCreateMsgIBCChannelOpenConfirm(
			parserParams.MsgCommonParams,

			ibc_model.MsgChannelOpenConfirmParams{
				RawMsgChannelOpenConfirm: rawMsg,
			},
		)}, possibleSignerAddresses
	}

	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])
	event := log.GetEventByType("channel_open_confirm")
	if event == nil {
		panic("missing `channel_open_confirm` event in TxsResult log")
	}

	if rawMsg.PortID == "" {
		rawMsg.PortID = event.MustGetAttributeByKey("port_id")
	}
	if rawMsg.ChannelID == "" {
		rawMsg.ChannelID = event.MustGetAttributeByKey("channel_id")
	}

	msgChannelOpenConfirmParams := ibc_model.MsgChannelOpenConfirmParams{
		RawMsgChannelOpenConfirm: rawMsg,

		CounterpartyChannelID: event.MustGetAttributeByKey("counterparty_channel_id"),
		CounterpartyPortID:    event.MustGetAttributeByKey("counterparty_port_id"),
		ConnectionID:          event.MustGetAttributeByKey("connection_id"),
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	possibleSignerAddresses = append(possibleSignerAddresses, msgChannelOpenConfirmParams.Signer)

	return []command.Command{command_usecase.NewCreateMsgIBCChannelOpenConfirm(
		parserParams.MsgCommonParams,

		msgChannelOpenConfirmParams,
	)}, possibleSignerAddresses
}

func ParseMsgUpdateClient(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var headerType interface{}
	if header, ok := parserParams.Msg["header"].(map[string]interface{}); ok {
		headerType = header["@type"]
	} else if header, ok := parserParams.Msg["client_message"].(map[string]interface{}); ok {
		headerType = header["@type"]
	}

	switch headerType {
	case tendermintHeaderTypeV1:
		return parseMsgUpdateTendermintLightClient(
			parserParams.MsgCommonParams,
			parserParams.TxsResult,
			parserParams.MsgIndex,
			parserParams.Msg,
		)
	case soloMachineHeaderTypeV2:
		return parseMsgUpdateSolomachineLightClient(
			parserParams.MsgCommonParams,
			parserParams.TxsResult,
			parserParams.MsgIndex,
			parserParams.Msg,
		)
	case soloMachineHeaderTypeV3:
		return parseMsgUpdateSolomachineLightClient(
			parserParams.MsgCommonParams,
			parserParams.TxsResult,
			parserParams.MsgIndex,
			parserParams.Msg,
		)
	default:
		return []command.Command{}, []string{}

	}
}

func parseMsgUpdateTendermintLightClient(
	msgCommonParams event.MsgCommonParams,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msg map[string]interface{},
) ([]command.Command, []string) {
	var rawMsg ibc_model.RawMsgUpdateTendermintLightClient
	if err := mapstructure_utils.DefaultMapstructureDecoder.Decode(msg, &rawMsg); err != nil {
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

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, params.Signer)

		return []command.Command{command_usecase.NewCreateMsgIBCUpdateClient(
			msgCommonParams,

			params,
		)}, possibleSignerAddresses
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
		ConsensusHeight: MustParseHeight(event.MustGetAttributeByKey("consensus_height")),
		Signer:          rawMsg.Signer,
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	possibleSignerAddresses = append(possibleSignerAddresses, params.Signer)

	return []command.Command{command_usecase.NewCreateMsgIBCUpdateClient(
		msgCommonParams,

		params,
	)}, possibleSignerAddresses
}

func parseMsgUpdateSolomachineLightClient(
	msgCommonParams event.MsgCommonParams,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msg map[string]interface{},
) ([]command.Command, []string) {
	var rawMsg ibc_model.RawMsgUpdateSoloMachineLightClient
	if err := mapstructure_utils.DefaultMapstructureDecoder.Decode(msg, &rawMsg); err != nil {
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

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, params.Signer)

		return []command.Command{command_usecase.NewCreateMsgIBCUpdateClient(
			msgCommonParams,

			params,
		)}, possibleSignerAddresses
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
		ConsensusHeight: MustParseHeight(event.MustGetAttributeByKey("consensus_height")),
		Signer:          rawMsg.Signer,
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	possibleSignerAddresses = append(possibleSignerAddresses, params.Signer)

	return []command.Command{command_usecase.NewCreateMsgIBCUpdateClient(
		msgCommonParams,

		params,
	)}, possibleSignerAddresses
}

func MustParseHeight(height string) ibc_model.Height {
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
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var rawMsg ibc_model.RawMsgTransfer
	if err := mapstructure_utils.DefaultMapstructureDecoder.Decode(parserParams.Msg, &rawMsg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgTransfer: %v", err))
	}

	if parserParams.IsProposalInnerMsg {
		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, rawMsg.Sender)

		return []command.Command{command_usecase.NewCreateMsgIBCTransferTransfer(
			parserParams.MsgCommonParams,
			ibc_model.MsgTransferParams{
				RawMsgTransfer: rawMsg,
			},
		)}, possibleSignerAddresses
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, rawMsg.Sender)

		return []command.Command{command_usecase.NewCreateMsgIBCTransferTransfer(
			parserParams.MsgCommonParams,
			ibc_model.MsgTransferParams{
				RawMsgTransfer: rawMsg,
			},
		)}, possibleSignerAddresses
	}

	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])

	if parserParams.IsEthereumTxInnerMsg {
		events := log.GetEventsByType("send_packet")
		for _, event := range events {
			// Transfer application, MsgTransfer
			packetData := event.MustGetAttributeByKey("packet_data")

			var fungiblePacketData *ibc_model.FungibleTokenPacketData
			if unmarshalErr := jsoniter.Unmarshal([]byte(packetData), &fungiblePacketData); unmarshalErr != nil {
				panic("unable to parse `send_packet` event, key `packet_data`")
			}

			if fungiblePacketData != nil {
				msgTransferParams := ibc_model.MsgTransferParams{
					RawMsgTransfer: ibc_model.RawMsgTransfer{
						Token: ibc_model.MsgTransferToken{
							Denom:  fungiblePacketData.Denom,
							Amount: fungiblePacketData.Amount,
						},
						Sender:           fungiblePacketData.Sender,
						Receiver:         fungiblePacketData.Receiver,
						SourcePort:       event.MustGetAttributeByKey("packet_src_port"),
						SourceChannel:    event.MustGetAttributeByKey("packet_src_channel"),
						TimeoutHeight:    MustParseHeight(event.MustGetAttributeByKey("packet_timeout_height")),
						TimeoutTimestamp: event.MustGetAttributeByKey("packet_timeout_timestamp"),
					},

					PacketSequence:     typeconv.MustAtou64(event.MustGetAttributeByKey("packet_sequence")),
					DestinationPort:    event.MustGetAttributeByKey("packet_dst_port"),
					DestinationChannel: event.MustGetAttributeByKey("packet_dst_channel"),
					ChannelOrdering:    event.MustGetAttributeByKey("packet_channel_ordering"),
					ConnectionID:       event.MustGetAttributeByKey("packet_connection"),
					PacketData:         *fungiblePacketData,
				}

				// Getting possible signer address from Msg
				var possibleSignerAddresses []string
				possibleSignerAddresses = append(possibleSignerAddresses, msgTransferParams.Sender)

				return []command.Command{command_usecase.NewCreateMsgIBCTransferTransfer(
					parserParams.MsgCommonParams,

					msgTransferParams,
				)}, possibleSignerAddresses
			}
		}
	}

	event := log.GetEventByType("send_packet")
	if event == nil {
		panic("missing `send_packet` event in TxsResult log")
	}

	packetData := event.MustGetAttributeByKey("packet_data")
	var fungiblePacketData ibc_model.FungibleTokenPacketData
	if unmarshalErr := jsoniter.Unmarshal([]byte(packetData), &fungiblePacketData); unmarshalErr != nil {
		panic("unable to parse `send_packet` event, key `packet_data`")
	}

	msgTransferParams := ibc_model.MsgTransferParams{
		RawMsgTransfer: rawMsg,

		PacketSequence:     typeconv.MustAtou64(event.MustGetAttributeByKey("packet_sequence")),
		DestinationPort:    event.MustGetAttributeByKey("packet_dst_port"),
		DestinationChannel: event.MustGetAttributeByKey("packet_dst_channel"),
		ChannelOrdering:    event.MustGetAttributeByKey("packet_channel_ordering"),
		ConnectionID:       event.MustGetAttributeByKey("packet_connection"),
		PacketData:         fungiblePacketData,
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	possibleSignerAddresses = append(possibleSignerAddresses, msgTransferParams.Sender)

	return []command.Command{command_usecase.NewCreateMsgIBCTransferTransfer(
		parserParams.MsgCommonParams,

		msgTransferParams,
	)}, possibleSignerAddresses
}

func ParseMsgRecvPacket(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var rawMsg ibc_model.RawMsgRecvPacket
	if err := mapstructure_utils.DefaultMapstructureDecoder.Decode(parserParams.Msg, &rawMsg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgRecvPacket: %v", err))
	}

	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])

	recvPacketEvent := log.GetEventByType("recv_packet")
	if recvPacketEvent == nil {
		parserParams.Logger.Warnf("missing `recv_packet` event in TxsResult log on TxHash: %s", parserParams.MsgCommonParams.TxHash)
	}

	// Transfer application, MsgTransfer
	var rawFungibleTokenPacketData ibc_model.FungibleTokenPacketData
	rawPacketData, err := base64_internal.DecodeString(rawMsg.Packet.Data)
	if err != nil {
		rawFungibleTokenPacketData = ibc_model.FungibleTokenPacketData{}
	} else {
		if err = json.Unmarshal(rawPacketData, &rawFungibleTokenPacketData); err != nil {
			// try to get `packet` data from `recv_packet` events
			if recvPacketEvent.HasAttribute("packet_data_hex") {
				packetDataHex := recvPacketEvent.MustGetAttributeByKey("packet_data_hex")
				var packetData []byte
				packetData, err = hex.DecodeString(packetDataHex)
				if err == nil {
					if err = json.Unmarshal(packetData, &rawFungibleTokenPacketData); err != nil {
						parserParams.Logger.Warnf("error unmarshalling packet data hex")
					}
				}
			} else {
				rawFungibleTokenPacketData = ibc_model.FungibleTokenPacketData{}
			}
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

	if !IsPacketMsgTransfer(rawMsg.Packet, &parserParams.TxsResult.Log[parserParams.MsgIndex]) {
		// unsupported application
		return []command.Command{}, []string{}
	}

	var packetSequence uint64
	var channelOrdering string
	var connectionID string

	if recvPacketEvent.HasAttribute("packet_sequence") {
		sequence := recvPacketEvent.MustGetAttributeByKey("packet_sequence")
		rawMsg.Packet.Sequence = sequence
		packetSequence = typeconv.MustAtou64(sequence)
	}
	if recvPacketEvent.HasAttribute("packet_channel_ordering") {
		channelOrdering = recvPacketEvent.MustGetAttributeByKey("packet_channel_ordering")
	}
	if recvPacketEvent.HasAttribute("packet_connection") {
		connectionID = recvPacketEvent.MustGetAttributeByKey("packet_connection")
	}
	if recvPacketEvent.HasAttribute("packet_timeout_height") {
		packetTimeoutHeight := recvPacketEvent.MustGetAttributeByKey("packet_timeout_height")
		var revisionHeight uint64
		revisionHeight, err = strconv.ParseUint(packetTimeoutHeight, 10, 64)
		if err == nil {
			rawMsg.Packet.TimeoutHeight.RevisionHeight = revisionHeight
		}
	}
	if recvPacketEvent.HasAttribute("packet_timeout_timestamp") {
		rawMsg.Packet.TimeoutTimestamp = recvPacketEvent.MustGetAttributeByKey("packet_timeout_timestamp")
	}
	if recvPacketEvent.HasAttribute("packet_sequence") {
		rawMsg.Packet.Sequence = recvPacketEvent.MustGetAttributeByKey("packet_sequence")
	}
	if recvPacketEvent.HasAttribute("packet_src_port") {
		rawMsg.Packet.SourcePort = recvPacketEvent.MustGetAttributeByKey("packet_src_port")
	}
	if recvPacketEvent.HasAttribute("packet_src_channel") {
		rawMsg.Packet.SourceChannel = recvPacketEvent.MustGetAttributeByKey("packet_src_channel")
	}
	if recvPacketEvent.HasAttribute("packet_dst_port") {
		rawMsg.Packet.DestinationPort = recvPacketEvent.MustGetAttributeByKey("packet_dst_port")
	}
	if recvPacketEvent.HasAttribute("packet_dst_channel") {
		rawMsg.Packet.DestinationChannel = recvPacketEvent.MustGetAttributeByKey("packet_dst_channel")
	}

	fungibleTokenPacketEvent := log.GetEventByType("fungible_token_packet")
	if fungibleTokenPacketEvent == nil {
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

	var maybeDenominationTrace *ibc_model.MsgRecvPacketFungibleTokenDenominationTrace
	denominationTraceEvent := log.GetEventByType("denomination_trace")
	if denominationTraceEvent != nil {
		maybeDenominationTrace = &ibc_model.MsgRecvPacketFungibleTokenDenominationTrace{
			Hash:  denominationTraceEvent.MustGetAttributeByKey("trace_hash"),
			Denom: denominationTraceEvent.MustGetAttributeByKey("denom"),
		}
	}

	writeAckEvent := log.GetEventByType("write_acknowledgement")
	var packetAck ibc_model.MsgRecvPacketPacketAck
	if writeAckEvent != nil {
		json.MustUnmarshalFromString(writeAckEvent.MustGetAttributeByKey("packet_ack"), &packetAck)
	}

	msgRecvPacketParams := ibc_model.MsgRecvPacketParams{
		RawMsgRecvPacket: rawMsg,

		Application: "transfer",
		MessageType: "/ibc.applications.transfer.v1.MsgTransfer",
		MaybeFungibleTokenPacketData: &ibc_model.MsgRecvPacketFungibleTokenPacketData{
			FungibleTokenPacketData: rawFungibleTokenPacketData,
			Success:                 fungibleTokenPacketEvent.MustGetAttributeByKey("success") == "true",
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

func ParseMsgAcknowledgement(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var rawMsg ibc_model.RawMsgAcknowledgement
	if err := mapstructure_utils.DefaultMapstructureDecoder.Decode(parserParams.Msg, &rawMsg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgAcknowledgement: %v", err))
	}

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
		msgAcknowledgementParams := ibc_model.MsgAcknowledgementParams{
			RawMsgAcknowledgement: rawMsg,

			Application: "transfer",
			MessageType: "/ibc.applications.transfer.v1.MsgTransfer",
			MaybeFungibleTokenPacketData: &ibc_model.MsgAcknowledgementFungibleTokenPacketData{
				FungibleTokenPacketData: rawFungibleTokenPacketData,
				Success:                 false,
			},
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, msgAcknowledgementParams.Signer)

		return []command.Command{command_usecase.NewCreateMsgIBCAcknowledgement(
			parserParams.MsgCommonParams,

			msgAcknowledgementParams,
		)}, possibleSignerAddresses
	}

	if !IsPacketMsgTransfer(rawMsg.Packet, &parserParams.TxsResult.Log[parserParams.MsgIndex]) {
		// unsupported application
		return []command.Command{}, []string{}
	}

	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])

	acknowledgePacketEvents := log.GetEventsByType("acknowledge_packet")

	var packetSequence uint64
	var channelOrdering string
	var connectionID string
	for _, acknowledgePacketEvent := range acknowledgePacketEvents {
		if acknowledgePacketEvent.HasAttribute("packet_sequence") {
			sequence := acknowledgePacketEvent.MustGetAttributeByKey("packet_sequence")
			rawMsg.Packet.Sequence = sequence
			packetSequence = typeconv.MustAtou64(sequence)
		}
		if acknowledgePacketEvent.HasAttribute("packet_channel_ordering") {
			channelOrdering = acknowledgePacketEvent.MustGetAttributeByKey("packet_channel_ordering")
		}
		if acknowledgePacketEvent.HasAttribute("packet_connection") {
			connectionID = acknowledgePacketEvent.MustGetAttributeByKey("packet_connection")
		}
		if acknowledgePacketEvent.HasAttribute("packet_src_channel") {
			rawMsg.Packet.SourceChannel = acknowledgePacketEvent.MustGetAttributeByKey("packet_src_channel")
		}
		if acknowledgePacketEvent.HasAttribute("packet_src_port") {
			rawMsg.Packet.SourcePort = acknowledgePacketEvent.MustGetAttributeByKey("packet_src_port")
		}
		if acknowledgePacketEvent.HasAttribute("packet_dst_port") {
			rawMsg.Packet.DestinationPort = acknowledgePacketEvent.MustGetAttributeByKey("packet_dst_port")
		}
		if acknowledgePacketEvent.HasAttribute("packet_dst_channel") {
			rawMsg.Packet.DestinationChannel = acknowledgePacketEvent.MustGetAttributeByKey("packet_dst_channel")
		}
	}

	fungibleTokenPacketEvents := log.GetEventsByType("fungible_token_packet")
	if fungibleTokenPacketEvents == nil {
		msgAlreadyRelayedAcknowledgementParams := ibc_model.MsgAlreadyRelayedAcknowledgementParams{
			RawMsgAcknowledgement: rawMsg,

			Application: "transfer",
			MessageType: "/ibc.applications.transfer.v1.MsgTransfer",
			MaybeFungibleTokenPacketData: &ibc_model.MsgAlreadyRelayedAcknowledgementFungibleTokenPacketData{
				FungibleTokenPacketData: rawFungibleTokenPacketData,
			},

			PacketSequence:  packetSequence,
			ChannelOrdering: channelOrdering,
			ConnectionID:    connectionID,
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, msgAlreadyRelayedAcknowledgementParams.Signer)

		return []command.Command{command_usecase.NewCreateMsgAlreadyRelayedIBCAcknowledgement(
			parserParams.MsgCommonParams,

			msgAlreadyRelayedAcknowledgementParams,
		)}, possibleSignerAddresses
	}

	var success bool
	var acknowledgement string
	var maybeErr *string
	for _, fungibleTokenPacketEvent := range fungibleTokenPacketEvents {
		if fungibleTokenPacketEvent.HasAttribute("success") {
			success = bytes.Equal(
				[]byte(fungibleTokenPacketEvent.MustGetAttributeByKey("success")),
				[]byte{byte(1)},
			)
		}
		if fungibleTokenPacketEvent.HasAttribute("sender") {
			rawFungibleTokenPacketData.Sender = fungibleTokenPacketEvent.MustGetAttributeByKey("sender")
		}
		if fungibleTokenPacketEvent.HasAttribute("receiver") {
			rawFungibleTokenPacketData.Receiver = fungibleTokenPacketEvent.MustGetAttributeByKey("receiver")
		}
		if fungibleTokenPacketEvent.HasAttribute("denom") {
			rawFungibleTokenPacketData.Denom = fungibleTokenPacketEvent.MustGetAttributeByKey("denom")
		}
		if fungibleTokenPacketEvent.HasAttribute("amount") {
			err := json.Unmarshal([]byte(fungibleTokenPacketEvent.MustGetAttributeByKey("amount")), &rawFungibleTokenPacketData.Amount)
			if err != nil {
				parserParams.Logger.Errorf("error unmarshalling amount", err)
			}
		}
		if fungibleTokenPacketEvent.HasAttribute("acknowledgement") {
			acknowledgement = fungibleTokenPacketEvent.MustGetAttributeByKey("acknowledgement")
		}
		if fungibleTokenPacketEvent.HasAttribute("error") {
			maybeErr = fungibleTokenPacketEvent.GetAttributeByKey("error")
		}
	}

	msgAcknowledgementParams := ibc_model.MsgAcknowledgementParams{
		RawMsgAcknowledgement: rawMsg,

		Application: "transfer",
		MessageType: "/ibc.applications.transfer.v1.MsgTransfer",
		MaybeFungibleTokenPacketData: &ibc_model.MsgAcknowledgementFungibleTokenPacketData{
			FungibleTokenPacketData: rawFungibleTokenPacketData,
			// https://github.com/cosmos/ibc-go/blob/e98838612a4fa5d240e392aad3409db5ec428f50/modules/apps/transfer/module.go#L327
			Success:         success,
			Acknowledgement: acknowledgement,
			// https://github.com/cosmos/ibc-go/blob/e98838612a4fa5d240e392aad3409db5ec428f50/modules/apps/transfer/module.go#L364
			MaybeError: maybeErr,
		},

		PacketSequence:  packetSequence,
		ChannelOrdering: channelOrdering,
		ConnectionID:    connectionID,
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	possibleSignerAddresses = append(possibleSignerAddresses, msgAcknowledgementParams.Signer)

	return []command.Command{command_usecase.NewCreateMsgIBCAcknowledgement(
		parserParams.MsgCommonParams,

		msgAcknowledgementParams,
	)}, possibleSignerAddresses
}

func IsPacketMsgTransfer(
	packet ibc_model.Packet,
	txsResultLog *model.BlockResultsTxsResultLog,
) bool {
	log := utils.NewParsedTxsResultLog(txsResultLog)

	// check for evm msgAcknowledgement
	acknowledgePacketEvents := log.GetEventsByType("acknowledge_packet")
	for _, acknowledgePacketEvent := range acknowledgePacketEvents {
		if acknowledgePacketEvent.HasAttribute("packet_dst_port") {
			if acknowledgePacketEvent.MustGetAttributeByKey("packet_dst_port") == "transfer" {
				return true
			}
		}
	}

	// check for evm msgTimeout
	timeoutPacketEvent := log.GetEventByType("timeout_packet")
	if timeoutPacketEvent != nil {
		if timeoutPacketEvent.HasAttribute("packet_dst_port") {
			if timeoutPacketEvent.MustGetAttributeByKey("packet_dst_port") == "transfer" {
				return true
			}
		}
	}

	// check for evm msgRecvPacket
	recvPacketEvents := log.GetEventsByType("recv_packet")
	if len(recvPacketEvents) > 0 {
		for _, recvPacketEvent := range recvPacketEvents {
			if recvPacketEvent.HasAttribute("packet_dst_port") {
				if recvPacketEvent.MustGetAttributeByKey("packet_dst_port") == "transfer" {
					return true
				}
			}
		}
	}

	// check for cosmos layer msgAcknowledgement
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

func IsPacketMsgSubmitTx(
	packet ibc_model.Packet,
) bool {
	if packet.DestinationPort != "icahost" {
		return false
	}

	var rawInterchainAccountPacketData ibc_model.InterchainAccountPacketData
	rawPacketData, decodeDataErr := base64.StdEncoding.DecodeString(packet.Data)
	if decodeDataErr != nil {
		return false
	}
	if unmarshalErr := jsoniter.Unmarshal(rawPacketData, &rawInterchainAccountPacketData); unmarshalErr != nil {
		return false
	}

	return true
}

func ParseMsgTimeout(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var rawMsg ibc_model.RawMsgTimeout
	if err := mapstructure_utils.DefaultMapstructureDecoder.Decode(parserParams.Msg, &rawMsg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgTimeout: %v", err))
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		msgTimeoutParams := ibc_model.MsgTimeoutParams{
			RawMsgTimeout: rawMsg,

			Application: "transfer",
			MessageType: "/ibc.applications.transfer.v1.MsgTransfer",
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, msgTimeoutParams.Signer)

		return []command.Command{command_usecase.NewCreateMsgIBCTimeout(
			parserParams.MsgCommonParams,

			msgTimeoutParams,
		)}, possibleSignerAddresses
	}

	if !IsPacketMsgTransfer(rawMsg.Packet, &parserParams.TxsResult.Log[parserParams.MsgIndex]) {
		// unsupported application
		return []command.Command{}, []string{}
	}

	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])

	// Transfer application, MsgTransfer
	timeoutEvent := log.GetEventByType("timeout")

	var rawFungibleTokenPacketData ibc_model.FungibleTokenPacketData
	rawPacketData := base64_internal.MustDecodeString(rawMsg.Packet.Data)
	if err := json.Unmarshal(rawPacketData, &rawFungibleTokenPacketData); err != nil {
		rawFungibleTokenPacketData = ibc_model.FungibleTokenPacketData{}
	}

	if timeoutEvent != nil {
		if timeoutEvent.HasAttribute("refund_receiver") {
			rawFungibleTokenPacketData.Sender = timeoutEvent.MustGetAttributeByKey("refund_receiver")
		}
		if timeoutEvent.HasAttribute("refund_denom") {
			rawFungibleTokenPacketData.Denom = timeoutEvent.MustGetAttributeByKey("refund_denom")
		}
		if timeoutEvent.HasAttribute("refund_amount") {
			err := json.Unmarshal([]byte(timeoutEvent.MustGetAttributeByKey("refund_amount")), &rawFungibleTokenPacketData.Amount)
			if err != nil {
				parserParams.Logger.Errorf("error unmarshalling refund amount", err)
			}
		}
	}

	timeoutPacketEvent := log.GetEventByType("timeout_packet")
	if timeoutPacketEvent == nil {
		parserParams.Logger.Warnf("missing `timeout_packet` event in TxsResult log: ", parserParams.MsgCommonParams.TxHash)
		return []command.Command{}, []string{}
	}

	if timeoutPacketEvent.HasAttribute("packet_sequence") {
		rawMsg.Packet.Sequence = timeoutPacketEvent.MustGetAttributeByKey("packet_sequence")
	}
	if timeoutPacketEvent.HasAttribute("packet_src_port") {
		rawMsg.Packet.SourcePort = timeoutPacketEvent.MustGetAttributeByKey("packet_src_port")
	}
	if timeoutPacketEvent.HasAttribute("packet_src_channel") {
		rawMsg.Packet.SourceChannel = timeoutPacketEvent.MustGetAttributeByKey("packet_src_channel")
	}
	if timeoutPacketEvent.HasAttribute("packet_dst_port") {
		rawMsg.Packet.DestinationPort = timeoutPacketEvent.MustGetAttributeByKey("packet_dst_port")
	}
	if timeoutPacketEvent.HasAttribute("packet_dst_channel") {
		rawMsg.Packet.DestinationChannel = timeoutPacketEvent.MustGetAttributeByKey("packet_dst_channel")
	}

	msgTimeoutParams := ibc_model.MsgTimeoutParams{
		RawMsgTimeout: rawMsg,

		Application: "transfer",
		MessageType: "/ibc.applications.transfer.v1.MsgTransfer",
		MaybeMsgTransfer: &ibc_model.MsgTimeoutMsgTransfer{
			RefundReceiver: rawFungibleTokenPacketData.Sender,
			RefundDenom:    rawFungibleTokenPacketData.Denom,
			RefundAmount:   rawFungibleTokenPacketData.Amount,
		},

		PacketTimeoutHeight: MustParseHeight(
			timeoutPacketEvent.MustGetAttributeByKey("packet_timeout_height"),
		),
		PacketTimeoutTimestamp: typeconv.MustAtou64(
			timeoutPacketEvent.MustGetAttributeByKey("packet_timeout_timestamp"),
		),

		PacketSequence:  typeconv.MustAtou64(timeoutPacketEvent.MustGetAttributeByKey("packet_sequence")),
		ChannelOrdering: timeoutPacketEvent.MustGetAttributeByKey("packet_channel_ordering"),
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	possibleSignerAddresses = append(possibleSignerAddresses, msgTimeoutParams.Signer)

	if timeoutEvent == nil {
		return []command.Command{command_usecase.NewCreateMsgAlreadyRelayedIBCTimeout(
			parserParams.MsgCommonParams,

			msgTimeoutParams,
		)}, possibleSignerAddresses
	} else {
		return []command.Command{command_usecase.NewCreateMsgIBCTimeout(
			parserParams.MsgCommonParams,

			msgTimeoutParams,
		)}, possibleSignerAddresses
	}
}

func ParseMsgTimeoutOnClose(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var rawMsg ibc_model.RawMsgTimeoutOnClose
	if err := mapstructure_utils.DefaultMapstructureDecoder.Decode(parserParams.Msg, &rawMsg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgTimeoutOnClose: %v", err))
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		msgTimeoutOnCloseParams := ibc_model.MsgTimeoutOnCloseParams{
			RawMsgTimeoutOnClose: rawMsg,

			Application: "transfer",
			MessageType: "/ibc.applications.transfer.v1.MsgTransfer",
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, msgTimeoutOnCloseParams.Signer)

		return []command.Command{command_usecase.NewCreateMsgIBCTimeoutOnClose(
			parserParams.MsgCommonParams,

			msgTimeoutOnCloseParams,
		)}, possibleSignerAddresses
	}

	if !IsPacketMsgTransfer(rawMsg.Packet, &parserParams.TxsResult.Log[parserParams.MsgIndex]) {
		// unsupported application
		return []command.Command{}, []string{}
	}

	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])

	// Transfer application, MsgTransfer
	var rawFungibleTokenPacketData ibc_model.FungibleTokenPacketData
	rawPacketData := base64_internal.MustDecodeString(rawMsg.Packet.Data)
	json.MustUnmarshal(rawPacketData, &rawFungibleTokenPacketData)

	timeoutPacketEvent := log.GetEventByType("timeout_packet")
	if timeoutPacketEvent == nil {
		panic("missing `timeout_packet` event in TxsResult log")
	}

	msgTimeoutOnCloseParams := ibc_model.MsgTimeoutOnCloseParams{
		RawMsgTimeoutOnClose: rawMsg,

		Application: "transfer",
		MessageType: "/ibc.applications.transfer.v1.MsgTransfer",
		MaybeMsgTransfer: &ibc_model.MsgTimeoutMsgTransfer{
			RefundReceiver: rawFungibleTokenPacketData.Sender,
			RefundDenom:    rawFungibleTokenPacketData.Denom,
			RefundAmount:   rawFungibleTokenPacketData.Amount,
		},

		PacketTimeoutHeight: MustParseHeight(
			timeoutPacketEvent.MustGetAttributeByKey("packet_timeout_height"),
		),
		PacketTimeoutTimestamp: typeconv.MustAtou64(
			timeoutPacketEvent.MustGetAttributeByKey("packet_timeout_timestamp"),
		),

		PacketSequence:  typeconv.MustAtou64(timeoutPacketEvent.MustGetAttributeByKey("packet_sequence")),
		ChannelOrdering: timeoutPacketEvent.MustGetAttributeByKey("packet_channel_ordering"),
	}

	timeoutEvent := log.GetEventByType("timeout")

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	possibleSignerAddresses = append(possibleSignerAddresses, msgTimeoutOnCloseParams.Signer)

	if timeoutEvent == nil {
		return []command.Command{command_usecase.NewCreateMsgAlreadyRelayedIBCTimeoutOnClose(
			parserParams.MsgCommonParams,

			msgTimeoutOnCloseParams,
		)}, possibleSignerAddresses
	} else {
		return []command.Command{command_usecase.NewCreateMsgIBCTimeoutOnClose(
			parserParams.MsgCommonParams,

			msgTimeoutOnCloseParams,
		)}, possibleSignerAddresses
	}
}

func ParseMsgChannelCloseInit(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var rawMsg ibc_model.RawMsgChannelCloseInit
	if err := mapstructure_utils.DefaultMapstructureDecoder.Decode(parserParams.Msg, &rawMsg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgChannelCloseInit: %v", err))
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, rawMsg.Signer)

		return []command.Command{command_usecase.NewCreateMsgIBCChannelCloseInit(
			parserParams.MsgCommonParams,

			ibc_model.MsgChannelCloseInitParams{
				RawMsgChannelCloseInit: rawMsg,
			},
		)}, possibleSignerAddresses
	}

	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])
	event := log.GetEventByType("channel_close_init")
	if event == nil {
		panic("missing `channel_close_init` event in TxsResult log")
	}

	if parserParams.IsEthereumTxInnerMsg {
		msgChannelCloseInitParams := ibc_model.MsgChannelCloseInitParams{
			RawMsgChannelCloseInit: ibc_model.RawMsgChannelCloseInit{
				PortID:    event.MustGetAttributeByKey("port_id"),
				ChannelID: event.MustGetAttributeByKey("channel_id"),
			},

			CounterpartyPortID:    event.MustGetAttributeByKey("counterparty_port_id"),
			CounterpartyChannelID: event.MustGetAttributeByKey("counterparty_channel_id"),
			ConnectionID:          event.MustGetAttributeByKey("connection_id"),
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, msgChannelCloseInitParams.Signer)

		return []command.Command{command_usecase.NewCreateMsgIBCChannelCloseInit(
			parserParams.MsgCommonParams,

			msgChannelCloseInitParams,
		)}, possibleSignerAddresses
	}

	msgChannelCloseInitParams := ibc_model.MsgChannelCloseInitParams{
		RawMsgChannelCloseInit: rawMsg,

		CounterpartyPortID:    event.MustGetAttributeByKey("counterparty_port_id"),
		CounterpartyChannelID: event.MustGetAttributeByKey("counterparty_channel_id"),
		ConnectionID:          event.MustGetAttributeByKey("connection_id"),
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	possibleSignerAddresses = append(possibleSignerAddresses, msgChannelCloseInitParams.Signer)

	return []command.Command{command_usecase.NewCreateMsgIBCChannelCloseInit(
		parserParams.MsgCommonParams,

		msgChannelCloseInitParams,
	)}, possibleSignerAddresses
}

func ParseMsgChannelCloseConfirm(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var rawMsg ibc_model.RawMsgChannelCloseConfirm
	if err := mapstructure_utils.DefaultMapstructureDecoder.Decode(parserParams.Msg, &rawMsg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgChannelCloseConfirm: %v", err))
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, rawMsg.Signer)

		return []command.Command{command_usecase.NewCreateMsgIBCChannelCloseConfirm(
			parserParams.MsgCommonParams,

			ibc_model.MsgChannelCloseConfirmParams{
				RawMsgChannelCloseConfirm: rawMsg,
			},
		)}, possibleSignerAddresses
	}

	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])
	event := log.GetEventByType("channel_close_confirm")
	if event == nil {
		panic("missing `channel_close_confirm` event in TxsResult log")
	}

	if parserParams.IsEthereumTxInnerMsg {
		msgChannelCloseConfirmParams := ibc_model.MsgChannelCloseConfirmParams{
			RawMsgChannelCloseConfirm: ibc_model.RawMsgChannelCloseConfirm{
				PortID:    event.MustGetAttributeByKey("port_id"),
				ChannelID: event.MustGetAttributeByKey("channel_id"),
			},

			CounterpartyPortID:    event.MustGetAttributeByKey("counterparty_port_id"),
			CounterpartyChannelID: event.MustGetAttributeByKey("counterparty_channel_id"),
			ConnectionID:          event.MustGetAttributeByKey("connection_id"),
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		possibleSignerAddresses = append(possibleSignerAddresses, msgChannelCloseConfirmParams.Signer)

		return []command.Command{command_usecase.NewCreateMsgIBCChannelCloseConfirm(
			parserParams.MsgCommonParams,

			msgChannelCloseConfirmParams,
		)}, possibleSignerAddresses
	}

	msgChannelCloseConfirmParams := ibc_model.MsgChannelCloseConfirmParams{
		RawMsgChannelCloseConfirm: rawMsg,

		CounterpartyPortID:    event.MustGetAttributeByKey("counterparty_port_id"),
		CounterpartyChannelID: event.MustGetAttributeByKey("counterparty_channel_id"),
		ConnectionID:          event.MustGetAttributeByKey("connection_id"),
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	possibleSignerAddresses = append(possibleSignerAddresses, msgChannelCloseConfirmParams.Signer)

	return []command.Command{command_usecase.NewCreateMsgIBCChannelCloseConfirm(
		parserParams.MsgCommonParams,

		msgChannelCloseConfirmParams,
	)}, possibleSignerAddresses
}
