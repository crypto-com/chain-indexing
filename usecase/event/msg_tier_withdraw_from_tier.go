package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_TIER_WITHDRAW_FROM_TIER = "/chainmain.tieredrewards.v1.MsgWithdrawFromTier"
const MSG_TIER_WITHDRAW_FROM_TIER_CREATED = "/chainmain.tieredrewards.v1.MsgWithdrawFromTier.Created"
const MSG_TIER_WITHDRAW_FROM_TIER_FAILED = "/chainmain.tieredrewards.v1.MsgWithdrawFromTier.Failed"

type MsgTierWithdrawFromTier struct {
	MsgBase
	Owner      string `json:"owner"`
	PositionId string `json:"positionId"`
}

func NewMsgTierWithdrawFromTier(msgCommonParams MsgCommonParams, params model.MsgWithdrawFromTierParams) *MsgTierWithdrawFromTier {
	return &MsgTierWithdrawFromTier{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_TIER_WITHDRAW_FROM_TIER,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),
		params.Owner,
		params.PositionId,
	}
}

func (event *MsgTierWithdrawFromTier) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

func (event *MsgTierWithdrawFromTier) String() string {
	return render.Render(event)
}

func DecodeMsgTierWithdrawFromTier(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()
	var event *MsgTierWithdrawFromTier
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}
	return event, nil
}
