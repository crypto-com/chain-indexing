package parser

import (
	"fmt"
	"strconv"
	"time"

	"github.com/crypto-com/chain-indexing/entity/command"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"

	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/external/tmcosmosutils"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/projection/validator/constants"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/crypto-com/chain-indexing/usecase/model/genesis"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
	mapstructure_utils "github.com/crypto-com/chain-indexing/usecase/parser/utils/mapstructure"
)

func ParseBlockTxsMsgToCommands(
	parserManager *utils.CosmosParserManager,
	blockHeight int64,
	blockResults *model.BlockResults,
	txs []model.Tx,
	accountAddressPrefix string,
	stakingDenom string,
) ([]command.Command, []string, error) {
	commands := make([]command.Command, 0)
	var addresses []string

	for i, tx := range txs {
		txHash := tx.TxResponse.TxHash
		txSuccess := true
		txsResult := blockResults.TxsResults[i]

		if txsResult.Code != 0 {
			txSuccess = false
		}

		for msgIndex, msg := range tx.Tx.Body.Messages {
			msgCommonParams := event.MsgCommonParams{
				BlockHeight: blockHeight,
				TxHash:      txHash,
				TxSuccess:   txSuccess,
				MsgIndex:    msgIndex,
			}

			var msgCommands []command.Command
			var possibleSignerAddresses []string

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
				"/ibc.core.channel.v1.MsgChannelCloseInit",
				"/ibc.core.channel.v1.MsgChannelCloseConfirm",

				// ibc applications transfer
				"/ibc.applications.transfer.v1.MsgTransfer",

				// cosmos authz
				"/cosmos.authz.v1beta1.MsgGrant",
				"/cosmos.authz.v1beta1.MsgRevoke",
				// FIXME: https://github.com/crypto-com/chain-indexing/issues/673
				//"/cosmos.authz.v1beta1.MsgExec",

				// cosmos feegrant
				"/cosmos.feegrant.v1beta1.MsgGrantAllowance",
				"/cosmos.feegrant.v1beta1.MsgRevokeAllowance",

				// cosmos vesting
				"/cosmos.vesting.v1beta1.MsgCreateVestingAccount",

				// ethermint evm
				"/ethermint.evm.v1.MsgEthereumTx":
				parser := parserManager.GetParser(utils.CosmosParserKey(msgType.(string)), utils.ParserBlockHeight(blockHeight))

				msgCommands, possibleSignerAddresses = parser(utils.CosmosParserParams{
					AddressPrefix:   accountAddressPrefix,
					StakingDenom:    stakingDenom,
					TxsResult:       txsResult,
					MsgCommonParams: msgCommonParams,
					Msg:             msg,
					MsgIndex:        msgIndex,
					ParserManager:   parserManager,
					Logger:          parserManager.GetLogger(),
				})
			}
			addresses = append(addresses, possibleSignerAddresses...)
			commands = append(commands, msgCommands...)
		}
	}
	return commands, addresses, nil
}

func ParseMsgSend(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if parserParams.Msg != nil {
		if fromAddress, ok := parserParams.Msg["from_address"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, fromAddress.(string))
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

	return []command.Command{command_usecase.NewCreateMsgSend(
		parserParams.MsgCommonParams,

		event.MsgSendCreatedParams{
			FromAddress: parserParams.Msg["from_address"].(string),
			ToAddress:   parserParams.Msg["to_address"].(string),
			Amount:      amount,
		},
	)}, possibleSignerAddresses
}

func ParseMsgMultiSend(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	rawInputs, _ := parserParams.Msg["inputs"].([]interface{})
	inputs := make([]model.MsgMultiSendInput, 0, len(rawInputs))
	possibleSignerAddresses := make([]string, 0, len(rawInputs))
	for _, rawInput := range rawInputs {
		input, _ := rawInput.(map[string]interface{})

		if fromAddress, ok := input["address"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, fromAddress.(string))
		}

		amountInterface := input["coins"].([]interface{})
		amount, err := tmcosmosutils.NewCoinsFromAmountInterface(amountInterface)
		if err != nil {
			amount = make([]coin.Coin, 0)
			for i := 0; i < len(amountInterface); i++ {
				amount = append(amount, coin.Coin{})
			}
		}

		inputs = append(inputs, model.MsgMultiSendInput{
			Address: input["address"].(string),
			Amount:  amount,
		})
	}

	rawOutputs, _ := parserParams.Msg["outputs"].([]interface{})
	outputs := make([]model.MsgMultiSendOutput, 0, len(rawOutputs))
	for _, rawOutput := range rawOutputs {
		output, _ := rawOutput.(map[string]interface{})

		amountInterface := output["coins"].([]interface{})
		amount, err := tmcosmosutils.NewCoinsFromAmountInterface(amountInterface)
		if err != nil {
			amount = make([]coin.Coin, 0)
			for i := 0; i < len(amountInterface); i++ {
				amount = append(amount, coin.Coin{})
			}
		}

		outputs = append(outputs, model.MsgMultiSendOutput{
			Address: output["address"].(string),
			Amount:  amount,
		})
	}

	return []command.Command{command_usecase.NewCreateMsgMultiSend(
		parserParams.MsgCommonParams,

		model.MsgMultiSendParams{
			Inputs:  inputs,
			Outputs: outputs,
		},
	)}, possibleSignerAddresses
}

