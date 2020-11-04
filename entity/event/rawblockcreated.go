package event

import (
	"github.com/crypto-com/chainindex/ddd/event"
	"github.com/crypto-com/chainindex/entity/model"
	"github.com/luci/go-render/render"
)

const RawBlockCreatedEventName = "RawBlockCreated"

type RawBlockCreatedEvent struct {
	event.Base

	rawBlock *model.RawBlock
}

func NewRawBlockCreatedEvent(rawBlock *model.RawBlock) *RawBlockCreatedEvent {
	return &RawBlockCreatedEvent{
		event.NewBase(),

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

func (event *RawBlockCreatedEvent) String() string {
	return render.Render(event)
}
