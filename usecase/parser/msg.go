package parser

import (
	"fmt"

	"github.com/crypto-com/chainindex/entity/command"
	"github.com/crypto-com/chainindex/usecase/coin"
	command_usecase "github.com/crypto-com/chainindex/usecase/command"
	"github.com/crypto-com/chainindex/usecase/event"
	"github.com/crypto-com/chainindex/usecase/model"
)

func ParseMsgToCommands(
	txDecoder *TxDecoder,
	block *model.Block,
	blockResults *model.BlockResults,
) []command.Command {
	commands := make([]command.Command, 0)

	blockHeight := block.Height
	for _, txHex := range block.Txs {
		txHash := TxHash(txHex)
		tx, err := txDecoder.Decode(txHex)
		if err != nil {
			panic(fmt.Sprintf("error decoding transaction: %v", err))
		}
		for msgIndex, msg := range tx.Body.Messages {
			if msg["@type"] == "/cosmos.bank.v1beta1.MsgSend" {
				commands = append(commands, command_usecase.NewCreateMsgSend(
					blockHeight,

					txHash,
					msgIndex,
					event.MsgSendCreatedParams{
						FromAddress: msg["from_address"].(string),
						ToAddress:   msg["to_address"].(string),
						Amount:      sumAmountInterfaces(msg["amount"].([]interface{})),
					},
				))
			}
		}
	}

	//for i, txsResult := range blockResults.TxsResults {
	//	for _, log := range txsResult.Log {
	//		txsResultLog := NewParsedTxsResultLog(&log)
	//		if !txsResultLog.HasEvent("message") {
	//			continue
	//		}
	//
	//		msgType := getMsgType(txsResultLog)
	//		if msgType.Module == "bank" && msgType.Action == "send" {
	//			transferEvent := txsResultLog.GetEvent("transfer")
	//			if transferEvent == nil {
	//				panic("expected bank.send message to have transfer event, but found none")
	//			}
	//
	//			amount := transferEvent.MustGetAttribute("amount")
	//			commands = append(commands, command_usecase.NewCreateMsgSend(
	//				blockHeight,
	//
	//				txHash,
	//				log.MsgIndex,
	//				event.MsgSendCreatedParams{
	//					FromAddress: transferEvent.MustGetAttribute("sender"),
	//					ToAddress:   transferEvent.MustGetAttribute("recipient"),
	//					Amount:      coin.MustNewCoinFromString(TrimAmountDenom(amount)),
	//				},
	//			))
	//		}
	//	}
	//}

	return commands
}

func sumAmountInterfaces(amounts []interface{}) coin.Coin {
	for _, rawAmount := range amounts {
		amount, _ := rawAmount.(map[string]interface{})
		return coin.MustNewCoinFromString(amount["amount"].(string))
	}

	return coin.Zero()
}

// Returns message type from block results TxsResult rawLog. Returns nil if no message is found.
//func getMsgType(txsResultLog *ParsedTxsResultLog) *msgType {
//	messageEvent := txsResultLog.GetEvent("message")
//	if messageEvent == nil {
//		return nil
//	}
//	return &msgType{
//		Module: messageEvent.MustGetAttribute("module"),
//		Action: messageEvent.MustGetAttribute("action"),
//	}
//}
//
//type msgType struct {
//	Module string
//	Action string
//}