func ParseMsgSetWithdrawAddress(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if parserParams.Msg != nil {
		if delegatorAddress, ok := parserParams.Msg["delegator_address"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, delegatorAddress.(string))
		}
	}

	return []command.Command{command_usecase.NewCreateMsgSetWithdrawAddress(
		parserParams.MsgCommonParams,

		model.MsgSetWithdrawAddressParams{
			DelegatorAddress: parserParams.Msg["delegator_address"].(string),
			WithdrawAddress:  parserParams.Msg["withdraw_address"].(string),
		},
	)}, possibleSignerAddresses
}

func ParseMsgWithdrawDelegatorReward(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if parserParams.Msg != nil {
		if delegatorAddress, ok := parserParams.Msg["delegator_address"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, delegatorAddress.(string))
		}
	}

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
		)}, possibleSignerAddresses
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
	)}, possibleSignerAddresses
}

func ParseMsgWithdrawValidatorCommission(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if parserParams.Msg != nil {
		if validatorAddress, ok := parserParams.Msg["validator_address"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, validatorAddress.(string))
		}
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		return []command.Command{command_usecase.NewCreateMsgWithdrawValidatorCommission(
			parserParams.MsgCommonParams,

			model.MsgWithdrawValidatorCommissionParams{
				ValidatorAddress: parserParams.Msg["validator_address"].(string),
				RecipientAddress: "",
				Amount:           coin.NewEmptyCoins(),
			},
		)}, possibleSignerAddresses
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
	)}, possibleSignerAddresses
}

func ParseMsgFundCommunityPool(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if parserParams.Msg != nil {
		if depositor, ok := parserParams.Msg["depositor"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, depositor.(string))
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

	return []command.Command{command_usecase.NewCreateMsgFundCommunityPool(
		parserParams.MsgCommonParams,

		model.MsgFundCommunityPoolParams{
			Depositor: parserParams.Msg["depositor"].(string),
			Amount:    amount,
		},
	)}, possibleSignerAddresses
}

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
		cmds, possibleSignerAddresses = parseMsgSubmitParamChangeProposal(parserParams.MsgCommonParams.TxSuccess, parserParams.TxsResult, parserParams.MsgIndex, parserParams.MsgCommonParams, parserParams.Msg, rawContent)
	} else if proposalContent.Type == "/cosmos.distribution.v1beta1.CommunityPoolSpendProposal" {
		cmds, possibleSignerAddresses = parseMsgSubmitCommunityFundSpendProposal(parserParams.MsgCommonParams.TxSuccess, parserParams.TxsResult, parserParams.MsgIndex, parserParams.MsgCommonParams, parserParams.Msg, rawContent)
	} else if proposalContent.Type == "/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal" {
		cmds, possibleSignerAddresses = parseMsgSubmitSoftwareUpgradeProposal(parserParams.MsgCommonParams.TxSuccess, parserParams.TxsResult, parserParams.MsgIndex, parserParams.MsgCommonParams, parserParams.Msg, rawContent)
	} else if proposalContent.Type == "/cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal" {
		cmds, possibleSignerAddresses = parseMsgSubmitCancelSoftwareUpgradeProposal(parserParams.MsgCommonParams.TxSuccess, parserParams.TxsResult, parserParams.MsgIndex, parserParams.MsgCommonParams, parserParams.Msg, rawContent)
	} else if proposalContent.Type == "/cosmos.gov.v1beta1.TextProposal" {
		cmds, possibleSignerAddresses = parseMsgSubmitTextProposal(parserParams.MsgCommonParams.TxSuccess, parserParams.TxsResult, parserParams.MsgIndex, parserParams.MsgCommonParams, parserParams.Msg, rawContent)
	} else {
		parserParams.Logger.Errorf("unrecognzied govenance proposal type `%s`", proposalContent.Type)
		cmds, possibleSignerAddresses = parseMsgSubmitUnknownProposal(parserParams.MsgCommonParams.TxSuccess, parserParams.TxsResult, parserParams.MsgIndex, parserParams.MsgCommonParams, parserParams.Msg, rawContent)
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

func parseMsgSubmitParamChangeProposal(
	txSuccess bool,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
	rawContent []byte,
) ([]command.Command, []string) {
	var proposalContent model.MsgSubmitParamChangeProposalContent
	if err := jsoniter.Unmarshal(rawContent, &proposalContent); err != nil {
		panic("error decoding param change proposal content")
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if msg != nil {
		if proposer, ok := msg["proposer"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, proposer.(string))
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
		return []command.Command{command_usecase.NewCreateMsgSubmitParamChangeProposal(
			msgCommonParams,

			model.MsgSubmitParamChangeProposalParams{
				MaybeProposalId: nil,
				Content:         proposalContent,
				ProposerAddress: msg["proposer"].(string),
				InitialDeposit:  initialDepositAmount,
			},
		)}, possibleSignerAddresses
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
			InitialDeposit:  initialDepositAmount,
		},
	)}, possibleSignerAddresses
}

