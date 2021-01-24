package cosmosapp

type UnbondingResp struct {
	UnbondingResponses []UnbondingResponse `json:"unbonding_responses"`
	Pagination         Pagination          `json:"pagination"`
}

type UnbondingResponse struct {
	DelegatorAddress string           `json:"delegator_address"`
	ValidatorAddress string           `json:"validator_address"`
	Entries          []UnbondingEntry `json:"entries"`
}

type UnbondingEntry struct {
	CreationHeight string `json:"creation_height"`
	CompletionTime string `json:"completion_time"`
	InitialBalance string `json:"initial_balance"`
	Balance        string `json:"balance"`
}
