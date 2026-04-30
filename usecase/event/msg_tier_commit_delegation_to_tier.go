package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

const MSG_TIER_COMMIT_DELEGATION_TO_TIER = "/chainmain.tieredrewards.v1.MsgCommitDelegationToTier"
const MSG_TIER_COMMIT_DELEGATION_TO_TIER_CREATED = "/chainmain.tieredrewards.v1.MsgCommitDelegationToTier.Created"
const MSG_TIER_COMMIT_DELEGATION_TO_TIER_FAILED = "/chainmain.tieredrewards.v1.MsgCommitDelegationToTier.Failed"

type MsgTierCommitDelegationToTier struct {
	MsgBase
	DelegatorAddress       string `json:"delegatorAddress"`
	ValidatorAddress       string `json:"validatorAddress"`
	Amount                 string `json:"amount"`
	Id                     uint32 `json:"id"`
	TriggerExitImmediately bool   `json:"triggerExitImmediately"`
}

func NewMsgTierCommitDelegationToTier(msgCommonParams MsgCommonParams, params model.MsgCommitDelegationToTierParams) *MsgTierCommitDelegationToTier {
	return &MsgTierCommitDelegationToTier{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_TIER_COMMIT_DELEGATION_TO_TIER,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),
		params.DelegatorAddress,
		params.ValidatorAddress,
		params.Amount,
		params.Id,
		params.TriggerExitImmediately,
	}
}

func (event *MsgTierCommitDelegationToTier) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

func (event *MsgTierCommitDelegationToTier) String() string {
	return render.Render(event)
}

func DecodeMsgTierCommitDelegationToTier(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()
	var event *MsgTierCommitDelegationToTier
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}
	return event, nil
}
