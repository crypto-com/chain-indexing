package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_SUBMIT_COMMUNITY_POOL_SPEND_PROPOSAL = "/cosmos.distribution.v1beta1.CommunityPoolSpendProposal"
const MSG_SUBMIT_COMMUNITY_POOL_SPEND_PROPOSAL_CREATED = "MsgSubmitCommunityPoolSpendProposalCreated"
const MSG_SUBMIT_COMMUNITY_POOL_SPEND_PROPOSAL_FAILED = "MsgSubmitCommunityPoolSpendProposalFailed"

type MsgSubmitCommunityPoolSpendProposal struct {
	MsgBase

	model.MsgSubmitCommunityPoolSpendProposalParams
}

func NewMsgSubmitCommunityPoolSpendProposal(
	msgCommonParams MsgCommonParams,
	params model.MsgSubmitCommunityPoolSpendProposalParams,
) *MsgSubmitCommunityPoolSpendProposal {
	return &MsgSubmitCommunityPoolSpendProposal{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_SUBMIT_COMMUNITY_POOL_SPEND_PROPOSAL,
			MsgSuccess:      MSG_SUBMIT_COMMUNITY_POOL_SPEND_PROPOSAL_CREATED,
			MsgFailed:       MSG_SUBMIT_COMMUNITY_POOL_SPEND_PROPOSAL_FAILED,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

func (event *MsgSubmitCommunityPoolSpendProposal) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgSubmitCommunityPoolSpendProposal) String() string {
	return render.Render(event)
}

func DecodeMsgSubmitCommunityPoolSpendProposal(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgSubmitCommunityPoolSpendProposal
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
