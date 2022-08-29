package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_SUBMIT_CANCEL_SOFTWARE_UPGRADE_PROPOSAL = "/cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal"
const MSG_SUBMIT_CANCEL_SOFTWARE_UPGRADE_PROPOSAL_CREATED = "/cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal.Created"
const MSG_SUBMIT_CANCEL_SOFTWARE_UPGRADE_PROPOSAL_FAILED = "/cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal.Failed"

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
			MsgName:         MSG_SUBMIT_CANCEL_SOFTWARE_UPGRADE_PROPOSAL,
			Version:         1,
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
