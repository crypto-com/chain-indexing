package parser

import (
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type ParsedTxsResultLog struct {
	typeIndex map[string][]int

	rawLog *model.BlockResultsTxsResultLog
}

func NewParsedTxsResultLog(txsResultLog *model.BlockResultsTxsResultLog) *ParsedTxsResultLog {
	log := &ParsedTxsResultLog{
		make(map[string][]int),

		txsResultLog,
	}

	for i, event := range txsResultLog.Events {
		if _, exist := log.typeIndex[event.Type]; !exist {
			log.typeIndex[event.Type] = make([]int, 0)
		}
		log.typeIndex[event.Type] = append(log.typeIndex[event.Type], i)
	}
	return log
}

func (log *ParsedTxsResultLog) HasEvent(t string) bool {
	_, ok := log.typeIndex[t]
	return ok
}

// Get the last event by type
func (log *ParsedTxsResultLog) GetEventByType(t string) *ParsedTxsResultLogEvent {
	if !log.HasEvent(t) {
		return nil
	}
	lastIndex := len(log.typeIndex[t]) - 1
	return NewParsedTxsResultLogEvent(&log.rawLog.Events[log.typeIndex[t][lastIndex]])
}

// Get first event by type
func (log *ParsedTxsResultLog) GetFirstEventByType(t string) *ParsedTxsResultLogEvent {
	if !log.HasEvent(t) {
		return nil
	}
	return NewParsedTxsResultLogEvent(&log.rawLog.Events[log.typeIndex[t][0]])
}

// Get all events by type
func (log *ParsedTxsResultLog) GetEventsByType(t string) []*ParsedTxsResultLogEvent {
	if !log.HasEvent(t) {
		return nil
	}

	logEvents := make([]*ParsedTxsResultLogEvent, 0, len(log.typeIndex[t]))
	for _, index := range log.typeIndex[t] {
		logEvents = append(logEvents, NewParsedTxsResultLogEvent(&log.rawLog.Events[index]))
	}

	return logEvents
}
