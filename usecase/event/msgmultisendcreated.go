package event

import (
	"bytes"

	"github.com/crypto-com/chainindex/usecase/model"

	entity_event "github.com/crypto-com/chainindex/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_MULTI_SEND_CREATED_NAME = "MsgSendCreated"

type MsgMultiSendCreated struct {
	entity_event.Base

	TxHash   string                     `json:"txHash"`
	MsgIndex int                        `json:"msgIndex"`
	Inputs   []model.MsgMultiSendInput  `json:"inputs"`
	Outputs  []model.MsgMultiSendOutput `json:"outputs"`
}

func NewMsgMultiSendCreated(blockHeight int64, txHash string, msgIndex int, params model.MsgMultiSendParams) *MsgMultiSendCreated {
	return &MsgMultiSendCreated{
		entity_event.NewBase(entity_event.BaseParams{
			Name:        MSG_MULTI_SEND_CREATED_NAME,
			Version:     1,
			BlockHeight: blockHeight,
		}),

		txHash,
		msgIndex,
		params.Inputs,
		params.Outputs,
	}
}

func (event *MsgMultiSendCreated) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgMultiSendCreated) String() string {
	return render.Render(event)
}

func DecodeMsgMultiSendCreated(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgMultiSendCreated
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
