package model

import (
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/luci/go-render/render"
)

// RawBlock defines the structure for TendermintApp /block API response JSON
type RawBlock struct {
	BlockID struct {
		Hash  string `json:"hash" fake:"{blockheight}"`
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
			ChainID     string          `json:"chain_id"`
			Height      string          `json:"height"`
			Time        utctime.UTCTime `json:"time"`
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
			Evidence []BlockEvidence `json:"evidence"`
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
}

func (rawBlock *RawBlock) String() string {
	return render.Render(rawBlock)
}

// RawBlockSignature defines the structure for signatures in /block API
type RawBlockSignature struct {
	BlockIDFlag      int             `json:"block_id_flag"`
	ValidatorAddress string          `json:"validator_address"`
	Timestamp        utctime.UTCTime `json:"timestamp"`
	Signature        *string         `json:"signature"`
}
