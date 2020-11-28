package parser

import (
	"github.com/crypto-com/chain-indexing/entity/command"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

func ParseEndBlockEventsCommands(blockHeight int64, endBlockEvents []model.BlockResultsEvent) ([]command.Command, error) {
	commands := make([]command.Command, 0)

	for i, event := range endBlockEvents {
		if event.Type == "transfer" {
			transferEvent := NewParsedTxsResultLogEvent(&endBlockEvents[i])

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
		} else if event.Type == "complete_unbonding" {
			completeBondingEvent := NewParsedTxsResultLogEvent(&endBlockEvents[i])
			amountValue := completeBondingEvent.MustGetAttributeByKey("amount")

			commands = append(commands, command_usecase.NewCreateCompleteBonding(
				blockHeight,
				model.CompleteBondingParams{
					Delegator: completeBondingEvent.MustGetAttributeByKey("delegator"),
					Validator: completeBondingEvent.MustGetAttributeByKey("validator"),
					Amount:    coin.MustNewCoinFromString(TrimAmountDenom(amountValue)),
				},
			))
		} else if event.Type == "active_proposal" {
			activeProposalEvent := NewParsedTxsResultLogEvent(&endBlockEvents[i])

			commands = append(commands, command_usecase.NewEndProposal(
				blockHeight,
				activeProposalEvent.MustGetAttributeByKey("proposal_id"),
				activeProposalEvent.MustGetAttributeByKey("proposal_result"),
			))
		} else if event.Type == "inactive_proposal" {
			activeProposalEvent := NewParsedTxsResultLogEvent(&endBlockEvents[i])

			commands = append(commands, command_usecase.NewInactiveProposal(
				blockHeight,
				activeProposalEvent.MustGetAttributeByKey("proposal_id"),
				activeProposalEvent.MustGetAttributeByKey("proposal_result"),
			))
		}
	}

	return commands, nil
}
