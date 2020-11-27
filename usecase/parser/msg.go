package parser

import (
	"fmt"
	"strconv"

	"github.com/crypto-com/chainindex/internal/primptr"

	"github.com/crypto-com/chainindex/internal/utctime"

	jsoniter "github.com/json-iterator/go"

	"github.com/crypto-com/chainindex/entity/command"
	"github.com/crypto-com/chainindex/usecase/coin"
	command_usecase "github.com/crypto-com/chainindex/usecase/command"
	"github.com/crypto-com/chainindex/usecase/event"
	"github.com/crypto-com/chainindex/usecase/model"
)

func ParseBlockResultsTxsMsgToCommands(
	txDecoder *TxDecoder,
	block *model.Block,
	blockResults *model.BlockResults,
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
			switch msg["@type"] {
			case "/cosmos.bank.v1beta1.MsgSend":
				msgCommands = parseMsgSend(msgCommonParams, msg)
			case "/cosmos.bank.v1beta1.MsgMultiSend":
				msgCommands = parseMsgMultiSend(msgCommonParams, msg)
			case "/cosmos.distribution.v1beta1.MsgSetWithdrawAddress":
				msgCommands = parseMsgSetWithdrawAddress(msgCommonParams, msg)
			case "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward":
				msgCommands = parseMsgWithdrawDelegatorReward(txSuccess, txsResult, msgIndex, msgCommonParams, msg)
			case "/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission":
				msgCommands = parseMsgWithdrawValidatorCommission(txSuccess, txsResult, msgIndex, msgCommonParams, msg)
			case "/cosmos.distribution.v1beta1.MsgFundCommunityPool":
				msgCommands = parseMsgFundCommunityPool(msgCommonParams, msg)
			case "/cosmos.gov.v1beta1.MsgSubmitProposal":
				msgCommands = parseMsgSubmitProposal(txSuccess, txsResult, msgIndex, msgCommonParams, msg)
			case "/cosmos.gov.v1beta1.MsgVote":
				msgCommands = parseMsgVote(msgCommonParams, msg)
			case "/cosmos.gov.v1beta1.MsgDeposit":
				msgCommands = parseMsgDeposit(msgCommonParams, msg)
			case "/cosmos.staking.v1beta1.MsgDelegate":
				msgCommands = parseMsgDelegate(msgCommonParams, msg)
			case "/cosmos.staking.v1beta1.MsgUndelegate":
				msgCommands = parseMsgUndelegate(msgCommonParams, msg)
			case "/cosmos.staking.v1beta1.MsgBeginRedelegate":
				msgCommands = parseMsgBeginRedelegate(msgCommonParams, msg)
			case "/cosmos.slashing.v1beta1.MsgUnjail":
				msgCommands = parseMsgUnjail(msgCommonParams, msg)
			case "/cosmos.staking.v1beta1.MsgCreateValidator":
				msgCommands = parseMsgCreateValidator(msgCommonParams, msg)
			case "/cosmos.staking.v1beta1.MsgEditValidator":
				msgCommands = parseMsgEditValidator(msgCommonParams, msg)
			}

			commands = append(commands, msgCommands...)
		}
	}

	return commands, nil
}

func parseMsgSend(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) []command.Command {
	return []command.Command{command_usecase.NewCreateMsgSend(
		msgCommonParams,

		event.MsgSendCreatedParams{
			FromAddress: msg["from_address"].(string),
			ToAddress:   msg["to_address"].(string),
			Amount:      sumAmountInterfaces(msg["amount"].([]interface{})),
		},
	)}
}

func parseMsgMultiSend(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) []command.Command {
	rawInputs, _ := msg["inputs"].([]interface{})
	inputs := make([]model.MsgMultiSendInput, 0, len(rawInputs))
	for _, rawInput := range rawInputs {
		input, _ := rawInput.(map[string]interface{})
		inputs = append(inputs, model.MsgMultiSendInput{
			Address: input["address"].(string),
			Amount:  sumAmountInterfaces(input["coins"].([]interface{})),
		})
	}

	rawOutputs, _ := msg["outputs"].([]interface{})
	outputs := make([]model.MsgMultiSendOutput, 0, len(rawOutputs))
	for _, rawOutput := range rawOutputs {
		output, _ := rawOutput.(map[string]interface{})
		outputs = append(outputs, model.MsgMultiSendOutput{
			Address: output["address"].(string),
			Amount:  sumAmountInterfaces(output["coins"].([]interface{})),
		})
	}

	return []command.Command{command_usecase.NewCreateMsgMultiSend(
		msgCommonParams,

		model.MsgMultiSendParams{
			Inputs:  inputs,
			Outputs: outputs,
		},
	)}
}

