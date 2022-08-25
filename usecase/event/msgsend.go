package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/coin"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_SEND = "/cosmos.bank.v1beta1.MsgSend"
const MSG_SEND_CREATED = "MsgSendCreated"
const MSG_SEND_FAILED = "MsgSendFailed"

type MsgSend struct {
	MsgBase

	FromAddress string     `json:"fromAddress"`
	ToAddress   string     `json:"toAddress"`
	Amount      coin.Coins `json:"amount"`
}

func NewMsgSend(msgCommonParams MsgCommonParams, params MsgSendCreatedParams) *MsgSend {
	return &MsgSend{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_SEND,
			MsgSuccess:      MSG_SEND_CREATED,
			MsgFailed:       MSG_SEND_FAILED,
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
	Amount      coin.Coins
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