func parseMsgSubmitCommunityFundSpendProposal(
	txSuccess bool,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
	rawContent []byte,
) ([]command.Command, []string) {
	var rawProposalContent model.RawMsgSubmitCommunityPoolSpendProposalContent
	if err := jsoniter.Unmarshal(rawContent, &rawProposalContent); err != nil {
		panic("error decoding community pool spend proposal content")
	}
	amount, err := tmcosmosutils.NewCoinsFromAmountInterface(rawProposalContent.Amount)
	if err != nil {
		amount = make([]coin.Coin, 0)
		for i := 0; i < len(rawProposalContent.Amount); i++ {
			amount = append(amount, coin.Coin{})
		}
	}
	proposalContent := model.MsgSubmitCommunityPoolSpendProposalContent{
		Type:             rawProposalContent.Type,
		Title:            rawProposalContent.Title,
		Description:      rawProposalContent.Description,
		RecipientAddress: rawProposalContent.RecipientAddress,
		Amount:           amount,
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if msg != nil {
		if proposer, ok := msg["proposer"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, proposer.(string))
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
		return []command.Command{command_usecase.NewCreateMsgSubmitCommunityPoolSpendProposal(
			msgCommonParams,

			model.MsgSubmitCommunityPoolSpendProposalParams{
				MaybeProposalId: nil,
				Content:         proposalContent,
				ProposerAddress: msg["proposer"].(string),
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

	return []command.Command{command_usecase.NewCreateMsgSubmitCommunityPoolSpendProposal(
		msgCommonParams,

		model.MsgSubmitCommunityPoolSpendProposalParams{
			MaybeProposalId: proposalId,
			Content:         proposalContent,
			ProposerAddress: msg["proposer"].(string),
			InitialDeposit:  initialDepositAmount,
		},
	)}, possibleSignerAddresses
}

func parseMsgSubmitSoftwareUpgradeProposal(
	txSuccess bool,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
	rawContent []byte,
) ([]command.Command, []string) {
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

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if msg != nil {
		if proposer, ok := msg["proposer"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, proposer.(string))
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

			model.MsgSubmitSoftwareUpgradeProposalParams{
				MaybeProposalId: nil,
				Content:         proposalContent,
				ProposerAddress: msg["proposer"].(string),
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

		model.MsgSubmitSoftwareUpgradeProposalParams{
			MaybeProposalId: proposalId,
			Content:         proposalContent,
			ProposerAddress: msg["proposer"].(string),
			InitialDeposit:  initialDepositAmount,
		},
	)}, possibleSignerAddresses
}

func parseMsgSubmitCancelSoftwareUpgradeProposal(
	txSuccess bool,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
	rawContent []byte,
) ([]command.Command, []string) {
	var proposalContent model.MsgSubmitCancelSoftwareUpgradeProposalContent
	if err := jsoniter.Unmarshal(rawContent, &proposalContent); err != nil {
		panic("error decoding software upgrade proposal content")
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if msg != nil {
		if proposer, ok := msg["proposer"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, proposer.(string))
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

			model.MsgSubmitCancelSoftwareUpgradeProposalParams{
				MaybeProposalId: nil,
				Content:         proposalContent,
				ProposerAddress: msg["proposer"].(string),
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
			ProposerAddress: msg["proposer"].(string),
			InitialDeposit:  initialDepositAmount,
		},
	)}, possibleSignerAddresses
}

func parseMsgSubmitTextProposal(
	txSuccess bool,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
	rawContent []byte,
) ([]command.Command, []string) {
	var proposalContent model.MsgSubmitTextProposalContent
	if err := jsoniter.Unmarshal(rawContent, &proposalContent); err != nil {
		panic("error decoding text proposal content")
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if msg != nil {
		if proposer, ok := msg["proposer"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, proposer.(string))
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
		return []command.Command{command_usecase.NewCreateMsgSubmitTextProposal(
			msgCommonParams,

			model.MsgSubmitTextProposalParams{
				MaybeProposalId: nil,
				Content:         proposalContent,
				ProposerAddress: msg["proposer"].(string),
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

	return []command.Command{command_usecase.NewCreateMsgSubmitTextProposal(
		msgCommonParams,

		model.MsgSubmitTextProposalParams{
			MaybeProposalId: proposalId,
			Content:         proposalContent,
			ProposerAddress: msg["proposer"].(string),
			InitialDeposit:  initialDepositAmount,
		},
	)}, possibleSignerAddresses
}

func parseMsgSubmitUnknownProposal(
	txSuccess bool,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
	rawContent []byte,
) ([]command.Command, []string) {
	var proposalContent model.MsgSubmitUnknownProposalContent
	if err := jsoniter.Unmarshal(rawContent, &proposalContent); err != nil {
		panic("error decoding unknown proposal content")
	}

	if err := jsoniter.Unmarshal(rawContent, &proposalContent.RawContent); err != nil {
		panic("error decoding unknown proposal rawContent")
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if msg != nil {
		if proposer, ok := msg["proposer"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, proposer.(string))
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
		return []command.Command{command_usecase.NewCreateMsgSubmitUnknownProposal(
			msgCommonParams,

			model.MsgSubmitUnknownProposalParams{
				MaybeProposalId: nil,
				Content:         proposalContent,
				ProposerAddress: msg["proposer"].(string),
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

	return []command.Command{command_usecase.NewCreateMsgSubmitUnknownProposal(
		msgCommonParams,

		model.MsgSubmitUnknownProposalParams{
			MaybeProposalId: proposalId,
			Content:         proposalContent,
			ProposerAddress: msg["proposer"].(string),
			InitialDeposit:  initialDepositAmount,
		},
	)}, possibleSignerAddresses
}

func ParseMsgVote(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if parserParams.Msg != nil {
		if voter, ok := parserParams.Msg["voter"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, voter.(string))
		}
	}

	return []command.Command{command_usecase.NewCreateMsgVote(
		parserParams.MsgCommonParams,

		model.MsgVoteParams{
			ProposalId: parserParams.Msg["proposal_id"].(string),
			Voter:      parserParams.Msg["voter"].(string),
			Option:     parserParams.Msg["option"].(string),
		},
	)}, possibleSignerAddresses
}

func ParseMsgDeposit(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if parserParams.Msg != nil {
		if depositor, ok := parserParams.Msg["depositor"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, depositor.(string))
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

	cmds := []command.Command{command_usecase.NewCreateMsgDeposit(
		parserParams.MsgCommonParams,

		model.MsgDepositParams{
			ProposalId: parserParams.Msg["proposal_id"].(string),
			Depositor:  parserParams.Msg["depositor"].(string),
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

func ParseMsgDelegate(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	amountValue, _ := parserParams.Msg["amount"].(map[string]interface{})
	amount, amountErr := tmcosmosutils.NewCoinFromAmountInterface(amountValue)
	if amountErr != nil {
		amount = coin.Coin{}
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if parserParams.Msg != nil {
		if delegatorAddress, ok := parserParams.Msg["delegator_address"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, delegatorAddress.(string))
		}
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		return []command.Command{command_usecase.NewCreateMsgDelegate(
			parserParams.MsgCommonParams,

			model.MsgDelegateParams{
				DelegatorAddress:   parserParams.Msg["delegator_address"].(string),
				ValidatorAddress:   parserParams.Msg["validator_address"].(string),
				Amount:             amount,
				AutoClaimedRewards: coin.Coin{},
			},
		)}, possibleSignerAddresses
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
	)}, possibleSignerAddresses
}

func ParseMsgUndelegate(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	amountValue, _ := parserParams.Msg["amount"].(map[string]interface{})
	amount, amountErr := tmcosmosutils.NewCoinFromAmountInterface(amountValue)
	if amountErr != nil {
		amount = coin.Coin{}
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if parserParams.Msg != nil {
		if delegatorAddress, ok := parserParams.Msg["delegator_address"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, delegatorAddress.(string))
		}
	}

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
		)}, possibleSignerAddresses
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
	)}, possibleSignerAddresses
}

func ParseMsgBeginRedelegate(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	amountValue, _ := parserParams.Msg["amount"].(map[string]interface{})
	amount, amountErr := tmcosmosutils.NewCoinFromAmountInterface(amountValue)
	if amountErr != nil {
		amount = coin.Coin{}
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if parserParams.Msg != nil {
		if delegatorAddress, ok := parserParams.Msg["delegator_address"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, delegatorAddress.(string))
		}
	}

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
		)}, possibleSignerAddresses
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
	)}, possibleSignerAddresses
}

func ParseMsgUnjail(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if parserParams.Msg != nil {
		if validatorAddr, ok := parserParams.Msg["validator_addr"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, validatorAddr.(string))
		}
	}

	return []command.Command{command_usecase.NewCreateMsgUnjail(
		parserParams.MsgCommonParams,

		model.MsgUnjailParams{
			ValidatorAddr: parserParams.Msg["validator_addr"].(string),
		},
	)}, possibleSignerAddresses
}

func parseGenesisGenTxsMsgCreateValidator(
	msg map[string]interface{},
) []command.Command {
	amountValue, _ := msg["value"].(map[string]interface{})
	amount, amountErr := tmcosmosutils.NewCoinFromAmountInterface(amountValue)
	if amountErr != nil {
		amount = coin.Coin{}
	}
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
) ([]command.Command, []string) {
	amountValue, _ := parserParams.Msg["value"].(map[string]interface{})
	amount, amountErr := tmcosmosutils.NewCoinFromAmountInterface(amountValue)
	if amountErr != nil {
		amount = coin.Coin{}
	}
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

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if parserParams.Msg != nil {
		if delegatorAddress, ok := parserParams.Msg["delegator_address"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, delegatorAddress.(string))
		}
		if validatorAddress, ok := parserParams.Msg["validator_address"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, validatorAddress.(string))
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
	)}, possibleSignerAddresses
}

func ParseMsgEditValidator(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
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

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if parserParams.Msg != nil {
		if validatorAddress, ok := parserParams.Msg["validator_address"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, validatorAddress.(string))
		}
	}

	return []command.Command{command_usecase.NewCreateMsgEditValidator(
		parserParams.MsgCommonParams,

		model.MsgEditValidatorParams{
			Description:            description,
			ValidatorAddress:       parserParams.Msg["validator_address"].(string),
			MaybeCommissionRate:    maybeCommissionRate,
			MaybeMinSelfDelegation: maybeMinSelfDelegation,
		},
	)}, possibleSignerAddresses
}

func ParseMsgNFTIssueDenom(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if parserParams.Msg != nil {
		if sender, ok := parserParams.Msg["sender"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, sender.(string))
		}
	}

	return []command.Command{command_usecase.NewCreateMsgNFTIssueDenom(
		parserParams.MsgCommonParams,

		model.MsgNFTIssueDenomParams{
			DenomId:   parserParams.Msg["id"].(string),
			DenomName: parserParams.Msg["name"].(string),
			Schema:    parserParams.Msg["schema"].(string),
			Sender:    parserParams.Msg["sender"].(string),
		},
	)}, possibleSignerAddresses
}

func ParseMsgNFTMintNFT(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if parserParams.Msg != nil {
		if sender, ok := parserParams.Msg["sender"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, sender.(string))
		}
	}

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
	)}, possibleSignerAddresses
}

func ParseMsgNFTTransferNFT(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if parserParams.Msg != nil {
		if sender, ok := parserParams.Msg["sender"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, sender.(string))
		}
	}

	return []command.Command{command_usecase.NewCreateMsgNFTTransferNFT(
		parserParams.MsgCommonParams,

		model.MsgNFTTransferNFTParams{
			TokenId:   parserParams.Msg["id"].(string),
			DenomId:   parserParams.Msg["denom_id"].(string),
			Sender:    parserParams.Msg["sender"].(string),
			Recipient: parserParams.Msg["recipient"].(string),
		},
	)}, possibleSignerAddresses
}

func ParseMsgNFTEditNFT(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if parserParams.Msg != nil {
		if sender, ok := parserParams.Msg["sender"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, sender.(string))
		}
	}

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
	)}, possibleSignerAddresses
}

