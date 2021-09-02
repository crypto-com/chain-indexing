package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_IBC_GRANT_ALLOWANCE = "MsgGrantAllowance"
const MSG_IBC_GRANT_ALLOWANCE_CREATED = "MsgGrantAllowanceCreated"
const MSG_IBC_GRANT_ALLOWANCE_FAILED = "MsgGrantAllowanceFailed"

type MsgIBCGrantAllowance struct {
	MsgBase

	Params ibc_model.MsgGrantAllowanceParams `json:"params"`
}

func NewMsgIBCGrantAllowance(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgGrantAllowanceParams,
) *MsgIBCGrantAllowance {
	return &MsgIBCGrantAllowance{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_IBC_GRANT_ALLOWANCE,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgIBCGrantAllowance) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgIBCGrantAllowance) String() string {
	return render.Render(event)
}

func DecodeMsgIBCGrantAllowance(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgIBCGrantAllowance
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
