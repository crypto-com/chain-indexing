package event

import (
	"bytes"

	"github.com/crypto-com/chainindex/usecase/coin"

	entity_event "github.com/crypto-com/chainindex/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_SEND = "MsgSend"
const MSG_SEND_CREATED = "MsgSendCreated"
const MSG_SEND_FAILED = "MsgSendFailed"

type MsgSend struct {
	MsgBase

	FromAddress string    `json:"fromAddress"`
	ToAddress   string    `json:"toAddress"`
	Amount      coin.Coin `json:"amount"`
}

func NewMsgSend(msgCommonParams MsgCommonParams, params MsgSendCreatedParams) *MsgSend {
	return &MsgSend{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_SEND,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

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

func (event *MsgSend) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgSend) String() string {
	return render.Render(event)
}

func DecodeMsgSend(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgSend
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
