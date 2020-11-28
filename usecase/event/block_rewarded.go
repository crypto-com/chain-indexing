package event

import (
	"bytes"

	jsoniter "github.com/json-iterator/go"

	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/luci/go-render/render"
)

const BLOCK_REWARDED = "BlockRewarded"

type BlockRewarded struct {
	event_entity.Base

	Validator string `json:"validator"`
	Amount    string `json:"amount"`
}

func NewBlockRewarded(blockHeight int64, validator string, amount string) *BlockRewarded {
	return &BlockRewarded{
		event_entity.NewBase(event_entity.BaseParams{
			Name:        BLOCK_REWARDED,
			Version:     1,
			BlockHeight: blockHeight,
		}),

		validator,
		amount,
	}

}
func (event *BlockRewarded) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *BlockRewarded) String() string {
	return render.Render(event)
}

func DecodeBlockRewarded(encoded []byte) (event_entity.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *BlockRewarded
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
