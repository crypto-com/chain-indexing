package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

const VOTES_TOTAL_TABLE_NAME = "view_proposal_votes_total"

type VotesTotal struct {
	*view.Total
}

func NewVotesTotal(rdbHandle *rdb.Handle) *VotesTotal {
	return &VotesTotal{
		view.NewTotal(rdbHandle, VOTES_TOTAL_TABLE_NAME),
	}
}
