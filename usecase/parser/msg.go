package parser

import (
	"fmt"
	"strconv"
	"time"

	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
	"github.com/mitchellh/mapstructure"

	"github.com/crypto-com/chain-indexing/projection/validator/constants"
	"github.com/crypto-com/chain-indexing/usecase/model/genesis"

	"github.com/crypto-com/chain-indexing/internal/tmcosmosutils"

	"github.com/crypto-com/chain-indexing/internal/primptr"

	"github.com/crypto-com/chain-indexing/internal/utctime"

	jsoniter "github.com/json-iterator/go"

	"github.com/crypto-com/chain-indexing/entity/command"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

func ParseBlockResultsTxsMsgToCommands(
	parserManager *utils.CosmosParserManager,
	txDecoder *utils.TxDecoder,
	block *model.Block,
	blockResults *model.BlockResults,
	accountAddressPrefix string,
	stakingDenom string,
) ([]command.Command, error) {
	commands := make([]command.Command, 0)

	blockHeight := block.Height
	for i, txHex := range block.Txs {
		txHash := TxHash(txHex)
		txSuccess := true
		txsResult := blockResults.TxsResults[i]

		if txsResult.Code != 0 {
			txSuccess = false
		}
		tx, err := txDecoder.Decode(txHex)
		if err != nil {
			panic(fmt.Sprintf("error decoding transaction: %v", err))
		}

		for msgIndex, msg := range tx.Body.Messages {
			msgCommonParams := event.MsgCommonParams{
				BlockHeight: blockHeight,
				TxHash:      txHash,
				TxSuccess:   txSuccess,
				MsgIndex:    msgIndex,
			}

			var msgCommands []command.Command

			msgType := msg["@type"]
			switch msgType {
			case
				// cosmos bank
				"/cosmos.bank.v1beta1.MsgSend",
				"/cosmos.bank.v1beta1.MsgMultiSend",

				// cosmos distribution
				"/cosmos.distribution.v1beta1.MsgSetWithdrawAddress",
				"/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward",
				"/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission",
				"/cosmos.distribution.v1beta1.MsgFundCommunityPool",

				// cosmos gov
				"/cosmos.gov.v1beta1.MsgSubmitProposal",
				"/cosmos.gov.v1beta1.MsgVote",
				"/cosmos.gov.v1beta1.MsgDeposit",

				// cosmos staking
				"/cosmos.staking.v1beta1.MsgDelegate",
				"/cosmos.staking.v1beta1.MsgUndelegate",
				"/cosmos.staking.v1beta1.MsgBeginRedelegate",
				"/cosmos.staking.v1beta1.MsgCreateValidator",
				"/cosmos.staking.v1beta1.MsgEditValidator",

				// cosmos slashing
				"/cosmos.slashing.v1beta1.MsgUnjail",

				// chainmain nft
				"/chainmain.nft.v1.MsgIssueDenom",
				"/chainmain.nft.v1.MsgMintNFT",
				"/chainmain.nft.v1.MsgTransferNFT",
				"/chainmain.nft.v1.MsgEditNFT",
				"/chainmain.nft.v1.MsgBurnNFT",

				// ibc core client
				"/ibc.core.client.v1.MsgCreateClient",
				"/ibc.core.client.v1.MsgUpdateClient",

				// ibc core connection
				"/ibc.core.connection.v1.MsgConnectionOpenInit",
				"/ibc.core.connection.v1.MsgConnectionOpenTry",
				"/ibc.core.connection.v1.MsgConnectionOpenAck",
				"/ibc.core.connection.v1.MsgConnectionOpenConfirm",

				// ibc core channel
				"/ibc.core.channel.v1.MsgChannelOpenInit",
				"/ibc.core.channel.v1.MsgChannelOpenTry",
				"/ibc.core.channel.v1.MsgChannelOpenAck",
				"/ibc.core.channel.v1.MsgChannelOpenConfirm",
				"/ibc.core.channel.v1.MsgRecvPacket",
				"/ibc.core.channel.v1.MsgAcknowledgement",
				"/ibc.core.channel.v1.MsgTimeout",
				"/ibc.core.channel.v1.MsgTimeoutOnClose",

				// ibc applications transfer
				"/ibc.applications.transfer.v1.MsgTransfer",

				// cosmos authz
				"/cosmos.authz.v1beta1.MsgGrant",
				"/cosmos.authz.v1beta1.MsgRevoke",
				"/cosmos.authz.v1beta1.MsgExec",

				// cosmos feegrant
				"/cosmos.feegrant.v1beta1.MsgGrantAllowance",
				"/cosmos.feegrant.v1beta1.MsgRevokeAllowance",

				// cosmos vesting
				"/cosmos.vesting.v1beta1.MsgCreateVestingAccount":
				parser := parserManager.GetParser(utils.CosmosParserKey(msgType.(string)), utils.ParserBlockHeight(blockHeight))
				msgCommands = parser(utils.CosmosParserParams{
					AddressPrefix:   accountAddressPrefix,
					StakingDenom:    stakingDenom,
					TxsResult:       txsResult,
					MsgCommonParams: msgCommonParams,
					Msg:             msg,
					MsgIndex:        msgIndex,
				})
			}

			commands = append(commands, msgCommands...)
		}
	}

	return commands, nil
}

func ParseMsgSend(
	parserParams utils.CosmosParserParams,
) []command.Command {
	return []command.Command{command_usecase.NewCreateMsgSend(
		parserParams.MsgCommonParams,

		event.MsgSendCreatedParams{
			FromAddress: parserParams.Msg["from_address"].(string),
			ToAddress:   parserParams.Msg["to_address"].(string),
			Amount:      tmcosmosutils.MustNewCoinsFromAmountInterface(parserParams.Msg["amount"].([]interface{})),
		},
	)}
}

func ParseMsgMultiSend(
	parserParams utils.CosmosParserParams,
) []command.Command {
	rawInputs, _ := parserParams.Msg["inputs"].([]interface{})
	inputs := make([]model.MsgMultiSendInput, 0, len(rawInputs))
	for _, rawInput := range rawInputs {
		input, _ := rawInput.(map[string]interface{})
		inputs = append(inputs, model.MsgMultiSendInput{
			Address: input["address"].(string),
			Amount:  tmcosmosutils.MustNewCoinsFromAmountInterface(input["coins"].([]interface{})),
		})
	}

	rawOutputs, _ := parserParams.Msg["outputs"].([]interface{})
	outputs := make([]model.MsgMultiSendOutput, 0, len(rawOutputs))
	for _, rawOutput := range rawOutputs {
		output, _ := rawOutput.(map[string]interface{})
		outputs = append(outputs, model.MsgMultiSendOutput{
			Address: output["address"].(string),
			Amount:  tmcosmosutils.MustNewCoinsFromAmountInterface(output["coins"].([]interface{})),
		})
	}

	return []command.Command{command_usecase.NewCreateMsgMultiSend(
		parserParams.MsgCommonParams,

		model.MsgMultiSendParams{
			Inputs:  inputs,
			Outputs: outputs,
		},
	)}
}

func ParseMsgSetWithdrawAddress(
	parserParams utils.CosmosParserParams,
) []command.Command {
	return []command.Command{command_usecase.NewCreateMsgSetWithdrawAddress(
		parserParams.MsgCommonParams,

		model.MsgSetWithdrawAddressParams{
			DelegatorAddress: parserParams.Msg["delegator_address"].(string),
			WithdrawAddress:  parserParams.Msg["withdraw_address"].(string),
		},
	)}
}

func ParseMsgWithdrawDelegatorReward(
	parserParams utils.CosmosParserParams,
) []command.Command {
	if !parserParams.MsgCommonParams.TxSuccess {
		delegatorAddress, _ := parserParams.Msg["delegator_address"].(string)
		return []command.Command{command_usecase.NewCreateMsgWithdrawDelegatorReward(
			parserParams.MsgCommonParams,

			model.MsgWithdrawDelegatorRewardParams{
				DelegatorAddress: delegatorAddress,
				ValidatorAddress: parserParams.Msg["validator_address"].(string),
				RecipientAddress: delegatorAddress,
				Amount:           coin.NewEmptyCoins(),
			},
		)}
	}
	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])
	var recipient string
	var amount coin.Coins
	// When there is no reward withdrew, `transfer` event would not exist
	if event := log.GetEventByType("transfer"); event == nil {
		recipient, _ = parserParams.Msg["delegator_address"].(string)
		amount = coin.NewEmptyCoins()
	} else {
		recipient = event.MustGetAttributeByKey("recipient")
		amountValue := event.MustGetAttributeByKey("amount")
		amount = coin.MustParseCoinsNormalized(amountValue)
	}

	return []command.Command{command_usecase.NewCreateMsgWithdrawDelegatorReward(
		parserParams.MsgCommonParams,

		model.MsgWithdrawDelegatorRewardParams{
			DelegatorAddress: parserParams.Msg["delegator_address"].(string),
			ValidatorAddress: parserParams.Msg["validator_address"].(string),
			RecipientAddress: recipient,
			Amount:           amount,
		},
	)}
}

