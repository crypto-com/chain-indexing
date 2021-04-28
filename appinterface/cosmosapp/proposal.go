package cosmosapp

type Proposal struct {
	ProposalID       string           `json:"proposal_id"`
	Content          Content          `json:"content"`
	Status           string           `json:"status"`
	FinalTallyResult FinalTallyResult `json:"final_tally_result"`
	SubmitTime       string           `json:"submit_time"`
	DepositEndTime   string           `json:"deposit_end_time"`
	TotalDeposit     []TotalDeposit   `json:"total_deposit"`
	VotingStartTime  string           `json:"voting_start_time"`
	VotingEndTime    string           `json:"voting_end_time"`
}

type Content struct {
	Type        string   `json:"@type"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Changes     []Change `json:"changes"`
}

type Change struct {
	Subspace string `json:"subspace"`
	Key      string `json:"key"`
	Value    string `json:"value"`
}

type FinalTallyResult struct {
	Yes        string `json:"yes"`
	Abstain    string `json:"abstain"`
	No         string `json:"no"`
	NoWithVeto string `json:"no_with_veto"`
}

type TotalDeposit struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type Tally struct {
	Yes        string `json:"yes"`
	Abstain    string `json:"abstain"`
	No         string `json:"no"`
	NoWithVeto string `json:"no_with_veto"`
}
