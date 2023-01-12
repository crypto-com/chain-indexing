package event

import (
	"bytes"

	model "github.com/crypto-com/chain-indexing/usecase/model/v1"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_SUBMIT_PROPOSAL = "/cosmos.gov.v1.MsgSubmitProposal"
const MSG_SUBMIT_PROPOSAL_CREATED = "/cosmos.gov.v1.MsgSubmitProposal.Created"
const MSG_SUBMIT_PROPOSAL_FAILED = "/cosmos.gov.v1.MsgSubmitProposal.Failed"

type MsgSubmitProposal struct {
	MsgBase

	model.MsgSubmitProposalParams
}

func NewMsgSubmitProposal(
	msgCommonParams MsgCommonParams,
	params model.MsgSubmitProposalParams,
) *MsgSubmitProposal {
	return &MsgSubmitProposal{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_SUBMIT_PROPOSAL,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

func (event *MsgSubmitProposal) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgSubmitProposal) String() string {
	return render.Render(event)
}

func DecodeMsgSubmitProposal(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgSubmitProposal
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
