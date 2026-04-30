package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_TIER_ADD_TO_TIER_POSITION = "/chainmain.tieredrewards.v1.MsgAddToTierPosition"
const MSG_TIER_ADD_TO_TIER_POSITION_CREATED = "/chainmain.tieredrewards.v1.MsgAddToTierPosition.Created"
const MSG_TIER_ADD_TO_TIER_POSITION_FAILED = "/chainmain.tieredrewards.v1.MsgAddToTierPosition.Failed"

type MsgTierAddToTierPosition struct {
	MsgBase
	Owner      string `json:"owner"`
	PositionId string `json:"positionId"`
	Amount     string `json:"amount"`
}

func NewMsgTierAddToTierPosition(msgCommonParams MsgCommonParams, params model.MsgAddToTierPositionParams) *MsgTierAddToTierPosition {
	return &MsgTierAddToTierPosition{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_TIER_ADD_TO_TIER_POSITION,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),
		params.Owner,
		params.PositionId,
		params.Amount,
	}
}

func (event *MsgTierAddToTierPosition) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

func (event *MsgTierAddToTierPosition) String() string {
	return render.Render(event)
}

func DecodeMsgTierAddToTierPosition(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()
	var event *MsgTierAddToTierPosition
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}
	return event, nil
}
