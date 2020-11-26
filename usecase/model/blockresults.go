package model

import (
	"math/big"

	"github.com/luci/go-render/render"
)

type BlockResults struct {
	Height                int64                             `json:"height"`
	TxsResults            []BlockResultsTxsResult           `json:"txsResults"`
	BeginBlockEvents      []BlockResultsEvent               `json:"beginBlockEvents"`
	EndBlockEvents        []BlockResultsEvent               `json:"endBlockEvents"`
	ValidatorUpdates      []BlockResultsValidatorUpdate     `json:"validatorUpdates"`
	ConsensusParamUpdates BlockResultsConsensusParamUpdates `json:"consensusParamUpdates"`
}

func (results *BlockResults) String() string {
	return render.Render(results)
}

type BlockResultsTxsResult struct {
	Code      int                        `json:"code"`
	Data      string                     `json:"data"`
	Log       []BlockResultsTxsResultLog `json:"log"`
	RawLog    string                     `json:"rawLog"`
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
	Pubkey     BlockResultsValidatorPubKey `json:"pubkey"`
	Address    string                      `json:"address"`
	MaybePower *big.Int                    `json:"power"`
}

type BlockResultsValidatorPubKey struct {
	Type   string `json:"type"`
	Pubkey string `json:"pubkey"`
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
	MaxBytes        string `json:"maxBytes"`
}

type BlockResultsConsensusParamsUpdatesValidator struct {
	PubKeyTypes []string `json:"pubKeyTypes"`
}
