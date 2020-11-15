package event

import (
	"bytes"

	"github.com/crypto-com/chainindex/usecase/coin"
	"github.com/crypto-com/chainindex/usecase/model"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"

	entity_event "github.com/crypto-com/chainindex/entity/event"
)

const TRANSACTION_FAILED_NAME = "TransactionFailed"

type TransactionFailed struct {
	entity_event.Base

	TxHash    string    `json:"txHash"`
	Code      int       `json:"code"`
	Log       string    `json:"log"`
	MsgCount  int       `json:"msgCount"`
	Fee       coin.Coin `json:"fee"`
	GasWanted string    `json:"gasWanted"`
	GasUsed   string    `json:"gasUsed"`
}

func NewTransactionFailed(blockHeight int64, params model.CreateTransactionParams) *TransactionFailed {
	return &TransactionFailed{
		Base: entity_event.NewBase(entity_event.BaseParams{
			Name:        TRANSACTION_FAILED_NAME,
			Version:     1,
			BlockHeight: blockHeight,
		}),

		TxHash:    params.TxHash,
		Code:      params.Code,
		Log:       params.Log,
		MsgCount:  params.MsgCount,
		Fee:       params.Fee,
		GasWanted: params.GasWanted,
		GasUsed:   params.GasUsed,
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
