package parser

import (
	"encoding/hex"
	"fmt"
	"strings"

	commandentity "github.com/crypto-com/chain-indexing/entity/command"
	"github.com/crypto-com/chain-indexing/external/json"
	"github.com/crypto-com/chain-indexing/internal/typeconv"
	"github.com/crypto-com/chain-indexing/projection/block_raw_event/types"
	"github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
	jsoniter "github.com/json-iterator/go"
)

func ParseBlockResultsTxsResults(
	block *model.Block,
	blockResults *model.BlockResults,
) ([]commandentity.Command, error) {
	cmds := make([]commandentity.Command, 0)

	for i := range blockResults.TxsResults {
		txHex := block.Txs[i]

		parsedCmds := parseCronosSendToIBC(block.Height, txHex, &blockResults.TxsResults[i])
		cmds = append(cmds, parsedCmds...)

		for j := range blockResults.TxsResults[i].Events {
			parseBlockRawEventCmd := command.NewCreateBlockRawEvent(block.Height, model.CreateBlockRawEventParams{
				BlockHash:  block.Hash,
				BlockTime:  block.Time,
				FromResult: types.TXS_RESULTS,
				Data: model.DataParams{
					Type:    blockResults.TxsResults[i].Events[j].Type,
					Content: blockResults.TxsResults[i].Events[j],
				},
			})

			cmds = append(cmds, parseBlockRawEventCmd)
		}
	}

	return cmds, nil
}

func parseCronosSendToIBC(
	blockHeight int64,
	txHex string,
	txResults *model.BlockResultsTxsResult,
) []commandentity.Command {
	isEthereumTx := false
	var ethereumTxHash string

	var maybeIBCSendPacketEvent *utils.ParsedTxsResultLogEvent
	for i, event := range txResults.Events {
		if event.Type == "ethereum_tx" {
			isEthereumTx = true
			ethereumTxEvent := utils.NewParsedTxsResultLogEvent(&txResults.Events[i])
			ethereumTxHash = ethereumTxEvent.MustGetAttributeByKey("ethereumTxHash")
		} else if event.Type == "send_packet" {
			maybeIBCSendPacketEvent = utils.NewParsedTxsResultLogEvent(&txResults.Events[i])
		}
	}

	if !isEthereumTx {
		return nil
	}

	if maybeIBCSendPacketEvent == nil {
		return nil
	}

	fmt.Println("===> in parseCronosSendToIBC", maybeIBCSendPacketEvent)
	params := model.RawCronosSendToIBCParams{
		PacketChannelOrdering:  maybeIBCSendPacketEvent.MustGetAttributeByKey("packet_channel_ordering"),
		PacketConnection:       maybeIBCSendPacketEvent.MustGetAttributeByKey("packet_connection"),
		PacketDstChannel:       maybeIBCSendPacketEvent.MustGetAttributeByKey("packet_dst_channel"),
		PacketDstPort:          maybeIBCSendPacketEvent.MustGetAttributeByKey("packet_dst_port"),
		PacketSequence:         maybeIBCSendPacketEvent.MustGetAttributeByKey("packet_sequence"),
		PacketSrcChannel:       maybeIBCSendPacketEvent.MustGetAttributeByKey("packet_src_channel"),
		PacketSrcPort:          maybeIBCSendPacketEvent.MustGetAttributeByKey("packet_src_port"),
		PacketTimeoutHeight:    maybeIBCSendPacketEvent.MustGetAttributeByKey("packet_timeout_height"),
		PacketTimeoutTimestamp: maybeIBCSendPacketEvent.MustGetAttributeByKey("packet_timeout_timestamp"),
	}

	var fungibleTokenPacketData model.FungibleTokenPacketData
	if maybeIBCSendPacketEvent.HasAttribute("packet_data") {
		packetData := maybeIBCSendPacketEvent.MustGetAttributeByKey("packet_data")
		params.PacketData = packetData
		if unmarshalErr := jsoniter.Unmarshal([]byte(packetData), &fungibleTokenPacketData); unmarshalErr != nil {
			panic("unable to parse `send_packet` event, key `packet_data`")
		}
	}

	if maybeIBCSendPacketEvent.HasAttribute("packet_data_hex") {
		packetDataHex := maybeIBCSendPacketEvent.MustGetAttributeByKey("packet_data_hex")
		params.PacketDataHex = packetDataHex

		var packetData []byte
		packetData, err := hex.DecodeString(packetDataHex)
		if err == nil {
			if err = json.Unmarshal(packetData, &fungibleTokenPacketData); err != nil {
				panic("unable to parse `send_packet` event, key `packet_data_hex`")
			}
		}
	}

	timeoutHeight := mustParseCosmosSendToIBCTimeoutHeight(params.PacketTimeoutHeight)

	cmd := command.NewCreateCronosSendToIBC(blockHeight, model.CronosSendToIBCParams{
		TxHash:         TxHash(txHex),
		EthereumTxHash: ethereumTxHash,
		SourcePort:     params.PacketSrcPort,
		SourceChannel:  params.PacketSrcChannel,
		Token: model.CronosSendToIBCToken{
			Denom:  fungibleTokenPacketData.Denom,
			Amount: fungibleTokenPacketData.Amount,
		},
		Sender:   fungibleTokenPacketData.Sender,
		Receiver: fungibleTokenPacketData.Receiver,
		TimeoutHeight: model.CronosSendToIBCHeight{
			RevisionNumber: timeoutHeight.RevisionNumber,
			RevisionHeight: timeoutHeight.RevisionHeight,
		},
		TimeoutTimestamp:   params.PacketTimeoutTimestamp,
		PacketSequence:     typeconv.MustAtou64(params.PacketSequence),
		PacketDataHex:      params.PacketDataHex,
		DestinationPort:    params.PacketDstPort,
		DestinationChannel: params.PacketDstChannel,
		ChannelOrdering:    params.PacketChannelOrdering,
		ConnectionID:       params.PacketConnection,
	})

	return []commandentity.Command{cmd}
}

func mustParseCosmosSendToIBCTimeoutHeight(height string) model.CronosSendToIBCHeight {
	heightTokens := strings.Split(height, "-")
	if len(heightTokens) != 2 {
		panic(fmt.Errorf("invalid height: %s", height))
	}

	revisionNumber := typeconv.MustAtou64(heightTokens[0])
	revisionHeight := typeconv.MustAtou64(heightTokens[1])
	return model.CronosSendToIBCHeight{
		RevisionNumber: revisionNumber,
		RevisionHeight: revisionHeight,
	}
}
