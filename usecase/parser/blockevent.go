package parser

import (
	"fmt"

	"github.com/crypto-com/chainindex/usecase/model"
)

// ParsedTxsResultLogEvent is only usable in txs_results log because it has unique attribute key.
// For txs_results.events, begin_block_events, end_block_events, they have to be parsed manually.
type ParsedTxsResultLogEvent struct {
	keyIndex map[string]int

	rawEvent *model.BlockResultsEvent
}

func NewParsedTxsResultLogEvent(rawEvent *model.BlockResultsEvent) *ParsedTxsResultLogEvent {
	event := &ParsedTxsResultLogEvent{
		make(map[string]int),

		rawEvent,
	}

	for i, attribute := range rawEvent.Attributes {
		if event.HasAttribute(attribute.Key) {
			panic(fmt.Sprintf("duplciated attribute key %s", attribute.Key))
		}
		event.keyIndex[attribute.Key] = i
	}

	return event
}

func (log *ParsedTxsResultLogEvent) HasAttribute(key string) bool {
	_, ok := log.keyIndex[key]
	return ok
}

func (log *ParsedTxsResultLogEvent) MustGetAttributeByKey(key string) string {
	attr := log.GetAttributeByKey(key)
	if attr == nil {
		panic(fmt.Sprintf("expected block_results event to have attribute %s, but found none", key))
	}
	return *attr
}

func (log *ParsedTxsResultLogEvent) GetAttributeByKey(key string) *string {
	if !log.HasAttribute(key) {
		return nil
	}
	return &log.rawEvent.Attributes[log.keyIndex[key]].Value
}
