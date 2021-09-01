package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_IBC_REVOKE = "MsgRevoke"
const MSG_IBC_REVOKE_CREATED = "MsgRevokeCreated"
const MSG_IBC_REVOKE_FAILED = "MsgRevokeFailed"

type MsgIBCRevoke struct {
	MsgBase

	Params ibc_model.MsgRevokeParams `json:"params"`
}

func NewMsgIBCRevoke(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgRevokeParams,
) *MsgIBCRevoke {
	return &MsgIBCRevoke{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_IBC_REVOKE,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgIBCRevoke) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgIBCRevoke) String() string {
	return render.Render(event)
}

func DecodeMsgIBCRevoke(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgIBCRevoke
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
