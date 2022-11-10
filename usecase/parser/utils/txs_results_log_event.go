package utils

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/usecase/model"
)

// ParsedTxsResultLogEvent is only usable in txs_results log because it has unique attribute key.
// For txs_results.events, begin_block_events, end_block_events, they have to be parsed manually.
type ParsedTxsResultLogEvent struct {
	keyIndex map[string]int

	rawEvent model.BlockResultsEvent
}

func NewParsedTxsResultLogEvent(rawEvent *model.BlockResultsEvent) *ParsedTxsResultLogEvent {
	event := &ParsedTxsResultLogEvent{
		make(map[string]int),

		*rawEvent,
	}

	for i, attribute := range rawEvent.Attributes {
		if event.HasAttribute(attribute.Key) {
			panic(fmt.Sprintf("duplicated attribute key `%s`", attribute.Key))
		}
		event.keyIndex[attribute.Key] = i
	}

	return event
}

// In the txs_results log, multiple event of the same types are grouped
// together into single event. This method pars txs_results log and split into
// a new event when a same key name appears.
func NewParsedTxsResultLogEventsSplitByKey(
	rawEvent *model.BlockResultsEvent,
) []ParsedTxsResultLogEvent {
	events := make([]ParsedTxsResultLogEvent, 0)
	event := ParsedTxsResultLogEvent{
		make(map[string]int),

		*rawEvent,
	}

	for i, attribute := range rawEvent.Attributes {
		if event.HasAttribute(attribute.Key) {
			events = append(events, event)

			event = ParsedTxsResultLogEvent{
				make(map[string]int),

				*rawEvent,
			}
		}
		event.keyIndex[attribute.Key] = i
	}
	events = append(events, event)

	return events
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
