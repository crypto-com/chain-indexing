package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_WITHDRAW_DELEGATOR_REWARD = "MsgWithdrawDelegatorReward"
const MSG_WITHDRAW_DELEGATOR_REWARD_CREATED = "MsgWithdrawDelegatorRewardCreated"
const MSG_WITHDRAW_DELEGATOR_REWARD_FAILED = "MsgWithdrawDelegatorRewardFailed"

type MsgWithdrawDelegatorReward struct {
	MsgBase

	model.MsgWithdrawDelegatorRewardParams
}

func NewMsgWithdrawDelegatorReward(
	msgCommonParams MsgCommonParams,
	params model.MsgWithdrawDelegatorRewardParams,
) *MsgWithdrawDelegatorReward {
	return &MsgWithdrawDelegatorReward{
		NewMsgBase(MsgBaseParams{
			MsgName: MSG_WITHDRAW_DELEGATOR_REWARD,
			Version: 1,

			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

func (event *MsgWithdrawDelegatorReward) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgWithdrawDelegatorReward) String() string {
	return render.Render(event)
}

func DecodeMsgWithdrawDelegatorReward(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgWithdrawDelegatorReward
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
