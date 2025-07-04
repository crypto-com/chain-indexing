package cosmos_gov_v1

import (
	"fmt"
	"strconv"
	"time"

	"github.com/crypto-com/chain-indexing/entity/command"
	"github.com/crypto-com/chain-indexing/external/tmcosmosutils"
	"github.com/mitchellh/mapstructure"

	"github.com/crypto-com/chain-indexing/usecase/coin"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	model_gov_v1 "github.com/crypto-com/chain-indexing/usecase/model/gov/v1"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"

	mapstructure_utils "github.com/crypto-com/chain-indexing/usecase/parser/utils/mapstructure"
)

func ParseMsgDeposit(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if parserParams.Msg != nil {
		if depositor, ok := parserParams.Msg["depositor"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, utils.AddressParse(depositor.(string)))
		}
	}

	amountInterface := parserParams.Msg["amount"].([]interface{})
	amount, err := tmcosmosutils.NewCoinsFromAmountInterface(amountInterface)
	if err != nil {
		amount = make([]coin.Coin, 0)
		for i := 0; i < len(amountInterface); i++ {
			amount = append(amount, coin.Coin{})
		}
	}

	cmds := []command.Command{command_usecase.NewCreateMsgDepositV1(
		parserParams.MsgCommonParams,

		model_gov_v1.MsgDepositParams{
			ProposalId: parserParams.Msg["proposal_id"].(string),
			Depositor:  utils.AddressParse(parserParams.Msg["depositor"].(string)),
			Amount:     amount,
		},
	)}

	if !parserParams.MsgCommonParams.TxSuccess {
		return cmds, possibleSignerAddresses
	}

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

	return cmds, possibleSignerAddresses
}

func ParseMsgSubmitProposal(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var rawMsg model_gov_v1.RawMsgSubmitProposal
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			mapstructure_utils.StringToDurationHookFunc(),
			mapstructure_utils.StringToByteSliceHookFunc(),
		),
		Result: &rawMsg,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		panic(fmt.Errorf("error creating ParseMsgSubmitProposal decoder: %v", decoderErr))
	}
	if err := decoder.Decode(parserParams.Msg); err != nil {
		panic(fmt.Errorf("error decoding ParseMsgSubmitProposal: %v", err))
	}

	var cmds []command.Command
	var possibleSignerAddresses []string

	rawMsg.Proposer = utils.AddressParse(rawMsg.Proposer)
	possibleSignerAddresses = append(possibleSignerAddresses, rawMsg.Proposer)

	blockHeight := parserParams.MsgCommonParams.BlockHeight

	msgs, ok := parserParams.Msg["messages"].([]interface{})
	if !ok {
		panic(fmt.Errorf("error parsing MsgSubmitProposal.msgs to []interface{}: %v", parserParams.Msg["messages"]))
	}

	initialDepositAmount, err := tmcosmosutils.NewCoinsFromAmountInterface(rawMsg.InitialDeposit)
	if err != nil {
		initialDepositAmount = make([]coin.Coin, 0)
		for i := 0; i < len(rawMsg.InitialDeposit); i++ {
			initialDepositAmount = append(initialDepositAmount, coin.Coin{})
		}
	}

	for innerMsgIndex, innerMsgInterface := range msgs {
		innerMsg, ok := innerMsgInterface.(map[string]interface{})
		if !ok {
			panic(fmt.Errorf("error parsing MsgSubmitProposal.msgs[%v] to map[string]interface{}: %v", innerMsgIndex, innerMsgInterface))
		}

		innerMsgType, ok := innerMsg["@type"].(string)
		if !ok {
			panic(fmt.Errorf("error missing '@type' in MsgSubmitProposal.msgs[%v]: %v", innerMsgIndex, innerMsg))
		}

		switch innerMsgType {
		case "/cosmos.gov.v1.MsgExecLegacyContent",
			"/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade",
			"/cosmos.upgrade.v1beta1.MsgCancelUpgrade",
			"/ethermint.feemarket.v1.MsgUpdateParams",
			"/cosmos.consensus.v1.MsgUpdateParams",
			"/ibc.core.client.v1.MsgRecoverClient",
			"/cosmos.staking.v1beta1.MsgUpdateParams",
			"/cosmos.gov.v1.MsgUpdateParams":
			break
		default:
			//nolint:gosec
			parser := parserParams.ParserManager.GetParser(utils.CosmosParserKey(innerMsgType), utils.ParserBlockHeight(blockHeight))

			msgCommands, signers := parser(utils.CosmosParserParams{
				AddressPrefix:      parserParams.AddressPrefix,
				StakingDenom:       parserParams.StakingDenom,
				TxsResult:          parserParams.TxsResult,
				MsgCommonParams:    parserParams.MsgCommonParams,
				Msg:                innerMsg,
				MsgIndex:           parserParams.MsgIndex,
				ParserManager:      parserParams.ParserManager,
				IsProposalInnerMsg: true,
			})

			possibleSignerAddresses = append(possibleSignerAddresses, signers...)
			cmds = append(cmds, msgCommands...)
		}
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		return []command.Command{command_usecase.NewCreateMsgSubmitProposal(
			parserParams.MsgCommonParams,

			model_gov_v1.MsgSubmitProposalParams{
				MaybeProposalId: nil,
				Messages:        rawMsg.Messages,
				InitialDeposit:  initialDepositAmount,
				Proposer:        rawMsg.Proposer,
				Metadata:        rawMsg.Metadata,
			},
		)}, possibleSignerAddresses
	}

	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])
	events := log.GetEventsByType("submit_proposal")

	var proposalId *string

	for _, event := range events {
		if proposalId != nil {
			break
		}

		if event.HasAttribute("msg_index") {
			msgIndex, err := strconv.Atoi(event.MustGetAttributeByKey("msg_index"))
			if err != nil {
				panic("error on parsing `msg_index` to int")
			}
			if msgIndex != parserParams.MsgIndex {
				continue
			}
		}

		if event.HasAttribute("proposal_id") {
			proposalId = event.GetAttributeByKey("proposal_id")
		}

		if event.HasAttribute("voting_period_start") {
			cmds = append(cmds, command_usecase.NewStartProposalVotingPeriod(
				parserParams.MsgCommonParams.BlockHeight, event.MustGetAttributeByKey("voting_period_start"),
			))
		}
	}

	if proposalId == nil {
		panic("missing `proposal_id` in `submit_proposal` event of TxsResult log")
	}

	return append([]command.Command{
		command_usecase.NewCreateMsgSubmitProposal(
			parserParams.MsgCommonParams,

			model_gov_v1.MsgSubmitProposalParams{
				MaybeProposalId: proposalId,
				Messages:        rawMsg.Messages,
				InitialDeposit:  initialDepositAmount,
				Proposer:        rawMsg.Proposer,
				Metadata:        rawMsg.Metadata,
			},
		),
	}, cmds...), possibleSignerAddresses
}