func ParseMsgNFTBurnNFT(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if parserParams.Msg != nil {
		if sender, ok := parserParams.Msg["sender"]; ok {
			possibleSignerAddresses = append(possibleSignerAddresses, sender.(string))
		}
	}

	return []command.Command{command_usecase.NewCreateMsgNFTBurnNFT(
		parserParams.MsgCommonParams,

		model.MsgNFTBurnNFTParams{
			DenomId: parserParams.Msg["denom_id"].(string),
			TokenId: parserParams.Msg["id"].(string),
			Sender:  parserParams.Msg["sender"].(string),
		},
	)}, possibleSignerAddresses
}

func ParseMsgGrant(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	authType := parserParams.Msg["grant"].(map[string]interface{})["authorization"].(map[string]interface{})["@type"]

	switch authType {
	case "/cosmos.bank.v1beta1.SendAuthorization":
		return parseRawMsgSendGrant(parserParams.MsgCommonParams, parserParams.Msg)
	case "/cosmos.staking.v1beta1.StakeAuthorization":
		return parseRawMsgStackGrant(parserParams.MsgCommonParams, parserParams.Msg)
	case "/cosmos.authz.v1beta1.GenericAuthorization":
		return parseRawMsgGenericGrant(parserParams.MsgCommonParams, parserParams.Msg)
	default:
		return []command.Command{}, []string{}
	}
}

