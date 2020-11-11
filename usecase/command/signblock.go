package command

import (
	entity_event "github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/internal/utctime"
	"github.com/crypto-com/chainindex/usecase/event/blocksigned"
)

type SignBlock struct {
	blockHeight      int64
	validatorAddress string
	isProposer       bool
	timestamp        utctime.UTCTime
	signature        string
}

func NewSignBlock(blockHeight int64, validatorAddress string, isProposer bool, timestamp utctime.UTCTime, signature string) *SignBlock {
	return &SignBlock{
		blockHeight,

		validatorAddress,
		isProposer,
		timestamp,
		signature,
	}
}

func (_ *SignBlock) Name() string {
	return "SignBlock"
}

func (_ *SignBlock) Version() int {
	return 0
}

func (cmd *SignBlock) Exec() (entity_event.Event, error) {
	evt := blocksigned.New(
		cmd.blockHeight,
		cmd.validatorAddress,
		cmd.isProposer,
		cmd.timestamp,
		cmd.signature,
	)
	return evt, nil
}
