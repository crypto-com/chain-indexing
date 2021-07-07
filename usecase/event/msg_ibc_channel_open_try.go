package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_IBC_CHANNEL_OPEN_TRY = "MsgChannelOpenTry"
const MSG_IBC_CHANNEL_OPEN_TRY_CREATED = "MsgChannelOpenTryCreated"
const MSG_IBC_CHANNEL_OPEN_TRY_FAILED = "MsgChannelOpenTryFailed"

type MsgIBCChannelOpenTry struct {
	MsgBase

	Params ibc_model.MsgChannelOpenTryParams `json:"params"`
}

func NewMsgIBCChannelOpenTry(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgChannelOpenTryParams,
) *MsgIBCChannelOpenTry {
	return &MsgIBCChannelOpenTry{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_IBC_CHANNEL_OPEN_TRY,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgIBCChannelOpenTry) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgIBCChannelOpenTry) String() string {
	return render.Render(event)
}

func DecodeMsgIBCChannelOpenTry(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgIBCChannelOpenTry
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
