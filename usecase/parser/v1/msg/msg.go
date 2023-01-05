package v1_parser

import (
	"fmt"
	"strconv"
	"time"

	"github.com/crypto-com/chain-indexing/entity/command"
	"github.com/crypto-com/chain-indexing/external/tmcosmosutils"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/mitchellh/mapstructure"

	"github.com/crypto-com/chain-indexing/usecase/coin"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	v1_model "github.com/crypto-com/chain-indexing/usecase/model/v1"
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

		v1_model.MsgDepositParams{
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
	var rawMsg v1_model.RawMsgSubmitProposal
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

	rawMsg.Proposer = utils.AddressParse(rawMsg.Proposer)

	var cmds []command.Command
	var possibleSignerAddresses []string

	blockHeight := parserParams.MsgCommonParams.BlockHeight

	msgs, ok := parserParams.Msg["messages"].([]interface{})
	if !ok {
		panic(fmt.Errorf("error parsing MsgSubmitProposal.msgs to []interface{}: %v", parserParams.Msg["messages"]))
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

		innerMsg["initial_deposit"] = parserParams.Msg["initial_deposit"]
		innerMsg["proposer"] = rawMsg.Proposer
		innerMsg["metadata"] = rawMsg.Metadata

		parser := parserParams.ParserManager.GetParser(utils.CosmosParserKey(innerMsgType), utils.ParserBlockHeight(blockHeight))

		msgCommands, signers := parser(utils.CosmosParserParams{
			AddressPrefix:   parserParams.AddressPrefix,
			StakingDenom:    parserParams.StakingDenom,
			TxsResult:       parserParams.TxsResult,
			MsgCommonParams: parserParams.MsgCommonParams,
			Msg:             innerMsg,
			MsgIndex:        parserParams.MsgIndex,
			ParserManager:   parserParams.ParserManager,
		})

		possibleSignerAddresses = append(possibleSignerAddresses, signers...)
		cmds = append(cmds, msgCommands...)
	}

	return cmds, possibleSignerAddresses
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

		v1_model.MsgVoteParams{
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

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if parserParams.Msg != nil {
		if voter, ok := parserParams.Msg["voter"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, utils.AddressParse(voter.(string)))
		}
	}

	return []command.Command{command_usecase.NewCreateMsgVoteV1(
		parserParams.MsgCommonParams,

		v1_model.MsgVoteParams{
			ProposalId: parserParams.Msg["proposal_id"].(string),
			Voter:      utils.AddressParse(parserParams.Msg["voter"].(string)),
			Option:     parserParams.Msg["option"].(string),
			Metadata:   parserParams.Msg["metadata"].(string),
		},
	)}, possibleSignerAddresses
}

func ParseMsgSoftwareUpgrade(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var rawProposalMsg v1_model.RawMsgSoftwareUpgrade
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			mapstructure_utils.StringToDurationHookFunc(),
			mapstructure_utils.StringToByteSliceHookFunc(),
		),
		Result: &rawProposalMsg,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		panic(fmt.Errorf("error creating ParseMsgExec decoder: %v", decoderErr))
	}
	if err := decoder.Decode(parserParams.Msg); err != nil {
		panic(fmt.Errorf("error decoding ParseMsgExec: %v", err))
	}

	height, err := strconv.ParseInt(rawProposalMsg.Plan.Height, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("error parsing software upgrade proposal plan height: %v", err))
	}
	proposalPlan := v1_model.MsgSoftwareUpgradePlan{
		Name:                rawProposalMsg.Plan.Name,
		Time:                utctime.FromTime(rawProposalMsg.Plan.Time),
		Height:              height,
		Info:                rawProposalMsg.Plan.Info,
		UpgradedClientState: rawProposalMsg.Plan.UpgradedClientState,
	}

	// Getting possible signer address from Msg
	var cmds []command.Command
	var possibleSignerAddresses []string

	if parserParams.Msg != nil {
		if proposer, ok := parserParams.Msg["proposer"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, utils.AddressParse(proposer.(string)))
		}
	}

	var metadata string
	if parserParams.Msg != nil {
		if msgMetadata, ok := parserParams.Msg["metadata"]; ok {
			metadata = msgMetadata.(string)
		}
	}

	initialDepositAmountInterface := parserParams.Msg["initial_deposit"].([]interface{})
	initialDepositAmount, err := tmcosmosutils.NewCoinsFromAmountInterface(initialDepositAmountInterface)
	if err != nil {
		initialDepositAmount = make([]coin.Coin, 0)
		for i := 0; i < len(initialDepositAmountInterface); i++ {
			initialDepositAmount = append(initialDepositAmount, coin.Coin{})
		}
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		return []command.Command{command_usecase.NewCreateMsgSoftwareUpgrade(
			parserParams.MsgCommonParams,

			v1_model.MsgSoftwareUpgradeParams{
				MaybeProposalId: nil,
				Proposer:        utils.AddressParse(parserParams.Msg["proposer"].(string)),
				InitialDeposit:  initialDepositAmount,
				Metadata:        metadata,
				Authority:       utils.AddressParse(parserParams.Msg["proposer"].(string)),
				Plan:            proposalPlan,
			},
		)}, possibleSignerAddresses
	}
	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])
	// When there is no reward withdrew, `transfer` event would not exist
	event := log.GetEventByType("submit_proposal")
	if event == nil {
		panic("missing `submit_proposal` event in TxsResult log")
	}

	if event.HasAttribute("voting_period_start") {
		cmds = append(cmds, command_usecase.NewStartProposalVotingPeriod(
			parserParams.MsgCommonParams.BlockHeight, event.MustGetAttributeByKey("voting_period_start"),
		))
	}

	proposalId := event.GetAttributeByKey("proposal_id")
	if proposalId == nil {
		panic("missing `proposal_id` in `submit_proposal` event of TxsResult log")
	}

	return append([]command.Command{
		command_usecase.NewCreateMsgSoftwareUpgrade(
			parserParams.MsgCommonParams,

			v1_model.MsgSoftwareUpgradeParams{
				MaybeProposalId: proposalId,
				Proposer:        utils.AddressParse(parserParams.Msg["proposer"].(string)),
				InitialDeposit:  initialDepositAmount,
				Metadata:        metadata,
				Authority:       utils.AddressParse(parserParams.Msg["proposer"].(string)),
				Plan:            proposalPlan,
			},
		)}, cmds...), possibleSignerAddresses
}

