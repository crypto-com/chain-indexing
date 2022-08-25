package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_GRANT_ALLOWANCE = "/cosmos.feegrant.v1beta1.MsgGrantAllowance"
const MSG_GRANT_ALLOWANCE_CREATED = "MsgGrantAllowanceCreated"
const MSG_GRANT_ALLOWANCE_FAILED = "MsgGrantAllowanceFailed"

type MsgGrantAllowance struct {
	MsgBase

	Params model.MsgGrantAllowanceParams `json:"params"`
}

func NewMsgGrantAllowance(
	msgCommonParams MsgCommonParams,
	params model.MsgGrantAllowanceParams,
) *MsgGrantAllowance {
	return &MsgGrantAllowance{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_GRANT_ALLOWANCE,
			MsgSuccess:      MSG_GRANT_ALLOWANCE_CREATED,
			MsgFailed:       MSG_GRANT_ALLOWANCE_FAILED,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgGrantAllowance) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgGrantAllowance) String() string {
	return render.Render(event)
}

func DecodeMsgGrantAllowance(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgGrantAllowance
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