func ParseMsgWithdrawValidatorCommission(
	parserParams utils.CosmosParserParams,
) []command.Command {
	if !parserParams.MsgCommonParams.TxSuccess {
		return []command.Command{command_usecase.NewCreateMsgWithdrawValidatorCommission(
			parserParams.MsgCommonParams,

			model.MsgWithdrawValidatorCommissionParams{
				ValidatorAddress: parserParams.Msg["validator_address"].(string),
				RecipientAddress: "",
				Amount:           coin.NewEmptyCoins(),
			},
		)}
	}
	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])
	var recipient string
	var amount coin.Coins
	// When there is no reward withdrew, `transfer` event would not exist
	if event := log.GetEventByType("transfer"); event == nil {
		recipient, _ = parserParams.Msg["delegator_address"].(string)
		amount = coin.NewEmptyCoins()
	} else {
		recipient = event.MustGetAttributeByKey("recipient")
		amountValue := event.MustGetAttributeByKey("amount")
		amount = coin.MustParseCoinsNormalized(amountValue)
	}

	return []command.Command{command_usecase.NewCreateMsgWithdrawValidatorCommission(
		parserParams.MsgCommonParams,

		model.MsgWithdrawValidatorCommissionParams{
			ValidatorAddress: parserParams.Msg["validator_address"].(string),
			RecipientAddress: recipient,
			Amount:           amount,
		},
	)}
}

func ParseMsgFundCommunityPool(
	parserParams utils.CosmosParserParams,
) []command.Command {
	return []command.Command{command_usecase.NewCreateMsgFundCommunityPool(
		parserParams.MsgCommonParams,

		model.MsgFundCommunityPoolParams{
			Depositor: parserParams.Msg["depositor"].(string),
			Amount:    tmcosmosutils.MustNewCoinsFromAmountInterface(parserParams.Msg["amount"].([]interface{})),
		},
	)}
}

func ParseMsgSubmitProposal(
	parserParams utils.CosmosParserParams,
) []command.Command {
	rawContent, err := jsoniter.Marshal(parserParams.Msg["content"])
	if err != nil {
		panic(fmt.Sprintf("error encoding proposal content: %v", err))
	}
	var proposalContent model.MsgSubmitProposalContent
	if err := jsoniter.Unmarshal(rawContent, &proposalContent); err != nil {
		panic(fmt.Sprintf("error decoding proposal content: %v", err))
	}

	var cmds []command.Command
	if proposalContent.Type == "/cosmos.params.v1beta1.ParameterChangeProposal" {
		cmds = parseMsgSubmitParamChangeProposal(parserParams.MsgCommonParams.TxSuccess, parserParams.TxsResult, parserParams.MsgIndex, parserParams.MsgCommonParams, parserParams.Msg, rawContent)
	} else if proposalContent.Type == "/cosmos.distribution.v1beta1.CommunityPoolSpendProposal" {
		cmds = parseMsgSubmitCommunityFundSpendProposal(parserParams.MsgCommonParams.TxSuccess, parserParams.TxsResult, parserParams.MsgIndex, parserParams.MsgCommonParams, parserParams.Msg, rawContent)
	} else if proposalContent.Type == "/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal" {
		cmds = parseMsgSubmitSoftwareUpgradeProposal(parserParams.MsgCommonParams.TxSuccess, parserParams.TxsResult, parserParams.MsgIndex, parserParams.MsgCommonParams, parserParams.Msg, rawContent)
	} else if proposalContent.Type == "/cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal" {
		cmds = parseMsgSubmitCancelSoftwareUpgradeProposal(parserParams.MsgCommonParams.TxSuccess, parserParams.TxsResult, parserParams.MsgIndex, parserParams.MsgCommonParams, parserParams.Msg, rawContent)
	} else if proposalContent.Type == "/cosmos.gov.v1beta1.TextProposal" {
		cmds = parseMsgSubmitTextProposal(parserParams.MsgCommonParams.TxSuccess, parserParams.TxsResult, parserParams.MsgIndex, parserParams.MsgCommonParams, parserParams.Msg, rawContent)
	} else {
		panic(fmt.Sprintf("unrecognzied govenance proposal type `%s`", proposalContent.Type))
	}

	if parserParams.MsgCommonParams.TxSuccess {
		log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])
		logEvent := log.GetEventByType("submit_proposal")
		if logEvent == nil {
			panic("missing `submit_proposal` event in TxsResult log")
		}

		if logEvent.HasAttribute("voting_period_start") {
			cmds = append(cmds, command_usecase.NewStartProposalVotingPeriod(
				parserParams.MsgCommonParams.BlockHeight, logEvent.MustGetAttributeByKey("voting_period_start"),
			))
		}
	}

	return cmds
}

