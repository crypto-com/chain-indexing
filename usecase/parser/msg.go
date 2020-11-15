package parser

import (
	"github.com/crypto-com/chainindex/entity/command"
	"github.com/crypto-com/chainindex/usecase/coin"
	command2 "github.com/crypto-com/chainindex/usecase/command"
	"github.com/crypto-com/chainindex/usecase/event"
	"github.com/crypto-com/chainindex/usecase/model"
)

func ParseMsgToCommands(block *model.Block, blockResults *model.BlockResults) []command.Command {
	commands := make([]command.Command, 0)

	blockHeight := block.Height
	for i, txsResult := range blockResults.TxsResults {
		txHex := block.Txs[i]
		txHash := TxHash(txHex)

		for _, log := range txsResult.Log {
			txsResultLog := NewParsedTxsResultLog(&log)
			if !txsResultLog.HasEvent("message") {
				continue
			}

			msgType := getMsgType(txsResultLog)
			if msgType.Module == "bank" && msgType.Action == "send" {
				transferEvent := txsResultLog.GetEvent("transfer")
				if transferEvent == nil {
					panic("expected bank.send message to have transfer event, but found none")
				}

				amount := transferEvent.MustGetAttribute("amount")
				commands = append(commands, command2.NewCreateMsgSend(
					blockHeight,

					txHash,
					log.MsgIndex,
					event.MsgSendCreatedParams{
						FromAddress: transferEvent.MustGetAttribute("sender"),
						ToAddress:   transferEvent.MustGetAttribute("recipient"),
						Amount:      coin.MustNewCoinFromString(TrimAmountDenom(amount)),
					},
				))
			}
		}
	}

	return commands
}

// Returns message type from block results TxsResult rawLog. Returns nil if no message is found.
func getMsgType(txsResultLog *ParsedTxsResultLog) *msgType {
	messageEvent := txsResultLog.GetEvent("message")
	if messageEvent == nil {
		return nil
	}
	return &msgType{
		Module: messageEvent.MustGetAttribute("module"),
		Action: messageEvent.MustGetAttribute("action"),
	}
}

type msgType struct {
	Module string
	Action string
}
