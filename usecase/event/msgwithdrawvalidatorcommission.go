package event

import (
	"bytes"

	"github.com/crypto-com/chainindex/usecase/model"

	entity_event "github.com/crypto-com/chainindex/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_WITHDRAW_VALIDATOR_COMMISSION = "MsgWithdrawValidatorCommission"
const MSG_WITHDRAW_VALIDATOR_COMMISSION_CREATED = "MsgWithdrawValidatorCommissionCreated"
const MSG_WITHDRAW_VALIDATOR_COMMISSION_FAILED = "MsgWithdrawValidatorCommissionFailed"

type MsgWithdrawValidatorCommission struct {
	MsgBase

	model.MsgWithdrawValidatorCommissionParams
}

func NewMsgWithdrawValidatorCommission(
	msgCommonParams MsgCommonParams,
	params model.MsgWithdrawValidatorCommissionParams,
) *MsgWithdrawValidatorCommission {
	return &MsgWithdrawValidatorCommission{
		NewMsgBase(MsgBaseParams{
			MsgName: MSG_WITHDRAW_DELEGATOR_REWARD,
			Version: 1,

			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

func (event *MsgWithdrawValidatorCommission) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgWithdrawValidatorCommission) String() string {
	return render.Render(event)
}

func DecodeMsgWithdrawValidatorCommission(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgWithdrawValidatorCommission
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
