package event

import (
	"bytes"

	jsoniter "github.com/json-iterator/go"

	event_entity "github.com/crypto-com/chainindex/entity/event"
	"github.com/luci/go-render/render"
)

const BLOCK_PROPOSER_REWARDED = "BlockProposerRewarded"

type BlockProposerRewarded struct {
	event_entity.Base

	Validator string `json:"validator"`
	Amount    string `json:"amount"`
}

func NewProposerRewarded(blockHeight int64, validator string, amount string) *BlockProposerRewarded {
	return &BlockProposerRewarded{
		event_entity.NewBase(event_entity.BaseParams{
			Name:        BLOCK_PROPOSER_REWARDED,
			Version:     1,
			BlockHeight: blockHeight,
		}),

		validator,
		amount,
	}

}
func (event *BlockProposerRewarded) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *BlockProposerRewarded) String() string {
	return render.Render(event)
}

func DecodeProposerRewarded(encoded []byte) (event_entity.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *BlockProposerRewarded
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
