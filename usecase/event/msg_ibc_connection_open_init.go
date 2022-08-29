package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_IBC_CONNECTION_OPEN_INIT = "/ibc.core.connection.v1.MsgConnectionOpenInit"
const MSG_IBC_CONNECTION_OPEN_INIT_CREATED = "/ibc.core.connection.v1.MsgConnectionOpenInit.Created"
const MSG_IBC_CONNECTION_OPEN_INIT_FAILED = "/ibc.core.connection.v1.MsgConnectionOpenInit.Failed"

type MsgIBCConnectionOpenInit struct {
	MsgBase

	Params ibc_model.MsgConnectionOpenInitParams `json:"params"`
}

func NewMsgIBCConnectionOpenInit(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgConnectionOpenInitParams,
) *MsgIBCConnectionOpenInit {
	return &MsgIBCConnectionOpenInit{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_IBC_CONNECTION_OPEN_INIT,
			Version:         1,

			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgIBCConnectionOpenInit) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgIBCConnectionOpenInit) String() string {
	return render.Render(event)
}

func DecodeMsgIBCConnectionOpenInit(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgIBCConnectionOpenInit
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
