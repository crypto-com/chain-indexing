package ibc

type MsgConnectionOpenAckParams struct {
	MsgConnectionOpenAckBaseParams

	MaybeTendermintClientState *TendermintLightClientState `json:"maybeTendermintClientState"`

	ClientID             string `json:"clientId"`
	CounterpartyClientID string `json:"counterpartyClientId"`
}

type RawMsgConnectionOpenAckTendermintClient struct {
	MsgConnectionOpenAckBaseParams `mapstructure:",squash"`

	TendermintClientState TendermintLightClientState `mapstructure:"client_state" json:"clientState"`
}

type MsgConnectionOpenAckBaseParams struct {
	ConnectionID             string                        `mapstructure:"connection_id" json:"connectionId"`
	CounterpartyConnectionID string                        `mapstructure:"counterparty_connection_id" json:"counterpartyConnectionId"`
	Version                  ConnectionCounterpartyVersion `mapstructure:"version" json:"version"`
	ProofHeight              Height                        `mapstructure:"proof_height" json:"proofHeight"`
	ProofTry                 []byte                        `mapstructure:"proof_try" json:"proofTry"`
	ProofClient              []byte                        `mapstructure:"proof_client" json:"proofClient"`
	ProofConsensus           []byte                        `mapstructure:"proof_consensus" json:"proofConsensus"`
	ConsensusHeight          Height                        `mapstructure:"consensus_height" json:"consensusHeight"`
	Signer                   string                        `mapstructure:"signer" json:"signer"`
}
