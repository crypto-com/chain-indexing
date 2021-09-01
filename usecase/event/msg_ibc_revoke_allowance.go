package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_IBC_REVOKE_ALLOWANCE = "MsgRevokeAllowance"
const MSG_IBC_REVOKE_ALLOWANCE_CREATED = "MsgRevokeAllowanceCreated"
const MSG_IBC_REVOKE_ALLOWANCE_FAILED = "MsgRevokeAllowanceFailed"

type MsgIBCRevokeAllowance struct {
	MsgBase

	Params ibc_model.MsgRevokeAllowanceParams `json:"params"`
}

func NewMsgIBCRevokeAllowance(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgRevokeAllowanceParams,
) *MsgIBCRevokeAllowance {
	return &MsgIBCRevokeAllowance{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_IBC_REVOKE_ALLOWANCE,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgIBCRevokeAllowance) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgIBCRevokeAllowance) String() string {
	return render.Render(event)
}

func DecodeMsgIBCRevokeAllowance(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgIBCRevokeAllowance
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
