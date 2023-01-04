package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/event"
	model "github.com/crypto-com/chain-indexing/usecase/model/v1"

	"github.com/crypto-com/chain-indexing/usecase/coin"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_DEPOSIT = "/cosmos.gov.v1.MsgDeposit"
const MSG_DEPOSIT_CREATED = "/cosmos.gov.v1.MsgDeposit.Created"
const MSG_DEPOSIT_FAILED = "/cosmos.gov.v1.MsgDeposit.Failed"

type MsgDeposit struct {
	event.MsgBase

	ProposalId string     `json:"proposalId"`
	Depositor  string     `json:"depositor"`
	Amount     coin.Coins `json:"amount"`
}

func NewMsgDeposit(msgCommonParams event.MsgCommonParams, params model.MsgDepositParams) *MsgDeposit {
	return &MsgDeposit{
		event.NewMsgBase(event.MsgBaseParams{
			MsgName:         MSG_DEPOSIT,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params.ProposalId,
		params.Depositor,
		params.Amount,
	}
}

func (event *MsgDeposit) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgDeposit) String() string {
	return render.Render(event)
}

func DecodeMsgDeposit(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgDeposit
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
