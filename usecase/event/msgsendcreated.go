package event

import (
	"bytes"

	"github.com/crypto-com/chainindex/usecase/coin"

	entity_event "github.com/crypto-com/chainindex/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_SEND_CREATED_NAME = "MsgSendCreated"

type MsgSendCreated struct {
	entity_event.Base

	TxHash      string    `json:"txHash"`
	MsgIndex    int       `json:"msgIndex"`
	FromAddress string    `json:"fromAddress"`
	ToAddress   string    `json:"toAddress"`
	Amount      coin.Coin `json:"amount"`
}

func NewMsgSendCreated(blockHeight int64, txHash string, msgIndex int, params MsgSendCreatedParams) *MsgSendCreated {
	return &MsgSendCreated{
		entity_event.NewBase(entity_event.BaseParams{
			Name:        MSG_SEND_CREATED_NAME,
			Version:     1,
			BlockHeight: blockHeight,
		}),

		txHash,
		msgIndex,
		params.FromAddress,
		params.ToAddress,
		params.Amount,
	}
}

type MsgSendCreatedParams struct {
	FromAddress string
	ToAddress   string
	Amount      coin.Coin
}

func (event *MsgSendCreated) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgSendCreated) String() string {
	return render.Render(event)
}

func DecodeMsgSendCreated(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgSendCreated
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
