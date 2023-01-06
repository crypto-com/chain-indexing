package v1_model

type MsgVoteParams struct {
	ProposalId string `json:"proposalId"`
	Voter      string `json:"voter"`
	Option     string `json:"option"`
	Metadata   string `json:"metadata"`
}
type MsgVoteWeightedParams struct {
	ProposalId  string       `json:"proposalId"`
	Voter       string       `json:"voter"`
	VoteOptions []VoteOption `json:"vote_options"`
	Metadata    string       `json:"metadata"`
}
type RawMsgVoteWeight struct {
	ProposalId string       `json:"proposalId"`
	Voter      string       `json:"voter"`
	Options    []VoteOption `json:"options"`
	Metadata   string       `json:"metadata"`
}

type VoteOption struct {
	Option string `json:"option"`
	Weight string `json:"weight"`
}
