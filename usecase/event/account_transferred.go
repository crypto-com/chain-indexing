package event

import (
	"bytes"

	jsoniter "github.com/json-iterator/go"

	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/luci/go-render/render"
)

const ACCOUNT_TRANSFERRED = "AccountTransferred"

type AccountTransferred struct {
	event_entity.Base

	Sender    string    `json:"sender"`
	Recipient string    `json:"recipient"`
	Amount    coin.Coin `json:"amount"`
}

func NewAccountTransferred(blockHeight int64, params model.AccountTransferParams) *AccountTransferred {
	return &AccountTransferred{
		event_entity.NewBase(event_entity.BaseParams{
			Name:        ACCOUNT_TRANSFERRED,
			Version:     1,
			BlockHeight: blockHeight,
		}),

		params.Sender,
		params.Recipient,
		params.Amount,
	}

}
func (event *AccountTransferred) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *AccountTransferred) String() string {
	return render.Render(event)
}

func DecodeAccountTransferred(encoded []byte) (event_entity.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *AccountTransferred
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
