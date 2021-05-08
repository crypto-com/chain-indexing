package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_UPDATE_CLIENT = "MsgUpdateClient"
const MSG_UPDATE_CLIENT_CREATED = "MsgUpdateClientCreated"
const MSG_UPDATE_CLIENT_FAILED = "MsgUpdateClientFailed"

type MsgUpdateClient struct {
	MsgBase

	ibc_model.MsgUpdateClientParams
}

// NewMsgDelegate creates a new instance of MsgDelegate
func NewMsgUpdateClient(msgCommonParams MsgCommonParams, params ibc_model.MsgUpdateClientParams) *MsgUpdateClient {
	return &MsgUpdateClient{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_UPDATE_CLIENT,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgUpdateClient) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgUpdateClient) String() string {
	return render.Render(event)
}

// DecodeMsgDelegate decodes the event from encoded bytes
func DecodeMsgUpdateClient(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgUpdateClient
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
