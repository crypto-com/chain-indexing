package msg

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"

	eth_types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/crypto-com/chain-indexing/entity/command"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/crypto-com/chain-indexing/usecase/parser/ibc"
	"github.com/crypto-com/chain-indexing/usecase/parser/icaauth"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
	mapstructure_utils "github.com/crypto-com/chain-indexing/usecase/parser/utils/mapstructure"
)

func ParseMsgEthereumTx(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var rawMsg model.RawMsgEthereumTx
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
		panic(fmt.Errorf("error creating RawMsgEthereumTx decoder: %v", decoderErr))
	}
	if err := decoder.Decode(parserParams.Msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgEthereumTx: %v", err))
	}

	fromBytes, err := base64.StdEncoding.DecodeString(rawMsg.From)
	if err != nil {
		panic(fmt.Errorf("error decoding from address: %v", err))
	}
	fromAddress := hex.EncodeToString(fromBytes)
	rawMsg.From = fmt.Sprintf("0x%s", utils.AddressParse(fromAddress))

	if !parserParams.MsgCommonParams.TxSuccess {
		// FIXME: https://github.com/crypto-com/chain-indexing/issues/730
		msgEthereumTxParams := model.MsgEthereumTxParams{
			RawMsgEthereumTx: rawMsg,
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		// FIXME: https://github.com/crypto-com/chain-indexing/issues/729
		// possibleSignerAddresses = append(possibleSignerAddresses, msgEthereumTxParams.From)

		return []command.Command{command_usecase.NewCreateMsgEthereumTx(
			parserParams.MsgCommonParams,

			msgEthereumTxParams,
		)}, possibleSignerAddresses
	}

	// FIXME: https://github.com/crypto-com/chain-indexing/issues/730
	msgEthereumTxParams := model.MsgEthereumTxParams{
		RawMsgEthereumTx: rawMsg,
	}

	var commands []command.Command
	var possibleSignerAddresses []string

	events := utils.NewParsedTxsResultsEvents(parserParams.TxsResult.Events)
	log := events.ParsedEventToTxsResultLog()
	parserParams.TxsResult.Log = log

	rawTxData, err := hex.DecodeString(rawMsg.Raw[2:]) // Remove hex prefix "0x..."
	if err != nil {
		panic("error decoding raw tx data")
	}

	// decode the raw tx data
	var tx eth_types.Transaction
	reader := bytes.NewReader(rawTxData)
	err = tx.DecodeRLP(rlp.NewStream(reader, 0))
	if err != nil {
		err = tx.UnmarshalBinary(rawTxData)
		if err != nil {
			panic(fmt.Errorf("error unmarshalling tx: %v", err))
		}
	}

	// unmarshal the raw tx data into `decodeRaw` field
	var txBytes []byte
	txBytes, err = tx.MarshalJSON()
	if err != nil {
		panic(fmt.Errorf("error marshalling tx: %v", err))
	}

	var decodedRaw model.DecodedRaw
	err = json.Unmarshal(txBytes, &decodedRaw)
	if err != nil {
		panic(fmt.Errorf("error unmarshalling tx: %v", err))
	}
	msgEthereumTxParams.DecodedRaw = &decodedRaw

	if len(tx.Data()) > 0 {
		inputData := tx.Data()

		// parse IBC msgChannelOpenInit
		if events.HasEvent("channel_open_init") {
			channelOpenInitEvents := events.GetEventsByType("channel_open_init")
			if len(channelOpenInitEvents) > 0 {
				for _, event := range channelOpenInitEvents {
					counterpartyPortId := event.GetAttributeByKey("counterparty_port_id")
					if counterpartyPortId != nil {
						if *counterpartyPortId == "icahost" {
							parserParams.IsEthereumTxInnerMsg = true
							cmds, signers := icaauth.ParseMsgRegisterAccount(parserParams)
							commands = append(commands, cmds...)
							possibleSignerAddresses = append(possibleSignerAddresses, signers...)
							break
						}
					} else {
						parserParams.IsEthereumTxInnerMsg = true
						cmds, signers := ibc.ParseMsgChannelOpenInit(parserParams)
						commands = append(commands, cmds...)
						possibleSignerAddresses = append(possibleSignerAddresses, signers...)
						break
					}
				}
			}
		} else if events.HasEvent("channel_open_try") {
			// parse IBC msgChannelOpenTry
			if events.GetEventByType("channel_open_try") != nil {
				parserParams.IsEthereumTxInnerMsg = true
				cmds, signers := ibc.ParseMsgChannelOpenTry(parserParams)
				commands = append(commands, cmds...)
				possibleSignerAddresses = append(possibleSignerAddresses, signers...)
			}
		} else if events.HasEvent("channel_open_ack") {
			// parse IBC msgChannelOpenAck

			if events.GetEventByType("channel_open_ack") != nil {
				parserParams.IsEthereumTxInnerMsg = true
				cmds, signers := ibc.ParseMsgChannelOpenAck(parserParams)
				commands = append(commands, cmds...)
				possibleSignerAddresses = append(possibleSignerAddresses, signers...)
			}
		} else if events.HasEvent("channel_open_confirm") {
			// parse IBC msgChannelOpenConfirm

			if events.GetEventByType("channel_open_confirm") != nil {
				parserParams.IsEthereumTxInnerMsg = true
				cmds, signers := ibc.ParseMsgChannelOpenConfirm(parserParams)
				commands = append(commands, cmds...)
				possibleSignerAddresses = append(possibleSignerAddresses, signers...)
			}
		} else if events.HasEvent("submit_msgs_result") {
			// parse ICA msgSendTx
			if events.GetEventByType("submit_msgs_result") != nil {
				parserParams.IsEthereumTxInnerMsg = true
				cmds, signers := icaauth.ParseMsgSubmitTx(parserParams)
				commands = append(commands, cmds...)
				possibleSignerAddresses = append(possibleSignerAddresses, signers...)
			}
		} else if events.HasEvent("acknowledge_packet") {
			// parse MsgAcknowledgement
			if events.GetEventByType("acknowledge_packet") != nil {
				parserParams.IsEthereumTxInnerMsg = true
				cmds, signers := ibc.ParseMsgAcknowledgement(parserParams)
				commands = append(commands, cmds...)
				possibleSignerAddresses = append(possibleSignerAddresses, signers...)
			}
		} else if events.HasEvent("timeout_packet") {
			// parse MsgTimeout

			if events.GetEventByType("timeout_packet") != nil {
				parserParams.IsEthereumTxInnerMsg = true
				cmds, signers := ibc.ParseMsgTimeout(parserParams)
				commands = append(commands, cmds...)
				possibleSignerAddresses = append(possibleSignerAddresses, signers...)
			}
		} else if events.HasEvent("recv_packet") {
			// parse MsgRecvPacket
			recvPacketEvents := events.GetEventsByType("recv_packet")
			if len(recvPacketEvents) > 0 {
				parserParams.IsEthereumTxInnerMsg = true
				cmds, signers := ibc.ParseMsgRecvPacket(parserParams)
				commands = append(commands, cmds...)
				possibleSignerAddresses = append(possibleSignerAddresses, signers...)
			}
		} else if events.HasEvent("channel_close_init") {
			// parse msgChannelCloseInit
			if events.GetEventByType("channel_close_init") != nil {
				parserParams.IsEthereumTxInnerMsg = true
				cmds, signers := ibc.ParseMsgChannelCloseInit(parserParams)
				commands = append(commands, cmds...)
				possibleSignerAddresses = append(possibleSignerAddresses, signers...)
			}
		} else if events.HasEvent("channel_close_confirm") {
			// parse msgChannelCloseConfirm
			if events.GetEventByType("channel_close_confirm") != nil {
				parserParams.IsEthereumTxInnerMsg = true
				cmds, signers := ibc.ParseMsgChannelCloseConfirm(parserParams)
				commands = append(commands, cmds...)
				possibleSignerAddresses = append(possibleSignerAddresses, signers...)
			}
		} else if events.HasEvent("connection_open_init") {
			// parse msgConnectionOpenInit
			if events.GetEventByType("connection_open_init") != nil {
				parserParams.IsEthereumTxInnerMsg = true
				cmds, signers := ibc.ParseMsgConnectionOpenInit(parserParams)
				commands = append(commands, cmds...)
				possibleSignerAddresses = append(possibleSignerAddresses, signers...)
			}
		} else if events.HasEvent("connection_open_try") {
			// parse msgConnectionOpenTry
			if events.GetEventByType("connection_open_try") != nil {
				parserParams.IsEthereumTxInnerMsg = true
				cmds, signers := ibc.ParseMsgConnectionOpenTry(parserParams)
				commands = append(commands, cmds...)
				possibleSignerAddresses = append(possibleSignerAddresses, signers...)
			}
		} else if events.HasEvent("connection_open_ack") {
			// parse msgConnectionOpenAck
			if events.GetEventByType("connection_open_ack") != nil {
				parserParams.IsEthereumTxInnerMsg = true
				cmds, signers := ibc.ParseMsgConnectionOpenAck(parserParams)
				commands = append(commands, cmds...)
				possibleSignerAddresses = append(possibleSignerAddresses, signers...)
			}
		} else if events.HasEvent("connection_open_confirm") {
			// parse msgConnectionOpenConfirm
			if events.GetEventByType("connection_open_confirm") != nil {
				parserParams.IsEthereumTxInnerMsg = true
				cmds, signers := ibc.ParseMsgConnectionOpenConfirm(parserParams)
				commands = append(commands, cmds...)
				possibleSignerAddresses = append(possibleSignerAddresses, signers...)
			}
		} else if events.HasEvent("send_packet") {
			// parse msgTransfer
			sendEvents := events.GetEventsByType("send_packet")
			if len(sendEvents) > 0 {
				parserParams.IsEthereumTxInnerMsg = true
				cmds, signers := ibc.ParseMsgTransfer(parserParams)
				commands = append(commands, cmds...)
				possibleSignerAddresses = append(possibleSignerAddresses, signers...)
			}
		} else if events.HasEvent("create_client") {
			// parse msgCreateClient
			sendEvents := events.GetEventsByType("create_client")
			if len(sendEvents) > 0 {
				msg, err := parserParams.EthereumTxInnerMsgDecoder.DecodeCosmosMsgFromTxInput(inputData, "MsgCreateClient")
				if err != nil {
					panic(fmt.Errorf("error deserializing MsgCreateClient: %v", err))
				}

				parserParams.Msg = msg
				parserParams.IsEthereumTxInnerMsg = true
				cmds, signers := ibc.ParseMsgCreateClient(parserParams)
				commands = append(commands, cmds...)
				possibleSignerAddresses = append(possibleSignerAddresses, signers...)
			}
		} else if events.HasEvent("update_client") {
			// parse MsgUpdateClient
			sendEvents := events.GetEventsByType("update_client")
			if len(sendEvents) > 0 {
				msg, err := parserParams.EthereumTxInnerMsgDecoder.DecodeCosmosMsgFromTxInput(inputData, "MsgUpdateClient")
				if err != nil {
					panic(fmt.Errorf("error deserializing MsgUpdateClient: %v", err))
				}

				parserParams.Msg = msg
				parserParams.IsEthereumTxInnerMsg = true
				cmds, signers := ibc.ParseMsgUpdateClient(parserParams)
				commands = append(commands, cmds...)
				possibleSignerAddresses = append(possibleSignerAddresses, signers...)
			}
		}
	}

	// Getting possible signer address from Msg
	// FIXME: https://github.com/crypto-com/chain-indexing/issues/729
	// possibleSignerAddresses = append(possibleSignerAddresses, msgEthereumTxParams.From)

	return append(commands, command_usecase.NewCreateMsgEthereumTx(
		parserParams.MsgCommonParams,

		msgEthereumTxParams,
	)), possibleSignerAddresses
}

func ParseMsgEventsToLog(
	parser utils.CosmosParser,
) utils.CosmosParser {
	return func(
		parserParams utils.CosmosParserParams,
	) ([]command.Command, []string) {
		events := utils.NewParsedTxsResultsEvents(parserParams.TxsResult.Events)
		log := events.ParsedEventToTxsResultLog()
		parserParams.TxsResult.Log = log

		return parser(utils.CosmosParserParams{
			AddressPrefix:   parserParams.AddressPrefix,
			StakingDenom:    parserParams.StakingDenom,
			TxsResult:       parserParams.TxsResult,
			MsgCommonParams: parserParams.MsgCommonParams,
			Msg:             parserParams.Msg,
			MsgIndex:        parserParams.MsgIndex,
			ParserManager:   parserParams.ParserManager,
		})
	}
}