func ParseMsgCancelUpgrade(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var rawProposalMsg v1_model.RawMsgCancelUpgrade
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			mapstructure_utils.StringToDurationHookFunc(),
			mapstructure_utils.StringToByteSliceHookFunc(),
		),
		Result: &rawProposalMsg,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		panic(fmt.Errorf("error creating ParseMsgExec decoder: %v", decoderErr))
	}
	if err := decoder.Decode(parserParams.Msg); err != nil {
		panic(fmt.Errorf("error decoding ParseMsgExec: %v", err))
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if parserParams.Msg != nil {
		if proposer, ok := parserParams.Msg["proposer"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, utils.AddressParse(proposer.(string)))
		}
	}

	var metadata string
	if parserParams.Msg != nil {
		if msgMetadata, ok := parserParams.Msg["metadata"]; ok {
			metadata = msgMetadata.(string)
		}
	}

	initialDepositAmountInterface := parserParams.Msg["initial_deposit"].([]interface{})
	initialDepositAmount, err := tmcosmosutils.NewCoinsFromAmountInterface(initialDepositAmountInterface)
	if err != nil {
		initialDepositAmount = make([]coin.Coin, 0)
		for i := 0; i < len(initialDepositAmountInterface); i++ {
			initialDepositAmount = append(initialDepositAmount, coin.Coin{})
		}
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		return []command.Command{command_usecase.NewCreateMsgCancelUpgrade(
			parserParams.MsgCommonParams,

			v1_model.MsgCancelUpgradeParams{
				MaybeProposalId: nil,
				Authority:       parserParams.Msg["authority"].(string),
				Proposer:        utils.AddressParse(parserParams.Msg["proposer"].(string)),
				InitialDeposit:  initialDepositAmount,
				Metadata:        metadata,
			},
		)}, possibleSignerAddresses
	}
	log := utils.NewParsedTxsResultLog(&parserParams.TxsResult.Log[parserParams.MsgIndex])
	// When there is no reward withdrew, `transfer` event would not exist
	event := log.GetEventByType("submit_proposal")
	if event == nil {
		panic("missing `submit_proposal` event in TxsResult log")
	}
	proposalId := event.GetAttributeByKey("proposal_id")
	if proposalId == nil {
		panic("missing `proposal_id` in `submit_proposal` event of TxsResult log")
	}

	return []command.Command{command_usecase.NewCreateMsgCancelUpgrade(
		parserParams.MsgCommonParams,

		v1_model.MsgCancelUpgradeParams{
			MaybeProposalId: proposalId,
			Authority:       parserParams.Msg["authority"].(string),
			Proposer:        utils.AddressParse(parserParams.Msg["proposer"].(string)),
			InitialDeposit:  initialDepositAmount,
			Metadata:        parserParams.Msg["metadata"].(string),
		},
	)}, possibleSignerAddresses
}
