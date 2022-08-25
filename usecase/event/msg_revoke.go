package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_REVOKE = "/cosmos.authz.v1beta1.MsgRevoke"
const MSG_REVOKE_CREATED = "MsgRevokeCreated"
const MSG_REVOKE_FAILED = "MsgRevokeFailed"

type MsgRevoke struct {
	MsgBase

	Params model.MsgRevokeParams `json:"params"`
}

func NewMsgRevoke(
	msgCommonParams MsgCommonParams,
	params model.MsgRevokeParams,
) *MsgRevoke {
	return &MsgRevoke{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_REVOKE,
			MsgSuccess:      MSG_REVOKE_CREATED,
			MsgFailed:       MSG_REVOKE_FAILED,
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
