package parser

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/usecase/model"
)

func parseInnerMsgsEvents(
	innerMsgType string,
	events []model.BlockResultsEvent,
) ([]model.BlockResultsEvent, []model.BlockResultsEvent) {
	var filteredEvents []model.BlockResultsEvent
	var validateEvents []model.BlockResultsEvent
	var remainingEvents []model.BlockResultsEvent

	for _, event := range events {
		if event.Type == "tx" {
			continue
		}
		if event.Type == "message" && event.Attributes[0].Key == "action" {
			continue
		}
		filteredEvents = append(filteredEvents, event)
	}
	switch innerMsgType {
	case "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward":
		validateEvents, remainingEvents = withdrawDelegatorReward(filteredEvents)

		fmt.Println("===> validateEvents:", validateEvents)
		fmt.Println("===> remainingEvents:", remainingEvents)
	}
	return validateEvents, remainingEvents
}

func withdrawDelegatorReward(events []model.BlockResultsEvent,
) ([]model.BlockResultsEvent, []model.BlockResultsEvent) {
	var withdrawRewardEvent []model.BlockResultsEvent
	var remainingEvents []model.BlockResultsEvent

	validateType := []string{"coin_spent", "coin_received", "transfer", "message", "withdraw_rewards", "message"}
	validateTypeWithoutAmount := []string{"withdraw_rewards", "message"}

	withdrawRewardEvent, remainingEvents = validateEvents(validateType, events)

	if len(withdrawRewardEvent) <= 0 {
		withdrawRewardEvent, remainingEvents = validateEvents(validateTypeWithoutAmount, events)
	}

	return withdrawRewardEvent, remainingEvents
}

func validateEvents(validateType []string, events []model.BlockResultsEvent) ([]model.BlockResultsEvent, []model.BlockResultsEvent) {
	var validateEvents []model.BlockResultsEvent
	var remainingEvents []model.BlockResultsEvent
	index := 0

	for _, event := range events {
		if validateType[index] == event.Type && len(validateEvents) < len(validateType) {
			validateEvents = append(validateEvents, event)

		} else {
			remainingEvents = append(remainingEvents, event)
		}

		if index < len(validateType)-1 {
			index++
		}
	}
	return validateEvents, remainingEvents
}
