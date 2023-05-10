package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

const VALIDATOR_BLOCK_COMMITMENTS_TABLE_NAME = "view_validator_block_commitments_total"

type ValidatorBlockCommitmentsTotal interface {
	Set(string, int64) error
	Increment(string, int64) error
	IncrementAll([]string, int64) error
	DecrementAll([]string, int64) error
	FindBy(string) (int64, error)
	SumBy([]string) (int64, error)
}

type ValidatorBlockCommitmentsTotalView struct {
	*view.Total
}

func NewValidatorBlockCommitmentsTotal(rdbHandle *rdb.Handle) ValidatorBlockCommitmentsTotal {
	return &ValidatorBlockCommitmentsTotalView{
		view.NewTotal(rdbHandle, VALIDATOR_BLOCK_COMMITMENTS_TABLE_NAME),
	}
}
