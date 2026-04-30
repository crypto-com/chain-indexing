package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

const MSG_TIER_CLAIM_TIER_REWARDS = "/chainmain.tieredrewards.v1.MsgClaimTierRewards"
const MSG_TIER_CLAIM_TIER_REWARDS_CREATED = "/chainmain.tieredrewards.v1.MsgClaimTierRewards.Created"
const MSG_TIER_CLAIM_TIER_REWARDS_FAILED = "/chainmain.tieredrewards.v1.MsgClaimTierRewards.Failed"

type MsgTierClaimTierRewards struct {
	MsgBase
	Owner       string   `json:"owner"`
	PositionIds []string `json:"position_ids"`
}

func NewMsgTierClaimTierRewards(msgCommonParams MsgCommonParams, params model.MsgClaimTierRewardsParams) *MsgTierClaimTierRewards {
	return &MsgTierClaimTierRewards{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_TIER_CLAIM_TIER_REWARDS,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),
		params.Owner,
		params.PositionIds,
	}
}

func (event *MsgTierClaimTierRewards) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

func (event *MsgTierClaimTierRewards) String() string {
	return render.Render(event)
}

func DecodeMsgTierClaimTierRewards(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()
	var event *MsgTierClaimTierRewards
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}
	return event, nil
}