func parseMsgSubmitParamChangeProposal(
	txSuccess bool,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
	rawContent []byte,
) []command.Command {
	var proposalContent model.MsgSubmitParamChangeProposalContent
	if err := jsoniter.Unmarshal(rawContent, &proposalContent); err != nil {
		panic("error decoding param change proposal content")
	}

	if !txSuccess {
		return []command.Command{command_usecase.NewCreateMsgSubmitParamChangeProposal(
			msgCommonParams,

			model.MsgSubmitParamChangeProposalParams{
				MaybeProposalId: nil,
				Content:         proposalContent,
				ProposerAddress: msg["proposer"].(string),
				InitialDeposit: tmcosmosutils.MustNewCoinsFromAmountInterface(
					msg["initial_deposit"].([]interface{}),
				),
			},
		)}
	}
	log := utils.NewParsedTxsResultLog(&txsResult.Log[msgIndex])
	event := log.GetEventByType("submit_proposal")
	if event == nil {
		panic("missing `submit_proposal` event in TxsResult log")
	}
	proposalId := event.GetAttributeByKey("proposal_id")
	if proposalId == nil {
		panic("missing `proposal_id` in `submit_proposal` event of TxsResult log")
	}

	return []command.Command{command_usecase.NewCreateMsgSubmitParamChangeProposal(
		msgCommonParams,

		model.MsgSubmitParamChangeProposalParams{
			MaybeProposalId: proposalId,
			Content:         proposalContent,
			ProposerAddress: msg["proposer"].(string),
			InitialDeposit: tmcosmosutils.MustNewCoinsFromAmountInterface(
				msg["initial_deposit"].([]interface{}),
			),
		},
	)}
}

func parseMsgSubmitCommunityFundSpendProposal(
	txSuccess bool,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
	rawContent []byte,
) []command.Command {
	var rawProposalContent model.RawMsgSubmitCommunityPoolSpendProposalContent
	if err := jsoniter.Unmarshal(rawContent, &rawProposalContent); err != nil {
		panic("error decoding community pool spend proposal content")
	}
	proposalContent := model.MsgSubmitCommunityPoolSpendProposalContent{
		Type:             rawProposalContent.Type,
		Title:            rawProposalContent.Title,
		Description:      rawProposalContent.Description,
		RecipientAddress: rawProposalContent.RecipientAddress,
		Amount:           tmcosmosutils.MustNewCoinsFromAmountInterface(rawProposalContent.Amount),
	}

	if !txSuccess {
		return []command.Command{command_usecase.NewCreateMsgSubmitCommunityPoolSpendProposal(
			msgCommonParams,

			model.MsgSubmitCommunityPoolSpendProposalParams{
				MaybeProposalId: nil,
				Content:         proposalContent,
				ProposerAddress: msg["proposer"].(string),
				InitialDeposit: tmcosmosutils.MustNewCoinsFromAmountInterface(
					msg["initial_deposit"].([]interface{}),
				),
			},
		)}
	}
	log := utils.NewParsedTxsResultLog(&txsResult.Log[msgIndex])
	// When there is no reward withdrew, `transfer` event would not exist
	event := log.GetEventByType("submit_proposal")
	if event == nil {
		panic("missing `submit_proposal` event in TxsResult log")
	}
	proposalId := event.GetAttributeByKey("proposal_id")
	if proposalId == nil {
		panic("missing `proposal_id` in `submit_proposal` event of TxsResult log")
	}

	return []command.Command{command_usecase.NewCreateMsgSubmitCommunityPoolSpendProposal(
		msgCommonParams,

		model.MsgSubmitCommunityPoolSpendProposalParams{
			MaybeProposalId: proposalId,
			Content:         proposalContent,
			ProposerAddress: msg["proposer"].(string),
			InitialDeposit: tmcosmosutils.MustNewCoinsFromAmountInterface(
				msg["initial_deposit"].([]interface{}),
			),
		},
	)}
}

func parseMsgSubmitSoftwareUpgradeProposal(
	txSuccess bool,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
	rawContent []byte,
) []command.Command {
	var rawProposalContent model.RawMsgSubmitSoftwareUpgradeProposalContent
	if err := jsoniter.Unmarshal(rawContent, &rawProposalContent); err != nil {
		panic("error decoding software upgrade proposal content")
	}

	height, err := strconv.ParseInt(rawProposalContent.Plan.Height, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("error parsing software upgrade proposal plan height: %v", err))
	}
	proposalContent := model.MsgSubmitSoftwareUpgradeProposalContent{
		Type:        rawProposalContent.Type,
		Title:       rawProposalContent.Title,
		Description: rawProposalContent.Description,
		Plan: model.MsgSubmitSoftwareUpgradeProposalPlan{
			Name:   rawProposalContent.Plan.Name,
			Time:   utctime.FromTime(rawProposalContent.Plan.Time),
			Height: height,
			Info:   rawProposalContent.Plan.Info,
		},
	}

	if !txSuccess {
		return []command.Command{command_usecase.NewCreateMsgSubmitSoftwareUpgradeProposal(
			msgCommonParams,

			model.MsgSubmitSoftwareUpgradeProposalParams{
				MaybeProposalId: nil,
				Content:         proposalContent,
				ProposerAddress: msg["proposer"].(string),
				InitialDeposit: tmcosmosutils.MustNewCoinsFromAmountInterface(
					msg["initial_deposit"].([]interface{}),
				),
			},
		)}
	}
	log := utils.NewParsedTxsResultLog(&txsResult.Log[msgIndex])
	// When there is no reward withdrew, `transfer` event would not exist
	event := log.GetEventByType("submit_proposal")
	if event == nil {
		panic("missing `submit_proposal` event in TxsResult log")
	}
	proposalId := event.GetAttributeByKey("proposal_id")
	if proposalId == nil {
		panic("missing `proposal_id` in `submit_proposal` event of TxsResult log")
	}

	return []command.Command{command_usecase.NewCreateMsgSubmitSoftwareUpgradeProposal(
		msgCommonParams,

		model.MsgSubmitSoftwareUpgradeProposalParams{
			MaybeProposalId: proposalId,
			Content:         proposalContent,
			ProposerAddress: msg["proposer"].(string),
			InitialDeposit: tmcosmosutils.MustNewCoinsFromAmountInterface(
				msg["initial_deposit"].([]interface{}),
			),
		},
	)}
}

