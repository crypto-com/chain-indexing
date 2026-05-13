package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_TIER_CLAIM_TIER_REWARDS = "/chainmain.tieredrewards.v1.MsgClaimTierRewards"
const MSG_TIER_CLAIM_TIER_REWARDS_CREATED = "/chainmain.tieredrewards.v1.MsgClaimTierRewards.Created"
const MSG_TIER_CLAIM_TIER_REWARDS_FAILED = "/chainmain.tieredrewards.v1.MsgClaimTierRewards.Failed"

type MsgTierClaimTierRewards struct {
	MsgBase
	Owner        string   `json:"owner"`
	PositionIds  []string `json:"positionIds"`
	BaseRewards  string   `json:"baseRewards"`
	BonusRewards string   `json:"bonusRewards"`
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
		params.BaseRewards,
		params.BonusRewards,
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
