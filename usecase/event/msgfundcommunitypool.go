package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/crypto-com/chain-indexing/usecase/model"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_FUND_COMMUNITY_POOL = "/cosmos.distribution.v1beta1.MsgFundCommunityPool"
const MSG_FUND_COMMUNITY_POOL_CREATED = "/cosmos.distribution.v1beta1.MsgFundCommunityPool.Created"
const MSG_FUND_COMMUNITY_POOL_FAILED = "/cosmos.distribution.v1beta1.MsgFundCommunityPool.Failed"

type MsgFundCommunityPool struct {
	MsgBase

	Depositor string     `json:"depositor"`
	Amount    coin.Coins `json:"amount"`
}

func NewMsgFundCommunityPool(msgCommonParams MsgCommonParams, params model.MsgFundCommunityPoolParams) *MsgFundCommunityPool {
	return &MsgFundCommunityPool{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_FUND_COMMUNITY_POOL,
			Version:         1,
			
			MsgCommonParams: msgCommonParams,
		}),

		params.Depositor,
		params.Amount,
	}
}

func (event *MsgFundCommunityPool) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgFundCommunityPool) String() string {
	return render.Render(event)
}

func DecodeMsgFundCommunityPool(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgFundCommunityPool
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
