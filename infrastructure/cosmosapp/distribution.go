package cosmosapp

import "github.com/crypto-com/chain-indexing/usecase/coin"

// A simplified distribution params response data schema
type DistributionParamsResp struct {
	Params DistributionParams `json:"params"`
}

type DistributionParams struct {
	CommunityTax string `json:"community_tax"`
}

type CommunityPoolResp struct {
	Pool coin.DecCoins `json:"pool"`
}
