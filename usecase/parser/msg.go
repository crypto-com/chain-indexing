package parser

import (
	"fmt"

	"github.com/crypto-com/chainindex/entity/command"
	"github.com/crypto-com/chainindex/usecase/coin"
	command_usecase "github.com/crypto-com/chainindex/usecase/command"
	"github.com/crypto-com/chainindex/usecase/event"
	"github.com/crypto-com/chainindex/usecase/model"
)

func ParseMsgToCommands(
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

			var msgCommand command.Command
			switch msg["@type"] {
			case "/cosmos.bank.v1beta1.MsgSend":
				msgCommand = parseMsgSend(msgCommonParams, msg)
			case "/cosmos.bank.v1beta1.MsgMultiSend":
				msgCommand = parseMsgMultiSend(msgCommonParams, msg)
			case "/cosmos.distribution.v1beta1.MsgSetWithdrawAddress":
				msgCommand = parseMsgSetWithdrawAddress(msgCommonParams, msg)
			case "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward":
				msgCommand = parseMsgWithdrawDelegatorReward(txSuccess, txsResult, msgIndex, msgCommonParams, msg)
			case "/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission":
				msgCommand = parseMsgWithdrawValidatorCommission(txSuccess, txsResult, msgIndex, msgCommonParams, msg)
			case "/cosmos.distribution.v1beta1.MsgFundCommunityPool":
				msgCommand = parseMsgFundCommunityPool(msgCommonParams, msg)
			case "/cosmos.staking.v1beta1.MsgDelegate":
				msgCommand = parseMsgDelegate(msgCommonParams, msg)
			case "/cosmos.staking.v1beta1.MsgUndelegate":
				msgCommand = parseMsgUndelegate(msgCommonParams, msg)
			case "/cosmos.staking.v1beta1.MsgBeginRedelegate":
				msgCommand = parseMsgBeginRedelegate(msgCommonParams, msg)
			case "/cosmos.slashing.v1beta1.MsgUnjail":
				msgCommand = parseMsgUnjail(msgCommonParams, msg)
			}

			commands = append(commands, msgCommand)
		}
	}

	return commands, nil
}

func parseMsgSend(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) *command_usecase.CreateMsgSend {
	return command_usecase.NewCreateMsgSend(
		msgCommonParams,

		event.MsgSendCreatedParams{
			FromAddress: msg["from_address"].(string),
			ToAddress:   msg["to_address"].(string),
			Amount:      sumAmountInterfaces(msg["amount"].([]interface{})),
		},
	)
}

func parseMsgMultiSend(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) *command_usecase.CreateMsgMultiSend {
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

	return command_usecase.NewCreateMsgMultiSend(
		msgCommonParams,

		model.MsgMultiSendParams{
			Inputs:  inputs,
			Outputs: outputs,
		},
	)
}

func parseMsgSetWithdrawAddress(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) *command_usecase.CreateMsgSetWithdrawAddress {
	return command_usecase.NewCreateMsgSetWithdrawAddress(
		msgCommonParams,

		model.MsgSetWithdrawAddressParams{
			DelegatorAddress: msg["delegator_address"].(string),
			WithdrawAddress:  msg["withdraw_address"].(string),
		},
	)
}

func parseMsgWithdrawDelegatorReward(
	txSuccess bool,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) *command_usecase.CreateMsgWithdrawDelegatorReward {
	if !txSuccess {
		delegatorAddress, _ := msg["delegator_address"].(string)
		return command_usecase.NewCreateMsgWithdrawDelegatorReward(
			msgCommonParams,

			model.MsgWithdrawDelegatorRewardParams{
				DelegatorAddress: delegatorAddress,
				ValidatorAddress: msg["validator_address"].(string),
				RecipientAddress: delegatorAddress,
				Amount:           coin.Zero(),
			},
		)
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

	return command_usecase.NewCreateMsgWithdrawDelegatorReward(
		msgCommonParams,

		model.MsgWithdrawDelegatorRewardParams{
			DelegatorAddress: msg["delegator_address"].(string),
			ValidatorAddress: msg["validator_address"].(string),
			RecipientAddress: recipient,
			Amount:           amount,
		},
	)
}

func parseMsgWithdrawValidatorCommission(
	txSuccess bool,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) *command_usecase.CreateMsgWithdrawValidatorCommission {
	if !txSuccess {
		return command_usecase.NewCreateMsgWithdrawValidatorCommission(
			msgCommonParams,

			model.MsgWithdrawValidatorCommissionParams{
				ValidatorAddress: msg["validator_address"].(string),
				RecipientAddress: msg["delegator_address"].(string),
				Amount:           coin.Zero(),
			},
		)
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

	return command_usecase.NewCreateMsgWithdrawValidatorCommission(
		msgCommonParams,

		model.MsgWithdrawValidatorCommissionParams{
			ValidatorAddress: msg["validator_address"].(string),
			RecipientAddress: recipient,
			Amount:           amount,
		},
	)
}

func parseMsgFundCommunityPool(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) *command_usecase.CreateMsgFundCommunityPool {
	return command_usecase.NewCreateMsgFundCommunityPool(
		msgCommonParams,

		model.MsgFundCommunityPoolParams{
			Depositor: msg["depositor"].(string),
			Amount:    sumAmountInterfaces(msg["amount"].([]interface{})),
		},
	)
}

func parseMsgDelegate(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) *command_usecase.CreateMsgDelegate {
	amountValue, _ := msg["amount"].(map[string]interface{})
	amount := coin.MustNewCoinFromString(amountValue["amount"].(string))

	return command_usecase.NewCreateMsgDelegate(
		msgCommonParams,

		model.MsgDelegateParams{
			DelegatorAddress: msg["delegator_address"].(string),
			ValidatorAddress: msg["validator_address"].(string),
			Amount:           amount,
		},
	)
}

func parseMsgUndelegate(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) *command_usecase.CreateMsgUndelegate {
	amountValue, _ := msg["amount"].(map[string]interface{})
	amount := coin.MustNewCoinFromString(amountValue["amount"].(string))

	return command_usecase.NewCreateMsgUndelegate(
		msgCommonParams,

		model.MsgUndelegateParams{
			DelegatorAddress: msg["delegator_address"].(string),
			ValidatorAddress: msg["validator_address"].(string),
			Amount:           amount,
		},
	)
}

func parseMsgBeginRedelegate(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) *command_usecase.CreateMsgBeginRedelegate {
	amountValue, _ := msg["amount"].(map[string]interface{})
	amount := coin.MustNewCoinFromString(amountValue["amount"].(string))

	return command_usecase.NewCreateMsgBeginRedelegate(
		msgCommonParams,

		model.MsgBeginRedelegateParams{
			DelegatorAddress:    msg["delegator_address"].(string),
			ValidatorSrcAddress: msg["validator_src_address"].(string),
			ValidatorDstAddress: msg["validator_dst_address"].(string),
			Amount:              amount,
		},
	)
}

func parseMsgUnjail(
	msgCommonParams event.MsgCommonParams,
	msg map[string]interface{},
) *command_usecase.CreateMsgUnjail {
	return command_usecase.NewCreateMsgUnjail(
		msgCommonParams,

		model.MsgUnjailParams{
			ValidatorAddr: msg["validator_addr"].(string),
		},
	)
}

func sumAmountInterfaces(amounts []interface{}) coin.Coin {
	sum := coin.Zero()
	for _, rawAmount := range amounts {
		amount, _ := rawAmount.(map[string]interface{})
		sum, _ = sum.Add(coin.MustNewCoinFromString(amount["amount"].(string)))
	}

	return sum
}
