package parser

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/entity/command"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
)

func ParseBeginBlockEventsCommands(
	blockHeight int64,
	beginBlockEvents []model.BlockResultsEvent,
	bondingDenom string,
) ([]command.Command, error) {
	commands := make([]command.Command, 0)

	var proposerReward struct {
		address string
		amount  string
	}
	for i, event := range beginBlockEvents {
		if event.Type == "proposer_reward" {
			proposerRewardEvent := utils.NewParsedTxsResultLogEvent(&beginBlockEvents[i])
			proposerReward.address = proposerRewardEvent.MustGetAttributeByKey("validator")
			proposerReward.amount = proposerRewardEvent.MustGetAttributeByKey("amount")
			break
		}
	}
	for i, event := range beginBlockEvents {
		if event.Type == "transfer" {
			transferEvent := utils.NewParsedTxsResultLogEvent(&beginBlockEvents[i])

			// TODO: Missing Multi-send support (https://github.com/crypto-com/chain-indexing/issues/682)
			if !transferEvent.HasAttribute("sender") {
				continue
			}

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
			mintEvent := utils.NewParsedTxsResultLogEvent(&beginBlockEvents[i])

			amount := mintEvent.MustGetAttributeByKey("amount")
			if amount == "" {
				continue
			}
			annualProvisions := mintEvent.MustGetAttributeByKey("annual_provisions")
			if annualProvisions == "" {
				continue
			}
			commands = append(commands, command_usecase.NewCreateMint(
				blockHeight,
				model.MintParams{
					BondedRatio:      mintEvent.MustGetAttributeByKey("bonded_ratio"),
					Inflation:        mintEvent.MustGetAttributeByKey("inflation"),
					AnnualProvisions: coin.MustParseDecCoin(fmt.Sprintf("%s%s", annualProvisions, bondingDenom)),
					Amount:           coin.MustParseCoinsNormalized(fmt.Sprintf("%s%s", amount, bondingDenom)),
				},
			))
		} else if event.Type == "proposer_reward" {
			proposerRewardEvent := utils.NewParsedTxsResultLogEvent(&beginBlockEvents[i])
			amount := proposerRewardEvent.MustGetAttributeByKey("amount")
			if amount == "" {
				continue
			}

			validator := proposerRewardEvent.MustGetAttributeByKey("validator")
			commands = append(commands, command_usecase.NewCreateBlockProposerReward(
				blockHeight, validator, coin.MustParseDecCoins(amount),
			))
		} else if event.Type == "rewards" {
			proposerRewardEvent := utils.NewParsedTxsResultLogEvent(&beginBlockEvents[i])

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
			proposerRewardEvent := utils.NewParsedTxsResultLogEvent(&beginBlockEvents[i])
			amount := proposerRewardEvent.MustGetAttributeByKey("amount")
			if amount == "" {
				continue
			}

			validator := proposerRewardEvent.MustGetAttributeByKey("validator")
			commands = append(commands, command_usecase.NewCreateBlockCommission(
				blockHeight, validator, coin.MustParseDecCoins(amount),
			))
		} else if event.Type == "slash" {
			slashEvent := utils.NewParsedTxsResultLogEvent(&beginBlockEvents[i])

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

func ParseBeginBlockRawEventsCommands(
	blockHeight int64,
	blockHash string,
	blockTime utctime.UTCTime,
	beginBlockEvents []model.BlockResultsEvent,
) ([]command.Command, error) {
	commands := make([]command.Command, 0)

	for _, event := range beginBlockEvents {
		parseRawBlockEventCmd := command_usecase.NewCreateRawBlockEvent(blockHeight, model.CreateRawBlockEventParams{
			BlockHash:  blockHash,
			BlockTime:  blockTime,
			FromResult: "BeginBlockEvent",
			RawData: model.RawDataParams{
				Type:    event.Type,
				Content: event,
			},
		})

		commands = append(commands, parseRawBlockEventCmd)
	}

	return commands, nil
}
