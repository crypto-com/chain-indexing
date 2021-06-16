package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_IBC_UPDATE_CLIENT = "MsgUpdateClientInit"
const MSG_IBC_UPDATE_CLIENT_CREATED = "MsgUpdateClientInitCreated"
const MSG_IBC_UPDATE_CLIENT_FAILED = "MsgUpdateClientInitFailed"

type MsgIBCUpdateClient struct {
	MsgBase

	ibc_model.MsgUpdateClientParams
}

// NewMsgDelegate creates a new instance of MsgDelegate
func NewMsgIBCUpdateClient(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgUpdateClientParams,
) *MsgIBCUpdateClient {
	return &MsgIBCUpdateClient{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_IBC_UPDATE_CLIENT,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgIBCUpdateClient) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgIBCUpdateClient) String() string {
	return render.Render(event)
}

func DecodeMsgIBCUpdateClient(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgIBCUpdateClient
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
