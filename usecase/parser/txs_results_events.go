package parser

import (
	"github.com/crypto-com/chain-indexing/entity/command"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
)

func ParseTxsResultsEventsCommands(
	blockHeight int64,
	txsResults []model.BlockResultsTxsResult,
	bondingDenom string,
) ([]command.Command, error) {
	commands := make([]command.Command, 0)
	for _, txsResult := range txsResults {
		for i, event := range txsResult.Events {
			if event.Type == "coin_spent" {
				coinSpentEvent := utils.NewParsedTxsResultLogEvent(&txsResult.Events[i])

				amount := coinSpentEvent.MustGetAttributeByKey("amount")
				if amount == "" {
					continue
				}
				commands = append(commands, command_usecase.NewCreateCoinSpent(
					blockHeight,
					model.CoinSpentParams{
						Address: coinSpentEvent.MustGetAttributeByKey("spender"),
						Amount:  coin.MustParseCoinsNormalized(amount),
					}))
			} else if event.Type == "coin_received" {
				coinReceivedEvent := utils.NewParsedTxsResultLogEvent(&txsResult.Events[i])

				amount := coinReceivedEvent.MustGetAttributeByKey("amount")
				if amount == "" {
					continue
				}
				commands = append(commands, command_usecase.NewCreateCoinReceived(
					blockHeight,
					model.CoinReceivedParams{
						Address: coinReceivedEvent.MustGetAttributeByKey("receiver"),
						Amount:  coin.MustParseCoinsNormalized(amount),
					}))
			} else if event.Type == "coinbase" {
				coinMintEvent := utils.NewParsedTxsResultLogEvent(&txsResult.Events[i])

				amount := coinMintEvent.MustGetAttributeByKey("amount")
				if amount == "" {
					continue
				}
				commands = append(commands, command_usecase.NewCreateCoinMint(
					blockHeight,
					model.CoinMintParams{
						Address: coinMintEvent.MustGetAttributeByKey("minter"),
						Amount:  coin.MustParseCoinsNormalized(amount),
					}))
			} else if event.Type == "burn" {
				coinBurnEvent := utils.NewParsedTxsResultLogEvent(&txsResult.Events[i])

				amount := coinBurnEvent.MustGetAttributeByKey("amount")
				if amount == "" {
					continue
				}
				commands = append(commands, command_usecase.NewCreateCoinBurn(
					blockHeight,
					model.CoinBurnParams{
						Address: coinBurnEvent.MustGetAttributeByKey("burner"),
						Amount:  coin.MustParseCoinsNormalized(amount),
					}))
			}
		}
	}

	return commands, nil
}
