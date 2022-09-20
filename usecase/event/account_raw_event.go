package event

import (
	"bytes"

	"github.com/luci/go-render/render"

	jsoniter "github.com/json-iterator/go"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	usecase_model "github.com/crypto-com/chain-indexing/usecase/model"
)

const ACCOUNT_RAW_EVENT_CREATED = "AccountRawEventCreated"

type AccountRawEventCreated struct {
	entity_event.Base
	Account string `json:"account"`
	Event model.BlockEvent `json:"event"`
}

func NewAccountRawEventCreated(blockHeight int64, account string, event usecase_model.BlockEvent) *AccountRawEventCreated {
	return &AccountRawEventCreated{
		entity_event.NewBase(entity_event.BaseParams{
			Name:        ACCOUNT_RAW_EVENT_CREATED,
			Version:     1,
			BlockHeight: blockHeight,
		}),
		account,
		event,
	}
}

func (event *AccountRawEventCreated) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *AccountRawEventCreated) String() string {
	return render.Render(event)
}

func DecodeAccountRawEventCreated(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *AccountRawEventCreated
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
