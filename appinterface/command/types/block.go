// Current file defines all block command types' Payload
package types

import "time"

// BlockSignature defines the parsed signatures
type BlockSignature struct {
	BlockIDFlag      int
	ValidatorAddress string
	Timestamp        time.Time
	Signature        string
}

// Block defines CreateBlock cmd payload
type Block struct {
	Height          uint64
	Hash            string
	Time            time.Time
	AppHash         string
	ProposerAddress string
	Txs             []string
	Signatures      []BlockSignature
}
