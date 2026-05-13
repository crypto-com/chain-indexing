package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_TIER_REDELEGATE = "/chainmain.tieredrewards.v1.MsgTierRedelegate"
const MSG_TIER_REDELEGATE_CREATED = "/chainmain.tieredrewards.v1.MsgTierRedelegate.Created"
const MSG_TIER_REDELEGATE_FAILED = "/chainmain.tieredrewards.v1.MsgTierRedelegate.Failed"

type MsgTierRedelegate struct {
	MsgBase
	Owner          string `json:"owner"`
	PositionId     string `json:"positionId"`
	DstValidator   string `json:"dstValidator"`
	CompletionTime string `json:"completionTime"`
	UnbondingId    string `json:"unbondingId"`
}

func NewMsgTierRedelegate(msgCommonParams MsgCommonParams, params model.MsgTierRedelegateParams) *MsgTierRedelegate {
	return &MsgTierRedelegate{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_TIER_REDELEGATE,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),
		params.Owner,
		params.PositionId,
		params.DstValidator,
		params.CompletionTime,
		params.UnbondingId,
	}
}

func (event *MsgTierRedelegate) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

func (event *MsgTierRedelegate) String() string {
	return render.Render(event)
}

func DecodeMsgTierRedelegate(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()
	var event *MsgTierRedelegate
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}
	return event, nil
}
