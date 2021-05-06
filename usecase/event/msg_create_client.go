package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_CREATE_CLIENT = "MsgCreateClient"
const MSG_CREATE_CLIENT_CREATED = "MsgCreateClientCreated"
const MSG_CREATE_CLIENT_FAILED = "MsgCreateClientFailed"

type MsgCreateClient struct {
	MsgBase

	ibc_model.MsgCreateClientParams
}

// NewMsgDelegate creates a new instance of MsgDelegate
func NewMsgCreateClient(msgCommonParams MsgCommonParams, params ibc_model.MsgCreateClientParams) *MsgCreateClient {
	return &MsgCreateClient{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_CREATE_CLIENT,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgCreateClient) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgCreateClient) String() string {
	return render.Render(event)
}

// DecodeMsgDelegate decodes the event from encoded bytes
func DecodeMsgCreateClient(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgCreateClient
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
