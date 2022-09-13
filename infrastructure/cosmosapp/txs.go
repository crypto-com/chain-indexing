package cosmosapp

import (
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type TxsResp struct {
	Tx         model.CosmosTx `json:"tx"`
	TxResponse TxResponse     `json:"tx_response"`
}

type TxResponse struct {
	Height    string             `json:"height,omitempty"`
	TxHash    string             `json:"txhash,omitempty"`
	Codespace string             `json:"codespace,omitempty"`
	Code      uint32             `json:"code,omitempty"`
	Data      string             `json:"data,omitempty"`
	RawLog    string             `json:"raw_log,omitempty"`
	Logs      []TxResponseLog    `json:"logs"`
	Info      string             `json:"info,omitempty"`
	GasWanted string             `json:"gas_wanted,omitempty"`
	GasUsed   string             `json:"gas_used,omitempty"`
	Tx        TxResponseTx       `json:"tx,omitempty"`
	Timestamp string             `json:"timestamp,omitempty"`
	Events    []model.BlockEvent `json:"events"`
}

type TxResponseTx struct {
	Type string `json:"@type"`
	model.CosmosTx
}

type TxResponseLog struct {
	MsgIndex uint32             `json:"msg_index,omitempty"`
	Log      string             `json:"log,omitempty"`
	Events   []model.BlockEvent `json:"events"`
}
