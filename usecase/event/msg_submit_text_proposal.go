package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_SUBMIT_TEXT_PROPOSAL = "/cosmos.gov.v1beta1.TextProposal"
const MSG_SUBMIT_TEXT_PROPOSAL_CREATED = "/cosmos.gov.v1beta1.TextProposal.Created"
const MSG_SUBMIT_TEXT_PROPOSAL_FAILED = "/cosmos.gov.v1beta1.TextProposal.Failed"

type MsgSubmitTextProposal struct {
	MsgBase

	model.MsgSubmitTextProposalParams
}

func NewMsgSubmitTextProposal(
	msgCommonParams MsgCommonParams,
	params model.MsgSubmitTextProposalParams,
) *MsgSubmitTextProposal {
	return &MsgSubmitTextProposal{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_SUBMIT_TEXT_PROPOSAL,
			Version:         1,

			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

func (event *MsgSubmitTextProposal) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgSubmitTextProposal) String() string {
	return render.Render(event)
}

func DecodeMsgSubmitTextProposal(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgSubmitTextProposal
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
