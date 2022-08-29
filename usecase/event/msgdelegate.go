package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/crypto-com/chain-indexing/usecase/model"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_DELEGATE = "/cosmos.staking.v1beta1.MsgDelegate"
const MSG_DELEGATE_CREATED = "/cosmos.staking.v1beta1.MsgDelegate.Created"
const MSG_DELEGATE_FAILED = "/cosmos.staking.v1beta1.MsgDelegate.Failed"

// MsgDelegate defines a Cosmos SDK message for performing a delegation of coins
// from a delegator to a validator.
type MsgDelegate struct {
	MsgBase

	DelegatorAddress   string    `json:"delegatorAddress"`
	ValidatorAddress   string    `json:"validatorAddress"`
	Amount             coin.Coin `json:"amount"`
	AutoClaimedRewards coin.Coin `json:"autoClaimedRewards"`
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
		params.AutoClaimedRewards,
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
