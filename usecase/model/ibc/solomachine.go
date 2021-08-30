package ibc

type SoloMachineLightClientState struct {
	Type                     string                               `mapstructure:"@type" json:"@type"`
	Sequence                 uint64                               `mapstructure:"sequence" json:"sequence"`
	IsFrozen                 bool                                 `mapstructure:"isFrozen" json:"isFrozen"`
	ConsensusState           SoloMachineLightClientConsensusState `mapstructure:"consensusState" json:"consensusState"`
	AllowUpdateAfterProposal bool                                 `mapstructure:"allowUpdateAfterProposal" json:"allowUpdateAfterProposal"`
}

type SoloMachineLightClientConsensusState struct {
	PublicKey   SoloMachineLightClientConsensusStatePublicKey `mapstructure:"publicKey" json:"publicKey"`
	Diversifier string                                        `mapstructure:"diversifier" json:"diversifier"`
	Timestamp   uint64                                        `mapstructure:"timestamp" json:"timestamp"`
}

type SoloMachineLightClientConsensusStatePublicKey struct {
	Type string `mapstructure:"type" json:"type"`
	Key  string `mapstructure:"key" json:"key"`
}
