package event

import (
	"bytes"

	model_gov_v1 "github.com/crypto-com/chain-indexing/usecase/model/gov/v1"

	"github.com/crypto-com/chain-indexing/usecase/coin"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_DEPOSIT_V1 = "/cosmos.gov.v1.MsgDeposit"
const MSG_DEPOSIT_V1_CREATED = "/cosmos.gov.v1.MsgDeposit.Created"
const MSG_DEPOSIT_V1_FAILED = "/cosmos.gov.v1.MsgDeposit.Failed"

type MsgDepositV1 struct {
	MsgBase

	ProposalId string     `json:"proposalId"`
	Depositor  string     `json:"depositor"`
	Amount     coin.Coins `json:"amount"`
}

func NewMsgDepositV1(msgCommonParams MsgCommonParams, params model_gov_v1.MsgDepositParams) *MsgDepositV1 {
	return &MsgDepositV1{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_DEPOSIT_V1,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params.ProposalId,
		params.Depositor,
		params.Amount,
	}
}

func (event *MsgDepositV1) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgDepositV1) String() string {
	return render.Render(event)
}

func DecodeMsgDepositV1(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgDepositV1
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
