package blocksigned

import (
	"bytes"

	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"

	entity_event "github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/internal/utctime"
)

const NAME = "BlockSigned"

type BlockSigned struct {
	entity_event.Base

	ValidatorAddress string          `json:"validatorAddress"`
	IsProposer       bool            `json:"isProposer"`
	Timestamp        utctime.UTCTime `json:"timestamp"`
	Signature        string          `json:"signature"`
}

func New(
	blockHeight int64,
	validatorAddress string,
	isProposer bool,
	timestamp utctime.UTCTime,
	signature string,
) *BlockSigned {
	return &BlockSigned{
		entity_event.NewBase(blockHeight),

		validatorAddress,
		isProposer,
		timestamp,
		signature,
	}
}

func (_ *BlockSigned) Name() string {
	return NAME
}

func (_ *BlockSigned) Version() int {
	return 1
}

func (event *BlockSigned) String() string {
	return render.Render(event)
}

func Decode(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *BlockSigned
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
