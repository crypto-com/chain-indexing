package ibc

type MsgUpdateClientParams struct {
	MaybeTendermintLightClientUpdate *TendermintLightClientUpdate `json:"maybeTendermintLightClientUpdate"`
	// TODO: SoloMachine and Localhost LightClient
	ClientID string `json:"clientId"`
	Signer   string `json:"signer"`
}

type TendermintLightClientUpdate struct {
	Header TendermintLightClientUpdateHeader `json:"header"`
}

type TendermintLightClientUpdateHeader struct {
	Type              string                                       `json:"@type"`
	SignedHeader      TendermintLightClientUpdateSignedHeader      `json:"signed_header"`
	ValidatorSet      TendermintLightClientUpdateTrustedValidators `json:"validator_set"`
	TrustedHeight     TendermintLightClientUpdateTrustedHeight     `json:"trusted_height"`
	TrustedValidators TendermintLightClientUpdateTrustedValidators `json:"trusted_validators"`
}

type TendermintLightClientUpdateSignedHeader struct {
	Header TendermintLightClientUpdateSignedHeaderHeader `json:"header"`
	Commit TendermintLightClientUpdateCommit             `json:"commit"`
}

type TendermintLightClientUpdateCommit struct {
	Height     string                                 `json:"height"`
	Round      int64                                  `json:"round"`
	BlockID    TendermintLightClientUpdateBlockID     `json:"block_id"`
	Signatures []TendermintLightClientUpdateSignature `json:"signatures"`
}

type TendermintLightClientUpdateBlockID struct {
	Hash          string                                   `json:"hash"`
	PartSetHeader TendermintLightClientUpdatePartSetHeader `json:"part_set_header"`
}

type TendermintLightClientUpdatePartSetHeader struct {
	Total int64  `json:"total"`
	Hash  string `json:"hash"`
}

type TendermintLightClientUpdateSignature struct {
	BlockIDFlag      string `json:"block_id_flag"`
	ValidatorAddress string `json:"validator_address"`
	Timestamp        string `json:"timestamp"`
	Signature        string `json:"signature"`
}

type TendermintLightClientUpdateSignedHeaderHeader struct {
	Version            TendermintLightClientUpdateVersion `json:"version"`
	ChainID            string                             `json:"chain_id"`
	Height             string                             `json:"height"`
	Time               string                             `json:"time"`
	LastBlockID        TendermintLightClientUpdateBlockID `json:"last_block_id"`
	LastCommitHash     string                             `json:"last_commit_hash"`
	DataHash           string                             `json:"dataHash"`
	ValidatorsHash     string                             `json:"validators_hash"`
	NextValidatorsHash string                             `json:"next_validators_hash"`
	ConsensusHash      string                             `json:"consensus_hash"`
	AppHash            string                             `json:"appHash"`
	LastResultsHash    string                             `json:"last_results_hash"`
	EvidenceHash       string                             `json:"evidence_hash"`
	ProposerAddress    string                             `json:"proposer_address"`
}

type TendermintLightClientUpdateVersion struct {
	Block string `json:"block"`
	App   string `json:"app"`
}

type TendermintLightClientUpdateTrustedHeight struct {
	RevisionNumber string `json:"revision_number"`
	RevisionHeight string `json:"revision_height"`
}

type TendermintLightClientUpdateTrustedValidators struct {
	Validators       []TendermintLightClientUpdateProposer `json:"validators"`
	Proposer         TendermintLightClientUpdateProposer   `json:"proposer"`
	TotalVotingPower string                                `json:"total_voting_power"`
}

type TendermintLightClientUpdateProposer struct {
	Address          string                            `json:"address"`
	PubKey           TendermintLightClientUpdatePubKey `json:"pubKey"`
	VotingPower      string                            `json:"voting_power"`
	ProposerPriority string                            `json:"proposer_priority"`
}

type TendermintLightClientUpdatePubKey struct {
	Ed25519 string `json:"ed25519"`
}
