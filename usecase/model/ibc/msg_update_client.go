package ibc

import "time"

type MsgUpdateClientParams struct {
	MaybeTendermintLightClientUpdate  *TendermintLightClientUpdate  `json:"maybeTendermintLightClientUpdate"`
	MaybeSoloMachineLightClientUpdate *SoloMachineLightClientUpdate `json:"maybeSoloMachineLightClientUpdate"`
	// TODO: Localhost LightClient

	ClientID        string `json:"clientId"`
	ClientType      string `json:"clientType"`
	ConsensusHeight Height `json:"consensusHeight"`
	Signer          string `json:"signer"`
}

type TendermintLightClientUpdate struct {
	Header TendermintLightClientHeader `json:"header"`
}

type SoloMachineLightClientUpdate struct {
	Header SoloMachineLightClientHeader `json:"header"`
}

type RawMsgUpdateTendermintLightClient struct {
	Type     string                      `mapstructure:"@type" json:"@type"`
	ClientID string                      `mapstructure:"client_id" json:"clientId"`
	Header   TendermintLightClientHeader `mapstructure:"header" json:"header"`
	Signer   string                      `mapstructure:"signer" json:"signer"`
}

type RawMsgUpdateSoloMachineLightClient struct {
	Type     string                       `mapstructure:"@type" json:"@type"`
	ClientID string                       `mapstructure:"client_id" json:"clientId"`
	Header   SoloMachineLightClientHeader `mapstructure:"header" json:"header"`
	Signer   string                       `mapstructure:"signer" json:"signer"`
}

type TendermintLightClientHeader struct {
	Type              string                                 `mapstructure:"@type" json:"@type"`
	SignedHeader      TendermintLightClientSignedHeader      `mapstructure:"signed_header" json:"signedHeader"`
	ValidatorSet      TendermintLightClientTrustedValidators `mapstructure:"validator_set" json:"validatorSet"`
	TrustedHeight     TendermintLightClientTrustedHeight     `mapstructure:"trusted_height" json:"trustedHeight"`
	TrustedValidators TendermintLightClientTrustedValidators `mapstructure:"trusted_validators" json:"trustedValidators"`
}

type TendermintLightClientSignedHeader struct {
	Header TendermintLightClientSignedHeaderHeader `mapstructure:"header" json:"header"`
	Commit TendermintLightClientCommit             `mapstructure:"commit" json:"commit"`
}

type TendermintLightClientCommit struct {
	Height     int64                            `mapstructure:"height" json:"height,string"`
	Round      int32                            `mapstructure:"round" json:"round"`
	BlockID    TendermintLightClientBlockID     `mapstructure:"block_id" json:"blockId"`
	Signatures []TendermintLightClientSignature `mapstructure:"signatures" json:"signatures"`
}

type TendermintLightClientBlockID struct {
	Hash          []byte                             `mapstructure:"hash" json:"hash"`
	PartSetHeader TendermintLightClientPartSetHeader `mapstructure:"part_set_header" json:"partSetHeader"`
}

type TendermintLightClientPartSetHeader struct {
	Total uint32 `mapstructure:"total" json:"total"`
	Hash  []byte `mapstructure:"hash" json:"hash"`
}

type TendermintLightClientSignature struct {
	BlockIDFlag      string    `mapstructure:"block_id_flag" json:"blockIdFlag"`
	ValidatorAddress []byte    `mapstructure:"validator_address" json:"validatorAddress"`
	Timestamp        time.Time `mapstructure:"timestamp" json:"timestamp"`
	Signature        []byte    `mapstructure:"signature" json:"signature"`
}

type TendermintLightClientSignedHeaderHeader struct {
	Version            TendermintLightClientVersion `mapstructure:"version" json:"version"`
	ChainID            string                       `mapstructure:"chain_id" json:"chainId"`
	Height             int64                        `mapstructure:"height" json:"height,string"`
	Time               time.Time                    `mapstructure:"time" json:"time"`
	LastBlockID        TendermintLightClientBlockID `mapstructure:"last_block_id" json:"lastBlockId"`
	LastCommitHash     []byte                       `mapstructure:"last_commit_hash" json:"lastCommitHash"`
	DataHash           []byte                       `mapstructure:"data_hash" json:"dataHash"`
	ValidatorsHash     []byte                       `mapstructure:"validators_hash" json:"validatorsHash"`
	NextValidatorsHash []byte                       `mapstructure:"next_validators_hash" json:"nextValidatorsHash"`
	ConsensusHash      []byte                       `mapstructure:"consensus_hash" json:"consensusHash"`
	AppHash            []byte                       `mapstructure:"app_hash" json:"appHash"`
	LastResultsHash    []byte                       `mapstructure:"last_results_hash" json:"lastResultsHash"`
	EvidenceHash       []byte                       `mapstructure:"evidence_hash" json:"evidenceHash"`
	ProposerAddress    []byte                       `mapstructure:"proposer_address" json:"proposerAddress"`
}

type TendermintLightClientVersion struct {
	Block uint64 `mapstructure:"block" json:"block,string"`
	App   string `mapstructure:"app" json:"app"`
}

type TendermintLightClientTrustedHeight struct {
	RevisionNumber uint64 `mapstructure:"revision_number" json:"revisionNumber,string"`
	RevisionHeight uint64 `mapstructure:"revision_height" json:"revisionHeight,string"`
}

type TendermintLightClientTrustedValidators struct {
	Validators       []TendermintLightClientProposer `mapstructure:"validators" json:"validators"`
	Proposer         TendermintLightClientProposer   `mapstructure:"proposer" json:"proposer"`
	TotalVotingPower string                          `mapstructure:"total_voting_power" json:"totalVotingPower"`
}

type TendermintLightClientProposer struct {
	Address          []byte                      `mapstructure:"address" json:"address"`
	PubKey           TendermintLightClientPubKey `mapstructure:"pub_key" json:"pubKey"`
	VotingPower      int64                       `mapstructure:"voting_power" json:"votingPower,string"`
	ProposerPriority int64                       `mapstructure:"proposer_priority" json:"proposerPriority,string"`
}

type TendermintLightClientPubKey struct {
	MaybeEd25519   []byte `mapstructure:"ed25519" json:"ed25519,omitempty"`
	MaybeSecp256K1 []byte `mapstructure:"secp256k1" json:"secp256k1,omitempty"`
}

type SoloMachineLightClientHeader struct {
	Type           string                       `mapstructure:"@type" json:"@type"`
	Sequence       uint64                       `mapstructure:"sequence" json:"sequence"`
	Timestamp      uint64                       `mapstructure:"timestamp" json:"timestamp"`
	Signature      []byte                       `mapstructure:"signature" json:"signature"`
	NewPublicKey   SoloMachineLightClientPubKey `mapstructure:"new_public_key" json:"newPublicKey"`
	NewDiversifier string                       `mapstructure:"new_diversifier" json:"NewDiversifier"`
}

type SoloMachineLightClientPubKey struct {
	Type string `mapstructure:"@type" json:"@type"`
	Key  string `mapstructure:"key" json:"key"`
}
