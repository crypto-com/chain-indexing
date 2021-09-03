package ibc

type MsgCreateClientParams struct {
	MaybeTendermintLightClient  *TendermintLightClient  `json:"maybeTendermintLightClient"`
	MaybeSoloMachineLightClient *SoloMachineLightClient `json:"maybeSoloMachineLightClient"`
	MaybeLocalhostLightClient   *LocalhostLightClient   `json:"maybeLocalhostLightClient"`
	// TODO: Test localhost LightClient
	Signer string `json:"signer"`

	ClientID   string `json:"clientId"`
	ClientType string `json:"clientType"`
}

type TendermintLightClient struct {
	TendermintClientState               TendermintLightClientState          `json:"clientState"`
	TendermintLightClientConsensusState TendermintLightClientConsensusState `json:"consensusState"`
}

type SoloMachineLightClient struct {
	SoloMachineClientState               SoloMachineLightClientState          `json:"clientState"`
	SoloMachineLightClientConsensusState SoloMachineLightClientConsensusState `json:"consensusState"`
}

type LocalhostLightClient struct {
	LocalhostClientState LocalhostLightClientState `json:"clientState"`
}

type RawMsgCreateTendermintLightClient struct {
	Type           string                              `mapstructure:"@type" json:"@type"`
	ClientState    TendermintLightClientState          `mapstructure:"client_state" json:"clientState"`
	ConsensusState TendermintLightClientConsensusState `mapstructure:"consensus_state" json:"consensusState"`
	Signer         string                              `mapstructure:"signer" json:"signer"`
}

type RawMsgCreateSoloMachineLightClient struct {
	Type           string                               `mapstructure:"@type" json:"@type"`
	ClientState    SoloMachineLightClientState          `mapstructure:"client_state" json:"clientState"`
	ConsensusState SoloMachineLightClientConsensusState `mapstructure:"consensus_state" json:"consensusState"`
	Signer         string                               `mapstructure:"signer" json:"signer"`
}

type RawMsgCreateLocalhostLightClient struct {
	Type        string                    `mapstructure:"@type" json:"@type"`
	ClientState LocalhostLightClientState `mapstructure:"client_state" json:"clientState"`
	Signer      string                    `mapstructure:"signer" json:"signer"`
}
