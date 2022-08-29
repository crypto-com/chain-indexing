package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_SUBMIT_PARAM_CHANGE_PROPOSAL = "/cosmos.params.v1beta1.ParameterChangeProposal"
const MSG_SUBMIT_PARAM_CHANGE_PROPOSAL_CREATED = "/cosmos.params.v1beta1.ParameterChangeProposal.Created"
const MSG_SUBMIT_PARAM_CHANGE_PROPOSAL_FAILED = "/cosmos.params.v1beta1.ParameterChangeProposal.Failed"

type MsgSubmitParamChangeProposal struct {
	MsgBase

	model.MsgSubmitParamChangeProposalParams
}

func NewMsgSubmitParamChangeProposal(
	msgCommonParams MsgCommonParams,
	params model.MsgSubmitParamChangeProposalParams,
) *MsgSubmitParamChangeProposal {
	return &MsgSubmitParamChangeProposal{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_SUBMIT_PARAM_CHANGE_PROPOSAL,
			Version:         1,

			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

func (event *MsgSubmitParamChangeProposal) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgSubmitParamChangeProposal) String() string {
	return render.Render(event)
}

func DecodeMsgSubmitParamChangeProposal(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgSubmitParamChangeProposal
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
