package model

type RawBlockResults struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Height                string                     `json:"height"`
		TxsEvents             []RawBlockResultsTxResult  `json:"txs_results"`
		BeginBlockEvents      []RawBlockResultsEvent     `json:"begin_block_events"`
		EndBlockEvents        []RawBlockResultsEvent     `json:"end_block_events"`
		ValidatorUpdates      []RawBlockResultsValidator `json:"validator_updates"`
		ConsensusParamUpdates RawConsensusParamUpdates   `json:"consensus_param_updates"`
	} `json:"result"`
}

type RawBlockResultsTxResult struct {
	Code      int                    `json:"code"`
	Data      interface{}            `json:"data"`
	Log       string                 `json:"log"`
	Info      string                 `json:"info"`
	GasWanted string                 `json:"gasWanted"`
	GasUsed   string                 `json:"gasUsed"`
	Events    []RawBlockResultsEvent `json:"events"`
	Codespace string                 `json:"codespace"`
}

type RawBlockResultsEvent struct {
	Type       string `json:"type"`
	Attributes []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"attributes"`
}

type RawBlockResultsValidator struct {
	PubKey struct {
		Sum struct {
			Type string `json:"type"`
			Value struct {
				Ed25519 string `json:"ed25519"`
			} `json:"value"`
		} `json:"Sum"`
	} `json:"pub_key"`
	Power *string `json:"power"`
}

type RawConsensusParamUpdates struct {
	Block struct {
		MaxBytes string `json:"max_bytes"`
		MaxGas   string `json:"max_gas"`
	} `json:"block"`
	Evidence struct {
		MaxAgeNumBlocks string `json:"max_age_num_blocks"`
		MaxAgeDuration  string `json:"max_age_duration"`
	} `json:"evidence"`
	Validator struct {
		PubKeyTypes []string `json:"pub_key_types"`
	} `json:"validator"`
}
