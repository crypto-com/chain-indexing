package model

import (
	"github.com/crypto-com/chainindex/internal/utctime"
)

type Block struct {
	Height          int64            `json:"height" fake:"{blockheight}"`
	Hash            string           `json:"hash" fake:"{blockhash}"`
	Time            utctime.UTCTime  `json:"time" fake:"{utctime}"`
	AppHash         string           `json:"appHash" fake:"{blockapphash}"`
	ProposerAddress string           `json:"proposerAddress" fake:"{validatoraddress}"`
	Txs             []string         `json:"txs" fake:"skip"`
	Signatures      []BlockSignature `json:"signature" fakesize:"3"`
}

type BlockSignature struct {
	BlockIdFlag      int             `json:"blockIdFlag" fake:"+int"`
	ValidatorAddress string          `json:"validatorAddress" fake:"{validatoraddress}"`
	Timestamp        utctime.UTCTime `json:"timestamp" fake:"{utctime}"`
	Signature        string          `json:"signature" fake:"{commitsignature}"`
}
