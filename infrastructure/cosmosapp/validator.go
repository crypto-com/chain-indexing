package cosmosapp

import cosmosapp_interface "github.com/crypto-com/chain-indexing/appinterface/cosmosapp"

type ValidatorResp struct {
	Validator cosmosapp_interface.Validator `json:"validator"`
}