func parseMsgSetWithdrawAddress(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) []command.Command {
	return []command.Command{command_usecase.NewCreateMsgSetWithdrawAddress(
		msgCommonParams,

		model.MsgSetWithdrawAddressParams{
			DelegatorAddress: msg["delegator_address"].(string),
			WithdrawAddress:  msg["withdraw_address"].(string),
		},
	)}
}

func parseMsgWithdrawDelegatorReward(
	txSuccess bool,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) []command.Command {
	if !txSuccess {
		delegatorAddress, _ := msg["delegator_address"].(string)
		return []command.Command{command_usecase.NewCreateMsgWithdrawDelegatorReward(
			msgCommonParams,

			model.MsgWithdrawDelegatorRewardParams{
				DelegatorAddress: delegatorAddress,
				ValidatorAddress: msg["validator_address"].(string),
				RecipientAddress: delegatorAddress,
				Amount:           coin.Zero(),
			},
		)}
	}
	log := NewParsedTxsResultLog(&txsResult.Log[msgIndex])
	var recipient string
	var amount coin.Coin
	// When there is no reward withdrew, `transfer` event would not exist
	if event := log.GetEventByType("transfer"); event == nil {
		recipient, _ = msg["delegator_address"].(string)
		amount = coin.Zero()
	} else {
		recipient = event.MustGetAttributeByKey("recipient")
		amountValue := event.MustGetAttributeByKey("amount")
		amount = coin.MustNewCoinFromString(TrimAmountDenom(amountValue))
	}

	return []command.Command{command_usecase.NewCreateMsgWithdrawDelegatorReward(
		msgCommonParams,

		model.MsgWithdrawDelegatorRewardParams{
			DelegatorAddress: msg["delegator_address"].(string),
			ValidatorAddress: msg["validator_address"].(string),
			RecipientAddress: recipient,
			Amount:           amount,
		},
	)}
}

func parseMsgWithdrawValidatorCommission(
	txSuccess bool,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) []command.Command {
	if !txSuccess {
		return []command.Command{command_usecase.NewCreateMsgWithdrawValidatorCommission(
			msgCommonParams,

			model.MsgWithdrawValidatorCommissionParams{
				ValidatorAddress: msg["validator_address"].(string),
				RecipientAddress: msg["delegator_address"].(string),
				Amount:           coin.Zero(),
			},
		)}
	}
	log := NewParsedTxsResultLog(&txsResult.Log[msgIndex])
	var recipient string
	var amount coin.Coin
	// When there is no reward withdrew, `transfer` event would not exist
	if event := log.GetEventByType("transfer"); event == nil {
		recipient, _ = msg["delegator_address"].(string)
		amount = coin.Zero()
	} else {
		recipient = event.MustGetAttributeByKey("recipient")
		amountValue := event.MustGetAttributeByKey("amount")
		amount = coin.MustNewCoinFromString(TrimAmountDenom(amountValue))
	}

	return []command.Command{command_usecase.NewCreateMsgWithdrawValidatorCommission(
		msgCommonParams,

		model.MsgWithdrawValidatorCommissionParams{
			ValidatorAddress: msg["validator_address"].(string),
			RecipientAddress: recipient,
			Amount:           amount,
		},
	)}
}

func parseMsgFundCommunityPool(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) []command.Command {
	return []command.Command{command_usecase.NewCreateMsgFundCommunityPool(
		msgCommonParams,

		model.MsgFundCommunityPoolParams{
			Depositor: msg["depositor"].(string),
			Amount:    sumAmountInterfaces(msg["amount"].([]interface{})),
		},
	)}
}

