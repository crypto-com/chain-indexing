package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/crypto-com/chain-indexing/usecase/model"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
)

const RAW_TRANSACTION_FAILED = "RawTransactionFailed"

type RawTransactionFailed struct {
	entity_event.Base

	TxHash        string                    `json:"txHash"`
	Index         int                       `json:"index"`
	Code          int                       `json:"code"`
	Log           string                    `json:"log"`
	Signers       []model.TransactionSigner `json:"signers"`
	Fee           coin.Coins                `json:"fee"`
	FeePayer      string                    `json:"feePayer"`
	FeeGranter    string                    `json:"feeGranter"`
	GasWanted     int                       `json:"gasWanted"`
	GasUsed       int                       `json:"gasUsed"`
	Memo          string                    `json:"memo"`
	TimeoutHeight int64                     `json:"timeoutHeight"`
	Messages      []map[string]interface{}  `json:"messages"`
}

func NewRawTransactionFailed(blockHeight int64, params model.CreateRawTransactionParams) *RawTransactionFailed {
	return &RawTransactionFailed{
		Base: entity_event.NewBase(entity_event.BaseParams{
			Name:        RAW_TRANSACTION_FAILED,
			Version:     1,
			BlockHeight: blockHeight,
		}),

		TxHash:        params.TxHash,
		Index:         params.Index,
		Code:          params.Code,
		Log:           params.Log,
		Signers:       params.Signers,
		Fee:           params.Fee,
		FeePayer:      params.FeePayer,
		FeeGranter:    params.FeeGranter,
		GasWanted:     params.GasWanted,
		GasUsed:       params.GasUsed,
		Memo:          params.Memo,
		TimeoutHeight: params.TimeoutHeight,
		Messages:      params.Messages,
	}
}

func (event *RawTransactionFailed) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *RawTransactionFailed) String() string {
	return render.Render(event)
}

func DecodeRawTransactionFailed(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *RawTransactionFailed
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
