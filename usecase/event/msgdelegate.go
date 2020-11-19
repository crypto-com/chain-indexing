package event

import (
	"bytes"

	"github.com/crypto-com/chainindex/usecase/coin"
	"github.com/crypto-com/chainindex/usecase/model"

	entity_event "github.com/crypto-com/chainindex/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_DELEGATE = "MsgDelegate"
const MSG_DELEGATE_CREATED = "MsgDelegateCreated"
const MSG_DELEGATE_FAILED = "MsgDelegateFailed"

// MsgDelegate defines a Cosmos SDK message for performing a delegation of coins
// from a delegator to a validator.
type MsgDelegate struct {
	MsgBase

	DelegatorAddress string    `json:"delegatorAddress"`
	ValidatorAddress string    `json:"validatorAddress"`
	Amount           coin.Coin `json:"amount"`
}

// NewMsgDelegate creates a new instance of MsgDelegate
func NewMsgDelegate(msgCommonParams MsgCommonParams, params model.MsgDelegateParams) *MsgDelegate {
	return &MsgDelegate{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_DELEGATE,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params.DelegatorAddress,
		params.ValidatorAddress,
		params.Amount,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgDelegate) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgDelegate) String() string {
	return render.Render(event)
}

// DecodeMsgDelegate decodes the event from encoded bytes
func DecodeMsgDelegate(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgDelegate
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
