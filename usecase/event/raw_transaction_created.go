package event

import (
	"bytes"

	"github.com/luci/go-render/render"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/crypto-com/chain-indexing/usecase/model"
	jsoniter "github.com/json-iterator/go"
)

const RAW_TRANSACTION_CREATED = "RawTransactionCreated"

type RawTransactionCreated struct {
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

const RAW_TRANSACTION_SIGNER_SECP256K1 = "/cosmos.crypto.secp256k1.PubKey"
const RAW_TRANSACTION_SIGNER_MULTISIG_LEGACY_AMINO = "/cosmos.crypto.multisig.LegacyAminoPubKey"

func NewRawTransactionCreated(blockHeight int64, params model.CreateRawTransactionParams) *RawTransactionCreated {
	return &RawTransactionCreated{
		Base: entity_event.NewBase(entity_event.BaseParams{
			Name:        RAW_TRANSACTION_CREATED,
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

func (event *RawTransactionCreated) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *RawTransactionCreated) String() string {
	return render.Render(event)
}

func DecodeRawTransactionCreated(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *RawTransactionCreated
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
