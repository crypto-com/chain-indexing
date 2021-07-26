package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/crypto-com/chain-indexing/usecase/model"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
)

const TRANSACTION_FAILED = "TransactionFailed"

type TransactionFailed struct {
	entity_event.Base

	TxHash   string                    `json:"txHash"`
	Index    int                       `json:"index"`
	Code     int                       `json:"code"`
	Log      string                    `json:"log"`
	MsgCount int                       `json:"msgCount"`
	Signers  []model.TransactionSigner `json:"signers"`
	// Deprecated
	Senders       []model.TransactionSigner `json:"senders"`
	Fee           coin.Coins                `json:"fee"`
	FeePayer      string                    `json:"feePayer"`
	FeeGranter    string                    `json:"feeGranter"`
	GasWanted     int                       `json:"gasWanted"`
	GasUsed       int                       `json:"gasUsed"`
	Memo          string                    `json:"memo"`
	TimeoutHeight int64                     `json:"timeoutHeight"`
}

func NewTransactionFailed(blockHeight int64, params model.CreateTransactionParams) *TransactionFailed {
	return &TransactionFailed{
		Base: entity_event.NewBase(entity_event.BaseParams{
			Name:        TRANSACTION_FAILED,
			Version:     1,
			BlockHeight: blockHeight,
		}),

		TxHash:        params.TxHash,
		Index:         params.Index,
		Code:          params.Code,
		Log:           params.Log,
		MsgCount:      params.MsgCount,
		Signers:       params.Signers,
		Senders:       params.Signers,
		Fee:           params.Fee,
		FeePayer:      params.FeePayer,
		FeeGranter:    params.FeeGranter,
		GasWanted:     params.GasWanted,
		GasUsed:       params.GasUsed,
		Memo:          params.Memo,
		TimeoutHeight: params.TimeoutHeight,
	}
}

func (event *TransactionFailed) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *TransactionFailed) String() string {
	return render.Render(event)
}

func DecodeTransactionFailed(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *TransactionFailed
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
