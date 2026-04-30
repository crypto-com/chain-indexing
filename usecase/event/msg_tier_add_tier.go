package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

const MSG_TIER_ADD_TIER = "/chainmain.tieredrewards.v1.MsgAddTier"
const MSG_TIER_ADD_TIER_CREATED = "/chainmain.tieredrewards.v1.MsgAddTier.Created"
const MSG_TIER_ADD_TIER_FAILED = "/chainmain.tieredrewards.v1.MsgAddTier.Failed"

type MsgTierAddTier struct {
	MsgBase
	Authority string `json:"authority"`
	Tier      string `json:"tier"`
}

func NewMsgTierAddTier(msgCommonParams MsgCommonParams, params model.MsgAddTierParams) *MsgTierAddTier {
	return &MsgTierAddTier{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_TIER_ADD_TIER,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),
		params.Authority,
		params.Tier,
	}
}

func (event *MsgTierAddTier) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

func (event *MsgTierAddTier) String() string {
	return render.Render(event)
}

func DecodeMsgTierAddTier(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()
	var event *MsgTierAddTier
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}
	return event, nil
}
