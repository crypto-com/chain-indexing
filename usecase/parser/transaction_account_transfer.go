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
		for i, event := range txsResult.Events {
			if event.Type == "transfer" {
				transferEvent := NewParsedTxsResultLogEvent(&txsResult.Events[i])
				commands = append(commands, command_usecase.NewCreateAccountTransfer(
					blockHeight, model.AccountTransferParams{
						Recipient: transferEvent.MustGetAttributeByKey("recipient"),
						Sender:    transferEvent.MustGetAttributeByKey("sender"),
						Amount: coin.MustNewCoinFromString(TrimAmountDenom(
							transferEvent.MustGetAttributeByKey("amount"),
						)),
					}))
			}
		}
	}

	return commands, nil
}
