package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_CONNECTION_OPEN_INIT = "MsgConnectionOpenInit"
const MSG_CONNECTION_OPEN_INIT_CREATED = "MsgConnectionOpenInitCreated"
const MSG_CONNECTION_OPEN_INIT_FAILED = "MsgConnectionOpenInitFailed"

type MsgConnectionOpenInit struct {
	MsgBase

	ibc_model.MsgConnectionOpenInitParams
}

// NewMsgDelegate creates a new instance of MsgDelegate
func NewMsgConnectionOpenInit(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgConnectionOpenInitParams,
) *MsgConnectionOpenInit {
	return &MsgConnectionOpenInit{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_CONNECTION_OPEN_INIT,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgConnectionOpenInit) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgConnectionOpenInit) String() string {
	return render.Render(event)
}

// DecodeMsgDelegate decodes the event from encoded bytes
func DecodeMsgConnectionOpenInit(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgConnectionOpenInit
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