func parseMsgSubmitCancelSoftwareUpgradeProposal(
	txSuccess bool,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
	rawContent []byte,
) []command.Command {
	var proposalContent model.MsgSubmitCancelSoftwareUpgradeProposalContent
	if err := jsoniter.Unmarshal(rawContent, &proposalContent); err != nil {
		panic("error decoding software upgrade proposal content")
	}

	if !txSuccess {
		return []command.Command{command_usecase.NewCreateMsgSubmitCancelSoftwareUpgradeProposal(
			msgCommonParams,

			model.MsgSubmitCancelSoftwareUpgradeProposalParams{
				MaybeProposalId: nil,
				Content:         proposalContent,
				ProposerAddress: msg["proposer"].(string),
				InitialDeposit: tmcosmosutils.MustNewCoinsFromAmountInterface(
					msg["initial_deposit"].([]interface{}),
				),
			},
		)}
	}
	log := utils.NewParsedTxsResultLog(&txsResult.Log[msgIndex])
	// When there is no reward withdrew, `transfer` event would not exist
	event := log.GetEventByType("submit_proposal")
	if event == nil {
		panic("missing `submit_proposal` event in TxsResult log")
	}
	proposalId := event.GetAttributeByKey("proposal_id")
	if proposalId == nil {
		panic("missing `proposal_id` in `submit_proposal` event of TxsResult log")
	}

	return []command.Command{command_usecase.NewCreateMsgSubmitCancelSoftwareUpgradeProposal(
		msgCommonParams,

		model.MsgSubmitCancelSoftwareUpgradeProposalParams{
			MaybeProposalId: proposalId,
			Content:         proposalContent,
			ProposerAddress: msg["proposer"].(string),
			InitialDeposit: tmcosmosutils.MustNewCoinsFromAmountInterface(
				msg["initial_deposit"].([]interface{}),
			),
		},
	)}
}

func parseMsgSubmitTextProposal(
	txSuccess bool,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
	rawContent []byte,
) []command.Command {
	var proposalContent model.MsgSubmitTextProposalContent
	if err := jsoniter.Unmarshal(rawContent, &proposalContent); err != nil {
		panic("error decoding text proposal content")
	}

	if !txSuccess {
		return []command.Command{command_usecase.NewCreateMsgSubmitTextProposal(
			msgCommonParams,

			model.MsgSubmitTextProposalParams{
				MaybeProposalId: nil,
				Content:         proposalContent,
				ProposerAddress: msg["proposer"].(string),
				InitialDeposit: tmcosmosutils.MustNewCoinsFromAmountInterface(
					msg["initial_deposit"].([]interface{}),
				),
			},
		)}
	}
	log := utils.NewParsedTxsResultLog(&txsResult.Log[msgIndex])
	// When there is no reward withdrew, `transfer` event would not exist
	event := log.GetEventByType("submit_proposal")
	if event == nil {
		panic("missing `submit_proposal` event in TxsResult log")
	}
	proposalId := event.GetAttributeByKey("proposal_id")
	if proposalId == nil {
		panic("missing `proposal_id` in `submit_proposal` event of TxsResult log")
	}

	return []command.Command{command_usecase.NewCreateMsgSubmitTextProposal(
		msgCommonParams,

		model.MsgSubmitTextProposalParams{
			MaybeProposalId: proposalId,
			Content:         proposalContent,
			ProposerAddress: msg["proposer"].(string),
			InitialDeposit: tmcosmosutils.MustNewCoinsFromAmountInterface(
				msg["initial_deposit"].([]interface{}),
			),
		},
	)}
}

func ParseMsgVote(
	parserParams utils.CosmosParserParams,
) []command.Command {
	return []command.Command{command_usecase.NewCreateMsgVote(
		parserParams.MsgCommonParams,

		model.MsgVoteParams{
			ProposalId: parserParams.Msg["proposal_id"].(string),
			Voter:      parserParams.Msg["voter"].(string),
			Option:     parserParams.Msg["option"].(string),
		},
	)}
}

func ParseMsgDeposit(
	parserParams utils.CosmosParserParams,
) []command.Command {
	cmds := []command.Command{command_usecase.NewCreateMsgDeposit(
		parserParams.MsgCommonParams,

		model.MsgDepositParams{
			ProposalId: parserParams.Msg["proposal_id"].(string),
			Depositor:  parserParams.Msg["depositor"].(string),
			Amount:     tmcosmosutils.MustNewCoinsFromAmountInterface(parserParams.Msg["amount"].([]interface{})),
		},
	)}

	if parserParams.MsgCommonParams.TxSuccess {
		log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])
		logEvents := log.GetEventsByType("proposal_deposit")
		if logEvents == nil {
			panic("missing `proposal_deposit` event in TxsResult log")
		}

		for _, logEvent := range logEvents {
			if logEvent.HasAttribute("voting_period_start") {
				cmds = append(cmds, command_usecase.NewStartProposalVotingPeriod(
					parserParams.MsgCommonParams.BlockHeight, logEvent.MustGetAttributeByKey("voting_period_start"),
				))
				break
			}
		}
	}

	return cmds
}

