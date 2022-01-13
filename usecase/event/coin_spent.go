package event

import (
	"bytes"

	jsoniter "github.com/json-iterator/go"

	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/luci/go-render/render"
)

const COIN_SPENT = "CoinSpent"

type CoinSpent struct {
	event_entity.Base

	Address string     `json:"address"`
	Amount  coin.Coins `json:"amount"`
}

func NewCoinSpent(blockHeight int64, params model.CoinSpentParams) *CoinSpent {
	return &CoinSpent{
		event_entity.NewBase(event_entity.BaseParams{
			Name:        COIN_SPENT,
			Version:     1,
			BlockHeight: blockHeight,
		}),

		params.Address,
		params.Amount,
	}

}
func (event *CoinSpent) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *CoinSpent) String() string {
	return render.Render(event)
}

func DecodeCoinSpent(encoded []byte) (event_entity.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *CoinSpent
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
