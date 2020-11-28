package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	"github.com/crypto-com/chain-indexing/usecase/coin"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_DEPOSIT = "MsgDeposit"
const MSG_DEPOSIT_CREATED = "MsgDepositCreated"
const MSG_DEPOSIT_FAILED = "MsgDepositFailed"

type MsgDeposit struct {
	MsgBase

	ProposalId string    `json:"proposalId"`
	Depositor  string    `json:"depositor"`
	Amount     coin.Coin `json:"amount"`
}

func NewMsgDeposit(msgCommonParams MsgCommonParams, params model.MsgDepositParams) *MsgDeposit {
	return &MsgDeposit{
		NewMsgBase(MsgBaseParams{
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
