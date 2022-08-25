package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_IBC_CONNECTION_OPEN_CONFIRM = "/ibc.core.connection.v1.MsgConnectionOpenConfirm"
const MSG_IBC_CONNECTION_OPEN_CONFIRM_CREATED = "MsgConnectionOpenConfirmCreated"
const MSG_IBC_CONNECTION_OPEN_CONFIRM_FAILED = "MsgConnectionOpenConfirmFailed"

type MsgIBCConnectionOpenConfirm struct {
	MsgBase

	Params ibc_model.MsgConnectionOpenConfirmParams `json:"params"`
}

func NewMsgIBCConnectionOpenConfirm(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgConnectionOpenConfirmParams,
) *MsgIBCConnectionOpenConfirm {
	return &MsgIBCConnectionOpenConfirm{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_IBC_CONNECTION_OPEN_CONFIRM,
			MsgSuccess:      MSG_IBC_CONNECTION_OPEN_CONFIRM_CREATED,
			MsgFailed:       MSG_IBC_CONNECTION_OPEN_CONFIRM_FAILED,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgIBCConnectionOpenConfirm) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgIBCConnectionOpenConfirm) String() string {
	return render.Render(event)
}

// DecodeMsgIBCConnectionOpenConfirm decodes the event from encoded bytes
func DecodeMsgIBCConnectionOpenConfirm(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgIBCConnectionOpenConfirm
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
