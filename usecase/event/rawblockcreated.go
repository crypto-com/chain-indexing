package event

import (
	"fmt"
	"strconv"

	"github.com/luci/go-render/render"

	entity_event "github.com/crypto-com/chainindex/entity/event"
	usecase_model "github.com/crypto-com/chainindex/usecase/model"
)

const RAW_BLOCK_CREATED = "RawBlockCreated"

type RawBlockCreated struct {
	entity_event.Base

	RawBlock *usecase_model.RawBlock `json:"rawBlock"`
}

func NewRawBlockCreated(rawBlock *usecase_model.RawBlock) *RawBlockCreated {
	height, err := strconv.ParseInt(rawBlock.Result.Block.Header.Height, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Missing block height in raw block: %v", err))
	}

	return &RawBlockCreated{
		entity_event.NewBase(height),

		rawBlock,
	}
}

func (_ *RawBlockCreated) Name() string {
	return RAW_BLOCK_CREATED
}

func (_ *RawBlockCreated) Version() int {
	return 1
}

func (evt *RawBlockCreated) String() string {
	return render.Render(evt)
}
