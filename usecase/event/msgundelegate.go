package event

import (
	"bytes"

	"github.com/crypto-com/chainindex/internal/utctime"

	"github.com/crypto-com/chainindex/usecase/coin"
	"github.com/crypto-com/chainindex/usecase/model"

	entity_event "github.com/crypto-com/chainindex/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_UNDELEGATE = "MsgUndelegate"
const MSG_UNDELEGATE_CREATED = "MsgUndelegateCreated"
const MSG_UNDELEGATE_FAILED = "MsgUndelegateFailed"

// MsgUndelegate defines a Cosmos SDK message for performing a undelegation of coins
// from a validator to a delegator.
type MsgUndelegate struct {
	MsgBase

	DelegatorAddress string          `json:"delegatorAddress"`
	ValidatorAddress string          `json:"validatorAddress"`
	Amount           coin.Coin       `json:"amount"`
	UnbondCompleteAt utctime.UTCTime `json:"unbondCompleteAt"`
}

// NewMsgUndelegate creates a new instance of MsgUndelegate
func NewMsgUndelegate(msgCommonParams MsgCommonParams, params model.MsgUndelegateParams) *MsgUndelegate {
	return &MsgUndelegate{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_UNDELEGATE,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params.DelegatorAddress,
		params.ValidatorAddress,
		params.Amount,
		params.UnbondCompleteAt,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgUndelegate) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgUndelegate) String() string {
	return render.Render(event)
}

// DecodeMsgUndelegate decodes the event from encoded bytes
func DecodeMsgUndelegate(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgUndelegate
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
