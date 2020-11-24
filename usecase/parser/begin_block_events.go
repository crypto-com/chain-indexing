package parser

import (
	"github.com/crypto-com/chainindex/entity/command"
	"github.com/crypto-com/chainindex/usecase/coin"
	command_usecase "github.com/crypto-com/chainindex/usecase/command"
	"github.com/crypto-com/chainindex/usecase/model"
)

func ParseBeginBlockEventsCommands(blockHeight int64, beginBlockEvents []model.BlockResultsEvent) ([]command.Command, error) {
	commands := make([]command.Command, 0)

	for i, event := range beginBlockEvents {
		if event.Type == "transfer" {
			transferEvent := NewParsedTxsResultLogEvent(&beginBlockEvents[i])

			amount := transferEvent.MustGetAttributeByKey("amount")
			if amount == "" {
				continue
			}
			commands = append(commands, command_usecase.NewCreateAccountTransfer(
				blockHeight, model.AccountTransferParams{
					Recipient: transferEvent.MustGetAttributeByKey("recipient"),
					Sender:    transferEvent.MustGetAttributeByKey("sender"),
					Amount:    coin.MustNewCoinFromString(TrimAmountDenom(amount)),
				}))
		}
	}

	return commands, nil
}
