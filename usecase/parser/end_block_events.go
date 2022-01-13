package parser

import (
	"github.com/crypto-com/chain-indexing/entity/command"
	"github.com/crypto-com/chain-indexing/internal/typeconv"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
)

func ParseEndBlockEventsCommands(blockHeight int64, endBlockEvents []model.BlockResultsEvent) ([]command.Command, error) {
	commands := make([]command.Command, 0)

	for i, event := range endBlockEvents {
		if event.Type == "transfer" {
			transferEvent := utils.NewParsedTxsResultLogEvent(&endBlockEvents[i])

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
		} else if event.Type == "complete_unbonding" {
			completeBondingEvent := utils.NewParsedTxsResultLogEvent(&endBlockEvents[i])
			amountValue := completeBondingEvent.MustGetAttributeByKey("amount")

			commands = append(commands, command_usecase.NewCreateCompleteBonding(
				blockHeight,
				model.CompleteBondingParams{
					Delegator: completeBondingEvent.MustGetAttributeByKey("delegator"),
					Validator: completeBondingEvent.MustGetAttributeByKey("validator"),
					Amount:    coin.MustParseCoinsNormalized(amountValue),
				},
			))
		} else if event.Type == "active_proposal" {
			activeProposalEvent := utils.NewParsedTxsResultLogEvent(&endBlockEvents[i])

			commands = append(commands, command_usecase.NewEndProposal(
				blockHeight,
				activeProposalEvent.MustGetAttributeByKey("proposal_id"),
				activeProposalEvent.MustGetAttributeByKey("proposal_result"),
			))
		} else if event.Type == "inactive_proposal" {
			activeProposalEvent := utils.NewParsedTxsResultLogEvent(&endBlockEvents[i])

			commands = append(commands, command_usecase.NewInactiveProposal(
				blockHeight,
				activeProposalEvent.MustGetAttributeByKey("proposal_id"),
				activeProposalEvent.MustGetAttributeByKey("proposal_result"),
			))
		} else if event.Type == "ethereum_send_to_cosmos_handled" {
			ethereumSendToCosmosHandledEvent := utils.NewParsedTxsResultLogEvent(&endBlockEvents[i])

			commands = append(commands, command_usecase.NewGravityHandleEthereumSendToCosmos(
				blockHeight,
				model.GravityEthereumSendToCosmosHandledEventParams{
					Module:                    ethereumSendToCosmosHandledEvent.MustGetAttributeByKey("module"),
					Sender:                    ethereumSendToCosmosHandledEvent.MustGetAttributeByKey("sender"),
					Receiver:                  ethereumSendToCosmosHandledEvent.MustGetAttributeByKey("receiver"),
					Amount:                    coin.MustParseCoinsNormalized(ethereumSendToCosmosHandledEvent.MustGetAttributeByKey("amount")),
					BridgeChainId:             typeconv.MustAtou64(ethereumSendToCosmosHandledEvent.MustGetAttributeByKey("bridge_chain_id")),
					EthereumTokenContract:     ethereumSendToCosmosHandledEvent.MustGetAttributeByKey("ethereum_token_contract"),
					Nonce:                     typeconv.MustAtou64(ethereumSendToCosmosHandledEvent.MustGetAttributeByKey("nonce")),
					EthereumEventVoteRecordId: []byte(ethereumSendToCosmosHandledEvent.MustGetAttributeByKey("ethereum_event_vote_record_id")),
				},
			))
		} else if event.Type == "coin_spent" {
			coinSpentEvent := utils.NewParsedTxsResultLogEvent(&endBlockEvents[i])

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
			coinReceivedEvent := utils.NewParsedTxsResultLogEvent(&endBlockEvents[i])

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
			coinMintEvent := utils.NewParsedTxsResultLogEvent(&endBlockEvents[i])

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
			coinBurnEvent := utils.NewParsedTxsResultLogEvent(&endBlockEvents[i])

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

	return commands, nil
}
