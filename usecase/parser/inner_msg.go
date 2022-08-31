package parser

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
)

func ParseTxsResultsEvents(
	innerMsgIndex int,
	innerMsg map[string]interface{},
	logs model.BlockResultsTxsResultLog,
) []model.BlockResultsTxsResultLog {
	var resultLog []model.BlockResultsTxsResultLog

	// parse events with index
	parsedEvents := utils.NewParsedTxsResultsEvents(logs.Events)
	extractedEvents := ParseInnerMsgsEvents(innerMsg, innerMsgIndex, parsedEvents)

	log := model.BlockResultsTxsResultLog{
		MsgIndex: innerMsgIndex,
		Events:   extractedEvents,
	}
	resultLog = append(resultLog, log)

	return resultLog
}

func ParseInnerMsgsEvents(
	innerMsg map[string]interface{},
	innerMsgIndex int,
	parsedEvents *utils.ParsedTxsResultsEvents,
) []model.BlockResultsEvent {
	var extractedEvents []model.BlockResultsEvent

	innerMsgType, ok := innerMsg["@type"].(string)
	if !ok {
		panic(fmt.Errorf("error missing '@type' in MsgExec.msgs[%v]: %v", innerMsgIndex, innerMsg))
	}

	switch innerMsgType {
	// bank
	case "/cosmos.bank.v1beta1.MsgSend":
		extractedEvents = MsgSend(innerMsgIndex, innerMsg, parsedEvents)

	// distribution
	case "/cosmos.distribution.v1beta1.MsgSetWithdrawAddress":
		extractedEvents = MsgSetWithdrawAddress(innerMsgIndex, innerMsg, parsedEvents)
	case "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward":
		extractedEvents = MsgWithdrawDelegatorReward(innerMsgIndex, innerMsg, parsedEvents)
	case "/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission":
		extractedEvents = MsgWithdrawValidatorCommission(innerMsgIndex, innerMsg, parsedEvents)
	case "/cosmos.distribution.v1beta1.MsgFundCommunityPool":
		extractedEvents = MsgFundCommunityPool(innerMsgIndex, innerMsg, parsedEvents)

	// cosmos staking
	case "/cosmos.staking.v1beta1.MsgDelegate":
		extractedEvents = MsgDelegate(innerMsgIndex, innerMsg, parsedEvents)
		// case "/cosmos.staking.v1beta1.MsgUndelegate":
		// 	extractedEvents = MsgUndelegate(innerMsgIndex, innerMsg, parsedEvents)
		// case "/cosmos.staking.v1beta1.MsgBeginRedelegate":
		// 	extractedEvents = MsgBeginRedelegate(innerMsgIndex, innerMsg, parsedEvents)
	}

	return extractedEvents
}

func MsgSend(
	innerMsgIndex int, innerMsg map[string]interface{}, parsedEvents *utils.ParsedTxsResultsEvents,
) []model.BlockResultsEvent {
	eventTypes := []string{"coin_spent", "coin_received", "transfer", "message"}

	return extract(innerMsgIndex, innerMsg, eventTypes, []string{}, parsedEvents)
}

func MsgSetWithdrawAddress(innerMsgIndex int, innerMsg map[string]interface{}, parsedEvents *utils.ParsedTxsResultsEvents,
) []model.BlockResultsEvent {
	eventTypes := []string{"set_withdraw_address", "message"}

	return extract(innerMsgIndex, innerMsg, eventTypes, []string{}, parsedEvents)
}

func MsgWithdrawDelegatorReward(innerMsgIndex int, innerMsg map[string]interface{}, parsedEvents *utils.ParsedTxsResultsEvents,
) []model.BlockResultsEvent {
	eventTypes := []string{"coin_spent", "coin_received", "transfer", "message", "withdraw_rewards"}

	eventTypesNoTransfer := []string{"message", "withdraw_rewards"}
	return extract(innerMsgIndex, innerMsg, eventTypes, eventTypesNoTransfer, parsedEvents)
}

func MsgWithdrawValidatorCommission(innerMsgIndex int, innerMsg map[string]interface{}, parsedEvents *utils.ParsedTxsResultsEvents,
) []model.BlockResultsEvent {
	eventTypes := []string{"coin_spent", "coin_received", "transfer", "message", "withdraw_commission"}

	eventTypesNoTransfer := []string{"message", "withdraw_commission"}

	return extract(innerMsgIndex, innerMsg, eventTypes, eventTypesNoTransfer, parsedEvents)
}

func MsgFundCommunityPool(innerMsgIndex int, innerMsg map[string]interface{}, parsedEvents *utils.ParsedTxsResultsEvents,
) []model.BlockResultsEvent {
	eventTypes := []string{"coin_spent", "coin_received", "transfer", "message"}

	eventTypesNoTransfer := []string{"message"}

	return extract(innerMsgIndex, innerMsg, eventTypes, eventTypesNoTransfer, parsedEvents)
}

func MsgDelegate(innerMsgIndex int, innerMsg map[string]interface{}, parsedEvents *utils.ParsedTxsResultsEvents,
) []model.BlockResultsEvent {
	eventTypes := []string{"coin_spent", "coin_received", "transfer", "delegate", "message"}

	eventTypesNoTransfer := []string{"coin_spent", "coin_received", "delegate", "message"}

	return extract(innerMsgIndex, innerMsg, eventTypes, eventTypesNoTransfer, parsedEvents)
}

func extract(innerMsgIndex int, innerMsg map[string]interface{}, eventTypes []string, eventTypesRetry []string, parsedEvents *utils.ParsedTxsResultsEvents) []model.BlockResultsEvent {
	// extract events
	extractedEvents := getMsgEvents(innerMsgIndex, innerMsg, eventTypes, parsedEvents)

	// retry
	if len(extractedEvents) <= 0 {
		extractedEvents = getMsgEvents(innerMsgIndex, innerMsg, eventTypesRetry, parsedEvents)
	}

	return extractedEvents
}

func getMsgEvents(innerMsgIndex int, innerMsg map[string]interface{}, eventTypes []string, parsedEvents *utils.ParsedTxsResultsEvents,
) []model.BlockResultsEvent {
	var extractedEvents []model.BlockResultsEvent

	for _, eventType := range eventTypes {

		events := parsedEvents.GetRawEvents()

		for i := range events {
			splitBykey := utils.NewParsedTxsResultLogEventsSplitByKey(&events[i])

			if len(splitBykey)-1 >= innerMsgIndex {

				rawEvent := splitBykey[innerMsgIndex].GetRawEvents()
				keyIndex := splitBykey[innerMsgIndex].GetKeyIndex()

				var extractedAttributes model.BlockResultsEvent

				// check event amount with inner msg amount
				amount := splitBykey[innerMsgIndex].GetAttributeByKey("amount")
				if amount == nil {
					continue
				}

				innerMsgAmount, ok := innerMsg["amount"].(string)
				if ok && *amount != innerMsgAmount {
					continue
				}

				for _, key := range keyIndex {
					extractedAttributes.Type = rawEvent.Type
					extractedAttributes.Attributes = append(extractedAttributes.Attributes, rawEvent.Attributes[key])
				}

				// check attribute key
				if rawEvent.Type == eventType {
					extractedEvents = append(extractedEvents, extractedAttributes)

				}
			}
		}
	}

	return extractedEvents
}
