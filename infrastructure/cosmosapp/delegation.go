package cosmosapp

import cosmosapp_interface "github.com/crypto-com/chain-indexing/appinterface/cosmosapp"

type DelegationResp struct {
	DelegationResponses []cosmosapp_interface.DelegationResponse `json:"delegation_responses"`
	Pagination          Pagination                               `json:"pagination"`
}
