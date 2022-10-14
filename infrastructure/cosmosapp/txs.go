package cosmosapp

import (
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type TxsResp struct {
	Tx         model.CosmosTx `json:"tx"`
	TxResponse struct {
		Height    string `json:"height,omitempty"`
		TxHash    string `json:"txhash,omitempty"`
		Codespace string `json:"codespace,omitempty"`
		Code      uint32 `json:"code,omitempty"`
		Data      string `json:"data,omitempty"`
		RawLog    string `json:"raw_log,omitempty"`
		Logs      []struct {
			MsgIndex uint32             `json:"msg_index,omitempty"`
			Log      string             `json:"log,omitempty"`
			Events   []model.BlockEvent `json:"events"`
		} `json:"logs"`
		Info      string `json:"info,omitempty"`
		GasWanted string `json:"gas_wanted,omitempty"`
		GasUsed   string `json:"gas_used,omitempty"`
		Tx        struct {
			Type string `json:"@type"`
			model.CosmosTx
		} `json:"tx,omitempty"`
		Timestamp string             `json:"timestamp,omitempty"`
		Events    []model.BlockEvent `json:"events"`
	} `json:"tx_response"`
}

type TxsByHeightResp struct {
	Txs     []model.CosmosTx `json:"txs"`
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
			Evidence []model.BlockEvidence `json:"evidence"`
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
			Signatures []struct {
				BlockIDFlag      string          `json:"block_id_flag"`
				ValidatorAddress string          `json:"validator_address"`
				Timestamp        utctime.UTCTime `json:"timestamp"`
				Signature        *string         `json:"signature"`
			} `json:"signatures"`
		} `json:"last_commit"`
	} `json:"block"`
}
