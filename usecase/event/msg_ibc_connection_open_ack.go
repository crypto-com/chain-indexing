package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_IBC_CONNECTION_OPEN_ACK = "/ibc.core.connection.v1.MsgConnectionOpenAck"
const MSG_IBC_CONNECTION_OPEN_ACK_CREATED = "/ibc.core.connection.v1.MsgConnectionOpenAck.Created"
const MSG_IBC_CONNECTION_OPEN_ACK_FAILED = "/ibc.core.connection.v1.MsgConnectionOpenAck.Failed"

type MsgIBCConnectionOpenAck struct {
	MsgBase

	Params ibc_model.MsgConnectionOpenAckParams `json:"params"`
}

func NewMsgIBCConnectionOpenAck(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgConnectionOpenAckParams,
) *MsgIBCConnectionOpenAck {
	return &MsgIBCConnectionOpenAck{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_IBC_CONNECTION_OPEN_ACK,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgIBCConnectionOpenAck) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgIBCConnectionOpenAck) String() string {
	return render.Render(event)
}

// DecodeMsgIBCConnectionOpenAck decodes the event from encoded bytes
func DecodeMsgIBCConnectionOpenAck(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgIBCConnectionOpenAck
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
