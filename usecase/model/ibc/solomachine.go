package ibc

type SoloMachineLightClientState struct {
	Type                     string `mapstructure:"@type" json:"@type"`
	Sequence                 uint64 `mapstructure:"sequence" json:"sequence"`
	IsFrozen                 bool   `mapstructure:"is_frozen" json:"isFrozen"`
	AllowUpdateAfterProposal bool   `mapstructure:"allowUpdate_after_proposal" json:"allowUpdateAfterProposal"`
}

type SoloMachineLightClientConsensusState struct {
	PublicKey   SoloMachineLightClientConsensusStatePublicKey `mapstructure:"public_key" json:"publicKey"`
	Diversifier string                                        `mapstructure:"diversifier" json:"diversifier"`
	Timestamp   uint64                                        `mapstructure:"timestamp" json:"timestamp"`
}

type SoloMachineLightClientConsensusStatePublicKey struct {
	Type string `mapstructure:"@type" json:"@type"`
	Key  string `mapstructure:"key" json:"key"`
}
