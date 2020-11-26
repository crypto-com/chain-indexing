package event

import (
	"bytes"

	"github.com/crypto-com/chainindex/usecase/model/genesis"

	jsoniter "github.com/json-iterator/go"

	event_entity "github.com/crypto-com/chainindex/entity/event"
	"github.com/luci/go-render/render"
)

const GENESIS_CREATED = "GenesisCreated"

type GenesisCreated struct {
	event_entity.Base

	Genesis genesis.Genesis
}

func NewGenesisCreated(genesis genesis.Genesis) *GenesisCreated {
	return &GenesisCreated{
		event_entity.NewBase(event_entity.BaseParams{
			Name:        GENESIS_CREATED,
			Version:     1,
			BlockHeight: int64(0),
		}),

		genesis,
	}

}
func (event *GenesisCreated) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *GenesisCreated) String() string {
	return render.Render(event)
}

func DecodeGenesisCreated(encoded []byte) (event_entity.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *GenesisCreated
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