func parseMsgSubmitProposal(
	txSuccess bool,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) []command.Command {
	rawContent, err := jsoniter.Marshal(msg["content"])
	if err != nil {
		panic(fmt.Sprintf("error encoding proposal content: %v", err))
	}
	var proposalContent model.MsgSubmitProposalContent
	if err := jsoniter.Unmarshal(rawContent, &proposalContent); err != nil {
		panic(fmt.Sprintf("error decoding proposal content: %v", err))
	}

	if proposalContent.Type == "/cosmos.params.v1beta1.ParameterChangeProposal" {
		return parseMsgSubmitParamChangeProposal(txSuccess, txsResult, msgIndex, msgCommonParams, msg, rawContent)
	} else if proposalContent.Type == "/cosmos.distribution.v1beta1.CommunityPoolSpendProposal" {
		return parseMsgSubmitCommunityFundSpendProposal(txSuccess, txsResult, msgIndex, msgCommonParams, msg, rawContent)
	} else if proposalContent.Type == "/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal" {
		return parseMsgSubmitSoftwareUpgradeProposal(txSuccess, txsResult, msgIndex, msgCommonParams, msg, rawContent)
	} else if proposalContent.Type == "/cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal" {
		return parseMsgSubmitCancelSoftwareUpgradeProposal(txSuccess, txsResult, msgIndex, msgCommonParams, msg, rawContent)
	}
	panic(fmt.Sprintf("unrecognzied govenance proposal type `%s`", proposalContent.Type))
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
				InitialDeposit:  sumAmountInterfaces(msg["initial_deposit"].([]interface{})),
			},
		)}
	}
	log := NewParsedTxsResultLog(&txsResult.Log[msgIndex])
	// When there is no reward withdrew, `transfer` event would not exist
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
			InitialDeposit:  sumAmountInterfaces(msg["initial_deposit"].([]interface{})),
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
		Amount:           sumAmountInterfaces(rawProposalContent.Amount),
	}

	if !txSuccess {
		return []command.Command{command_usecase.NewCreateMsgSubmitCommunityPoolSpendProposal(
			msgCommonParams,

			model.MsgSubmitCommunityPoolSpendProposalParams{
				MaybeProposalId: nil,
				Content:         proposalContent,
				ProposerAddress: msg["proposer"].(string),
				InitialDeposit:  sumAmountInterfaces(msg["initial_deposit"].([]interface{})),
			},
		)}
	}
	log := NewParsedTxsResultLog(&txsResult.Log[msgIndex])
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
			InitialDeposit:  sumAmountInterfaces(msg["initial_deposit"].([]interface{})),
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
				InitialDeposit:  sumAmountInterfaces(msg["initial_deposit"].([]interface{})),
			},
		)}
	}
	log := NewParsedTxsResultLog(&txsResult.Log[msgIndex])
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
			InitialDeposit:  sumAmountInterfaces(msg["initial_deposit"].([]interface{})),
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
				InitialDeposit:  sumAmountInterfaces(msg["initial_deposit"].([]interface{})),
			},
		)}
	}
	log := NewParsedTxsResultLog(&txsResult.Log[msgIndex])
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
			InitialDeposit:  sumAmountInterfaces(msg["initial_deposit"].([]interface{})),
		},
	)}
}

func parseMsgVote(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) []command.Command {
	return []command.Command{command_usecase.NewCreateMsgVote(
		msgCommonParams,

		model.MsgVoteParams{
			ProposalId: msg["proposal_id"].(string),
			Voter:      msg["voter"].(string),
			Option:     msg["option"].(string),
		},
	)}
}

func parseMsgDeposit(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) []command.Command {
	return []command.Command{command_usecase.NewCreateMsgDeposit(
		msgCommonParams,

		model.MsgDepositParams{
			ProposalId: msg["proposal_id"].(string),
			Depositor:  msg["depositor"].(string),
			Amount:     sumAmountInterfaces(msg["amount"].([]interface{})),
		},
	)}
}

func parseMsgDelegate(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) []command.Command {
	amountValue, _ := msg["amount"].(map[string]interface{})
	amount := coin.MustNewCoinFromString(amountValue["amount"].(string))

	return []command.Command{command_usecase.NewCreateMsgDelegate(
		msgCommonParams,

		model.MsgDelegateParams{
			DelegatorAddress: msg["delegator_address"].(string),
			ValidatorAddress: msg["validator_address"].(string),
			Amount:           amount,
		},
	)}
}

func parseMsgUndelegate(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) []command.Command {
	amountValue, _ := msg["amount"].(map[string]interface{})
	amount := coin.MustNewCoinFromString(amountValue["amount"].(string))

	return []command.Command{command_usecase.NewCreateMsgUndelegate(
		msgCommonParams,

		model.MsgUndelegateParams{
			DelegatorAddress: msg["delegator_address"].(string),
			ValidatorAddress: msg["validator_address"].(string),
			Amount:           amount,
		},
	)}
}

