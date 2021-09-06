package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_REVOKE = "MsgRevoke"
const MSG_REVOKE_CREATED = "MsgRevokeCreated"
const MSG_REVOKE_FAILED = "MsgRevokeFailed"

type MsgRevoke struct {
	MsgBase

	Params ibc_model.MsgRevokeParams `json:"params"`
}

func NewMsgRevoke(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgRevokeParams,
) *MsgRevoke {
	return &MsgRevoke{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_REVOKE,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgRevoke) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgRevoke) String() string {
	return render.Render(event)
}

func DecodeMsgRevoke(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgRevoke
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
