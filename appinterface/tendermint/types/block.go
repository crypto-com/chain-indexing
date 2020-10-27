package types

import "time"

// BlockResp defines the structure for tendermint /block API response JSON
type BlockResp struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		BlockID struct {
			Hash  string `json:"hash"`
			Parts struct {
				Total int    `json:"total"`
				Hash  string `json:"hash"`
			} `json:"parts"`
		} `json:"block_id"`
		Block struct {
			Header struct {
				Version struct {
					Block string `json:"block"`
				} `json:"version"`
				ChainID     string    `json:"chain_id"`
				Height      string    `json:"height"`
				Time        time.Time `json:"time"`
				LastBlockID struct {
					Hash  string `json:"hash"`
					Parts struct {
						Total int    `json:"total"`
						Hash  string `json:"hash"`
					} `json:"parts"`
				} `json:"last_block_id"`
				LastCommitHash     string `json:"last_commit_hash"`
				DataHash           string `json:"data_hash"`
				ValidatorsHash     string `json:"validators_hash"`
				NextValidatorsHash string `json:"next_validators_hash"`
				ConsensusHash      string `json:"consensus_hash"`
				AppHash            string `json:"app_hash"`
				LastResultsHash    string `json:"last_results_hash"`
				EvidenceHash       string `json:"evidence_hash"`
				ProposerAddress    string `json:"proposer_address"`
			} `json:"header"`
			Data struct {
				Txs []string `json:"txs"`
			} `json:"data"`
			Evidence struct {
				Evidence []struct {
					Type  string `json:"type"`
					Value struct {
						PubKey struct {
							Type  string `json:"type"`
							Value string `json:"value"`
						} `json:"PubKey"`
						VoteA struct {
							Type    int    `json:"type"`
							Height  string `json:"height"`
							Round   int    `json:"round"`
							BlockID struct {
								Hash  string `json:"hash"`
								Parts struct {
									Total int    `json:"total"`
									Hash  string `json:"hash"`
								} `json:"parts"`
							} `json:"block_id"`
							Timestamp        time.Time `json:"timestamp"`
							ValidatorAddress string    `json:"validator_address"`
							ValidatorIndex   string    `json:"validator_index"`
							Signature        string    `json:"signature"`
						} `json:"VoteA"`
						VoteB struct {
							Type    int    `json:"type"`
							Height  string `json:"height"`
							Round   int    `json:"round"`
							BlockID struct {
								Hash  string `json:"hash"`
								Parts struct {
									Total int    `json:"total"`
									Hash  string `json:"hash"`
								} `json:"parts"`
							} `json:"block_id"`
							Timestamp        time.Time `json:"timestamp"`
							ValidatorAddress string    `json:"validator_address"`
							ValidatorIndex   string    `json:"validator_index"`
							Signature        string    `json:"signature"`
						} `json:"VoteB"`
					} `json:"value"`
				} `json:"evidence"`
			} `json:"evidence"`
			LastCommit struct {
				Height  string `json:"height"`
				Round   int    `json:"round"`
				BlockID struct {
					Hash  string `json:"hash"`
					Parts struct {
						Total int    `json:"total"`
						Hash  string `json:"hash"`
					} `json:"parts"`
				} `json:"block_id"`
				Signatures []RawBlockSignature `json:"signatures"`
			} `json:"last_commit"`
		} `json:"block"`
	} `json:"result"`
}

// RawBlockSignature defines the structure for signatures in /block API
type RawBlockSignature struct {
	BlockIDFlag      int       `json:"block_id_flag"`
	ValidatorAddress string    `json:"validator_address"`
	Timestamp        time.Time `json:"timestamp"`
	Signature        *string   `json:"signature"`
}

// BlockSignature defines the parsed signatures
type BlockSignature struct {
	BlockIDFlag      int
	ValidatorAddress string
	Timestamp        time.Time
	Signature        string
}

// Block defines parsed block data
type Block struct {
	Height          uint64
	Hash            string
	Time            time.Time
	AppHash         string
	ProposerAddress string
	Txs             []string
	Signatures      []BlockSignature
}
