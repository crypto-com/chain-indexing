package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_IBC_CHANNEL_OPEN_CONFIRM = "/ibc.core.channel.v1.MsgChannelOpenConfirm"
const MSG_IBC_CHANNEL_OPEN_CONFIRM_CREATED = "/ibc.core.channel.v1.MsgChannelOpenConfirm.Created"
const MSG_IBC_CHANNEL_OPEN_CONFIRM_FAILED = "/ibc.core.channel.v1.MsgChannelOpenConfirm.Failed"

type MsgIBCChannelOpenConfirm struct {
	MsgBase

	Params ibc_model.MsgChannelOpenConfirmParams `json:"params"`
}

func NewMsgIBCChannelOpenConfirm(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgChannelOpenConfirmParams,
) *MsgIBCChannelOpenConfirm {
	return &MsgIBCChannelOpenConfirm{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_IBC_CHANNEL_OPEN_CONFIRM,
			Version:         1,

			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgIBCChannelOpenConfirm) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgIBCChannelOpenConfirm) String() string {
	return render.Render(event)
}

func DecodeMsgIBCChannelOpenConfirm(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgIBCChannelOpenConfirm
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
