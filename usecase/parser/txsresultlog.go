package parser

import (
	"github.com/crypto-com/chainindex/usecase/model"
)

type ParsedTxsResultLog struct {
	typeIndex map[string]int

	rawLog *model.BlockResultsTxsResultLog
}

func NewParsedTxsResultLog(txsResultLog *model.BlockResultsTxsResultLog) *ParsedTxsResultLog {
	log := &ParsedTxsResultLog{
		make(map[string]int),

		txsResultLog,
	}

	for i, event := range txsResultLog.Events {
		log.typeIndex[event.Type] = i
	}
	return log
}

func (log *ParsedTxsResultLog) HasEvent(t string) bool {
	_, ok := log.typeIndex[t]
	return ok
}

func (log *ParsedTxsResultLog) GetEvent(t string) *ParsedTxsResultLogEvent {
	if !log.HasEvent(t) {
		return nil
	}
	return NewParsedTxsResultLogEvent(&log.rawLog.Events[log.typeIndex[t]])
}