func parseMsgBeginRedelegate(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) []command.Command {
	amountValue, _ := msg["amount"].(map[string]interface{})
	amount := coin.MustNewCoinFromString(amountValue["amount"].(string))

	return []command.Command{command_usecase.NewCreateMsgBeginRedelegate(
		msgCommonParams,

		model.MsgBeginRedelegateParams{
			DelegatorAddress:    msg["delegator_address"].(string),
			ValidatorSrcAddress: msg["validator_src_address"].(string),
			ValidatorDstAddress: msg["validator_dst_address"].(string),
			Amount:              amount,
		},
	)}
}

func parseMsgUnjail(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) []command.Command {
	return []command.Command{command_usecase.NewCreateMsgUnjail(
		msgCommonParams,

		model.MsgUnjailParams{
			ValidatorAddr: msg["validator_addr"].(string),
		},
	)}
}

func parseMsgCreateValidator(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) []command.Command {
	amountValue, _ := msg["value"].(map[string]interface{})
	amount := coin.MustNewCoinFromString(amountValue["amount"].(string))

	description := model.MsgValidatorDescription{
		Moniker:         "",
		Identity:        "",
		Website:         "",
		SecurityContact: "",
		Details:         "",
	}

	if descriptionJSON, ok := msg["description"].(map[string]interface{}); ok {
		description = model.MsgValidatorDescription{
			Moniker:         descriptionJSON["moniker"].(string),
			Identity:        descriptionJSON["identity"].(string),
			Website:         descriptionJSON["website"].(string),
			SecurityContact: descriptionJSON["security_contact"].(string),
			Details:         descriptionJSON["details"].(string),
		}
	}

	commission := model.MsgValidatorCommission{
		Rate:          "",
		MaxRate:       "",
		MaxChangeRate: "",
	}
	if commissionJSON, ok := msg["commission"].(map[string]interface{}); ok {
		commission = model.MsgValidatorCommission{
			Rate:          commissionJSON["rate"].(string),
			MaxRate:       commissionJSON["max_rate"].(string),
			MaxChangeRate: commissionJSON["max_change_rate"].(string),
		}
	}

	return []command.Command{command_usecase.NewCreateMsgCreateValidator(
		msgCommonParams,

		model.MsgCreateValidatorParams{
			Description:       description,
			Commission:        commission,
			MinSelfDelegation: msg["min_self_delegation"].(string),
			DelegatorAddress:  msg["delegator_address"].(string),
			ValidatorAddress:  msg["validator_address"].(string),
			Pubkey:            msg["pubkey"].(string),
			Amount:            amount,
		},
	)}
}

func parseMsgEditValidator(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) []command.Command {
	var description model.MsgValidatorDescription
	if descriptionJSON, ok := msg["description"].(map[string]interface{}); ok {
		description = model.MsgValidatorDescription{
			Moniker:         descriptionJSON["moniker"].(string),
			Identity:        descriptionJSON["identity"].(string),
			Website:         descriptionJSON["website"].(string),
			SecurityContact: descriptionJSON["security_contact"].(string),
			Details:         descriptionJSON["details"].(string),
		}
	}

	var maybeCommissionRate *string
	if msg["commission_rate"] != nil {
		maybeCommissionRate = primptr.String(msg["commission_rate"].(string))
	}

	var maybeMinSelfDelegation *string
	if msg["min_self_delegation"] != nil {
		maybeMinSelfDelegation = primptr.String(msg["min_self_delegation"].(string))
	}

	return []command.Command{command_usecase.NewCreateMsgEditValidator(
		msgCommonParams,

		model.MsgEditValidatorParams{
			Description:            description,
			ValidatorAddress:       msg["validator_address"].(string),
			MaybeCommissionRate:    maybeCommissionRate,
			MaybeMinSelfDelegation: maybeMinSelfDelegation,
		},
	)}
}

func sumAmountInterfaces(amounts []interface{}) coin.Coin {
	sum := coin.Zero()
	for _, rawAmount := range amounts {
		amount, _ := rawAmount.(map[string]interface{})
		sum, _ = sum.Add(coin.MustNewCoinFromString(amount["amount"].(string)))
	}

	return sum
}
