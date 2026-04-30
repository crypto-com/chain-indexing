package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_TIER_EXIT_TIER_WITH_DELEGATION = "/chainmain.tieredrewards.v1.MsgExitTierWithDelegation"
const MSG_TIER_EXIT_TIER_WITH_DELEGATION_CREATED = "/chainmain.tieredrewards.v1.MsgExitTierWithDelegation.Created"
const MSG_TIER_EXIT_TIER_WITH_DELEGATION_FAILED = "/chainmain.tieredrewards.v1.MsgExitTierWithDelegation.Failed"

type MsgTierExitTierWithDelegation struct {
	MsgBase
	Owner      string `json:"owner"`
	PositionId string `json:"positionId"`
	Amount     string `json:"amount"`
}

func NewMsgTierExitTierWithDelegation(msgCommonParams MsgCommonParams, params model.MsgExitTierWithDelegationParams) *MsgTierExitTierWithDelegation {
	return &MsgTierExitTierWithDelegation{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_TIER_EXIT_TIER_WITH_DELEGATION,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),
		params.Owner,
		params.PositionId,
		params.Amount,
	}
}

func (event *MsgTierExitTierWithDelegation) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

func (event *MsgTierExitTierWithDelegation) String() string {
	return render.Render(event)
}

func DecodeMsgTierExitTierWithDelegation(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()
	var event *MsgTierExitTierWithDelegation
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}
	return event, nil
}
