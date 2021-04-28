package cosmosapp

import cosmosapp_interface "github.com/crypto-com/chain-indexing/appinterface/cosmosapp"

type ProposalsResp struct {
	MaybeProposalsResponse []cosmosapp_interface.Proposal `json:"proposals"`
	MaybePagination        *Pagination                    `json:"pagination"`
	// On error
	MaybeCode    *int    `json:"code"`
	MaybeMessage *string `json:"message"`
}

type ProposalResp struct {
	Proposal cosmosapp_interface.Proposal `json:"proposal"`
}

type TallyResp struct {
	Tally cosmosapp_interface.Tally `json:"tally"`
}
