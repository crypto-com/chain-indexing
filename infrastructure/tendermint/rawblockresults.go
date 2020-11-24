package tendermint

type RawBlockResults struct {
	Height                string                               `json:"height"`
	TxsResults            []RawBlockResultsTxsResult           `json:"txs_results"`
	BeginBlockEvents      []RawBlockResultsEvent               `json:"begin_block_events"`
	EndBlockEvents        []RawBlockResultsEvent               `json:"end_block_events"`
	ValidatorUpdates      []RawBlockResultsValidatorUpdate     `json:"validator_updates"`
	ConsensusParamUpdates RawBlockResultsConsensusParamUpdates `json:"consensus_param_updates"`
}

type RawBlockResultsTxsResult struct {
	Code      int                    `json:"code"`
	Data      string                 `json:"data"`
	Log       string                 `json:"log"`
	Info      string                 `json:"info"`
	GasWanted string                 `json:"gas_wanted"`
	GasUsed   string                 `json:"gas_used"`
	Events    []RawBlockResultsEvent `json:"events"`
	Codespace string                 `json:"codespace"`
}

type RawBlockResultsTxsResultLog struct {
	MaybeMsgIndex *int                   `json:"msg_index"`
	Events        []RawBlockResultsEvent `json:"events"`
}

type RawBlockResultsEvent struct {
	Type       string                          `json:"type"`
	Attributes []RawBlockResultsEventAttribute `json:"attributes"`
}

type RawBlockResultsEventAttribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Index bool   `json:"index"`
}

type RawBlockResultsValidatorUpdate struct {
	PubKey struct {
		Sum struct {
			Type  string `json:"type"`
			Value struct {
				Ed25519 string `json:"ed25519"`
			} `json:"value"`
		} `json:"Sum"`
	} `json:"pub_key"`
	MaybePower string `json:"power,omitempty"`
}

type RawBlockResultsConsensusParamUpdates struct {
	Block struct {
		MaxBytes string `json:"max_bytes"`
		MaxGas   string `json:"max_gas"`
	} `json:"block"`
	Evidence struct {
		MaxAgeNumBlocks string `json:"max_age_num_blocks"`
		MaxAgeDuration  string `json:"max_age_duration"`
		MaxBytes        string `json:"max_bytes"`
	} `json:"evidence"`
	Validator struct {
		PubKeyTypes []string `json:"pub_key_types"`
	} `json:"validator"`
}
