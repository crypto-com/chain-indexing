package parser

import (
	"fmt"
	"strconv"

	"github.com/crypto-com/chain-indexing/entity/command"
	"github.com/crypto-com/chain-indexing/external/tmcosmosutils"
	"github.com/crypto-com/chain-indexing/external/utctime"
	jsoniter "github.com/json-iterator/go"

	"github.com/crypto-com/chain-indexing/usecase/coin"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
)

func ParseMsgSubmitProposal(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	rawContent, err := jsoniter.Marshal(parserParams.Msg["content"])
	if err != nil {
		panic(fmt.Sprintf("error encoding proposal content: %v", err))
	}
	var proposalContent model.MsgSubmitProposalContent
	if err := jsoniter.Unmarshal(rawContent, &proposalContent); err != nil {
		panic(fmt.Sprintf("error decoding proposal content: %v", err))
	}

	var cmds []command.Command
	var possibleSignerAddresses []string
	if proposalContent.Type == "/cosmos.params.v1beta1.ParameterChangeProposal" {
		// cmds, possibleSignerAddresses = parseMsgSubmitParamChangeProposal(parserParams.MsgCommonParams.TxSuccess, parserParams.TxsResult, parserParams.MsgIndex, parserParams.MsgCommonParams, parserParams.Msg, rawContent)
	} else if proposalContent.Type == "/cosmos.distribution.v1beta1.CommunityPoolSpendProposal" {
		// cmds, possibleSignerAddresses = parseMsgSubmitCommunityFundSpendProposal(parserParams.MsgCommonParams.TxSuccess, parserParams.TxsResult, parserParams.MsgIndex, parserParams.MsgCommonParams, parserParams.Msg, rawContent)
	} else if proposalContent.Type == "/cosmos.upgrade.v1beta1.SoftwareUpgrade" {
		cmds, possibleSignerAddresses = parseMsgSubmitSoftwareUpgrade(parserParams.MsgCommonParams.TxSuccess, parserParams.TxsResult, parserParams.MsgIndex, parserParams.MsgCommonParams, parserParams.Msg, rawContent)
	} else if proposalContent.Type == "/cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal" {
		// cmds, possibleSignerAddresses = parseMsgSubmitCancelSoftwareUpgradeProposal(parserParams.MsgCommonParams.TxSuccess, parserParams.TxsResult, parserParams.MsgIndex, parserParams.MsgCommonParams, parserParams.Msg, rawContent)
	} else if proposalContent.Type == "/cosmos.gov.v1beta1.TextProposal" {
		// cmds, possibleSignerAddresses = parseMsgSubmitTextProposal(parserParams.MsgCommonParams.TxSuccess, parserParams.TxsResult, parserParams.MsgIndex, parserParams.MsgCommonParams, parserParams.Msg, rawContent)
	} else {
		parserParams.Logger.Errorf("unrecognzied govenance proposal type `%s`", proposalContent.Type)
		fmt.Println("===> 2a. proposalContent.Type: ", proposalContent.Type)

		// cmds, possibleSignerAddresses = parseMsgSubmitUnknownProposal(parserParams.MsgCommonParams.TxSuccess, parserParams.TxsResult, parserParams.MsgIndex, parserParams.MsgCommonParams, parserParams.Msg, rawContent)
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		return cmds, possibleSignerAddresses
	}

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

	return cmds, possibleSignerAddresses
}

func parseMsgSubmitSoftwareUpgrade(
	txSuccess bool,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
	rawContent []byte,
) ([]command.Command, []string) {
	var rawProposalContent model.RawMsgSubmitSoftwareUpgradeContent
	if err := jsoniter.Unmarshal(rawContent, &rawProposalContent); err != nil {
		panic("error decoding software upgrade proposal content")
	}

	height, err := strconv.ParseInt(rawProposalContent.Plan.Height, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("error parsing software upgrade proposal plan height: %v", err))
	}
	proposalContent := model.MsgSubmitSoftwareUpgradeProposalContent{
		Type: rawProposalContent.Type,
		Plan: model.MsgSubmitSoftwareUpgradeProposalPlan{
			Name:   rawProposalContent.Plan.Name,
			Time:   utctime.FromTime(rawProposalContent.Plan.Time),
			Height: height,
			Info:   rawProposalContent.Plan.Info,
		},
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if msg != nil {
		if proposer, ok := msg["proposer"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, utils.AddressParse(proposer.(string)))
		}
	}

	initialDepositAmountInterface := msg["initial_deposit"].([]interface{})
	initialDepositAmount, err := tmcosmosutils.NewCoinsFromAmountInterface(initialDepositAmountInterface)
	if err != nil {
		initialDepositAmount = make([]coin.Coin, 0)
		for i := 0; i < len(initialDepositAmountInterface); i++ {
			initialDepositAmount = append(initialDepositAmount, coin.Coin{})
		}
	}

	if !txSuccess {
		return []command.Command{command_usecase.NewCreateMsgSubmitSoftwareUpgradeProposal(
			msgCommonParams,

			model.MsgSubmitSoftwareUpgradeParams{
				MaybeProposalId: nil,
				Metadata:        "",
				Messages:        proposalContent,
				ProposerAddress: utils.AddressParse(msg["proposer"].(string)),
				InitialDeposit:  initialDepositAmount,
			},
		)}, possibleSignerAddresses
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

		model.MsgSubmitSoftwareUpgradeParams{
			MaybeProposalId: proposalId,
			Metadata:        "",
			Messages:        proposalContent,
			ProposerAddress: utils.AddressParse(msg["proposer"].(string)),
			InitialDeposit:  initialDepositAmount,
		},
	)}, possibleSignerAddresses
}

func parseMsgSubmitCancelSoftwareUpgrade(
	txSuccess bool,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
	rawContent []byte,
) ([]command.Command, []string) {
	var proposalContent model.MsgSubmitCancelSoftwareUpgradeMessages
	if err := jsoniter.Unmarshal(rawContent, &proposalContent); err != nil {
		panic("error decoding software upgrade proposal content")
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if msg != nil {
		if proposer, ok := msg["proposer"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, utils.AddressParse(proposer.(string)))
		}
	}

	var metadata string
	if msg != nil {
		if msgMetadata, ok := msg["metadata"]; ok {
			metadata = msgMetadata
		}
	}

	initialDepositAmountInterface := msg["initial_deposit"].([]interface{})
	initialDepositAmount, err := tmcosmosutils.NewCoinsFromAmountInterface(initialDepositAmountInterface)
	if err != nil {
		initialDepositAmount = make([]coin.Coin, 0)
		for i := 0; i < len(initialDepositAmountInterface); i++ {
			initialDepositAmount = append(initialDepositAmount, coin.Coin{})
		}
	}

	if !txSuccess {
		return []command.Command{command_usecase.NewCreateMsgSubmitCancelSoftwareUpgradeProposal(
			msgCommonParams,

			model.MsgSubmitCancelSoftwareUpgradeParams{
				MaybeProposalId: nil,
				Metadata:        metadata,
				Messages:        proposalContent,
				ProposerAddress: utils.AddressParse(msg["proposer"].(string)),
				InitialDeposit:  initialDepositAmount,
			},
		)}, possibleSignerAddresses
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
			ProposerAddress: utils.AddressParse(msg["proposer"].(string)),
			InitialDeposit:  initialDepositAmount,
		},
	)}, possibleSignerAddresses
}
