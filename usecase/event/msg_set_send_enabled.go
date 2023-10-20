package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_SET_SEND_ENABLED = "/cosmos.bank.v1beta1.MsgSetSendEnabled"
const MSG_SET_SEND_ENABLED_CREATED = "/cosmos.bank.v1beta1.MsgSetSendEnabled.Created"
const MSG_SET_SEND_ENABLED_FAILED = "/cosmos.bank.v1beta1.MsgSetSendEnabled.Failed"

type MsgSetSendEnabled struct {
	MsgBase

	Params model.MsgSetSendEnabled `json:"params"`
}

func NewMsgSetSendEnabled(
	msgCommonParams MsgCommonParams,
	params model.MsgSetSendEnabled,
) *MsgSetSendEnabled {
	return &MsgSetSendEnabled{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_SET_SEND_ENABLED,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgSetSendEnabled) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgSetSendEnabled) String() string {
	return render.Render(event)
}

func DecodeMsgSetSendEnabled(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgSetSendEnabled
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
