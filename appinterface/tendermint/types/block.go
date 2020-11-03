package types

import (
	"github.com/crypto-com/chainindex/internal/utctime"
)

// BlockSignature defines the parsed signatures
type BlockSignature struct {
	BlockIDFlag      int
	ValidatorAddress string
	Timestamp        utctime.UTCTime
	Signature        string
}

// Block defines parsed block data
type Block struct {
	Height          int64
	Hash            string
	Time            utctime.UTCTime
	AppHash         string
	ProposerAddress string
	Txs             []string
	Signatures      []BlockSignature
}
