package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

const MSG_TIER_LOCK_TIER = "/chainmain.tieredrewards.v1.MsgLockTier"
const MSG_TIER_LOCK_TIER_CREATED = "/chainmain.tieredrewards.v1.MsgLockTier.Created"
const MSG_TIER_LOCK_TIER_FAILED = "/chainmain.tieredrewards.v1.MsgLockTier.Failed"

type MsgTierLockTier struct {
	MsgBase
	Owner                  string `json:"owner"`
	Id                     uint32 `json:"id"`
	Amount                 string `json:"amount"`
	ValidatorAddress       string `json:"validator_address"`
	TriggerExitImmediately bool   `json:"trigger_exit_immediately"`
}

func NewMsgTierLockTier(msgCommonParams MsgCommonParams, params model.MsgLockTierParams) *MsgTierLockTier {
	return &MsgTierLockTier{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_TIER_LOCK_TIER,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),
		params.Owner,
		params.Id,
		params.Amount,
		params.ValidatorAddress,
		params.TriggerExitImmediately,
	}
}

func (event *MsgTierLockTier) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

func (event *MsgTierLockTier) String() string {
	return render.Render(event)
}

func DecodeMsgTierLockTier(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()
	var event *MsgTierLockTier
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}
	return event, nil
}
