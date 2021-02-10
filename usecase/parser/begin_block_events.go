package parser

import (
	"github.com/crypto-com/chain-indexing/entity/command"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

func ParseBeginBlockEventsCommands(blockHeight int64, beginBlockEvents []model.BlockResultsEvent) ([]command.Command, error) {
	commands := make([]command.Command, 0)

	var proposerReward struct {
		address string
		amount  string
	}
	for i, event := range beginBlockEvents {
		if event.Type == "proposer_reward" {
			proposerRewardEvent := NewParsedTxsResultLogEvent(&beginBlockEvents[i])
			proposerReward.address = proposerRewardEvent.MustGetAttributeByKey("validator")
			proposerReward.amount = proposerRewardEvent.MustGetAttributeByKey("amount")
			break
		}
	}
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
					Amount:    coin.MustParseCoinsNormalized(amount),
				}))
		} else if event.Type == "mint" {
			mintEvent := NewParsedTxsResultLogEvent(&beginBlockEvents[i])
			commands = append(commands, command_usecase.NewCreateMint(
				blockHeight,
				model.MintParams{
					BondedRatio:      mintEvent.MustGetAttributeByKey("bonded_ratio"),
					Inflation:        mintEvent.MustGetAttributeByKey("inflation"),
					AnnualProvisions: mintEvent.MustGetAttributeByKey("annual_provisions"),
					Amount:           mintEvent.MustGetAttributeByKey("amount"),
				},
			))
		} else if event.Type == "proposer_reward" {
			proposerRewardEvent := NewParsedTxsResultLogEvent(&beginBlockEvents[i])
			amount := proposerRewardEvent.MustGetAttributeByKey("amount")
			if amount == "" {
				continue
			}

			validator := proposerRewardEvent.MustGetAttributeByKey("validator")
			commands = append(commands, command_usecase.NewCreateBlockProposerReward(
				blockHeight, validator, coin.MustParseDecCoins(amount),
			))
		} else if event.Type == "rewards" {
			proposerRewardEvent := NewParsedTxsResultLogEvent(&beginBlockEvents[i])

			amount := proposerRewardEvent.MustGetAttributeByKey("amount")
			if amount == "" {
				continue
			}

			validator := proposerRewardEvent.MustGetAttributeByKey("validator")
			if validator == proposerReward.address && amount == proposerReward.amount {
				// prevent proposer reward and block reward double count
				continue
			}

			commands = append(commands, command_usecase.NewCreateBlockReward(
				blockHeight, validator, coin.MustParseDecCoins(amount),
			))
		} else if event.Type == "commission" {
			proposerRewardEvent := NewParsedTxsResultLogEvent(&beginBlockEvents[i])
			amount := proposerRewardEvent.MustGetAttributeByKey("amount")
			if amount == "" {
				continue
			}

			validator := proposerRewardEvent.MustGetAttributeByKey("validator")
			commands = append(commands, command_usecase.NewCreateBlockCommission(
				blockHeight, validator, coin.MustParseDecCoins(amount),
			))
		} else if event.Type == "slash" {
			slashEvent := NewParsedTxsResultLogEvent(&beginBlockEvents[i])

			if slashEvent.HasAttribute("address") {
				commands = append(commands, command_usecase.NewSlashValidator(
					blockHeight,
					model.SlashValidatorParams{
						ConsensusNodeAddress: slashEvent.MustGetAttributeByKey("address"),
						SlashedPower:         slashEvent.MustGetAttributeByKey("power"),
						Reason:               slashEvent.MustGetAttributeByKey("reason"),
					},
				))
				// Slash is always accompanied by jailed
				commands = append(commands, command_usecase.NewJailValidator(
					blockHeight,
					slashEvent.MustGetAttributeByKey("address"),
					slashEvent.MustGetAttributeByKey("reason"),
				))
			}
		}
	}

	return commands, nil
}
