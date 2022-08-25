package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_IBC_CONNECTION_OPEN_TRY = "/ibc.core.connection.v1.MsgConnectionOpenTry"
const MSG_IBC_CONNECTION_OPEN_TRY_CREATED = "MsgConnectionOpenTryCreated"
const MSG_IBC_CONNECTION_OPEN_TRY_FAILED = "MsgConnectionOpenTryFailed"

type MsgIBCConnectionOpenTry struct {
	MsgBase

	Params ibc_model.MsgConnectionOpenTryParams `json:"params"`
}

func NewMsgIBCConnectionOpenTry(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgConnectionOpenTryParams,
) *MsgIBCConnectionOpenTry {
	return &MsgIBCConnectionOpenTry{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_IBC_CONNECTION_OPEN_TRY,
			MsgSuccess:      MSG_IBC_CONNECTION_OPEN_TRY_CREATED,
			MsgFailed:       MSG_IBC_CONNECTION_OPEN_TRY_FAILED,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgIBCConnectionOpenTry) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgIBCConnectionOpenTry) String() string {
	return render.Render(event)
}

// DecodeMsgIBCConnectionOpenTry decodes the event from encoded bytes
func DecodeMsgIBCConnectionOpenTry(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgIBCConnectionOpenTry
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
