package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	"github.com/crypto-com/chain-indexing/usecase/coin"

	jsoniter "github.com/json-iterator/go"

	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/luci/go-render/render"
)

const UNBONDING_COMPLETED = "UnbondingCompleted"

type BondingCompleted struct {
	event_entity.Base

	Delegator string     `json:"delegator"`
	Validator string     `json:"validator"`
	Amount    coin.Coins `json:"amount"`
}

func NewUnbondingCompleted(blockHeight int64, params model.CompleteBondingParams) *BondingCompleted {
	return &BondingCompleted{
		event_entity.NewBase(event_entity.BaseParams{
			Name:        UNBONDING_COMPLETED,
			Version:     1,
			BlockHeight: blockHeight,
		}),

		params.Delegator,
		params.Validator,
		params.Amount,
	}

}
func (event *BondingCompleted) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *BondingCompleted) String() string {
	return render.Render(event)
}

func DecodeUnbondingCompleted(encoded []byte) (event_entity.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *BondingCompleted
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
