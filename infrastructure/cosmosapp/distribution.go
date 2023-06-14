package cosmosapp

// A simplified distribution params response data schema
type DistributionParamsResp struct {
	Params DistributionParams `json:"params"`
}

type DistributionParams struct {
	CommunityTax string `json:"community_tax"`
}