func ParseMsgDelegate(
	parserParams utils.CosmosParserParams,
) []command.Command {
	amountValue, _ := parserParams.Msg["amount"].(map[string]interface{})
	amount := tmcosmosutils.MustNewCoinFromAmountInterface(amountValue)

	if !parserParams.MsgCommonParams.TxSuccess {
		return []command.Command{command_usecase.NewCreateMsgDelegate(
			parserParams.MsgCommonParams,

			model.MsgDelegateParams{
				DelegatorAddress:   parserParams.Msg["delegator_address"].(string),
				ValidatorAddress:   parserParams.Msg["validator_address"].(string),
				Amount:             amount,
				AutoClaimedRewards: coin.Coin{},
			},
		)}
	}
	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])

	moduleAccounts := tmcosmosutils.NewModuleAccounts(parserParams.AddressPrefix)
	transferEvents := log.GetEventsByType("transfer")
	autoClaimedRewards := coin.NewZeroCoin(parserParams.StakingDenom)
	for _, transferEvent := range transferEvents {
		sender := transferEvent.MustGetAttributeByKey("sender")
		if sender != moduleAccounts.Distribution {
			continue
		}

		amount := transferEvent.MustGetAttributeByKey("amount")
		coin, coinErr := coin.ParseCoinNormalized(amount)
		if coinErr != nil {
			panic(fmt.Errorf("error parsing auto claimed rewards amount: %v", coinErr))
		}
		autoClaimedRewards = autoClaimedRewards.Add(coin)
	}

	return []command.Command{command_usecase.NewCreateMsgDelegate(
		parserParams.MsgCommonParams,

		model.MsgDelegateParams{
			DelegatorAddress:   parserParams.Msg["delegator_address"].(string),
			ValidatorAddress:   parserParams.Msg["validator_address"].(string),
			Amount:             amount,
			AutoClaimedRewards: autoClaimedRewards,
		},
	)}
}

func ParseMsgUndelegate(
	parserParams utils.CosmosParserParams,
) []command.Command {
	amountValue, _ := parserParams.Msg["amount"].(map[string]interface{})
	amount := tmcosmosutils.MustNewCoinFromAmountInterface(amountValue)

	if !parserParams.MsgCommonParams.TxSuccess {
		return []command.Command{command_usecase.NewCreateMsgUndelegate(
			parserParams.MsgCommonParams,

			model.MsgUndelegateParams{
				DelegatorAddress:      parserParams.Msg["delegator_address"].(string),
				ValidatorAddress:      parserParams.Msg["validator_address"].(string),
				MaybeUnbondCompleteAt: nil,
				Amount:                amount,
				AutoClaimedRewards:    coin.Coin{},
			},
		)}
	}
	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])
	// When there is no reward withdrew, `transfer` event would not exist
	unbondEvent := log.GetEventByType("unbond")
	if unbondEvent == nil {
		panic("missing `unbond` event in TxsResult log")
	}
	unbondCompletionTime, unbondCompletionTimeErr := utctime.Parse(
		time.RFC3339, unbondEvent.MustGetAttributeByKey("completion_time"),
	)
	if unbondCompletionTimeErr != nil {
		panic(fmt.Sprintf("error parsing unbond completion time: %v", unbondCompletionTimeErr))
	}

	moduleAccounts := tmcosmosutils.NewModuleAccounts(parserParams.AddressPrefix)
	transferEvents := log.GetEventsByType("transfer")
	autoClaimedRewards := coin.NewZeroCoin(parserParams.StakingDenom)
	for _, transferEvent := range transferEvents {
		sender := transferEvent.MustGetAttributeByKey("sender")
		if sender != moduleAccounts.Distribution {
			continue
		}

		amount := transferEvent.MustGetAttributeByKey("amount")
		coin, coinErr := coin.ParseCoinNormalized(amount)
		if coinErr != nil {
			panic(fmt.Errorf("error parsing auto claimed rewards amount: %v", coinErr))
		}
		autoClaimedRewards = autoClaimedRewards.Add(coin)
	}

	return []command.Command{command_usecase.NewCreateMsgUndelegate(
		parserParams.MsgCommonParams,

		model.MsgUndelegateParams{
			DelegatorAddress:      parserParams.Msg["delegator_address"].(string),
			ValidatorAddress:      parserParams.Msg["validator_address"].(string),
			MaybeUnbondCompleteAt: &unbondCompletionTime,
			Amount:                amount,
			AutoClaimedRewards:    autoClaimedRewards,
		},
	)}
}

func ParseMsgBeginRedelegate(
	parserParams utils.CosmosParserParams,
) []command.Command {
	amountValue, _ := parserParams.Msg["amount"].(map[string]interface{})
	amount := tmcosmosutils.MustNewCoinFromAmountInterface(amountValue)

	if !parserParams.MsgCommonParams.TxSuccess {
		return []command.Command{command_usecase.NewCreateMsgBeginRedelegate(
			parserParams.MsgCommonParams,

			model.MsgBeginRedelegateParams{
				DelegatorAddress:    parserParams.Msg["delegator_address"].(string),
				ValidatorSrcAddress: parserParams.Msg["validator_src_address"].(string),
				ValidatorDstAddress: parserParams.Msg["validator_dst_address"].(string),
				Amount:              amount,
				AutoClaimedRewards:  coin.Coin{},
			},
		)}
	}

	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])
	moduleAccounts := tmcosmosutils.NewModuleAccounts(parserParams.AddressPrefix)
	transferEvents := log.GetEventsByType("transfer")
	autoClaimedRewards := coin.NewZeroCoin(parserParams.StakingDenom)
	for _, transferEvent := range transferEvents {
		sender := transferEvent.MustGetAttributeByKey("sender")
		if sender != moduleAccounts.Distribution {
			continue
		}

		amount := transferEvent.MustGetAttributeByKey("amount")
		coin, coinErr := coin.ParseCoinNormalized(amount)
		if coinErr != nil {
			panic(fmt.Errorf("error parsing auto claimed rewards amount: %v", coinErr))
		}
		autoClaimedRewards = autoClaimedRewards.Add(coin)
	}

	return []command.Command{command_usecase.NewCreateMsgBeginRedelegate(
		parserParams.MsgCommonParams,

		model.MsgBeginRedelegateParams{
			DelegatorAddress:    parserParams.Msg["delegator_address"].(string),
			ValidatorSrcAddress: parserParams.Msg["validator_src_address"].(string),
			ValidatorDstAddress: parserParams.Msg["validator_dst_address"].(string),
			Amount:              amount,
			AutoClaimedRewards:  autoClaimedRewards,
		},
	)}
}

func ParseMsgUnjail(
	parserParams utils.CosmosParserParams,
) []command.Command {
	return []command.Command{command_usecase.NewCreateMsgUnjail(
		parserParams.MsgCommonParams,

		model.MsgUnjailParams{
			ValidatorAddr: parserParams.Msg["validator_addr"].(string),
		},
	)}
}

