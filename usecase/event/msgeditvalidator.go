package event

import (
	"bytes"

	"github.com/crypto-com/chainindex/usecase/model"

	entity_event "github.com/crypto-com/chainindex/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_EDIT_VALIDATOR = "MsgEditValidator"
const MSG_EDIT_VALIDATOR_CREATED = "MsgEditValidatorCreated"
const MSG_EDIT_VALIDATOR_FAILED = "MsgEditValidatorFailed"

type MsgEditValidator struct {
	MsgBase

	ValidatorAddress  string `json:"validatorAddress"`
	CommissionRate    string `json:"commissionRate"`
	MinSelfDelegation string `json:"minSelfDelegation"`
}

func NewMsgEditValidator(msgCommonParams MsgCommonParams, params model.MsgEditValidatorParams) *MsgEditValidator {
	return &MsgEditValidator{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_EDIT_VALIDATOR,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params.ValidatorAddress,
		params.CommissionRate,
		params.MinSelfDelegation,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgEditValidator) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgEditValidator) String() string {
	return render.Render(event)
}

func DecodeMsgEditValidator(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgEditValidator
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
