package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_REVOKE_ALLOWANCE = "MsgRevokeAllowance"
const MSG_REVOKE_ALLOWANCE_CREATED = "MsgRevokeAllowanceCreated"
const MSG_REVOKE_ALLOWANCE_FAILED = "MsgRevokeAllowanceFailed"

type MsgRevokeAllowance struct {
	MsgBase

	Params ibc_model.MsgRevokeAllowanceParams `json:"params"`
}

func NewMsgRevokeAllowance(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgRevokeAllowanceParams,
) *MsgRevokeAllowance {
	return &MsgRevokeAllowance{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_REVOKE_ALLOWANCE,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgRevokeAllowance) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgRevokeAllowance) String() string {
	return render.Render(event)
}

func DecodeMsgRevokeAllowance(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgRevokeAllowance
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
