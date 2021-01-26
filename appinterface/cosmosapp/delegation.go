package cosmosapp

type DelegationResponse struct {
	Delegation Delegation        `json:"delegation"`
	Balance    DelegationBalance `json:"balance"`
}

type Delegation struct {
	DelegatorAddress string `json:"delegator_address"`
	ValidatorAddress string `json:"validator_address"`
	Shares           string `json:"shares"`
}

type DelegationBalance struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}
