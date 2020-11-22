package parser

import (
	"github.com/crypto-com/chainindex/entity/command"
	"github.com/crypto-com/chainindex/usecase/coin"
	command_usecase "github.com/crypto-com/chainindex/usecase/command"
	"github.com/crypto-com/chainindex/usecase/model"
)

func ParseTxAccountTransferCommands(
	blockHeight int64,
	txsResults []model.BlockResultsTxsResult,
) ([]command.Command, error) {
	commands := make([]command.Command, 0)
	for _, txsResult := range txsResults {
		var lastSender string
		for i, event := range txsResult.Events {
			if event.Type == "message" {
				messageEvent := NewParsedTxsResultLogEvent(&txsResult.Events[i])
				if messageEvent.HasAttribute("sender") {
					lastSender = messageEvent.MustGetAttributeByKey("sender")
				}
			} else if event.Type == "transfer" {
				transferEvent := NewParsedTxsResultLogEvent(&txsResult.Events[i])

				amount := transferEvent.MustGetAttributeByKey("amount")
				if amount == "" {
					continue
				}

				var sender string
				if transferEvent.HasAttribute("sender") {
					sender = transferEvent.MustGetAttributeByKey("sender")
				} else {
					sender = lastSender
				}
				commands = append(commands, command_usecase.NewCreateAccountTransfer(
					blockHeight, model.AccountTransferParams{
						Recipient: transferEvent.MustGetAttributeByKey("recipient"),
						Sender:    sender,
						Amount:    coin.MustNewCoinFromString(TrimAmountDenom(amount)),
					}))
			}
		}
	}

	return commands, nil
}
