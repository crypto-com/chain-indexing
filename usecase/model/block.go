package model

import (
	"time"

	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/luci/go-render/render"
)

type Block struct {
	Height          int64            `json:"height" fake:"{blockheight}"`
	Hash            string           `json:"hash" fake:"{blockhash}"`
	Time            utctime.UTCTime  `json:"time" fake:"{utctime}"`
	AppHash         string           `json:"appHash" fake:"{blockapphash}"`
	ProposerAddress string           `json:"proposerAddress" fake:"{validatoraddress}"`
	Txs             []string         `json:"txs" fake:"skip"`
	Signatures      []BlockSignature `json:"signature" fakesize:"3"`
	Evidences       []BlockEvidence  `json:"evidences" fake:"skip"`
}

func (block *Block) String() string {
	return render.Render(block)
}

type BlockSignature struct {
	BlockIdFlag      int             `json:"blockIdFlag" fake:"+int"`
	ValidatorAddress string          `json:"validatorAddress" fake:"{validatoraddress}"`
	Timestamp        utctime.UTCTime `json:"timestamp" fake:"{utctime}"`
	Signature        string          `json:"signature" fake:"{commitsignature}"`
}

type BlockEvidence struct {
	Type  string `json:"type"`
	Value struct {
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
			ValidatorIndex   int       `json:"validator_index"`
			Signature        string    `json:"signature"`
		} `json:"vote_a"`
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
			ValidatorIndex   int       `json:"validator_index"`
			Signature        string    `json:"signature"`
		} `json:"vote_b"`
		// TODO: Breaking changes to snake case in the future
		TotalVotingPower string    `json:"TotalVotingPower"`
		ValidatorPower   string    `json:"ValidatorPower"`
		Timestamp        time.Time `json:"Timestamp"`
	} `json:"value"`
}
