package model

import "math/big"

type BlockResults struct {
	Height                int64
	TxsResults            []BlockResultsTxsResult
	BeginBlockEvents      []BlockResultsEvent
	EndBlockEvents        []BlockResultsEvent
	ValidatorUpdates      []BlockResultsValidatorUpdate
	ConsensusParamUpdates BlockResultsConsensusParamUpdates
}

type BlockResultsTxsResult struct {
	Code      int                        `json:"code"`
	Data      string                     `json:"data"`
	Log       []BlockResultsTxsResultLog `json:"log"`
	Info      string                     `json:"info"`
	GasWanted string                     `json:"gasWanted"`
	GasUsed   string                     `json:"gasUsed"`
	Events    []BlockResultsEvent        `json:"events"`
	Codespace string                     `json:"codespace"`
}

type BlockResultsTxsResultLog struct {
	MsgIndex int                 `json:"msgIndex"`
	Events   []BlockResultsEvent `json:"events"`
}

type BlockResultsEvent struct {
	Type       string                       `json:"type"`
	Attributes []BlockResultsEventAttribute `json:"attributes"`
}

type BlockResultsEventAttribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type BlockResultsValidatorUpdate struct {
	PubKey BlockResultsValidatorPubKey
	Power  *big.Int `json:"power"`
}

type BlockResultsValidatorPubKey struct {
	Type    string `json:"type"`
	PubKey  string `json:"pubKey"`
	Address string `json:"address"`
}

type BlockResultsConsensusParamUpdates struct {
	Block     BlockResultsConsensusParamUpdatesBlock      `json:"block"`
	Evidence  BlockResultsConsensusParamUpdatesEvidence   `json:"evidence"`
	Validator BlockResultsConsensusParamsUpdatesValidator `json:"validator"`
}

type BlockResultsConsensusParamUpdatesBlock struct {
	MaxBytes string `json:"maxBytes"`
	MaxGas   string `json:"maxGas"`
}

type BlockResultsConsensusParamUpdatesEvidence struct {
	MaxAgeNumBlocks string `json:"maxAgeNumBlocks"`
	MaxAgeDuration  string `json:"maxAgeDuration"`
}

type BlockResultsConsensusParamsUpdatesValidator struct {
	PubKeyTypes []string `json:"pubKeyTypes"`
}