func parseGenesisGenTxsMsgCreateValidator(
	msg map[string]interface{},
) []command.Command {
	amountValue, _ := msg["value"].(map[string]interface{})
	amount := tmcosmosutils.MustNewCoinFromAmountInterface(amountValue)
	tendermintPubkey, _ := msg["pubkey"].(map[string]interface{})
	description := model.ValidatorDescription{
		Moniker:         "",
		Identity:        "",
		Website:         "",
		SecurityContact: "",
		Details:         "",
	}

	if descriptionJSON, ok := msg["description"].(map[string]interface{}); ok {
		description = model.ValidatorDescription{
			Moniker:         descriptionJSON["moniker"].(string),
			Identity:        descriptionJSON["identity"].(string),
			Website:         descriptionJSON["website"].(string),
			SecurityContact: descriptionJSON["security_contact"].(string),
			Details:         descriptionJSON["details"].(string),
		}
	}

	commission := model.ValidatorCommission{
		Rate:          "",
		MaxRate:       "",
		MaxChangeRate: "",
	}
	if commissionJSON, ok := msg["commission"].(map[string]interface{}); ok {
		commission = model.ValidatorCommission{
			Rate:          commissionJSON["rate"].(string),
			MaxRate:       commissionJSON["max_rate"].(string),
			MaxChangeRate: commissionJSON["max_change_rate"].(string),
		}
	}

	return []command.Command{command_usecase.NewCreateGenesisValidator(
		genesis.CreateGenesisValidatorParams{
			// Genesis validator are always bonded
			// TODO: What if gen_txs contains more validators than maximum validators
			Status:            constants.BONDED,
			Description:       description,
			Commission:        commission,
			MinSelfDelegation: msg["min_self_delegation"].(string),
			DelegatorAddress:  msg["delegator_address"].(string),
			ValidatorAddress:  msg["validator_address"].(string),
			TendermintPubkey:  tendermintPubkey["key"].(string),
			Amount:            amount,
			Jailed:            false,
		},
	)}
}

func ParseMsgCreateValidator(
	parserParams utils.CosmosParserParams,
) []command.Command {
	amountValue, _ := parserParams.Msg["value"].(map[string]interface{})
	amount := tmcosmosutils.MustNewCoinFromAmountInterface(amountValue)
	tendermintPubkey, _ := parserParams.Msg["pubkey"].(map[string]interface{})
	description := model.ValidatorDescription{
		Moniker:         "",
		Identity:        "",
		Website:         "",
		SecurityContact: "",
		Details:         "",
	}

	if descriptionJSON, ok := parserParams.Msg["description"].(map[string]interface{}); ok {
		description = model.ValidatorDescription{
			Moniker:         descriptionJSON["moniker"].(string),
			Identity:        descriptionJSON["identity"].(string),
			Website:         descriptionJSON["website"].(string),
			SecurityContact: descriptionJSON["security_contact"].(string),
			Details:         descriptionJSON["details"].(string),
		}
	}

	commission := model.ValidatorCommission{
		Rate:          "",
		MaxRate:       "",
		MaxChangeRate: "",
	}
	if commissionJSON, ok := parserParams.Msg["commission"].(map[string]interface{}); ok {
		commission = model.ValidatorCommission{
			Rate:          commissionJSON["rate"].(string),
			MaxRate:       commissionJSON["max_rate"].(string),
			MaxChangeRate: commissionJSON["max_change_rate"].(string),
		}
	}

	return []command.Command{command_usecase.NewCreateMsgCreateValidator(
		parserParams.MsgCommonParams,

		model.MsgCreateValidatorParams{
			Description:       description,
			Commission:        commission,
			MinSelfDelegation: parserParams.Msg["min_self_delegation"].(string),
			DelegatorAddress:  parserParams.Msg["delegator_address"].(string),
			ValidatorAddress:  parserParams.Msg["validator_address"].(string),
			TendermintPubkey:  tendermintPubkey["key"].(string),
			Amount:            amount,
		},
	)}
}

func ParseMsgEditValidator(
	parserParams utils.CosmosParserParams,
) []command.Command {
	var description model.ValidatorDescription
	if descriptionJSON, ok := parserParams.Msg["description"].(map[string]interface{}); ok {
		description = model.ValidatorDescription{
			Moniker:         descriptionJSON["moniker"].(string),
			Identity:        descriptionJSON["identity"].(string),
			Website:         descriptionJSON["website"].(string),
			SecurityContact: descriptionJSON["security_contact"].(string),
			Details:         descriptionJSON["details"].(string),
		}
	}

	var maybeCommissionRate *string
	if parserParams.Msg["commission_rate"] != nil {
		maybeCommissionRate = primptr.String(parserParams.Msg["commission_rate"].(string))
	}

	var maybeMinSelfDelegation *string
	if parserParams.Msg["min_self_delegation"] != nil {
		maybeMinSelfDelegation = primptr.String(parserParams.Msg["min_self_delegation"].(string))
	}

	return []command.Command{command_usecase.NewCreateMsgEditValidator(
		parserParams.MsgCommonParams,

		model.MsgEditValidatorParams{
			Description:            description,
			ValidatorAddress:       parserParams.Msg["validator_address"].(string),
			MaybeCommissionRate:    maybeCommissionRate,
			MaybeMinSelfDelegation: maybeMinSelfDelegation,
		},
	)}
}

func ParseMsgNFTIssueDenom(
	parserParams utils.CosmosParserParams,
) []command.Command {
	return []command.Command{command_usecase.NewCreateMsgNFTIssueDenom(
		parserParams.MsgCommonParams,

		model.MsgNFTIssueDenomParams{
			DenomId:   parserParams.Msg["id"].(string),
			DenomName: parserParams.Msg["name"].(string),
			Schema:    parserParams.Msg["schema"].(string),
			Sender:    parserParams.Msg["sender"].(string),
		},
	)}
}

func ParseMsgNFTMintNFT(
	parserParams utils.CosmosParserParams,
) []command.Command {
	return []command.Command{command_usecase.NewCreateMsgNFTMintNFT(
		parserParams.MsgCommonParams,

		model.MsgNFTMintNFTParams{
			DenomId:   parserParams.Msg["denom_id"].(string),
			TokenId:   parserParams.Msg["id"].(string),
			TokenName: parserParams.Msg["name"].(string),
			URI:       parserParams.Msg["uri"].(string),
			Data:      parserParams.Msg["data"].(string),
			Sender:    parserParams.Msg["sender"].(string),
			Recipient: parserParams.Msg["recipient"].(string),
		},
	)}
}

