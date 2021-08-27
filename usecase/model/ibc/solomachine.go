package ibc

type SoloMachineLightClientState struct {
	Sequence                 uint64
	IsFrozen                 bool
	ConsensusState           SoloMachineLightClientConsensusState
	AllowUpdateAfterProposal bool
}

type SoloMachineLightClientConsensusState struct {
	PublicKey   string
	Diversifier string
	Timestamp   uint64
}
