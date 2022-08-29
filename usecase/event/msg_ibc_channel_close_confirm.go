package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_IBC_CHANNEL_CLOSE_CONFIRM = "/ibc.core.channel.v1.MsgChannelCloseConfirm"
const MSG_IBC_CHANNEL_CLOSE_CONFIRM_CREATED = "/ibc.core.channel.v1.MsgChannelCloseConfirm.Created"
const MSG_IBC_CHANNEL_CLOSE_CONFIRM_FAILED = "/ibc.core.channel.v1.MsgChannelCloseConfirm.Failed"

type MsgIBCChannelCloseConfirm struct {
	MsgBase

	Params ibc_model.MsgChannelCloseConfirmParams `json:"params"`
}

func NewMsgIBCChannelCloseConfirm(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgChannelCloseConfirmParams,
) *MsgIBCChannelCloseConfirm {
	return &MsgIBCChannelCloseConfirm{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_IBC_CHANNEL_CLOSE_CONFIRM,
			Version:         1,

			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgIBCChannelCloseConfirm) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgIBCChannelCloseConfirm) String() string {
	return render.Render(event)
}

func DecodeMsgIBCChannelCloseConfirm(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgIBCChannelCloseConfirm
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
