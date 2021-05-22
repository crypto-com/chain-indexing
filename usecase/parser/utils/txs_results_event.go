package utils

import (
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type ParsedTxsResultsEvents struct {
	typeIndex map[string][]int

	rawEvents []model.BlockResultsEvent
}

func NewParsedTxsResultsEvents(txsResultsEvents []model.BlockResultsEvent) *ParsedTxsResultsEvents {
	log := &ParsedTxsResultsEvents{
		make(map[string][]int),

		txsResultsEvents,
	}

	for i, event := range txsResultsEvents {
		if _, exist := log.typeIndex[event.Type]; !exist {
			log.typeIndex[event.Type] = make([]int, 0)
		}
		log.typeIndex[event.Type] = append(log.typeIndex[event.Type], i)
	}
	return log
}

func (log *ParsedTxsResultsEvents) HasEvent(t string) bool {
	_, ok := log.typeIndex[t]
	return ok
}

// Get the last event by type
func (log *ParsedTxsResultsEvents) GetEventByType(t string) *ParsedTxsResultLogEvent {
	if !log.HasEvent(t) {
		return nil
	}
	lastIndex := len(log.typeIndex[t]) - 1
	return NewParsedTxsResultLogEvent(&log.rawEvents[log.typeIndex[t][lastIndex]])
}

// Get first event by type
func (log *ParsedTxsResultsEvents) GetFirstEventByType(t string) *ParsedTxsResultLogEvent {
	if !log.HasEvent(t) {
		return nil
	}
	return NewParsedTxsResultLogEvent(&log.rawEvents[log.typeIndex[t][0]])
}

// Get all events by type
func (log *ParsedTxsResultsEvents) GetEventsByType(t string) []*ParsedTxsResultLogEvent {
	if !log.HasEvent(t) {
		return nil
	}

	logEvents := make([]*ParsedTxsResultLogEvent, 0, len(log.typeIndex[t]))
	for _, index := range log.typeIndex[t] {
		logEvents = append(logEvents, NewParsedTxsResultLogEvent(&log.rawEvents[index]))
	}

	return logEvents
}
