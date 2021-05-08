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
	SignedHeader      TendermintLightClientUpdateSignedHeader      `json:"signedHeader"`
	ValidatorSet      TendermintLightClientUpdateTrustedValidators `json:"validatorSet"`
	TrustedHeight     TendermintLightClientUpdateTrustedHeight     `json:"trustedHeight"`
	TrustedValidators TendermintLightClientUpdateTrustedValidators `json:"trustedValidators"`
}

type TendermintLightClientUpdateSignedHeader struct {
	Header TendermintLightClientUpdateSignedHeaderHeader `json:"header"`
	Commit TendermintLightClientUpdateCommit             `json:"commit"`
}

type TendermintLightClientUpdateCommit struct {
	Height     string                                 `json:"height"`
	Round      int64                                  `json:"round"`
	BlockID    TendermintLightClientUpdateBlockID     `json:"blockId"`
	Signatures []TendermintLightClientUpdateSignature `json:"signatures"`
}

type TendermintLightClientUpdateBlockID struct {
	Hash          string                                   `json:"hash"`
	PartSetHeader TendermintLightClientUpdatePartSetHeader `json:"partSetHeader"`
}

type TendermintLightClientUpdatePartSetHeader struct {
	Total int64  `json:"total"`
	Hash  string `json:"hash"`
}

type TendermintLightClientUpdateSignature struct {
	BlockIDFlag      string `json:"blockIdFlag"`
	ValidatorAddress string `json:"validatorAddress"`
	Timestamp        string `json:"timestamp"`
	Signature        string `json:"signature"`
}

type TendermintLightClientUpdateSignedHeaderHeader struct {
	Version            TendermintLightClientUpdateVersion `json:"version"`
	ChainID            string                             `json:"chainId"`
	Height             string                             `json:"height"`
	Time               string                             `json:"time"`
	LastBlockID        TendermintLightClientUpdateBlockID `json:"lastBlockId"`
	LastCommitHash     string                             `json:"lastCommitHash"`
	DataHash           string                             `json:"dataHash"`
	ValidatorsHash     string                             `json:"validatorsHash"`
	NextValidatorsHash string                             `json:"nextValidatorsHash"`
	ConsensusHash      string                             `json:"consensusHash"`
	AppHash            string                             `json:"appHash"`
	LastResultsHash    string                             `json:"lastResultsHash"`
	EvidenceHash       string                             `json:"evidenceHash"`
	ProposerAddress    string                             `json:"proposerAddress"`
}

type TendermintLightClientUpdateVersion struct {
	Block string `json:"block"`
	App   string `json:"app"`
}

type TendermintLightClientUpdateTrustedHeight struct {
	RevisionNumber string `json:"revisionNumber"`
	RevisionHeight string `json:"revisionHeight"`
}

type TendermintLightClientUpdateTrustedValidators struct {
	Validators       []TendermintLightClientUpdateProposer `json:"validators"`
	Proposer         TendermintLightClientUpdateProposer   `json:"proposer"`
	TotalVotingPower string                                `json:"totalVotingPower"`
}

type TendermintLightClientUpdateProposer struct {
	Address          string                            `json:"address"`
	PubKey           TendermintLightClientUpdatePubKey `json:"pubKey"`
	VotingPower      string                            `json:"votingPower"`
	ProposerPriority string                            `json:"proposerPriority"`
}

type TendermintLightClientUpdatePubKey struct {
	Ed25519 string `json:"ed25519"`
}