func ParseMsgNFTTransferNFT(
	parserParams utils.CosmosParserParams,
) []command.Command {
	return []command.Command{command_usecase.NewCreateMsgNFTTransferNFT(
		parserParams.MsgCommonParams,

		model.MsgNFTTransferNFTParams{
			TokenId:   parserParams.Msg["id"].(string),
			DenomId:   parserParams.Msg["denom_id"].(string),
			Sender:    parserParams.Msg["sender"].(string),
			Recipient: parserParams.Msg["recipient"].(string),
		},
	)}
}

func ParseMsgNFTEditNFT(
	parserParams utils.CosmosParserParams,
) []command.Command {
	return []command.Command{command_usecase.NewCreateMsgNFTEditNFT(
		parserParams.MsgCommonParams,

		model.MsgNFTEditNFTParams{
			DenomId:   parserParams.Msg["denom_id"].(string),
			TokenId:   parserParams.Msg["id"].(string),
			TokenName: parserParams.Msg["name"].(string),
			URI:       parserParams.Msg["uri"].(string),
			Data:      parserParams.Msg["data"].(string),
			Sender:    parserParams.Msg["sender"].(string),
		},
	)}
}

func ParseMsgNFTBurnNFT(
	parserParams utils.CosmosParserParams,
) []command.Command {
	return []command.Command{command_usecase.NewCreateMsgNFTBurnNFT(
		parserParams.MsgCommonParams,

		model.MsgNFTBurnNFTParams{
			DenomId: parserParams.Msg["denom_id"].(string),
			TokenId: parserParams.Msg["id"].(string),
			Sender:  parserParams.Msg["sender"].(string),
		},
	)}
}

func ParseMsgGrant(
	parserParams utils.CosmosParserParams,
) []command.Command {
	authType := parserParams.Msg["grant"].(map[string]interface{})["authorization"].(map[string]interface{})["@type"]

	switch authType {
	case "/cosmos.bank.v1beta1.SendAuthorization":
		return parseRawMsgSendGrant(parserParams.MsgCommonParams, parserParams.Msg)
	case "/cosmos.bank.v1beta1.StakeAuthorization":
		return parseRawMsgStackGrant(parserParams.MsgCommonParams, parserParams.Msg)
	case "/cosmos.bank.v1beta1.GenericAuthorization":
		return parseRawMsgGenericGrant(parserParams.MsgCommonParams, parserParams.Msg)
	default:
		return []command.Command{}
	}
}

func parseRawMsgSendGrant(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) []command.Command {
	var rawMsg model.RawMsgSendGrant
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			utils.StringToDurationHookFunc(),
			utils.StringToByteSliceHookFunc(),
		),
		Result: &rawMsg,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		panic(fmt.Errorf("error creating RawMsgSendGrant decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgSendGrant: %v", err))
	}

	if !msgCommonParams.TxSuccess {
		params := model.MsgGrantParams{
			MaybeSendGrant: &rawMsg,
		}

		return []command.Command{command_usecase.NewCreateMsgGrant(
			msgCommonParams,

			params,
		)}
	}

	params := model.MsgGrantParams{
		MaybeSendGrant: &rawMsg,
	}

	return []command.Command{command_usecase.NewCreateMsgGrant(
		msgCommonParams,

		params,
	)}
}

func parseRawMsgStackGrant(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) []command.Command {
	var rawMsg model.RawMsgStakeGrant
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			utils.StringToDurationHookFunc(),
			utils.StringToByteSliceHookFunc(),
		),
		Result: &rawMsg,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		panic(fmt.Errorf("error creating RawMsgStakeGrant decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgStakeGrant: %v", err))
	}

	if !msgCommonParams.TxSuccess {
		params := model.MsgGrantParams{
			MaybeStakeGrant: &rawMsg,
		}

		return []command.Command{command_usecase.NewCreateMsgGrant(
			msgCommonParams,

			params,
		)}
	}

	params := model.MsgGrantParams{
		MaybeStakeGrant: &rawMsg,
	}

	return []command.Command{command_usecase.NewCreateMsgGrant(
		msgCommonParams,

		params,
	)}
}

func parseRawMsgGenericGrant(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) []command.Command {
	var rawMsg model.RawMsgGenericGrant
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			utils.StringToDurationHookFunc(),
			utils.StringToByteSliceHookFunc(),
		),
		Result: &rawMsg,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		panic(fmt.Errorf("error creating RawMsgGenericGrant decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgGenericGrant: %v", err))
	}

	if !msgCommonParams.TxSuccess {
		params := model.MsgGrantParams{
			MaybeGenericGrant: &rawMsg,
		}

		return []command.Command{command_usecase.NewCreateMsgGrant(
			msgCommonParams,

			params,
		)}
	}

	params := model.MsgGrantParams{
		MaybeGenericGrant: &rawMsg,
	}

	return []command.Command{command_usecase.NewCreateMsgGrant(
		msgCommonParams,

		params,
	)}
}

func ParseMsgRevoke(
	parserParams utils.CosmosParserParams,
) []command.Command {
	var rawMsg model.RawMsgRevoke
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			utils.StringToDurationHookFunc(),
			utils.StringToByteSliceHookFunc(),
		),
		Result: &rawMsg,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		panic(fmt.Errorf("error creating RawMsgRevoke decoder: %v", decoderErr))
	}
	if err := decoder.Decode(parserParams.Msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgRevoke: %v", err))
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		revokeParams := model.MsgRevokeParams{
			RawMsgRevoke: rawMsg,
		}

		return []command.Command{command_usecase.NewCreateMsgRevoke(
			parserParams.MsgCommonParams,

			revokeParams,
		)}
	}

	revokeParams := model.MsgRevokeParams{
		RawMsgRevoke: rawMsg,
	}

	return []command.Command{command_usecase.NewCreateMsgRevoke(
		parserParams.MsgCommonParams,

		revokeParams,
	)}
}

func ParseMsgExec(
	parserParams utils.CosmosParserParams,
) []command.Command {
	var rawMsg model.RawMsgExec
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			utils.StringToDurationHookFunc(),
			utils.StringToByteSliceHookFunc(),
		),
		Result: &rawMsg,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		panic(fmt.Errorf("error creating ParseMsgExec decoder: %v", decoderErr))
	}
	if err := decoder.Decode(parserParams.Msg); err != nil {
		panic(fmt.Errorf("error decoding ParseMsgExec: %v", err))
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		execParams := model.MsgExecParams{
			RawMsgExec: rawMsg,
		}

		return []command.Command{command_usecase.NewCreateMsgExec(
			parserParams.MsgCommonParams,

			execParams,
		)}
	}

	execParams := model.MsgExecParams{
		RawMsgExec: rawMsg,
	}

	return []command.Command{command_usecase.NewCreateMsgExec(
		parserParams.MsgCommonParams,

		execParams,
	)}
}

