package event

import (
	"bytes"

	jsoniter "github.com/json-iterator/go"

	event_entity "github.com/crypto-com/chainindex/entity/event"
	"github.com/luci/go-render/render"
)

const BLOCK_COMMISSIONED = "BlockCommissioned"

type BlockCommissioned struct {
	event_entity.Base

	Validator string `json:"validator"`
	Amount    string `json:"amount"`
}

func NewBlockCommissioned(blockHeight int64, validator string, amount string) *BlockCommissioned {
	return &BlockCommissioned{
		event_entity.NewBase(event_entity.BaseParams{
			Name:        BLOCK_COMMISSIONED,
			Version:     1,
			BlockHeight: blockHeight,
		}),

		validator,
		amount,
	}

}
func (event *BlockCommissioned) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *BlockCommissioned) String() string {
	return render.Render(event)
}

func DecodeBlockCommissioned(encoded []byte) (event_entity.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *BlockCommissioned
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
