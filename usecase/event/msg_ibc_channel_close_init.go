package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_IBC_CHANNEL_CLOSE_INIT = "/ibc.core.channel.v1.MsgChannelCloseInit"
const MSG_IBC_CHANNEL_CLOSE_INIT_CREATED = "/ibc.core.channel.v1.MsgChannelCloseInit.Created"
const MSG_IBC_CHANNEL_CLOSE_INIT_FAILED = "/ibc.core.channel.v1.MsgChannelCloseInit.Failed"

type MsgIBCChannelCloseInit struct {
	MsgBase

	Params ibc_model.MsgChannelCloseInitParams `json:"params"`
}

func NewMsgIBCChannelCloseInit(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgChannelCloseInitParams,
) *MsgIBCChannelCloseInit {
	return &MsgIBCChannelCloseInit{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_IBC_CHANNEL_CLOSE_INIT,
			Version:         1,

			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgIBCChannelCloseInit) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgIBCChannelCloseInit) String() string {
	return render.Render(event)
}

func DecodeMsgIBCChannelCloseInit(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgIBCChannelCloseInit
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
