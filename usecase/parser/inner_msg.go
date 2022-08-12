package parser

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
)

type EventType struct {
	Type  string `json:"type"`
	Count int64  `json:"count"`
}

func ParseTxsResultsEvents(
	msgs []interface{},
	events []model.BlockResultsEvent,
) []model.BlockResultsTxsResultLog {
	var filteredEvents []model.BlockResultsEvent
	var resultLog []model.BlockResultsTxsResultLog

	// remove event type = tx and message action = MsgExec
	for _, event := range events {
		if event.Type == "message" && event.Attributes[0].Key == "action" {
			continue
		}
		filteredEvents = append(filteredEvents, event)
	}

	parsedEvents := utils.NewParsedTxsResultsEvents(filteredEvents)
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

	return resultLog
}

func ParseInnerMsgsEvents(
	innerMsgType string,
	innerMsgIndex int,
	parsedEvents *utils.ParsedTxsResultsEvents,
) []model.BlockResultsEvent {
	var extractedEvents []model.BlockResultsEvent
	switch innerMsgType {

	// cosmos bank
	case "/cosmos.bank.v1beta1.MsgSend":
		extractedEvents = MsgSend(parsedEvents, innerMsgIndex)

	// cosmos distribution
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

func MsgSend(events *utils.ParsedTxsResultsEvents,
	innerMsgIndex int,
) []model.BlockResultsEvent {
	var extractedEvents []model.BlockResultsEvent
	eventTypes := []EventType{
		{
			Type:  "coin_spent",
			Count: 1,
		},
		{
			Type:  "coin_received",
			Count: 1,
		},
		{
			Type:  "transfer",
			Count: 1,
		},
		{
			Type:  "message",
			Count: 2,
		},
	}

	// extract events
	extractedEvents = extractMsgEvents(eventTypes, events)

	return extractedEvents
}

func MsgSetWithdrawAddress(events *utils.ParsedTxsResultsEvents,
	innerMsgIndex int,
) []model.BlockResultsEvent {
	var extractedEvents []model.BlockResultsEvent
	//
	eventTypes := []EventType{
		{
			Type:  "set_withdraw_address",
			Count: 1,
		},
		{
			Type:  "message",
			Count: 1,
		},
	}

	// extract events
	extractedEvents = extractMsgEvents(eventTypes, events)
	return extractedEvents
}

func MsgWithdrawDelegatorReward(events *utils.ParsedTxsResultsEvents,
	innerMsgIndex int,
) []model.BlockResultsEvent {
	var extractedEvents []model.BlockResultsEvent
	eventTypes := []EventType{
		{
			Type:  "coin_spent",
			Count: 1,
		},
		{
			Type:  "coin_received",
			Count: 1,
		},
		{
			Type:  "transfer",
			Count: 1,
		},
		{
			Type:  "message",
			Count: 2,
		},
		{
			Type:  "withdraw_rewards",
			Count: 1,
		},
	}
	eventTypesWithoutAmount := []EventType{
		{
			Type:  "message",
			Count: 1,
		},
		{
			Type:  "withdraw_rewards",
			Count: 1,
		},
	}

	// extract events
	extractedEvents = extractMsgEvents(eventTypes, events)

	if len(extractedEvents) <= 0 {
		// extract events
		extractedEvents = extractMsgEvents(eventTypesWithoutAmount, events)
	}

	return extractedEvents
}

func MsgWithdrawValidatorCommission(events *utils.ParsedTxsResultsEvents,
	innerMsgIndex int,
) []model.BlockResultsEvent {
	var extractedEvents []model.BlockResultsEvent
	eventTypes := []EventType{
		{
			Type:  "coin_spent",
			Count: 1,
		},
		{
			Type:  "coin_received",
			Count: 1,
		},
		{
			Type:  "transfer",
			Count: 1,
		},
		{
			Type:  "message",
			Count: 2,
		},
		{
			Type:  "withdraw_commission",
			Count: 1,
		},
	}

	eventTypesWithoutAmount := []EventType{
		{
			Type:  "message",
			Count: 1,
		},
		{
			Type:  "withdraw_commission",
			Count: 1,
		},
	}
	// extract events
	extractedEvents = extractMsgEvents(eventTypes, events)
	if len(extractedEvents) <= 0 {
		// extract events
		extractedEvents = extractMsgEvents(eventTypesWithoutAmount, events)
	}

	return extractedEvents
}

func MsgFundCommunityPool(events *utils.ParsedTxsResultsEvents,
	innerMsgIndex int,
) []model.BlockResultsEvent {
	var extractedEvents []model.BlockResultsEvent
	eventTypes := []EventType{
		{
			Type:  "coin_spent",
			Count: 1,
		},
		{
			Type:  "coin_received",
			Count: 1,
		},
		{
			Type:  "transfer",
			Count: 1,
		},
		{
			Type:  "message",
			Count: 2,
		},
	}

	eventTypesWithoutAmount := []EventType{
		{
			Type:  "message",
			Count: 1,
		},
	}

	// extract events
	extractedEvents = extractMsgEvents(eventTypes, events)
	if len(extractedEvents) <= 0 {
		// extract events
		extractedEvents = extractMsgEvents(eventTypesWithoutAmount, events)
	}

	return extractedEvents
}

func extractMsgEvents(eventTypes []EventType, events *utils.ParsedTxsResultsEvents) []model.BlockResultsEvent {
	var extractedEvents []model.BlockResultsEvent

	for _, eventType := range eventTypes {
		for i := 0; i < int(eventType.Count); i++ {
			if i > len(events.GetTypeIndex(eventType.Type)) || len(events.GetTypeIndex(eventType.Type)) <= 0 {
				return []model.BlockResultsEvent{}
			}
			eventIndex := events.GetTypeIndex(eventType.Type)[0]
			extractedEvents = append(extractedEvents, events.GetRawEvents()[eventIndex])
			events.RemoveIndexType(eventType.Type, 0)

		}
	}

	return extractedEvents
}
