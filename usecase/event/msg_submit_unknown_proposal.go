package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_SUBMIT_UNKNOWN_PROPOSAL = "MsgSubmitUnknownProposal"
const MSG_SUBMIT_UNKNOWN_PROPOSAL_CREATED = "MsgSubmitUnknownProposal.Created"
const MSG_SUBMIT_UNKNOWN_PROPOSAL_FAILED = "MsgSubmitUnknownProposal.Failed"

type MsgSubmitUnknownProposal struct {
	MsgBase

	model.MsgSubmitUnknownProposalParams
}

func NewMsgSubmitUnknownProposal(
	msgCommonParams MsgCommonParams,
	params model.MsgSubmitUnknownProposalParams,
) *MsgSubmitUnknownProposal {
	return &MsgSubmitUnknownProposal{
		NewMsgBase(MsgBaseParams{
			MsgName: MSG_SUBMIT_UNKNOWN_PROPOSAL,
			Version: 1,

			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

func (event *MsgSubmitUnknownProposal) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgSubmitUnknownProposal) String() string {
	return render.Render(event)
}

func DecodeMsgSubmitUnknownProposal(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgSubmitUnknownProposal
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
