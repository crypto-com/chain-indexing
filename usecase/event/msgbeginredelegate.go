package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/crypto-com/chain-indexing/usecase/model"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_BEGIN_REDELEGATE = "/cosmos.staking.v1beta1.MsgBeginRedelegate"
const MSG_BEGIN_REDELEGATE_CREATED = "MsgBeginRedelegateCreated"
const MSG_BEGIN_REDELEGATE_FAILED = "MsgBeginRedelegateFailed"

// MsgBeginRedelegate defines a Cosmos SDK message for performing a redelegation
// of coins from a delegator and source validator to a destination validator.
type MsgBeginRedelegate struct {
	MsgBase

	DelegatorAddress    string    `json:"delegatorAddress"`
	ValidatorSrcAddress string    `json:"validatorSrcAddress"`
	ValidatorDstAddress string    `json:"validatorDstAddress"`
	Amount              coin.Coin `json:"amount"`
	AutoClaimedRewards  coin.Coin `json:"autoClaimedRewards"`
}

// NewMsgBeginRedelegate creates a new instance of MsgBeginRedelegate
func NewMsgBeginRedelegate(msgCommonParams MsgCommonParams, params model.MsgBeginRedelegateParams) *MsgBeginRedelegate {
	return &MsgBeginRedelegate{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_BEGIN_REDELEGATE,
			MsgSuccess:      MSG_BEGIN_REDELEGATE_CREATED,
			MsgFailed:       MSG_BEGIN_REDELEGATE_FAILED,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params.DelegatorAddress,
		params.ValidatorSrcAddress,
		params.ValidatorDstAddress,
		params.Amount,
		params.AutoClaimedRewards,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgBeginRedelegate) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgBeginRedelegate) String() string {
	return render.Render(event)
}

// DecodeMsgBeginRedelegate decodes the event from encoded bytes
func DecodeMsgBeginRedelegate(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgBeginRedelegate
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
