package event

import "github.com/crypto-com/chainindex/entity/model"

const RawBlockCreatedEventName = "RawBlockCreated"

type RawBlockCreatedEvent struct {
	rawBlock *model.RawBlock
}

func NewRawBlockCreatedEvent(rawBlock *model.RawBlock) *RawBlockCreatedEvent {
	return &RawBlockCreatedEvent{
		rawBlock,
	}
}

func (_ *RawBlockCreatedEvent) Name() string {
	return RawBlockCreatedEventName
}

func (_ *RawBlockCreatedEvent) Version() int {
	return 1
}

func (event *RawBlockCreatedEvent) Payload() interface{} {
	return event.rawBlock
}
