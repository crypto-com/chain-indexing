package event

import (
	"bytes"

	"github.com/crypto-com/chainindex/usecase/model"

	entity_event "github.com/crypto-com/chainindex/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_SUBMIT_CANCEL_SOFTWARE_UPGRADE_PROPOSAL = "MsgSubmitCancelSoftwareUpgradeProposal"
const MSG_SUBMIT_CANCEL_SOFTWARE_UPGRADE_PROPOSAL_CREATED = "MsgSubmitCancelSoftwareUpgradeProposalCreated"
const MSG_SUBMIT_CANCEL_SOFTWARE_UPGRADE_PROPOSAL_FAILED = "MsgSubmitCancelSoftwareUpgradeProposalFailed"

type MsgSubmitCancelSoftwareUpgradeProposal struct {
	MsgBase

	model.MsgSubmitCancelSoftwareUpgradeProposalParams
}

func NewMsgSubmitCancelSoftwareUpgradeProposal(
	msgCommonParams MsgCommonParams,
	params model.MsgSubmitCancelSoftwareUpgradeProposalParams,
) *MsgSubmitCancelSoftwareUpgradeProposal {
	return &MsgSubmitCancelSoftwareUpgradeProposal{
		NewMsgBase(MsgBaseParams{
			MsgName: MSG_SUBMIT_CANCEL_SOFTWARE_UPGRADE_PROPOSAL,
			Version: 1,

			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

func (event *MsgSubmitCancelSoftwareUpgradeProposal) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgSubmitCancelSoftwareUpgradeProposal) String() string {
	return render.Render(event)
}

func DecodeMsgSubmitCancelSoftwareUpgradeProposal(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgSubmitCancelSoftwareUpgradeProposal
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
