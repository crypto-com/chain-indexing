package event

import (
	"fmt"
	"github.com/crypto-com/chainindex/entity/model"
	"github.com/luci/go-render/render"
	"strconv"
)

const RawBlockCreatedEventName = "RawBlockCreated"

type RawBlockCreatedEvent struct {
	BlockBase

	rawBlock *model.RawBlock
}

func NewRawBlockCreatedEvent(rawBlock *model.RawBlock) *RawBlockCreatedEvent {
	height, err := strconv.ParseInt(rawBlock.Result.Block.Header.Height, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Missing block height in raw block: %v", err))
	}

	return &RawBlockCreatedEvent{
		NewBlockBase(height),
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
