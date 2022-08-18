package parser

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
)

func ParseTxsResultsEvents(
	msgs []interface{},
	logs []model.BlockResultsTxsResultLog,
) []model.BlockResultsTxsResultLog {
	var resultLog []model.BlockResultsTxsResultLog
	var parsedEvents *utils.ParsedTxsResultsEvents

	if len(logs) == 1 && len(logs) != len(msgs) {
		log := logs[0]
		parsedEvents = utils.NewParsedTxsResultsEvents(log.Events)

		for innerMsgIndex, innerMsgInterface := range msgs {
			innerMsg, ok := innerMsgInterface.(map[string]interface{})
			if !ok {
				panic(fmt.Errorf("error parsing MsgExec.msgs[%v] to map[string]interface{}: %v", innerMsgIndex, innerMsgInterface))
			}

			innerMsgType, ok := innerMsg["@type"].(string)
			if !ok {
				panic(fmt.Errorf("error missing '@type' in MsgExec.msgs[%v]: %v", innerMsgIndex, innerMsg))
			}

			validateEvents := ParseInnerMsgsEvents(innerMsgType, innerMsgIndex, parsedEvents)

			log := model.BlockResultsTxsResultLog{
				MsgIndex: innerMsgIndex,
				Events:   validateEvents,
			}
			resultLog = append(resultLog, log)
		}
	}

	return resultLog
}

func ParseInnerMsgsEvents(
	innerMsgType string,
	innerMsgIndex int,
	parsedEvents *utils.ParsedTxsResultsEvents,
) []model.BlockResultsEvent {
	var extractedEvents []model.BlockResultsEvent
	switch innerMsgType {

	// bank
	case "/cosmos.bank.v1beta1.MsgSend":
		extractedEvents = MsgSend(parsedEvents, innerMsgIndex)

	// distribution
	case "/cosmos.distribution.v1beta1.MsgSetWithdrawAddress":
		extractedEvents = MsgSetWithdrawAddress(parsedEvents, innerMsgIndex)
	case "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward":
		extractedEvents = MsgWithdrawDelegatorReward(parsedEvents, innerMsgIndex)
	case "/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission":
		extractedEvents = MsgWithdrawValidatorCommission(parsedEvents, innerMsgIndex)
	case "/cosmos.distribution.v1beta1.MsgFundCommunityPool":
		extractedEvents = MsgFundCommunityPool(parsedEvents, innerMsgIndex)
	}

	return extractedEvents
}

func MsgSend(parsedEvents *utils.ParsedTxsResultsEvents,
	innerMsgIndex int,
) []model.BlockResultsEvent {
	eventTypes := []string{"coin_spent", "coin_received", "transfer", "message"}

	return extract(innerMsgIndex, eventTypes, []string{}, parsedEvents)
}

func MsgSetWithdrawAddress(parsedEvents *utils.ParsedTxsResultsEvents,
	innerMsgIndex int,
) []model.BlockResultsEvent {
	eventTypes := []string{"set_withdraw_address", "message"}

	return extract(innerMsgIndex, eventTypes, []string{}, parsedEvents)
}

func MsgWithdrawDelegatorReward(parsedEvents *utils.ParsedTxsResultsEvents,
	innerMsgIndex int,
) []model.BlockResultsEvent {
	eventTypes := []string{"coin_spent", "coin_received", "transfer", "message", "withdraw_rewards"}

	eventTypesNoTransfer := []string{
		"message",
		"withdraw_rewards",
	}
	return extract(innerMsgIndex, eventTypes, eventTypesNoTransfer, parsedEvents)
}

func MsgWithdrawValidatorCommission(parsedEvents *utils.ParsedTxsResultsEvents,
	innerMsgIndex int,
) []model.BlockResultsEvent {
	eventTypes := []string{"coin_spent", "coin_received", "transfer", "message", "withdraw_commission"}

	eventTypesNoTransfer := []string{"message", "withdraw_commission"}

	return extract(innerMsgIndex, eventTypes, eventTypesNoTransfer, parsedEvents)
}

func MsgFundCommunityPool(parsedEvents *utils.ParsedTxsResultsEvents,
	innerMsgIndex int,
) []model.BlockResultsEvent {
	eventTypes := []string{"coin_spent", "coin_received", "transfer", "message"}

	eventTypesNoTransfer := []string{"message"}

	return extract(innerMsgIndex, eventTypes, eventTypesNoTransfer, parsedEvents)
}

func extract(innerMsgIndex int, eventTypes []string, eventTypesRetry []string, parsedEvents *utils.ParsedTxsResultsEvents) []model.BlockResultsEvent {
	// extract events
	extractedEvents := getMsgEvents(innerMsgIndex, eventTypes, parsedEvents)

	// retry
	if len(extractedEvents) <= 0 {
		extractedEvents = getMsgEvents(innerMsgIndex, eventTypesRetry, parsedEvents)
	}

	return extractedEvents
}

func getMsgEvents(innerMsgIndex int, eventTypes []string, parsedEvents *utils.ParsedTxsResultsEvents) []model.BlockResultsEvent {
	var extractedEvents []model.BlockResultsEvent

	for _, eventType := range eventTypes {

		events := parsedEvents.GetRawEvents()

		for _, event := range events {
			splitBykey := utils.NewParsedTxsResultLogEventsSplitByKey(&event)

			rawEvent := splitBykey[innerMsgIndex].GetRawEvents()
			keyIndex := splitBykey[innerMsgIndex].GetKeyIndex()

			var extractedAttributes model.BlockResultsEvent

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

	return extractedEvents
}
