package event

import (
	"bytes"

	"github.com/crypto-com/chainindex/usecase/model"

	"github.com/crypto-com/chainindex/usecase/coin"

	jsoniter "github.com/json-iterator/go"

	event_entity "github.com/crypto-com/chainindex/entity/event"
	"github.com/luci/go-render/render"
)

const BONDING_COMPLETED = "BondingCompleted"

type BondingCompleted struct {
	event_entity.Base

	Delegator string    `json:"delegator"`
	Validator string    `json:"validator"`
	Amount    coin.Coin `json:"amount"`
}

func NewBondingCompleted(blockHeight int64, params model.CompleteBondingParams) *BondingCompleted {
	return &BondingCompleted{
		event_entity.NewBase(event_entity.BaseParams{
			Name:        BONDING_COMPLETED,
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

func DecodeBondingCompleted(encoded []byte) (event_entity.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *BondingCompleted
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
