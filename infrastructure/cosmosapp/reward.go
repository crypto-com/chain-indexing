package cosmosapp

type DelegatorRewardResp struct {
	Rewards []DelegatorReward `json:"rewards"`
	Total   []DelegatorTotal  `json:"total"`
}

type DelegatorReward struct {
	ValidatorAddress string           `json:"validator_address"`
	Reward           []DelegatorTotal `json:"reward"`
}

type DelegatorTotal struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}