func ParseMsgVote(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if parserParams.Msg != nil {
		if voter, ok := parserParams.Msg["voter"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, utils.AddressParse(voter.(string)))
		}
	}

	return []command.Command{command_usecase.NewCreateMsgVoteV1(
		parserParams.MsgCommonParams,

		model_gov_v1.MsgVoteParams{
			ProposalId: parserParams.Msg["proposal_id"].(string),
			Voter:      utils.AddressParse(parserParams.Msg["voter"].(string)),
			Option:     parserParams.Msg["option"].(string),
			Metadata:   parserParams.Msg["metadata"].(string),
		},
	)}, possibleSignerAddresses
}

func ParseMsgVoteWeighted(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var rawMsg model_gov_v1.RawMsgVoteWeight
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			mapstructure_utils.StringToDurationHookFunc(),
			mapstructure_utils.StringToByteSliceHookFunc(),
		),
		Result: &rawMsg,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		panic(fmt.Errorf("error creating ParseMsgSubmitProposal decoder: %v", decoderErr))
	}
	if err := decoder.Decode(parserParams.Msg); err != nil {
		panic(fmt.Errorf("error decoding ParseMsgSubmitProposal: %v", err))
	}

	var cmds []command.Command
	var possibleSignerAddresses []string

	if rawMsg.Options != nil {
		// Getting possible signer address from Msg
		if voter, ok := parserParams.Msg["voter"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, utils.AddressParse(voter.(string)))
		}

		cmds = append(cmds, command_usecase.NewCreateMsgVoteWeightedV1(
			parserParams.MsgCommonParams,

			model_gov_v1.MsgVoteWeightedParams{
				ProposalId:  parserParams.Msg["proposal_id"].(string),
				Voter:       utils.AddressParse(parserParams.Msg["voter"].(string)),
				VoteOptions: rawMsg.Options,
				Metadata:    parserParams.Msg["metadata"].(string),
			},
		))

	}

	return cmds, possibleSignerAddresses
}