func ParseMsgGrantAllowance(
	parserParams utils.CosmosParserParams,
) []command.Command {
	allowanceType := parserParams.Msg["allowance"].(map[string]interface{})["@type"]

	switch allowanceType {
	case "/cosmos.feegrant.v1beta1.BasicAllowance":
		return parseRawMsgGrantBasicAllowance(parserParams.MsgCommonParams, parserParams.Msg)
	case "/cosmos.feegrant.v1beta1.PeriodicAllowance":
		return parseRawMsgGrantPeriodicAllowance(parserParams.MsgCommonParams, parserParams.Msg)
	case "/cosmos.feegrant.v1beta1.AllowedMsgAllowance":
		return parseRawMsgGrantAllowedMsgAllowance(parserParams.MsgCommonParams, parserParams.Msg)
	default:
		return []command.Command{}
	}
}

func parseRawMsgGrantBasicAllowance(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) []command.Command {
	var rawMsg model.RawMsgGrantBasicAllowance
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			utils.StringToDurationHookFunc(),
			utils.StringToByteSliceHookFunc(),
		),
		Result: &rawMsg,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		panic(fmt.Errorf("error creating RawMsgGrantBasicAllowance decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgGrantBasicAllowance: %v", err))
	}

	if !msgCommonParams.TxSuccess {
		params := model.MsgGrantAllowanceParams{
			MaybeBasicAllowance: &rawMsg,
		}

		return []command.Command{command_usecase.NewCreateMsgGrantAllowance(
			msgCommonParams,

			params,
		)}
	}

	params := model.MsgGrantAllowanceParams{
		MaybeBasicAllowance: &rawMsg,
	}

	return []command.Command{command_usecase.NewCreateMsgGrantAllowance(
		msgCommonParams,

		params,
	)}
}

func parseRawMsgGrantPeriodicAllowance(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) []command.Command {
	var rawMsg model.RawMsgGrantPeriodicAllowance
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			utils.StringToDurationHookFunc(),
			utils.StringToByteSliceHookFunc(),
		),
		Result: &rawMsg,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		panic(fmt.Errorf("error creating RawMsgGrantPeriodicAllowance decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgGrantPeriodicAllowance: %v", err))
	}

	if !msgCommonParams.TxSuccess {
		params := model.MsgGrantAllowanceParams{
			MaybePeriodicAllowance: &rawMsg,
		}

		return []command.Command{command_usecase.NewCreateMsgGrantAllowance(
			msgCommonParams,

			params,
		)}
	}

	params := model.MsgGrantAllowanceParams{
		MaybePeriodicAllowance: &rawMsg,
	}

	return []command.Command{command_usecase.NewCreateMsgGrantAllowance(
		msgCommonParams,

		params,
	)}
}

func parseRawMsgGrantAllowedMsgAllowance(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) []command.Command {
	var rawMsg model.RawMsgGrantAllowedMsgAllowance
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			utils.StringToDurationHookFunc(),
			utils.StringToByteSliceHookFunc(),
		),
		Result: &rawMsg,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		panic(fmt.Errorf("error creating RawMsgGrantAllowedMsgAllowance decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgGrantAllowedMsgAllowance: %v", err))
	}

	if !msgCommonParams.TxSuccess {
		params := model.MsgGrantAllowanceParams{
			MaybeAllowedMsgAllowance: &rawMsg,
		}

		return []command.Command{command_usecase.NewCreateMsgGrantAllowance(
			msgCommonParams,

			params,
		)}
	}

	params := model.MsgGrantAllowanceParams{
		MaybeAllowedMsgAllowance: &rawMsg,
	}

	return []command.Command{command_usecase.NewCreateMsgGrantAllowance(
		msgCommonParams,

		params,
	)}
}

func ParseMsgRevokeAllowance(
	parserParams utils.CosmosParserParams,
) []command.Command {
	var rawMsg model.RawMsgRevokeAllowance
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			utils.StringToDurationHookFunc(),
			utils.StringToByteSliceHookFunc(),
		),
		Result: &rawMsg,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		panic(fmt.Errorf("error creating RawMsgRevokeAllowance decoder: %v", decoderErr))
	}
	if err := decoder.Decode(parserParams.Msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgRevokeAllowance: %v", err))
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		revokeAllowanceParams := model.MsgRevokeAllowanceParams{
			RawMsgRevokeAllowance: rawMsg,
		}

		return []command.Command{command_usecase.NewCreateMsgRevokeAllowance(
			parserParams.MsgCommonParams,

			revokeAllowanceParams,
		)}
	}

	revokeAllowanceParams := model.MsgRevokeAllowanceParams{
		RawMsgRevokeAllowance: rawMsg,
	}

	return []command.Command{command_usecase.NewCreateMsgRevokeAllowance(
		parserParams.MsgCommonParams,

		revokeAllowanceParams,
	)}
}

func ParseMsgCreateVestingAccount(
	parserParams utils.CosmosParserParams,
) []command.Command {
	var rawMsg model.RawMsgCreateVestingAccount
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			utils.StringToDurationHookFunc(),
			utils.StringToByteSliceHookFunc(),
		),
		Result: &rawMsg,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		panic(fmt.Errorf("error creating RawMsgCreateVestingAccount decoder: %v", decoderErr))
	}
	if err := decoder.Decode(parserParams.Msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgCreateVestingAccount: %v", err))
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		msgCreateVestingAccountParams := model.MsgCreateVestingAccountParams{
			RawMsgCreateVestingAccount: rawMsg,
		}

		return []command.Command{command_usecase.NewCreateMsgCreateVestingAccount(
			parserParams.MsgCommonParams,

			msgCreateVestingAccountParams,
		)}
	}

	msgCreateVestingAccountParams := model.MsgCreateVestingAccountParams{
		RawMsgCreateVestingAccount: rawMsg,
	}

	return []command.Command{command_usecase.NewCreateMsgCreateVestingAccount(
		parserParams.MsgCommonParams,

		msgCreateVestingAccountParams,
	)}
}
