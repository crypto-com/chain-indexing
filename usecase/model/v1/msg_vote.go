package v1_model

type MsgVoteParams struct {
	ProposalId string `json:"proposalId"`
	Voter      string `json:"voter"`
	Option     string `json:"option"`
	Metadata   string `json:"metadata"`
}
