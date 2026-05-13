package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_TIER_TRIGGER_EXIT_FROM_TIER = "/chainmain.tieredrewards.v1.MsgTriggerExitFromTier"
const MSG_TIER_TRIGGER_EXIT_FROM_TIER_CREATED = "/chainmain.tieredrewards.v1.MsgTriggerExitFromTier.Created"
const MSG_TIER_TRIGGER_EXIT_FROM_TIER_FAILED = "/chainmain.tieredrewards.v1.MsgTriggerExitFromTier.Failed"

type MsgTierTriggerExitFromTier struct {
	MsgBase
	Owner        string `json:"owner"`
	PositionId   string `json:"positionId"`
	ExitUnlockAt string `json:"exitUnlockAt"`
}

func NewMsgTierTriggerExitFromTier(msgCommonParams MsgCommonParams, params model.MsgTriggerExitFromTierParams) *MsgTierTriggerExitFromTier {
	return &MsgTierTriggerExitFromTier{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_TIER_TRIGGER_EXIT_FROM_TIER,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),
		params.Owner,
		params.PositionId,
		params.ExitUnlockAt,
	}
}

func (event *MsgTierTriggerExitFromTier) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

func (event *MsgTierTriggerExitFromTier) String() string {
	return render.Render(event)
}

func DecodeMsgTierTriggerExitFromTier(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()
	var event *MsgTierTriggerExitFromTier
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}
	return event, nil
}
