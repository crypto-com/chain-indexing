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
) []command.Command {
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
			if msg["@type"] == "/cosmos.bank.v1beta1.MsgSend" {
				commands = append(commands, command_usecase.NewCreateMsgSend(
					msgCommonParams,

					event.MsgSendCreatedParams{
						FromAddress: msg["from_address"].(string),
						ToAddress:   msg["to_address"].(string),
						Amount:      sumAmountInterfaces(msg["amount"].([]interface{})),
					},
				))
			} else if msg["@type"] == "/cosmos.bank.v1beta1.MsgMultiSend" {
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

				commands = append(commands, command_usecase.NewCreateMsgMultiSend(
					msgCommonParams,

					model.MsgMultiSendParams{
						Inputs:  inputs,
						Outputs: outputs,
					},
				))
			} else if msg["@type"] == "/cosmos.distribution.v1beta1.MsgSetWithdrawAddress" {
				commands = append(commands, command_usecase.NewCreateMsgSetWithdrawAddress(
					msgCommonParams,

					model.MsgSetWithdrawAddressParams{
						DelegatorAddress: msg["delegator_address"].(string),
						WithdrawAddress:  msg["withdraw_address"].(string),
					},
				))
			} else if msg["@type"] == "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward" {
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

				commands = append(commands, command_usecase.NewCreateMsgWithdrawDelegatorReward(
					msgCommonParams,

					model.MsgWithdrawDelegatorRewardParams{
						DelegatorAddress: msg["delegator_address"].(string),
						ValidatorAddress: msg["validator_address"].(string),
						RecipientAddress: recipient,
						Amount:           amount,
					},
				))
			} else if msg["@type"] == "/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission" {
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

				commands = append(commands, command_usecase.NewCreateMsgWithdrawValidatorCommission(
					msgCommonParams,

					model.MsgWithdrawValidatorCommissionParams{
						ValidatorAddress: msg["validator_address"].(string),
						RecipientAddress: recipient,
						Amount:           amount,
					},
				))
			} else if msg["@type"] == "/cosmos.distribution.v1beta1.MsgFundCommunityPool" {
				commands = append(commands, command_usecase.NewCreateMsgFundCommunityPool(
					msgCommonParams,

					model.MsgFundCommunityPoolParams{
						Depositor: msg["depositor"].(string),
						Amount:    sumAmountInterfaces(msg["amount"].([]interface{})),
					},
				))
			} else if msg["@type"] == "/cosmos.staking.v1beta1.MsgDelegate" {
				amountValue, _ := msg["amount"].(map[string]interface{})
				amount := coin.MustNewCoinFromString(amountValue["amount"].(string))

				commands = append(commands, command_usecase.NewCreateMsgDelegate(
					msgCommonParams,

					model.MsgDelegateParams{
						DelegatorAddress: msg["delegator_address"].(string),
						ValidatorAddress: msg["validator_address"].(string),
						Amount:           amount,
					},
				))
			} else if msg["@type"] == "/cosmos.staking.v1beta1.MsgUndelegate" {
				amountValue, _ := msg["amount"].(map[string]interface{})
				amount := coin.MustNewCoinFromString(amountValue["amount"].(string))

				commands = append(commands, command_usecase.NewCreateMsgUndelegate(
					msgCommonParams,

					model.MsgUndelegateParams{
						DelegatorAddress: msg["delegator_address"].(string),
						ValidatorAddress: msg["validator_address"].(string),
						Amount:           amount,
					},
				))
			}
		}
	}

	//for i, txsResult := range blockResults.TxsResults {
	//	for _, log := range txsResult.Log {
	//		txsResultLog := NewParsedTxsResultLog(&log)
	//		if !txsResultLog.HasEvent("message") {
	//			continue
	//		}
	//
	//		msgType := getMsgType(txsResultLog)
	//		if msgType.Module == "bank" && msgType.Action == "send" {
	//			transferEvent := txsResultLog.GetEventByType("transfer")
	//			if transferEvent == nil {
	//				panic("expected bank.send message to have transfer event, but found none")
	//			}
	//
	//			amount := transferEvent.MustGetAttributeByKey("amount")
	//			commands = append(commands, command_usecase.NewCreateMsgSend(
	//				blockHeight,
	//
	//				txHash,
	//				log.MsgIndex,
	//				event.MsgSendCreatedParams{
	//					FromAddress: transferEvent.MustGetAttributeByKey("sender"),
	//					ToAddress:   transferEvent.MustGetAttributeByKey("recipient"),
	//					Amount:      coin.MustNewCoinFromString(TrimAmountDenom(amount)),
	//				},
	//			))
	//		}
	//	}
	//}

	return commands
}

func sumAmountInterfaces(amounts []interface{}) coin.Coin {
	for _, rawAmount := range amounts {
		amount, _ := rawAmount.(map[string]interface{})
		return coin.MustNewCoinFromString(amount["amount"].(string))
	}

	return coin.Zero()
}

// Returns message type from block results TxsResult rawLog. Returns nil if no message is found.
//func getMsgType(txsResultLog *ParsedTxsResultLog) *msgType {
//	messageEvent := txsResultLog.GetEventByType("message")
//	if messageEvent == nil {
//		return nil
//	}
//	return &msgType{
//		Module: messageEvent.MustGetAttributeByKey("module"),
//		Action: messageEvent.MustGetAttributeByKey("action"),
//	}
//}
//
//type msgType struct {
//	Module string
//	Action string
//}
