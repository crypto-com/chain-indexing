package event

import (
	"bytes"

	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"

	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

const MINTED = "Minted"

type Minted struct {
	event_entity.Base

	BondedRatio      string     `json:"bondedRatio"`
	Inflation        string     `json:"inflation"`
	AnnualProvisions coin.Coins `json:"annualProvisions"`
	Amount           coin.Coins `json:"amount"`
}

func NewMinted(blockHeight int64, params model.MintParams) *Minted {
	return &Minted{
		event_entity.NewBase(event_entity.BaseParams{
			Name:        MINTED,
			Version:     1,
			BlockHeight: blockHeight,
		}),

		params.BondedRatio,
		params.Inflation,
		params.AnnualProvisions,
		params.Amount,
	}

}
func (event *Minted) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *Minted) String() string {
	return render.Render(event)
}

func DecodeMinted(encoded []byte) (event_entity.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *Minted
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
