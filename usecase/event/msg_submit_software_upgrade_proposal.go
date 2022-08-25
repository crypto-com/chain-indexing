package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_SUBMIT_SOFTWARE_UPGRADE_PROPOSAL = "/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal"
const MSG_SUBMIT_SOFTWARE_UPGRADE_PROPOSAL_CREATED = "MsgSubmitSoftwareUpgradeProposalCreated"
const MSG_SUBMIT_SOFTWARE_UPGRADE_PROPOSAL_FAILED = "MsgSubmitSoftwareUpgradeProposalFailed"

type MsgSubmitSoftwareUpgradeProposal struct {
	MsgBase

	model.MsgSubmitSoftwareUpgradeProposalParams
}

func NewMsgSubmitSoftwareUpgradeProposal(
	msgCommonParams MsgCommonParams,
	params model.MsgSubmitSoftwareUpgradeProposalParams,
) *MsgSubmitSoftwareUpgradeProposal {
	return &MsgSubmitSoftwareUpgradeProposal{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_SUBMIT_SOFTWARE_UPGRADE_PROPOSAL,
			MsgSuccess:      MSG_SUBMIT_SOFTWARE_UPGRADE_PROPOSAL_CREATED,
			MsgFailed:       MSG_SUBMIT_SOFTWARE_UPGRADE_PROPOSAL_FAILED,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

func (event *MsgSubmitSoftwareUpgradeProposal) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgSubmitSoftwareUpgradeProposal) String() string {
	return render.Render(event)
}

func DecodeMsgSubmitSoftwareUpgradeProposal(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgSubmitSoftwareUpgradeProposal
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
