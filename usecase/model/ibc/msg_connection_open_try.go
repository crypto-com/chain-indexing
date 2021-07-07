package ibc

type MsgConnectionOpenTryParams struct {
	MsgConnectionOpenTryBaseParams

	MaybeTendermintClientState *TendermintLightClientState `json:"maybeTendermintClientState"`

	ConnectionID string `json:"connectionId"`
}

type RawMsgConnectionOpenTryTendermintClient struct {
	MsgConnectionOpenTryBaseParams `mapstructure:",squash"`

	TendermintClientState TendermintLightClientState `mapstructure:"client_state" json:"clientState"`
}

type MsgConnectionOpenTryBaseParams struct {
	ClientID             string                          `mapstructure:"client_id" json:"clientId"`
	PreviousConnectionID string                          `mapstructure:"previous_connection_id" json:"previousConnectionId"`
	Counterparty         ConnectionCounterparty          `mapstructure:"counterparty" json:"counterparty"`
	DelayPeriod          uint64                          `mapstructure:"delay_period" json:"delayPeriod,string"`
	CounterpartyVersions []ConnectionCounterpartyVersion `mapstructure:"counterparty_versions" json:"counterpartyVersions"`
	ProofHeight          Height                          `mapstructure:"proof_height" json:"proofHeight"`
	ProofInit            string                          `mapstructure:"proof_init" json:"proofInit"`
	ProofClient          string                          `mapstructure:"proof_client" json:"proofClient"`
	ProofConsensus       string                          `mapstructure:"proof_consensus" json:"proofConsensus"`
	ConsensusHeight      Height                          `mapstructure:"consensus_height" json:"consensusHeight"`
	Signer               string                          `mapstructure:"signer" json:"signer"`
}