func parseRawMsgSendGrant(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) ([]command.Command, []string) {
	var rawMsg model.RawMsgSendGrant
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
		panic(fmt.Errorf("error creating RawMsgSendGrant decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgSendGrant: %v", err))
	}

	if !msgCommonParams.TxSuccess {
		params := model.MsgGrantParams{
			MaybeSendGrant: &rawMsg,
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string

		if params.MaybeSendGrant != nil {
			if params.MaybeSendGrant.Granter != "" {
				possibleSignerAddresses = append(possibleSignerAddresses, params.MaybeSendGrant.Granter)
			}
		}

		return []command.Command{command_usecase.NewCreateMsgGrant(
			msgCommonParams,

			params,
		)}, possibleSignerAddresses
	}

	params := model.MsgGrantParams{
		MaybeSendGrant: &rawMsg,
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if params.MaybeSendGrant != nil {
		if params.MaybeSendGrant.Granter != "" {
			possibleSignerAddresses = append(possibleSignerAddresses, params.MaybeSendGrant.Granter)
		}
	}

	return []command.Command{command_usecase.NewCreateMsgGrant(
		msgCommonParams,

		params,
	)}, possibleSignerAddresses
}

func parseRawMsgStackGrant(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) ([]command.Command, []string) {
	var rawMsg model.RawMsgStakeGrant
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
		panic(fmt.Errorf("error creating RawMsgStakeGrant decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgStakeGrant: %v", err))
	}

	if !msgCommonParams.TxSuccess {
		params := model.MsgGrantParams{
			MaybeStakeGrant: &rawMsg,
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		if params.MaybeStakeGrant != nil {
			if params.MaybeStakeGrant.Granter != "" {
				possibleSignerAddresses = append(possibleSignerAddresses, params.MaybeStakeGrant.Granter)
			}
		}

		return []command.Command{command_usecase.NewCreateMsgGrant(
			msgCommonParams,

			params,
		)}, possibleSignerAddresses
	}

	params := model.MsgGrantParams{
		MaybeStakeGrant: &rawMsg,
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if params.MaybeStakeGrant != nil {
		if params.MaybeStakeGrant.Granter != "" {
			possibleSignerAddresses = append(possibleSignerAddresses, params.MaybeStakeGrant.Granter)
		}
	}

	return []command.Command{command_usecase.NewCreateMsgGrant(
		msgCommonParams,

		params,
	)}, possibleSignerAddresses
}

func parseRawMsgGenericGrant(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) ([]command.Command, []string) {
	var rawMsg model.RawMsgGenericGrant
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
		panic(fmt.Errorf("error creating RawMsgGenericGrant decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgGenericGrant: %v", err))
	}

	if !msgCommonParams.TxSuccess {
		params := model.MsgGrantParams{
			MaybeGenericGrant: &rawMsg,
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		if params.MaybeGenericGrant != nil {
			if params.MaybeGenericGrant.Granter != "" {
				possibleSignerAddresses = append(possibleSignerAddresses, params.MaybeGenericGrant.Granter)
			}
		}

		return []command.Command{command_usecase.NewCreateMsgGrant(
			msgCommonParams,

			params,
		)}, possibleSignerAddresses
	}

	params := model.MsgGrantParams{
		MaybeGenericGrant: &rawMsg,
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if params.MaybeGenericGrant != nil {
		if params.MaybeGenericGrant.Granter != "" {
			possibleSignerAddresses = append(possibleSignerAddresses, params.MaybeGenericGrant.Granter)
		}
	}

	return []command.Command{command_usecase.NewCreateMsgGrant(
		msgCommonParams,

		params,
	)}, possibleSignerAddresses
}

func ParseMsgRevoke(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var rawMsg model.RawMsgRevoke
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
		panic(fmt.Errorf("error creating RawMsgRevoke decoder: %v", decoderErr))
	}
	if err := decoder.Decode(parserParams.Msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgRevoke: %v", err))
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		revokeParams := model.MsgRevokeParams{
			RawMsgRevoke: rawMsg,
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		if revokeParams.RawMsgRevoke.Granter != "" {
			possibleSignerAddresses = append(possibleSignerAddresses, revokeParams.RawMsgRevoke.Granter)
		}

		return []command.Command{command_usecase.NewCreateMsgRevoke(
			parserParams.MsgCommonParams,

			revokeParams,
		)}, possibleSignerAddresses
	}

	revokeParams := model.MsgRevokeParams{
		RawMsgRevoke: rawMsg,
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if revokeParams.RawMsgRevoke.Granter != "" {
		possibleSignerAddresses = append(possibleSignerAddresses, revokeParams.RawMsgRevoke.Granter)
	}

	return []command.Command{command_usecase.NewCreateMsgRevoke(
		parserParams.MsgCommonParams,

		revokeParams,
	)}, possibleSignerAddresses
}

func ParseMsgExec(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var rawMsg model.RawMsgExec
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
		panic(fmt.Errorf("error creating ParseMsgExec decoder: %v", decoderErr))
	}
	if err := decoder.Decode(parserParams.Msg); err != nil {
		panic(fmt.Errorf("error decoding ParseMsgExec: %v", err))
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		execParams := model.MsgExecParams{
			RawMsgExec: rawMsg,
		}

		innerCommands := parseMsgExecInnerMsgs(parserParams)

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		if execParams.RawMsgExec.Grantee != "" {
			possibleSignerAddresses = append(possibleSignerAddresses, execParams.RawMsgExec.Grantee)
		}

		return append([]command.Command{command_usecase.NewCreateMsgExec(
			parserParams.MsgCommonParams,

			execParams,
		)}, innerCommands...), possibleSignerAddresses
	}

	execParams := model.MsgExecParams{
		RawMsgExec: rawMsg,
	}

	innerCommands := parseMsgExecInnerMsgs(parserParams)

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if execParams.RawMsgExec.Grantee != "" {
		possibleSignerAddresses = append(possibleSignerAddresses, execParams.RawMsgExec.Grantee)
	}

	return append([]command.Command{command_usecase.NewCreateMsgExec(
		parserParams.MsgCommonParams,

		execParams,
	)}, innerCommands...), possibleSignerAddresses
}

func parseMsgExecInnerMsgs(
	parserParams utils.CosmosParserParams,
) []command.Command {

	var commands []command.Command

	blockHeight := parserParams.MsgCommonParams.BlockHeight

	msgs, ok := parserParams.Msg["msgs"].([]interface{})
	if !ok {
		panic(fmt.Errorf("error parsing MsgExec.msgs to []interface{}: %v", parserParams.Msg["msgs"]))
	}

	for innerMsgIndex, innerMsgInterface := range msgs {
		innerMsg, ok := innerMsgInterface.(map[string]interface{})
		if !ok {
			panic(fmt.Errorf("error parsing MsgExec.msgs[%v] to map[string]interface{}: %v", innerMsgIndex, innerMsgInterface))
		}

		innerMsgType, ok := innerMsg["@type"].(string)
		if !ok {
			panic(fmt.Errorf("error missing '@type' in MsgExec.msgs[%v]: %v", innerMsgIndex, innerMsg))
		}

		parser := parserParams.ParserManager.GetParser(utils.CosmosParserKey(innerMsgType), utils.ParserBlockHeight(blockHeight))

		msgCommands, _ := parser(utils.CosmosParserParams{
			AddressPrefix:   parserParams.AddressPrefix,
			StakingDenom:    parserParams.StakingDenom,
			TxsResult:       parserParams.TxsResult,
			MsgCommonParams: parserParams.MsgCommonParams,
			Msg:             innerMsg,
			MsgIndex:        parserParams.MsgIndex,
			ParserManager:   parserParams.ParserManager,
		})
		commands = append(commands, msgCommands...)
	}

	return commands
}

func ParseMsgGrantAllowance(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	allowanceType := parserParams.Msg["allowance"].(map[string]interface{})["@type"]

	switch allowanceType {
	case "/cosmos.feegrant.v1beta1.BasicAllowance":
		return parseRawMsgGrantBasicAllowance(parserParams.MsgCommonParams, parserParams.Msg)
	case "/cosmos.feegrant.v1beta1.PeriodicAllowance":
		return parseRawMsgGrantPeriodicAllowance(parserParams.MsgCommonParams, parserParams.Msg)
	case "/cosmos.feegrant.v1beta1.AllowedMsgAllowance":
		return parseRawMsgGrantAllowedMsgAllowance(parserParams.MsgCommonParams, parserParams.Msg)
	default:
		return []command.Command{}, []string{}
	}
}

func parseRawMsgGrantBasicAllowance(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) ([]command.Command, []string) {
	var rawMsg model.RawMsgGrantBasicAllowance
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
		panic(fmt.Errorf("error creating RawMsgGrantBasicAllowance decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgGrantBasicAllowance: %v", err))
	}

	if !msgCommonParams.TxSuccess {
		params := model.MsgGrantAllowanceParams{
			MaybeBasicAllowance: &rawMsg,
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		if params.MaybeBasicAllowance != nil {
			if params.MaybeBasicAllowance.Granter != "" {
				possibleSignerAddresses = append(possibleSignerAddresses, params.MaybeBasicAllowance.Granter)
			}
		}

		return []command.Command{command_usecase.NewCreateMsgGrantAllowance(
			msgCommonParams,

			params,
		)}, possibleSignerAddresses
	}

	params := model.MsgGrantAllowanceParams{
		MaybeBasicAllowance: &rawMsg,
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if params.MaybeBasicAllowance != nil {
		if params.MaybeBasicAllowance.Granter != "" {
			possibleSignerAddresses = append(possibleSignerAddresses, params.MaybeBasicAllowance.Granter)
		}
	}

	return []command.Command{command_usecase.NewCreateMsgGrantAllowance(
		msgCommonParams,

		params,
	)}, possibleSignerAddresses
}

func parseRawMsgGrantPeriodicAllowance(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) ([]command.Command, []string) {
	var rawMsg model.RawMsgGrantPeriodicAllowance
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
		panic(fmt.Errorf("error creating RawMsgGrantPeriodicAllowance decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgGrantPeriodicAllowance: %v", err))
	}

	if !msgCommonParams.TxSuccess {
		params := model.MsgGrantAllowanceParams{
			MaybePeriodicAllowance: &rawMsg,
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		if params.MaybePeriodicAllowance != nil {
			if params.MaybePeriodicAllowance.Granter != "" {
				possibleSignerAddresses = append(possibleSignerAddresses, params.MaybePeriodicAllowance.Granter)
			}
		}

		return []command.Command{command_usecase.NewCreateMsgGrantAllowance(
			msgCommonParams,

			params,
		)}, possibleSignerAddresses
	}

	params := model.MsgGrantAllowanceParams{
		MaybePeriodicAllowance: &rawMsg,
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if params.MaybePeriodicAllowance != nil {
		if params.MaybePeriodicAllowance.Granter != "" {
			possibleSignerAddresses = append(possibleSignerAddresses, params.MaybePeriodicAllowance.Granter)
		}
	}

	return []command.Command{command_usecase.NewCreateMsgGrantAllowance(
		msgCommonParams,

		params,
	)}, possibleSignerAddresses
}

func parseRawMsgGrantAllowedMsgAllowance(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) ([]command.Command, []string) {
	var rawMsg model.RawMsgGrantAllowedMsgAllowance
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
		panic(fmt.Errorf("error creating RawMsgGrantAllowedMsgAllowance decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgGrantAllowedMsgAllowance: %v", err))
	}

	if !msgCommonParams.TxSuccess {
		params := model.MsgGrantAllowanceParams{
			MaybeAllowedMsgAllowance: &rawMsg,
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		if params.MaybeAllowedMsgAllowance != nil {
			if params.MaybeAllowedMsgAllowance.Granter != "" {
				possibleSignerAddresses = append(possibleSignerAddresses, params.MaybeAllowedMsgAllowance.Granter)
			}
		}

		return []command.Command{command_usecase.NewCreateMsgGrantAllowance(
			msgCommonParams,

			params,
		)}, possibleSignerAddresses
	}

	params := model.MsgGrantAllowanceParams{
		MaybeAllowedMsgAllowance: &rawMsg,
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if params.MaybeAllowedMsgAllowance != nil {
		if params.MaybeAllowedMsgAllowance.Granter != "" {
			possibleSignerAddresses = append(possibleSignerAddresses, params.MaybeAllowedMsgAllowance.Granter)
		}
	}

	return []command.Command{command_usecase.NewCreateMsgGrantAllowance(
		msgCommonParams,

		params,
	)}, possibleSignerAddresses
}

func ParseMsgRevokeAllowance(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var rawMsg model.RawMsgRevokeAllowance
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
		panic(fmt.Errorf("error creating RawMsgRevokeAllowance decoder: %v", decoderErr))
	}
	if err := decoder.Decode(parserParams.Msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgRevokeAllowance: %v", err))
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		revokeAllowanceParams := model.MsgRevokeAllowanceParams{
			RawMsgRevokeAllowance: rawMsg,
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		if revokeAllowanceParams.RawMsgRevokeAllowance.Granter != "" {
			possibleSignerAddresses = append(possibleSignerAddresses, revokeAllowanceParams.RawMsgRevokeAllowance.Granter)
		}

		return []command.Command{command_usecase.NewCreateMsgRevokeAllowance(
			parserParams.MsgCommonParams,

			revokeAllowanceParams,
		)}, possibleSignerAddresses
	}

	revokeAllowanceParams := model.MsgRevokeAllowanceParams{
		RawMsgRevokeAllowance: rawMsg,
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if revokeAllowanceParams.RawMsgRevokeAllowance.Granter != "" {
		possibleSignerAddresses = append(possibleSignerAddresses, revokeAllowanceParams.RawMsgRevokeAllowance.Granter)
	}

	return []command.Command{command_usecase.NewCreateMsgRevokeAllowance(
		parserParams.MsgCommonParams,

		revokeAllowanceParams,
	)}, possibleSignerAddresses
}

func ParseMsgCreateVestingAccount(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var rawMsg model.RawMsgCreateVestingAccount
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
		panic(fmt.Errorf("error creating RawMsgCreateVestingAccount decoder: %v", decoderErr))
	}
	if err := decoder.Decode(parserParams.Msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgCreateVestingAccount: %v", err))
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		msgCreateVestingAccountParams := model.MsgCreateVestingAccountParams{
			RawMsgCreateVestingAccount: rawMsg,
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		if msgCreateVestingAccountParams.RawMsgCreateVestingAccount.FromAddress != "" {
			possibleSignerAddresses = append(possibleSignerAddresses, msgCreateVestingAccountParams.RawMsgCreateVestingAccount.FromAddress)
		}

		return []command.Command{command_usecase.NewCreateMsgCreateVestingAccount(
			parserParams.MsgCommonParams,

			msgCreateVestingAccountParams,
		)}, possibleSignerAddresses
	}

	msgCreateVestingAccountParams := model.MsgCreateVestingAccountParams{
		RawMsgCreateVestingAccount: rawMsg,
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	if msgCreateVestingAccountParams.RawMsgCreateVestingAccount.FromAddress != "" {
		possibleSignerAddresses = append(possibleSignerAddresses, msgCreateVestingAccountParams.RawMsgCreateVestingAccount.FromAddress)
	}

	return []command.Command{command_usecase.NewCreateMsgCreateVestingAccount(
		parserParams.MsgCommonParams,

		msgCreateVestingAccountParams,
	)}, possibleSignerAddresses
}

func ParseMsgEthereumTx(
	parserParams utils.CosmosParserParams,
) ([]command.Command, []string) {
	var rawMsg model.RawMsgEthereumTx
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
		panic(fmt.Errorf("error creating RawMsgEthereumTx decoder: %v", decoderErr))
	}
	if err := decoder.Decode(parserParams.Msg); err != nil {
		panic(fmt.Errorf("error decoding RawMsgEthereumTx: %v", err))
	}

	if !parserParams.MsgCommonParams.TxSuccess {
		// FIXME: https://github.com/crypto-com/chain-indexing/issues/730
		msgEthereumTxParams := model.MsgEthereumTxParams{
			RawMsgEthereumTx: rawMsg,
		}

		// Getting possible signer address from Msg
		var possibleSignerAddresses []string
		// FIXME: https://github.com/crypto-com/chain-indexing/issues/729
		// possibleSignerAddresses = append(possibleSignerAddresses, msgEthereumTxParams.From)

		return []command.Command{command_usecase.NewCreateMsgEthereumTx(
			parserParams.MsgCommonParams,

			msgEthereumTxParams,
		)}, possibleSignerAddresses
	}

	// FIXME: https://github.com/crypto-com/chain-indexing/issues/730
	msgEthereumTxParams := model.MsgEthereumTxParams{
		RawMsgEthereumTx: rawMsg,
	}

	// Getting possible signer address from Msg
	var possibleSignerAddresses []string
	// FIXME: https://github.com/crypto-com/chain-indexing/issues/729
	// possibleSignerAddresses = append(possibleSignerAddresses, msgEthereumTxParams.From)

	return []command.Command{command_usecase.NewCreateMsgEthereumTx(
		parserParams.MsgCommonParams,

		msgEthereumTxParams,
	)}, possibleSignerAddresses
}
