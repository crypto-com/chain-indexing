package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

const VOTES_TOTAL_TABLE_NAME = "view_proposal_votes_total"

type VotesTotal interface {
	Set(string, int64) error
	Increment(string, int64) error
	IncrementAll([]string, int64) error
	DecrementAll([]string, int64) error
	FindBy(string) (int64, error)
	SumBy([]string) (int64, error)
}

type VotesTotalView struct {
	*view.Total
}

func NewVotesTotalView(rdbHandle *rdb.Handle) VotesTotal {
	return &VotesTotalView{
		view.NewTotal(rdbHandle, VOTES_TOTAL_TABLE_NAME),
	}
}
