package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_UNJAIL = "MsgUnjail"
const MSG_UNJAIL_CREATED = "MsgUnjailCreated"
const MSG_UNJAIL_FAILED = "MsgUnjailFailed"

// MsgUnjail defines a Cosmos SDK message for unjailing a jailed validator, thus returning
// them into the bonded validator set, so they can begin receiving provisions
// and rewards again.
type MsgUnjail struct {
	MsgBase

	ValidatorAddr string `json:"validatorAddress"`
}

// NewMsgUnjail creates a new instance of MsgUnjail
func NewMsgUnjail(msgCommonParams MsgCommonParams, params model.MsgUnjailParams) *MsgUnjail {
	return &MsgUnjail{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_UNJAIL,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params.ValidatorAddr,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgUnjail) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgUnjail) String() string {
	return render.Render(event)
}

// DecodeMsgUnjail decodes the event from encoded bytes
func DecodeMsgUnjail(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgUnjail
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
